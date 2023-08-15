"use strict";
// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
Object.defineProperty(exports, "__esModule", { value: true });
exports.VirtualMachineCD = void 0;
const pulumi = require("@pulumi/pulumi");
const utilities = require("../utilities");
class VirtualMachineCD extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachineCD resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name, id, opts) {
        return new VirtualMachineCD(name, undefined, Object.assign(Object.assign({}, opts), { id: id }));
    }
    /**
     * Returns true if the given object is an instance of VirtualMachineCD.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj) {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualMachineCD.__pulumiType;
    }
    /**
     * Create a VirtualMachineCD resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name, args, opts) {
        let resourceInputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.cdrom_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cdrom_id'");
            }
            if ((!args || args.cloudspace_id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cloudspace_id'");
            }
            if ((!args || args.customerID === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerID'");
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
            resourceInputs["cdrom_id"] = args ? args.cdrom_id : undefined;
            resourceInputs["cloudspace_id"] = args ? args.cloudspace_id : undefined;
            resourceInputs["customerID"] = (args === null || args === void 0 ? void 0 : args.customerID) ? pulumi.secret(args.customerID) : undefined;
            resourceInputs["token"] = (args === null || args === void 0 ? void 0 : args.token) ? pulumi.secret(args.token) : undefined;
            resourceInputs["url"] = (args === null || args === void 0 ? void 0 : args.url) ? pulumi.secret(args.url) : undefined;
            resourceInputs["vm_id"] = args ? args.vm_id : undefined;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["disk_size"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        else {
            resourceInputs["cdrom_id"] = undefined /*out*/;
            resourceInputs["cloudspace_id"] = undefined /*out*/;
            resourceInputs["customerID"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["disk_size"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["vm_id"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VirtualMachineCD.__pulumiType, name, resourceInputs, opts);
    }
}
exports.VirtualMachineCD = VirtualMachineCD;
/** @internal */
VirtualMachineCD.__pulumiType = 'vco:virtual_machine:VirtualMachineCD';
//# sourceMappingURL=virtualMachineCD.js.map