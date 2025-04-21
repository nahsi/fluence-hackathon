package datasource_vm_instances

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudlesslabs/terraform-provider-fluence/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the data source implements the core DataSource interface.
var _ datasource.DataSource = (*vmInstancesDataSource)(nil)

// NewVmInstancesDataSource is the factory function called by Terraform.
func NewVmInstancesDataSource() datasource.DataSource {
	return &vmInstancesDataSource{}
}

type vmInstancesDataSource struct {
	cfg client.Config
}

type vmInstancesDataSourceModel struct {
	Raw types.String `tfsdk:"raw"`
}

// Metadata returns the data source type name.
func (d *vmInstancesDataSource) Metadata(
	_ context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_vm_instances"
}

// Schema defines the schema for the data source.
func (d *vmInstancesDataSource) Schema(
	_ context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"raw": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Raw JSON returned by **GET /vms/v2**.",
			},
		},
	}
}

// Configure makes the provider-config available to the data source.
func (d *vmInstancesDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	resp *datasource.ConfigureResponse,
) {
	// Pull the shared client.Config from provider
	if cfg, ok := req.ProviderData.(client.Config); ok {
		d.cfg = cfg
	}
}

// Read queries the API and sets the state accordingly.
func (d *vmInstancesDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	// build the request URL
	url := fmt.Sprintf("%s/vms/v2", d.cfg.APIURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		resp.Diagnostics.AddError("Request build error", err.Error())
		return
	}
	httpReq.Header.Set("Authorization", "Bearer "+d.cfg.APIKey)

	httpResp, err := d.cfg.HTTPClient.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError("HTTP error", err.Error())
		return
	}
	defer httpResp.Body.Close()

	// non‑2xx → diagnostic
	if httpResp.StatusCode >= 400 {
		b, _ := io.ReadAll(httpResp.Body)
		resp.Diagnostics.AddError(
			"API error",
			fmt.Sprintf("status %d: %s", httpResp.StatusCode, string(b)),
		)
		return
	}

	// decode and re‑encode for pretty JSON
	var dec interface{}
	if err := json.NewDecoder(httpResp.Body).Decode(&dec); err != nil {
		resp.Diagnostics.AddError("Decode error", err.Error())
		return
	}
	out, _ := json.MarshalIndent(dec, "", "  ")

	// set state
	var state vmInstancesDataSourceModel
	state.Raw = types.StringValue(string(out))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...) 
}
