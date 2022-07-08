// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class Device extends pulumi.CustomResource {
    /**
     * Get an existing Device resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceState, opts?: pulumi.CustomResourceOptions): Device {
        return new Device(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'unifi:index/device:Device';

    /**
     * Returns true if the given object is an instance of Device.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Device {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Device.__pulumiType;
    }

    /**
     * Specifies whether this device should be disabled.
     */
    public /*out*/ readonly disabled!: pulumi.Output<boolean>;
    /**
     * The MAC address of the device. This can be specified so that the provider can take control of a device (since devices are created through adoption).
     */
    public readonly mac!: pulumi.Output<string>;
    /**
     * The name of the device.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Settings overrides for specific switch ports.
     */
    public readonly portOverrides!: pulumi.Output<outputs.DevicePortOverride[] | undefined>;
    /**
     * The name of the site to associate the device with.
     */
    public readonly site!: pulumi.Output<string>;

    /**
     * Create a Device resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DeviceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceArgs | DeviceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeviceState | undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["mac"] = state ? state.mac : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["portOverrides"] = state ? state.portOverrides : undefined;
            resourceInputs["site"] = state ? state.site : undefined;
        } else {
            const args = argsOrState as DeviceArgs | undefined;
            resourceInputs["mac"] = args ? args.mac : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["portOverrides"] = args ? args.portOverrides : undefined;
            resourceInputs["site"] = args ? args.site : undefined;
            resourceInputs["disabled"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Device.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Device resources.
 */
export interface DeviceState {
    /**
     * Specifies whether this device should be disabled.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * The MAC address of the device. This can be specified so that the provider can take control of a device (since devices are created through adoption).
     */
    mac?: pulumi.Input<string>;
    /**
     * The name of the device.
     */
    name?: pulumi.Input<string>;
    /**
     * Settings overrides for specific switch ports.
     */
    portOverrides?: pulumi.Input<pulumi.Input<inputs.DevicePortOverride>[]>;
    /**
     * The name of the site to associate the device with.
     */
    site?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Device resource.
 */
export interface DeviceArgs {
    /**
     * The MAC address of the device. This can be specified so that the provider can take control of a device (since devices are created through adoption).
     */
    mac?: pulumi.Input<string>;
    /**
     * The name of the device.
     */
    name?: pulumi.Input<string>;
    /**
     * Settings overrides for specific switch ports.
     */
    portOverrides?: pulumi.Input<pulumi.Input<inputs.DevicePortOverride>[]>;
    /**
     * The name of the site to associate the device with.
     */
    site?: pulumi.Input<string>;
}
