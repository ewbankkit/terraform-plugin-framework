package tfsdk

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type testServeDataSourceTypeValidateConfig struct{}

func (dt testServeDataSourceTypeValidateConfig) GetSchema(_ context.Context) (Schema, []*tfprotov6.Diagnostic) {
	return Schema{
		Attributes: map[string]Attribute{
			"string": {
				Type:     types.StringType,
				Optional: true,
			},
		},
	}, nil
}

func (dt testServeDataSourceTypeValidateConfig) NewDataSource(_ context.Context, p Provider) (DataSource, []*tfprotov6.Diagnostic) {
	provider, ok := p.(*testServeProvider)
	if !ok {
		prov, ok := p.(*testServeProviderWithMetaSchema)
		if !ok {
			panic(fmt.Sprintf("unexpected provider type %T", p))
		}
		provider = prov.testServeProvider
	}
	return testServeDataSourceValidateConfig{
		provider: provider,
	}, nil
}

var testServeDataSourceTypeValidateConfigSchema = &tfprotov6.Schema{
	Block: &tfprotov6.SchemaBlock{
		Attributes: []*tfprotov6.SchemaAttribute{
			{
				Name:     "string",
				Optional: true,
				Type:     tftypes.String,
			},
		},
	},
}

var testServeDataSourceTypeValidateConfigType = tftypes.Object{
	AttributeTypes: map[string]tftypes.Type{
		"string": tftypes.String,
	},
}

type testServeDataSourceValidateConfig struct {
	provider *testServeProvider
}

func (r testServeDataSourceValidateConfig) Read(ctx context.Context, req ReadDataSourceRequest, resp *ReadDataSourceResponse) {
}

func (r testServeDataSourceValidateConfig) ValidateConfig(ctx context.Context, req ValidateDataSourceConfigRequest, resp *ValidateDataSourceConfigResponse) {
	r.provider.validateDataSourceConfigCalledDataSourceType = "test_validate_config"
	r.provider.validateDataSourceConfigImpl(ctx, req, resp)
}