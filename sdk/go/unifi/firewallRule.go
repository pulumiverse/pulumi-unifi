// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package unifi

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `FirewallRule` manages an individual firewall rule on the gateway.
//
// ## Import
//
// # import using the ID from the controller API/UI
//
// ```sh
//  $ pulumi import unifi:index/firewallRule:FirewallRule my_rule 5f7080eb6b8969064f80494f
// ```
type FirewallRule struct {
	pulumi.CustomResourceState

	// The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
	Action pulumi.StringOutput `pulumi:"action"`
	// The destination address of the firewall rule.
	DstAddress pulumi.StringPtrOutput `pulumi:"dstAddress"`
	// The destination firewall group IDs of the firewall rule.
	DstFirewallGroupIds pulumi.StringArrayOutput `pulumi:"dstFirewallGroupIds"`
	// The destination network ID of the firewall rule.
	DstNetworkId pulumi.StringPtrOutput `pulumi:"dstNetworkId"`
	// The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	DstNetworkType pulumi.StringPtrOutput `pulumi:"dstNetworkType"`
	// The destination port of the firewall rule.
	DstPort pulumi.StringPtrOutput `pulumi:"dstPort"`
	// ICMP type name.
	IcmpTypename pulumi.StringPtrOutput `pulumi:"icmpTypename"`
	// Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
	IpSec pulumi.StringPtrOutput `pulumi:"ipSec"`
	// Enable logging for the firewall rule.
	Logging pulumi.BoolPtrOutput `pulumi:"logging"`
	// The name of the firewall rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The protocol of the rule.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
	RuleIndex pulumi.IntOutput `pulumi:"ruleIndex"`
	// The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
	Ruleset pulumi.StringOutput `pulumi:"ruleset"`
	// The name of the site to associate the firewall rule with.
	Site pulumi.StringOutput `pulumi:"site"`
	// The source address for the firewall rule.
	SrcAddress pulumi.StringPtrOutput `pulumi:"srcAddress"`
	// The source firewall group IDs for the firewall rule.
	SrcFirewallGroupIds pulumi.StringArrayOutput `pulumi:"srcFirewallGroupIds"`
	// The source MAC address of the firewall rule.
	SrcMac pulumi.StringPtrOutput `pulumi:"srcMac"`
	// The source network ID for the firewall rule.
	SrcNetworkId pulumi.StringPtrOutput `pulumi:"srcNetworkId"`
	// The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	SrcNetworkType pulumi.StringPtrOutput `pulumi:"srcNetworkType"`
	// Match where the state is established.
	StateEstablished pulumi.BoolPtrOutput `pulumi:"stateEstablished"`
	// Match where the state is invalid.
	StateInvalid pulumi.BoolPtrOutput `pulumi:"stateInvalid"`
	// Match where the state is new.
	StateNew pulumi.BoolPtrOutput `pulumi:"stateNew"`
	// Match where the state is related.
	StateRelated pulumi.BoolPtrOutput `pulumi:"stateRelated"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.RuleIndex == nil {
		return nil, errors.New("invalid value for required argument 'RuleIndex'")
	}
	if args.Ruleset == nil {
		return nil, errors.New("invalid value for required argument 'Ruleset'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallRule
	err := ctx.RegisterResource("unifi:index/firewallRule:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("unifi:index/firewallRule:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
	Action *string `pulumi:"action"`
	// The destination address of the firewall rule.
	DstAddress *string `pulumi:"dstAddress"`
	// The destination firewall group IDs of the firewall rule.
	DstFirewallGroupIds []string `pulumi:"dstFirewallGroupIds"`
	// The destination network ID of the firewall rule.
	DstNetworkId *string `pulumi:"dstNetworkId"`
	// The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	DstNetworkType *string `pulumi:"dstNetworkType"`
	// The destination port of the firewall rule.
	DstPort *string `pulumi:"dstPort"`
	// ICMP type name.
	IcmpTypename *string `pulumi:"icmpTypename"`
	// Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
	IpSec *string `pulumi:"ipSec"`
	// Enable logging for the firewall rule.
	Logging *bool `pulumi:"logging"`
	// The name of the firewall rule.
	Name *string `pulumi:"name"`
	// The protocol of the rule.
	Protocol *string `pulumi:"protocol"`
	// The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
	RuleIndex *int `pulumi:"ruleIndex"`
	// The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
	Ruleset *string `pulumi:"ruleset"`
	// The name of the site to associate the firewall rule with.
	Site *string `pulumi:"site"`
	// The source address for the firewall rule.
	SrcAddress *string `pulumi:"srcAddress"`
	// The source firewall group IDs for the firewall rule.
	SrcFirewallGroupIds []string `pulumi:"srcFirewallGroupIds"`
	// The source MAC address of the firewall rule.
	SrcMac *string `pulumi:"srcMac"`
	// The source network ID for the firewall rule.
	SrcNetworkId *string `pulumi:"srcNetworkId"`
	// The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	SrcNetworkType *string `pulumi:"srcNetworkType"`
	// Match where the state is established.
	StateEstablished *bool `pulumi:"stateEstablished"`
	// Match where the state is invalid.
	StateInvalid *bool `pulumi:"stateInvalid"`
	// Match where the state is new.
	StateNew *bool `pulumi:"stateNew"`
	// Match where the state is related.
	StateRelated *bool `pulumi:"stateRelated"`
}

type FirewallRuleState struct {
	// The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
	Action pulumi.StringPtrInput
	// The destination address of the firewall rule.
	DstAddress pulumi.StringPtrInput
	// The destination firewall group IDs of the firewall rule.
	DstFirewallGroupIds pulumi.StringArrayInput
	// The destination network ID of the firewall rule.
	DstNetworkId pulumi.StringPtrInput
	// The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	DstNetworkType pulumi.StringPtrInput
	// The destination port of the firewall rule.
	DstPort pulumi.StringPtrInput
	// ICMP type name.
	IcmpTypename pulumi.StringPtrInput
	// Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
	IpSec pulumi.StringPtrInput
	// Enable logging for the firewall rule.
	Logging pulumi.BoolPtrInput
	// The name of the firewall rule.
	Name pulumi.StringPtrInput
	// The protocol of the rule.
	Protocol pulumi.StringPtrInput
	// The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
	RuleIndex pulumi.IntPtrInput
	// The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
	Ruleset pulumi.StringPtrInput
	// The name of the site to associate the firewall rule with.
	Site pulumi.StringPtrInput
	// The source address for the firewall rule.
	SrcAddress pulumi.StringPtrInput
	// The source firewall group IDs for the firewall rule.
	SrcFirewallGroupIds pulumi.StringArrayInput
	// The source MAC address of the firewall rule.
	SrcMac pulumi.StringPtrInput
	// The source network ID for the firewall rule.
	SrcNetworkId pulumi.StringPtrInput
	// The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	SrcNetworkType pulumi.StringPtrInput
	// Match where the state is established.
	StateEstablished pulumi.BoolPtrInput
	// Match where the state is invalid.
	StateInvalid pulumi.BoolPtrInput
	// Match where the state is new.
	StateNew pulumi.BoolPtrInput
	// Match where the state is related.
	StateRelated pulumi.BoolPtrInput
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
	Action string `pulumi:"action"`
	// The destination address of the firewall rule.
	DstAddress *string `pulumi:"dstAddress"`
	// The destination firewall group IDs of the firewall rule.
	DstFirewallGroupIds []string `pulumi:"dstFirewallGroupIds"`
	// The destination network ID of the firewall rule.
	DstNetworkId *string `pulumi:"dstNetworkId"`
	// The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	DstNetworkType *string `pulumi:"dstNetworkType"`
	// The destination port of the firewall rule.
	DstPort *string `pulumi:"dstPort"`
	// ICMP type name.
	IcmpTypename *string `pulumi:"icmpTypename"`
	// Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
	IpSec *string `pulumi:"ipSec"`
	// Enable logging for the firewall rule.
	Logging *bool `pulumi:"logging"`
	// The name of the firewall rule.
	Name *string `pulumi:"name"`
	// The protocol of the rule.
	Protocol string `pulumi:"protocol"`
	// The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
	RuleIndex int `pulumi:"ruleIndex"`
	// The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
	Ruleset string `pulumi:"ruleset"`
	// The name of the site to associate the firewall rule with.
	Site *string `pulumi:"site"`
	// The source address for the firewall rule.
	SrcAddress *string `pulumi:"srcAddress"`
	// The source firewall group IDs for the firewall rule.
	SrcFirewallGroupIds []string `pulumi:"srcFirewallGroupIds"`
	// The source MAC address of the firewall rule.
	SrcMac *string `pulumi:"srcMac"`
	// The source network ID for the firewall rule.
	SrcNetworkId *string `pulumi:"srcNetworkId"`
	// The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	SrcNetworkType *string `pulumi:"srcNetworkType"`
	// Match where the state is established.
	StateEstablished *bool `pulumi:"stateEstablished"`
	// Match where the state is invalid.
	StateInvalid *bool `pulumi:"stateInvalid"`
	// Match where the state is new.
	StateNew *bool `pulumi:"stateNew"`
	// Match where the state is related.
	StateRelated *bool `pulumi:"stateRelated"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
	Action pulumi.StringInput
	// The destination address of the firewall rule.
	DstAddress pulumi.StringPtrInput
	// The destination firewall group IDs of the firewall rule.
	DstFirewallGroupIds pulumi.StringArrayInput
	// The destination network ID of the firewall rule.
	DstNetworkId pulumi.StringPtrInput
	// The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	DstNetworkType pulumi.StringPtrInput
	// The destination port of the firewall rule.
	DstPort pulumi.StringPtrInput
	// ICMP type name.
	IcmpTypename pulumi.StringPtrInput
	// Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
	IpSec pulumi.StringPtrInput
	// Enable logging for the firewall rule.
	Logging pulumi.BoolPtrInput
	// The name of the firewall rule.
	Name pulumi.StringPtrInput
	// The protocol of the rule.
	Protocol pulumi.StringInput
	// The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
	RuleIndex pulumi.IntInput
	// The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
	Ruleset pulumi.StringInput
	// The name of the site to associate the firewall rule with.
	Site pulumi.StringPtrInput
	// The source address for the firewall rule.
	SrcAddress pulumi.StringPtrInput
	// The source firewall group IDs for the firewall rule.
	SrcFirewallGroupIds pulumi.StringArrayInput
	// The source MAC address of the firewall rule.
	SrcMac pulumi.StringPtrInput
	// The source network ID for the firewall rule.
	SrcNetworkId pulumi.StringPtrInput
	// The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	SrcNetworkType pulumi.StringPtrInput
	// Match where the state is established.
	StateEstablished pulumi.BoolPtrInput
	// Match where the state is invalid.
	StateInvalid pulumi.BoolPtrInput
	// Match where the state is new.
	StateNew pulumi.BoolPtrInput
	// Match where the state is related.
	StateRelated pulumi.BoolPtrInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}

type FirewallRuleInput interface {
	pulumi.Input

	ToFirewallRuleOutput() FirewallRuleOutput
	ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput
}

func (*FirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (i *FirewallRule) ToFirewallRuleOutput() FirewallRuleOutput {
	return i.ToFirewallRuleOutputWithContext(context.Background())
}

func (i *FirewallRule) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleOutput)
}

// FirewallRuleArrayInput is an input type that accepts FirewallRuleArray and FirewallRuleArrayOutput values.
// You can construct a concrete instance of `FirewallRuleArrayInput` via:
//
//          FirewallRuleArray{ FirewallRuleArgs{...} }
type FirewallRuleArrayInput interface {
	pulumi.Input

	ToFirewallRuleArrayOutput() FirewallRuleArrayOutput
	ToFirewallRuleArrayOutputWithContext(context.Context) FirewallRuleArrayOutput
}

type FirewallRuleArray []FirewallRuleInput

func (FirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRule)(nil)).Elem()
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return i.ToFirewallRuleArrayOutputWithContext(context.Background())
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleArrayOutput)
}

// FirewallRuleMapInput is an input type that accepts FirewallRuleMap and FirewallRuleMapOutput values.
// You can construct a concrete instance of `FirewallRuleMapInput` via:
//
//          FirewallRuleMap{ "key": FirewallRuleArgs{...} }
type FirewallRuleMapInput interface {
	pulumi.Input

	ToFirewallRuleMapOutput() FirewallRuleMapOutput
	ToFirewallRuleMapOutputWithContext(context.Context) FirewallRuleMapOutput
}

type FirewallRuleMap map[string]FirewallRuleInput

func (FirewallRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRule)(nil)).Elem()
}

func (i FirewallRuleMap) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return i.ToFirewallRuleMapOutputWithContext(context.Background())
}

func (i FirewallRuleMap) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleMapOutput)
}

type FirewallRuleOutput struct{ *pulumi.OutputState }

func (FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (o FirewallRuleOutput) ToFirewallRuleOutput() FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return o
}

// The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
func (o FirewallRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// The destination address of the firewall rule.
func (o FirewallRuleOutput) DstAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.DstAddress }).(pulumi.StringPtrOutput)
}

// The destination firewall group IDs of the firewall rule.
func (o FirewallRuleOutput) DstFirewallGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringArrayOutput { return v.DstFirewallGroupIds }).(pulumi.StringArrayOutput)
}

// The destination network ID of the firewall rule.
func (o FirewallRuleOutput) DstNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.DstNetworkId }).(pulumi.StringPtrOutput)
}

// The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
func (o FirewallRuleOutput) DstNetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.DstNetworkType }).(pulumi.StringPtrOutput)
}

// The destination port of the firewall rule.
func (o FirewallRuleOutput) DstPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.DstPort }).(pulumi.StringPtrOutput)
}

// ICMP type name.
func (o FirewallRuleOutput) IcmpTypename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.IcmpTypename }).(pulumi.StringPtrOutput)
}

// Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
func (o FirewallRuleOutput) IpSec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.IpSec }).(pulumi.StringPtrOutput)
}

// Enable logging for the firewall rule.
func (o FirewallRuleOutput) Logging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.BoolPtrOutput { return v.Logging }).(pulumi.BoolPtrOutput)
}

// The name of the firewall rule.
func (o FirewallRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The protocol of the rule.
func (o FirewallRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
func (o FirewallRuleOutput) RuleIndex() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.IntOutput { return v.RuleIndex }).(pulumi.IntOutput)
}

// The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
func (o FirewallRuleOutput) Ruleset() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Ruleset }).(pulumi.StringOutput)
}

// The name of the site to associate the firewall rule with.
func (o FirewallRuleOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

// The source address for the firewall rule.
func (o FirewallRuleOutput) SrcAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.SrcAddress }).(pulumi.StringPtrOutput)
}

// The source firewall group IDs for the firewall rule.
func (o FirewallRuleOutput) SrcFirewallGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringArrayOutput { return v.SrcFirewallGroupIds }).(pulumi.StringArrayOutput)
}

// The source MAC address of the firewall rule.
func (o FirewallRuleOutput) SrcMac() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.SrcMac }).(pulumi.StringPtrOutput)
}

// The source network ID for the firewall rule.
func (o FirewallRuleOutput) SrcNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.SrcNetworkId }).(pulumi.StringPtrOutput)
}

// The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
func (o FirewallRuleOutput) SrcNetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringPtrOutput { return v.SrcNetworkType }).(pulumi.StringPtrOutput)
}

// Match where the state is established.
func (o FirewallRuleOutput) StateEstablished() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.BoolPtrOutput { return v.StateEstablished }).(pulumi.BoolPtrOutput)
}

// Match where the state is invalid.
func (o FirewallRuleOutput) StateInvalid() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.BoolPtrOutput { return v.StateInvalid }).(pulumi.BoolPtrOutput)
}

// Match where the state is new.
func (o FirewallRuleOutput) StateNew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.BoolPtrOutput { return v.StateNew }).(pulumi.BoolPtrOutput)
}

// Match where the state is related.
func (o FirewallRuleOutput) StateRelated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.BoolPtrOutput { return v.StateRelated }).(pulumi.BoolPtrOutput)
}

type FirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (FirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRule)(nil)).Elem()
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) Index(i pulumi.IntInput) FirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallRule {
		return vs[0].([]*FirewallRule)[vs[1].(int)]
	}).(FirewallRuleOutput)
}

type FirewallRuleMapOutput struct{ *pulumi.OutputState }

func (FirewallRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRule)(nil)).Elem()
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) MapIndex(k pulumi.StringInput) FirewallRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallRule {
		return vs[0].(map[string]*FirewallRule)[vs[1].(string)]
	}).(FirewallRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleInput)(nil)).Elem(), &FirewallRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleArrayInput)(nil)).Elem(), FirewallRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleMapInput)(nil)).Elem(), FirewallRuleMap{})
	pulumi.RegisterOutputType(FirewallRuleOutput{})
	pulumi.RegisterOutputType(FirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(FirewallRuleMapOutput{})
}
