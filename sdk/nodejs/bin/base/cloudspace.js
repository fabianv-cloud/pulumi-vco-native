"use strict";
// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
exports.Cloudspace = void 0;
const pulumi = require("@pulumi/pulumi");
const utilities = require("../utilities");
class Cloudspace extends pulumi.CustomResource {
    /**
     * Get an existing Cloudspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name, id, opts) {
        return new Cloudspace(name, undefined, Object.assign(Object.assign({}, opts), { id: id }));
    }
    /**
     * Returns true if the given object is an instance of Cloudspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj) {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cloudspace.__pulumiType;
    }
    /**
     * Create a Cloudspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name, args, opts) {
        let resourceInputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.external_network_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'external_network_id'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.private === undefined) && !opts.urn) {
                throw new Error("Missing required property 'private'");
            }
            if ((!args || args.private_network === undefined) && !opts.urn) {
                throw new Error("Missing required property 'private_network'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["cdrom_id"] = args ? args.cdrom_id : undefined;
            resourceInputs["customerID"] = (args === null || args === void 0 ? void 0 : args.customerID) ? pulumi.secret(args.customerID) : undefined;
            resourceInputs["disk_size"] = args ? args.disk_size : undefined;
            resourceInputs["external_network_id"] = args ? args.external_network_id : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["image_id"] = args ? args.image_id : undefined;
            resourceInputs["local_domain"] = args ? args.local_domain : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["memory"] = args ? args.memory : undefined;
            resourceInputs["memory_quota"] = args ? args.memory_quota : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parent_cloudspace_id"] = args ? args.parent_cloudspace_id : undefined;
            resourceInputs["private"] = args ? args.private : undefined;
            resourceInputs["private_network"] = args ? args.private_network : undefined;
            resourceInputs["public_ip_quota"] = args ? args.public_ip_quota : undefined;
            resourceInputs["token"] = (args === null || args === void 0 ? void 0 : args.token) ? pulumi.secret(args.token) : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["url"] = (args === null || args === void 0 ? void 0 : args.url) ? pulumi.secret(args.url) : undefined;
            resourceInputs["vcpu_quota"] = args ? args.vcpu_quota : undefined;
            resourceInputs["vcpus"] = args ? args.vcpus : undefined;
            resourceInputs["vdisk_space_quota"] = args ? args.vdisk_space_quota : undefined;
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["cloudspace_mode"] = undefined /*out*/;
            resourceInputs["creation_time"] = undefined /*out*/;
            resourceInputs["external_network_ip"] = undefined /*out*/;
            resourceInputs["router_type"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["update_time"] = undefined /*out*/;
        }
        else {
            resourceInputs["cdrom_id"] = undefined /*out*/;
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["cloudspace_mode"] = undefined /*out*/;
            resourceInputs["creation_time"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["disk_size"] = undefined /*out*/;
            resourceInputs["external_network_id"] = undefined /*out*/;
            resourceInputs["external_network_ip"] = undefined /*out*/;
            resourceInputs["host"] = undefined /*out*/;
            resourceInputs["image_id"] = undefined /*out*/;
            resourceInputs["local_domain"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["memory"] = undefined /*out*/;
            resourceInputs["memory_quota"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parent_cloudspace_id"] = undefined /*out*/;
            resourceInputs["private"] = undefined /*out*/;
            resourceInputs["private_network"] = undefined /*out*/;
            resourceInputs["public_ip_quota"] = undefined /*out*/;
            resourceInputs["router_type"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["update_time"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["vcpu_quota"] = undefined /*out*/;
            resourceInputs["vcpus"] = undefined /*out*/;
            resourceInputs["vdisk_space_quota"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cloudspace.__pulumiType, name, resourceInputs, opts);
    }
}
exports.Cloudspace = Cloudspace;
/** @internal */
Cloudspace.__pulumiType = 'vco:base:Cloudspace';
//# sourceMappingURL=cloudspace.js.map