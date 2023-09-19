// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"simple-plain-schema/example/internal"
)

var _ = internal.GetEnvOrDefault

type Foo struct {
	A bool    `pulumi:"a"`
	B *bool   `pulumi:"b"`
	C int     `pulumi:"c"`
	D *int    `pulumi:"d"`
	E string  `pulumi:"e"`
	F *string `pulumi:"f"`
}

type FooArgs struct {
	A pulumix.Input[bool]    `pulumi:"a"`
	B pulumix.Input[*bool]   `pulumi:"b"`
	C pulumix.Input[int]     `pulumi:"c"`
	D pulumix.Input[*int]    `pulumi:"d"`
	E pulumix.Input[string]  `pulumi:"e"`
	F pulumix.Input[*string] `pulumi:"f"`
}

func (FooArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Foo)(nil)).Elem()
}

func (i FooArgs) ToFooOutput() FooOutput {
	return i.ToFooOutputWithContext(context.Background())
}

func (i FooArgs) ToFooOutputWithContext(ctx context.Context) FooOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FooOutput)
}

func (i *FooArgs) ToOutput(ctx context.Context) pulumix.Output[*FooArgs] {
	return pulumix.Val(i)
}

type FooOutput struct{ *pulumi.OutputState }

func (FooOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Foo)(nil)).Elem()
}

func (o FooOutput) ToFooOutput() FooOutput {
	return o
}

func (o FooOutput) ToFooOutputWithContext(ctx context.Context) FooOutput {
	return o
}

func (o FooOutput) ToOutput(ctx context.Context) pulumix.Output[Foo] {
	return pulumix.Output[Foo]{
		OutputState: o.OutputState,
	}
}

func (o FooOutput) A() pulumix.Output[bool] {
	return pulumix.Apply[Foo](o, func(v Foo) bool { return v.A })
}

func (o FooOutput) B() pulumix.Output[*bool] {
	return pulumix.Apply[Foo](o, func(v Foo) *bool { return v.B })
}

func (o FooOutput) C() pulumix.Output[int] {
	return pulumix.Apply[Foo](o, func(v Foo) int { return v.C })
}

func (o FooOutput) D() pulumix.Output[*int] {
	return pulumix.Apply[Foo](o, func(v Foo) *int { return v.D })
}

func (o FooOutput) E() pulumix.Output[string] {
	return pulumix.Apply[Foo](o, func(v Foo) string { return v.E })
}

func (o FooOutput) F() pulumix.Output[*string] {
	return pulumix.Apply[Foo](o, func(v Foo) *string { return v.F })
}

func init() {
	pulumi.RegisterOutputType(FooOutput{})
}
