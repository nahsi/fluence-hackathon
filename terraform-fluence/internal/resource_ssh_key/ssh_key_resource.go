package resource_ssh_key

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudlesslabs/terraform-provider-fluence/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// apiSshKey models the JSON returned by the Fluence API for an SSH key.
type apiSshKey struct {
	Active      bool   `json:"active"`
	Algorithm   string `json:"algorithm"`
	Comment     string `json:"comment"`
	CreatedAt   string `json:"createdAt"`
	Fingerprint string `json:"fingerprint"`
}

// NewSshKeyResource returns a new SSH key resource.
func NewSshKeyResource() resource.Resource {
	return &sshKeyResource{}
}

type sshKeyResource struct {
	cfg client.Config
}

// Configure stores the provider data.
func (r *sshKeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if cfg, ok := req.ProviderData.(client.Config); ok {
		r.cfg = cfg
	}
}

// Metadata sets the resource type name.
func (r *sshKeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_key"
}

// Schema points at the generated schema for SSH keys.
func (r *sshKeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = SshKeyResourceSchema(ctx)
}

// Create uploads a new SSH key and sets computed fields in state.
func (r *sshKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SshKeyModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)  // read name & public_key
	if resp.Diagnostics.HasError() {
		return
	}

	// Build API request
	payload := map[string]string{
		"name":      plan.Name.ValueString(),
		"publicKey": plan.PublicKey.ValueString(),
	}
	body, _ := json.Marshal(payload)

	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodPost, r.cfg.APIURL+"/ssh_keys", bytes.NewReader(body))
	httpReq.Header.Set("Authorization", "Bearer "+r.cfg.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+r.cfg.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := r.cfg.HTTPClient.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("API Error", err.Error())
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK && httpResp.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(httpResp.Body)
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("status %d: %s", httpResp.StatusCode, string(b)))
		return
	}

	// Decode API response (single object)
	var k apiSshKey
	if err := json.NewDecoder(httpResp.Body).Decode(&k); err != nil {
		resp.Diagnostics.AddError("Decode Error", err.Error())
		return
	}

	// Only overwrite computed attributes
	plan.Active = types.BoolValue(k.Active)
	plan.Algorithm = types.StringValue(k.Algorithm)
	plan.Comment = types.StringValue(k.Comment)
	plan.CreatedAt = types.StringValue(k.CreatedAt)
	plan.Fingerprint = types.StringValue(k.Fingerprint)

	// Persist the new state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...) 
}

// Read refreshes the state with the latest API data.
func (r *sshKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SshKeyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...) // load current state
	if resp.Diagnostics.HasError() {
		return
	}

	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, r.cfg.APIURL+"/ssh_keys", nil)
	httpReq.Header.Set("Authorization", "Bearer "+r.cfg.APIKey)

	httpResp, err := r.cfg.HTTPClient.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("API Error", err.Error())
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(httpResp.Body)
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("status %d: %s", httpResp.StatusCode, string(b)))
		return
	}

	var apiKeys []apiSshKey
	if err := json.NewDecoder(httpResp.Body).Decode(&apiKeys); err != nil {
		resp.Diagnostics.AddError("Decode Error", err.Error())
		return
	}

	// Match the key by fingerprint and update computed fields
	for _, k := range apiKeys {
		if k.Fingerprint == state.Fingerprint.ValueString() {
			state.Active = types.BoolValue(k.Active)
			state.Algorithm = types.StringValue(k.Algorithm)
			state.Comment = types.StringValue(k.Comment)
			state.CreatedAt = types.StringValue(k.CreatedAt)
			resp.Diagnostics.Append(resp.State.Set(ctx, &state)...) 
			return
		}
	}

	// Not found => remove from state
	resp.State.RemoveResource(ctx)
}

// Update recreates the SSH key (no PATCH support).
func (r *sshKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Simply call Create with the new plan
	r.Create(ctx, resource.CreateRequest{Plan: req.Plan}, (*resource.CreateResponse)(resp))
}

// Delete removes the SSH key by fingerprint.
func (r *sshKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SshKeyModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)  
	if resp.Diagnostics.HasError() {
		return
	}

	// Build delete request
	body, _ := json.Marshal(map[string]string{"fingerprint": state.Fingerprint.ValueString()})
	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodDelete, r.cfg.APIURL+"/ssh_keys", bytes.NewReader(body))
	httpReq.Header.Set("Authorization", "Bearer "+r.cfg.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := r.cfg.HTTPClient.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("API Error", err.Error())
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK && httpResp.StatusCode != http.StatusNoContent {
		b, _ := io.ReadAll(httpResp.Body)
		resp.Diagnostics.AddError("API Error", fmt.Sprintf("status %d: %s", httpResp.StatusCode, string(b)))
	}
}
