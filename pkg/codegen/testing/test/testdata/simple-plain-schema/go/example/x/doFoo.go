// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"simple-plain-schema/example/internal"
)

func DoFoo(ctx *pulumi.Context, args *DoFooArgs, opts ...pulumi.InvokeOption) error {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv struct{}
	err := ctx.Invoke("example::doFoo", args, &rv, opts...)
	return err
}

type DoFooArgs struct {
	Foo Foo `pulumi:"foo"`
}
