// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package unifi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/internal"
)

// `RadiusProfile` manages RADIUS profiles.
type RadiusProfile struct {
	pulumi.CustomResourceState

	// Specifies whether to use RADIUS accounting. Defaults to `false`.
	AccountingEnabled pulumi.BoolPtrOutput `pulumi:"accountingEnabled"`
	// RADIUS accounting servers.
	AcctServers RadiusProfileAcctServerArrayOutput `pulumi:"acctServers"`
	// RADIUS authentication servers.
	AuthServers RadiusProfileAuthServerArrayOutput `pulumi:"authServers"`
	// Specifies whether to use interim_update. Defaults to `false`.
	InterimUpdateEnabled pulumi.BoolPtrOutput `pulumi:"interimUpdateEnabled"`
	// Specifies interimUpdate interval. Defaults to `3600`.
	InterimUpdateInterval pulumi.IntPtrOutput `pulumi:"interimUpdateInterval"`
	// The name of the profile.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the site to associate the settings with.
	Site pulumi.StringOutput `pulumi:"site"`
	// Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
	UseUsgAcctServer pulumi.BoolPtrOutput `pulumi:"useUsgAcctServer"`
	// Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
	UseUsgAuthServer pulumi.BoolPtrOutput `pulumi:"useUsgAuthServer"`
	// Specifies whether to use vlan on wired connections. Defaults to `false`.
	VlanEnabled pulumi.BoolPtrOutput `pulumi:"vlanEnabled"`
	// Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
	VlanWlanMode pulumi.StringPtrOutput `pulumi:"vlanWlanMode"`
}

// NewRadiusProfile registers a new resource with the given unique name, arguments, and options.
func NewRadiusProfile(ctx *pulumi.Context,
	name string, args *RadiusProfileArgs, opts ...pulumi.ResourceOption) (*RadiusProfile, error) {
	if args == nil {
		args = &RadiusProfileArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RadiusProfile
	err := ctx.RegisterResource("unifi:index/radiusProfile:RadiusProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRadiusProfile gets an existing RadiusProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRadiusProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RadiusProfileState, opts ...pulumi.ResourceOption) (*RadiusProfile, error) {
	var resource RadiusProfile
	err := ctx.ReadResource("unifi:index/radiusProfile:RadiusProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RadiusProfile resources.
type radiusProfileState struct {
	// Specifies whether to use RADIUS accounting. Defaults to `false`.
	AccountingEnabled *bool `pulumi:"accountingEnabled"`
	// RADIUS accounting servers.
	AcctServers []RadiusProfileAcctServer `pulumi:"acctServers"`
	// RADIUS authentication servers.
	AuthServers []RadiusProfileAuthServer `pulumi:"authServers"`
	// Specifies whether to use interim_update. Defaults to `false`.
	InterimUpdateEnabled *bool `pulumi:"interimUpdateEnabled"`
	// Specifies interimUpdate interval. Defaults to `3600`.
	InterimUpdateInterval *int `pulumi:"interimUpdateInterval"`
	// The name of the profile.
	Name *string `pulumi:"name"`
	// The name of the site to associate the settings with.
	Site *string `pulumi:"site"`
	// Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
	UseUsgAcctServer *bool `pulumi:"useUsgAcctServer"`
	// Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
	UseUsgAuthServer *bool `pulumi:"useUsgAuthServer"`
	// Specifies whether to use vlan on wired connections. Defaults to `false`.
	VlanEnabled *bool `pulumi:"vlanEnabled"`
	// Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
	VlanWlanMode *string `pulumi:"vlanWlanMode"`
}

type RadiusProfileState struct {
	// Specifies whether to use RADIUS accounting. Defaults to `false`.
	AccountingEnabled pulumi.BoolPtrInput
	// RADIUS accounting servers.
	AcctServers RadiusProfileAcctServerArrayInput
	// RADIUS authentication servers.
	AuthServers RadiusProfileAuthServerArrayInput
	// Specifies whether to use interim_update. Defaults to `false`.
	InterimUpdateEnabled pulumi.BoolPtrInput
	// Specifies interimUpdate interval. Defaults to `3600`.
	InterimUpdateInterval pulumi.IntPtrInput
	// The name of the profile.
	Name pulumi.StringPtrInput
	// The name of the site to associate the settings with.
	Site pulumi.StringPtrInput
	// Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
	UseUsgAcctServer pulumi.BoolPtrInput
	// Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
	UseUsgAuthServer pulumi.BoolPtrInput
	// Specifies whether to use vlan on wired connections. Defaults to `false`.
	VlanEnabled pulumi.BoolPtrInput
	// Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
	VlanWlanMode pulumi.StringPtrInput
}

func (RadiusProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*radiusProfileState)(nil)).Elem()
}

type radiusProfileArgs struct {
	// Specifies whether to use RADIUS accounting. Defaults to `false`.
	AccountingEnabled *bool `pulumi:"accountingEnabled"`
	// RADIUS accounting servers.
	AcctServers []RadiusProfileAcctServer `pulumi:"acctServers"`
	// RADIUS authentication servers.
	AuthServers []RadiusProfileAuthServer `pulumi:"authServers"`
	// Specifies whether to use interim_update. Defaults to `false`.
	InterimUpdateEnabled *bool `pulumi:"interimUpdateEnabled"`
	// Specifies interimUpdate interval. Defaults to `3600`.
	InterimUpdateInterval *int `pulumi:"interimUpdateInterval"`
	// The name of the profile.
	Name *string `pulumi:"name"`
	// The name of the site to associate the settings with.
	Site *string `pulumi:"site"`
	// Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
	UseUsgAcctServer *bool `pulumi:"useUsgAcctServer"`
	// Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
	UseUsgAuthServer *bool `pulumi:"useUsgAuthServer"`
	// Specifies whether to use vlan on wired connections. Defaults to `false`.
	VlanEnabled *bool `pulumi:"vlanEnabled"`
	// Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
	VlanWlanMode *string `pulumi:"vlanWlanMode"`
}

// The set of arguments for constructing a RadiusProfile resource.
type RadiusProfileArgs struct {
	// Specifies whether to use RADIUS accounting. Defaults to `false`.
	AccountingEnabled pulumi.BoolPtrInput
	// RADIUS accounting servers.
	AcctServers RadiusProfileAcctServerArrayInput
	// RADIUS authentication servers.
	AuthServers RadiusProfileAuthServerArrayInput
	// Specifies whether to use interim_update. Defaults to `false`.
	InterimUpdateEnabled pulumi.BoolPtrInput
	// Specifies interimUpdate interval. Defaults to `3600`.
	InterimUpdateInterval pulumi.IntPtrInput
	// The name of the profile.
	Name pulumi.StringPtrInput
	// The name of the site to associate the settings with.
	Site pulumi.StringPtrInput
	// Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
	UseUsgAcctServer pulumi.BoolPtrInput
	// Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
	UseUsgAuthServer pulumi.BoolPtrInput
	// Specifies whether to use vlan on wired connections. Defaults to `false`.
	VlanEnabled pulumi.BoolPtrInput
	// Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
	VlanWlanMode pulumi.StringPtrInput
}

func (RadiusProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*radiusProfileArgs)(nil)).Elem()
}

type RadiusProfileInput interface {
	pulumi.Input

	ToRadiusProfileOutput() RadiusProfileOutput
	ToRadiusProfileOutputWithContext(ctx context.Context) RadiusProfileOutput
}

func (*RadiusProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**RadiusProfile)(nil)).Elem()
}

func (i *RadiusProfile) ToRadiusProfileOutput() RadiusProfileOutput {
	return i.ToRadiusProfileOutputWithContext(context.Background())
}

func (i *RadiusProfile) ToRadiusProfileOutputWithContext(ctx context.Context) RadiusProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RadiusProfileOutput)
}

func (i *RadiusProfile) ToOutput(ctx context.Context) pulumix.Output[*RadiusProfile] {
	return pulumix.Output[*RadiusProfile]{
		OutputState: i.ToRadiusProfileOutputWithContext(ctx).OutputState,
	}
}

// RadiusProfileArrayInput is an input type that accepts RadiusProfileArray and RadiusProfileArrayOutput values.
// You can construct a concrete instance of `RadiusProfileArrayInput` via:
//
//	RadiusProfileArray{ RadiusProfileArgs{...} }
type RadiusProfileArrayInput interface {
	pulumi.Input

	ToRadiusProfileArrayOutput() RadiusProfileArrayOutput
	ToRadiusProfileArrayOutputWithContext(context.Context) RadiusProfileArrayOutput
}

type RadiusProfileArray []RadiusProfileInput

func (RadiusProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RadiusProfile)(nil)).Elem()
}

func (i RadiusProfileArray) ToRadiusProfileArrayOutput() RadiusProfileArrayOutput {
	return i.ToRadiusProfileArrayOutputWithContext(context.Background())
}

func (i RadiusProfileArray) ToRadiusProfileArrayOutputWithContext(ctx context.Context) RadiusProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RadiusProfileArrayOutput)
}

func (i RadiusProfileArray) ToOutput(ctx context.Context) pulumix.Output[[]*RadiusProfile] {
	return pulumix.Output[[]*RadiusProfile]{
		OutputState: i.ToRadiusProfileArrayOutputWithContext(ctx).OutputState,
	}
}

// RadiusProfileMapInput is an input type that accepts RadiusProfileMap and RadiusProfileMapOutput values.
// You can construct a concrete instance of `RadiusProfileMapInput` via:
//
//	RadiusProfileMap{ "key": RadiusProfileArgs{...} }
type RadiusProfileMapInput interface {
	pulumi.Input

	ToRadiusProfileMapOutput() RadiusProfileMapOutput
	ToRadiusProfileMapOutputWithContext(context.Context) RadiusProfileMapOutput
}

type RadiusProfileMap map[string]RadiusProfileInput

func (RadiusProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RadiusProfile)(nil)).Elem()
}

func (i RadiusProfileMap) ToRadiusProfileMapOutput() RadiusProfileMapOutput {
	return i.ToRadiusProfileMapOutputWithContext(context.Background())
}

func (i RadiusProfileMap) ToRadiusProfileMapOutputWithContext(ctx context.Context) RadiusProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RadiusProfileMapOutput)
}

func (i RadiusProfileMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RadiusProfile] {
	return pulumix.Output[map[string]*RadiusProfile]{
		OutputState: i.ToRadiusProfileMapOutputWithContext(ctx).OutputState,
	}
}

type RadiusProfileOutput struct{ *pulumi.OutputState }

func (RadiusProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RadiusProfile)(nil)).Elem()
}

func (o RadiusProfileOutput) ToRadiusProfileOutput() RadiusProfileOutput {
	return o
}

func (o RadiusProfileOutput) ToRadiusProfileOutputWithContext(ctx context.Context) RadiusProfileOutput {
	return o
}

func (o RadiusProfileOutput) ToOutput(ctx context.Context) pulumix.Output[*RadiusProfile] {
	return pulumix.Output[*RadiusProfile]{
		OutputState: o.OutputState,
	}
}

// Specifies whether to use RADIUS accounting. Defaults to `false`.
func (o RadiusProfileOutput) AccountingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RadiusProfile) pulumi.BoolPtrOutput { return v.AccountingEnabled }).(pulumi.BoolPtrOutput)
}

// RADIUS accounting servers.
func (o RadiusProfileOutput) AcctServers() RadiusProfileAcctServerArrayOutput {
	return o.ApplyT(func(v *RadiusProfile) RadiusProfileAcctServerArrayOutput { return v.AcctServers }).(RadiusProfileAcctServerArrayOutput)
}

// RADIUS authentication servers.
func (o RadiusProfileOutput) AuthServers() RadiusProfileAuthServerArrayOutput {
	return o.ApplyT(func(v *RadiusProfile) RadiusProfileAuthServerArrayOutput { return v.AuthServers }).(RadiusProfileAuthServerArrayOutput)
}

// Specifies whether to use interim_update. Defaults to `false`.
func (o RadiusProfileOutput) InterimUpdateEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RadiusProfile) pulumi.BoolPtrOutput { return v.InterimUpdateEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies interimUpdate interval. Defaults to `3600`.
func (o RadiusProfileOutput) InterimUpdateInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RadiusProfile) pulumi.IntPtrOutput { return v.InterimUpdateInterval }).(pulumi.IntPtrOutput)
}

// The name of the profile.
func (o RadiusProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RadiusProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the site to associate the settings with.
func (o RadiusProfileOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *RadiusProfile) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

// Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
func (o RadiusProfileOutput) UseUsgAcctServer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RadiusProfile) pulumi.BoolPtrOutput { return v.UseUsgAcctServer }).(pulumi.BoolPtrOutput)
}

// Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
func (o RadiusProfileOutput) UseUsgAuthServer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RadiusProfile) pulumi.BoolPtrOutput { return v.UseUsgAuthServer }).(pulumi.BoolPtrOutput)
}

// Specifies whether to use vlan on wired connections. Defaults to `false`.
func (o RadiusProfileOutput) VlanEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RadiusProfile) pulumi.BoolPtrOutput { return v.VlanEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to “.
func (o RadiusProfileOutput) VlanWlanMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RadiusProfile) pulumi.StringPtrOutput { return v.VlanWlanMode }).(pulumi.StringPtrOutput)
}

type RadiusProfileArrayOutput struct{ *pulumi.OutputState }

func (RadiusProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RadiusProfile)(nil)).Elem()
}

func (o RadiusProfileArrayOutput) ToRadiusProfileArrayOutput() RadiusProfileArrayOutput {
	return o
}

func (o RadiusProfileArrayOutput) ToRadiusProfileArrayOutputWithContext(ctx context.Context) RadiusProfileArrayOutput {
	return o
}

func (o RadiusProfileArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RadiusProfile] {
	return pulumix.Output[[]*RadiusProfile]{
		OutputState: o.OutputState,
	}
}

func (o RadiusProfileArrayOutput) Index(i pulumi.IntInput) RadiusProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RadiusProfile {
		return vs[0].([]*RadiusProfile)[vs[1].(int)]
	}).(RadiusProfileOutput)
}

type RadiusProfileMapOutput struct{ *pulumi.OutputState }

func (RadiusProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RadiusProfile)(nil)).Elem()
}

func (o RadiusProfileMapOutput) ToRadiusProfileMapOutput() RadiusProfileMapOutput {
	return o
}

func (o RadiusProfileMapOutput) ToRadiusProfileMapOutputWithContext(ctx context.Context) RadiusProfileMapOutput {
	return o
}

func (o RadiusProfileMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RadiusProfile] {
	return pulumix.Output[map[string]*RadiusProfile]{
		OutputState: o.OutputState,
	}
}

func (o RadiusProfileMapOutput) MapIndex(k pulumi.StringInput) RadiusProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RadiusProfile {
		return vs[0].(map[string]*RadiusProfile)[vs[1].(string)]
	}).(RadiusProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RadiusProfileInput)(nil)).Elem(), &RadiusProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*RadiusProfileArrayInput)(nil)).Elem(), RadiusProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RadiusProfileMapInput)(nil)).Elem(), RadiusProfileMap{})
	pulumi.RegisterOutputType(RadiusProfileOutput{})
	pulumi.RegisterOutputType(RadiusProfileArrayOutput{})
	pulumi.RegisterOutputType(RadiusProfileMapOutput{})
}
