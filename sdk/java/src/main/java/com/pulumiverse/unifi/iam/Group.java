// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.unifi.iam;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumiverse.unifi.Utilities;
import com.pulumiverse.unifi.iam.GroupArgs;
import com.pulumiverse.unifi.iam.inputs.GroupState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `unifi.iam.Group` manages a user group (called &#34;client group&#34; in the UI), which can be used to limit bandwidth for groups of users.
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
 * import com.pulumi.unifi.iam.Group;
 * import com.pulumi.unifi.iam.GroupArgs;
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
 *         var wifi = new Group(&#34;wifi&#34;, GroupArgs.builder()
 *             .qosRateMaxDown(2000)
 *             .qosRateMaxUp(10)
 *             .build());
 * 
 *         // 10kbps
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * import using the ID
 * 
 * ```sh
 * $ pulumi import unifi:iam/group:Group wifi 5fe6261995fe130013456a36
 * ```
 * 
 */
@ResourceType(type="unifi:iam/group:Group")
public class Group extends com.pulumi.resources.CustomResource {
    /**
     * The name of the user group.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the user group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The QOS maximum download rate. Defaults to `-1`.
     * 
     */
    @Export(name="qosRateMaxDown", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> qosRateMaxDown;

    /**
     * @return The QOS maximum download rate. Defaults to `-1`.
     * 
     */
    public Output<Optional<Integer>> qosRateMaxDown() {
        return Codegen.optional(this.qosRateMaxDown);
    }
    /**
     * The QOS maximum upload rate. Defaults to `-1`.
     * 
     */
    @Export(name="qosRateMaxUp", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> qosRateMaxUp;

    /**
     * @return The QOS maximum upload rate. Defaults to `-1`.
     * 
     */
    public Output<Optional<Integer>> qosRateMaxUp() {
        return Codegen.optional(this.qosRateMaxUp);
    }
    /**
     * The name of the site to associate the user group with.
     * 
     */
    @Export(name="site", refs={String.class}, tree="[0]")
    private Output<String> site;

    /**
     * @return The name of the site to associate the user group with.
     * 
     */
    public Output<String> site() {
        return this.site;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Group(String name) {
        this(name, GroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Group(String name, @Nullable GroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Group(String name, @Nullable GroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("unifi:iam/group:Group", name, args == null ? GroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Group(String name, Output<String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("unifi:iam/group:Group", name, state, makeResourceOptions(options, id));
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
    public static Group get(String name, Output<String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Group(name, id, state, options);
    }
}
