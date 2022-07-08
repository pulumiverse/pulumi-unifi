// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package unifi

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `User` manages a user (or "client" in the UI) of the network, these are identified by unique MAC addresses.
//
// Users are created in the controller when observed on the network, so the resource defaults to allowing itself to just take over management of a MAC address, but this can be turned off.
type User struct {
	pulumi.CustomResourceState

	// Specifies whether this resource should just take over control of an existing user. Defaults to `true`.
	AllowExisting pulumi.BoolPtrOutput `pulumi:"allowExisting"`
	// Specifies whether this user should be blocked from the network.
	Blocked pulumi.BoolPtrOutput `pulumi:"blocked"`
	// Override the device fingerprint.
	DevIdOverride pulumi.IntPtrOutput `pulumi:"devIdOverride"`
	// A fixed IPv4 address for this user.
	FixedIp pulumi.StringPtrOutput `pulumi:"fixedIp"`
	// The hostname of the user.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The IP address of the user.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The MAC address of the user.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// The name of the user.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network ID for this user.
	NetworkId pulumi.StringPtrOutput `pulumi:"networkId"`
	// A note with additional information for the user.
	Note pulumi.StringPtrOutput `pulumi:"note"`
	// The name of the site to associate the user with.
	Site pulumi.StringOutput `pulumi:"site"`
	// Specifies whether this resource should tell the controller to "forget" the user on destroy. Defaults to `false`.
	SkipForgetOnDestroy pulumi.BoolPtrOutput `pulumi:"skipForgetOnDestroy"`
	// The user group ID for the user.
	UserGroupId pulumi.StringPtrOutput `pulumi:"userGroupId"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Mac == nil {
		return nil, errors.New("invalid value for required argument 'Mac'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource User
	err := ctx.RegisterResource("unifi:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("unifi:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Specifies whether this resource should just take over control of an existing user. Defaults to `true`.
	AllowExisting *bool `pulumi:"allowExisting"`
	// Specifies whether this user should be blocked from the network.
	Blocked *bool `pulumi:"blocked"`
	// Override the device fingerprint.
	DevIdOverride *int `pulumi:"devIdOverride"`
	// A fixed IPv4 address for this user.
	FixedIp *string `pulumi:"fixedIp"`
	// The hostname of the user.
	Hostname *string `pulumi:"hostname"`
	// The IP address of the user.
	Ip *string `pulumi:"ip"`
	// The MAC address of the user.
	Mac *string `pulumi:"mac"`
	// The name of the user.
	Name *string `pulumi:"name"`
	// The network ID for this user.
	NetworkId *string `pulumi:"networkId"`
	// A note with additional information for the user.
	Note *string `pulumi:"note"`
	// The name of the site to associate the user with.
	Site *string `pulumi:"site"`
	// Specifies whether this resource should tell the controller to "forget" the user on destroy. Defaults to `false`.
	SkipForgetOnDestroy *bool `pulumi:"skipForgetOnDestroy"`
	// The user group ID for the user.
	UserGroupId *string `pulumi:"userGroupId"`
}

type UserState struct {
	// Specifies whether this resource should just take over control of an existing user. Defaults to `true`.
	AllowExisting pulumi.BoolPtrInput
	// Specifies whether this user should be blocked from the network.
	Blocked pulumi.BoolPtrInput
	// Override the device fingerprint.
	DevIdOverride pulumi.IntPtrInput
	// A fixed IPv4 address for this user.
	FixedIp pulumi.StringPtrInput
	// The hostname of the user.
	Hostname pulumi.StringPtrInput
	// The IP address of the user.
	Ip pulumi.StringPtrInput
	// The MAC address of the user.
	Mac pulumi.StringPtrInput
	// The name of the user.
	Name pulumi.StringPtrInput
	// The network ID for this user.
	NetworkId pulumi.StringPtrInput
	// A note with additional information for the user.
	Note pulumi.StringPtrInput
	// The name of the site to associate the user with.
	Site pulumi.StringPtrInput
	// Specifies whether this resource should tell the controller to "forget" the user on destroy. Defaults to `false`.
	SkipForgetOnDestroy pulumi.BoolPtrInput
	// The user group ID for the user.
	UserGroupId pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// Specifies whether this resource should just take over control of an existing user. Defaults to `true`.
	AllowExisting *bool `pulumi:"allowExisting"`
	// Specifies whether this user should be blocked from the network.
	Blocked *bool `pulumi:"blocked"`
	// Override the device fingerprint.
	DevIdOverride *int `pulumi:"devIdOverride"`
	// A fixed IPv4 address for this user.
	FixedIp *string `pulumi:"fixedIp"`
	// The MAC address of the user.
	Mac string `pulumi:"mac"`
	// The name of the user.
	Name *string `pulumi:"name"`
	// The network ID for this user.
	NetworkId *string `pulumi:"networkId"`
	// A note with additional information for the user.
	Note *string `pulumi:"note"`
	// The name of the site to associate the user with.
	Site *string `pulumi:"site"`
	// Specifies whether this resource should tell the controller to "forget" the user on destroy. Defaults to `false`.
	SkipForgetOnDestroy *bool `pulumi:"skipForgetOnDestroy"`
	// The user group ID for the user.
	UserGroupId *string `pulumi:"userGroupId"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Specifies whether this resource should just take over control of an existing user. Defaults to `true`.
	AllowExisting pulumi.BoolPtrInput
	// Specifies whether this user should be blocked from the network.
	Blocked pulumi.BoolPtrInput
	// Override the device fingerprint.
	DevIdOverride pulumi.IntPtrInput
	// A fixed IPv4 address for this user.
	FixedIp pulumi.StringPtrInput
	// The MAC address of the user.
	Mac pulumi.StringInput
	// The name of the user.
	Name pulumi.StringPtrInput
	// The network ID for this user.
	NetworkId pulumi.StringPtrInput
	// A note with additional information for the user.
	Note pulumi.StringPtrInput
	// The name of the site to associate the user with.
	Site pulumi.StringPtrInput
	// Specifies whether this resource should tell the controller to "forget" the user on destroy. Defaults to `false`.
	SkipForgetOnDestroy pulumi.BoolPtrInput
	// The user group ID for the user.
	UserGroupId pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//          UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//          UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// Specifies whether this resource should just take over control of an existing user. Defaults to `true`.
func (o UserOutput) AllowExisting() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.AllowExisting }).(pulumi.BoolPtrOutput)
}

// Specifies whether this user should be blocked from the network.
func (o UserOutput) Blocked() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.Blocked }).(pulumi.BoolPtrOutput)
}

// Override the device fingerprint.
func (o UserOutput) DevIdOverride() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *User) pulumi.IntPtrOutput { return v.DevIdOverride }).(pulumi.IntPtrOutput)
}

// A fixed IPv4 address for this user.
func (o UserOutput) FixedIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.FixedIp }).(pulumi.StringPtrOutput)
}

// The hostname of the user.
func (o UserOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// The IP address of the user.
func (o UserOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The MAC address of the user.
func (o UserOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Mac }).(pulumi.StringOutput)
}

// The name of the user.
func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network ID for this user.
func (o UserOutput) NetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.NetworkId }).(pulumi.StringPtrOutput)
}

// A note with additional information for the user.
func (o UserOutput) Note() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Note }).(pulumi.StringPtrOutput)
}

// The name of the site to associate the user with.
func (o UserOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

// Specifies whether this resource should tell the controller to "forget" the user on destroy. Defaults to `false`.
func (o UserOutput) SkipForgetOnDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.SkipForgetOnDestroy }).(pulumi.BoolPtrOutput)
}

// The user group ID for the user.
func (o UserOutput) UserGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.UserGroupId }).(pulumi.StringPtrOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
