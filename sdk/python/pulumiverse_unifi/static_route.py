# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['StaticRouteArgs', 'StaticRoute']

@pulumi.input_type
class StaticRouteArgs:
    def __init__(__self__, *,
                 distance: pulumi.Input[int],
                 network: pulumi.Input[str],
                 type: pulumi.Input[str],
                 interface: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 next_hop: Optional[pulumi.Input[str]] = None,
                 site: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a StaticRoute resource.
        :param pulumi.Input[int] distance: The distance of the static route.
        :param pulumi.Input[str] network: The network subnet address.
        :param pulumi.Input[str] type: The type of static route. Can be `interface-route`, `nexthop-route`, or `blackhole`.
        :param pulumi.Input[str] interface: The interface of the static route (only valid for `interface-route` type). This can be `WAN1`, `WAN2`, or a network ID.
        :param pulumi.Input[str] name: The name of the static route.
        :param pulumi.Input[str] next_hop: The next hop of the static route (only valid for `nexthop-route` type).
        :param pulumi.Input[str] site: The name of the site to associate the static route with.
        """
        pulumi.set(__self__, "distance", distance)
        pulumi.set(__self__, "network", network)
        pulumi.set(__self__, "type", type)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if next_hop is not None:
            pulumi.set(__self__, "next_hop", next_hop)
        if site is not None:
            pulumi.set(__self__, "site", site)

    @property
    @pulumi.getter
    def distance(self) -> pulumi.Input[int]:
        """
        The distance of the static route.
        """
        return pulumi.get(self, "distance")

    @distance.setter
    def distance(self, value: pulumi.Input[int]):
        pulumi.set(self, "distance", value)

    @property
    @pulumi.getter
    def network(self) -> pulumi.Input[str]:
        """
        The network subnet address.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: pulumi.Input[str]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of static route. Can be `interface-route`, `nexthop-route`, or `blackhole`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        The interface of the static route (only valid for `interface-route` type). This can be `WAN1`, `WAN2`, or a network ID.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the static route.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nextHop")
    def next_hop(self) -> Optional[pulumi.Input[str]]:
        """
        The next hop of the static route (only valid for `nexthop-route` type).
        """
        return pulumi.get(self, "next_hop")

    @next_hop.setter
    def next_hop(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "next_hop", value)

    @property
    @pulumi.getter
    def site(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the site to associate the static route with.
        """
        return pulumi.get(self, "site")

    @site.setter
    def site(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "site", value)


@pulumi.input_type
class _StaticRouteState:
    def __init__(__self__, *,
                 distance: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 next_hop: Optional[pulumi.Input[str]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StaticRoute resources.
        :param pulumi.Input[int] distance: The distance of the static route.
        :param pulumi.Input[str] interface: The interface of the static route (only valid for `interface-route` type). This can be `WAN1`, `WAN2`, or a network ID.
        :param pulumi.Input[str] name: The name of the static route.
        :param pulumi.Input[str] network: The network subnet address.
        :param pulumi.Input[str] next_hop: The next hop of the static route (only valid for `nexthop-route` type).
        :param pulumi.Input[str] site: The name of the site to associate the static route with.
        :param pulumi.Input[str] type: The type of static route. Can be `interface-route`, `nexthop-route`, or `blackhole`.
        """
        if distance is not None:
            pulumi.set(__self__, "distance", distance)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if next_hop is not None:
            pulumi.set(__self__, "next_hop", next_hop)
        if site is not None:
            pulumi.set(__self__, "site", site)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def distance(self) -> Optional[pulumi.Input[int]]:
        """
        The distance of the static route.
        """
        return pulumi.get(self, "distance")

    @distance.setter
    def distance(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "distance", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        The interface of the static route (only valid for `interface-route` type). This can be `WAN1`, `WAN2`, or a network ID.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the static route.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        The network subnet address.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="nextHop")
    def next_hop(self) -> Optional[pulumi.Input[str]]:
        """
        The next hop of the static route (only valid for `nexthop-route` type).
        """
        return pulumi.get(self, "next_hop")

    @next_hop.setter
    def next_hop(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "next_hop", value)

    @property
    @pulumi.getter
    def site(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the site to associate the static route with.
        """
        return pulumi.get(self, "site")

    @site.setter
    def site(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "site", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of static route. Can be `interface-route`, `nexthop-route`, or `blackhole`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class StaticRoute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 next_hop: Optional[pulumi.Input[str]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `StaticRoute` manages a static route.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_unifi as unifi

        nexthop = unifi.StaticRoute("nexthop",
            type="nexthop-route",
            network="172.17.0.0/16",
            name="basic nexthop",
            distance=1,
            next_hop="172.16.0.1")
        blackhole = unifi.StaticRoute("blackhole",
            type="blackhole",
            network=blackhole_cidr,
            name="blackhole traffice to cidr",
            distance=1)
        interface = unifi.StaticRoute("interface",
            type="interface-route",
            network=wan2_cidr,
            name="send traffic over wan2",
            distance=1,
            interface="WAN2")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] distance: The distance of the static route.
        :param pulumi.Input[str] interface: The interface of the static route (only valid for `interface-route` type). This can be `WAN1`, `WAN2`, or a network ID.
        :param pulumi.Input[str] name: The name of the static route.
        :param pulumi.Input[str] network: The network subnet address.
        :param pulumi.Input[str] next_hop: The next hop of the static route (only valid for `nexthop-route` type).
        :param pulumi.Input[str] site: The name of the site to associate the static route with.
        :param pulumi.Input[str] type: The type of static route. Can be `interface-route`, `nexthop-route`, or `blackhole`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StaticRouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `StaticRoute` manages a static route.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_unifi as unifi

        nexthop = unifi.StaticRoute("nexthop",
            type="nexthop-route",
            network="172.17.0.0/16",
            name="basic nexthop",
            distance=1,
            next_hop="172.16.0.1")
        blackhole = unifi.StaticRoute("blackhole",
            type="blackhole",
            network=blackhole_cidr,
            name="blackhole traffice to cidr",
            distance=1)
        interface = unifi.StaticRoute("interface",
            type="interface-route",
            network=wan2_cidr,
            name="send traffic over wan2",
            distance=1,
            interface="WAN2")
        ```

        :param str resource_name: The name of the resource.
        :param StaticRouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StaticRouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 distance: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 next_hop: Optional[pulumi.Input[str]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StaticRouteArgs.__new__(StaticRouteArgs)

            if distance is None and not opts.urn:
                raise TypeError("Missing required property 'distance'")
            __props__.__dict__["distance"] = distance
            __props__.__dict__["interface"] = interface
            __props__.__dict__["name"] = name
            if network is None and not opts.urn:
                raise TypeError("Missing required property 'network'")
            __props__.__dict__["network"] = network
            __props__.__dict__["next_hop"] = next_hop
            __props__.__dict__["site"] = site
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(StaticRoute, __self__).__init__(
            'unifi:index/staticRoute:StaticRoute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            distance: Optional[pulumi.Input[int]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network: Optional[pulumi.Input[str]] = None,
            next_hop: Optional[pulumi.Input[str]] = None,
            site: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'StaticRoute':
        """
        Get an existing StaticRoute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] distance: The distance of the static route.
        :param pulumi.Input[str] interface: The interface of the static route (only valid for `interface-route` type). This can be `WAN1`, `WAN2`, or a network ID.
        :param pulumi.Input[str] name: The name of the static route.
        :param pulumi.Input[str] network: The network subnet address.
        :param pulumi.Input[str] next_hop: The next hop of the static route (only valid for `nexthop-route` type).
        :param pulumi.Input[str] site: The name of the site to associate the static route with.
        :param pulumi.Input[str] type: The type of static route. Can be `interface-route`, `nexthop-route`, or `blackhole`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StaticRouteState.__new__(_StaticRouteState)

        __props__.__dict__["distance"] = distance
        __props__.__dict__["interface"] = interface
        __props__.__dict__["name"] = name
        __props__.__dict__["network"] = network
        __props__.__dict__["next_hop"] = next_hop
        __props__.__dict__["site"] = site
        __props__.__dict__["type"] = type
        return StaticRoute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def distance(self) -> pulumi.Output[int]:
        """
        The distance of the static route.
        """
        return pulumi.get(self, "distance")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[Optional[str]]:
        """
        The interface of the static route (only valid for `interface-route` type). This can be `WAN1`, `WAN2`, or a network ID.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the static route.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[str]:
        """
        The network subnet address.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="nextHop")
    def next_hop(self) -> pulumi.Output[Optional[str]]:
        """
        The next hop of the static route (only valid for `nexthop-route` type).
        """
        return pulumi.get(self, "next_hop")

    @property
    @pulumi.getter
    def site(self) -> pulumi.Output[str]:
        """
        The name of the site to associate the static route with.
        """
        return pulumi.get(self, "site")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of static route. Can be `interface-route`, `nexthop-route`, or `blackhole`.
        """
        return pulumi.get(self, "type")

