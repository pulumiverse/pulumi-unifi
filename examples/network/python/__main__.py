"""A Python Pulumi program"""

import pulumi
from pulumiverse_unifi import network

config = pulumi.Config()
vlan_id = config.get_int("vlanId", 10)

vlan = network.Network(
    "vlan",
    purpose="corporate",
    subnet="10.0.0.1/24",
    vlan_id=vlan_id,
    dhcp_start="10.0.0.6",
    dhcp_stop="10.0.0.254",
    dhcp_enabled=True,

    # Adding "wanGateway" and ignoring changes to it is required due to serialisation issue.
    # See https://github.com/paultyng/terraform-provider-unifi/issues/196
    wan_gateway="0.0.0.0",
    opts=pulumi.ResourceOptions(ignore_changes=["wan_gateway"])
)

wan = network.Network(
    "wan",
    purpose="wan",
    wan_networkgroup="WAN",
    wan_type="pppoe",
    wan_ip="192.168.1.1",
    wan_egress_qos=1,
    wan_username="username",
    x_wan_password="password",
    wan_gateway="0.0.0.0"
)
