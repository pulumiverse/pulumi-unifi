// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Unifi
{
    /// <summary>
    /// `unifi.User` manages a user (or "client" in the UI) of the network, these are identified by unique MAC addresses.
    /// 
    /// Users are created in the controller when observed on the network, so the resource defaults to allowing itself to just take over management of a MAC address, but this can be turned off.
    /// </summary>
    [UnifiResourceType("unifi:index/user:User")]
    public partial class User : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether this resource should just take over control of an existing user. Defaults to `true`.
        /// </summary>
        [Output("allowExisting")]
        public Output<bool?> AllowExisting { get; private set; } = null!;

        /// <summary>
        /// Specifies whether this user should be blocked from the network.
        /// </summary>
        [Output("blocked")]
        public Output<bool?> Blocked { get; private set; } = null!;

        /// <summary>
        /// Override the device fingerprint.
        /// </summary>
        [Output("devIdOverride")]
        public Output<int?> DevIdOverride { get; private set; } = null!;

        /// <summary>
        /// A fixed IPv4 address for this user.
        /// </summary>
        [Output("fixedIp")]
        public Output<string?> FixedIp { get; private set; } = null!;

        /// <summary>
        /// The hostname of the user.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// The IP address of the user.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// The MAC address of the user.
        /// </summary>
        [Output("mac")]
        public Output<string> Mac { get; private set; } = null!;

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network ID for this user.
        /// </summary>
        [Output("networkId")]
        public Output<string?> NetworkId { get; private set; } = null!;

        /// <summary>
        /// A note with additional information for the user.
        /// </summary>
        [Output("note")]
        public Output<string?> Note { get; private set; } = null!;

        /// <summary>
        /// The name of the site to associate the user with.
        /// </summary>
        [Output("site")]
        public Output<string> Site { get; private set; } = null!;

        /// <summary>
        /// Specifies whether this resource should tell the controller to "forget" the user on destroy. Defaults to `false`.
        /// </summary>
        [Output("skipForgetOnDestroy")]
        public Output<bool?> SkipForgetOnDestroy { get; private set; } = null!;

        /// <summary>
        /// The user group ID for the user.
        /// </summary>
        [Output("userGroupId")]
        public Output<string?> UserGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("unifi:index/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("unifi:index/user:User", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/pulumiverse/pulumi-unifi/releases/download/${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether this resource should just take over control of an existing user. Defaults to `true`.
        /// </summary>
        [Input("allowExisting")]
        public Input<bool>? AllowExisting { get; set; }

        /// <summary>
        /// Specifies whether this user should be blocked from the network.
        /// </summary>
        [Input("blocked")]
        public Input<bool>? Blocked { get; set; }

        /// <summary>
        /// Override the device fingerprint.
        /// </summary>
        [Input("devIdOverride")]
        public Input<int>? DevIdOverride { get; set; }

        /// <summary>
        /// A fixed IPv4 address for this user.
        /// </summary>
        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        /// <summary>
        /// The MAC address of the user.
        /// </summary>
        [Input("mac", required: true)]
        public Input<string> Mac { get; set; } = null!;

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network ID for this user.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// A note with additional information for the user.
        /// </summary>
        [Input("note")]
        public Input<string>? Note { get; set; }

        /// <summary>
        /// The name of the site to associate the user with.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        /// <summary>
        /// Specifies whether this resource should tell the controller to "forget" the user on destroy. Defaults to `false`.
        /// </summary>
        [Input("skipForgetOnDestroy")]
        public Input<bool>? SkipForgetOnDestroy { get; set; }

        /// <summary>
        /// The user group ID for the user.
        /// </summary>
        [Input("userGroupId")]
        public Input<string>? UserGroupId { get; set; }

        public UserArgs()
        {
        }
    }

    public sealed class UserState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether this resource should just take over control of an existing user. Defaults to `true`.
        /// </summary>
        [Input("allowExisting")]
        public Input<bool>? AllowExisting { get; set; }

        /// <summary>
        /// Specifies whether this user should be blocked from the network.
        /// </summary>
        [Input("blocked")]
        public Input<bool>? Blocked { get; set; }

        /// <summary>
        /// Override the device fingerprint.
        /// </summary>
        [Input("devIdOverride")]
        public Input<int>? DevIdOverride { get; set; }

        /// <summary>
        /// A fixed IPv4 address for this user.
        /// </summary>
        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        /// <summary>
        /// The hostname of the user.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// The IP address of the user.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The MAC address of the user.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network ID for this user.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// A note with additional information for the user.
        /// </summary>
        [Input("note")]
        public Input<string>? Note { get; set; }

        /// <summary>
        /// The name of the site to associate the user with.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        /// <summary>
        /// Specifies whether this resource should tell the controller to "forget" the user on destroy. Defaults to `false`.
        /// </summary>
        [Input("skipForgetOnDestroy")]
        public Input<bool>? SkipForgetOnDestroy { get; set; }

        /// <summary>
        /// The user group ID for the user.
        /// </summary>
        [Input("userGroupId")]
        public Input<string>? UserGroupId { get; set; }

        public UserState()
        {
        }
    }
}
