// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CloudspaceArgs } from "./cloudspace";
export type Cloudspace = import("./cloudspace").Cloudspace;
export const Cloudspace: typeof import("./cloudspace").Cloudspace = null as any;
utilities.lazyLoad(exports, ["Cloudspace"], () => require("./cloudspace"));

export { DiskArgs } from "./disk";
export type Disk = import("./disk").Disk;
export const Disk: typeof import("./disk").Disk = null as any;
utilities.lazyLoad(exports, ["Disk"], () => require("./disk"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vco:base:Cloudspace":
                return new Cloudspace(name, <any>undefined, { urn })
            case "vco:base:Disk":
                return new Disk(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vco", "base", _module)
