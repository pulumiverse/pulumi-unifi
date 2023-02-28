// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.unifi.port;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ForwardArgs extends com.pulumi.resources.ResourceArgs {

    public static final ForwardArgs Empty = new ForwardArgs();

    /**
     * The destination port for the forwarding.
     * 
     */
    @Import(name="dstPort")
    private @Nullable Output<String> dstPort;

    /**
     * @return The destination port for the forwarding.
     * 
     */
    public Optional<Output<String>> dstPort() {
        return Optional.ofNullable(this.dstPort);
    }

    /**
     * Specifies whether the port forwarding rule is enabled or not. Defaults to `true`. This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
     * 
     * @deprecated
     * This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
     * 
     */
    @Deprecated /* This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration. */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Specifies whether the port forwarding rule is enabled or not. Defaults to `true`. This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
     * 
     * @deprecated
     * This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
     * 
     */
    @Deprecated /* This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration. */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * The IPv4 address to forward traffic to.
     * 
     */
    @Import(name="fwdIp")
    private @Nullable Output<String> fwdIp;

    /**
     * @return The IPv4 address to forward traffic to.
     * 
     */
    public Optional<Output<String>> fwdIp() {
        return Optional.ofNullable(this.fwdIp);
    }

    /**
     * The port to forward traffic to.
     * 
     */
    @Import(name="fwdPort")
    private @Nullable Output<String> fwdPort;

    /**
     * @return The port to forward traffic to.
     * 
     */
    public Optional<Output<String>> fwdPort() {
        return Optional.ofNullable(this.fwdPort);
    }

    /**
     * Specifies whether to log forwarded traffic or not. Defaults to `false`.
     * 
     */
    @Import(name="log")
    private @Nullable Output<Boolean> log;

    /**
     * @return Specifies whether to log forwarded traffic or not. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> log() {
        return Optional.ofNullable(this.log);
    }

    /**
     * The name of the port forwarding rule.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the port forwarding rule.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The port forwarding interface. Can be `wan`, `wan2`, or `both`.
     * 
     */
    @Import(name="portForwardInterface")
    private @Nullable Output<String> portForwardInterface;

    /**
     * @return The port forwarding interface. Can be `wan`, `wan2`, or `both`.
     * 
     */
    public Optional<Output<String>> portForwardInterface() {
        return Optional.ofNullable(this.portForwardInterface);
    }

    /**
     * The protocol for the port forwarding rule. Can be `tcp`, `udp`, or `tcp_udp`. Defaults to `tcp_udp`.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return The protocol for the port forwarding rule. Can be `tcp`, `udp`, or `tcp_udp`. Defaults to `tcp_udp`.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * The name of the site to associate the port forwarding rule with.
     * 
     */
    @Import(name="site")
    private @Nullable Output<String> site;

    /**
     * @return The name of the site to associate the port forwarding rule with.
     * 
     */
    public Optional<Output<String>> site() {
        return Optional.ofNullable(this.site);
    }

    /**
     * The source IPv4 address (or CIDR) of the port forwarding rule. For all traffic, specify `any`. Defaults to `any`.
     * 
     */
    @Import(name="srcIp")
    private @Nullable Output<String> srcIp;

    /**
     * @return The source IPv4 address (or CIDR) of the port forwarding rule. For all traffic, specify `any`. Defaults to `any`.
     * 
     */
    public Optional<Output<String>> srcIp() {
        return Optional.ofNullable(this.srcIp);
    }

    private ForwardArgs() {}

    private ForwardArgs(ForwardArgs $) {
        this.dstPort = $.dstPort;
        this.enabled = $.enabled;
        this.fwdIp = $.fwdIp;
        this.fwdPort = $.fwdPort;
        this.log = $.log;
        this.name = $.name;
        this.portForwardInterface = $.portForwardInterface;
        this.protocol = $.protocol;
        this.site = $.site;
        this.srcIp = $.srcIp;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ForwardArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ForwardArgs $;

        public Builder() {
            $ = new ForwardArgs();
        }

        public Builder(ForwardArgs defaults) {
            $ = new ForwardArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dstPort The destination port for the forwarding.
         * 
         * @return builder
         * 
         */
        public Builder dstPort(@Nullable Output<String> dstPort) {
            $.dstPort = dstPort;
            return this;
        }

        /**
         * @param dstPort The destination port for the forwarding.
         * 
         * @return builder
         * 
         */
        public Builder dstPort(String dstPort) {
            return dstPort(Output.of(dstPort));
        }

        /**
         * @param enabled Specifies whether the port forwarding rule is enabled or not. Defaults to `true`. This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
         * 
         * @return builder
         * 
         * @deprecated
         * This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
         * 
         */
        @Deprecated /* This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration. */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Specifies whether the port forwarding rule is enabled or not. Defaults to `true`. This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
         * 
         * @return builder
         * 
         * @deprecated
         * This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration.
         * 
         */
        @Deprecated /* This will attribute will be removed in a future release. Instead of disabling a port forwarding rule you can remove it from your configuration. */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param fwdIp The IPv4 address to forward traffic to.
         * 
         * @return builder
         * 
         */
        public Builder fwdIp(@Nullable Output<String> fwdIp) {
            $.fwdIp = fwdIp;
            return this;
        }

        /**
         * @param fwdIp The IPv4 address to forward traffic to.
         * 
         * @return builder
         * 
         */
        public Builder fwdIp(String fwdIp) {
            return fwdIp(Output.of(fwdIp));
        }

        /**
         * @param fwdPort The port to forward traffic to.
         * 
         * @return builder
         * 
         */
        public Builder fwdPort(@Nullable Output<String> fwdPort) {
            $.fwdPort = fwdPort;
            return this;
        }

        /**
         * @param fwdPort The port to forward traffic to.
         * 
         * @return builder
         * 
         */
        public Builder fwdPort(String fwdPort) {
            return fwdPort(Output.of(fwdPort));
        }

        /**
         * @param log Specifies whether to log forwarded traffic or not. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder log(@Nullable Output<Boolean> log) {
            $.log = log;
            return this;
        }

        /**
         * @param log Specifies whether to log forwarded traffic or not. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder log(Boolean log) {
            return log(Output.of(log));
        }

        /**
         * @param name The name of the port forwarding rule.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the port forwarding rule.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param portForwardInterface The port forwarding interface. Can be `wan`, `wan2`, or `both`.
         * 
         * @return builder
         * 
         */
        public Builder portForwardInterface(@Nullable Output<String> portForwardInterface) {
            $.portForwardInterface = portForwardInterface;
            return this;
        }

        /**
         * @param portForwardInterface The port forwarding interface. Can be `wan`, `wan2`, or `both`.
         * 
         * @return builder
         * 
         */
        public Builder portForwardInterface(String portForwardInterface) {
            return portForwardInterface(Output.of(portForwardInterface));
        }

        /**
         * @param protocol The protocol for the port forwarding rule. Can be `tcp`, `udp`, or `tcp_udp`. Defaults to `tcp_udp`.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The protocol for the port forwarding rule. Can be `tcp`, `udp`, or `tcp_udp`. Defaults to `tcp_udp`.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param site The name of the site to associate the port forwarding rule with.
         * 
         * @return builder
         * 
         */
        public Builder site(@Nullable Output<String> site) {
            $.site = site;
            return this;
        }

        /**
         * @param site The name of the site to associate the port forwarding rule with.
         * 
         * @return builder
         * 
         */
        public Builder site(String site) {
            return site(Output.of(site));
        }

        /**
         * @param srcIp The source IPv4 address (or CIDR) of the port forwarding rule. For all traffic, specify `any`. Defaults to `any`.
         * 
         * @return builder
         * 
         */
        public Builder srcIp(@Nullable Output<String> srcIp) {
            $.srcIp = srcIp;
            return this;
        }

        /**
         * @param srcIp The source IPv4 address (or CIDR) of the port forwarding rule. For all traffic, specify `any`. Defaults to `any`.
         * 
         * @return builder
         * 
         */
        public Builder srcIp(String srcIp) {
            return srcIp(Output.of(srcIp));
        }

        public ForwardArgs build() {
            return $;
        }
    }

}
