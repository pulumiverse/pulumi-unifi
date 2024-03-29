// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * `unifi.RadiusProfile` manages RADIUS profiles.
 */
export class RadiusProfile extends pulumi.CustomResource {
    /**
     * Get an existing RadiusProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RadiusProfileState, opts?: pulumi.CustomResourceOptions): RadiusProfile {
        return new RadiusProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'unifi:index/radiusProfile:RadiusProfile';

    /**
     * Returns true if the given object is an instance of RadiusProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RadiusProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RadiusProfile.__pulumiType;
    }

    /**
     * Specifies whether to use RADIUS accounting. Defaults to `false`.
     */
    public readonly accountingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * RADIUS accounting servers.
     */
    public readonly acctServers!: pulumi.Output<outputs.RadiusProfileAcctServer[] | undefined>;
    /**
     * RADIUS authentication servers.
     */
    public readonly authServers!: pulumi.Output<outputs.RadiusProfileAuthServer[] | undefined>;
    /**
     * Specifies whether to use interim_update. Defaults to `false`.
     */
    public readonly interimUpdateEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies interimUpdate interval. Defaults to `3600`.
     */
    public readonly interimUpdateInterval!: pulumi.Output<number | undefined>;
    /**
     * The name of the profile.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the site to associate the settings with.
     */
    public readonly site!: pulumi.Output<string>;
    /**
     * Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
     */
    public readonly useUsgAcctServer!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
     */
    public readonly useUsgAuthServer!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether to use vlan on wired connections. Defaults to `false`.
     */
    public readonly vlanEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
     */
    public readonly vlanWlanMode!: pulumi.Output<string | undefined>;

    /**
     * Create a RadiusProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RadiusProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RadiusProfileArgs | RadiusProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RadiusProfileState | undefined;
            resourceInputs["accountingEnabled"] = state ? state.accountingEnabled : undefined;
            resourceInputs["acctServers"] = state ? state.acctServers : undefined;
            resourceInputs["authServers"] = state ? state.authServers : undefined;
            resourceInputs["interimUpdateEnabled"] = state ? state.interimUpdateEnabled : undefined;
            resourceInputs["interimUpdateInterval"] = state ? state.interimUpdateInterval : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["site"] = state ? state.site : undefined;
            resourceInputs["useUsgAcctServer"] = state ? state.useUsgAcctServer : undefined;
            resourceInputs["useUsgAuthServer"] = state ? state.useUsgAuthServer : undefined;
            resourceInputs["vlanEnabled"] = state ? state.vlanEnabled : undefined;
            resourceInputs["vlanWlanMode"] = state ? state.vlanWlanMode : undefined;
        } else {
            const args = argsOrState as RadiusProfileArgs | undefined;
            resourceInputs["accountingEnabled"] = args ? args.accountingEnabled : undefined;
            resourceInputs["acctServers"] = args ? args.acctServers : undefined;
            resourceInputs["authServers"] = args ? args.authServers : undefined;
            resourceInputs["interimUpdateEnabled"] = args ? args.interimUpdateEnabled : undefined;
            resourceInputs["interimUpdateInterval"] = args ? args.interimUpdateInterval : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["site"] = args ? args.site : undefined;
            resourceInputs["useUsgAcctServer"] = args ? args.useUsgAcctServer : undefined;
            resourceInputs["useUsgAuthServer"] = args ? args.useUsgAuthServer : undefined;
            resourceInputs["vlanEnabled"] = args ? args.vlanEnabled : undefined;
            resourceInputs["vlanWlanMode"] = args ? args.vlanWlanMode : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RadiusProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RadiusProfile resources.
 */
export interface RadiusProfileState {
    /**
     * Specifies whether to use RADIUS accounting. Defaults to `false`.
     */
    accountingEnabled?: pulumi.Input<boolean>;
    /**
     * RADIUS accounting servers.
     */
    acctServers?: pulumi.Input<pulumi.Input<inputs.RadiusProfileAcctServer>[]>;
    /**
     * RADIUS authentication servers.
     */
    authServers?: pulumi.Input<pulumi.Input<inputs.RadiusProfileAuthServer>[]>;
    /**
     * Specifies whether to use interim_update. Defaults to `false`.
     */
    interimUpdateEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies interimUpdate interval. Defaults to `3600`.
     */
    interimUpdateInterval?: pulumi.Input<number>;
    /**
     * The name of the profile.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the site to associate the settings with.
     */
    site?: pulumi.Input<string>;
    /**
     * Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
     */
    useUsgAcctServer?: pulumi.Input<boolean>;
    /**
     * Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
     */
    useUsgAuthServer?: pulumi.Input<boolean>;
    /**
     * Specifies whether to use vlan on wired connections. Defaults to `false`.
     */
    vlanEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
     */
    vlanWlanMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RadiusProfile resource.
 */
export interface RadiusProfileArgs {
    /**
     * Specifies whether to use RADIUS accounting. Defaults to `false`.
     */
    accountingEnabled?: pulumi.Input<boolean>;
    /**
     * RADIUS accounting servers.
     */
    acctServers?: pulumi.Input<pulumi.Input<inputs.RadiusProfileAcctServer>[]>;
    /**
     * RADIUS authentication servers.
     */
    authServers?: pulumi.Input<pulumi.Input<inputs.RadiusProfileAuthServer>[]>;
    /**
     * Specifies whether to use interim_update. Defaults to `false`.
     */
    interimUpdateEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies interimUpdate interval. Defaults to `3600`.
     */
    interimUpdateInterval?: pulumi.Input<number>;
    /**
     * The name of the profile.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the site to associate the settings with.
     */
    site?: pulumi.Input<string>;
    /**
     * Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
     */
    useUsgAcctServer?: pulumi.Input<boolean>;
    /**
     * Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
     */
    useUsgAuthServer?: pulumi.Input<boolean>;
    /**
     * Specifies whether to use vlan on wired connections. Defaults to `false`.
     */
    vlanEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
     */
    vlanWlanMode?: pulumi.Input<string>;
}
