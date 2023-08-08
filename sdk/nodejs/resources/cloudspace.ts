// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Cloudspace extends pulumi.CustomResource {
    /**
     * Get an existing Cloudspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cloudspace {
        return new Cloudspace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vco:resources:Cloudspace';

    /**
     * Returns true if the given object is an instance of Cloudspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cloudspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cloudspace.__pulumiType;
    }

    public /*out*/ readonly cloudspace_id!: pulumi.Output<string>;
    public /*out*/ readonly cloudspace_mode!: pulumi.Output<string>;
    public /*out*/ readonly creation_time!: pulumi.Output<number>;
    public readonly customerID!: pulumi.Output<string>;
    public /*out*/ readonly external_network_id!: pulumi.Output<string>;
    public /*out*/ readonly external_network_ip!: pulumi.Output<string>;
    public readonly firewall_custom!: pulumi.Output<outputs.resources.FirewallCustom>;
    public readonly host!: pulumi.Output<string>;
    public readonly local_domain!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    public readonly memory_quota!: pulumi.Output<number>;
    public readonly name!: pulumi.Output<string>;
    public readonly parent_cloudspace_id!: pulumi.Output<string>;
    public readonly private!: pulumi.Output<boolean>;
    public /*out*/ readonly private_network!: pulumi.Output<string>;
    public readonly public_ip_quota!: pulumi.Output<number>;
    public /*out*/ readonly resource_limits!: pulumi.Output<outputs.resources.ResourceLimits>;
    public /*out*/ readonly router_type!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly token!: pulumi.Output<string>;
    public /*out*/ readonly update_time!: pulumi.Output<number>;
    public readonly url!: pulumi.Output<string>;
    public readonly vcpu_quota!: pulumi.Output<number>;
    public readonly vdisk_space_quota!: pulumi.Output<number>;

    /**
     * Create a Cloudspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudspaceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.externalNetworkID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalNetworkID'");
            }
            if ((!args || args.firewall_custom === undefined) && !opts.urn) {
                throw new Error("Missing required property 'firewall_custom'");
            }
            if ((!args || args.host === undefined) && !opts.urn) {
                throw new Error("Missing required property 'host'");
            }
            if ((!args || args.local_domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'local_domain'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.memory_quota === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memory_quota'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.parent_cloudspace_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parent_cloudspace_id'");
            }
            if ((!args || args.private === undefined) && !opts.urn) {
                throw new Error("Missing required property 'private'");
            }
            if ((!args || args.privateNetwork === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateNetwork'");
            }
            if ((!args || args.public_ip_quota === undefined) && !opts.urn) {
                throw new Error("Missing required property 'public_ip_quota'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            if ((!args || args.vcpu_quota === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vcpu_quota'");
            }
            if ((!args || args.vdisk_space_quota === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vdisk_space_quota'");
            }
            resourceInputs["customerID"] = args ? args.customerID : undefined;
            resourceInputs["externalNetworkID"] = args ? args.externalNetworkID : undefined;
            resourceInputs["firewall_custom"] = args ? args.firewall_custom : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["local_domain"] = args ? args.local_domain : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["memory_quota"] = args ? args.memory_quota : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parent_cloudspace_id"] = args ? args.parent_cloudspace_id : undefined;
            resourceInputs["private"] = args ? args.private : undefined;
            resourceInputs["privateNetwork"] = args ? args.privateNetwork : undefined;
            resourceInputs["public_ip_quota"] = args ? args.public_ip_quota : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["vcpu_quota"] = args ? args.vcpu_quota : undefined;
            resourceInputs["vdisk_space_quota"] = args ? args.vdisk_space_quota : undefined;
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["cloudspace_mode"] = undefined /*out*/;
            resourceInputs["creation_time"] = undefined /*out*/;
            resourceInputs["external_network_id"] = undefined /*out*/;
            resourceInputs["external_network_ip"] = undefined /*out*/;
            resourceInputs["private_network"] = undefined /*out*/;
            resourceInputs["resource_limits"] = undefined /*out*/;
            resourceInputs["router_type"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["update_time"] = undefined /*out*/;
        } else {
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["cloudspace_mode"] = undefined /*out*/;
            resourceInputs["creation_time"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["external_network_id"] = undefined /*out*/;
            resourceInputs["external_network_ip"] = undefined /*out*/;
            resourceInputs["firewall_custom"] = undefined /*out*/;
            resourceInputs["host"] = undefined /*out*/;
            resourceInputs["local_domain"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["memory_quota"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parent_cloudspace_id"] = undefined /*out*/;
            resourceInputs["private"] = undefined /*out*/;
            resourceInputs["private_network"] = undefined /*out*/;
            resourceInputs["public_ip_quota"] = undefined /*out*/;
            resourceInputs["resource_limits"] = undefined /*out*/;
            resourceInputs["router_type"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["update_time"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["vcpu_quota"] = undefined /*out*/;
            resourceInputs["vdisk_space_quota"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cloudspace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cloudspace resource.
 */
export interface CloudspaceArgs {
    customerID: pulumi.Input<string>;
    externalNetworkID: pulumi.Input<number>;
    firewall_custom: pulumi.Input<inputs.resources.FirewallCustomArgs>;
    host: pulumi.Input<string>;
    local_domain: pulumi.Input<string>;
    location: pulumi.Input<string>;
    memory_quota: pulumi.Input<number>;
    name: pulumi.Input<string>;
    parent_cloudspace_id: pulumi.Input<string>;
    private: pulumi.Input<boolean>;
    privateNetwork: pulumi.Input<string>;
    public_ip_quota: pulumi.Input<number>;
    token: pulumi.Input<string>;
    url: pulumi.Input<string>;
    vcpu_quota: pulumi.Input<number>;
    vdisk_space_quota: pulumi.Input<number>;
}
