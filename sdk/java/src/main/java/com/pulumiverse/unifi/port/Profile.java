// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.unifi.port;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumiverse.unifi.Utilities;
import com.pulumiverse.unifi.port.ProfileArgs;
import com.pulumiverse.unifi.port.inputs.ProfileState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `unifi.port.Profile` manages a port profile for use on network switches.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.unifi.Network;
 * import com.pulumi.unifi.NetworkArgs;
 * import com.pulumi.unifi.port.Profile;
 * import com.pulumi.unifi.port.ProfileArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var vlanId = config.get(&#34;vlanId&#34;).orElse(10);
 *         var vlan = new Network(&#34;vlan&#34;, NetworkArgs.builder()
 *             .purpose(&#34;corporate&#34;)
 *             .subnet(&#34;10.0.0.1/24&#34;)
 *             .vlanId(vlanId)
 *             .dhcpStart(&#34;10.0.0.6&#34;)
 *             .dhcpStop(&#34;10.0.0.254&#34;)
 *             .dhcpEnabled(true)
 *             .build());
 * 
 *         var poeDisabled = new Profile(&#34;poeDisabled&#34;, ProfileArgs.builder()
 *             .nativeNetworkconfId(vlan.id())
 *             .poeMode(&#34;off&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="unifi:port/profile:Profile")
public class Profile extends com.pulumi.resources.CustomResource {
    /**
     * Enable link auto negotiation for the port profile. When set to `true` this overrides `speed`. Defaults to `true`.
     * 
     */
    @Export(name="autoneg", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoneg;

    /**
     * @return Enable link auto negotiation for the port profile. When set to `true` this overrides `speed`. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> autoneg() {
        return Codegen.optional(this.autoneg);
    }
    /**
     * The type of 802.1X control to use. Can be `auto`, `force_authorized`, `force_unauthorized`, `mac_based` or `multi_host`. Defaults to `force_authorized`.
     * 
     */
    @Export(name="dot1xCtrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dot1xCtrl;

    /**
     * @return The type of 802.1X control to use. Can be `auto`, `force_authorized`, `force_unauthorized`, `mac_based` or `multi_host`. Defaults to `force_authorized`.
     * 
     */
    public Output<Optional<String>> dot1xCtrl() {
        return Codegen.optional(this.dot1xCtrl);
    }
    /**
     * The timeout, in seconds, to use when using the MAC Based 802.1X control. Can be between 0 and 65535 Defaults to `300`.
     * 
     */
    @Export(name="dot1xIdleTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> dot1xIdleTimeout;

    /**
     * @return The timeout, in seconds, to use when using the MAC Based 802.1X control. Can be between 0 and 65535 Defaults to `300`.
     * 
     */
    public Output<Optional<Integer>> dot1xIdleTimeout() {
        return Codegen.optional(this.dot1xIdleTimeout);
    }
    /**
     * The egress rate limit, in kpbs, for the port profile. Can be between `64` and `9999999`.
     * 
     */
    @Export(name="egressRateLimitKbps", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> egressRateLimitKbps;

    /**
     * @return The egress rate limit, in kpbs, for the port profile. Can be between `64` and `9999999`.
     * 
     */
    public Output<Optional<Integer>> egressRateLimitKbps() {
        return Codegen.optional(this.egressRateLimitKbps);
    }
    /**
     * Enable egress rate limiting for the port profile. Defaults to `false`.
     * 
     */
    @Export(name="egressRateLimitKbpsEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> egressRateLimitKbpsEnabled;

    /**
     * @return Enable egress rate limiting for the port profile. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> egressRateLimitKbpsEnabled() {
        return Codegen.optional(this.egressRateLimitKbpsEnabled);
    }
    /**
     * The type forwarding to use for the port profile. Can be `all`, `native`, `customize` or `disabled`. Defaults to `native`.
     * 
     */
    @Export(name="forward", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> forward;

    /**
     * @return The type forwarding to use for the port profile. Can be `all`, `native`, `customize` or `disabled`. Defaults to `native`.
     * 
     */
    public Output<Optional<String>> forward() {
        return Codegen.optional(this.forward);
    }
    /**
     * Enable full duplex for the port profile. Defaults to `false`.
     * 
     */
    @Export(name="fullDuplex", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> fullDuplex;

    /**
     * @return Enable full duplex for the port profile. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> fullDuplex() {
        return Codegen.optional(this.fullDuplex);
    }
    /**
     * Enable port isolation for the port profile. Defaults to `false`.
     * 
     */
    @Export(name="isolation", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isolation;

    /**
     * @return Enable port isolation for the port profile. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> isolation() {
        return Codegen.optional(this.isolation);
    }
    /**
     * Enable LLDP-MED for the port profile. Defaults to `true`.
     * 
     */
    @Export(name="lldpmedEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> lldpmedEnabled;

    /**
     * @return Enable LLDP-MED for the port profile. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> lldpmedEnabled() {
        return Codegen.optional(this.lldpmedEnabled);
    }
    /**
     * Enable LLDP-MED topology change notifications for the port profile.
     * 
     */
    @Export(name="lldpmedNotifyEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> lldpmedNotifyEnabled;

    /**
     * @return Enable LLDP-MED topology change notifications for the port profile.
     * 
     */
    public Output<Optional<Boolean>> lldpmedNotifyEnabled() {
        return Codegen.optional(this.lldpmedNotifyEnabled);
    }
    /**
     * The name of the port profile.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the port profile.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of network to use as the main network on the port profile.
     * 
     */
    @Export(name="nativeNetworkconfId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> nativeNetworkconfId;

    /**
     * @return The ID of network to use as the main network on the port profile.
     * 
     */
    public Output<Optional<String>> nativeNetworkconfId() {
        return Codegen.optional(this.nativeNetworkconfId);
    }
    /**
     * The operation mode for the port profile. Can only be `switch` Defaults to `switch`.
     * 
     */
    @Export(name="opMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> opMode;

    /**
     * @return The operation mode for the port profile. Can only be `switch` Defaults to `switch`.
     * 
     */
    public Output<Optional<String>> opMode() {
        return Codegen.optional(this.opMode);
    }
    /**
     * The POE mode for the port profile. Can be one of `auto`, `passv24`, `passthrough` or `off`.
     * 
     */
    @Export(name="poeMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> poeMode;

    /**
     * @return The POE mode for the port profile. Can be one of `auto`, `passv24`, `passthrough` or `off`.
     * 
     */
    public Output<Optional<String>> poeMode() {
        return Codegen.optional(this.poeMode);
    }
    /**
     * Enable port security for the port profile. Defaults to `false`.
     * 
     */
    @Export(name="portSecurityEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> portSecurityEnabled;

    /**
     * @return Enable port security for the port profile. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> portSecurityEnabled() {
        return Codegen.optional(this.portSecurityEnabled);
    }
    /**
     * The MAC addresses associated with the port security for the port profile.
     * 
     */
    @Export(name="portSecurityMacAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> portSecurityMacAddresses;

    /**
     * @return The MAC addresses associated with the port security for the port profile.
     * 
     */
    public Output<Optional<List<String>>> portSecurityMacAddresses() {
        return Codegen.optional(this.portSecurityMacAddresses);
    }
    /**
     * The priority queue 1 level for the port profile. Can be between 0 and 100.
     * 
     */
    @Export(name="priorityQueue1Level", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> priorityQueue1Level;

    /**
     * @return The priority queue 1 level for the port profile. Can be between 0 and 100.
     * 
     */
    public Output<Optional<Integer>> priorityQueue1Level() {
        return Codegen.optional(this.priorityQueue1Level);
    }
    /**
     * The priority queue 2 level for the port profile. Can be between 0 and 100.
     * 
     */
    @Export(name="priorityQueue2Level", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> priorityQueue2Level;

    /**
     * @return The priority queue 2 level for the port profile. Can be between 0 and 100.
     * 
     */
    public Output<Optional<Integer>> priorityQueue2Level() {
        return Codegen.optional(this.priorityQueue2Level);
    }
    /**
     * The priority queue 3 level for the port profile. Can be between 0 and 100.
     * 
     */
    @Export(name="priorityQueue3Level", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> priorityQueue3Level;

    /**
     * @return The priority queue 3 level for the port profile. Can be between 0 and 100.
     * 
     */
    public Output<Optional<Integer>> priorityQueue3Level() {
        return Codegen.optional(this.priorityQueue3Level);
    }
    /**
     * The priority queue 4 level for the port profile. Can be between 0 and 100.
     * 
     */
    @Export(name="priorityQueue4Level", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> priorityQueue4Level;

    /**
     * @return The priority queue 4 level for the port profile. Can be between 0 and 100.
     * 
     */
    public Output<Optional<Integer>> priorityQueue4Level() {
        return Codegen.optional(this.priorityQueue4Level);
    }
    /**
     * The name of the site to associate the port profile with.
     * 
     */
    @Export(name="site", refs={String.class}, tree="[0]")
    private Output<String> site;

    /**
     * @return The name of the site to associate the port profile with.
     * 
     */
    public Output<String> site() {
        return this.site;
    }
    /**
     * The link speed to set for the port profile. Can be one of `10`, `100`, `1000`, `2500`, `5000`, `10000`, `20000`, `25000`, `40000`, `50000` or `100000`
     * 
     */
    @Export(name="speed", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> speed;

    /**
     * @return The link speed to set for the port profile. Can be one of `10`, `100`, `1000`, `2500`, `5000`, `10000`, `20000`, `25000`, `40000`, `50000` or `100000`
     * 
     */
    public Output<Optional<Integer>> speed() {
        return Codegen.optional(this.speed);
    }
    /**
     * Enable broadcast Storm Control for the port profile. Defaults to `false`.
     * 
     */
    @Export(name="stormctrlBcastEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stormctrlBcastEnabled;

    /**
     * @return Enable broadcast Storm Control for the port profile. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> stormctrlBcastEnabled() {
        return Codegen.optional(this.stormctrlBcastEnabled);
    }
    /**
     * The broadcast Storm Control level for the port profile. Can be between 0 and 100.
     * 
     */
    @Export(name="stormctrlBcastLevel", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> stormctrlBcastLevel;

    /**
     * @return The broadcast Storm Control level for the port profile. Can be between 0 and 100.
     * 
     */
    public Output<Optional<Integer>> stormctrlBcastLevel() {
        return Codegen.optional(this.stormctrlBcastLevel);
    }
    /**
     * The broadcast Storm Control rate for the port profile. Can be between 0 and 14880000.
     * 
     */
    @Export(name="stormctrlBcastRate", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> stormctrlBcastRate;

    /**
     * @return The broadcast Storm Control rate for the port profile. Can be between 0 and 14880000.
     * 
     */
    public Output<Optional<Integer>> stormctrlBcastRate() {
        return Codegen.optional(this.stormctrlBcastRate);
    }
    /**
     * Enable multicast Storm Control for the port profile. Defaults to `false`.
     * 
     */
    @Export(name="stormctrlMcastEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stormctrlMcastEnabled;

    /**
     * @return Enable multicast Storm Control for the port profile. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> stormctrlMcastEnabled() {
        return Codegen.optional(this.stormctrlMcastEnabled);
    }
    /**
     * The multicast Storm Control level for the port profile. Can be between 0 and 100.
     * 
     */
    @Export(name="stormctrlMcastLevel", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> stormctrlMcastLevel;

    /**
     * @return The multicast Storm Control level for the port profile. Can be between 0 and 100.
     * 
     */
    public Output<Optional<Integer>> stormctrlMcastLevel() {
        return Codegen.optional(this.stormctrlMcastLevel);
    }
    /**
     * The multicast Storm Control rate for the port profile. Can be between 0 and 14880000.
     * 
     */
    @Export(name="stormctrlMcastRate", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> stormctrlMcastRate;

    /**
     * @return The multicast Storm Control rate for the port profile. Can be between 0 and 14880000.
     * 
     */
    public Output<Optional<Integer>> stormctrlMcastRate() {
        return Codegen.optional(this.stormctrlMcastRate);
    }
    /**
     * The type of Storm Control to use for the port profile. Can be one of `level` or `rate`.
     * 
     */
    @Export(name="stormctrlType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stormctrlType;

    /**
     * @return The type of Storm Control to use for the port profile. Can be one of `level` or `rate`.
     * 
     */
    public Output<Optional<String>> stormctrlType() {
        return Codegen.optional(this.stormctrlType);
    }
    /**
     * Enable unknown unicast Storm Control for the port profile. Defaults to `false`.
     * 
     */
    @Export(name="stormctrlUcastEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stormctrlUcastEnabled;

    /**
     * @return Enable unknown unicast Storm Control for the port profile. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> stormctrlUcastEnabled() {
        return Codegen.optional(this.stormctrlUcastEnabled);
    }
    /**
     * The unknown unicast Storm Control level for the port profile. Can be between 0 and 100.
     * 
     */
    @Export(name="stormctrlUcastLevel", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> stormctrlUcastLevel;

    /**
     * @return The unknown unicast Storm Control level for the port profile. Can be between 0 and 100.
     * 
     */
    public Output<Optional<Integer>> stormctrlUcastLevel() {
        return Codegen.optional(this.stormctrlUcastLevel);
    }
    /**
     * The unknown unicast Storm Control rate for the port profile. Can be between 0 and 14880000.
     * 
     */
    @Export(name="stormctrlUcastRate", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> stormctrlUcastRate;

    /**
     * @return The unknown unicast Storm Control rate for the port profile. Can be between 0 and 14880000.
     * 
     */
    public Output<Optional<Integer>> stormctrlUcastRate() {
        return Codegen.optional(this.stormctrlUcastRate);
    }
    /**
     * Enable spanning tree protocol on the port profile. Defaults to `true`.
     * 
     */
    @Export(name="stpPortMode", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stpPortMode;

    /**
     * @return Enable spanning tree protocol on the port profile. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> stpPortMode() {
        return Codegen.optional(this.stpPortMode);
    }
    /**
     * The IDs of networks to tag traffic with for the port profile.
     * 
     */
    @Export(name="taggedNetworkconfIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> taggedNetworkconfIds;

    /**
     * @return The IDs of networks to tag traffic with for the port profile.
     * 
     */
    public Output<Optional<List<String>>> taggedNetworkconfIds() {
        return Codegen.optional(this.taggedNetworkconfIds);
    }
    /**
     * The ID of network to use as the voice network on the port profile.
     * 
     */
    @Export(name="voiceNetworkconfId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> voiceNetworkconfId;

    /**
     * @return The ID of network to use as the voice network on the port profile.
     * 
     */
    public Output<Optional<String>> voiceNetworkconfId() {
        return Codegen.optional(this.voiceNetworkconfId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Profile(String name) {
        this(name, ProfileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Profile(String name, @Nullable ProfileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Profile(String name, @Nullable ProfileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("unifi:port/profile:Profile", name, args == null ? ProfileArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Profile(String name, Output<String> id, @Nullable ProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("unifi:port/profile:Profile", name, state, makeResourceOptions(options, id));
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
    public static Profile get(String name, Output<String> id, @Nullable ProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Profile(name, id, state, options);
    }
}
