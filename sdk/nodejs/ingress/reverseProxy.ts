// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ReverseProxy extends pulumi.CustomResource {
    /**
     * Get an existing ReverseProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReverseProxy {
        return new ReverseProxy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vco:ingress:ReverseProxy';

    /**
     * Returns true if the given object is an instance of ReverseProxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReverseProxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReverseProxy.__pulumiType;
    }

    public readonly back_end!: pulumi.Output<outputs.ingress.ReverseProxyBackend>;
    public readonly cloudspace_id!: pulumi.Output<string>;
    public readonly customerID!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string>;
    public readonly front_end!: pulumi.Output<outputs.ingress.ReverseProxyFrontEnd>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly reverseproxy_id!: pulumi.Output<string>;
    public readonly token!: pulumi.Output<string>;
    public /*out*/ readonly type!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ReverseProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReverseProxyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.back_end === undefined) && !opts.urn) {
                throw new Error("Missing required property 'back_end'");
            }
            if ((!args || args.cloudspace_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudspace_id'");
            }
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.front_end === undefined) && !opts.urn) {
                throw new Error("Missing required property 'front_end'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["back_end"] = args ? args.back_end : undefined;
            resourceInputs["cloudspace_id"] = args ? args.cloudspace_id : undefined;
            resourceInputs["customerID"] = args?.customerID ? pulumi.secret(args.customerID) : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["front_end"] = args ? args.front_end : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["token"] = args?.token ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = args?.url ? pulumi.secret(args.url) : undefined;
            resourceInputs["reverseproxy_id"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["back_end"] = undefined /*out*/;
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["front_end"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["reverseproxy_id"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReverseProxy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ReverseProxy resource.
 */
export interface ReverseProxyArgs {
    back_end: pulumi.Input<inputs.ingress.ReverseProxyBackendArgs>;
    cloudspace_id: pulumi.Input<string>;
    customerID: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    front_end: pulumi.Input<inputs.ingress.ReverseProxyFrontEndArgs>;
    name: pulumi.Input<string>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
