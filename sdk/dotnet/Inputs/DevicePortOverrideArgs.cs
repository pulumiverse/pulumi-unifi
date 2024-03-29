// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Unifi.Inputs
{

    public sealed class DevicePortOverrideArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of ports in the aggregate.
        /// </summary>
        [Input("aggregateNumPorts")]
        public Input<int>? AggregateNumPorts { get; set; }

        /// <summary>
        /// Human-readable name of the port.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Switch port number.
        /// </summary>
        [Input("number", required: true)]
        public Input<int> Number { get; set; } = null!;

        /// <summary>
        /// Operating mode of the port, valid values are `switch`, `mirror`, and `aggregate`. Defaults to `switch`.
        /// </summary>
        [Input("opMode")]
        public Input<string>? OpMode { get; set; }

        /// <summary>
        /// ID of the Port Profile used on this port.
        /// </summary>
        [Input("portProfileId")]
        public Input<string>? PortProfileId { get; set; }

        public DevicePortOverrideArgs()
        {
        }
        public static new DevicePortOverrideArgs Empty => new DevicePortOverrideArgs();
    }
}
