// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.unifi;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class SiteArgs extends com.pulumi.resources.ResourceArgs {

    public static final SiteArgs Empty = new SiteArgs();

    /**
     * The description of the site.
     * 
     */
    @Import(name="description", required=true)
    private Output<String> description;

    /**
     * @return The description of the site.
     * 
     */
    public Output<String> description() {
        return this.description;
    }

    private SiteArgs() {}

    private SiteArgs(SiteArgs $) {
        this.description = $.description;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SiteArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SiteArgs $;

        public Builder() {
            $ = new SiteArgs();
        }

        public Builder(SiteArgs defaults) {
            $ = new SiteArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the site.
         * 
         * @return builder
         * 
         */
        public Builder description(Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the site.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        public SiteArgs build() {
            $.description = Objects.requireNonNull($.description, "expected parameter 'description' to be non-null");
            return $;
        }
    }

}
