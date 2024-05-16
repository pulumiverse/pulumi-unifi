// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package unifi

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/internal"
)

// `DynamicDNS` manages dynamic DNS settings for different providers.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := unifi.NewDynamicDNS(ctx, "test", &unifi.DynamicDNSArgs{
//				Service:  pulumi.String("dyndns"),
//				HostName: pulumi.String("my-network.example.com"),
//				Server:   pulumi.String("domains.google.com"),
//				Login:    pulumi.Any(_var.Dns_login),
//				Password: pulumi.Any(_var.Dns_password),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DynamicDNS struct {
	pulumi.CustomResourceState

	// The host name to update in the dynamic DNS service.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The interface for the dynamic DNS. Can be `wan` or `wan2`. Defaults to `wan`.
	Interface pulumi.StringPtrOutput `pulumi:"interface"`
	// The server for the dynamic DNS service.
	Login pulumi.StringPtrOutput `pulumi:"login"`
	// The server for the dynamic DNS service.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The server for the dynamic DNS service.
	Server pulumi.StringPtrOutput `pulumi:"server"`
	// The Dynamic DNS service provider, various values are supported (for example `dyndns`, etc.).
	Service pulumi.StringOutput `pulumi:"service"`
	// The name of the site to associate the dynamic DNS with.
	Site pulumi.StringOutput `pulumi:"site"`
}

// NewDynamicDNS registers a new resource with the given unique name, arguments, and options.
func NewDynamicDNS(ctx *pulumi.Context,
	name string, args *DynamicDNSArgs, opts ...pulumi.ResourceOption) (*DynamicDNS, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DynamicDNS
	err := ctx.RegisterResource("unifi:index/dynamicDNS:DynamicDNS", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDynamicDNS gets an existing DynamicDNS resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDynamicDNS(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DynamicDNSState, opts ...pulumi.ResourceOption) (*DynamicDNS, error) {
	var resource DynamicDNS
	err := ctx.ReadResource("unifi:index/dynamicDNS:DynamicDNS", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DynamicDNS resources.
type dynamicDNSState struct {
	// The host name to update in the dynamic DNS service.
	HostName *string `pulumi:"hostName"`
	// The interface for the dynamic DNS. Can be `wan` or `wan2`. Defaults to `wan`.
	Interface *string `pulumi:"interface"`
	// The server for the dynamic DNS service.
	Login *string `pulumi:"login"`
	// The server for the dynamic DNS service.
	Password *string `pulumi:"password"`
	// The server for the dynamic DNS service.
	Server *string `pulumi:"server"`
	// The Dynamic DNS service provider, various values are supported (for example `dyndns`, etc.).
	Service *string `pulumi:"service"`
	// The name of the site to associate the dynamic DNS with.
	Site *string `pulumi:"site"`
}

type DynamicDNSState struct {
	// The host name to update in the dynamic DNS service.
	HostName pulumi.StringPtrInput
	// The interface for the dynamic DNS. Can be `wan` or `wan2`. Defaults to `wan`.
	Interface pulumi.StringPtrInput
	// The server for the dynamic DNS service.
	Login pulumi.StringPtrInput
	// The server for the dynamic DNS service.
	Password pulumi.StringPtrInput
	// The server for the dynamic DNS service.
	Server pulumi.StringPtrInput
	// The Dynamic DNS service provider, various values are supported (for example `dyndns`, etc.).
	Service pulumi.StringPtrInput
	// The name of the site to associate the dynamic DNS with.
	Site pulumi.StringPtrInput
}

func (DynamicDNSState) ElementType() reflect.Type {
	return reflect.TypeOf((*dynamicDNSState)(nil)).Elem()
}

type dynamicDNSArgs struct {
	// The host name to update in the dynamic DNS service.
	HostName string `pulumi:"hostName"`
	// The interface for the dynamic DNS. Can be `wan` or `wan2`. Defaults to `wan`.
	Interface *string `pulumi:"interface"`
	// The server for the dynamic DNS service.
	Login *string `pulumi:"login"`
	// The server for the dynamic DNS service.
	Password *string `pulumi:"password"`
	// The server for the dynamic DNS service.
	Server *string `pulumi:"server"`
	// The Dynamic DNS service provider, various values are supported (for example `dyndns`, etc.).
	Service string `pulumi:"service"`
	// The name of the site to associate the dynamic DNS with.
	Site *string `pulumi:"site"`
}

// The set of arguments for constructing a DynamicDNS resource.
type DynamicDNSArgs struct {
	// The host name to update in the dynamic DNS service.
	HostName pulumi.StringInput
	// The interface for the dynamic DNS. Can be `wan` or `wan2`. Defaults to `wan`.
	Interface pulumi.StringPtrInput
	// The server for the dynamic DNS service.
	Login pulumi.StringPtrInput
	// The server for the dynamic DNS service.
	Password pulumi.StringPtrInput
	// The server for the dynamic DNS service.
	Server pulumi.StringPtrInput
	// The Dynamic DNS service provider, various values are supported (for example `dyndns`, etc.).
	Service pulumi.StringInput
	// The name of the site to associate the dynamic DNS with.
	Site pulumi.StringPtrInput
}

func (DynamicDNSArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dynamicDNSArgs)(nil)).Elem()
}

type DynamicDNSInput interface {
	pulumi.Input

	ToDynamicDNSOutput() DynamicDNSOutput
	ToDynamicDNSOutputWithContext(ctx context.Context) DynamicDNSOutput
}

func (*DynamicDNS) ElementType() reflect.Type {
	return reflect.TypeOf((**DynamicDNS)(nil)).Elem()
}

func (i *DynamicDNS) ToDynamicDNSOutput() DynamicDNSOutput {
	return i.ToDynamicDNSOutputWithContext(context.Background())
}

func (i *DynamicDNS) ToDynamicDNSOutputWithContext(ctx context.Context) DynamicDNSOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicDNSOutput)
}

// DynamicDNSArrayInput is an input type that accepts DynamicDNSArray and DynamicDNSArrayOutput values.
// You can construct a concrete instance of `DynamicDNSArrayInput` via:
//
//	DynamicDNSArray{ DynamicDNSArgs{...} }
type DynamicDNSArrayInput interface {
	pulumi.Input

	ToDynamicDNSArrayOutput() DynamicDNSArrayOutput
	ToDynamicDNSArrayOutputWithContext(context.Context) DynamicDNSArrayOutput
}

type DynamicDNSArray []DynamicDNSInput

func (DynamicDNSArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DynamicDNS)(nil)).Elem()
}

func (i DynamicDNSArray) ToDynamicDNSArrayOutput() DynamicDNSArrayOutput {
	return i.ToDynamicDNSArrayOutputWithContext(context.Background())
}

func (i DynamicDNSArray) ToDynamicDNSArrayOutputWithContext(ctx context.Context) DynamicDNSArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicDNSArrayOutput)
}

// DynamicDNSMapInput is an input type that accepts DynamicDNSMap and DynamicDNSMapOutput values.
// You can construct a concrete instance of `DynamicDNSMapInput` via:
//
//	DynamicDNSMap{ "key": DynamicDNSArgs{...} }
type DynamicDNSMapInput interface {
	pulumi.Input

	ToDynamicDNSMapOutput() DynamicDNSMapOutput
	ToDynamicDNSMapOutputWithContext(context.Context) DynamicDNSMapOutput
}

type DynamicDNSMap map[string]DynamicDNSInput

func (DynamicDNSMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DynamicDNS)(nil)).Elem()
}

func (i DynamicDNSMap) ToDynamicDNSMapOutput() DynamicDNSMapOutput {
	return i.ToDynamicDNSMapOutputWithContext(context.Background())
}

func (i DynamicDNSMap) ToDynamicDNSMapOutputWithContext(ctx context.Context) DynamicDNSMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynamicDNSMapOutput)
}

type DynamicDNSOutput struct{ *pulumi.OutputState }

func (DynamicDNSOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DynamicDNS)(nil)).Elem()
}

func (o DynamicDNSOutput) ToDynamicDNSOutput() DynamicDNSOutput {
	return o
}

func (o DynamicDNSOutput) ToDynamicDNSOutputWithContext(ctx context.Context) DynamicDNSOutput {
	return o
}

// The host name to update in the dynamic DNS service.
func (o DynamicDNSOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *DynamicDNS) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// The interface for the dynamic DNS. Can be `wan` or `wan2`. Defaults to `wan`.
func (o DynamicDNSOutput) Interface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DynamicDNS) pulumi.StringPtrOutput { return v.Interface }).(pulumi.StringPtrOutput)
}

// The server for the dynamic DNS service.
func (o DynamicDNSOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DynamicDNS) pulumi.StringPtrOutput { return v.Login }).(pulumi.StringPtrOutput)
}

// The server for the dynamic DNS service.
func (o DynamicDNSOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DynamicDNS) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The server for the dynamic DNS service.
func (o DynamicDNSOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DynamicDNS) pulumi.StringPtrOutput { return v.Server }).(pulumi.StringPtrOutput)
}

// The Dynamic DNS service provider, various values are supported (for example `dyndns`, etc.).
func (o DynamicDNSOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *DynamicDNS) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// The name of the site to associate the dynamic DNS with.
func (o DynamicDNSOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *DynamicDNS) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

type DynamicDNSArrayOutput struct{ *pulumi.OutputState }

func (DynamicDNSArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DynamicDNS)(nil)).Elem()
}

func (o DynamicDNSArrayOutput) ToDynamicDNSArrayOutput() DynamicDNSArrayOutput {
	return o
}

func (o DynamicDNSArrayOutput) ToDynamicDNSArrayOutputWithContext(ctx context.Context) DynamicDNSArrayOutput {
	return o
}

func (o DynamicDNSArrayOutput) Index(i pulumi.IntInput) DynamicDNSOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DynamicDNS {
		return vs[0].([]*DynamicDNS)[vs[1].(int)]
	}).(DynamicDNSOutput)
}

type DynamicDNSMapOutput struct{ *pulumi.OutputState }

func (DynamicDNSMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DynamicDNS)(nil)).Elem()
}

func (o DynamicDNSMapOutput) ToDynamicDNSMapOutput() DynamicDNSMapOutput {
	return o
}

func (o DynamicDNSMapOutput) ToDynamicDNSMapOutputWithContext(ctx context.Context) DynamicDNSMapOutput {
	return o
}

func (o DynamicDNSMapOutput) MapIndex(k pulumi.StringInput) DynamicDNSOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DynamicDNS {
		return vs[0].(map[string]*DynamicDNS)[vs[1].(string)]
	}).(DynamicDNSOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DynamicDNSInput)(nil)).Elem(), &DynamicDNS{})
	pulumi.RegisterInputType(reflect.TypeOf((*DynamicDNSArrayInput)(nil)).Elem(), DynamicDNSArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DynamicDNSMapInput)(nil)).Elem(), DynamicDNSMap{})
	pulumi.RegisterOutputType(DynamicDNSOutput{})
	pulumi.RegisterOutputType(DynamicDNSArrayOutput{})
	pulumi.RegisterOutputType(DynamicDNSMapOutput{})
}
