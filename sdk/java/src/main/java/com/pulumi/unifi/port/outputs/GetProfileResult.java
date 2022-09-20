// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.unifi.port.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetProfileResult {
    /**
     * @return The ID of this port profile.
     * 
     */
    private String id;
    /**
     * @return The name of the port profile to look up. Defaults to `All`.
     * 
     */
    private @Nullable String name;
    /**
     * @return The name of the site the port profile is associated with.
     * 
     */
    private String site;

    private GetProfileResult() {}
    /**
     * @return The ID of this port profile.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the port profile to look up. Defaults to `All`.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The name of the site the port profile is associated with.
     * 
     */
    public String site() {
        return this.site;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProfileResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable String name;
        private String site;
        public Builder() {}
        public Builder(GetProfileResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.site = defaults.site;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder site(String site) {
            this.site = Objects.requireNonNull(site);
            return this;
        }
        public GetProfileResult build() {
            final var o = new GetProfileResult();
            o.id = id;
            o.name = name;
            o.site = site;
            return o;
        }
    }
}
