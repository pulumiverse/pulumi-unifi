// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package unifi

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/internal"
)

// `Account` manages a RADIUS user account
//
// To authenticate devices based on MAC address, use the MAC address as the username and password under client creation.
// Convert lowercase letters to uppercase, and also remove colons or periods from the MAC address.
//
// ATTENTION: If the user profile does not include a VLAN, the client will fall back to the untagged VLAN.
//
// NOTE: MAC-based authentication accounts can only be used for wireless and wired clients. L2TP remote access does not apply.
type Account struct {
	pulumi.CustomResourceState

	// The name of the account.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the network for this account
	NetworkId pulumi.StringPtrOutput `pulumi:"networkId"`
	// The password of the account.
	Password pulumi.StringOutput `pulumi:"password"`
	// The name of the site to associate the account with.
	Site pulumi.StringOutput `pulumi:"site"`
	// See [RFC 2868](https://www.rfc-editor.org/rfc/rfc2868) section 3.2 Defaults to `6`.
	TunnelMediumType pulumi.IntPtrOutput `pulumi:"tunnelMediumType"`
	// See [RFC 2868](https://www.rfc-editor.org/rfc/rfc2868) section 3.1 Defaults to `13`.
	TunnelType pulumi.IntPtrOutput `pulumi:"tunnelType"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Account
	err := ctx.RegisterResource("unifi:index/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("unifi:index/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The name of the account.
	Name *string `pulumi:"name"`
	// ID of the network for this account
	NetworkId *string `pulumi:"networkId"`
	// The password of the account.
	Password *string `pulumi:"password"`
	// The name of the site to associate the account with.
	Site *string `pulumi:"site"`
	// See [RFC 2868](https://www.rfc-editor.org/rfc/rfc2868) section 3.2 Defaults to `6`.
	TunnelMediumType *int `pulumi:"tunnelMediumType"`
	// See [RFC 2868](https://www.rfc-editor.org/rfc/rfc2868) section 3.1 Defaults to `13`.
	TunnelType *int `pulumi:"tunnelType"`
}

type AccountState struct {
	// The name of the account.
	Name pulumi.StringPtrInput
	// ID of the network for this account
	NetworkId pulumi.StringPtrInput
	// The password of the account.
	Password pulumi.StringPtrInput
	// The name of the site to associate the account with.
	Site pulumi.StringPtrInput
	// See [RFC 2868](https://www.rfc-editor.org/rfc/rfc2868) section 3.2 Defaults to `6`.
	TunnelMediumType pulumi.IntPtrInput
	// See [RFC 2868](https://www.rfc-editor.org/rfc/rfc2868) section 3.1 Defaults to `13`.
	TunnelType pulumi.IntPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name of the account.
	Name *string `pulumi:"name"`
	// ID of the network for this account
	NetworkId *string `pulumi:"networkId"`
	// The password of the account.
	Password string `pulumi:"password"`
	// The name of the site to associate the account with.
	Site *string `pulumi:"site"`
	// See [RFC 2868](https://www.rfc-editor.org/rfc/rfc2868) section 3.2 Defaults to `6`.
	TunnelMediumType *int `pulumi:"tunnelMediumType"`
	// See [RFC 2868](https://www.rfc-editor.org/rfc/rfc2868) section 3.1 Defaults to `13`.
	TunnelType *int `pulumi:"tunnelType"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the account.
	Name pulumi.StringPtrInput
	// ID of the network for this account
	NetworkId pulumi.StringPtrInput
	// The password of the account.
	Password pulumi.StringInput
	// The name of the site to associate the account with.
	Site pulumi.StringPtrInput
	// See [RFC 2868](https://www.rfc-editor.org/rfc/rfc2868) section 3.2 Defaults to `6`.
	TunnelMediumType pulumi.IntPtrInput
	// See [RFC 2868](https://www.rfc-editor.org/rfc/rfc2868) section 3.1 Defaults to `13`.
	TunnelType pulumi.IntPtrInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

func (i *Account) ToOutput(ctx context.Context) pulumix.Output[*Account] {
	return pulumix.Output[*Account]{
		OutputState: i.ToAccountOutputWithContext(ctx).OutputState,
	}
}

// AccountArrayInput is an input type that accepts AccountArray and AccountArrayOutput values.
// You can construct a concrete instance of `AccountArrayInput` via:
//
//	AccountArray{ AccountArgs{...} }
type AccountArrayInput interface {
	pulumi.Input

	ToAccountArrayOutput() AccountArrayOutput
	ToAccountArrayOutputWithContext(context.Context) AccountArrayOutput
}

type AccountArray []AccountInput

func (AccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (i AccountArray) ToAccountArrayOutput() AccountArrayOutput {
	return i.ToAccountArrayOutputWithContext(context.Background())
}

func (i AccountArray) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountArrayOutput)
}

func (i AccountArray) ToOutput(ctx context.Context) pulumix.Output[[]*Account] {
	return pulumix.Output[[]*Account]{
		OutputState: i.ToAccountArrayOutputWithContext(ctx).OutputState,
	}
}

// AccountMapInput is an input type that accepts AccountMap and AccountMapOutput values.
// You can construct a concrete instance of `AccountMapInput` via:
//
//	AccountMap{ "key": AccountArgs{...} }
type AccountMapInput interface {
	pulumi.Input

	ToAccountMapOutput() AccountMapOutput
	ToAccountMapOutputWithContext(context.Context) AccountMapOutput
}

type AccountMap map[string]AccountInput

func (AccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (i AccountMap) ToAccountMapOutput() AccountMapOutput {
	return i.ToAccountMapOutputWithContext(context.Background())
}

func (i AccountMap) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountMapOutput)
}

func (i AccountMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Account] {
	return pulumix.Output[map[string]*Account]{
		OutputState: i.ToAccountMapOutputWithContext(ctx).OutputState,
	}
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func (o AccountOutput) ToOutput(ctx context.Context) pulumix.Output[*Account] {
	return pulumix.Output[*Account]{
		OutputState: o.OutputState,
	}
}

// The name of the account.
func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the network for this account
func (o AccountOutput) NetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.NetworkId }).(pulumi.StringPtrOutput)
}

// The password of the account.
func (o AccountOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The name of the site to associate the account with.
func (o AccountOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

// See [RFC 2868](https://www.rfc-editor.org/rfc/rfc2868) section 3.2 Defaults to `6`.
func (o AccountOutput) TunnelMediumType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.IntPtrOutput { return v.TunnelMediumType }).(pulumi.IntPtrOutput)
}

// See [RFC 2868](https://www.rfc-editor.org/rfc/rfc2868) section 3.1 Defaults to `13`.
func (o AccountOutput) TunnelType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.IntPtrOutput { return v.TunnelType }).(pulumi.IntPtrOutput)
}

type AccountArrayOutput struct{ *pulumi.OutputState }

func (AccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (o AccountArrayOutput) ToAccountArrayOutput() AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Account] {
	return pulumix.Output[[]*Account]{
		OutputState: o.OutputState,
	}
}

func (o AccountArrayOutput) Index(i pulumi.IntInput) AccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Account {
		return vs[0].([]*Account)[vs[1].(int)]
	}).(AccountOutput)
}

type AccountMapOutput struct{ *pulumi.OutputState }

func (AccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (o AccountMapOutput) ToAccountMapOutput() AccountMapOutput {
	return o
}

func (o AccountMapOutput) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return o
}

func (o AccountMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Account] {
	return pulumix.Output[map[string]*Account]{
		OutputState: o.OutputState,
	}
}

func (o AccountMapOutput) MapIndex(k pulumi.StringInput) AccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Account {
		return vs[0].(map[string]*Account)[vs[1].(string)]
	}).(AccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountInput)(nil)).Elem(), &Account{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountArrayInput)(nil)).Elem(), AccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountMapInput)(nil)).Elem(), AccountMap{})
	pulumi.RegisterOutputType(AccountOutput{})
	pulumi.RegisterOutputType(AccountArrayOutput{})
	pulumi.RegisterOutputType(AccountMapOutput{})
}
