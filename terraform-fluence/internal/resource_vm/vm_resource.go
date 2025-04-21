package resource_vm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudlesslabs/terraform-provider-fluence/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// vmResourceModel holds Terraform state for fluence_vm.
// It includes optional constraints, instances, vm_configuration, and computed id (the Fluence vmId).
type vmResourceModel struct {
	Constraints     ConstraintsValue     `tfsdk:"constraints"`
	Instances       types.Int64          `tfsdk:"instances"`
	VmConfiguration VmConfigurationValue `tfsdk:"vm_configuration"`
	// id is the Fluence vmId
	Id              types.String         `tfsdk:"id"`
}

// NewVmResource returns a new vmResource implementation.
func NewVmResource() resource.Resource {
	return &vmResource{}
}

// vmResource is the Terraform resource for Fluence VMs.
type vmResource struct {
	cfg client.Config
}

// Configure assigns the provider data (client.Config) to the resource.
func (r *vmResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if cfg, ok := req.ProviderData.(client.Config); ok {
		r.cfg = cfg
	}
}

// Metadata sets the resource type name.
func (r *vmResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vm"
}

// Schema merges the generated schema and injects a computed "id" attribute.
func (r *vmResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	base := VmResourceSchema(ctx)
	// Computed vmId
	base.Attributes["id"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The Fluence vmId for this VM instance.",
	}
	resp.Schema = base
}

// Create provisions a new VM instance.
func (r *vmResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan vmResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...) // read inputs
	if resp.Diagnostics.HasError() {
		return
	}

	// Build API payload
	instances := plan.Instances.ValueInt64()
	cfg := plan.VmConfiguration

	// openPorts
	var openPorts []map[string]interface{}
	if !cfg.OpenPorts.IsNull() && !cfg.OpenPorts.IsUnknown() {
		for _, e := range cfg.OpenPorts.Elements() {
			op := e.(OpenPortsValue)
			openPorts = append(openPorts, map[string]interface{}{ 
				"port":     op.Port.ValueInt64(),
				"protocol": op.Protocol.ValueString(),
			})
		}
	}

	// sshKeys
	var sshKeys []string
	if !cfg.SshKeys.IsNull() && !cfg.SshKeys.IsUnknown() {
		for _, e := range cfg.SshKeys.Elements() {
			sshKeys = append(sshKeys, e.(types.String).ValueString())
		}
	}

	payload := map[string]interface{}{ 
		"instances": instances,
		"vmConfiguration": map[string]interface{}{ 
			"name":      cfg.Name.ValueString(),
			"osImage":   cfg.OsImage.ValueString(),
			"openPorts": openPorts,
			"sshKeys":   sshKeys,
		},
	}

	// Marshal request
	body, err := json.Marshal(payload)
	if err != nil {
		resp.Diagnostics.AddError("Marshal error", err.Error())
		return
	}

	reqURL := fmt.Sprintf("%s/vms/v3", r.cfg.APIURL)
	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(body))
	httpReq.Header.Set("Authorization", "Bearer "+r.cfg.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := r.cfg.HTTPClient.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("API error", err.Error())
		return
	}
	defer httpResp.Body.Close()

	// Accept 200 or 201
	if httpResp.StatusCode != http.StatusCreated && httpResp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(httpResp.Body)
		resp.Diagnostics.AddError("API error", fmt.Sprintf("status %d: %s", httpResp.StatusCode, string(b)))
		return
	}

	// Decode the API response to get vmId
	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		resp.Diagnostics.AddError("Read error", err.Error())
		return
	}

	var raw []map[string]interface{}
	if err := json.Unmarshal(respBody, &raw); err != nil {
		resp.Diagnostics.AddError("Decode error", fmt.Sprintf("unable to parse JSON: %s\nResponse body: %s", err.Error(), string(respBody)))
		return
	}
	if len(raw) == 0 {
		resp.Diagnostics.AddError("API error", "empty JSON array from VM creation")
		return
	}
	idVal, ok := raw[0]["vmId"].(string)
	if !ok || idVal == "" {
		resp.Diagnostics.AddError("API error", fmt.Sprintf("vmId field missing or empty in response: %v", raw[0]))
		return
	}
	plan.Id = types.StringValue(idVal)

	// Nullify unknown constraints
	if plan.Constraints.IsUnknown() {
		plan.Constraints = NewConstraintsValueNull()
	}

	// Save full state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)  
}

// Read refreshes state via GET /vms/v2
func (r *vmResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state vmResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)  
	if resp.Diagnostics.HasError() {
		return
	}

	reqURL := fmt.Sprintf("%s/vms/v2", r.cfg.APIURL)
	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	httpReq.Header.Set("Authorization", "Bearer "+r.cfg.APIKey)

	httpResp, err := r.cfg.HTTPClient.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("API error", err.Error())
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(httpResp.Body)
		resp.Diagnostics.AddError("API error", fmt.Sprintf("status %d: %s", httpResp.StatusCode, string(b)))
		return
	}

	var list []struct{ VmId string `json:"vmId"` }
	if err := json.NewDecoder(httpResp.Body).Decode(&list); err != nil {
		resp.Diagnostics.AddError("Decode error", err.Error())
		return
	}

	for _, inst := range list {
		if inst.VmId == state.Id.ValueString() {
			resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)  
			return
		}
	}

	resp.State.RemoveResource(ctx)
}

// Update forces replacement by delete then create
func (r *vmResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var del resource.DeleteResponse
	r.Delete(ctx, resource.DeleteRequest{State: req.State}, &del)

	var cre resource.CreateResponse
	r.Create(ctx, resource.CreateRequest{Plan: req.Plan}, &cre)

	resp.State = cre.State
	resp.Diagnostics.Append(cre.Diagnostics...)
}

// Delete tears down the VM via DELETE /vms/v1 using dealId, as expected by the API
func (r *vmResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state vmResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)  
	if resp.Diagnostics.HasError() {
		return
	}

	// Build delete payload using dealId (API expects dealId for deletion)
	body, _ := json.Marshal(map[string]string{"dealId": state.Id.ValueString()})
	reqURL := fmt.Sprintf("%s/vms/v1", r.cfg.APIURL)
	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodDelete, reqURL, bytes.NewReader(body))
	httpReq.Header.Set("Authorization", "Bearer "+r.cfg.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := r.cfg.HTTPClient.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("delete failed: %s", err.Error()))
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK && httpResp.StatusCode != http.StatusNoContent {
		b, _ := io.ReadAll(httpResp.Body)
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("status %d: %s", httpResp.StatusCode, string(b)))
		return
	}

	// On successful deletion, remove resource from state
	resp.State.RemoveResource(ctx)
}
