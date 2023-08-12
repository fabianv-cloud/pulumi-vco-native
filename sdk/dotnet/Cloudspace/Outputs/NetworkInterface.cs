// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Cloudspace.Outputs
{

    [OutputType]
    public sealed class NetworkInterface
    {
        public readonly string Device_name;
        public readonly string External_cloudspace_id;
        public readonly string Ip_address;
        public readonly string Mac_address;
        public readonly string Model;
        public readonly int Network_id;
        public readonly string Nic_type;

        [OutputConstructor]
        private NetworkInterface(
            string device_name,

            string external_cloudspace_id,

            string ip_address,

            string mac_address,

            string model,

            int network_id,

            string nic_type)
        {
            Device_name = device_name;
            External_cloudspace_id = external_cloudspace_id;
            Ip_address = ip_address;
            Mac_address = mac_address;
            Model = model;
            Network_id = network_id;
            Nic_type = nic_type;
        }
    }
}
