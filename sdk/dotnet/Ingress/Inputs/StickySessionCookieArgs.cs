// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vco.Ingress.Inputs
{

    public sealed class StickySessionCookieArgs : global::Pulumi.ResourceArgs
    {
        [Input("http_only")]
        public Input<bool>? Http_only { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("same_site")]
        public Input<string>? Same_site { get; set; }

        [Input("secure")]
        public Input<bool>? Secure { get; set; }

        public StickySessionCookieArgs()
        {
        }
        public static new StickySessionCookieArgs Empty => new StickySessionCookieArgs();
    }
}
