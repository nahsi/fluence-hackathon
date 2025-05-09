// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto6_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/internal/fromproto6"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwschema"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func TestValidateResourceConfigRequest(t *testing.T) {
	t.Parallel()

	testProto6Type := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"test_attribute": tftypes.String,
		},
	}

	testProto6Value := tftypes.NewValue(testProto6Type, map[string]tftypes.Value{
		"test_attribute": tftypes.NewValue(tftypes.String, "test-value"),
	})

	testProto6DynamicValue, err := tfprotov6.NewDynamicValue(testProto6Type, testProto6Value)

	if err != nil {
		t.Fatalf("unexpected error calling tfprotov6.NewDynamicValue(): %s", err)
	}

	testFwSchema := schema.Schema{
		Attributes: map[string]schema.Attribute{
			"test_attribute": schema.StringAttribute{
				Required: true,
			},
		},
	}

	testCases := map[string]struct {
		input               *tfprotov6.ValidateResourceConfigRequest
		resourceSchema      fwschema.Schema
		resource            resource.Resource
		expected            *fwserver.ValidateResourceConfigRequest
		expectedDiagnostics diag.Diagnostics
	}{
		"nil": {
			input:    nil,
			expected: nil,
		},
		"empty": {
			input:    &tfprotov6.ValidateResourceConfigRequest{},
			expected: &fwserver.ValidateResourceConfigRequest{},
		},
		"config-missing-schema": {
			input: &tfprotov6.ValidateResourceConfigRequest{
				Config: &testProto6DynamicValue,
			},
			expected: &fwserver.ValidateResourceConfigRequest{},
			expectedDiagnostics: diag.Diagnostics{
				diag.NewErrorDiagnostic(
					"Unable to Convert Configuration",
					"An unexpected error was encountered when converting the configuration from the protocol type. "+
						"This is always an issue in terraform-plugin-framework used to implement the provider and should be reported to the provider developers.\n\n"+
						"Please report this to the provider developer:\n\n"+
						"Missing schema.",
				),
			},
		},
		"config": {
			input: &tfprotov6.ValidateResourceConfigRequest{
				Config: &testProto6DynamicValue,
			},
			resourceSchema: testFwSchema,
			expected: &fwserver.ValidateResourceConfigRequest{
				Config: &tfsdk.Config{
					Raw:    testProto6Value,
					Schema: testFwSchema,
				},
			},
		},
		"client-capabilities": {
			input: &tfprotov6.ValidateResourceConfigRequest{
				Config: &testProto6DynamicValue,
				ClientCapabilities: &tfprotov6.ValidateResourceConfigClientCapabilities{
					WriteOnlyAttributesAllowed: true,
				},
			},
			resourceSchema: testFwSchema,
			expected: &fwserver.ValidateResourceConfigRequest{
				ClientCapabilities: resource.ValidateConfigClientCapabilities{
					WriteOnlyAttributesAllowed: true,
				},
				Config: &tfsdk.Config{
					Raw:    testProto6Value,
					Schema: testFwSchema,
				},
			},
		},
		"client-capabilities-not-set": {
			input: &tfprotov6.ValidateResourceConfigRequest{
				Config: &testProto6DynamicValue,
			},
			resourceSchema: testFwSchema,
			expected: &fwserver.ValidateResourceConfigRequest{
				ClientCapabilities: resource.ValidateConfigClientCapabilities{
					WriteOnlyAttributesAllowed: false,
				},
				Config: &tfsdk.Config{
					Raw:    testProto6Value,
					Schema: testFwSchema,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, diags := fromproto6.ValidateResourceConfigRequest(context.Background(), testCase.input, testCase.resource, testCase.resourceSchema)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}

			if diff := cmp.Diff(diags, testCase.expectedDiagnostics); diff != "" {
				t.Errorf("unexpected diagnostics difference: %s", diff)
			}
		})
	}
}
