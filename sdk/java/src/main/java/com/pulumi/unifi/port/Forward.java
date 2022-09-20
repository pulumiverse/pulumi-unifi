// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.unifi.port;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.unifi.Utilities;
import com.pulumi.unifi.port.ForwardArgs;
import com.pulumi.unifi.port.inputs.ForwardState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `unifi.port.Forward` manages a port forwarding rule on the gateway.
 * 
 */
@ResourceType(type="unifi:port/forward:Forward")
public class Forward extends com.pulumi.resources.CustomResource {
    /**
     * The destination port for the forwarding.
     * 
     */
    @Export(name="dstPort", type=String.class, parameters={})
    private Output</* @Nullable */ String> dstPort;

    /**
     * @return The destination port for the forwarding.
     * 
     */
    public Output<Optional<String>> dstPort() {
        return Codegen.optional(this.dstPort);
    }
    /**
     * Specifies whether the port forwarding rule is enabled or not. This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
     * 
     * @deprecated
     * This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
     * 
     */
    @Deprecated /* This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration. */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Specifies whether the port forwarding rule is enabled or not. This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * The IPv4 address to forward traffic to.
     * 
     */
    @Export(name="fwdIp", type=String.class, parameters={})
    private Output</* @Nullable */ String> fwdIp;

    /**
     * @return The IPv4 address to forward traffic to.
     * 
     */
    public Output<Optional<String>> fwdIp() {
        return Codegen.optional(this.fwdIp);
    }
    /**
     * The port to forward traffic to.
     * 
     */
    @Export(name="fwdPort", type=String.class, parameters={})
    private Output</* @Nullable */ String> fwdPort;

    /**
     * @return The port to forward traffic to.
     * 
     */
    public Output<Optional<String>> fwdPort() {
        return Codegen.optional(this.fwdPort);
    }
    /**
     * Specifies whether to log forwarded traffic or not. Defaults to `false`.
     * 
     */
    @Export(name="log", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> log;

    /**
     * @return Specifies whether to log forwarded traffic or not. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> log() {
        return Codegen.optional(this.log);
    }
    /**
     * The name of the port forwarding rule.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the port forwarding rule.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The port forwarding interface. Can be `wan`, `wan2`, or `both`.
     * 
     */
    @Export(name="portForwardInterface", type=String.class, parameters={})
    private Output</* @Nullable */ String> portForwardInterface;

    /**
     * @return The port forwarding interface. Can be `wan`, `wan2`, or `both`.
     * 
     */
    public Output<Optional<String>> portForwardInterface() {
        return Codegen.optional(this.portForwardInterface);
    }
    /**
     * The protocol for the port forwarding rule. Can be `tcp`, `udp`, or `tcp_udp`. Defaults to `tcp_udp`.
     * 
     */
    @Export(name="protocol", type=String.class, parameters={})
    private Output</* @Nullable */ String> protocol;

    /**
     * @return The protocol for the port forwarding rule. Can be `tcp`, `udp`, or `tcp_udp`. Defaults to `tcp_udp`.
     * 
     */
    public Output<Optional<String>> protocol() {
        return Codegen.optional(this.protocol);
    }
    /**
     * The name of the site to associate the port forwarding rule with.
     * 
     */
    @Export(name="site", type=String.class, parameters={})
    private Output<String> site;

    /**
     * @return The name of the site to associate the port forwarding rule with.
     * 
     */
    public Output<String> site() {
        return this.site;
    }
    /**
     * The source IPv4 address (or CIDR) of the port forwarding rule. For all traffic, specify `any`. Defaults to `any`.
     * 
     */
    @Export(name="srcIp", type=String.class, parameters={})
    private Output</* @Nullable */ String> srcIp;

    /**
     * @return The source IPv4 address (or CIDR) of the port forwarding rule. For all traffic, specify `any`. Defaults to `any`.
     * 
     */
    public Output<Optional<String>> srcIp() {
        return Codegen.optional(this.srcIp);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Forward(String name) {
        this(name, ForwardArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Forward(String name, @Nullable ForwardArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Forward(String name, @Nullable ForwardArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("unifi:port/forward:Forward", name, args == null ? ForwardArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Forward(String name, Output<String> id, @Nullable ForwardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("unifi:port/forward:Forward", name, state, makeResourceOptions(options, id));
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
    public static Forward get(String name, Output<String> id, @Nullable ForwardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Forward(name, id, state, options);
    }
}
