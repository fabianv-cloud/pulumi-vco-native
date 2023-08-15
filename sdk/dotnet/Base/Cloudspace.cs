// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Base
{
    [VcoResourceType("vco:base:Cloudspace")]
    public partial class Cloudspace : global::Pulumi.CustomResource
    {
        [Output("cdrom_id")]
        public Output<int?> Cdrom_id { get; private set; } = null!;

        [Output("cloudspace_id")]
        public Output<string> Cloudspace_id { get; private set; } = null!;

        [Output("cloudspace_mode")]
        public Output<string> Cloudspace_mode { get; private set; } = null!;

        [Output("creation_time")]
        public Output<int> Creation_time { get; private set; } = null!;

        [Output("customerID")]
        public Output<string> CustomerID { get; private set; } = null!;

        [Output("disk_size")]
        public Output<int?> Disk_size { get; private set; } = null!;

        [Output("external_network_id")]
        public Output<int> External_network_id { get; private set; } = null!;

        [Output("external_network_ip")]
        public Output<string> External_network_ip { get; private set; } = null!;

        [Output("host")]
        public Output<string?> Host { get; private set; } = null!;

        [Output("image_id")]
        public Output<int?> Image_id { get; private set; } = null!;

        [Output("local_domain")]
        public Output<string> Local_domain { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        [Output("memory")]
        public Output<int?> Memory { get; private set; } = null!;

        [Output("memory_quota")]
        public Output<int?> Memory_quota { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("parent_cloudspace_id")]
        public Output<string?> Parent_cloudspace_id { get; private set; } = null!;

        [Output("private")]
        public Output<bool> Private { get; private set; } = null!;

        [Output("private_network")]
        public Output<string> Private_network { get; private set; } = null!;

        [Output("public_ip_quota")]
        public Output<int?> Public_ip_quota { get; private set; } = null!;

        [Output("router_type")]
        public Output<string> Router_type { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        [Output("update_time")]
        public Output<int> Update_time { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        [Output("vcpu_quota")]
        public Output<int?> Vcpu_quota { get; private set; } = null!;

        [Output("vcpus")]
        public Output<int?> Vcpus { get; private set; } = null!;

        [Output("vdisk_space_quota")]
        public Output<int?> Vdisk_space_quota { get; private set; } = null!;


        /// <summary>
        /// Create a Cloudspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cloudspace(string name, CloudspaceArgs args, CustomResourceOptions? options = null)
            : base("vco:base:Cloudspace", name, args ?? new CloudspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cloudspace(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("vco:base:Cloudspace", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Cloudspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cloudspace Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Cloudspace(name, id, options);
        }
    }

    public sealed class CloudspaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("cdrom_id")]
        public Input<int>? Cdrom_id { get; set; }

        [Input("customerID", required: true)]
        private Input<string>? _customerID;
        public Input<string>? CustomerID
        {
            get => _customerID;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _customerID = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("disk_size")]
        public Input<int>? Disk_size { get; set; }

        [Input("external_network_id", required: true)]
        public Input<int> External_network_id { get; set; } = null!;

        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("image_id")]
        public Input<int>? Image_id { get; set; }

        [Input("local_domain")]
        public Input<string>? Local_domain { get; set; }

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("memory")]
        public Input<int>? Memory { get; set; }

        [Input("memory_quota")]
        public Input<int>? Memory_quota { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("parent_cloudspace_id")]
        public Input<string>? Parent_cloudspace_id { get; set; }

        [Input("private", required: true)]
        public Input<bool> Private { get; set; } = null!;

        [Input("private_network", required: true)]
        public Input<string> Private_network { get; set; } = null!;

        [Input("public_ip_quota")]
        public Input<int>? Public_ip_quota { get; set; }

        [Input("token", required: true)]
        private Input<string>? _token;
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("url", required: true)]
        private Input<string>? _url;
        public Input<string>? Url
        {
            get => _url;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _url = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("vcpu_quota")]
        public Input<int>? Vcpu_quota { get; set; }

        [Input("vcpus")]
        public Input<int>? Vcpus { get; set; }

        [Input("vdisk_space_quota")]
        public Input<int>? Vdisk_space_quota { get; set; }

        public CloudspaceArgs()
        {
        }
        public static new CloudspaceArgs Empty => new CloudspaceArgs();
    }
}
