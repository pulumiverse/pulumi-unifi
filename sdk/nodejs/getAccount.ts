// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `unifi.Account` data source can be used to retrieve RADIUS user accounts
 */
export function getAccount(args: GetAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("unifi:index/getAccount:getAccount", {
        "name": args.name,
        "site": args.site,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccount.
 */
export interface GetAccountArgs {
    /**
     * The name of the account to look up
     */
    name: string;
    /**
     * The name of the site the account is associated with.
     */
    site?: string;
}

/**
 * A collection of values returned by getAccount.
 */
export interface GetAccountResult {
    /**
     * The ID of this account.
     */
    readonly id: string;
    /**
     * The name of the account to look up
     */
    readonly name: string;
    /**
     * ID of the network for this account
     */
    readonly networkId: string;
    /**
     * The password of the account.
     */
    readonly password: string;
    /**
     * The name of the site the account is associated with.
     */
    readonly site: string;
    /**
     * See RFC2868 section 3.2
     */
    readonly tunnelMediumType: number;
    /**
     * See RFC2868 section 3.1
     */
    readonly tunnelType: number;
}

export function getAccountOutput(args: GetAccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountResult> {
    return pulumi.output(args).apply(a => getAccount(a, opts))
}

/**
 * A collection of arguments for invoking getAccount.
 */
export interface GetAccountOutputArgs {
    /**
     * The name of the account to look up
     */
    name: pulumi.Input<string>;
    /**
     * The name of the site the account is associated with.
     */
    site?: pulumi.Input<string>;
}
