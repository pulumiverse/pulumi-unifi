// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/internal"
)

// `firewall.Rule` manages an individual firewall rule on the gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/firewall"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			ipAddress := cfg.Require("ipAddress")
//			_, err := firewall.NewRule(ctx, "drop_all", &firewall.RuleArgs{
//				Name:       pulumi.String("drop all"),
//				Action:     pulumi.String("drop"),
//				Ruleset:    pulumi.String("LAN_IN"),
//				RuleIndex:  pulumi.Int(2011),
//				Protocol:   pulumi.String("all"),
//				DstAddress: pulumi.String(ipAddress),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// import using the ID from the controller API/UI
//
// ```sh
// $ pulumi import unifi:firewall/rule:Rule my_rule 5f7080eb6b8969064f80494f
// ```
type Rule struct {
	pulumi.CustomResourceState

	// The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
	Action pulumi.StringOutput `pulumi:"action"`
	// The destination address of the firewall rule.
	DstAddress pulumi.StringPtrOutput `pulumi:"dstAddress"`
	// The IPv6 destination address of the firewall rule.
	DstAddressIpv6 pulumi.StringPtrOutput `pulumi:"dstAddressIpv6"`
	// The destination firewall group IDs of the firewall rule.
	DstFirewallGroupIds pulumi.StringArrayOutput `pulumi:"dstFirewallGroupIds"`
	// The destination network ID of the firewall rule.
	DstNetworkId pulumi.StringPtrOutput `pulumi:"dstNetworkId"`
	// The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	DstNetworkType pulumi.StringPtrOutput `pulumi:"dstNetworkType"`
	// The destination port of the firewall rule.
	DstPort pulumi.StringPtrOutput `pulumi:"dstPort"`
	// Specifies whether the rule should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// ICMP type name.
	IcmpTypename pulumi.StringPtrOutput `pulumi:"icmpTypename"`
	// ICMPv6 type name.
	IcmpV6Typename pulumi.StringPtrOutput `pulumi:"icmpV6Typename"`
	// Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
	IpSec pulumi.StringPtrOutput `pulumi:"ipSec"`
	// Enable logging for the firewall rule.
	Logging pulumi.BoolPtrOutput `pulumi:"logging"`
	// The name of the firewall rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The protocol of the rule.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// The IPv6 protocol of the rule.
	ProtocolV6 pulumi.StringPtrOutput `pulumi:"protocolV6"`
	// The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
	RuleIndex pulumi.IntOutput `pulumi:"ruleIndex"`
	// The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
	Ruleset pulumi.StringOutput `pulumi:"ruleset"`
	// The name of the site to associate the firewall rule with.
	Site pulumi.StringOutput `pulumi:"site"`
	// The source address for the firewall rule.
	SrcAddress pulumi.StringPtrOutput `pulumi:"srcAddress"`
	// The IPv6 source address for the firewall rule.
	SrcAddressIpv6 pulumi.StringPtrOutput `pulumi:"srcAddressIpv6"`
	// The source firewall group IDs for the firewall rule.
	SrcFirewallGroupIds pulumi.StringArrayOutput `pulumi:"srcFirewallGroupIds"`
	// The source MAC address of the firewall rule.
	SrcMac pulumi.StringPtrOutput `pulumi:"srcMac"`
	// The source network ID for the firewall rule.
	SrcNetworkId pulumi.StringPtrOutput `pulumi:"srcNetworkId"`
	// The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	SrcNetworkType pulumi.StringPtrOutput `pulumi:"srcNetworkType"`
	// The source port of the firewall rule.
	SrcPort pulumi.StringPtrOutput `pulumi:"srcPort"`
	// Match where the state is established.
	StateEstablished pulumi.BoolPtrOutput `pulumi:"stateEstablished"`
	// Match where the state is invalid.
	StateInvalid pulumi.BoolPtrOutput `pulumi:"stateInvalid"`
	// Match where the state is new.
	StateNew pulumi.BoolPtrOutput `pulumi:"stateNew"`
	// Match where the state is related.
	StateRelated pulumi.BoolPtrOutput `pulumi:"stateRelated"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.RuleIndex == nil {
		return nil, errors.New("invalid value for required argument 'RuleIndex'")
	}
	if args.Ruleset == nil {
		return nil, errors.New("invalid value for required argument 'Ruleset'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rule
	err := ctx.RegisterResource("unifi:firewall/rule:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("unifi:firewall/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
	Action *string `pulumi:"action"`
	// The destination address of the firewall rule.
	DstAddress *string `pulumi:"dstAddress"`
	// The IPv6 destination address of the firewall rule.
	DstAddressIpv6 *string `pulumi:"dstAddressIpv6"`
	// The destination firewall group IDs of the firewall rule.
	DstFirewallGroupIds []string `pulumi:"dstFirewallGroupIds"`
	// The destination network ID of the firewall rule.
	DstNetworkId *string `pulumi:"dstNetworkId"`
	// The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	DstNetworkType *string `pulumi:"dstNetworkType"`
	// The destination port of the firewall rule.
	DstPort *string `pulumi:"dstPort"`
	// Specifies whether the rule should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// ICMP type name.
	IcmpTypename *string `pulumi:"icmpTypename"`
	// ICMPv6 type name.
	IcmpV6Typename *string `pulumi:"icmpV6Typename"`
	// Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
	IpSec *string `pulumi:"ipSec"`
	// Enable logging for the firewall rule.
	Logging *bool `pulumi:"logging"`
	// The name of the firewall rule.
	Name *string `pulumi:"name"`
	// The protocol of the rule.
	Protocol *string `pulumi:"protocol"`
	// The IPv6 protocol of the rule.
	ProtocolV6 *string `pulumi:"protocolV6"`
	// The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
	RuleIndex *int `pulumi:"ruleIndex"`
	// The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
	Ruleset *string `pulumi:"ruleset"`
	// The name of the site to associate the firewall rule with.
	Site *string `pulumi:"site"`
	// The source address for the firewall rule.
	SrcAddress *string `pulumi:"srcAddress"`
	// The IPv6 source address for the firewall rule.
	SrcAddressIpv6 *string `pulumi:"srcAddressIpv6"`
	// The source firewall group IDs for the firewall rule.
	SrcFirewallGroupIds []string `pulumi:"srcFirewallGroupIds"`
	// The source MAC address of the firewall rule.
	SrcMac *string `pulumi:"srcMac"`
	// The source network ID for the firewall rule.
	SrcNetworkId *string `pulumi:"srcNetworkId"`
	// The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	SrcNetworkType *string `pulumi:"srcNetworkType"`
	// The source port of the firewall rule.
	SrcPort *string `pulumi:"srcPort"`
	// Match where the state is established.
	StateEstablished *bool `pulumi:"stateEstablished"`
	// Match where the state is invalid.
	StateInvalid *bool `pulumi:"stateInvalid"`
	// Match where the state is new.
	StateNew *bool `pulumi:"stateNew"`
	// Match where the state is related.
	StateRelated *bool `pulumi:"stateRelated"`
}

type RuleState struct {
	// The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
	Action pulumi.StringPtrInput
	// The destination address of the firewall rule.
	DstAddress pulumi.StringPtrInput
	// The IPv6 destination address of the firewall rule.
	DstAddressIpv6 pulumi.StringPtrInput
	// The destination firewall group IDs of the firewall rule.
	DstFirewallGroupIds pulumi.StringArrayInput
	// The destination network ID of the firewall rule.
	DstNetworkId pulumi.StringPtrInput
	// The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	DstNetworkType pulumi.StringPtrInput
	// The destination port of the firewall rule.
	DstPort pulumi.StringPtrInput
	// Specifies whether the rule should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// ICMP type name.
	IcmpTypename pulumi.StringPtrInput
	// ICMPv6 type name.
	IcmpV6Typename pulumi.StringPtrInput
	// Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
	IpSec pulumi.StringPtrInput
	// Enable logging for the firewall rule.
	Logging pulumi.BoolPtrInput
	// The name of the firewall rule.
	Name pulumi.StringPtrInput
	// The protocol of the rule.
	Protocol pulumi.StringPtrInput
	// The IPv6 protocol of the rule.
	ProtocolV6 pulumi.StringPtrInput
	// The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
	RuleIndex pulumi.IntPtrInput
	// The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
	Ruleset pulumi.StringPtrInput
	// The name of the site to associate the firewall rule with.
	Site pulumi.StringPtrInput
	// The source address for the firewall rule.
	SrcAddress pulumi.StringPtrInput
	// The IPv6 source address for the firewall rule.
	SrcAddressIpv6 pulumi.StringPtrInput
	// The source firewall group IDs for the firewall rule.
	SrcFirewallGroupIds pulumi.StringArrayInput
	// The source MAC address of the firewall rule.
	SrcMac pulumi.StringPtrInput
	// The source network ID for the firewall rule.
	SrcNetworkId pulumi.StringPtrInput
	// The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	SrcNetworkType pulumi.StringPtrInput
	// The source port of the firewall rule.
	SrcPort pulumi.StringPtrInput
	// Match where the state is established.
	StateEstablished pulumi.BoolPtrInput
	// Match where the state is invalid.
	StateInvalid pulumi.BoolPtrInput
	// Match where the state is new.
	StateNew pulumi.BoolPtrInput
	// Match where the state is related.
	StateRelated pulumi.BoolPtrInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
	Action string `pulumi:"action"`
	// The destination address of the firewall rule.
	DstAddress *string `pulumi:"dstAddress"`
	// The IPv6 destination address of the firewall rule.
	DstAddressIpv6 *string `pulumi:"dstAddressIpv6"`
	// The destination firewall group IDs of the firewall rule.
	DstFirewallGroupIds []string `pulumi:"dstFirewallGroupIds"`
	// The destination network ID of the firewall rule.
	DstNetworkId *string `pulumi:"dstNetworkId"`
	// The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	DstNetworkType *string `pulumi:"dstNetworkType"`
	// The destination port of the firewall rule.
	DstPort *string `pulumi:"dstPort"`
	// Specifies whether the rule should be enabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// ICMP type name.
	IcmpTypename *string `pulumi:"icmpTypename"`
	// ICMPv6 type name.
	IcmpV6Typename *string `pulumi:"icmpV6Typename"`
	// Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
	IpSec *string `pulumi:"ipSec"`
	// Enable logging for the firewall rule.
	Logging *bool `pulumi:"logging"`
	// The name of the firewall rule.
	Name *string `pulumi:"name"`
	// The protocol of the rule.
	Protocol *string `pulumi:"protocol"`
	// The IPv6 protocol of the rule.
	ProtocolV6 *string `pulumi:"protocolV6"`
	// The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
	RuleIndex int `pulumi:"ruleIndex"`
	// The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
	Ruleset string `pulumi:"ruleset"`
	// The name of the site to associate the firewall rule with.
	Site *string `pulumi:"site"`
	// The source address for the firewall rule.
	SrcAddress *string `pulumi:"srcAddress"`
	// The IPv6 source address for the firewall rule.
	SrcAddressIpv6 *string `pulumi:"srcAddressIpv6"`
	// The source firewall group IDs for the firewall rule.
	SrcFirewallGroupIds []string `pulumi:"srcFirewallGroupIds"`
	// The source MAC address of the firewall rule.
	SrcMac *string `pulumi:"srcMac"`
	// The source network ID for the firewall rule.
	SrcNetworkId *string `pulumi:"srcNetworkId"`
	// The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	SrcNetworkType *string `pulumi:"srcNetworkType"`
	// The source port of the firewall rule.
	SrcPort *string `pulumi:"srcPort"`
	// Match where the state is established.
	StateEstablished *bool `pulumi:"stateEstablished"`
	// Match where the state is invalid.
	StateInvalid *bool `pulumi:"stateInvalid"`
	// Match where the state is new.
	StateNew *bool `pulumi:"stateNew"`
	// Match where the state is related.
	StateRelated *bool `pulumi:"stateRelated"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
	Action pulumi.StringInput
	// The destination address of the firewall rule.
	DstAddress pulumi.StringPtrInput
	// The IPv6 destination address of the firewall rule.
	DstAddressIpv6 pulumi.StringPtrInput
	// The destination firewall group IDs of the firewall rule.
	DstFirewallGroupIds pulumi.StringArrayInput
	// The destination network ID of the firewall rule.
	DstNetworkId pulumi.StringPtrInput
	// The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	DstNetworkType pulumi.StringPtrInput
	// The destination port of the firewall rule.
	DstPort pulumi.StringPtrInput
	// Specifies whether the rule should be enabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// ICMP type name.
	IcmpTypename pulumi.StringPtrInput
	// ICMPv6 type name.
	IcmpV6Typename pulumi.StringPtrInput
	// Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
	IpSec pulumi.StringPtrInput
	// Enable logging for the firewall rule.
	Logging pulumi.BoolPtrInput
	// The name of the firewall rule.
	Name pulumi.StringPtrInput
	// The protocol of the rule.
	Protocol pulumi.StringPtrInput
	// The IPv6 protocol of the rule.
	ProtocolV6 pulumi.StringPtrInput
	// The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
	RuleIndex pulumi.IntInput
	// The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
	Ruleset pulumi.StringInput
	// The name of the site to associate the firewall rule with.
	Site pulumi.StringPtrInput
	// The source address for the firewall rule.
	SrcAddress pulumi.StringPtrInput
	// The IPv6 source address for the firewall rule.
	SrcAddressIpv6 pulumi.StringPtrInput
	// The source firewall group IDs for the firewall rule.
	SrcFirewallGroupIds pulumi.StringArrayInput
	// The source MAC address of the firewall rule.
	SrcMac pulumi.StringPtrInput
	// The source network ID for the firewall rule.
	SrcNetworkId pulumi.StringPtrInput
	// The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
	SrcNetworkType pulumi.StringPtrInput
	// The source port of the firewall rule.
	SrcPort pulumi.StringPtrInput
	// Match where the state is established.
	StateEstablished pulumi.BoolPtrInput
	// Match where the state is invalid.
	StateInvalid pulumi.BoolPtrInput
	// Match where the state is new.
	StateNew pulumi.BoolPtrInput
	// Match where the state is related.
	StateRelated pulumi.BoolPtrInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

type RuleInput interface {
	pulumi.Input

	ToRuleOutput() RuleOutput
	ToRuleOutputWithContext(ctx context.Context) RuleOutput
}

func (*Rule) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

// RuleArrayInput is an input type that accepts RuleArray and RuleArrayOutput values.
// You can construct a concrete instance of `RuleArrayInput` via:
//
//	RuleArray{ RuleArgs{...} }
type RuleArrayInput interface {
	pulumi.Input

	ToRuleArrayOutput() RuleArrayOutput
	ToRuleArrayOutputWithContext(context.Context) RuleArrayOutput
}

type RuleArray []RuleInput

func (RuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (i RuleArray) ToRuleArrayOutput() RuleArrayOutput {
	return i.ToRuleArrayOutputWithContext(context.Background())
}

func (i RuleArray) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleArrayOutput)
}

// RuleMapInput is an input type that accepts RuleMap and RuleMapOutput values.
// You can construct a concrete instance of `RuleMapInput` via:
//
//	RuleMap{ "key": RuleArgs{...} }
type RuleMapInput interface {
	pulumi.Input

	ToRuleMapOutput() RuleMapOutput
	ToRuleMapOutputWithContext(context.Context) RuleMapOutput
}

type RuleMap map[string]RuleInput

func (RuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (i RuleMap) ToRuleMapOutput() RuleMapOutput {
	return i.ToRuleMapOutputWithContext(context.Background())
}

func (i RuleMap) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMapOutput)
}

type RuleOutput struct{ *pulumi.OutputState }

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

// The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
func (o RuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// The destination address of the firewall rule.
func (o RuleOutput) DstAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.DstAddress }).(pulumi.StringPtrOutput)
}

// The IPv6 destination address of the firewall rule.
func (o RuleOutput) DstAddressIpv6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.DstAddressIpv6 }).(pulumi.StringPtrOutput)
}

// The destination firewall group IDs of the firewall rule.
func (o RuleOutput) DstFirewallGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringArrayOutput { return v.DstFirewallGroupIds }).(pulumi.StringArrayOutput)
}

// The destination network ID of the firewall rule.
func (o RuleOutput) DstNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.DstNetworkId }).(pulumi.StringPtrOutput)
}

// The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
func (o RuleOutput) DstNetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.DstNetworkType }).(pulumi.StringPtrOutput)
}

// The destination port of the firewall rule.
func (o RuleOutput) DstPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.DstPort }).(pulumi.StringPtrOutput)
}

// Specifies whether the rule should be enabled. Defaults to `true`.
func (o RuleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// ICMP type name.
func (o RuleOutput) IcmpTypename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.IcmpTypename }).(pulumi.StringPtrOutput)
}

// ICMPv6 type name.
func (o RuleOutput) IcmpV6Typename() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.IcmpV6Typename }).(pulumi.StringPtrOutput)
}

// Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
func (o RuleOutput) IpSec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.IpSec }).(pulumi.StringPtrOutput)
}

// Enable logging for the firewall rule.
func (o RuleOutput) Logging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.BoolPtrOutput { return v.Logging }).(pulumi.BoolPtrOutput)
}

// The name of the firewall rule.
func (o RuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The protocol of the rule.
func (o RuleOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The IPv6 protocol of the rule.
func (o RuleOutput) ProtocolV6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.ProtocolV6 }).(pulumi.StringPtrOutput)
}

// The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
func (o RuleOutput) RuleIndex() pulumi.IntOutput {
	return o.ApplyT(func(v *Rule) pulumi.IntOutput { return v.RuleIndex }).(pulumi.IntOutput)
}

// The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
func (o RuleOutput) Ruleset() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Ruleset }).(pulumi.StringOutput)
}

// The name of the site to associate the firewall rule with.
func (o RuleOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

// The source address for the firewall rule.
func (o RuleOutput) SrcAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.SrcAddress }).(pulumi.StringPtrOutput)
}

// The IPv6 source address for the firewall rule.
func (o RuleOutput) SrcAddressIpv6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.SrcAddressIpv6 }).(pulumi.StringPtrOutput)
}

// The source firewall group IDs for the firewall rule.
func (o RuleOutput) SrcFirewallGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringArrayOutput { return v.SrcFirewallGroupIds }).(pulumi.StringArrayOutput)
}

// The source MAC address of the firewall rule.
func (o RuleOutput) SrcMac() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.SrcMac }).(pulumi.StringPtrOutput)
}

// The source network ID for the firewall rule.
func (o RuleOutput) SrcNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.SrcNetworkId }).(pulumi.StringPtrOutput)
}

// The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
func (o RuleOutput) SrcNetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.SrcNetworkType }).(pulumi.StringPtrOutput)
}

// The source port of the firewall rule.
func (o RuleOutput) SrcPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.SrcPort }).(pulumi.StringPtrOutput)
}

// Match where the state is established.
func (o RuleOutput) StateEstablished() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.BoolPtrOutput { return v.StateEstablished }).(pulumi.BoolPtrOutput)
}

// Match where the state is invalid.
func (o RuleOutput) StateInvalid() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.BoolPtrOutput { return v.StateInvalid }).(pulumi.BoolPtrOutput)
}

// Match where the state is new.
func (o RuleOutput) StateNew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.BoolPtrOutput { return v.StateNew }).(pulumi.BoolPtrOutput)
}

// Match where the state is related.
func (o RuleOutput) StateRelated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.BoolPtrOutput { return v.StateRelated }).(pulumi.BoolPtrOutput)
}

type RuleArrayOutput struct{ *pulumi.OutputState }

func (RuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (o RuleArrayOutput) ToRuleArrayOutput() RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) Index(i pulumi.IntInput) RuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].([]*Rule)[vs[1].(int)]
	}).(RuleOutput)
}

type RuleMapOutput struct{ *pulumi.OutputState }

func (RuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (o RuleMapOutput) ToRuleMapOutput() RuleMapOutput {
	return o
}

func (o RuleMapOutput) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return o
}

func (o RuleMapOutput) MapIndex(k pulumi.StringInput) RuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].(map[string]*Rule)[vs[1].(string)]
	}).(RuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleInput)(nil)).Elem(), &Rule{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleArrayInput)(nil)).Elem(), RuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleMapInput)(nil)).Elem(), RuleMap{})
	pulumi.RegisterOutputType(RuleOutput{})
	pulumi.RegisterOutputType(RuleArrayOutput{})
	pulumi.RegisterOutputType(RuleMapOutput{})
}
