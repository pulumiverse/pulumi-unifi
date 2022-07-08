// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Unifi.Outputs
{

    [OutputType]
    public sealed class DevicePortOverride
    {
        /// <summary>
        /// Human-readable name of the port.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Switch port number.
        /// </summary>
        public readonly int Number;
        /// <summary>
        /// ID of the Port Profile used on this port.
        /// </summary>
        public readonly string? PortProfileId;

        [OutputConstructor]
        private DevicePortOverride(
            string? name,

            int number,

            string? portProfileId)
        {
            Name = name;
            Number = number;
            PortProfileId = portProfileId;
        }
    }
}