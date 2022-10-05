// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package unifi

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `Wlan` manages a WiFi network / SSID.
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
//	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi"
//	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/iam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			vlanId := 10
//			if param := cfg.GetFloat64("vlanId"); param != 0 {
//				vlanId = param
//			}
//			defaultApGroup, err := unifi.GetApGroup(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultGroup, err := iam.LookupGroup(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			vlan, err := unifi.NewNetwork(ctx, "vlan", &unifi.NetworkArgs{
//				Purpose:     pulumi.String("corporate"),
//				Subnet:      pulumi.String("10.0.0.1/24"),
//				VlanId:      pulumi.Float64(vlanId),
//				DhcpStart:   pulumi.String("10.0.0.6"),
//				DhcpStop:    pulumi.String("10.0.0.254"),
//				DhcpEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = unifi.NewWlan(ctx, "wifi", &unifi.WlanArgs{
//				Passphrase:     pulumi.String("12345678"),
//				Security:       pulumi.String("wpapsk"),
//				Wpa3Support:    pulumi.Bool(true),
//				Wpa3Transition: pulumi.Bool(true),
//				PmfMode:        pulumi.String("optional"),
//				NetworkId:      vlan.ID(),
//				ApGroupIds: pulumi.StringArray{
//					pulumi.String(defaultApGroup.Id),
//				},
//				UserGroupId: pulumi.String(defaultGroup.Id),
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
// import from provider configured site
//
// ```sh
//
//	$ pulumi import unifi:index/wlan:Wlan mywlan 5dc28e5e9106d105bdc87217
//
// ```
//
//	import from another site
//
// ```sh
//
//	$ pulumi import unifi:index/wlan:Wlan mywlan bfa2l6i7:5dc28e5e9106d105bdc87217
//
// ```
type Wlan struct {
	pulumi.CustomResourceState

	// IDs of the AP groups to use for this network.
	ApGroupIds pulumi.StringArrayOutput `pulumi:"apGroupIds"`
	// Indicates whether or not to hide the SSID from broadcast.
	HideSsid pulumi.BoolPtrOutput `pulumi:"hideSsid"`
	// Indicates that this is a guest WLAN and should use guest behaviors.
	IsGuest pulumi.BoolPtrOutput `pulumi:"isGuest"`
	// Isolates stations on layer 2 (ethernet) level Defaults to `false`.
	L2Isolation pulumi.BoolPtrOutput `pulumi:"l2Isolation"`
	// Indicates whether or not the MAC filter is turned of for the network.
	MacFilterEnabled pulumi.BoolPtrOutput `pulumi:"macFilterEnabled"`
	// List of MAC addresses to filter (only valid if `macFilterEnabled` is `true`).
	MacFilterLists pulumi.StringArrayOutput `pulumi:"macFilterLists"`
	// MAC address filter policy (only valid if `macFilterEnabled` is `true`). Defaults to `deny`.
	MacFilterPolicy pulumi.StringPtrOutput `pulumi:"macFilterPolicy"`
	// Set minimum data rate control for 2G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `1000`, `2000`, `5500`, `6000`, `9000`, `11000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
	MinimumDataRate2gKbps pulumi.IntPtrOutput `pulumi:"minimumDataRate2gKbps"`
	// Set minimum data rate control for 5G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `6000`, `9000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
	MinimumDataRate5gKbps pulumi.IntPtrOutput `pulumi:"minimumDataRate5gKbps"`
	// Indicates whether or not Multicast Enhance is turned of for the network.
	MulticastEnhance pulumi.BoolPtrOutput `pulumi:"multicastEnhance"`
	// The SSID of the network.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the network for this SSID
	NetworkId pulumi.StringPtrOutput `pulumi:"networkId"`
	// Connect high performance clients to 5 GHz only Defaults to `true`.
	No2ghzOui pulumi.BoolPtrOutput `pulumi:"no2ghzOui"`
	// The passphrase for the network, this is only required if `security` is not set to `open`.
	Passphrase pulumi.StringPtrOutput `pulumi:"passphrase"`
	// Enable Protected Management Frames. This cannot be disabled if using WPA 3. Valid values are `required`, `optional` and `disabled`. Defaults to `disabled`.
	PmfMode pulumi.StringPtrOutput `pulumi:"pmfMode"`
	// ID of the RADIUS profile to use when security `wpaeap`. You can query this via the `getRadiusProfile` data source.
	RadiusProfileId pulumi.StringPtrOutput `pulumi:"radiusProfileId"`
	// Start and stop schedules for the WLAN
	Schedules WlanScheduleArrayOutput `pulumi:"schedules"`
	// The type of WiFi security for this network. Valid values are: `wpapsk`, `wpaeap`, and `open`.
	Security pulumi.StringOutput `pulumi:"security"`
	// The name of the site to associate the wlan with.
	Site pulumi.StringOutput `pulumi:"site"`
	// Enable Unscheduled Automatic Power Save Delivery Defaults to `false`.
	Uapsd pulumi.BoolPtrOutput `pulumi:"uapsd"`
	// ID of the user group to use for this network.
	UserGroupId pulumi.StringOutput `pulumi:"userGroupId"`
	// Radio band your WiFi network will use.
	WlanBand pulumi.StringPtrOutput `pulumi:"wlanBand"`
	// Enable WPA 3 support (security must be `wpapsk` and PMF must be turned on).
	Wpa3Support pulumi.BoolPtrOutput `pulumi:"wpa3Support"`
	// Enable WPA 3 and WPA 2 support (security must be `wpapsk` and `wpa3Support` must be true).
	Wpa3Transition pulumi.BoolPtrOutput `pulumi:"wpa3Transition"`
}

// NewWlan registers a new resource with the given unique name, arguments, and options.
func NewWlan(ctx *pulumi.Context,
	name string, args *WlanArgs, opts ...pulumi.ResourceOption) (*Wlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Security == nil {
		return nil, errors.New("invalid value for required argument 'Security'")
	}
	if args.UserGroupId == nil {
		return nil, errors.New("invalid value for required argument 'UserGroupId'")
	}
	if args.Passphrase != nil {
		args.Passphrase = pulumi.ToSecret(args.Passphrase).(pulumi.StringPtrOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"passphrase",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource Wlan
	err := ctx.RegisterResource("unifi:index/wlan:Wlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWlan gets an existing Wlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WlanState, opts ...pulumi.ResourceOption) (*Wlan, error) {
	var resource Wlan
	err := ctx.ReadResource("unifi:index/wlan:Wlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Wlan resources.
type wlanState struct {
	// IDs of the AP groups to use for this network.
	ApGroupIds []string `pulumi:"apGroupIds"`
	// Indicates whether or not to hide the SSID from broadcast.
	HideSsid *bool `pulumi:"hideSsid"`
	// Indicates that this is a guest WLAN and should use guest behaviors.
	IsGuest *bool `pulumi:"isGuest"`
	// Isolates stations on layer 2 (ethernet) level Defaults to `false`.
	L2Isolation *bool `pulumi:"l2Isolation"`
	// Indicates whether or not the MAC filter is turned of for the network.
	MacFilterEnabled *bool `pulumi:"macFilterEnabled"`
	// List of MAC addresses to filter (only valid if `macFilterEnabled` is `true`).
	MacFilterLists []string `pulumi:"macFilterLists"`
	// MAC address filter policy (only valid if `macFilterEnabled` is `true`). Defaults to `deny`.
	MacFilterPolicy *string `pulumi:"macFilterPolicy"`
	// Set minimum data rate control for 2G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `1000`, `2000`, `5500`, `6000`, `9000`, `11000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
	MinimumDataRate2gKbps *int `pulumi:"minimumDataRate2gKbps"`
	// Set minimum data rate control for 5G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `6000`, `9000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
	MinimumDataRate5gKbps *int `pulumi:"minimumDataRate5gKbps"`
	// Indicates whether or not Multicast Enhance is turned of for the network.
	MulticastEnhance *bool `pulumi:"multicastEnhance"`
	// The SSID of the network.
	Name *string `pulumi:"name"`
	// ID of the network for this SSID
	NetworkId *string `pulumi:"networkId"`
	// Connect high performance clients to 5 GHz only Defaults to `true`.
	No2ghzOui *bool `pulumi:"no2ghzOui"`
	// The passphrase for the network, this is only required if `security` is not set to `open`.
	Passphrase *string `pulumi:"passphrase"`
	// Enable Protected Management Frames. This cannot be disabled if using WPA 3. Valid values are `required`, `optional` and `disabled`. Defaults to `disabled`.
	PmfMode *string `pulumi:"pmfMode"`
	// ID of the RADIUS profile to use when security `wpaeap`. You can query this via the `getRadiusProfile` data source.
	RadiusProfileId *string `pulumi:"radiusProfileId"`
	// Start and stop schedules for the WLAN
	Schedules []WlanSchedule `pulumi:"schedules"`
	// The type of WiFi security for this network. Valid values are: `wpapsk`, `wpaeap`, and `open`.
	Security *string `pulumi:"security"`
	// The name of the site to associate the wlan with.
	Site *string `pulumi:"site"`
	// Enable Unscheduled Automatic Power Save Delivery Defaults to `false`.
	Uapsd *bool `pulumi:"uapsd"`
	// ID of the user group to use for this network.
	UserGroupId *string `pulumi:"userGroupId"`
	// Radio band your WiFi network will use.
	WlanBand *string `pulumi:"wlanBand"`
	// Enable WPA 3 support (security must be `wpapsk` and PMF must be turned on).
	Wpa3Support *bool `pulumi:"wpa3Support"`
	// Enable WPA 3 and WPA 2 support (security must be `wpapsk` and `wpa3Support` must be true).
	Wpa3Transition *bool `pulumi:"wpa3Transition"`
}

type WlanState struct {
	// IDs of the AP groups to use for this network.
	ApGroupIds pulumi.StringArrayInput
	// Indicates whether or not to hide the SSID from broadcast.
	HideSsid pulumi.BoolPtrInput
	// Indicates that this is a guest WLAN and should use guest behaviors.
	IsGuest pulumi.BoolPtrInput
	// Isolates stations on layer 2 (ethernet) level Defaults to `false`.
	L2Isolation pulumi.BoolPtrInput
	// Indicates whether or not the MAC filter is turned of for the network.
	MacFilterEnabled pulumi.BoolPtrInput
	// List of MAC addresses to filter (only valid if `macFilterEnabled` is `true`).
	MacFilterLists pulumi.StringArrayInput
	// MAC address filter policy (only valid if `macFilterEnabled` is `true`). Defaults to `deny`.
	MacFilterPolicy pulumi.StringPtrInput
	// Set minimum data rate control for 2G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `1000`, `2000`, `5500`, `6000`, `9000`, `11000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
	MinimumDataRate2gKbps pulumi.IntPtrInput
	// Set minimum data rate control for 5G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `6000`, `9000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
	MinimumDataRate5gKbps pulumi.IntPtrInput
	// Indicates whether or not Multicast Enhance is turned of for the network.
	MulticastEnhance pulumi.BoolPtrInput
	// The SSID of the network.
	Name pulumi.StringPtrInput
	// ID of the network for this SSID
	NetworkId pulumi.StringPtrInput
	// Connect high performance clients to 5 GHz only Defaults to `true`.
	No2ghzOui pulumi.BoolPtrInput
	// The passphrase for the network, this is only required if `security` is not set to `open`.
	Passphrase pulumi.StringPtrInput
	// Enable Protected Management Frames. This cannot be disabled if using WPA 3. Valid values are `required`, `optional` and `disabled`. Defaults to `disabled`.
	PmfMode pulumi.StringPtrInput
	// ID of the RADIUS profile to use when security `wpaeap`. You can query this via the `getRadiusProfile` data source.
	RadiusProfileId pulumi.StringPtrInput
	// Start and stop schedules for the WLAN
	Schedules WlanScheduleArrayInput
	// The type of WiFi security for this network. Valid values are: `wpapsk`, `wpaeap`, and `open`.
	Security pulumi.StringPtrInput
	// The name of the site to associate the wlan with.
	Site pulumi.StringPtrInput
	// Enable Unscheduled Automatic Power Save Delivery Defaults to `false`.
	Uapsd pulumi.BoolPtrInput
	// ID of the user group to use for this network.
	UserGroupId pulumi.StringPtrInput
	// Radio band your WiFi network will use.
	WlanBand pulumi.StringPtrInput
	// Enable WPA 3 support (security must be `wpapsk` and PMF must be turned on).
	Wpa3Support pulumi.BoolPtrInput
	// Enable WPA 3 and WPA 2 support (security must be `wpapsk` and `wpa3Support` must be true).
	Wpa3Transition pulumi.BoolPtrInput
}

func (WlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*wlanState)(nil)).Elem()
}

type wlanArgs struct {
	// IDs of the AP groups to use for this network.
	ApGroupIds []string `pulumi:"apGroupIds"`
	// Indicates whether or not to hide the SSID from broadcast.
	HideSsid *bool `pulumi:"hideSsid"`
	// Indicates that this is a guest WLAN and should use guest behaviors.
	IsGuest *bool `pulumi:"isGuest"`
	// Isolates stations on layer 2 (ethernet) level Defaults to `false`.
	L2Isolation *bool `pulumi:"l2Isolation"`
	// Indicates whether or not the MAC filter is turned of for the network.
	MacFilterEnabled *bool `pulumi:"macFilterEnabled"`
	// List of MAC addresses to filter (only valid if `macFilterEnabled` is `true`).
	MacFilterLists []string `pulumi:"macFilterLists"`
	// MAC address filter policy (only valid if `macFilterEnabled` is `true`). Defaults to `deny`.
	MacFilterPolicy *string `pulumi:"macFilterPolicy"`
	// Set minimum data rate control for 2G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `1000`, `2000`, `5500`, `6000`, `9000`, `11000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
	MinimumDataRate2gKbps *int `pulumi:"minimumDataRate2gKbps"`
	// Set minimum data rate control for 5G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `6000`, `9000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
	MinimumDataRate5gKbps *int `pulumi:"minimumDataRate5gKbps"`
	// Indicates whether or not Multicast Enhance is turned of for the network.
	MulticastEnhance *bool `pulumi:"multicastEnhance"`
	// The SSID of the network.
	Name *string `pulumi:"name"`
	// ID of the network for this SSID
	NetworkId *string `pulumi:"networkId"`
	// Connect high performance clients to 5 GHz only Defaults to `true`.
	No2ghzOui *bool `pulumi:"no2ghzOui"`
	// The passphrase for the network, this is only required if `security` is not set to `open`.
	Passphrase *string `pulumi:"passphrase"`
	// Enable Protected Management Frames. This cannot be disabled if using WPA 3. Valid values are `required`, `optional` and `disabled`. Defaults to `disabled`.
	PmfMode *string `pulumi:"pmfMode"`
	// ID of the RADIUS profile to use when security `wpaeap`. You can query this via the `getRadiusProfile` data source.
	RadiusProfileId *string `pulumi:"radiusProfileId"`
	// Start and stop schedules for the WLAN
	Schedules []WlanSchedule `pulumi:"schedules"`
	// The type of WiFi security for this network. Valid values are: `wpapsk`, `wpaeap`, and `open`.
	Security string `pulumi:"security"`
	// The name of the site to associate the wlan with.
	Site *string `pulumi:"site"`
	// Enable Unscheduled Automatic Power Save Delivery Defaults to `false`.
	Uapsd *bool `pulumi:"uapsd"`
	// ID of the user group to use for this network.
	UserGroupId string `pulumi:"userGroupId"`
	// Radio band your WiFi network will use.
	WlanBand *string `pulumi:"wlanBand"`
	// Enable WPA 3 support (security must be `wpapsk` and PMF must be turned on).
	Wpa3Support *bool `pulumi:"wpa3Support"`
	// Enable WPA 3 and WPA 2 support (security must be `wpapsk` and `wpa3Support` must be true).
	Wpa3Transition *bool `pulumi:"wpa3Transition"`
}

// The set of arguments for constructing a Wlan resource.
type WlanArgs struct {
	// IDs of the AP groups to use for this network.
	ApGroupIds pulumi.StringArrayInput
	// Indicates whether or not to hide the SSID from broadcast.
	HideSsid pulumi.BoolPtrInput
	// Indicates that this is a guest WLAN and should use guest behaviors.
	IsGuest pulumi.BoolPtrInput
	// Isolates stations on layer 2 (ethernet) level Defaults to `false`.
	L2Isolation pulumi.BoolPtrInput
	// Indicates whether or not the MAC filter is turned of for the network.
	MacFilterEnabled pulumi.BoolPtrInput
	// List of MAC addresses to filter (only valid if `macFilterEnabled` is `true`).
	MacFilterLists pulumi.StringArrayInput
	// MAC address filter policy (only valid if `macFilterEnabled` is `true`). Defaults to `deny`.
	MacFilterPolicy pulumi.StringPtrInput
	// Set minimum data rate control for 2G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `1000`, `2000`, `5500`, `6000`, `9000`, `11000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
	MinimumDataRate2gKbps pulumi.IntPtrInput
	// Set minimum data rate control for 5G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `6000`, `9000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
	MinimumDataRate5gKbps pulumi.IntPtrInput
	// Indicates whether or not Multicast Enhance is turned of for the network.
	MulticastEnhance pulumi.BoolPtrInput
	// The SSID of the network.
	Name pulumi.StringPtrInput
	// ID of the network for this SSID
	NetworkId pulumi.StringPtrInput
	// Connect high performance clients to 5 GHz only Defaults to `true`.
	No2ghzOui pulumi.BoolPtrInput
	// The passphrase for the network, this is only required if `security` is not set to `open`.
	Passphrase pulumi.StringPtrInput
	// Enable Protected Management Frames. This cannot be disabled if using WPA 3. Valid values are `required`, `optional` and `disabled`. Defaults to `disabled`.
	PmfMode pulumi.StringPtrInput
	// ID of the RADIUS profile to use when security `wpaeap`. You can query this via the `getRadiusProfile` data source.
	RadiusProfileId pulumi.StringPtrInput
	// Start and stop schedules for the WLAN
	Schedules WlanScheduleArrayInput
	// The type of WiFi security for this network. Valid values are: `wpapsk`, `wpaeap`, and `open`.
	Security pulumi.StringInput
	// The name of the site to associate the wlan with.
	Site pulumi.StringPtrInput
	// Enable Unscheduled Automatic Power Save Delivery Defaults to `false`.
	Uapsd pulumi.BoolPtrInput
	// ID of the user group to use for this network.
	UserGroupId pulumi.StringInput
	// Radio band your WiFi network will use.
	WlanBand pulumi.StringPtrInput
	// Enable WPA 3 support (security must be `wpapsk` and PMF must be turned on).
	Wpa3Support pulumi.BoolPtrInput
	// Enable WPA 3 and WPA 2 support (security must be `wpapsk` and `wpa3Support` must be true).
	Wpa3Transition pulumi.BoolPtrInput
}

func (WlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wlanArgs)(nil)).Elem()
}

type WlanInput interface {
	pulumi.Input

	ToWlanOutput() WlanOutput
	ToWlanOutputWithContext(ctx context.Context) WlanOutput
}

func (*Wlan) ElementType() reflect.Type {
	return reflect.TypeOf((**Wlan)(nil)).Elem()
}

func (i *Wlan) ToWlanOutput() WlanOutput {
	return i.ToWlanOutputWithContext(context.Background())
}

func (i *Wlan) ToWlanOutputWithContext(ctx context.Context) WlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WlanOutput)
}

// WlanArrayInput is an input type that accepts WlanArray and WlanArrayOutput values.
// You can construct a concrete instance of `WlanArrayInput` via:
//
//	WlanArray{ WlanArgs{...} }
type WlanArrayInput interface {
	pulumi.Input

	ToWlanArrayOutput() WlanArrayOutput
	ToWlanArrayOutputWithContext(context.Context) WlanArrayOutput
}

type WlanArray []WlanInput

func (WlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wlan)(nil)).Elem()
}

func (i WlanArray) ToWlanArrayOutput() WlanArrayOutput {
	return i.ToWlanArrayOutputWithContext(context.Background())
}

func (i WlanArray) ToWlanArrayOutputWithContext(ctx context.Context) WlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WlanArrayOutput)
}

// WlanMapInput is an input type that accepts WlanMap and WlanMapOutput values.
// You can construct a concrete instance of `WlanMapInput` via:
//
//	WlanMap{ "key": WlanArgs{...} }
type WlanMapInput interface {
	pulumi.Input

	ToWlanMapOutput() WlanMapOutput
	ToWlanMapOutputWithContext(context.Context) WlanMapOutput
}

type WlanMap map[string]WlanInput

func (WlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wlan)(nil)).Elem()
}

func (i WlanMap) ToWlanMapOutput() WlanMapOutput {
	return i.ToWlanMapOutputWithContext(context.Background())
}

func (i WlanMap) ToWlanMapOutputWithContext(ctx context.Context) WlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WlanMapOutput)
}

type WlanOutput struct{ *pulumi.OutputState }

func (WlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Wlan)(nil)).Elem()
}

func (o WlanOutput) ToWlanOutput() WlanOutput {
	return o
}

func (o WlanOutput) ToWlanOutputWithContext(ctx context.Context) WlanOutput {
	return o
}

// IDs of the AP groups to use for this network.
func (o WlanOutput) ApGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Wlan) pulumi.StringArrayOutput { return v.ApGroupIds }).(pulumi.StringArrayOutput)
}

// Indicates whether or not to hide the SSID from broadcast.
func (o WlanOutput) HideSsid() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.BoolPtrOutput { return v.HideSsid }).(pulumi.BoolPtrOutput)
}

// Indicates that this is a guest WLAN and should use guest behaviors.
func (o WlanOutput) IsGuest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.BoolPtrOutput { return v.IsGuest }).(pulumi.BoolPtrOutput)
}

// Isolates stations on layer 2 (ethernet) level Defaults to `false`.
func (o WlanOutput) L2Isolation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.BoolPtrOutput { return v.L2Isolation }).(pulumi.BoolPtrOutput)
}

// Indicates whether or not the MAC filter is turned of for the network.
func (o WlanOutput) MacFilterEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.BoolPtrOutput { return v.MacFilterEnabled }).(pulumi.BoolPtrOutput)
}

// List of MAC addresses to filter (only valid if `macFilterEnabled` is `true`).
func (o WlanOutput) MacFilterLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Wlan) pulumi.StringArrayOutput { return v.MacFilterLists }).(pulumi.StringArrayOutput)
}

// MAC address filter policy (only valid if `macFilterEnabled` is `true`). Defaults to `deny`.
func (o WlanOutput) MacFilterPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.StringPtrOutput { return v.MacFilterPolicy }).(pulumi.StringPtrOutput)
}

// Set minimum data rate control for 2G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `1000`, `2000`, `5500`, `6000`, `9000`, `11000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
func (o WlanOutput) MinimumDataRate2gKbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.IntPtrOutput { return v.MinimumDataRate2gKbps }).(pulumi.IntPtrOutput)
}

// Set minimum data rate control for 5G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `6000`, `9000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
func (o WlanOutput) MinimumDataRate5gKbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.IntPtrOutput { return v.MinimumDataRate5gKbps }).(pulumi.IntPtrOutput)
}

// Indicates whether or not Multicast Enhance is turned of for the network.
func (o WlanOutput) MulticastEnhance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.BoolPtrOutput { return v.MulticastEnhance }).(pulumi.BoolPtrOutput)
}

// The SSID of the network.
func (o WlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Wlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the network for this SSID
func (o WlanOutput) NetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.StringPtrOutput { return v.NetworkId }).(pulumi.StringPtrOutput)
}

// Connect high performance clients to 5 GHz only Defaults to `true`.
func (o WlanOutput) No2ghzOui() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.BoolPtrOutput { return v.No2ghzOui }).(pulumi.BoolPtrOutput)
}

// The passphrase for the network, this is only required if `security` is not set to `open`.
func (o WlanOutput) Passphrase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.StringPtrOutput { return v.Passphrase }).(pulumi.StringPtrOutput)
}

// Enable Protected Management Frames. This cannot be disabled if using WPA 3. Valid values are `required`, `optional` and `disabled`. Defaults to `disabled`.
func (o WlanOutput) PmfMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.StringPtrOutput { return v.PmfMode }).(pulumi.StringPtrOutput)
}

// ID of the RADIUS profile to use when security `wpaeap`. You can query this via the `getRadiusProfile` data source.
func (o WlanOutput) RadiusProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.StringPtrOutput { return v.RadiusProfileId }).(pulumi.StringPtrOutput)
}

// Start and stop schedules for the WLAN
func (o WlanOutput) Schedules() WlanScheduleArrayOutput {
	return o.ApplyT(func(v *Wlan) WlanScheduleArrayOutput { return v.Schedules }).(WlanScheduleArrayOutput)
}

// The type of WiFi security for this network. Valid values are: `wpapsk`, `wpaeap`, and `open`.
func (o WlanOutput) Security() pulumi.StringOutput {
	return o.ApplyT(func(v *Wlan) pulumi.StringOutput { return v.Security }).(pulumi.StringOutput)
}

// The name of the site to associate the wlan with.
func (o WlanOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *Wlan) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

// Enable Unscheduled Automatic Power Save Delivery Defaults to `false`.
func (o WlanOutput) Uapsd() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.BoolPtrOutput { return v.Uapsd }).(pulumi.BoolPtrOutput)
}

// ID of the user group to use for this network.
func (o WlanOutput) UserGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Wlan) pulumi.StringOutput { return v.UserGroupId }).(pulumi.StringOutput)
}

// Radio band your WiFi network will use.
func (o WlanOutput) WlanBand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.StringPtrOutput { return v.WlanBand }).(pulumi.StringPtrOutput)
}

// Enable WPA 3 support (security must be `wpapsk` and PMF must be turned on).
func (o WlanOutput) Wpa3Support() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.BoolPtrOutput { return v.Wpa3Support }).(pulumi.BoolPtrOutput)
}

// Enable WPA 3 and WPA 2 support (security must be `wpapsk` and `wpa3Support` must be true).
func (o WlanOutput) Wpa3Transition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Wlan) pulumi.BoolPtrOutput { return v.Wpa3Transition }).(pulumi.BoolPtrOutput)
}

type WlanArrayOutput struct{ *pulumi.OutputState }

func (WlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wlan)(nil)).Elem()
}

func (o WlanArrayOutput) ToWlanArrayOutput() WlanArrayOutput {
	return o
}

func (o WlanArrayOutput) ToWlanArrayOutputWithContext(ctx context.Context) WlanArrayOutput {
	return o
}

func (o WlanArrayOutput) Index(i pulumi.IntInput) WlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Wlan {
		return vs[0].([]*Wlan)[vs[1].(int)]
	}).(WlanOutput)
}

type WlanMapOutput struct{ *pulumi.OutputState }

func (WlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wlan)(nil)).Elem()
}

func (o WlanMapOutput) ToWlanMapOutput() WlanMapOutput {
	return o
}

func (o WlanMapOutput) ToWlanMapOutputWithContext(ctx context.Context) WlanMapOutput {
	return o
}

func (o WlanMapOutput) MapIndex(k pulumi.StringInput) WlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Wlan {
		return vs[0].(map[string]*Wlan)[vs[1].(string)]
	}).(WlanOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WlanInput)(nil)).Elem(), &Wlan{})
	pulumi.RegisterInputType(reflect.TypeOf((*WlanArrayInput)(nil)).Elem(), WlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WlanMapInput)(nil)).Elem(), WlanMap{})
	pulumi.RegisterOutputType(WlanOutput{})
	pulumi.RegisterOutputType(WlanArrayOutput{})
	pulumi.RegisterOutputType(WlanMapOutput{})
}
