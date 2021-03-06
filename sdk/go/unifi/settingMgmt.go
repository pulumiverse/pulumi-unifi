// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package unifi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `SettingMgmt` manages settings for a unifi site.
type SettingMgmt struct {
	pulumi.CustomResourceState

	// Automatically upgrade device firmware.
	AutoUpgrade pulumi.BoolPtrOutput `pulumi:"autoUpgrade"`
	// The name of the site to associate the settings with.
	Site pulumi.StringOutput `pulumi:"site"`
	// Enable SSH authentication.
	SshEnabled pulumi.BoolPtrOutput `pulumi:"sshEnabled"`
	// SSH key.
	SshKeys SettingMgmtSshKeyArrayOutput `pulumi:"sshKeys"`
}

// NewSettingMgmt registers a new resource with the given unique name, arguments, and options.
func NewSettingMgmt(ctx *pulumi.Context,
	name string, args *SettingMgmtArgs, opts ...pulumi.ResourceOption) (*SettingMgmt, error) {
	if args == nil {
		args = &SettingMgmtArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SettingMgmt
	err := ctx.RegisterResource("unifi:index/settingMgmt:SettingMgmt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSettingMgmt gets an existing SettingMgmt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSettingMgmt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SettingMgmtState, opts ...pulumi.ResourceOption) (*SettingMgmt, error) {
	var resource SettingMgmt
	err := ctx.ReadResource("unifi:index/settingMgmt:SettingMgmt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SettingMgmt resources.
type settingMgmtState struct {
	// Automatically upgrade device firmware.
	AutoUpgrade *bool `pulumi:"autoUpgrade"`
	// The name of the site to associate the settings with.
	Site *string `pulumi:"site"`
	// Enable SSH authentication.
	SshEnabled *bool `pulumi:"sshEnabled"`
	// SSH key.
	SshKeys []SettingMgmtSshKey `pulumi:"sshKeys"`
}

type SettingMgmtState struct {
	// Automatically upgrade device firmware.
	AutoUpgrade pulumi.BoolPtrInput
	// The name of the site to associate the settings with.
	Site pulumi.StringPtrInput
	// Enable SSH authentication.
	SshEnabled pulumi.BoolPtrInput
	// SSH key.
	SshKeys SettingMgmtSshKeyArrayInput
}

func (SettingMgmtState) ElementType() reflect.Type {
	return reflect.TypeOf((*settingMgmtState)(nil)).Elem()
}

type settingMgmtArgs struct {
	// Automatically upgrade device firmware.
	AutoUpgrade *bool `pulumi:"autoUpgrade"`
	// The name of the site to associate the settings with.
	Site *string `pulumi:"site"`
	// Enable SSH authentication.
	SshEnabled *bool `pulumi:"sshEnabled"`
	// SSH key.
	SshKeys []SettingMgmtSshKey `pulumi:"sshKeys"`
}

// The set of arguments for constructing a SettingMgmt resource.
type SettingMgmtArgs struct {
	// Automatically upgrade device firmware.
	AutoUpgrade pulumi.BoolPtrInput
	// The name of the site to associate the settings with.
	Site pulumi.StringPtrInput
	// Enable SSH authentication.
	SshEnabled pulumi.BoolPtrInput
	// SSH key.
	SshKeys SettingMgmtSshKeyArrayInput
}

func (SettingMgmtArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*settingMgmtArgs)(nil)).Elem()
}

type SettingMgmtInput interface {
	pulumi.Input

	ToSettingMgmtOutput() SettingMgmtOutput
	ToSettingMgmtOutputWithContext(ctx context.Context) SettingMgmtOutput
}

func (*SettingMgmt) ElementType() reflect.Type {
	return reflect.TypeOf((**SettingMgmt)(nil)).Elem()
}

func (i *SettingMgmt) ToSettingMgmtOutput() SettingMgmtOutput {
	return i.ToSettingMgmtOutputWithContext(context.Background())
}

func (i *SettingMgmt) ToSettingMgmtOutputWithContext(ctx context.Context) SettingMgmtOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingMgmtOutput)
}

// SettingMgmtArrayInput is an input type that accepts SettingMgmtArray and SettingMgmtArrayOutput values.
// You can construct a concrete instance of `SettingMgmtArrayInput` via:
//
//          SettingMgmtArray{ SettingMgmtArgs{...} }
type SettingMgmtArrayInput interface {
	pulumi.Input

	ToSettingMgmtArrayOutput() SettingMgmtArrayOutput
	ToSettingMgmtArrayOutputWithContext(context.Context) SettingMgmtArrayOutput
}

type SettingMgmtArray []SettingMgmtInput

func (SettingMgmtArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SettingMgmt)(nil)).Elem()
}

func (i SettingMgmtArray) ToSettingMgmtArrayOutput() SettingMgmtArrayOutput {
	return i.ToSettingMgmtArrayOutputWithContext(context.Background())
}

func (i SettingMgmtArray) ToSettingMgmtArrayOutputWithContext(ctx context.Context) SettingMgmtArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingMgmtArrayOutput)
}

// SettingMgmtMapInput is an input type that accepts SettingMgmtMap and SettingMgmtMapOutput values.
// You can construct a concrete instance of `SettingMgmtMapInput` via:
//
//          SettingMgmtMap{ "key": SettingMgmtArgs{...} }
type SettingMgmtMapInput interface {
	pulumi.Input

	ToSettingMgmtMapOutput() SettingMgmtMapOutput
	ToSettingMgmtMapOutputWithContext(context.Context) SettingMgmtMapOutput
}

type SettingMgmtMap map[string]SettingMgmtInput

func (SettingMgmtMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SettingMgmt)(nil)).Elem()
}

func (i SettingMgmtMap) ToSettingMgmtMapOutput() SettingMgmtMapOutput {
	return i.ToSettingMgmtMapOutputWithContext(context.Background())
}

func (i SettingMgmtMap) ToSettingMgmtMapOutputWithContext(ctx context.Context) SettingMgmtMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingMgmtMapOutput)
}

type SettingMgmtOutput struct{ *pulumi.OutputState }

func (SettingMgmtOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SettingMgmt)(nil)).Elem()
}

func (o SettingMgmtOutput) ToSettingMgmtOutput() SettingMgmtOutput {
	return o
}

func (o SettingMgmtOutput) ToSettingMgmtOutputWithContext(ctx context.Context) SettingMgmtOutput {
	return o
}

// Automatically upgrade device firmware.
func (o SettingMgmtOutput) AutoUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SettingMgmt) pulumi.BoolPtrOutput { return v.AutoUpgrade }).(pulumi.BoolPtrOutput)
}

// The name of the site to associate the settings with.
func (o SettingMgmtOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *SettingMgmt) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

// Enable SSH authentication.
func (o SettingMgmtOutput) SshEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SettingMgmt) pulumi.BoolPtrOutput { return v.SshEnabled }).(pulumi.BoolPtrOutput)
}

// SSH key.
func (o SettingMgmtOutput) SshKeys() SettingMgmtSshKeyArrayOutput {
	return o.ApplyT(func(v *SettingMgmt) SettingMgmtSshKeyArrayOutput { return v.SshKeys }).(SettingMgmtSshKeyArrayOutput)
}

type SettingMgmtArrayOutput struct{ *pulumi.OutputState }

func (SettingMgmtArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SettingMgmt)(nil)).Elem()
}

func (o SettingMgmtArrayOutput) ToSettingMgmtArrayOutput() SettingMgmtArrayOutput {
	return o
}

func (o SettingMgmtArrayOutput) ToSettingMgmtArrayOutputWithContext(ctx context.Context) SettingMgmtArrayOutput {
	return o
}

func (o SettingMgmtArrayOutput) Index(i pulumi.IntInput) SettingMgmtOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SettingMgmt {
		return vs[0].([]*SettingMgmt)[vs[1].(int)]
	}).(SettingMgmtOutput)
}

type SettingMgmtMapOutput struct{ *pulumi.OutputState }

func (SettingMgmtMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SettingMgmt)(nil)).Elem()
}

func (o SettingMgmtMapOutput) ToSettingMgmtMapOutput() SettingMgmtMapOutput {
	return o
}

func (o SettingMgmtMapOutput) ToSettingMgmtMapOutputWithContext(ctx context.Context) SettingMgmtMapOutput {
	return o
}

func (o SettingMgmtMapOutput) MapIndex(k pulumi.StringInput) SettingMgmtOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SettingMgmt {
		return vs[0].(map[string]*SettingMgmt)[vs[1].(string)]
	}).(SettingMgmtOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SettingMgmtInput)(nil)).Elem(), &SettingMgmt{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingMgmtArrayInput)(nil)).Elem(), SettingMgmtArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingMgmtMapInput)(nil)).Elem(), SettingMgmtMap{})
	pulumi.RegisterOutputType(SettingMgmtOutput{})
	pulumi.RegisterOutputType(SettingMgmtArrayOutput{})
	pulumi.RegisterOutputType(SettingMgmtMapOutput{})
}
