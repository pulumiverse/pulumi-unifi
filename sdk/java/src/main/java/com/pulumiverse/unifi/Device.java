// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.unifi;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumiverse.unifi.DeviceArgs;
import com.pulumiverse.unifi.Utilities;
import com.pulumiverse.unifi.inputs.DeviceState;
import com.pulumiverse.unifi.outputs.DevicePortOverride;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.unifi.port.PortFunctions;
 * import com.pulumi.unifi.port.inputs.GetProfileArgs;
 * import com.pulumi.unifi.port.Profile;
 * import com.pulumi.unifi.port.ProfileArgs;
 * import com.pulumi.unifi.Device;
 * import com.pulumi.unifi.DeviceArgs;
 * import com.pulumi.unifi.inputs.DevicePortOverrideArgs;
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
 *         final var disabled = PortFunctions.getProfile(GetProfileArgs.builder()
 *             .name(&#34;Disabled&#34;)
 *             .build());
 * 
 *         var poe = new Profile(&#34;poe&#34;, ProfileArgs.builder()        
 *             .forward(&#34;customize&#34;)
 *             .nativeNetworkconfId(var_.native_network_id())
 *             .taggedNetworkconfIds(var_.some_vlan_network_id())
 *             .poeMode(&#34;auto&#34;)
 *             .build());
 * 
 *         var us24Poe = new Device(&#34;us24Poe&#34;, DeviceArgs.builder()        
 *             .mac(&#34;01:23:45:67:89:AB&#34;)
 *             .portOverrides(            
 *                 DevicePortOverrideArgs.builder()
 *                     .number(1)
 *                     .name(&#34;port w/ poe&#34;)
 *                     .portProfileId(poe.id())
 *                     .build(),
 *                 DevicePortOverrideArgs.builder()
 *                     .number(2)
 *                     .name(&#34;disabled&#34;)
 *                     .portProfileId(disabled.applyValue(getProfileResult -&gt; getProfileResult.id()))
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="unifi:index/device:Device")
public class Device extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether this resource should tell the controller to adopt the device on create. Defaults to `true`.
     * 
     */
    @Export(name="allowAdoption", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowAdoption;

    /**
     * @return Specifies whether this resource should tell the controller to adopt the device on create. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> allowAdoption() {
        return Codegen.optional(this.allowAdoption);
    }
    /**
     * Specifies whether this device should be disabled.
     * 
     */
    @Export(name="disabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> disabled;

    /**
     * @return Specifies whether this device should be disabled.
     * 
     */
    public Output<Boolean> disabled() {
        return this.disabled;
    }
    /**
     * Specifies whether this resource should tell the controller to forget the device on destroy. Defaults to `true`.
     * 
     */
    @Export(name="forgetOnDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forgetOnDestroy;

    /**
     * @return Specifies whether this resource should tell the controller to forget the device on destroy. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> forgetOnDestroy() {
        return Codegen.optional(this.forgetOnDestroy);
    }
    /**
     * The MAC address of the device. This can be specified so that the provider can take control of a device (since devices are created through adoption).
     * 
     */
    @Export(name="mac", refs={String.class}, tree="[0]")
    private Output<String> mac;

    /**
     * @return The MAC address of the device. This can be specified so that the provider can take control of a device (since devices are created through adoption).
     * 
     */
    public Output<String> mac() {
        return this.mac;
    }
    /**
     * The name of the device.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the device.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Settings overrides for specific switch ports.
     * 
     */
    @Export(name="portOverrides", refs={List.class,DevicePortOverride.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DevicePortOverride>> portOverrides;

    /**
     * @return Settings overrides for specific switch ports.
     * 
     */
    public Output<Optional<List<DevicePortOverride>>> portOverrides() {
        return Codegen.optional(this.portOverrides);
    }
    /**
     * The name of the site to associate the device with.
     * 
     */
    @Export(name="site", refs={String.class}, tree="[0]")
    private Output<String> site;

    /**
     * @return The name of the site to associate the device with.
     * 
     */
    public Output<String> site() {
        return this.site;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Device(String name) {
        this(name, DeviceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Device(String name, @Nullable DeviceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Device(String name, @Nullable DeviceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("unifi:index/device:Device", name, args == null ? DeviceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Device(String name, Output<String> id, @Nullable DeviceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("unifi:index/device:Device", name, state, makeResourceOptions(options, id));
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
    public static Device get(String name, Output<String> id, @Nullable DeviceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Device(name, id, state, options);
    }
}
