// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ExposedDisk extends pulumi.CustomResource {
    /**
     * Get an existing ExposedDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ExposedDisk {
        return new ExposedDisk(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vco:disk:ExposedDisk';

    /**
     * Returns true if the given object is an instance of ExposedDisk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExposedDisk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExposedDisk.__pulumiType;
    }

    public readonly cloudspace_id!: pulumi.Output<string>;
    public readonly customerID!: pulumi.Output<string>;
    public readonly disk_id!: pulumi.Output<number>;
    public /*out*/ readonly endpoint!: pulumi.Output<outputs.disk.Endpoint>;
    public readonly iops!: pulumi.Output<number | undefined>;
    public readonly max_connections!: pulumi.Output<number | undefined>;
    public /*out*/ readonly protocol!: pulumi.Output<string>;
    public readonly token!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ExposedDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExposedDiskArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cloudspace_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudspace_id'");
            }
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.disk_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'disk_id'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["cloudspace_id"] = args ? args.cloudspace_id : undefined;
            resourceInputs["customerID"] = args ? args.customerID : undefined;
            resourceInputs["disk_id"] = args ? args.disk_id : undefined;
            resourceInputs["iops"] = args ? args.iops : undefined;
            resourceInputs["max_connections"] = args ? args.max_connections : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["protocol"] = undefined /*out*/;
        } else {
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["disk_id"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["iops"] = undefined /*out*/;
            resourceInputs["max_connections"] = undefined /*out*/;
            resourceInputs["protocol"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ExposedDisk.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ExposedDisk resource.
 */
export interface ExposedDiskArgs {
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    disk_id: pulumi.Input<number>;
    iops?: pulumi.Input<number>;
    max_connections?: pulumi.Input<number>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
