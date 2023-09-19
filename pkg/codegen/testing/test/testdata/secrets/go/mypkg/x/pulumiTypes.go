// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mypkg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"secrets/mypkg/internal"
)

var _ = internal.GetEnvOrDefault

type Config struct {
	Foo *string `pulumi:"foo"`
}

type ConfigArgs struct {
	Foo pulumix.Input[*string] `pulumi:"foo"`
}

func (ConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Config)(nil)).Elem()
}

func (i ConfigArgs) ToConfigOutput() ConfigOutput {
	return i.ToConfigOutputWithContext(context.Background())
}

func (i ConfigArgs) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigOutput)
}

func (i *ConfigArgs) ToOutput(ctx context.Context) pulumix.Output[*ConfigArgs] {
	return pulumix.Val(i)
}

type ConfigOutput struct{ *pulumi.OutputState }

func (ConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Config)(nil)).Elem()
}

func (o ConfigOutput) ToConfigOutput() ConfigOutput {
	return o
}

func (o ConfigOutput) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return o
}

func (o ConfigOutput) ToOutput(ctx context.Context) pulumix.Output[Config] {
	return pulumix.Output[Config]{
		OutputState: o.OutputState,
	}
}

func (o ConfigOutput) Foo() pulumix.Output[*string] {
	return pulumix.Apply[Config](o, func(v Config) *string { return v.Foo })
}

func init() {
	pulumi.RegisterOutputType(ConfigOutput{})
}
