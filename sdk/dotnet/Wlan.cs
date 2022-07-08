// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Unifi
{
    /// <summary>
    /// `unifi.Wlan` manages a WiFi network / SSID.
    /// 
    /// ## Import
    /// 
    /// # import from provider configured site
    /// 
    /// ```sh
    ///  $ pulumi import unifi:index/wlan:Wlan mywlan 5dc28e5e9106d105bdc87217
    /// ```
    /// 
    /// # import from another site
    /// 
    /// ```sh
    ///  $ pulumi import unifi:index/wlan:Wlan mywlan bfa2l6i7:5dc28e5e9106d105bdc87217
    /// ```
    /// </summary>
    [UnifiResourceType("unifi:index/wlan:Wlan")]
    public partial class Wlan : Pulumi.CustomResource
    {
        /// <summary>
        /// IDs of the AP groups to use for this network.
        /// </summary>
        [Output("apGroupIds")]
        public Output<ImmutableArray<string>> ApGroupIds { get; private set; } = null!;

        /// <summary>
        /// Indicates whether or not to hide the SSID from broadcast.
        /// </summary>
        [Output("hideSsid")]
        public Output<bool?> HideSsid { get; private set; } = null!;

        /// <summary>
        /// Indicates that this is a guest WLAN and should use guest behaviors.
        /// </summary>
        [Output("isGuest")]
        public Output<bool?> IsGuest { get; private set; } = null!;

        /// <summary>
        /// Isolates stations on layer 2 (ethernet) level Defaults to `false`.
        /// </summary>
        [Output("l2Isolation")]
        public Output<bool?> L2Isolation { get; private set; } = null!;

        /// <summary>
        /// Indicates whether or not the MAC filter is turned of for the network.
        /// </summary>
        [Output("macFilterEnabled")]
        public Output<bool?> MacFilterEnabled { get; private set; } = null!;

        /// <summary>
        /// List of MAC addresses to filter (only valid if `mac_filter_enabled` is `true`).
        /// </summary>
        [Output("macFilterLists")]
        public Output<ImmutableArray<string>> MacFilterLists { get; private set; } = null!;

        /// <summary>
        /// MAC address filter policy (only valid if `mac_filter_enabled` is `true`). Defaults to `deny`.
        /// </summary>
        [Output("macFilterPolicy")]
        public Output<string?> MacFilterPolicy { get; private set; } = null!;

        /// <summary>
        /// Set minimum data rate control for 2G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `1000`, `2000`, `5500`, `6000`, `9000`, `11000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
        /// </summary>
        [Output("minimumDataRate2gKbps")]
        public Output<int?> MinimumDataRate2gKbps { get; private set; } = null!;

        /// <summary>
        /// Set minimum data rate control for 5G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `6000`, `9000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
        /// </summary>
        [Output("minimumDataRate5gKbps")]
        public Output<int?> MinimumDataRate5gKbps { get; private set; } = null!;

        /// <summary>
        /// Indicates whether or not Multicast Enhance is turned of for the network.
        /// </summary>
        [Output("multicastEnhance")]
        public Output<bool?> MulticastEnhance { get; private set; } = null!;

        /// <summary>
        /// The SSID of the network.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the network for this SSID
        /// </summary>
        [Output("networkId")]
        public Output<string?> NetworkId { get; private set; } = null!;

        /// <summary>
        /// Connect high performance clients to 5 GHz only Defaults to `true`.
        /// </summary>
        [Output("no2ghzOui")]
        public Output<bool?> No2ghzOui { get; private set; } = null!;

        /// <summary>
        /// The passphrase for the network, this is only required if `security` is not set to `open`.
        /// </summary>
        [Output("passphrase")]
        public Output<string?> Passphrase { get; private set; } = null!;

        /// <summary>
        /// Enable Protected Management Frames. This cannot be disabled if using WPA 3. Valid values are `required`, `optional` and `disabled`. Defaults to `disabled`.
        /// </summary>
        [Output("pmfMode")]
        public Output<string?> PmfMode { get; private set; } = null!;

        /// <summary>
        /// ID of the RADIUS profile to use when security `wpaeap`. You can query this via the `unifi.getRadiusProfile` data source.
        /// </summary>
        [Output("radiusProfileId")]
        public Output<string?> RadiusProfileId { get; private set; } = null!;

        /// <summary>
        /// Start and stop schedules for the WLAN
        /// </summary>
        [Output("schedules")]
        public Output<ImmutableArray<Outputs.WlanSchedule>> Schedules { get; private set; } = null!;

        /// <summary>
        /// The type of WiFi security for this network. Valid values are: `wpapsk`, `wpaeap`, and `open`.
        /// </summary>
        [Output("security")]
        public Output<string> Security { get; private set; } = null!;

        /// <summary>
        /// The name of the site to associate the wlan with.
        /// </summary>
        [Output("site")]
        public Output<string> Site { get; private set; } = null!;

        /// <summary>
        /// Enable Unscheduled Automatic Power Save Delivery Defaults to `false`.
        /// </summary>
        [Output("uapsd")]
        public Output<bool?> Uapsd { get; private set; } = null!;

        /// <summary>
        /// ID of the user group to use for this network.
        /// </summary>
        [Output("userGroupId")]
        public Output<string> UserGroupId { get; private set; } = null!;

        /// <summary>
        /// Radio band your WiFi network will use.
        /// </summary>
        [Output("wlanBand")]
        public Output<string?> WlanBand { get; private set; } = null!;

        /// <summary>
        /// Enable WPA 3 support (security must be `wpapsk` and PMF must be turned on).
        /// </summary>
        [Output("wpa3Support")]
        public Output<bool?> Wpa3Support { get; private set; } = null!;

        /// <summary>
        /// Enable WPA 3 and WPA 2 support (security must be `wpapsk` and `wpa3_support` must be true).
        /// </summary>
        [Output("wpa3Transition")]
        public Output<bool?> Wpa3Transition { get; private set; } = null!;


        /// <summary>
        /// Create a Wlan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Wlan(string name, WlanArgs args, CustomResourceOptions? options = null)
            : base("unifi:index/wlan:Wlan", name, args ?? new WlanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Wlan(string name, Input<string> id, WlanState? state = null, CustomResourceOptions? options = null)
            : base("unifi:index/wlan:Wlan", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/pulumiverse/pulumi-unifi/releases/download/${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Wlan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Wlan Get(string name, Input<string> id, WlanState? state = null, CustomResourceOptions? options = null)
        {
            return new Wlan(name, id, state, options);
        }
    }

    public sealed class WlanArgs : Pulumi.ResourceArgs
    {
        [Input("apGroupIds")]
        private InputList<string>? _apGroupIds;

        /// <summary>
        /// IDs of the AP groups to use for this network.
        /// </summary>
        public InputList<string> ApGroupIds
        {
            get => _apGroupIds ?? (_apGroupIds = new InputList<string>());
            set => _apGroupIds = value;
        }

        /// <summary>
        /// Indicates whether or not to hide the SSID from broadcast.
        /// </summary>
        [Input("hideSsid")]
        public Input<bool>? HideSsid { get; set; }

        /// <summary>
        /// Indicates that this is a guest WLAN and should use guest behaviors.
        /// </summary>
        [Input("isGuest")]
        public Input<bool>? IsGuest { get; set; }

        /// <summary>
        /// Isolates stations on layer 2 (ethernet) level Defaults to `false`.
        /// </summary>
        [Input("l2Isolation")]
        public Input<bool>? L2Isolation { get; set; }

        /// <summary>
        /// Indicates whether or not the MAC filter is turned of for the network.
        /// </summary>
        [Input("macFilterEnabled")]
        public Input<bool>? MacFilterEnabled { get; set; }

        [Input("macFilterLists")]
        private InputList<string>? _macFilterLists;

        /// <summary>
        /// List of MAC addresses to filter (only valid if `mac_filter_enabled` is `true`).
        /// </summary>
        public InputList<string> MacFilterLists
        {
            get => _macFilterLists ?? (_macFilterLists = new InputList<string>());
            set => _macFilterLists = value;
        }

        /// <summary>
        /// MAC address filter policy (only valid if `mac_filter_enabled` is `true`). Defaults to `deny`.
        /// </summary>
        [Input("macFilterPolicy")]
        public Input<string>? MacFilterPolicy { get; set; }

        /// <summary>
        /// Set minimum data rate control for 2G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `1000`, `2000`, `5500`, `6000`, `9000`, `11000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
        /// </summary>
        [Input("minimumDataRate2gKbps")]
        public Input<int>? MinimumDataRate2gKbps { get; set; }

        /// <summary>
        /// Set minimum data rate control for 5G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `6000`, `9000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
        /// </summary>
        [Input("minimumDataRate5gKbps")]
        public Input<int>? MinimumDataRate5gKbps { get; set; }

        /// <summary>
        /// Indicates whether or not Multicast Enhance is turned of for the network.
        /// </summary>
        [Input("multicastEnhance")]
        public Input<bool>? MulticastEnhance { get; set; }

        /// <summary>
        /// The SSID of the network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the network for this SSID
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// Connect high performance clients to 5 GHz only Defaults to `true`.
        /// </summary>
        [Input("no2ghzOui")]
        public Input<bool>? No2ghzOui { get; set; }

        /// <summary>
        /// The passphrase for the network, this is only required if `security` is not set to `open`.
        /// </summary>
        [Input("passphrase")]
        public Input<string>? Passphrase { get; set; }

        /// <summary>
        /// Enable Protected Management Frames. This cannot be disabled if using WPA 3. Valid values are `required`, `optional` and `disabled`. Defaults to `disabled`.
        /// </summary>
        [Input("pmfMode")]
        public Input<string>? PmfMode { get; set; }

        /// <summary>
        /// ID of the RADIUS profile to use when security `wpaeap`. You can query this via the `unifi.getRadiusProfile` data source.
        /// </summary>
        [Input("radiusProfileId")]
        public Input<string>? RadiusProfileId { get; set; }

        [Input("schedules")]
        private InputList<Inputs.WlanScheduleArgs>? _schedules;

        /// <summary>
        /// Start and stop schedules for the WLAN
        /// </summary>
        public InputList<Inputs.WlanScheduleArgs> Schedules
        {
            get => _schedules ?? (_schedules = new InputList<Inputs.WlanScheduleArgs>());
            set => _schedules = value;
        }

        /// <summary>
        /// The type of WiFi security for this network. Valid values are: `wpapsk`, `wpaeap`, and `open`.
        /// </summary>
        [Input("security", required: true)]
        public Input<string> Security { get; set; } = null!;

        /// <summary>
        /// The name of the site to associate the wlan with.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        /// <summary>
        /// Enable Unscheduled Automatic Power Save Delivery Defaults to `false`.
        /// </summary>
        [Input("uapsd")]
        public Input<bool>? Uapsd { get; set; }

        /// <summary>
        /// ID of the user group to use for this network.
        /// </summary>
        [Input("userGroupId", required: true)]
        public Input<string> UserGroupId { get; set; } = null!;

        /// <summary>
        /// Radio band your WiFi network will use.
        /// </summary>
        [Input("wlanBand")]
        public Input<string>? WlanBand { get; set; }

        /// <summary>
        /// Enable WPA 3 support (security must be `wpapsk` and PMF must be turned on).
        /// </summary>
        [Input("wpa3Support")]
        public Input<bool>? Wpa3Support { get; set; }

        /// <summary>
        /// Enable WPA 3 and WPA 2 support (security must be `wpapsk` and `wpa3_support` must be true).
        /// </summary>
        [Input("wpa3Transition")]
        public Input<bool>? Wpa3Transition { get; set; }

        public WlanArgs()
        {
        }
    }

    public sealed class WlanState : Pulumi.ResourceArgs
    {
        [Input("apGroupIds")]
        private InputList<string>? _apGroupIds;

        /// <summary>
        /// IDs of the AP groups to use for this network.
        /// </summary>
        public InputList<string> ApGroupIds
        {
            get => _apGroupIds ?? (_apGroupIds = new InputList<string>());
            set => _apGroupIds = value;
        }

        /// <summary>
        /// Indicates whether or not to hide the SSID from broadcast.
        /// </summary>
        [Input("hideSsid")]
        public Input<bool>? HideSsid { get; set; }

        /// <summary>
        /// Indicates that this is a guest WLAN and should use guest behaviors.
        /// </summary>
        [Input("isGuest")]
        public Input<bool>? IsGuest { get; set; }

        /// <summary>
        /// Isolates stations on layer 2 (ethernet) level Defaults to `false`.
        /// </summary>
        [Input("l2Isolation")]
        public Input<bool>? L2Isolation { get; set; }

        /// <summary>
        /// Indicates whether or not the MAC filter is turned of for the network.
        /// </summary>
        [Input("macFilterEnabled")]
        public Input<bool>? MacFilterEnabled { get; set; }

        [Input("macFilterLists")]
        private InputList<string>? _macFilterLists;

        /// <summary>
        /// List of MAC addresses to filter (only valid if `mac_filter_enabled` is `true`).
        /// </summary>
        public InputList<string> MacFilterLists
        {
            get => _macFilterLists ?? (_macFilterLists = new InputList<string>());
            set => _macFilterLists = value;
        }

        /// <summary>
        /// MAC address filter policy (only valid if `mac_filter_enabled` is `true`). Defaults to `deny`.
        /// </summary>
        [Input("macFilterPolicy")]
        public Input<string>? MacFilterPolicy { get; set; }

        /// <summary>
        /// Set minimum data rate control for 2G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `1000`, `2000`, `5500`, `6000`, `9000`, `11000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
        /// </summary>
        [Input("minimumDataRate2gKbps")]
        public Input<int>? MinimumDataRate2gKbps { get; set; }

        /// <summary>
        /// Set minimum data rate control for 5G devices, in Kbps. Use `0` to disable minimum data rates. Valid values are: `6000`, `9000`, `12000`, `18000`, `24000`, `36000`, `48000`,  and `54000`.
        /// </summary>
        [Input("minimumDataRate5gKbps")]
        public Input<int>? MinimumDataRate5gKbps { get; set; }

        /// <summary>
        /// Indicates whether or not Multicast Enhance is turned of for the network.
        /// </summary>
        [Input("multicastEnhance")]
        public Input<bool>? MulticastEnhance { get; set; }

        /// <summary>
        /// The SSID of the network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the network for this SSID
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// Connect high performance clients to 5 GHz only Defaults to `true`.
        /// </summary>
        [Input("no2ghzOui")]
        public Input<bool>? No2ghzOui { get; set; }

        /// <summary>
        /// The passphrase for the network, this is only required if `security` is not set to `open`.
        /// </summary>
        [Input("passphrase")]
        public Input<string>? Passphrase { get; set; }

        /// <summary>
        /// Enable Protected Management Frames. This cannot be disabled if using WPA 3. Valid values are `required`, `optional` and `disabled`. Defaults to `disabled`.
        /// </summary>
        [Input("pmfMode")]
        public Input<string>? PmfMode { get; set; }

        /// <summary>
        /// ID of the RADIUS profile to use when security `wpaeap`. You can query this via the `unifi.getRadiusProfile` data source.
        /// </summary>
        [Input("radiusProfileId")]
        public Input<string>? RadiusProfileId { get; set; }

        [Input("schedules")]
        private InputList<Inputs.WlanScheduleGetArgs>? _schedules;

        /// <summary>
        /// Start and stop schedules for the WLAN
        /// </summary>
        public InputList<Inputs.WlanScheduleGetArgs> Schedules
        {
            get => _schedules ?? (_schedules = new InputList<Inputs.WlanScheduleGetArgs>());
            set => _schedules = value;
        }

        /// <summary>
        /// The type of WiFi security for this network. Valid values are: `wpapsk`, `wpaeap`, and `open`.
        /// </summary>
        [Input("security")]
        public Input<string>? Security { get; set; }

        /// <summary>
        /// The name of the site to associate the wlan with.
        /// </summary>
        [Input("site")]
        public Input<string>? Site { get; set; }

        /// <summary>
        /// Enable Unscheduled Automatic Power Save Delivery Defaults to `false`.
        /// </summary>
        [Input("uapsd")]
        public Input<bool>? Uapsd { get; set; }

        /// <summary>
        /// ID of the user group to use for this network.
        /// </summary>
        [Input("userGroupId")]
        public Input<string>? UserGroupId { get; set; }

        /// <summary>
        /// Radio band your WiFi network will use.
        /// </summary>
        [Input("wlanBand")]
        public Input<string>? WlanBand { get; set; }

        /// <summary>
        /// Enable WPA 3 support (security must be `wpapsk` and PMF must be turned on).
        /// </summary>
        [Input("wpa3Support")]
        public Input<bool>? Wpa3Support { get; set; }

        /// <summary>
        /// Enable WPA 3 and WPA 2 support (security must be `wpapsk` and `wpa3_support` must be true).
        /// </summary>
        [Input("wpa3Transition")]
        public Input<bool>? Wpa3Transition { get; set; }

        public WlanState()
        {
        }
    }
}
