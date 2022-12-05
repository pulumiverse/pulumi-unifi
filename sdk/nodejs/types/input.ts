// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export * as setting from "./setting/input";
export interface DevicePortOverride {
    /**
     * Human-readable name of the port.
     */
    name?: pulumi.Input<string>;
    /**
     * Switch port number.
     */
    number: pulumi.Input<number>;
    /**
     * ID of the Port Profile used on this port.
     */
    portProfileId?: pulumi.Input<string>;
}

export interface RadiusProfileAcctServer {
    /**
     * IP address of accounting service server.
     */
    ip: pulumi.Input<string>;
    /**
     * Port of accounting service. Defaults to `1813`.
     */
    port?: pulumi.Input<number>;
    /**
     * RADIUS secret.
     */
    xsecret: pulumi.Input<string>;
}

export interface RadiusProfileAuthServer {
    /**
     * IP address of authentication service server.
     */
    ip: pulumi.Input<string>;
    /**
     * Port of authentication service. Defaults to `1812`.
     */
    port?: pulumi.Input<number>;
    /**
     * RADIUS secret.
     */
    xsecret: pulumi.Input<string>;
}

export interface WlanSchedule {
    /**
     * Day of week for the block. Valid values are `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.
     */
    dayOfWeek: pulumi.Input<string>;
    /**
     * Length of the block in minutes.
     */
    duration: pulumi.Input<number>;
    /**
     * Name of the block.
     */
    name?: pulumi.Input<string>;
    /**
     * Start hour for the block (0-23).
     */
    startHour: pulumi.Input<number>;
    /**
     * Start minute for the block (0-59). Defaults to `0`.
     */
    startMinute?: pulumi.Input<number>;
}
