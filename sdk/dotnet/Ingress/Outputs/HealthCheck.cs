// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Ingress.Outputs
{

    [OutputType]
    public sealed class HealthCheck
    {
        public readonly int? Interval;
        public readonly string? Path;
        public readonly int? Port;
        public readonly string? Scheme;
        public readonly int? Timeout;

        [OutputConstructor]
        private HealthCheck(
            int? interval,

            string? path,

            int? port,

            string? scheme,

            int? timeout)
        {
            Interval = interval;
            Path = path;
            Port = port;
            Scheme = scheme;
            Timeout = timeout;
        }
    }
}
