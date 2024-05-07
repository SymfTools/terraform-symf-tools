package json

import (
	"context"

	"github.com/Boltairex/terraform-symf-tools/json/functions"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Enkapsulacja providera
var (
    _ provider.Provider = &JsonProvider{}
)

type JsonProvider struct {
	provider.ProviderWithFunctions

	version string
}

type DataStorage struct {
	test string
}

func New(version string) func() provider.Provider {

	return func() provider.Provider {
		return &JsonProvider{
			version: version,
		}
	}
}

func (p *JsonProvider) Metadata(context context.Context, request provider.MetadataRequest, response *provider.MetadataResponse) {
	tflog.Info(context, "Setting JsonProvider metadata...")
    response.TypeName = "symf"
    response.Version = p.version
}

// Schema should return the schema for this provider.
func (p *JsonProvider) Schema(context context.Context, request provider.SchemaRequest, response *provider.SchemaResponse) {
	tflog.Info(context, "Configuring JsonProvider...")

	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"test": schema.StringAttribute{
				MarkdownDescription: "Example provider attribute",
				Optional:            true,
			},
		},
	}
}

// Configure is called at the beginning of the provider lifecycle, when
// Terraform sends to the provider the values the user specified in the
// provider configuration block.
func (p *JsonProvider) Configure(context context.Context, request provider.ConfigureRequest, response *provider.ConfigureResponse) {
	var data DataStorage

	response.Diagnostics.Append(request.Config.Get(context, &data)...)

	if response.Diagnostics.HasError() {
		return
	}
}

// DataSources returns a slice of functions to instantiate each DataSource
// implementation.
func (p *JsonProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource {
		NewExampleDataSource,
	};
}

// Resources returns a slice of functions to instantiate each Resource
// implementation.
func (p *JsonProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource {
		NewExampleResource,
	};
}

func (p *JsonProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function {
		functions.NewJsonReplacer,
		NewExampleFunction,
	}
}