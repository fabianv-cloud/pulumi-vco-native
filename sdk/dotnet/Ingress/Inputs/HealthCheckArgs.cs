// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Ingress.Inputs
{

    public sealed class HealthCheckArgs : global::Pulumi.ResourceArgs
    {
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public HealthCheckArgs()
        {
        }
        public static new HealthCheckArgs Empty => new HealthCheckArgs();
    }
}
