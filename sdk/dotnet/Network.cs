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
    /// `unifi.Network` manages WAN/LAN/VLAN networks.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Unifi = Pulumiverse.Unifi;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var vlanId = config.GetDouble("vlanId") ?? 10;
    ///     var vlan = new Unifi.Network("vlan", new()
    ///     {
    ///         Purpose = "corporate",
    ///         Subnet = "10.0.0.1/24",
    ///         VlanId = vlanId,
    ///         DhcpStart = "10.0.0.6",
    ///         DhcpStop = "10.0.0.254",
    ///         DhcpEnabled = true,
    ///     });
    /// 
    ///     var wan = new Unifi.Network("wan", new()
    ///     {
    ///         Purpose = "wan",
    ///         WanNetworkgroup = "WAN",
    ///         WanType = "pppoe",
    ///         WanIp = "192.168.1.1",
    ///         WanEgressQos = 1,
    ///         WanUsername = "username",
    ///         XWanPassword = "password",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// import from provider configured site
    /// 
    /// ```sh
    /// $ pulumi import unifi:index/network:Network mynetwork 5dc28e5e9106d105bdc87217
    /// ```
    /// 
    /// import from another site
    /// 
    /// ```sh
    /// $ pulumi import unifi:index/network:Network mynetwork bfa2l6i7:5dc28e5e9106d105bdc87217
    /// ```
    /// 
    /// import network by name
    /// 
    /// ```sh
    /// $ pulumi import unifi:index/network:Network mynetwork name=LAN
    /// ```
    /// </summary>
    [UnifiResourceType("unifi:index/network:Network")]
    public partial class Network : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the IPv4 addresses for the DNS server to be returned from the DHCP server. Leave blank to disable this feature.
        /// </summary>
        [Output("dhcpDns")]
        public Output<ImmutableArray<string>> DhcpDns { get; private set; } = null!;

        /// <summary>
        /// Specifies whether DHCP is enabled or not on this network.
        /// </summary>
        [Output("dhcpEnabled")]
        public Output<bool?> DhcpEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the lease time for DHCP addresses in seconds. Defaults to `86400`.
        /// </summary>
        [Output("dhcpLease")]
        public Output<int?> DhcpLease { get; private set; } = null!;

        /// <summary>
        /// Specifies whether DHCP relay is enabled or not on this network.
        /// </summary>
        [Output("dhcpRelayEnabled")]
        public Output<bool?> DhcpRelayEnabled { get; private set; } = null!;

        /// <summary>
        /// The IPv4 address where the DHCP range of addresses starts.
        /// </summary>
        [Output("dhcpStart")]
        public Output<string?> DhcpStart { get; private set; } = null!;

        /// <summary>
        /// The IPv4 address where the DHCP range of addresses stops.
        /// </summary>
        [Output("dhcpStop")]
        public Output<string?> DhcpStop { get; private set; } = null!;

        /// <summary>
        /// Specifies the IPv6 addresses for the DNS server to be returned from the DHCP server. Used if `dhcp_v6_dns_auto` is set to `false`.
        /// </summary>
        [Output("dhcpV6Dns")]
        public Output<ImmutableArray<string>> DhcpV6Dns { get; private set; } = null!;

        /// <summary>
        /// Specifies DNS source to propagate. If set `false` the entries in `dhcp_v6_dns` are used, the upstream entries otherwise Defaults to `true`.
        /// </summary>
        [Output("dhcpV6DnsAuto")]
        public Output<bool?> DhcpV6DnsAuto { get; private set; } = null!;

        /// <summary>
        /// Enable stateful DHCPv6 for static configuration.
        /// </summary>
        [Output("dhcpV6Enabled")]
        public Output<bool?> DhcpV6Enabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the lease time for DHCPv6 addresses in seconds. Defaults to `86400`.
        /// </summary>
        [Output("dhcpV6Lease")]
        public Output<int?> DhcpV6Lease { get; private set; } = null!;

        /// <summary>
        /// Start address of the DHCPv6 range. Used in static DHCPv6 configuration.
        /// </summary>
        [Output("dhcpV6Start")]
        public Output<string?> DhcpV6Start { get; private set; } = null!;

        /// <summary>
        /// End address of the DHCPv6 range. Used in static DHCPv6 configuration.
        /// </summary>
        [Output("dhcpV6Stop")]
        public Output<string?> DhcpV6Stop { get; private set; } = null!;

        /// <summary>
        /// Toggles on the DHCP boot options. Should be set to true when you want to have dhcpd*boot*filename, and dhcpd*boot*server to take effect.
        /// </summary>
        [Output("dhcpdBootEnabled")]
        public Output<bool?> DhcpdBootEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies the file to PXE boot from on the dhcpd*boot*server.
        /// </summary>
        [Output("dhcpdBootFilename")]
        public Output<string?> DhcpdBootFilename { get; private set; } = null!;

        /// <summary>
        /// Specifies the IPv4 address of a TFTP server to network boot from.
        /// </summary>
        [Output("dhcpdBootServer")]
        public Output<string?> DhcpdBootServer { get; private set; } = null!;

        /// <summary>
        /// The domain name of this network.
        /// </summary>
        [Output("domainName")]
        public Output<string?> DomainName { get; private set; } = null!;

        /// <summary>
        /// Specifies whether IGMP snooping is enabled or not.
        /// </summary>
        [Output("igmpSnooping")]
        public Output<bool?> IgmpSnooping { get; private set; } = null!;

        /// <summary>
        /// Specifies whether this network should be allowed to access the internet or not. Defaults to `true`.
        /// </summary>
        [Output("internetAccessEnabled")]
        public Output<bool?> InternetAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies whether this network should be allowed to access other local networks or not. Defaults to `true`.
        /// </summary>
        [Output("intraNetworkAccessEnabled")]
        public Output<bool?> IntraNetworkAccessEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies which type of IPv6 connection to use. Must be one of either `static`, `pd`, or `none`. Defaults to `none`.
        /// </summary>
        [Output("ipv6InterfaceType")]
        public Output<string?> Ipv6InterfaceType { get; private set; } = null!;

        /// <summary>
        /// Specifies which WAN interface to use for IPv6 PD. Must be one of either `wan` or `wan2`.
        /// </summary>
        [Output("ipv6PdInterface")]
        public Output<string?> Ipv6PdInterface { get; private set; } = null!;

        /// <summary>
        /// Specifies the IPv6 Prefix ID.
        /// </summary>
        [Output("ipv6PdPrefixid")]
        public Output<string?> Ipv6PdPrefixid { get; private set; } = null!;

        /// <summary>
        /// Start address of the DHCPv6 range. Used if `ipv6_interface_type` is set to `pd`.
        /// </summary>
        [Output("ipv6PdStart")]
        public Output<string?> Ipv6PdStart { get; private set; } = null!;

        /// <summary>
        /// End address of the DHCPv6 range. Used if `ipv6_interface_type` is set to `pd`.
        /// </summary>
        [Output("ipv6PdStop")]
        public Output<string?> Ipv6PdStop { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable router advertisements or not.
        /// </summary>
        [Output("ipv6RaEnable")]
        public Output<bool?> Ipv6RaEnable { get; private set; } = null!;

        /// <summary>
        /// Lifetime in which the address can be used. Address becomes deprecated afterwards. Must be lower than or equal to `ipv6_ra_valid_lifetime` Defaults to `14400`.
        /// </summary>
        [Output("ipv6RaPreferredLifetime")]
        public Output<int?> Ipv6RaPreferredLifetime { get; private set; } = null!;

        /// <summary>
        /// IPv6 router advertisement priority. Must be one of either `high`, `medium`, or `low`
        /// </summary>
        [Output("ipv6RaPriority")]
        public Output<string?> Ipv6RaPriority { get; private set; } = null!;

        /// <summary>
        /// Total lifetime in which the address can be used. Must be equal to or greater than `ipv6_ra_preferred_lifetime`. Defaults to `86400`.
        /// </summary>
        [Output("ipv6RaValidLifetime")]
        public Output<int?> Ipv6RaValidLifetime { get; private set; } = null!;

        /// <summary>
        /// Specifies the static IPv6 subnet when `ipv6_interface_type` is 'static'.
        /// </summary>
        [Output("ipv6StaticSubnet")]
        public Output<string?> Ipv6StaticSubnet { get; private set; } = null!;

        /// <summary>
        /// Specifies whether Multicast DNS (mDNS) is enabled or not on the network (Controller &gt;=v7).
        /// </summary>
        [Output("multicastDns")]
        public Output<bool?> MulticastDns { get; private set; } = null!;

        /// <summary>
        /// The name of the network.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The group of the network. Defaults to `LAN`.
        /// </summary>
        [Output("networkGroup")]
        public Output<string?> NetworkGroup { get; private set; } = null!;

        /// <summary>
        /// The purpose of the network. Must be one of `corporate`, `guest`, `wan`, or `vlan-only`.
        /// </summary>
        [Output("purpose")]
        public Output<string> Purpose { get; private set; } = null!;

        /// <summary>
        /// The name of the site to associate the network with.
        /// </summary>
        [Output("site")]
        public Output<string> Site { get; private set; } = null!;

        /// <summary>
        /// The subnet of the network. Must be a valid CIDR address.
        /// </summary>
        [Output("subnet")]
        public Output<string?> Subnet { get; private set; } = null!;

        /// <summary>
        /// The VLAN ID of the network.
        /// </summary>
        [Output("vlanId")]
        public Output<int?> VlanId { get; private set; } = null!;

        /// <summary>
        /// Specifies the IPv6 prefix size to request from ISP. Must be between 48 and 64.
        /// </summary>
        [Output("wanDhcpV6PdSize")]
        public Output<int?> WanDhcpV6PdSize { get; private set; } = null!;

        /// <summary>
        /// DNS servers IPs of the WAN.
        /// </summary>
        [Output("wanDns")]
        public Output<ImmutableArray<string>> WanDns { get; private set; } = null!;

        /// <summary>
        /// Specifies the WAN egress quality of service. Defaults to `0`.
        /// </summary>
        [Output("wanEgressQos")]
        public Output<int?> WanEgressQos { get; private set; } = null!;

        /// <summary>
        /// The IPv4 gateway of the WAN.
        /// </summary>
        [Output("wanGateway")]
        public Output<string?> WanGateway { get; private set; } = null!;

        /// <summary>
        /// The IPv6 gateway of the WAN.
        /// </summary>
        [Output("wanGatewayV6")]
        public Output<string?> WanGatewayV6 { get; private set; } = null!;

        /// <summary>
        /// The IPv4 address of the WAN.
        /// </summary>
        [Output("wanIp")]
        public Output<string?> WanIp { get; private set; } = null!;

        /// <summary>
        /// The IPv6 address of the WAN.
        /// </summary>
        [Output("wanIpv6")]
        public Output<string?> WanIpv6 { get; private set; } = null!;

        /// <summary>
        /// The IPv4 netmask of the WAN.
        /// </summary>
        [Output("wanNetmask")]
        public Output<string?> WanNetmask { get; private set; } = null!;

        /// <summary>
        /// Specifies the WAN network group. Must be one of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
        /// </summary>
        [Output("wanNetworkgroup")]
        public Output<string?> WanNetworkgroup { get; private set; } = null!;

        /// <summary>
        /// The IPv6 prefix length of the WAN. Must be between 1 and 128.
        /// </summary>
        [Output("wanPrefixlen")]
        public Output<int?> WanPrefixlen { get; private set; } = null!;

        /// <summary>
        /// Specifies the IPV4 WAN connection type. Must be one of either `disabled`, `static`, `dhcp`, or `pppoe`.
        /// </summary>
        [Output("wanType")]
        public Output<string?> WanType { get; private set; } = null!;

        /// <summary>
        /// Specifies the IPV6 WAN connection type. Must be one of either `disabled`, `static`, or `dhcpv6`.
        /// </summary>
        [Output("wanTypeV6")]
        public Output<string?> WanTypeV6 { get; private set; } = null!;

        /// <summary>
        /// Specifies the IPV4 WAN username.
        /// </summary>
        [Output("wanUsername")]
        public Output<string?> WanUsername { get; private set; } = null!;

        /// <summary>
        /// Specifies the IPV4 WAN password.
        /// </summary>
        [Output("xWanPassword")]
        public Output<string?> XWanPassword { get; private set; } = null!;


        /// <summary>
        /// Create a Network resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Network(string name, NetworkArgs args, CustomResourceOptions? options = null)
            : base("unifi:index/network:Network", name, args ?? new NetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Network(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
            : base("unifi:index/network:Network", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Network resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Network Get(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new Network(name, id, state, options);
        }
    }

    public sealed class NetworkArgs : global::Pulumi.ResourceArgs
    {
        [Input("dhcpDns")]
        private InputList<string>? _dhcpDns;

        /// <summary>
        /// Specifies the IPv4 addresses for the DNS server to be returned from the DHCP server. Leave blank to disable this feature.
        /// </summary>
        public InputList<string> DhcpDns
        {
            get => _dhcpDns ?? (_dhcpDns = new InputList<string>());
            set => _dhcpDns = value;
        }

        /// <summary>
        /// Specifies whether DHCP is enabled or not on this network.
        /// </summary>
        [Input("dhcpEnabled")]
        public Input<bool>? DhcpEnabled { get; set; }

        /// <summary>
        /// Specifies the lease time for DHCP addresses in seconds. Defaults to `86400`.
        /// </summary>
        [Input("dhcpLease")]
        public Input<int>? DhcpLease { get; set; }

        /// <summary>
        /// Specifies whether DHCP relay is enabled or not on this network.
        /// </summary>
        [Input("dhcpRelayEnabled")]
        public Input<bool>? DhcpRelayEnabled { get; set; }

        /// <summary>
        /// The IPv4 address where the DHCP range of addresses starts.
        /// </summary>
        [Input("dhcpStart")]
        public Input<string>? DhcpStart { get; set; }

        /// <summary>
        /// The IPv4 address where the DHCP range of addresses stops.
        /// </summary>
        [Input("dhcpStop")]
        public Input<string>? DhcpStop { get; set; }

        [Input("dhcpV6Dns")]
        private InputList<string>? _dhcpV6Dns;

        /// <summary>
        /// Specifies the IPv6 addresses for the DNS server to be returned from the DHCP server. Used if `dhcp_v6_dns_auto` is set to `false`.
        /// </summary>
        public InputList<string> DhcpV6Dns
        {
            get => _dhcpV6Dns ?? (_dhcpV6Dns = new InputList<string>());
            set => _dhcpV6Dns = value;
        }

        /// <summary>
        /// Specifies DNS source to propagate. If set `false` the entries in `dhcp_v6_dns` are used, the upstream entries otherwise Defaults to `true`.
        /// </summary>
        [Input("dhcpV6DnsAuto")]
        public Input<bool>? DhcpV6DnsAuto { get; set; }

        /// <summary>
        /// Enable stateful DHCPv6 for static configuration.
        /// </summary>
        [Input("dhcpV6Enabled")]
        public Input<bool>? DhcpV6Enabled { get; set; }

        /// <summary>
        /// Specifies the lease time for DHCPv6 addresses in seconds. Defaults to `86400`.
        /// </summary>
        [Input("dhcpV6Lease")]
        public Input<int>? DhcpV6Lease { get; set; }

        /// <summary>
        /// Start address of the DHCPv6 range. Used in static DHCPv6 configuration.
        /// </summary>
        [Input("dhcpV6Start")]
        public Input<string>? DhcpV6Start { get; set; }

        /// <summary>
        /// End address of the DHCPv6 range. Used in static DHCPv6 configuration.
        /// </summary>
        [Input("dhcpV6Stop")]
        public Input<string>? DhcpV6Stop { get; set; }

        /// <summary>
        /// Toggles on the DHCP boot options. Should be set to true when you want to have dhcpd*boot*filename, and dhcpd*boot*server to take effect.
        /// </summary>
        [Input("dhcpdBootEnabled")]
        public Input<bool>? DhcpdBootEnabled { get; set; }

        /// <summary>
        /// Specifies the file to PXE boot from on the dhcpd*boot*server.
        /// </summary>
        [Input("dhcpdBootFilename")]
        public Input<string>? DhcpdBootFilename { get; set; }

        /// <summary>
        /// Specifies the IPv4 address of a TFTP server to network boot from.
        /// </summary>
        [Input("dhcpdBootServer")]
        public Input<string>? DhcpdBootServer { get; set; }

        /// <summary>
        /// The domain name of this network.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Specifies whether IGMP snooping is enabled or not.
        /// </summary>
        [Input("igmpSnooping")]
        public Input<bool>? IgmpSnooping { get; set; }

        /// <summary>
        /// Specifies whether this network should be allowed to access the internet or not. Defaults to `true`.
        /// </summary>
        [Input("internetAccessEnabled")]
        public Input<bool>? InternetAccessEnabled { get; set; }

        /// <summary>
        /// Specifies whether this network should be allowed to access other local networks or not. Defaults to `true`.
        /// </summary>
        [Input("intraNetworkAccessEnabled")]
        public Input<bool>? IntraNetworkAccessEnabled { get; set; }

        /// <summary>
        /// Specifies which type of IPv6 connection to use. Must be one of either `static`, `pd`, or `none`. Defaults to `none`.
        /// </summary>
        [Input("ipv6InterfaceType")]
        public Input<string>? Ipv6InterfaceType { get; set; }

        /// <summary>
        /// Specifies which WAN interface to use for IPv6 PD. Must be one of either `wan` or `wan2`.
        /// </summary>
        [Input("ipv6PdInterface")]
        public Input<string>? Ipv6PdInterface { get; set; }

        /// <summary>
        /// Specifies the IPv6 Prefix ID.
        /// </summary>
        [Input("ipv6PdPrefixid")]
        public Input<string>? Ipv6PdPrefixid { get; set; }

        /// <summary>
        /// Start address of the DHCPv6 range. Used if `ipv6_interface_type` is set to `pd`.
        /// </summary>
        [Input("ipv6PdStart")]
        public Input<string>? Ipv6PdStart { get; set; }

        /// <summary>
        /// End address of the DHCPv6 range. Used if `ipv6_interface_type` is set to `pd`.
        /// </summary>
        [Input("ipv6PdStop")]
        public Input<string>? Ipv6PdStop { get; set; }

        /// <summary>
        /// Specifies whether to enable router advertisements or not.
        /// </summary>
        [Input("ipv6RaEnable")]
        public Input<bool>? Ipv6RaEnable { get; set; }

        /// <summary>
        /// Lifetime in which the address can be used. Address becomes deprecated afterwards. Must be lower than or equal to `ipv6_ra_valid_lifetime` Defaults to `14400`.
        /// </summary>
        [Input("ipv6RaPreferredLifetime")]
        public Input<int>? Ipv6RaPreferredLifetime { get; set; }

        /// <summary>
        /// IPv6 router advertisement priority. Must be one of either `high`, `medium`, or `low`
        /// </summary>
        [Input("ipv6RaPriority")]
        public Input<string>? Ipv6RaPriority { get; set; }

        /// <summary>
        /// Total lifetime in which the address can be used. Must be equal to or greater than `ipv6_ra_preferred_lifetime`. Defaults to `86400`.
        /// </summary>
        [Input("ipv6RaValidLifetime")]
        public Input<int>? Ipv6RaValidLifetime { get; set; }

        /// <summary>
        /// Specifies the static IPv6 subnet when `ipv6_interface_type` is 'static'.
        /// </summary>
        [Input("ipv6StaticSubnet")]
        public Input<string>? Ipv6StaticSubnet { get; set; }

        /// <summary>
        /// Specifies whether Multicast DNS (mDNS) is enabled or not on the network (Controller &gt;=v7).
        /// </summary>
        [Input("multicastDns")]
        public Input<bool>? MulticastDns { get; set; }

        /// <summary>
        /// The name of the network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The group of the network. Defaults to `LAN`.
        /// </summary>
        [Input("networkGroup")]
        public Input<string>? NetworkGroup { get; set; }

        /// <summary>
        /// The purpose of the network. Must be one of `corporate`, `guest`, `wan`, or `vlan-only`.
        /// </summary>
        [Input("purpose", required: true)]
        public Input<string> Purpose { get; set; } = null!;

        /// <summary>
        /// The name of the site to associate the network with.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        /// <summary>
        /// The subnet of the network. Must be a valid CIDR address.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// The VLAN ID of the network.
        /// </summary>
        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        /// <summary>
        /// Specifies the IPv6 prefix size to request from ISP. Must be between 48 and 64.
        /// </summary>
        [Input("wanDhcpV6PdSize")]
        public Input<int>? WanDhcpV6PdSize { get; set; }

        [Input("wanDns")]
        private InputList<string>? _wanDns;

        /// <summary>
        /// DNS servers IPs of the WAN.
        /// </summary>
        public InputList<string> WanDns
        {
            get => _wanDns ?? (_wanDns = new InputList<string>());
            set => _wanDns = value;
        }

        /// <summary>
        /// Specifies the WAN egress quality of service. Defaults to `0`.
        /// </summary>
        [Input("wanEgressQos")]
        public Input<int>? WanEgressQos { get; set; }

        /// <summary>
        /// The IPv4 gateway of the WAN.
        /// </summary>
        [Input("wanGateway")]
        public Input<string>? WanGateway { get; set; }

        /// <summary>
        /// The IPv6 gateway of the WAN.
        /// </summary>
        [Input("wanGatewayV6")]
        public Input<string>? WanGatewayV6 { get; set; }

        /// <summary>
        /// The IPv4 address of the WAN.
        /// </summary>
        [Input("wanIp")]
        public Input<string>? WanIp { get; set; }

        /// <summary>
        /// The IPv6 address of the WAN.
        /// </summary>
        [Input("wanIpv6")]
        public Input<string>? WanIpv6 { get; set; }

        /// <summary>
        /// The IPv4 netmask of the WAN.
        /// </summary>
        [Input("wanNetmask")]
        public Input<string>? WanNetmask { get; set; }

        /// <summary>
        /// Specifies the WAN network group. Must be one of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
        /// </summary>
        [Input("wanNetworkgroup")]
        public Input<string>? WanNetworkgroup { get; set; }

        /// <summary>
        /// The IPv6 prefix length of the WAN. Must be between 1 and 128.
        /// </summary>
        [Input("wanPrefixlen")]
        public Input<int>? WanPrefixlen { get; set; }

        /// <summary>
        /// Specifies the IPV4 WAN connection type. Must be one of either `disabled`, `static`, `dhcp`, or `pppoe`.
        /// </summary>
        [Input("wanType")]
        public Input<string>? WanType { get; set; }

        /// <summary>
        /// Specifies the IPV6 WAN connection type. Must be one of either `disabled`, `static`, or `dhcpv6`.
        /// </summary>
        [Input("wanTypeV6")]
        public Input<string>? WanTypeV6 { get; set; }

        /// <summary>
        /// Specifies the IPV4 WAN username.
        /// </summary>
        [Input("wanUsername")]
        public Input<string>? WanUsername { get; set; }

        /// <summary>
        /// Specifies the IPV4 WAN password.
        /// </summary>
        [Input("xWanPassword")]
        public Input<string>? XWanPassword { get; set; }

        public NetworkArgs()
        {
        }
        public static new NetworkArgs Empty => new NetworkArgs();
    }

    public sealed class NetworkState : global::Pulumi.ResourceArgs
    {
        [Input("dhcpDns")]
        private InputList<string>? _dhcpDns;

        /// <summary>
        /// Specifies the IPv4 addresses for the DNS server to be returned from the DHCP server. Leave blank to disable this feature.
        /// </summary>
        public InputList<string> DhcpDns
        {
            get => _dhcpDns ?? (_dhcpDns = new InputList<string>());
            set => _dhcpDns = value;
        }

        /// <summary>
        /// Specifies whether DHCP is enabled or not on this network.
        /// </summary>
        [Input("dhcpEnabled")]
        public Input<bool>? DhcpEnabled { get; set; }

        /// <summary>
        /// Specifies the lease time for DHCP addresses in seconds. Defaults to `86400`.
        /// </summary>
        [Input("dhcpLease")]
        public Input<int>? DhcpLease { get; set; }

        /// <summary>
        /// Specifies whether DHCP relay is enabled or not on this network.
        /// </summary>
        [Input("dhcpRelayEnabled")]
        public Input<bool>? DhcpRelayEnabled { get; set; }

        /// <summary>
        /// The IPv4 address where the DHCP range of addresses starts.
        /// </summary>
        [Input("dhcpStart")]
        public Input<string>? DhcpStart { get; set; }

        /// <summary>
        /// The IPv4 address where the DHCP range of addresses stops.
        /// </summary>
        [Input("dhcpStop")]
        public Input<string>? DhcpStop { get; set; }

        [Input("dhcpV6Dns")]
        private InputList<string>? _dhcpV6Dns;

        /// <summary>
        /// Specifies the IPv6 addresses for the DNS server to be returned from the DHCP server. Used if `dhcp_v6_dns_auto` is set to `false`.
        /// </summary>
        public InputList<string> DhcpV6Dns
        {
            get => _dhcpV6Dns ?? (_dhcpV6Dns = new InputList<string>());
            set => _dhcpV6Dns = value;
        }

        /// <summary>
        /// Specifies DNS source to propagate. If set `false` the entries in `dhcp_v6_dns` are used, the upstream entries otherwise Defaults to `true`.
        /// </summary>
        [Input("dhcpV6DnsAuto")]
        public Input<bool>? DhcpV6DnsAuto { get; set; }

        /// <summary>
        /// Enable stateful DHCPv6 for static configuration.
        /// </summary>
        [Input("dhcpV6Enabled")]
        public Input<bool>? DhcpV6Enabled { get; set; }

        /// <summary>
        /// Specifies the lease time for DHCPv6 addresses in seconds. Defaults to `86400`.
        /// </summary>
        [Input("dhcpV6Lease")]
        public Input<int>? DhcpV6Lease { get; set; }

        /// <summary>
        /// Start address of the DHCPv6 range. Used in static DHCPv6 configuration.
        /// </summary>
        [Input("dhcpV6Start")]
        public Input<string>? DhcpV6Start { get; set; }

        /// <summary>
        /// End address of the DHCPv6 range. Used in static DHCPv6 configuration.
        /// </summary>
        [Input("dhcpV6Stop")]
        public Input<string>? DhcpV6Stop { get; set; }

        /// <summary>
        /// Toggles on the DHCP boot options. Should be set to true when you want to have dhcpd*boot*filename, and dhcpd*boot*server to take effect.
        /// </summary>
        [Input("dhcpdBootEnabled")]
        public Input<bool>? DhcpdBootEnabled { get; set; }

        /// <summary>
        /// Specifies the file to PXE boot from on the dhcpd*boot*server.
        /// </summary>
        [Input("dhcpdBootFilename")]
        public Input<string>? DhcpdBootFilename { get; set; }

        /// <summary>
        /// Specifies the IPv4 address of a TFTP server to network boot from.
        /// </summary>
        [Input("dhcpdBootServer")]
        public Input<string>? DhcpdBootServer { get; set; }

        /// <summary>
        /// The domain name of this network.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Specifies whether IGMP snooping is enabled or not.
        /// </summary>
        [Input("igmpSnooping")]
        public Input<bool>? IgmpSnooping { get; set; }

        /// <summary>
        /// Specifies whether this network should be allowed to access the internet or not. Defaults to `true`.
        /// </summary>
        [Input("internetAccessEnabled")]
        public Input<bool>? InternetAccessEnabled { get; set; }

        /// <summary>
        /// Specifies whether this network should be allowed to access other local networks or not. Defaults to `true`.
        /// </summary>
        [Input("intraNetworkAccessEnabled")]
        public Input<bool>? IntraNetworkAccessEnabled { get; set; }

        /// <summary>
        /// Specifies which type of IPv6 connection to use. Must be one of either `static`, `pd`, or `none`. Defaults to `none`.
        /// </summary>
        [Input("ipv6InterfaceType")]
        public Input<string>? Ipv6InterfaceType { get; set; }

        /// <summary>
        /// Specifies which WAN interface to use for IPv6 PD. Must be one of either `wan` or `wan2`.
        /// </summary>
        [Input("ipv6PdInterface")]
        public Input<string>? Ipv6PdInterface { get; set; }

        /// <summary>
        /// Specifies the IPv6 Prefix ID.
        /// </summary>
        [Input("ipv6PdPrefixid")]
        public Input<string>? Ipv6PdPrefixid { get; set; }

        /// <summary>
        /// Start address of the DHCPv6 range. Used if `ipv6_interface_type` is set to `pd`.
        /// </summary>
        [Input("ipv6PdStart")]
        public Input<string>? Ipv6PdStart { get; set; }

        /// <summary>
        /// End address of the DHCPv6 range. Used if `ipv6_interface_type` is set to `pd`.
        /// </summary>
        [Input("ipv6PdStop")]
        public Input<string>? Ipv6PdStop { get; set; }

        /// <summary>
        /// Specifies whether to enable router advertisements or not.
        /// </summary>
        [Input("ipv6RaEnable")]
        public Input<bool>? Ipv6RaEnable { get; set; }

        /// <summary>
        /// Lifetime in which the address can be used. Address becomes deprecated afterwards. Must be lower than or equal to `ipv6_ra_valid_lifetime` Defaults to `14400`.
        /// </summary>
        [Input("ipv6RaPreferredLifetime")]
        public Input<int>? Ipv6RaPreferredLifetime { get; set; }

        /// <summary>
        /// IPv6 router advertisement priority. Must be one of either `high`, `medium`, or `low`
        /// </summary>
        [Input("ipv6RaPriority")]
        public Input<string>? Ipv6RaPriority { get; set; }

        /// <summary>
        /// Total lifetime in which the address can be used. Must be equal to or greater than `ipv6_ra_preferred_lifetime`. Defaults to `86400`.
        /// </summary>
        [Input("ipv6RaValidLifetime")]
        public Input<int>? Ipv6RaValidLifetime { get; set; }

        /// <summary>
        /// Specifies the static IPv6 subnet when `ipv6_interface_type` is 'static'.
        /// </summary>
        [Input("ipv6StaticSubnet")]
        public Input<string>? Ipv6StaticSubnet { get; set; }

        /// <summary>
        /// Specifies whether Multicast DNS (mDNS) is enabled or not on the network (Controller &gt;=v7).
        /// </summary>
        [Input("multicastDns")]
        public Input<bool>? MulticastDns { get; set; }

        /// <summary>
        /// The name of the network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The group of the network. Defaults to `LAN`.
        /// </summary>
        [Input("networkGroup")]
        public Input<string>? NetworkGroup { get; set; }

        /// <summary>
        /// The purpose of the network. Must be one of `corporate`, `guest`, `wan`, or `vlan-only`.
        /// </summary>
        [Input("purpose")]
        public Input<string>? Purpose { get; set; }

        /// <summary>
        /// The name of the site to associate the network with.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        /// <summary>
        /// The subnet of the network. Must be a valid CIDR address.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// The VLAN ID of the network.
        /// </summary>
        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        /// <summary>
        /// Specifies the IPv6 prefix size to request from ISP. Must be between 48 and 64.
        /// </summary>
        [Input("wanDhcpV6PdSize")]
        public Input<int>? WanDhcpV6PdSize { get; set; }

        [Input("wanDns")]
        private InputList<string>? _wanDns;

        /// <summary>
        /// DNS servers IPs of the WAN.
        /// </summary>
        public InputList<string> WanDns
        {
            get => _wanDns ?? (_wanDns = new InputList<string>());
            set => _wanDns = value;
        }

        /// <summary>
        /// Specifies the WAN egress quality of service. Defaults to `0`.
        /// </summary>
        [Input("wanEgressQos")]
        public Input<int>? WanEgressQos { get; set; }

        /// <summary>
        /// The IPv4 gateway of the WAN.
        /// </summary>
        [Input("wanGateway")]
        public Input<string>? WanGateway { get; set; }

        /// <summary>
        /// The IPv6 gateway of the WAN.
        /// </summary>
        [Input("wanGatewayV6")]
        public Input<string>? WanGatewayV6 { get; set; }

        /// <summary>
        /// The IPv4 address of the WAN.
        /// </summary>
        [Input("wanIp")]
        public Input<string>? WanIp { get; set; }

        /// <summary>
        /// The IPv6 address of the WAN.
        /// </summary>
        [Input("wanIpv6")]
        public Input<string>? WanIpv6 { get; set; }

        /// <summary>
        /// The IPv4 netmask of the WAN.
        /// </summary>
        [Input("wanNetmask")]
        public Input<string>? WanNetmask { get; set; }

        /// <summary>
        /// Specifies the WAN network group. Must be one of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
        /// </summary>
        [Input("wanNetworkgroup")]
        public Input<string>? WanNetworkgroup { get; set; }

        /// <summary>
        /// The IPv6 prefix length of the WAN. Must be between 1 and 128.
        /// </summary>
        [Input("wanPrefixlen")]
        public Input<int>? WanPrefixlen { get; set; }

        /// <summary>
        /// Specifies the IPV4 WAN connection type. Must be one of either `disabled`, `static`, `dhcp`, or `pppoe`.
        /// </summary>
        [Input("wanType")]
        public Input<string>? WanType { get; set; }

        /// <summary>
        /// Specifies the IPV6 WAN connection type. Must be one of either `disabled`, `static`, or `dhcpv6`.
        /// </summary>
        [Input("wanTypeV6")]
        public Input<string>? WanTypeV6 { get; set; }

        /// <summary>
        /// Specifies the IPV4 WAN username.
        /// </summary>
        [Input("wanUsername")]
        public Input<string>? WanUsername { get; set; }

        /// <summary>
        /// Specifies the IPV4 WAN password.
        /// </summary>
        [Input("xWanPassword")]
        public Input<string>? XWanPassword { get; set; }

        public NetworkState()
        {
        }
        public static new NetworkState Empty => new NetworkState();
    }
}
