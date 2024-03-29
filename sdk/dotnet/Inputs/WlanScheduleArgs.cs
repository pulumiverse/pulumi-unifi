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

    public sealed class WlanScheduleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Day of week for the block. Valid values are `sun`, `mon`, `tue`, `wed`, `thu`, `fri`, `sat`.
        /// </summary>
        [Input("dayOfWeek", required: true)]
        public Input<string> DayOfWeek { get; set; } = null!;

        /// <summary>
        /// Length of the block in minutes.
        /// </summary>
        [Input("duration", required: true)]
        public Input<int> Duration { get; set; } = null!;

        /// <summary>
        /// Name of the block.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Start hour for the block (0-23).
        /// </summary>
        [Input("startHour", required: true)]
        public Input<int> StartHour { get; set; } = null!;

        /// <summary>
        /// Start minute for the block (0-59). Defaults to `0`.
        /// </summary>
        [Input("startMinute")]
        public Input<int>? StartMinute { get; set; }

        public WlanScheduleArgs()
        {
        }
        public static new WlanScheduleArgs Empty => new WlanScheduleArgs();
    }
}
