// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package setting

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/internal"
)

var _ = internal.GetEnvOrDefault

type MgmtSshKey struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Public SSH key.
	Key *string `pulumi:"key"`
	// Name of SSH key.
	Name string `pulumi:"name"`
	// Type of SSH key, e.g. ssh-rsa.
	Type string `pulumi:"type"`
}

// MgmtSshKeyInput is an input type that accepts MgmtSshKeyArgs and MgmtSshKeyOutput values.
// You can construct a concrete instance of `MgmtSshKeyInput` via:
//
//	MgmtSshKeyArgs{...}
type MgmtSshKeyInput interface {
	pulumi.Input

	ToMgmtSshKeyOutput() MgmtSshKeyOutput
	ToMgmtSshKeyOutputWithContext(context.Context) MgmtSshKeyOutput
}

type MgmtSshKeyArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput `pulumi:"comment"`
	// Public SSH key.
	Key pulumi.StringPtrInput `pulumi:"key"`
	// Name of SSH key.
	Name pulumi.StringInput `pulumi:"name"`
	// Type of SSH key, e.g. ssh-rsa.
	Type pulumi.StringInput `pulumi:"type"`
}

func (MgmtSshKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MgmtSshKey)(nil)).Elem()
}

func (i MgmtSshKeyArgs) ToMgmtSshKeyOutput() MgmtSshKeyOutput {
	return i.ToMgmtSshKeyOutputWithContext(context.Background())
}

func (i MgmtSshKeyArgs) ToMgmtSshKeyOutputWithContext(ctx context.Context) MgmtSshKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MgmtSshKeyOutput)
}

// MgmtSshKeyArrayInput is an input type that accepts MgmtSshKeyArray and MgmtSshKeyArrayOutput values.
// You can construct a concrete instance of `MgmtSshKeyArrayInput` via:
//
//	MgmtSshKeyArray{ MgmtSshKeyArgs{...} }
type MgmtSshKeyArrayInput interface {
	pulumi.Input

	ToMgmtSshKeyArrayOutput() MgmtSshKeyArrayOutput
	ToMgmtSshKeyArrayOutputWithContext(context.Context) MgmtSshKeyArrayOutput
}

type MgmtSshKeyArray []MgmtSshKeyInput

func (MgmtSshKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MgmtSshKey)(nil)).Elem()
}

func (i MgmtSshKeyArray) ToMgmtSshKeyArrayOutput() MgmtSshKeyArrayOutput {
	return i.ToMgmtSshKeyArrayOutputWithContext(context.Background())
}

func (i MgmtSshKeyArray) ToMgmtSshKeyArrayOutputWithContext(ctx context.Context) MgmtSshKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MgmtSshKeyArrayOutput)
}

type MgmtSshKeyOutput struct{ *pulumi.OutputState }

func (MgmtSshKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MgmtSshKey)(nil)).Elem()
}

func (o MgmtSshKeyOutput) ToMgmtSshKeyOutput() MgmtSshKeyOutput {
	return o
}

func (o MgmtSshKeyOutput) ToMgmtSshKeyOutputWithContext(ctx context.Context) MgmtSshKeyOutput {
	return o
}

// Comment.
func (o MgmtSshKeyOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MgmtSshKey) *string { return v.Comment }).(pulumi.StringPtrOutput)
}

// Public SSH key.
func (o MgmtSshKeyOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MgmtSshKey) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// Name of SSH key.
func (o MgmtSshKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MgmtSshKey) string { return v.Name }).(pulumi.StringOutput)
}

// Type of SSH key, e.g. ssh-rsa.
func (o MgmtSshKeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MgmtSshKey) string { return v.Type }).(pulumi.StringOutput)
}

type MgmtSshKeyArrayOutput struct{ *pulumi.OutputState }

func (MgmtSshKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MgmtSshKey)(nil)).Elem()
}

func (o MgmtSshKeyArrayOutput) ToMgmtSshKeyArrayOutput() MgmtSshKeyArrayOutput {
	return o
}

func (o MgmtSshKeyArrayOutput) ToMgmtSshKeyArrayOutputWithContext(ctx context.Context) MgmtSshKeyArrayOutput {
	return o
}

func (o MgmtSshKeyArrayOutput) Index(i pulumi.IntInput) MgmtSshKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MgmtSshKey {
		return vs[0].([]MgmtSshKey)[vs[1].(int)]
	}).(MgmtSshKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MgmtSshKeyInput)(nil)).Elem(), MgmtSshKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MgmtSshKeyArrayInput)(nil)).Elem(), MgmtSshKeyArray{})
	pulumi.RegisterOutputType(MgmtSshKeyOutput{})
	pulumi.RegisterOutputType(MgmtSshKeyArrayOutput{})
}
