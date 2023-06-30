// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package unifi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DevicePortOverride struct {
	// Number of ports in the aggregate.
	AggregateNumPorts *int `pulumi:"aggregateNumPorts"`
	// Human-readable name of the port.
	Name *string `pulumi:"name"`
	// Switch port number.
	Number int `pulumi:"number"`
	// Operating mode of the port, valid values are `switch`, `mirror`, and `aggregate`. Defaults to `switch`.
	OpMode *string `pulumi:"opMode"`
	// ID of the Port Profile used on this port.
	PortProfileId *string `pulumi:"portProfileId"`
}

// DevicePortOverrideInput is an input type that accepts DevicePortOverrideArgs and DevicePortOverrideOutput values.
// You can construct a concrete instance of `DevicePortOverrideInput` via:
//
//	DevicePortOverrideArgs{...}
type DevicePortOverrideInput interface {
	pulumi.Input

	ToDevicePortOverrideOutput() DevicePortOverrideOutput
	ToDevicePortOverrideOutputWithContext(context.Context) DevicePortOverrideOutput
}

type DevicePortOverrideArgs struct {
	// Number of ports in the aggregate.
	AggregateNumPorts pulumi.IntPtrInput `pulumi:"aggregateNumPorts"`
	// Human-readable name of the port.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Switch port number.
	Number pulumi.IntInput `pulumi:"number"`
	// Operating mode of the port, valid values are `switch`, `mirror`, and `aggregate`. Defaults to `switch`.
	OpMode pulumi.StringPtrInput `pulumi:"opMode"`
	// ID of the Port Profile used on this port.
	PortProfileId pulumi.StringPtrInput `pulumi:"portProfileId"`
}

func (DevicePortOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DevicePortOverride)(nil)).Elem()
}

func (i DevicePortOverrideArgs) ToDevicePortOverrideOutput() DevicePortOverrideOutput {
	return i.ToDevicePortOverrideOutputWithContext(context.Background())
}

func (i DevicePortOverrideArgs) ToDevicePortOverrideOutputWithContext(ctx context.Context) DevicePortOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicePortOverrideOutput)
}

// DevicePortOverrideArrayInput is an input type that accepts DevicePortOverrideArray and DevicePortOverrideArrayOutput values.
// You can construct a concrete instance of `DevicePortOverrideArrayInput` via:
//
//	DevicePortOverrideArray{ DevicePortOverrideArgs{...} }
type DevicePortOverrideArrayInput interface {
	pulumi.Input

	ToDevicePortOverrideArrayOutput() DevicePortOverrideArrayOutput
	ToDevicePortOverrideArrayOutputWithContext(context.Context) DevicePortOverrideArrayOutput
}

type DevicePortOverrideArray []DevicePortOverrideInput

func (DevicePortOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DevicePortOverride)(nil)).Elem()
}

func (i DevicePortOverrideArray) ToDevicePortOverrideArrayOutput() DevicePortOverrideArrayOutput {
	return i.ToDevicePortOverrideArrayOutputWithContext(context.Background())
}

func (i DevicePortOverrideArray) ToDevicePortOverrideArrayOutputWithContext(ctx context.Context) DevicePortOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicePortOverrideArrayOutput)
}

type DevicePortOverrideOutput struct{ *pulumi.OutputState }

func (DevicePortOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DevicePortOverride)(nil)).Elem()
}

func (o DevicePortOverrideOutput) ToDevicePortOverrideOutput() DevicePortOverrideOutput {
	return o
}

func (o DevicePortOverrideOutput) ToDevicePortOverrideOutputWithContext(ctx context.Context) DevicePortOverrideOutput {
	return o
}

// Number of ports in the aggregate.
func (o DevicePortOverrideOutput) AggregateNumPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DevicePortOverride) *int { return v.AggregateNumPorts }).(pulumi.IntPtrOutput)
}

// Human-readable name of the port.
func (o DevicePortOverrideOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DevicePortOverride) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Switch port number.
func (o DevicePortOverrideOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v DevicePortOverride) int { return v.Number }).(pulumi.IntOutput)
}

// Operating mode of the port, valid values are `switch`, `mirror`, and `aggregate`. Defaults to `switch`.
func (o DevicePortOverrideOutput) OpMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DevicePortOverride) *string { return v.OpMode }).(pulumi.StringPtrOutput)
}

// ID of the Port Profile used on this port.
func (o DevicePortOverrideOutput) PortProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DevicePortOverride) *string { return v.PortProfileId }).(pulumi.StringPtrOutput)
}

type DevicePortOverrideArrayOutput struct{ *pulumi.OutputState }

func (DevicePortOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DevicePortOverride)(nil)).Elem()
}

func (o DevicePortOverrideArrayOutput) ToDevicePortOverrideArrayOutput() DevicePortOverrideArrayOutput {
	return o
}

func (o DevicePortOverrideArrayOutput) ToDevicePortOverrideArrayOutputWithContext(ctx context.Context) DevicePortOverrideArrayOutput {
	return o
}

func (o DevicePortOverrideArrayOutput) Index(i pulumi.IntInput) DevicePortOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DevicePortOverride {
		return vs[0].([]DevicePortOverride)[vs[1].(int)]
	}).(DevicePortOverrideOutput)
}

type RadiusProfileAcctServer struct {
	// IP address of accounting service server.
	Ip string `pulumi:"ip"`
	// Port of accounting service. Defaults to `1813`.
	Port *int `pulumi:"port"`
	// RADIUS secret.
	Xsecret string `pulumi:"xsecret"`
}

// RadiusProfileAcctServerInput is an input type that accepts RadiusProfileAcctServerArgs and RadiusProfileAcctServerOutput values.
// You can construct a concrete instance of `RadiusProfileAcctServerInput` via:
//
//	RadiusProfileAcctServerArgs{...}
type RadiusProfileAcctServerInput interface {
	pulumi.Input

	ToRadiusProfileAcctServerOutput() RadiusProfileAcctServerOutput
	ToRadiusProfileAcctServerOutputWithContext(context.Context) RadiusProfileAcctServerOutput
}

type RadiusProfileAcctServerArgs struct {
	// IP address of accounting service server.
	Ip pulumi.StringInput `pulumi:"ip"`
	// Port of accounting service. Defaults to `1813`.
	Port pulumi.IntPtrInput `pulumi:"port"`
	// RADIUS secret.
	Xsecret pulumi.StringInput `pulumi:"xsecret"`
}

func (RadiusProfileAcctServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RadiusProfileAcctServer)(nil)).Elem()
}

func (i RadiusProfileAcctServerArgs) ToRadiusProfileAcctServerOutput() RadiusProfileAcctServerOutput {
	return i.ToRadiusProfileAcctServerOutputWithContext(context.Background())
}

func (i RadiusProfileAcctServerArgs) ToRadiusProfileAcctServerOutputWithContext(ctx context.Context) RadiusProfileAcctServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RadiusProfileAcctServerOutput)
}

// RadiusProfileAcctServerArrayInput is an input type that accepts RadiusProfileAcctServerArray and RadiusProfileAcctServerArrayOutput values.
// You can construct a concrete instance of `RadiusProfileAcctServerArrayInput` via:
//
//	RadiusProfileAcctServerArray{ RadiusProfileAcctServerArgs{...} }
type RadiusProfileAcctServerArrayInput interface {
	pulumi.Input

	ToRadiusProfileAcctServerArrayOutput() RadiusProfileAcctServerArrayOutput
	ToRadiusProfileAcctServerArrayOutputWithContext(context.Context) RadiusProfileAcctServerArrayOutput
}

type RadiusProfileAcctServerArray []RadiusProfileAcctServerInput

func (RadiusProfileAcctServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RadiusProfileAcctServer)(nil)).Elem()
}

func (i RadiusProfileAcctServerArray) ToRadiusProfileAcctServerArrayOutput() RadiusProfileAcctServerArrayOutput {
	return i.ToRadiusProfileAcctServerArrayOutputWithContext(context.Background())
}

func (i RadiusProfileAcctServerArray) ToRadiusProfileAcctServerArrayOutputWithContext(ctx context.Context) RadiusProfileAcctServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RadiusProfileAcctServerArrayOutput)
}

type RadiusProfileAcctServerOutput struct{ *pulumi.OutputState }

func (RadiusProfileAcctServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RadiusProfileAcctServer)(nil)).Elem()
}

func (o RadiusProfileAcctServerOutput) ToRadiusProfileAcctServerOutput() RadiusProfileAcctServerOutput {
	return o
}

func (o RadiusProfileAcctServerOutput) ToRadiusProfileAcctServerOutputWithContext(ctx context.Context) RadiusProfileAcctServerOutput {
	return o
}

// IP address of accounting service server.
func (o RadiusProfileAcctServerOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v RadiusProfileAcctServer) string { return v.Ip }).(pulumi.StringOutput)
}

// Port of accounting service. Defaults to `1813`.
func (o RadiusProfileAcctServerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RadiusProfileAcctServer) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// RADIUS secret.
func (o RadiusProfileAcctServerOutput) Xsecret() pulumi.StringOutput {
	return o.ApplyT(func(v RadiusProfileAcctServer) string { return v.Xsecret }).(pulumi.StringOutput)
}

type RadiusProfileAcctServerArrayOutput struct{ *pulumi.OutputState }

func (RadiusProfileAcctServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RadiusProfileAcctServer)(nil)).Elem()
}

func (o RadiusProfileAcctServerArrayOutput) ToRadiusProfileAcctServerArrayOutput() RadiusProfileAcctServerArrayOutput {
	return o
}

func (o RadiusProfileAcctServerArrayOutput) ToRadiusProfileAcctServerArrayOutputWithContext(ctx context.Context) RadiusProfileAcctServerArrayOutput {
	return o
}

func (o RadiusProfileAcctServerArrayOutput) Index(i pulumi.IntInput) RadiusProfileAcctServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RadiusProfileAcctServer {
		return vs[0].([]RadiusProfileAcctServer)[vs[1].(int)]
	}).(RadiusProfileAcctServerOutput)
}

type RadiusProfileAuthServer struct {
	// IP address of authentication service server.
	Ip string `pulumi:"ip"`
	// Port of authentication service. Defaults to `1812`.
	Port *int `pulumi:"port"`
	// RADIUS secret.
	Xsecret string `pulumi:"xsecret"`
}

// RadiusProfileAuthServerInput is an input type that accepts RadiusProfileAuthServerArgs and RadiusProfileAuthServerOutput values.
// You can construct a concrete instance of `RadiusProfileAuthServerInput` via:
//
//	RadiusProfileAuthServerArgs{...}
type RadiusProfileAuthServerInput interface {
	pulumi.Input

	ToRadiusProfileAuthServerOutput() RadiusProfileAuthServerOutput
	ToRadiusProfileAuthServerOutputWithContext(context.Context) RadiusProfileAuthServerOutput
}

type RadiusProfileAuthServerArgs struct {
	// IP address of authentication service server.
	Ip pulumi.StringInput `pulumi:"ip"`
	// Port of authentication service. Defaults to `1812`.
	Port pulumi.IntPtrInput `pulumi:"port"`
	// RADIUS secret.
	Xsecret pulumi.StringInput `pulumi:"xsecret"`
}

func (RadiusProfileAuthServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RadiusProfileAuthServer)(nil)).Elem()
}

func (i RadiusProfileAuthServerArgs) ToRadiusProfileAuthServerOutput() RadiusProfileAuthServerOutput {
	return i.ToRadiusProfileAuthServerOutputWithContext(context.Background())
}

func (i RadiusProfileAuthServerArgs) ToRadiusProfileAuthServerOutputWithContext(ctx context.Context) RadiusProfileAuthServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RadiusProfileAuthServerOutput)
}

// RadiusProfileAuthServerArrayInput is an input type that accepts RadiusProfileAuthServerArray and RadiusProfileAuthServerArrayOutput values.
// You can construct a concrete instance of `RadiusProfileAuthServerArrayInput` via:
//
//	RadiusProfileAuthServerArray{ RadiusProfileAuthServerArgs{...} }
type RadiusProfileAuthServerArrayInput interface {
	pulumi.Input

	ToRadiusProfileAuthServerArrayOutput() RadiusProfileAuthServerArrayOutput
	ToRadiusProfileAuthServerArrayOutputWithContext(context.Context) RadiusProfileAuthServerArrayOutput
}

type RadiusProfileAuthServerArray []RadiusProfileAuthServerInput

func (RadiusProfileAuthServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RadiusProfileAuthServer)(nil)).Elem()
}

func (i RadiusProfileAuthServerArray) ToRadiusProfileAuthServerArrayOutput() RadiusProfileAuthServerArrayOutput {
	return i.ToRadiusProfileAuthServerArrayOutputWithContext(context.Background())
}

func (i RadiusProfileAuthServerArray) ToRadiusProfileAuthServerArrayOutputWithContext(ctx context.Context) RadiusProfileAuthServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RadiusProfileAuthServerArrayOutput)
}

type RadiusProfileAuthServerOutput struct{ *pulumi.OutputState }

func (RadiusProfileAuthServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RadiusProfileAuthServer)(nil)).Elem()
}

func (o RadiusProfileAuthServerOutput) ToRadiusProfileAuthServerOutput() RadiusProfileAuthServerOutput {
	return o
}

func (o RadiusProfileAuthServerOutput) ToRadiusProfileAuthServerOutputWithContext(ctx context.Context) RadiusProfileAuthServerOutput {
	return o
}

// IP address of authentication service server.
func (o RadiusProfileAuthServerOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v RadiusProfileAuthServer) string { return v.Ip }).(pulumi.StringOutput)
}

// Port of authentication service. Defaults to `1812`.
func (o RadiusProfileAuthServerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RadiusProfileAuthServer) *int { return v.Port }).(pulumi.IntPtrOutput)
}

// RADIUS secret.
func (o RadiusProfileAuthServerOutput) Xsecret() pulumi.StringOutput {
	return o.ApplyT(func(v RadiusProfileAuthServer) string { return v.Xsecret }).(pulumi.StringOutput)
}

type RadiusProfileAuthServerArrayOutput struct{ *pulumi.OutputState }

func (RadiusProfileAuthServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RadiusProfileAuthServer)(nil)).Elem()
}

func (o RadiusProfileAuthServerArrayOutput) ToRadiusProfileAuthServerArrayOutput() RadiusProfileAuthServerArrayOutput {
	return o
}

func (o RadiusProfileAuthServerArrayOutput) ToRadiusProfileAuthServerArrayOutputWithContext(ctx context.Context) RadiusProfileAuthServerArrayOutput {
	return o
}

func (o RadiusProfileAuthServerArrayOutput) Index(i pulumi.IntInput) RadiusProfileAuthServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RadiusProfileAuthServer {
		return vs[0].([]RadiusProfileAuthServer)[vs[1].(int)]
	}).(RadiusProfileAuthServerOutput)
}

type WlanSchedule struct {
	// Day of week for the block. Valid values are `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.
	DayOfWeek string `pulumi:"dayOfWeek"`
	// Length of the block in minutes.
	Duration int `pulumi:"duration"`
	// Name of the block.
	Name *string `pulumi:"name"`
	// Start hour for the block (0-23).
	StartHour int `pulumi:"startHour"`
	// Start minute for the block (0-59). Defaults to `0`.
	StartMinute *int `pulumi:"startMinute"`
}

// WlanScheduleInput is an input type that accepts WlanScheduleArgs and WlanScheduleOutput values.
// You can construct a concrete instance of `WlanScheduleInput` via:
//
//	WlanScheduleArgs{...}
type WlanScheduleInput interface {
	pulumi.Input

	ToWlanScheduleOutput() WlanScheduleOutput
	ToWlanScheduleOutputWithContext(context.Context) WlanScheduleOutput
}

type WlanScheduleArgs struct {
	// Day of week for the block. Valid values are `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.
	DayOfWeek pulumi.StringInput `pulumi:"dayOfWeek"`
	// Length of the block in minutes.
	Duration pulumi.IntInput `pulumi:"duration"`
	// Name of the block.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Start hour for the block (0-23).
	StartHour pulumi.IntInput `pulumi:"startHour"`
	// Start minute for the block (0-59). Defaults to `0`.
	StartMinute pulumi.IntPtrInput `pulumi:"startMinute"`
}

func (WlanScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WlanSchedule)(nil)).Elem()
}

func (i WlanScheduleArgs) ToWlanScheduleOutput() WlanScheduleOutput {
	return i.ToWlanScheduleOutputWithContext(context.Background())
}

func (i WlanScheduleArgs) ToWlanScheduleOutputWithContext(ctx context.Context) WlanScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WlanScheduleOutput)
}

// WlanScheduleArrayInput is an input type that accepts WlanScheduleArray and WlanScheduleArrayOutput values.
// You can construct a concrete instance of `WlanScheduleArrayInput` via:
//
//	WlanScheduleArray{ WlanScheduleArgs{...} }
type WlanScheduleArrayInput interface {
	pulumi.Input

	ToWlanScheduleArrayOutput() WlanScheduleArrayOutput
	ToWlanScheduleArrayOutputWithContext(context.Context) WlanScheduleArrayOutput
}

type WlanScheduleArray []WlanScheduleInput

func (WlanScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WlanSchedule)(nil)).Elem()
}

func (i WlanScheduleArray) ToWlanScheduleArrayOutput() WlanScheduleArrayOutput {
	return i.ToWlanScheduleArrayOutputWithContext(context.Background())
}

func (i WlanScheduleArray) ToWlanScheduleArrayOutputWithContext(ctx context.Context) WlanScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WlanScheduleArrayOutput)
}

type WlanScheduleOutput struct{ *pulumi.OutputState }

func (WlanScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WlanSchedule)(nil)).Elem()
}

func (o WlanScheduleOutput) ToWlanScheduleOutput() WlanScheduleOutput {
	return o
}

func (o WlanScheduleOutput) ToWlanScheduleOutputWithContext(ctx context.Context) WlanScheduleOutput {
	return o
}

// Day of week for the block. Valid values are `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.
func (o WlanScheduleOutput) DayOfWeek() pulumi.StringOutput {
	return o.ApplyT(func(v WlanSchedule) string { return v.DayOfWeek }).(pulumi.StringOutput)
}

// Length of the block in minutes.
func (o WlanScheduleOutput) Duration() pulumi.IntOutput {
	return o.ApplyT(func(v WlanSchedule) int { return v.Duration }).(pulumi.IntOutput)
}

// Name of the block.
func (o WlanScheduleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WlanSchedule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Start hour for the block (0-23).
func (o WlanScheduleOutput) StartHour() pulumi.IntOutput {
	return o.ApplyT(func(v WlanSchedule) int { return v.StartHour }).(pulumi.IntOutput)
}

// Start minute for the block (0-59). Defaults to `0`.
func (o WlanScheduleOutput) StartMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WlanSchedule) *int { return v.StartMinute }).(pulumi.IntPtrOutput)
}

type WlanScheduleArrayOutput struct{ *pulumi.OutputState }

func (WlanScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WlanSchedule)(nil)).Elem()
}

func (o WlanScheduleArrayOutput) ToWlanScheduleArrayOutput() WlanScheduleArrayOutput {
	return o
}

func (o WlanScheduleArrayOutput) ToWlanScheduleArrayOutputWithContext(ctx context.Context) WlanScheduleArrayOutput {
	return o
}

func (o WlanScheduleArrayOutput) Index(i pulumi.IntInput) WlanScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WlanSchedule {
		return vs[0].([]WlanSchedule)[vs[1].(int)]
	}).(WlanScheduleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DevicePortOverrideInput)(nil)).Elem(), DevicePortOverrideArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DevicePortOverrideArrayInput)(nil)).Elem(), DevicePortOverrideArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RadiusProfileAcctServerInput)(nil)).Elem(), RadiusProfileAcctServerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RadiusProfileAcctServerArrayInput)(nil)).Elem(), RadiusProfileAcctServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RadiusProfileAuthServerInput)(nil)).Elem(), RadiusProfileAuthServerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RadiusProfileAuthServerArrayInput)(nil)).Elem(), RadiusProfileAuthServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WlanScheduleInput)(nil)).Elem(), WlanScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WlanScheduleArrayInput)(nil)).Elem(), WlanScheduleArray{})
	pulumi.RegisterOutputType(DevicePortOverrideOutput{})
	pulumi.RegisterOutputType(DevicePortOverrideArrayOutput{})
	pulumi.RegisterOutputType(RadiusProfileAcctServerOutput{})
	pulumi.RegisterOutputType(RadiusProfileAcctServerArrayOutput{})
	pulumi.RegisterOutputType(RadiusProfileAuthServerOutput{})
	pulumi.RegisterOutputType(RadiusProfileAuthServerArrayOutput{})
	pulumi.RegisterOutputType(WlanScheduleOutput{})
	pulumi.RegisterOutputType(WlanScheduleArrayOutput{})
}
