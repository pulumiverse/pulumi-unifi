// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/internal"
)

// `iam.User` retrieves properties of a user (or "client" in the UI) of the network by MAC address.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-unifi/sdk/go/unifi/iam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.LookupUser(ctx, &iam.LookupUserArgs{
//				Mac: "01:23:45:67:89:ab",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("unifi:iam/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// The MAC address of the user.
	Mac string `pulumi:"mac"`
	// The name of the site the user is associated with.
	Site *string `pulumi:"site"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	// Specifies whether this user should be blocked from the network.
	Blocked bool `pulumi:"blocked"`
	// Override the device fingerprint.
	DevIdOverride int `pulumi:"devIdOverride"`
	// fixed IPv4 address set for this user.
	FixedIp string `pulumi:"fixedIp"`
	// The hostname of the user.
	Hostname string `pulumi:"hostname"`
	// The ID of the user.
	Id string `pulumi:"id"`
	// The IP address of the user.
	Ip string `pulumi:"ip"`
	// The local DNS record for this user.
	LocalDnsRecord string `pulumi:"localDnsRecord"`
	// The MAC address of the user.
	Mac string `pulumi:"mac"`
	// The name of the user.
	Name string `pulumi:"name"`
	// The network ID for this user.
	NetworkId string `pulumi:"networkId"`
	// A note with additional information for the user.
	Note string `pulumi:"note"`
	// The name of the site the user is associated with.
	Site string `pulumi:"site"`
	// The user group ID for the user.
	UserGroupId string `pulumi:"userGroupId"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResultOutput, error) {
			args := v.(LookupUserArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupUserResult
			secret, err := ctx.InvokePackageRaw("unifi:iam/getUser:getUser", args, &rv, "", opts...)
			if err != nil {
				return LookupUserResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupUserResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupUserResultOutput), nil
			}
			return output, nil
		}).(LookupUserResultOutput)
}

// A collection of arguments for invoking getUser.
type LookupUserOutputArgs struct {
	// The MAC address of the user.
	Mac pulumi.StringInput `pulumi:"mac"`
	// The name of the site the user is associated with.
	Site pulumi.StringPtrInput `pulumi:"site"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

// Specifies whether this user should be blocked from the network.
func (o LookupUserResultOutput) Blocked() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupUserResult) bool { return v.Blocked }).(pulumi.BoolOutput)
}

// Override the device fingerprint.
func (o LookupUserResultOutput) DevIdOverride() pulumi.IntOutput {
	return o.ApplyT(func(v LookupUserResult) int { return v.DevIdOverride }).(pulumi.IntOutput)
}

// fixed IPv4 address set for this user.
func (o LookupUserResultOutput) FixedIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.FixedIp }).(pulumi.StringOutput)
}

// The hostname of the user.
func (o LookupUserResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The ID of the user.
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// The IP address of the user.
func (o LookupUserResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Ip }).(pulumi.StringOutput)
}

// The local DNS record for this user.
func (o LookupUserResultOutput) LocalDnsRecord() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.LocalDnsRecord }).(pulumi.StringOutput)
}

// The MAC address of the user.
func (o LookupUserResultOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Mac }).(pulumi.StringOutput)
}

// The name of the user.
func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// The network ID for this user.
func (o LookupUserResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

// A note with additional information for the user.
func (o LookupUserResultOutput) Note() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Note }).(pulumi.StringOutput)
}

// The name of the site the user is associated with.
func (o LookupUserResultOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Site }).(pulumi.StringOutput)
}

// The user group ID for the user.
func (o LookupUserResultOutput) UserGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.UserGroupId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
