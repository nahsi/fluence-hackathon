package provider_fluence

import (
	"context"
	"net/http"

	"github.com/cloudlesslabs/terraform-provider-fluence/internal/client"
	"github.com/cloudlesslabs/terraform-provider-fluence/internal/datasource_vm_instances"
	"github.com/cloudlesslabs/terraform-provider-fluence/internal/resource_ssh_key"
	"github.com/cloudlesslabs/terraform-provider-fluence/internal/resource_vm"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*fluenceProvider)(nil)

// New is the factory Terraform calls.
func New() func() provider.Provider {
	return func() provider.Provider { return &fluenceProvider{} }
}

type fluenceProvider struct{}

//──────────────────────── Metadata & Schema ────────────────────────

func (p *fluenceProvider) Metadata(
	_ context.Context,
	_ provider.MetadataRequest,
	resp *provider.MetadataResponse,
) {
	resp.TypeName = "fluence"
}

func (p *fluenceProvider) Schema(
	ctx context.Context,
	_ provider.SchemaRequest,
	resp *provider.SchemaResponse,
) {
	resp.Schema = FluenceProviderSchema(ctx)
}

//──────────────────────── Configure ────────────────────────────────

func (p *fluenceProvider) Configure(
	ctx context.Context,
	req provider.ConfigureRequest,
	resp *provider.ConfigureResponse,
) {
	// Decode the provider block into the generated model.
	var cfg FluenceModel
	diags := req.Config.Get(ctx, &cfg)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate api_key.
	if cfg.ApiKey.IsNull() || cfg.ApiKey.ValueString() == "" {
		resp.Diagnostics.AddError("Missing api_key", "`api_key` must be set")
		return
	}
	apiKey := cfg.ApiKey.ValueString()

	// Default api_url.
	apiURL := "https://api.fluence.dev"
	if !cfg.ApiUrl.IsNull() && cfg.ApiUrl.ValueString() != "" {
		apiURL = cfg.ApiUrl.ValueString()
	}

	// Shared HTTP client.
	httpClient := &http.Client{}

	// Hand the same config struct to every resource / data‑source.
	shared := client.Config{
		APIKey:     apiKey,
		APIURL:     apiURL,
		HTTPClient: httpClient,
	}
	resp.ResourceData = shared
	resp.DataSourceData = shared
}

//──────────────────────── Registration ─────────────────────────────

func (p *fluenceProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resource_ssh_key.NewSshKeyResource,
		resource_vm.NewVmResource,
	}
}

func (p *fluenceProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasource_vm_instances.NewVmInstancesDataSource,
	}
}
