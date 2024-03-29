// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.unifi.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNetworkArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNetworkArgs Empty = new GetNetworkArgs();

    /**
     * The ID of the network.
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return The ID of the network.
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * The name of the network.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the network.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The name of the site to associate the network with.
     * 
     */
    @Import(name="site")
    private @Nullable Output<String> site;

    /**
     * @return The name of the site to associate the network with.
     * 
     */
    public Optional<Output<String>> site() {
        return Optional.ofNullable(this.site);
    }

    private GetNetworkArgs() {}

    private GetNetworkArgs(GetNetworkArgs $) {
        this.id = $.id;
        this.name = $.name;
        this.site = $.site;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNetworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNetworkArgs $;

        public Builder() {
            $ = new GetNetworkArgs();
        }

        public Builder(GetNetworkArgs defaults) {
            $ = new GetNetworkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id The ID of the network.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The ID of the network.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param name The name of the network.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the network.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param site The name of the site to associate the network with.
         * 
         * @return builder
         * 
         */
        public Builder site(@Nullable Output<String> site) {
            $.site = site;
            return this;
        }

        /**
         * @param site The name of the site to associate the network with.
         * 
         * @return builder
         * 
         */
        public Builder site(String site) {
            return site(Output.of(site));
        }

        public GetNetworkArgs build() {
            return $;
        }
    }

}
