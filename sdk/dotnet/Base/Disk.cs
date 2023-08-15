// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Base
{
    [VcoResourceType("vco:base:Disk")]
    public partial class Disk : global::Pulumi.CustomResource
    {
        [Output("cloudspace_id")]
        public Output<string> Cloudspace_id { get; private set; } = null!;

        [Output("customerID")]
        public Output<string> CustomerID { get; private set; } = null!;

        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("detach")]
        public Output<bool?> Detach { get; private set; } = null!;

        [Output("disk_id")]
        public Output<int> Disk_id { get; private set; } = null!;

        [Output("disk_name")]
        public Output<string> Disk_name { get; private set; } = null!;

        [Output("disk_size")]
        public Output<int> Disk_size { get; private set; } = null!;

        [Output("disk_type")]
        public Output<string> Disk_type { get; private set; } = null!;

        [Output("exposed")]
        public Output<bool> Exposed { get; private set; } = null!;

        [Output("iops")]
        public Output<int?> Iops { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        [Output("model")]
        public Output<string> Model { get; private set; } = null!;

        [Output("order")]
        public Output<string> Order { get; private set; } = null!;

        [Output("permanently")]
        public Output<bool?> Permanently { get; private set; } = null!;

        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        [Output("vm_id")]
        public Output<int> Vm_id { get; private set; } = null!;


        /// <summary>
        /// Create a Disk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Disk(string name, DiskArgs args, CustomResourceOptions? options = null)
            : base("vco:base:Disk", name, args ?? new DiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Disk(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("vco:base:Disk", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Disk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Disk Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Disk(name, id, options);
        }
    }

    public sealed class DiskArgs : global::Pulumi.ResourceArgs
    {
        [Input("customerID", required: true)]
        public Input<string> CustomerID { get; set; } = null!;

        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("detach")]
        public Input<bool>? Detach { get; set; }

        [Input("disk_name", required: true)]
        public Input<string> Disk_name { get; set; } = null!;

        [Input("disk_size", required: true)]
        public Input<int> Disk_size { get; set; } = null!;

        [Input("disk_type")]
        public Input<string>? Disk_type { get; set; }

        [Input("iops")]
        public Input<int>? Iops { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("permanently")]
        public Input<bool>? Permanently { get; set; }

        [Input("token", required: true)]
        public Input<string> Token { get; set; } = null!;

        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        [Input("vm_id")]
        public Input<string>? Vm_id { get; set; }

        public DiskArgs()
        {
        }
        public static new DiskArgs Empty => new DiskArgs();
    }
}
