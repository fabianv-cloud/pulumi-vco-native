// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Virtual_machine
{
    [VcoResourceType("vco:virtual_machine:VirtualMachineDisk")]
    public partial class VirtualMachineDisk : global::Pulumi.CustomResource
    {
        [Output("cloudspace_id")]
        public Output<string> Cloudspace_id { get; private set; } = null!;

        [Output("customerID")]
        public Output<string> CustomerID { get; private set; } = null!;

        [Output("disk_id")]
        public Output<int> Disk_id { get; private set; } = null!;

        [Output("success")]
        public Output<bool> Success { get; private set; } = null!;

        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        [Output("vm_id")]
        public Output<int> Vm_id { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualMachineDisk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachineDisk(string name, VirtualMachineDiskArgs args, CustomResourceOptions? options = null)
            : base("vco:virtual_machine:VirtualMachineDisk", name, args ?? new VirtualMachineDiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachineDisk(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("vco:virtual_machine:VirtualMachineDisk", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualMachineDisk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualMachineDisk Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualMachineDisk(name, id, options);
        }
    }

    public sealed class VirtualMachineDiskArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudspace_id", required: true)]
        public Input<string> Cloudspace_id { get; set; } = null!;

        [Input("customerID", required: true)]
        public Input<string> CustomerID { get; set; } = null!;

        [Input("disk_id", required: true)]
        public Input<int> Disk_id { get; set; } = null!;

        [Input("token", required: true)]
        public Input<string> Token { get; set; } = null!;

        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        [Input("vm_id", required: true)]
        public Input<int> Vm_id { get; set; } = null!;

        public VirtualMachineDiskArgs()
        {
        }
        public static new VirtualMachineDiskArgs Empty => new VirtualMachineDiskArgs();
    }
}
