"use strict";
// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
exports.ExternalNetwork = void 0;
const pulumi = require("@pulumi/pulumi");
const utilities = require("../utilities");
class ExternalNetwork extends pulumi.CustomResource {
    /**
     * Get an existing ExternalNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name, id, opts) {
        return new ExternalNetwork(name, undefined, Object.assign(Object.assign({}, opts), { id: id }));
    }
    /**
     * Returns true if the given object is an instance of ExternalNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj) {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExternalNetwork.__pulumiType;
    }
    /**
     * Create a ExternalNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name, args, opts) {
        let resourceInputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cloudspace_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudspace_id'");
            }
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
            }
            if ((!args || args.external_network_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'external_network_id'");
            }
            if ((!args || args.external_network_ip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'external_network_ip'");
            }
            if ((!args || args.external_network_type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'external_network_type'");
            }
            if ((!args || args.metric === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metric'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["cloudspace_id"] = args ? args.cloudspace_id : undefined;
            resourceInputs["customerID"] = (args === null || args === void 0 ? void 0 : args.customerID) ? pulumi.secret(args.customerID) : undefined;
            resourceInputs["external_network_id"] = args ? args.external_network_id : undefined;
            resourceInputs["external_network_ip"] = args ? args.external_network_ip : undefined;
            resourceInputs["external_network_type"] = args ? args.external_network_type : undefined;
            resourceInputs["metric"] = args ? args.metric : undefined;
            resourceInputs["token"] = (args === null || args === void 0 ? void 0 : args.token) ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = (args === null || args === void 0 ? void 0 : args.url) ? pulumi.secret(args.url) : undefined;
        }
        else {
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["external_network_id"] = undefined /*out*/;
            resourceInputs["external_network_ip"] = undefined /*out*/;
            resourceInputs["external_network_type"] = undefined /*out*/;
            resourceInputs["metric"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ExternalNetwork.__pulumiType, name, resourceInputs, opts);
    }
}
exports.ExternalNetwork = ExternalNetwork;
/** @internal */
ExternalNetwork.__pulumiType = 'vco:cloudspace:ExternalNetwork';
//# sourceMappingURL=externalNetwork.js.map