// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package unifi

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `FirewallGroup` manages groups of addresses or ports for use in firewall rules (`FirewallRule`).
type FirewallGroup struct {
	pulumi.CustomResourceState

	// The members of the firewall group.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The name of the firewall group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the site to associate the firewall group with.
	Site pulumi.StringOutput `pulumi:"site"`
	// The type of the firewall group. Must be one of: `address-group`, `port-group`, or `ipv6-address-group`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallGroup(ctx *pulumi.Context,
	name string, args *FirewallGroupArgs, opts ...pulumi.ResourceOption) (*FirewallGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallGroup
	err := ctx.RegisterResource("unifi:index/firewallGroup:FirewallGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallGroup gets an existing FirewallGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallGroupState, opts ...pulumi.ResourceOption) (*FirewallGroup, error) {
	var resource FirewallGroup
	err := ctx.ReadResource("unifi:index/firewallGroup:FirewallGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallGroup resources.
type firewallGroupState struct {
	// The members of the firewall group.
	Members []string `pulumi:"members"`
	// The name of the firewall group.
	Name *string `pulumi:"name"`
	// The name of the site to associate the firewall group with.
	Site *string `pulumi:"site"`
	// The type of the firewall group. Must be one of: `address-group`, `port-group`, or `ipv6-address-group`.
	Type *string `pulumi:"type"`
}

type FirewallGroupState struct {
	// The members of the firewall group.
	Members pulumi.StringArrayInput
	// The name of the firewall group.
	Name pulumi.StringPtrInput
	// The name of the site to associate the firewall group with.
	Site pulumi.StringPtrInput
	// The type of the firewall group. Must be one of: `address-group`, `port-group`, or `ipv6-address-group`.
	Type pulumi.StringPtrInput
}

func (FirewallGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallGroupState)(nil)).Elem()
}

type firewallGroupArgs struct {
	// The members of the firewall group.
	Members []string `pulumi:"members"`
	// The name of the firewall group.
	Name *string `pulumi:"name"`
	// The name of the site to associate the firewall group with.
	Site *string `pulumi:"site"`
	// The type of the firewall group. Must be one of: `address-group`, `port-group`, or `ipv6-address-group`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a FirewallGroup resource.
type FirewallGroupArgs struct {
	// The members of the firewall group.
	Members pulumi.StringArrayInput
	// The name of the firewall group.
	Name pulumi.StringPtrInput
	// The name of the site to associate the firewall group with.
	Site pulumi.StringPtrInput
	// The type of the firewall group. Must be one of: `address-group`, `port-group`, or `ipv6-address-group`.
	Type pulumi.StringInput
}

func (FirewallGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallGroupArgs)(nil)).Elem()
}

type FirewallGroupInput interface {
	pulumi.Input

	ToFirewallGroupOutput() FirewallGroupOutput
	ToFirewallGroupOutputWithContext(ctx context.Context) FirewallGroupOutput
}

func (*FirewallGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallGroup)(nil)).Elem()
}

func (i *FirewallGroup) ToFirewallGroupOutput() FirewallGroupOutput {
	return i.ToFirewallGroupOutputWithContext(context.Background())
}

func (i *FirewallGroup) ToFirewallGroupOutputWithContext(ctx context.Context) FirewallGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallGroupOutput)
}

// FirewallGroupArrayInput is an input type that accepts FirewallGroupArray and FirewallGroupArrayOutput values.
// You can construct a concrete instance of `FirewallGroupArrayInput` via:
//
//          FirewallGroupArray{ FirewallGroupArgs{...} }
type FirewallGroupArrayInput interface {
	pulumi.Input

	ToFirewallGroupArrayOutput() FirewallGroupArrayOutput
	ToFirewallGroupArrayOutputWithContext(context.Context) FirewallGroupArrayOutput
}

type FirewallGroupArray []FirewallGroupInput

func (FirewallGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallGroup)(nil)).Elem()
}

func (i FirewallGroupArray) ToFirewallGroupArrayOutput() FirewallGroupArrayOutput {
	return i.ToFirewallGroupArrayOutputWithContext(context.Background())
}

func (i FirewallGroupArray) ToFirewallGroupArrayOutputWithContext(ctx context.Context) FirewallGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallGroupArrayOutput)
}

// FirewallGroupMapInput is an input type that accepts FirewallGroupMap and FirewallGroupMapOutput values.
// You can construct a concrete instance of `FirewallGroupMapInput` via:
//
//          FirewallGroupMap{ "key": FirewallGroupArgs{...} }
type FirewallGroupMapInput interface {
	pulumi.Input

	ToFirewallGroupMapOutput() FirewallGroupMapOutput
	ToFirewallGroupMapOutputWithContext(context.Context) FirewallGroupMapOutput
}

type FirewallGroupMap map[string]FirewallGroupInput

func (FirewallGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallGroup)(nil)).Elem()
}

func (i FirewallGroupMap) ToFirewallGroupMapOutput() FirewallGroupMapOutput {
	return i.ToFirewallGroupMapOutputWithContext(context.Background())
}

func (i FirewallGroupMap) ToFirewallGroupMapOutputWithContext(ctx context.Context) FirewallGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallGroupMapOutput)
}

type FirewallGroupOutput struct{ *pulumi.OutputState }

func (FirewallGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallGroup)(nil)).Elem()
}

func (o FirewallGroupOutput) ToFirewallGroupOutput() FirewallGroupOutput {
	return o
}

func (o FirewallGroupOutput) ToFirewallGroupOutputWithContext(ctx context.Context) FirewallGroupOutput {
	return o
}

// The members of the firewall group.
func (o FirewallGroupOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallGroup) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the firewall group.
func (o FirewallGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the site to associate the firewall group with.
func (o FirewallGroupOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallGroup) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

// The type of the firewall group. Must be one of: `address-group`, `port-group`, or `ipv6-address-group`.
func (o FirewallGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type FirewallGroupArrayOutput struct{ *pulumi.OutputState }

func (FirewallGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallGroup)(nil)).Elem()
}

func (o FirewallGroupArrayOutput) ToFirewallGroupArrayOutput() FirewallGroupArrayOutput {
	return o
}

func (o FirewallGroupArrayOutput) ToFirewallGroupArrayOutputWithContext(ctx context.Context) FirewallGroupArrayOutput {
	return o
}

func (o FirewallGroupArrayOutput) Index(i pulumi.IntInput) FirewallGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallGroup {
		return vs[0].([]*FirewallGroup)[vs[1].(int)]
	}).(FirewallGroupOutput)
}

type FirewallGroupMapOutput struct{ *pulumi.OutputState }

func (FirewallGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallGroup)(nil)).Elem()
}

func (o FirewallGroupMapOutput) ToFirewallGroupMapOutput() FirewallGroupMapOutput {
	return o
}

func (o FirewallGroupMapOutput) ToFirewallGroupMapOutputWithContext(ctx context.Context) FirewallGroupMapOutput {
	return o
}

func (o FirewallGroupMapOutput) MapIndex(k pulumi.StringInput) FirewallGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallGroup {
		return vs[0].(map[string]*FirewallGroup)[vs[1].(string)]
	}).(FirewallGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallGroupInput)(nil)).Elem(), &FirewallGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallGroupArrayInput)(nil)).Elem(), FirewallGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallGroupMapInput)(nil)).Elem(), FirewallGroupMap{})
	pulumi.RegisterOutputType(FirewallGroupOutput{})
	pulumi.RegisterOutputType(FirewallGroupArrayOutput{})
	pulumi.RegisterOutputType(FirewallGroupMapOutput{})
}
