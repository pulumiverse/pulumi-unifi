// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `unifi.FirewallRule` manages an individual firewall rule on the gateway.
 *
 * ## Import
 *
 * # import using the ID from the controller API/UI
 *
 * ```sh
 *  $ pulumi import unifi:index/firewallRule:FirewallRule my_rule 5f7080eb6b8969064f80494f
 * ```
 */
export class FirewallRule extends pulumi.CustomResource {
    /**
     * Get an existing FirewallRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallRuleState, opts?: pulumi.CustomResourceOptions): FirewallRule {
        return new FirewallRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'unifi:index/firewallRule:FirewallRule';

    /**
     * Returns true if the given object is an instance of FirewallRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallRule.__pulumiType;
    }

    /**
     * The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * The destination address of the firewall rule.
     */
    public readonly dstAddress!: pulumi.Output<string | undefined>;
    /**
     * The destination firewall group IDs of the firewall rule.
     */
    public readonly dstFirewallGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The destination network ID of the firewall rule.
     */
    public readonly dstNetworkId!: pulumi.Output<string | undefined>;
    /**
     * The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     */
    public readonly dstNetworkType!: pulumi.Output<string | undefined>;
    /**
     * The destination port of the firewall rule.
     */
    public readonly dstPort!: pulumi.Output<string | undefined>;
    /**
     * ICMP type name.
     */
    public readonly icmpTypename!: pulumi.Output<string | undefined>;
    /**
     * Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
     */
    public readonly ipSec!: pulumi.Output<string | undefined>;
    /**
     * Enable logging for the firewall rule.
     */
    public readonly logging!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the firewall rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The protocol of the rule.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
     */
    public readonly ruleIndex!: pulumi.Output<number>;
    /**
     * The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
     */
    public readonly ruleset!: pulumi.Output<string>;
    /**
     * The name of the site to associate the firewall rule with.
     */
    public readonly site!: pulumi.Output<string>;
    /**
     * The source address for the firewall rule.
     */
    public readonly srcAddress!: pulumi.Output<string | undefined>;
    /**
     * The source firewall group IDs for the firewall rule.
     */
    public readonly srcFirewallGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The source MAC address of the firewall rule.
     */
    public readonly srcMac!: pulumi.Output<string | undefined>;
    /**
     * The source network ID for the firewall rule.
     */
    public readonly srcNetworkId!: pulumi.Output<string | undefined>;
    /**
     * The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     */
    public readonly srcNetworkType!: pulumi.Output<string | undefined>;
    /**
     * Match where the state is established.
     */
    public readonly stateEstablished!: pulumi.Output<boolean | undefined>;
    /**
     * Match where the state is invalid.
     */
    public readonly stateInvalid!: pulumi.Output<boolean | undefined>;
    /**
     * Match where the state is new.
     */
    public readonly stateNew!: pulumi.Output<boolean | undefined>;
    /**
     * Match where the state is related.
     */
    public readonly stateRelated!: pulumi.Output<boolean | undefined>;

    /**
     * Create a FirewallRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallRuleArgs | FirewallRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallRuleState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["dstAddress"] = state ? state.dstAddress : undefined;
            resourceInputs["dstFirewallGroupIds"] = state ? state.dstFirewallGroupIds : undefined;
            resourceInputs["dstNetworkId"] = state ? state.dstNetworkId : undefined;
            resourceInputs["dstNetworkType"] = state ? state.dstNetworkType : undefined;
            resourceInputs["dstPort"] = state ? state.dstPort : undefined;
            resourceInputs["icmpTypename"] = state ? state.icmpTypename : undefined;
            resourceInputs["ipSec"] = state ? state.ipSec : undefined;
            resourceInputs["logging"] = state ? state.logging : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["ruleIndex"] = state ? state.ruleIndex : undefined;
            resourceInputs["ruleset"] = state ? state.ruleset : undefined;
            resourceInputs["site"] = state ? state.site : undefined;
            resourceInputs["srcAddress"] = state ? state.srcAddress : undefined;
            resourceInputs["srcFirewallGroupIds"] = state ? state.srcFirewallGroupIds : undefined;
            resourceInputs["srcMac"] = state ? state.srcMac : undefined;
            resourceInputs["srcNetworkId"] = state ? state.srcNetworkId : undefined;
            resourceInputs["srcNetworkType"] = state ? state.srcNetworkType : undefined;
            resourceInputs["stateEstablished"] = state ? state.stateEstablished : undefined;
            resourceInputs["stateInvalid"] = state ? state.stateInvalid : undefined;
            resourceInputs["stateNew"] = state ? state.stateNew : undefined;
            resourceInputs["stateRelated"] = state ? state.stateRelated : undefined;
        } else {
            const args = argsOrState as FirewallRuleArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.ruleIndex === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleIndex'");
            }
            if ((!args || args.ruleset === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleset'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["dstAddress"] = args ? args.dstAddress : undefined;
            resourceInputs["dstFirewallGroupIds"] = args ? args.dstFirewallGroupIds : undefined;
            resourceInputs["dstNetworkId"] = args ? args.dstNetworkId : undefined;
            resourceInputs["dstNetworkType"] = args ? args.dstNetworkType : undefined;
            resourceInputs["dstPort"] = args ? args.dstPort : undefined;
            resourceInputs["icmpTypename"] = args ? args.icmpTypename : undefined;
            resourceInputs["ipSec"] = args ? args.ipSec : undefined;
            resourceInputs["logging"] = args ? args.logging : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["ruleIndex"] = args ? args.ruleIndex : undefined;
            resourceInputs["ruleset"] = args ? args.ruleset : undefined;
            resourceInputs["site"] = args ? args.site : undefined;
            resourceInputs["srcAddress"] = args ? args.srcAddress : undefined;
            resourceInputs["srcFirewallGroupIds"] = args ? args.srcFirewallGroupIds : undefined;
            resourceInputs["srcMac"] = args ? args.srcMac : undefined;
            resourceInputs["srcNetworkId"] = args ? args.srcNetworkId : undefined;
            resourceInputs["srcNetworkType"] = args ? args.srcNetworkType : undefined;
            resourceInputs["stateEstablished"] = args ? args.stateEstablished : undefined;
            resourceInputs["stateInvalid"] = args ? args.stateInvalid : undefined;
            resourceInputs["stateNew"] = args ? args.stateNew : undefined;
            resourceInputs["stateRelated"] = args ? args.stateRelated : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallRule resources.
 */
export interface FirewallRuleState {
    /**
     * The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
     */
    action?: pulumi.Input<string>;
    /**
     * The destination address of the firewall rule.
     */
    dstAddress?: pulumi.Input<string>;
    /**
     * The destination firewall group IDs of the firewall rule.
     */
    dstFirewallGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The destination network ID of the firewall rule.
     */
    dstNetworkId?: pulumi.Input<string>;
    /**
     * The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     */
    dstNetworkType?: pulumi.Input<string>;
    /**
     * The destination port of the firewall rule.
     */
    dstPort?: pulumi.Input<string>;
    /**
     * ICMP type name.
     */
    icmpTypename?: pulumi.Input<string>;
    /**
     * Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
     */
    ipSec?: pulumi.Input<string>;
    /**
     * Enable logging for the firewall rule.
     */
    logging?: pulumi.Input<boolean>;
    /**
     * The name of the firewall rule.
     */
    name?: pulumi.Input<string>;
    /**
     * The protocol of the rule.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
     */
    ruleIndex?: pulumi.Input<number>;
    /**
     * The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
     */
    ruleset?: pulumi.Input<string>;
    /**
     * The name of the site to associate the firewall rule with.
     */
    site?: pulumi.Input<string>;
    /**
     * The source address for the firewall rule.
     */
    srcAddress?: pulumi.Input<string>;
    /**
     * The source firewall group IDs for the firewall rule.
     */
    srcFirewallGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The source MAC address of the firewall rule.
     */
    srcMac?: pulumi.Input<string>;
    /**
     * The source network ID for the firewall rule.
     */
    srcNetworkId?: pulumi.Input<string>;
    /**
     * The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     */
    srcNetworkType?: pulumi.Input<string>;
    /**
     * Match where the state is established.
     */
    stateEstablished?: pulumi.Input<boolean>;
    /**
     * Match where the state is invalid.
     */
    stateInvalid?: pulumi.Input<boolean>;
    /**
     * Match where the state is new.
     */
    stateNew?: pulumi.Input<boolean>;
    /**
     * Match where the state is related.
     */
    stateRelated?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a FirewallRule resource.
 */
export interface FirewallRuleArgs {
    /**
     * The action of the firewall rule. Must be one of `drop`, `accept`, or `reject`.
     */
    action: pulumi.Input<string>;
    /**
     * The destination address of the firewall rule.
     */
    dstAddress?: pulumi.Input<string>;
    /**
     * The destination firewall group IDs of the firewall rule.
     */
    dstFirewallGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The destination network ID of the firewall rule.
     */
    dstNetworkId?: pulumi.Input<string>;
    /**
     * The destination network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     */
    dstNetworkType?: pulumi.Input<string>;
    /**
     * The destination port of the firewall rule.
     */
    dstPort?: pulumi.Input<string>;
    /**
     * ICMP type name.
     */
    icmpTypename?: pulumi.Input<string>;
    /**
     * Specify whether the rule matches on IPsec packets. Can be one of `match-ipset` or `match-none`.
     */
    ipSec?: pulumi.Input<string>;
    /**
     * Enable logging for the firewall rule.
     */
    logging?: pulumi.Input<boolean>;
    /**
     * The name of the firewall rule.
     */
    name?: pulumi.Input<string>;
    /**
     * The protocol of the rule.
     */
    protocol: pulumi.Input<string>;
    /**
     * The index of the rule. Must be >= 2000 < 3000 or >= 4000 < 5000.
     */
    ruleIndex: pulumi.Input<number>;
    /**
     * The ruleset for the rule. This is from the perspective of the security gateway. Must be one of `WAN_IN`, `WAN_OUT`, `WAN_LOCAL`, `LAN_IN`, `LAN_OUT`, `LAN_LOCAL`, `GUEST_IN`, `GUEST_OUT`, `GUEST_LOCAL`, `WANv6_IN`, `WANv6_OUT`, `WANv6_LOCAL`, `LANv6_IN`, `LANv6_OUT`, `LANv6_LOCAL`, `GUESTv6_IN`, `GUESTv6_OUT`, or `GUESTv6_LOCAL`.
     */
    ruleset: pulumi.Input<string>;
    /**
     * The name of the site to associate the firewall rule with.
     */
    site?: pulumi.Input<string>;
    /**
     * The source address for the firewall rule.
     */
    srcAddress?: pulumi.Input<string>;
    /**
     * The source firewall group IDs for the firewall rule.
     */
    srcFirewallGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The source MAC address of the firewall rule.
     */
    srcMac?: pulumi.Input<string>;
    /**
     * The source network ID for the firewall rule.
     */
    srcNetworkId?: pulumi.Input<string>;
    /**
     * The source network type of the firewall rule. Can be one of `ADDRv4` or `NETv4`. Defaults to `NETv4`.
     */
    srcNetworkType?: pulumi.Input<string>;
    /**
     * Match where the state is established.
     */
    stateEstablished?: pulumi.Input<boolean>;
    /**
     * Match where the state is invalid.
     */
    stateInvalid?: pulumi.Input<boolean>;
    /**
     * Match where the state is new.
     */
    stateNew?: pulumi.Input<boolean>;
    /**
     * Match where the state is related.
     */
    stateRelated?: pulumi.Input<boolean>;
}
