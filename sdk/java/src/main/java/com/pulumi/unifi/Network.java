// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.unifi;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.unifi.NetworkArgs;
import com.pulumi.unifi.Utilities;
import com.pulumi.unifi.inputs.NetworkState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `unifi.Network` manages WAN/LAN/VLAN networks.
 * 
 * ## Import
 * 
 * # import from provider configured site
 * 
 * ```sh
 *  $ pulumi import unifi:index/network:Network mynetwork 5dc28e5e9106d105bdc87217
 * ```
 * 
 * # import from another site
 * 
 * ```sh
 *  $ pulumi import unifi:index/network:Network mynetwork bfa2l6i7:5dc28e5e9106d105bdc87217
 * ```
 * 
 * # import network by name
 * 
 * ```sh
 *  $ pulumi import unifi:index/network:Network mynetwork name=LAN
 * ```
 * 
 */
@ResourceType(type="unifi:index/network:Network")
public class Network extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the IPv4 addresses for the DNS server to be returned from the DHCP server. Leave blank to disable this feature.
     * 
     */
    @Export(name="dhcpDns", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> dhcpDns;

    /**
     * @return Specifies the IPv4 addresses for the DNS server to be returned from the DHCP server. Leave blank to disable this feature.
     * 
     */
    public Output<Optional<List<String>>> dhcpDns() {
        return Codegen.optional(this.dhcpDns);
    }
    /**
     * Specifies whether DHCP is enabled or not on this network.
     * 
     */
    @Export(name="dhcpEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dhcpEnabled;

    /**
     * @return Specifies whether DHCP is enabled or not on this network.
     * 
     */
    public Output<Optional<Boolean>> dhcpEnabled() {
        return Codegen.optional(this.dhcpEnabled);
    }
    /**
     * Specifies the lease time for DHCP addresses. Defaults to `86400`.
     * 
     */
    @Export(name="dhcpLease", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> dhcpLease;

    /**
     * @return Specifies the lease time for DHCP addresses. Defaults to `86400`.
     * 
     */
    public Output<Optional<Integer>> dhcpLease() {
        return Codegen.optional(this.dhcpLease);
    }
    /**
     * Specifies whether DHCP relay is enabled or not on this network.
     * 
     */
    @Export(name="dhcpRelayEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dhcpRelayEnabled;

    /**
     * @return Specifies whether DHCP relay is enabled or not on this network.
     * 
     */
    public Output<Optional<Boolean>> dhcpRelayEnabled() {
        return Codegen.optional(this.dhcpRelayEnabled);
    }
    /**
     * The IPv4 address where the DHCP range of addresses starts.
     * 
     */
    @Export(name="dhcpStart", type=String.class, parameters={})
    private Output</* @Nullable */ String> dhcpStart;

    /**
     * @return The IPv4 address where the DHCP range of addresses starts.
     * 
     */
    public Output<Optional<String>> dhcpStart() {
        return Codegen.optional(this.dhcpStart);
    }
    /**
     * The IPv4 address where the DHCP range of addresses stops.
     * 
     */
    @Export(name="dhcpStop", type=String.class, parameters={})
    private Output</* @Nullable */ String> dhcpStop;

    /**
     * @return The IPv4 address where the DHCP range of addresses stops.
     * 
     */
    public Output<Optional<String>> dhcpStop() {
        return Codegen.optional(this.dhcpStop);
    }
    /**
     * Toggles on the DHCP boot options. Should be set to true when you want to have dhcpd*boot*filename, and dhcpd*boot*server to take effect.
     * 
     */
    @Export(name="dhcpdBootEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dhcpdBootEnabled;

    /**
     * @return Toggles on the DHCP boot options. Should be set to true when you want to have dhcpd*boot*filename, and dhcpd*boot*server to take effect.
     * 
     */
    public Output<Optional<Boolean>> dhcpdBootEnabled() {
        return Codegen.optional(this.dhcpdBootEnabled);
    }
    /**
     * Specifies the file to PXE boot from on the dhcpd*boot*server.
     * 
     */
    @Export(name="dhcpdBootFilename", type=String.class, parameters={})
    private Output</* @Nullable */ String> dhcpdBootFilename;

    /**
     * @return Specifies the file to PXE boot from on the dhcpd*boot*server.
     * 
     */
    public Output<Optional<String>> dhcpdBootFilename() {
        return Codegen.optional(this.dhcpdBootFilename);
    }
    /**
     * Specifies the IPv4 address of a TFTP server to network boot from.
     * 
     */
    @Export(name="dhcpdBootServer", type=String.class, parameters={})
    private Output</* @Nullable */ String> dhcpdBootServer;

    /**
     * @return Specifies the IPv4 address of a TFTP server to network boot from.
     * 
     */
    public Output<Optional<String>> dhcpdBootServer() {
        return Codegen.optional(this.dhcpdBootServer);
    }
    /**
     * The domain name of this network.
     * 
     */
    @Export(name="domainName", type=String.class, parameters={})
    private Output</* @Nullable */ String> domainName;

    /**
     * @return The domain name of this network.
     * 
     */
    public Output<Optional<String>> domainName() {
        return Codegen.optional(this.domainName);
    }
    /**
     * Specifies whether IGMP snooping is enabled or not.
     * 
     */
    @Export(name="igmpSnooping", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> igmpSnooping;

    /**
     * @return Specifies whether IGMP snooping is enabled or not.
     * 
     */
    public Output<Optional<Boolean>> igmpSnooping() {
        return Codegen.optional(this.igmpSnooping);
    }
    /**
     * Specifies which type of IPv6 connection to use. Defaults to `none`.
     * 
     */
    @Export(name="ipv6InterfaceType", type=String.class, parameters={})
    private Output</* @Nullable */ String> ipv6InterfaceType;

    /**
     * @return Specifies which type of IPv6 connection to use. Defaults to `none`.
     * 
     */
    public Output<Optional<String>> ipv6InterfaceType() {
        return Codegen.optional(this.ipv6InterfaceType);
    }
    /**
     * Specifies which WAN interface to use for IPv6 PD.
     * 
     */
    @Export(name="ipv6PdInterface", type=String.class, parameters={})
    private Output</* @Nullable */ String> ipv6PdInterface;

    /**
     * @return Specifies which WAN interface to use for IPv6 PD.
     * 
     */
    public Output<Optional<String>> ipv6PdInterface() {
        return Codegen.optional(this.ipv6PdInterface);
    }
    /**
     * Specifies the IPv6 Prefix ID.
     * 
     */
    @Export(name="ipv6PdPrefixid", type=String.class, parameters={})
    private Output</* @Nullable */ String> ipv6PdPrefixid;

    /**
     * @return Specifies the IPv6 Prefix ID.
     * 
     */
    public Output<Optional<String>> ipv6PdPrefixid() {
        return Codegen.optional(this.ipv6PdPrefixid);
    }
    /**
     * Specifies whether to enable router advertisements or not.
     * 
     */
    @Export(name="ipv6RaEnable", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> ipv6RaEnable;

    /**
     * @return Specifies whether to enable router advertisements or not.
     * 
     */
    public Output<Optional<Boolean>> ipv6RaEnable() {
        return Codegen.optional(this.ipv6RaEnable);
    }
    /**
     * Specifies the static IPv6 subnet when ipv6*interface*type is &#39;static&#39;.
     * 
     */
    @Export(name="ipv6StaticSubnet", type=String.class, parameters={})
    private Output</* @Nullable */ String> ipv6StaticSubnet;

    /**
     * @return Specifies the static IPv6 subnet when ipv6*interface*type is &#39;static&#39;.
     * 
     */
    public Output<Optional<String>> ipv6StaticSubnet() {
        return Codegen.optional(this.ipv6StaticSubnet);
    }
    /**
     * The name of the network.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the network.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The group of the network. Defaults to `LAN`.
     * 
     */
    @Export(name="networkGroup", type=String.class, parameters={})
    private Output</* @Nullable */ String> networkGroup;

    /**
     * @return The group of the network. Defaults to `LAN`.
     * 
     */
    public Output<Optional<String>> networkGroup() {
        return Codegen.optional(this.networkGroup);
    }
    /**
     * The purpose of the network. Must be one of `corporate`, `guest`, `wan`, or `vlan-only`.
     * 
     */
    @Export(name="purpose", type=String.class, parameters={})
    private Output<String> purpose;

    /**
     * @return The purpose of the network. Must be one of `corporate`, `guest`, `wan`, or `vlan-only`.
     * 
     */
    public Output<String> purpose() {
        return this.purpose;
    }
    /**
     * The name of the site to associate the network with.
     * 
     */
    @Export(name="site", type=String.class, parameters={})
    private Output<String> site;

    /**
     * @return The name of the site to associate the network with.
     * 
     */
    public Output<String> site() {
        return this.site;
    }
    /**
     * The subnet of the network. Must be a valid CIDR address.
     * 
     */
    @Export(name="subnet", type=String.class, parameters={})
    private Output</* @Nullable */ String> subnet;

    /**
     * @return The subnet of the network. Must be a valid CIDR address.
     * 
     */
    public Output<Optional<String>> subnet() {
        return Codegen.optional(this.subnet);
    }
    /**
     * The VLAN ID of the network.
     * 
     */
    @Export(name="vlanId", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> vlanId;

    /**
     * @return The VLAN ID of the network.
     * 
     */
    public Output<Optional<Integer>> vlanId() {
        return Codegen.optional(this.vlanId);
    }
    /**
     * DNS servers IPs of the WAN.
     * 
     */
    @Export(name="wanDns", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> wanDns;

    /**
     * @return DNS servers IPs of the WAN.
     * 
     */
    public Output<Optional<List<String>>> wanDns() {
        return Codegen.optional(this.wanDns);
    }
    /**
     * Specifies the WAN egress quality of service. Defaults to `0`.
     * 
     */
    @Export(name="wanEgressQos", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> wanEgressQos;

    /**
     * @return Specifies the WAN egress quality of service. Defaults to `0`.
     * 
     */
    public Output<Optional<Integer>> wanEgressQos() {
        return Codegen.optional(this.wanEgressQos);
    }
    /**
     * The IPv4 gateway of the WAN.
     * 
     */
    @Export(name="wanGateway", type=String.class, parameters={})
    private Output</* @Nullable */ String> wanGateway;

    /**
     * @return The IPv4 gateway of the WAN.
     * 
     */
    public Output<Optional<String>> wanGateway() {
        return Codegen.optional(this.wanGateway);
    }
    /**
     * The IPv4 address of the WAN.
     * 
     */
    @Export(name="wanIp", type=String.class, parameters={})
    private Output</* @Nullable */ String> wanIp;

    /**
     * @return The IPv4 address of the WAN.
     * 
     */
    public Output<Optional<String>> wanIp() {
        return Codegen.optional(this.wanIp);
    }
    /**
     * The IPv4 netmask of the WAN.
     * 
     */
    @Export(name="wanNetmask", type=String.class, parameters={})
    private Output</* @Nullable */ String> wanNetmask;

    /**
     * @return The IPv4 netmask of the WAN.
     * 
     */
    public Output<Optional<String>> wanNetmask() {
        return Codegen.optional(this.wanNetmask);
    }
    /**
     * Specifies the WAN network group. Must be one of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
     * 
     */
    @Export(name="wanNetworkgroup", type=String.class, parameters={})
    private Output</* @Nullable */ String> wanNetworkgroup;

    /**
     * @return Specifies the WAN network group. Must be one of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
     * 
     */
    public Output<Optional<String>> wanNetworkgroup() {
        return Codegen.optional(this.wanNetworkgroup);
    }
    /**
     * Specifies the IPV4 WAN connection type. Must be one of either `disabled`, `static`, `dhcp`, or `pppoe`.
     * 
     */
    @Export(name="wanType", type=String.class, parameters={})
    private Output</* @Nullable */ String> wanType;

    /**
     * @return Specifies the IPV4 WAN connection type. Must be one of either `disabled`, `static`, `dhcp`, or `pppoe`.
     * 
     */
    public Output<Optional<String>> wanType() {
        return Codegen.optional(this.wanType);
    }
    /**
     * Specifies the IPV4 WAN username.
     * 
     */
    @Export(name="wanUsername", type=String.class, parameters={})
    private Output</* @Nullable */ String> wanUsername;

    /**
     * @return Specifies the IPV4 WAN username.
     * 
     */
    public Output<Optional<String>> wanUsername() {
        return Codegen.optional(this.wanUsername);
    }
    /**
     * Specifies the IPV4 WAN password.
     * 
     */
    @Export(name="xWanPassword", type=String.class, parameters={})
    private Output</* @Nullable */ String> xWanPassword;

    /**
     * @return Specifies the IPV4 WAN password.
     * 
     */
    public Output<Optional<String>> xWanPassword() {
        return Codegen.optional(this.xWanPassword);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Network(String name) {
        this(name, NetworkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Network(String name, NetworkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Network(String name, NetworkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("unifi:index/network:Network", name, args == null ? NetworkArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Network(String name, Output<String> id, @Nullable NetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("unifi:index/network:Network", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Network get(String name, Output<String> id, @Nullable NetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Network(name, id, state, options);
    }
}