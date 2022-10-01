import * as pulumi from "@pulumi/pulumi";
import * as unifi from "@pulumiverse/unifi";

const config = new pulumi.Config();
const vlanId = config.getNumber("vlanId") || 10;

const vlan = new unifi.Network(
  "vlan",
  {
    purpose: "corporate",
    subnet: "10.0.0.1/24",
    vlanId: vlanId,
    dhcpStart: "10.0.0.6",
    dhcpStop: "10.0.0.254",
    dhcpEnabled: true,

    // Adding "wanGateway" and ignoring changes to it is required due to serialisation issue.
    // See https://github.com/paultyng/terraform-provider-unifi/issues/196
    wanGateway: "0.0.0.0",
  },
  { ignoreChanges: ["wanGateway"] }
);

const wan = new unifi.Network("wan", {
  purpose: "wan",
  wanNetworkgroup: "WAN",
  wanType: "pppoe",
  wanIp: "192.168.1.1",
  wanEgressQos: 1,
  wanUsername: "username",
  xWanPassword: "password",

  wanGateway: "0.0.0.0",
});
