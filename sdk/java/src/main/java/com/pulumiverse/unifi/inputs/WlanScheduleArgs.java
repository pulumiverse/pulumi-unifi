// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.unifi.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WlanScheduleArgs extends com.pulumi.resources.ResourceArgs {

    public static final WlanScheduleArgs Empty = new WlanScheduleArgs();

    /**
     * Day of week for the block. Valid values are `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.
     * 
     */
    @Import(name="dayOfWeek", required=true)
    private Output<String> dayOfWeek;

    /**
     * @return Day of week for the block. Valid values are `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.
     * 
     */
    public Output<String> dayOfWeek() {
        return this.dayOfWeek;
    }

    /**
     * Length of the block in minutes.
     * 
     */
    @Import(name="duration", required=true)
    private Output<Integer> duration;

    /**
     * @return Length of the block in minutes.
     * 
     */
    public Output<Integer> duration() {
        return this.duration;
    }

    /**
     * Name of the block.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the block.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Start hour for the block (0-23).
     * 
     */
    @Import(name="startHour", required=true)
    private Output<Integer> startHour;

    /**
     * @return Start hour for the block (0-23).
     * 
     */
    public Output<Integer> startHour() {
        return this.startHour;
    }

    /**
     * Start minute for the block (0-59). Defaults to `0`.
     * 
     */
    @Import(name="startMinute")
    private @Nullable Output<Integer> startMinute;

    /**
     * @return Start minute for the block (0-59). Defaults to `0`.
     * 
     */
    public Optional<Output<Integer>> startMinute() {
        return Optional.ofNullable(this.startMinute);
    }

    private WlanScheduleArgs() {}

    private WlanScheduleArgs(WlanScheduleArgs $) {
        this.dayOfWeek = $.dayOfWeek;
        this.duration = $.duration;
        this.name = $.name;
        this.startHour = $.startHour;
        this.startMinute = $.startMinute;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WlanScheduleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WlanScheduleArgs $;

        public Builder() {
            $ = new WlanScheduleArgs();
        }

        public Builder(WlanScheduleArgs defaults) {
            $ = new WlanScheduleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dayOfWeek Day of week for the block. Valid values are `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.
         * 
         * @return builder
         * 
         */
        public Builder dayOfWeek(Output<String> dayOfWeek) {
            $.dayOfWeek = dayOfWeek;
            return this;
        }

        /**
         * @param dayOfWeek Day of week for the block. Valid values are `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.
         * 
         * @return builder
         * 
         */
        public Builder dayOfWeek(String dayOfWeek) {
            return dayOfWeek(Output.of(dayOfWeek));
        }

        /**
         * @param duration Length of the block in minutes.
         * 
         * @return builder
         * 
         */
        public Builder duration(Output<Integer> duration) {
            $.duration = duration;
            return this;
        }

        /**
         * @param duration Length of the block in minutes.
         * 
         * @return builder
         * 
         */
        public Builder duration(Integer duration) {
            return duration(Output.of(duration));
        }

        /**
         * @param name Name of the block.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the block.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param startHour Start hour for the block (0-23).
         * 
         * @return builder
         * 
         */
        public Builder startHour(Output<Integer> startHour) {
            $.startHour = startHour;
            return this;
        }

        /**
         * @param startHour Start hour for the block (0-23).
         * 
         * @return builder
         * 
         */
        public Builder startHour(Integer startHour) {
            return startHour(Output.of(startHour));
        }

        /**
         * @param startMinute Start minute for the block (0-59). Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder startMinute(@Nullable Output<Integer> startMinute) {
            $.startMinute = startMinute;
            return this;
        }

        /**
         * @param startMinute Start minute for the block (0-59). Defaults to `0`.
         * 
         * @return builder
         * 
         */
        public Builder startMinute(Integer startMinute) {
            return startMinute(Output.of(startMinute));
        }

        public WlanScheduleArgs build() {
            $.dayOfWeek = Objects.requireNonNull($.dayOfWeek, "expected parameter 'dayOfWeek' to be non-null");
            $.duration = Objects.requireNonNull($.duration, "expected parameter 'duration' to be non-null");
            $.startHour = Objects.requireNonNull($.startHour, "expected parameter 'startHour' to be non-null");
            return $;
        }
    }

}
