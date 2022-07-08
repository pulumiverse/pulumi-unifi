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
    /// `unifi.FirewallGroup` manages groups of addresses or ports for use in firewall rules (`unifi.FirewallRule`).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Unifi = Pulumiverse.Unifi;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var laptopIps = config.RequireObject&lt;dynamic&gt;("laptopIps");
    ///         var canPrint = new Unifi.FirewallGroup("canPrint", new Unifi.FirewallGroupArgs
    ///         {
    ///             Type = "address-group",
    ///             Members = laptopIps,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [UnifiResourceType("unifi:index/firewallGroup:FirewallGroup")]
    public partial class FirewallGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The members of the firewall group.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The name of the firewall group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the site to associate the firewall group with.
        /// </summary>
        [Output("site")]
        public Output<string> Site { get; private set; } = null!;

        /// <summary>
        /// The type of the firewall group. Must be one of: `address-group`, `port-group`, or `ipv6-address-group`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallGroup(string name, FirewallGroupArgs args, CustomResourceOptions? options = null)
            : base("unifi:index/firewallGroup:FirewallGroup", name, args ?? new FirewallGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallGroup(string name, Input<string> id, FirewallGroupState? state = null, CustomResourceOptions? options = null)
            : base("unifi:index/firewallGroup:FirewallGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallGroup Get(string name, Input<string> id, FirewallGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallGroup(name, id, state, options);
        }
    }

    public sealed class FirewallGroupArgs : Pulumi.ResourceArgs
    {
        [Input("members", required: true)]
        private InputList<string>? _members;

        /// <summary>
        /// The members of the firewall group.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The name of the firewall group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the site to associate the firewall group with.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        /// <summary>
        /// The type of the firewall group. Must be one of: `address-group`, `port-group`, or `ipv6-address-group`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FirewallGroupArgs()
        {
        }
    }

    public sealed class FirewallGroupState : Pulumi.ResourceArgs
    {
        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// The members of the firewall group.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The name of the firewall group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the site to associate the firewall group with.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        /// <summary>
        /// The type of the firewall group. Must be one of: `address-group`, `port-group`, or `ipv6-address-group`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FirewallGroupState()
        {
        }
    }
}