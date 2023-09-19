// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mypkg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"output-funcs/mypkg/internal"
)

// Another failing example. A list of SSIS object metadata.
// API Version: 2018-06-01.
func GetIntegrationRuntimeObjectMetadatum(ctx *pulumi.Context, args *GetIntegrationRuntimeObjectMetadatumArgs, opts ...pulumi.InvokeOption) (*GetIntegrationRuntimeObjectMetadatumResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIntegrationRuntimeObjectMetadatumResult
	err := ctx.Invoke("mypkg::getIntegrationRuntimeObjectMetadatum", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeObjectMetadatumArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The integration runtime name.
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	// Metadata path.
	MetadataPath *string `pulumi:"metadataPath"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A list of SSIS object metadata.
type GetIntegrationRuntimeObjectMetadatumResult struct {
	// The link to the next page of results, if any remaining results exist.
	NextLink *string `pulumi:"nextLink"`
	// List of SSIS object metadata.
	Value []interface{} `pulumi:"value"`
}

func GetIntegrationRuntimeObjectMetadatumOutput(ctx *pulumi.Context, args GetIntegrationRuntimeObjectMetadatumOutputArgs, opts ...pulumi.InvokeOption) GetIntegrationRuntimeObjectMetadatumResultOutput {
	outputResult := pulumix.ApplyErr[*GetIntegrationRuntimeObjectMetadatumArgs](args.ToOutput(), func(plainArgs *GetIntegrationRuntimeObjectMetadatumArgs) (*GetIntegrationRuntimeObjectMetadatumResult, error) {
		return GetIntegrationRuntimeObjectMetadatum(ctx, plainArgs, opts...)
	})

	return pulumix.Cast[GetIntegrationRuntimeObjectMetadatumResultOutput, *GetIntegrationRuntimeObjectMetadatumResult](outputResult)
}

type GetIntegrationRuntimeObjectMetadatumOutputArgs struct {
	// The factory name.
	FactoryName pulumix.Input[string] `pulumi:"factoryName"`
	// The integration runtime name.
	IntegrationRuntimeName pulumix.Input[string] `pulumi:"integrationRuntimeName"`
	// Metadata path.
	MetadataPath pulumix.Input[*string] `pulumi:"metadataPath"`
	// The resource group name.
	ResourceGroupName pulumix.Input[string] `pulumi:"resourceGroupName"`
}

func (args GetIntegrationRuntimeObjectMetadatumOutputArgs) ToOutput() pulumix.Output[*GetIntegrationRuntimeObjectMetadatumArgs] {
	allArgs := pulumix.All(
		args.FactoryName.ToOutput(context.Background()).AsAny(),
		args.IntegrationRuntimeName.ToOutput(context.Background()).AsAny(),
		args.MetadataPath.ToOutput(context.Background()).AsAny(),
		args.ResourceGroupName.ToOutput(context.Background()).AsAny())
	return pulumix.Apply[[]any](allArgs, func(resolvedArgs []interface{}) *GetIntegrationRuntimeObjectMetadatumArgs {
		return &GetIntegrationRuntimeObjectMetadatumArgs{
			FactoryName:            resolvedArgs[0].(string),
			IntegrationRuntimeName: resolvedArgs[1].(string),
			MetadataPath:           resolvedArgs[2].(*string),
			ResourceGroupName:      resolvedArgs[3].(string),
		}
	})
}

type GetIntegrationRuntimeObjectMetadatumResultOutput struct{ *pulumi.OutputState }

func (GetIntegrationRuntimeObjectMetadatumResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeObjectMetadatumResult)(nil)).Elem()
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) ToOutput(context.Context) pulumix.Output[*GetIntegrationRuntimeObjectMetadatumResult] {
	return pulumix.Output[*GetIntegrationRuntimeObjectMetadatumResult]{
		OutputState: o.OutputState,
	}
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) NextLink() pulumix.Output[*string] {
	return pulumix.Apply[*GetIntegrationRuntimeObjectMetadatumResult](o, func(v *GetIntegrationRuntimeObjectMetadatumResult) *string { return v.NextLink })
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) Value() pulumix.ArrayOutput[any] {
	value := pulumix.Apply[*GetIntegrationRuntimeObjectMetadatumResult](o, func(v *GetIntegrationRuntimeObjectMetadatumResult) []interface{} { return v.Value })
	return pulumix.ArrayOutput[any]{
		OutputState: value.OutputState,
	}
}
