// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.unifi.iam.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetUserArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserArgs Empty = new GetUserArgs();

    /**
     * The MAC address of the user.
     * 
     */
    @Import(name="mac", required=true)
    private Output<String> mac;

    /**
     * @return The MAC address of the user.
     * 
     */
    public Output<String> mac() {
        return this.mac;
    }

    /**
     * The name of the site the user is associated with.
     * 
     */
    @Import(name="site")
    private @Nullable Output<String> site;

    /**
     * @return The name of the site the user is associated with.
     * 
     */
    public Optional<Output<String>> site() {
        return Optional.ofNullable(this.site);
    }

    private GetUserArgs() {}

    private GetUserArgs(GetUserArgs $) {
        this.mac = $.mac;
        this.site = $.site;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserArgs $;

        public Builder() {
            $ = new GetUserArgs();
        }

        public Builder(GetUserArgs defaults) {
            $ = new GetUserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param mac The MAC address of the user.
         * 
         * @return builder
         * 
         */
        public Builder mac(Output<String> mac) {
            $.mac = mac;
            return this;
        }

        /**
         * @param mac The MAC address of the user.
         * 
         * @return builder
         * 
         */
        public Builder mac(String mac) {
            return mac(Output.of(mac));
        }

        /**
         * @param site The name of the site the user is associated with.
         * 
         * @return builder
         * 
         */
        public Builder site(@Nullable Output<String> site) {
            $.site = site;
            return this;
        }

        /**
         * @param site The name of the site the user is associated with.
         * 
         * @return builder
         * 
         */
        public Builder site(String site) {
            return site(Output.of(site));
        }

        public GetUserArgs build() {
            $.mac = Objects.requireNonNull($.mac, "expected parameter 'mac' to be non-null");
            return $;
        }
    }

}