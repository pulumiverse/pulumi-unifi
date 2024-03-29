// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { MgmtArgs, MgmtState } from "./mgmt";
export type Mgmt = import("./mgmt").Mgmt;
export const Mgmt: typeof import("./mgmt").Mgmt = null as any;
utilities.lazyLoad(exports, ["Mgmt"], () => require("./mgmt"));

export { RadiusArgs, RadiusState } from "./radius";
export type Radius = import("./radius").Radius;
export const Radius: typeof import("./radius").Radius = null as any;
utilities.lazyLoad(exports, ["Radius"], () => require("./radius"));

export { USGArgs, USGState } from "./usg";
export type USG = import("./usg").USG;
export const USG: typeof import("./usg").USG = null as any;
utilities.lazyLoad(exports, ["USG"], () => require("./usg"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "unifi:setting/mgmt:Mgmt":
                return new Mgmt(name, <any>undefined, { urn })
            case "unifi:setting/radius:Radius":
                return new Radius(name, <any>undefined, { urn })
            case "unifi:setting/uSG:USG":
                return new USG(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("unifi", "setting/mgmt", _module)
pulumi.runtime.registerResourceModule("unifi", "setting/radius", _module)
pulumi.runtime.registerResourceModule("unifi", "setting/uSG", _module)
