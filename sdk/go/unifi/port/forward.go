// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package port

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/internal"
)

// `port.Forward` manages a port forwarding rule on the gateway.
type Forward struct {
	pulumi.CustomResourceState

	// The destination port for the forwarding.
	DstPort pulumi.StringPtrOutput `pulumi:"dstPort"`
	// Specifies whether the port forwarding rule is enabled or not. Defaults to `true`. This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
	//
	// Deprecated: This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The IPv4 address to forward traffic to.
	FwdIp pulumi.StringPtrOutput `pulumi:"fwdIp"`
	// The port to forward traffic to.
	FwdPort pulumi.StringPtrOutput `pulumi:"fwdPort"`
	// Specifies whether to log forwarded traffic or not. Defaults to `false`.
	Log pulumi.BoolPtrOutput `pulumi:"log"`
	// The name of the port forwarding rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The port forwarding interface. Can be `wan`, `wan2`, or `both`.
	PortForwardInterface pulumi.StringPtrOutput `pulumi:"portForwardInterface"`
	// The protocol for the port forwarding rule. Can be `tcp`, `udp`, or `tcpUdp`. Defaults to `tcpUdp`.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// The name of the site to associate the port forwarding rule with.
	Site pulumi.StringOutput `pulumi:"site"`
	// The source IPv4 address (or CIDR) of the port forwarding rule. For all traffic, specify `any`. Defaults to `any`.
	SrcIp pulumi.StringPtrOutput `pulumi:"srcIp"`
}

// NewForward registers a new resource with the given unique name, arguments, and options.
func NewForward(ctx *pulumi.Context,
	name string, args *ForwardArgs, opts ...pulumi.ResourceOption) (*Forward, error) {
	if args == nil {
		args = &ForwardArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Forward
	err := ctx.RegisterResource("unifi:port/forward:Forward", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetForward gets an existing Forward resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetForward(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ForwardState, opts ...pulumi.ResourceOption) (*Forward, error) {
	var resource Forward
	err := ctx.ReadResource("unifi:port/forward:Forward", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Forward resources.
type forwardState struct {
	// The destination port for the forwarding.
	DstPort *string `pulumi:"dstPort"`
	// Specifies whether the port forwarding rule is enabled or not. Defaults to `true`. This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
	//
	// Deprecated: This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
	Enabled *bool `pulumi:"enabled"`
	// The IPv4 address to forward traffic to.
	FwdIp *string `pulumi:"fwdIp"`
	// The port to forward traffic to.
	FwdPort *string `pulumi:"fwdPort"`
	// Specifies whether to log forwarded traffic or not. Defaults to `false`.
	Log *bool `pulumi:"log"`
	// The name of the port forwarding rule.
	Name *string `pulumi:"name"`
	// The port forwarding interface. Can be `wan`, `wan2`, or `both`.
	PortForwardInterface *string `pulumi:"portForwardInterface"`
	// The protocol for the port forwarding rule. Can be `tcp`, `udp`, or `tcpUdp`. Defaults to `tcpUdp`.
	Protocol *string `pulumi:"protocol"`
	// The name of the site to associate the port forwarding rule with.
	Site *string `pulumi:"site"`
	// The source IPv4 address (or CIDR) of the port forwarding rule. For all traffic, specify `any`. Defaults to `any`.
	SrcIp *string `pulumi:"srcIp"`
}

type ForwardState struct {
	// The destination port for the forwarding.
	DstPort pulumi.StringPtrInput
	// Specifies whether the port forwarding rule is enabled or not. Defaults to `true`. This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
	//
	// Deprecated: This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
	Enabled pulumi.BoolPtrInput
	// The IPv4 address to forward traffic to.
	FwdIp pulumi.StringPtrInput
	// The port to forward traffic to.
	FwdPort pulumi.StringPtrInput
	// Specifies whether to log forwarded traffic or not. Defaults to `false`.
	Log pulumi.BoolPtrInput
	// The name of the port forwarding rule.
	Name pulumi.StringPtrInput
	// The port forwarding interface. Can be `wan`, `wan2`, or `both`.
	PortForwardInterface pulumi.StringPtrInput
	// The protocol for the port forwarding rule. Can be `tcp`, `udp`, or `tcpUdp`. Defaults to `tcpUdp`.
	Protocol pulumi.StringPtrInput
	// The name of the site to associate the port forwarding rule with.
	Site pulumi.StringPtrInput
	// The source IPv4 address (or CIDR) of the port forwarding rule. For all traffic, specify `any`. Defaults to `any`.
	SrcIp pulumi.StringPtrInput
}

func (ForwardState) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardState)(nil)).Elem()
}

type forwardArgs struct {
	// The destination port for the forwarding.
	DstPort *string `pulumi:"dstPort"`
	// Specifies whether the port forwarding rule is enabled or not. Defaults to `true`. This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
	//
	// Deprecated: This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
	Enabled *bool `pulumi:"enabled"`
	// The IPv4 address to forward traffic to.
	FwdIp *string `pulumi:"fwdIp"`
	// The port to forward traffic to.
	FwdPort *string `pulumi:"fwdPort"`
	// Specifies whether to log forwarded traffic or not. Defaults to `false`.
	Log *bool `pulumi:"log"`
	// The name of the port forwarding rule.
	Name *string `pulumi:"name"`
	// The port forwarding interface. Can be `wan`, `wan2`, or `both`.
	PortForwardInterface *string `pulumi:"portForwardInterface"`
	// The protocol for the port forwarding rule. Can be `tcp`, `udp`, or `tcpUdp`. Defaults to `tcpUdp`.
	Protocol *string `pulumi:"protocol"`
	// The name of the site to associate the port forwarding rule with.
	Site *string `pulumi:"site"`
	// The source IPv4 address (or CIDR) of the port forwarding rule. For all traffic, specify `any`. Defaults to `any`.
	SrcIp *string `pulumi:"srcIp"`
}

// The set of arguments for constructing a Forward resource.
type ForwardArgs struct {
	// The destination port for the forwarding.
	DstPort pulumi.StringPtrInput
	// Specifies whether the port forwarding rule is enabled or not. Defaults to `true`. This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
	//
	// Deprecated: This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
	Enabled pulumi.BoolPtrInput
	// The IPv4 address to forward traffic to.
	FwdIp pulumi.StringPtrInput
	// The port to forward traffic to.
	FwdPort pulumi.StringPtrInput
	// Specifies whether to log forwarded traffic or not. Defaults to `false`.
	Log pulumi.BoolPtrInput
	// The name of the port forwarding rule.
	Name pulumi.StringPtrInput
	// The port forwarding interface. Can be `wan`, `wan2`, or `both`.
	PortForwardInterface pulumi.StringPtrInput
	// The protocol for the port forwarding rule. Can be `tcp`, `udp`, or `tcpUdp`. Defaults to `tcpUdp`.
	Protocol pulumi.StringPtrInput
	// The name of the site to associate the port forwarding rule with.
	Site pulumi.StringPtrInput
	// The source IPv4 address (or CIDR) of the port forwarding rule. For all traffic, specify `any`. Defaults to `any`.
	SrcIp pulumi.StringPtrInput
}

func (ForwardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*forwardArgs)(nil)).Elem()
}

type ForwardInput interface {
	pulumi.Input

	ToForwardOutput() ForwardOutput
	ToForwardOutputWithContext(ctx context.Context) ForwardOutput
}

func (*Forward) ElementType() reflect.Type {
	return reflect.TypeOf((**Forward)(nil)).Elem()
}

func (i *Forward) ToForwardOutput() ForwardOutput {
	return i.ToForwardOutputWithContext(context.Background())
}

func (i *Forward) ToForwardOutputWithContext(ctx context.Context) ForwardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardOutput)
}

// ForwardArrayInput is an input type that accepts ForwardArray and ForwardArrayOutput values.
// You can construct a concrete instance of `ForwardArrayInput` via:
//
//	ForwardArray{ ForwardArgs{...} }
type ForwardArrayInput interface {
	pulumi.Input

	ToForwardArrayOutput() ForwardArrayOutput
	ToForwardArrayOutputWithContext(context.Context) ForwardArrayOutput
}

type ForwardArray []ForwardInput

func (ForwardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Forward)(nil)).Elem()
}

func (i ForwardArray) ToForwardArrayOutput() ForwardArrayOutput {
	return i.ToForwardArrayOutputWithContext(context.Background())
}

func (i ForwardArray) ToForwardArrayOutputWithContext(ctx context.Context) ForwardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardArrayOutput)
}

// ForwardMapInput is an input type that accepts ForwardMap and ForwardMapOutput values.
// You can construct a concrete instance of `ForwardMapInput` via:
//
//	ForwardMap{ "key": ForwardArgs{...} }
type ForwardMapInput interface {
	pulumi.Input

	ToForwardMapOutput() ForwardMapOutput
	ToForwardMapOutputWithContext(context.Context) ForwardMapOutput
}

type ForwardMap map[string]ForwardInput

func (ForwardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Forward)(nil)).Elem()
}

func (i ForwardMap) ToForwardMapOutput() ForwardMapOutput {
	return i.ToForwardMapOutputWithContext(context.Background())
}

func (i ForwardMap) ToForwardMapOutputWithContext(ctx context.Context) ForwardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardMapOutput)
}

type ForwardOutput struct{ *pulumi.OutputState }

func (ForwardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Forward)(nil)).Elem()
}

func (o ForwardOutput) ToForwardOutput() ForwardOutput {
	return o
}

func (o ForwardOutput) ToForwardOutputWithContext(ctx context.Context) ForwardOutput {
	return o
}

// The destination port for the forwarding.
func (o ForwardOutput) DstPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Forward) pulumi.StringPtrOutput { return v.DstPort }).(pulumi.StringPtrOutput)
}

// Specifies whether the port forwarding rule is enabled or not. Defaults to `true`. This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
//
// Deprecated: This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
func (o ForwardOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Forward) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The IPv4 address to forward traffic to.
func (o ForwardOutput) FwdIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Forward) pulumi.StringPtrOutput { return v.FwdIp }).(pulumi.StringPtrOutput)
}

// The port to forward traffic to.
func (o ForwardOutput) FwdPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Forward) pulumi.StringPtrOutput { return v.FwdPort }).(pulumi.StringPtrOutput)
}

// Specifies whether to log forwarded traffic or not. Defaults to `false`.
func (o ForwardOutput) Log() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Forward) pulumi.BoolPtrOutput { return v.Log }).(pulumi.BoolPtrOutput)
}

// The name of the port forwarding rule.
func (o ForwardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Forward) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The port forwarding interface. Can be `wan`, `wan2`, or `both`.
func (o ForwardOutput) PortForwardInterface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Forward) pulumi.StringPtrOutput { return v.PortForwardInterface }).(pulumi.StringPtrOutput)
}

// The protocol for the port forwarding rule. Can be `tcp`, `udp`, or `tcpUdp`. Defaults to `tcpUdp`.
func (o ForwardOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Forward) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The name of the site to associate the port forwarding rule with.
func (o ForwardOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *Forward) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

// The source IPv4 address (or CIDR) of the port forwarding rule. For all traffic, specify `any`. Defaults to `any`.
func (o ForwardOutput) SrcIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Forward) pulumi.StringPtrOutput { return v.SrcIp }).(pulumi.StringPtrOutput)
}

type ForwardArrayOutput struct{ *pulumi.OutputState }

func (ForwardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Forward)(nil)).Elem()
}

func (o ForwardArrayOutput) ToForwardArrayOutput() ForwardArrayOutput {
	return o
}

func (o ForwardArrayOutput) ToForwardArrayOutputWithContext(ctx context.Context) ForwardArrayOutput {
	return o
}

func (o ForwardArrayOutput) Index(i pulumi.IntInput) ForwardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Forward {
		return vs[0].([]*Forward)[vs[1].(int)]
	}).(ForwardOutput)
}

type ForwardMapOutput struct{ *pulumi.OutputState }

func (ForwardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Forward)(nil)).Elem()
}

func (o ForwardMapOutput) ToForwardMapOutput() ForwardMapOutput {
	return o
}

func (o ForwardMapOutput) ToForwardMapOutputWithContext(ctx context.Context) ForwardMapOutput {
	return o
}

func (o ForwardMapOutput) MapIndex(k pulumi.StringInput) ForwardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Forward {
		return vs[0].(map[string]*Forward)[vs[1].(string)]
	}).(ForwardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardInput)(nil)).Elem(), &Forward{})
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardArrayInput)(nil)).Elem(), ForwardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ForwardMapInput)(nil)).Elem(), ForwardMap{})
	pulumi.RegisterOutputType(ForwardOutput{})
	pulumi.RegisterOutputType(ForwardArrayOutput{})
	pulumi.RegisterOutputType(ForwardMapOutput{})
}
