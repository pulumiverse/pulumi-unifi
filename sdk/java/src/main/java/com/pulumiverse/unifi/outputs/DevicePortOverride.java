// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.unifi.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DevicePortOverride {
    /**
     * @return Number of ports in the aggregate.
     * 
     */
    private @Nullable Integer aggregateNumPorts;
    /**
     * @return Human-readable name of the port.
     * 
     */
    private @Nullable String name;
    /**
     * @return Switch port number.
     * 
     */
    private Integer number;
    /**
     * @return Operating mode of the port, valid values are `switch`, `mirror`, and `aggregate`. Defaults to `switch`.
     * 
     */
    private @Nullable String opMode;
    /**
     * @return ID of the Port Profile used on this port.
     * 
     */
    private @Nullable String portProfileId;

    private DevicePortOverride() {}
    /**
     * @return Number of ports in the aggregate.
     * 
     */
    public Optional<Integer> aggregateNumPorts() {
        return Optional.ofNullable(this.aggregateNumPorts);
    }
    /**
     * @return Human-readable name of the port.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return Switch port number.
     * 
     */
    public Integer number() {
        return this.number;
    }
    /**
     * @return Operating mode of the port, valid values are `switch`, `mirror`, and `aggregate`. Defaults to `switch`.
     * 
     */
    public Optional<String> opMode() {
        return Optional.ofNullable(this.opMode);
    }
    /**
     * @return ID of the Port Profile used on this port.
     * 
     */
    public Optional<String> portProfileId() {
        return Optional.ofNullable(this.portProfileId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DevicePortOverride defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer aggregateNumPorts;
        private @Nullable String name;
        private Integer number;
        private @Nullable String opMode;
        private @Nullable String portProfileId;
        public Builder() {}
        public Builder(DevicePortOverride defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aggregateNumPorts = defaults.aggregateNumPorts;
    	      this.name = defaults.name;
    	      this.number = defaults.number;
    	      this.opMode = defaults.opMode;
    	      this.portProfileId = defaults.portProfileId;
        }

        @CustomType.Setter
        public Builder aggregateNumPorts(@Nullable Integer aggregateNumPorts) {
            this.aggregateNumPorts = aggregateNumPorts;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder number(Integer number) {
            this.number = Objects.requireNonNull(number);
            return this;
        }
        @CustomType.Setter
        public Builder opMode(@Nullable String opMode) {
            this.opMode = opMode;
            return this;
        }
        @CustomType.Setter
        public Builder portProfileId(@Nullable String portProfileId) {
            this.portProfileId = portProfileId;
            return this;
        }
        public DevicePortOverride build() {
            final var o = new DevicePortOverride();
            o.aggregateNumPorts = aggregateNumPorts;
            o.name = name;
            o.number = number;
            o.opMode = opMode;
            o.portProfileId = portProfileId;
            return o;
        }
    }
}
