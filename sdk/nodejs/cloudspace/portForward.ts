// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class PortForward extends pulumi.CustomResource {
    /**
     * Get an existing PortForward resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PortForward {
        return new PortForward(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vco:cloudspace:PortForward';

    /**
     * Returns true if the given object is an instance of PortForward.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PortForward {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PortForward.__pulumiType;
    }

    public readonly cloudspace_id!: pulumi.Output<string>;
    public readonly customerID!: pulumi.Output<string>;
    public readonly local_port!: pulumi.Output<number>;
    public readonly nested_cs_id!: pulumi.Output<string | undefined>;
    public /*out*/ readonly portforward_id!: pulumi.Output<string>;
    public readonly protocol!: pulumi.Output<string>;
    public readonly public_ip!: pulumi.Output<string>;
    public readonly public_port!: pulumi.Output<number>;
    public readonly till_public_port!: pulumi.Output<number | undefined>;
    public readonly token!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string>;
    public readonly vm_id!: pulumi.Output<number>;

    /**
     * Create a PortForward resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PortForwardArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cloudspace_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudspace_id'");
            }
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.local_port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'local_port'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.public_ip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'public_ip'");
            }
            if ((!args || args.public_port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'public_port'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            if ((!args || args.vm_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vm_id'");
            }
            resourceInputs["cloudspace_id"] = args ? args.cloudspace_id : undefined;
            resourceInputs["customerID"] = args?.customerID ? pulumi.secret(args.customerID) : undefined;
            resourceInputs["local_port"] = args ? args.local_port : undefined;
            resourceInputs["nested_cs_id"] = args ? args.nested_cs_id : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["public_ip"] = args ? args.public_ip : undefined;
            resourceInputs["public_port"] = args ? args.public_port : undefined;
            resourceInputs["till_public_port"] = args ? args.till_public_port : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = args?.url ? pulumi.secret(args.url) : undefined;
            resourceInputs["vm_id"] = args ? args.vm_id : undefined;
            resourceInputs["portforward_id"] = undefined /*out*/;
        } else {
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["local_port"] = undefined /*out*/;
            resourceInputs["nested_cs_id"] = undefined /*out*/;
            resourceInputs["portforward_id"] = undefined /*out*/;
            resourceInputs["protocol"] = undefined /*out*/;
            resourceInputs["public_ip"] = undefined /*out*/;
            resourceInputs["public_port"] = undefined /*out*/;
            resourceInputs["till_public_port"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["vm_id"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PortForward.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PortForward resource.
 */
export interface PortForwardArgs {
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    local_port: pulumi.Input<number>;
    nested_cs_id?: pulumi.Input<string>;
    protocol: pulumi.Input<string>;
    public_ip: pulumi.Input<string>;
    public_port: pulumi.Input<number>;
    till_public_port?: pulumi.Input<number>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
    vm_id: pulumi.Input<number>;
}
