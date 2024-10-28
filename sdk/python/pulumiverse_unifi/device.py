# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DeviceArgs', 'Device']

@pulumi.input_type
class DeviceArgs:
    def __init__(__self__, *,
                 allow_adoption: Optional[pulumi.Input[bool]] = None,
                 forget_on_destroy: Optional[pulumi.Input[bool]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port_overrides: Optional[pulumi.Input[Sequence[pulumi.Input['DevicePortOverrideArgs']]]] = None,
                 site: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Device resource.
        :param pulumi.Input[bool] allow_adoption: Specifies whether this resource should tell the controller to adopt the device on create. Defaults to `true`.
        :param pulumi.Input[bool] forget_on_destroy: Specifies whether this resource should tell the controller to forget the device on destroy. Defaults to `true`.
        :param pulumi.Input[str] mac: The MAC address of the device. This can be specified so that the provider can take control of a device (since devices are created through adoption).
        :param pulumi.Input[str] name: The name of the device.
        :param pulumi.Input[Sequence[pulumi.Input['DevicePortOverrideArgs']]] port_overrides: Settings overrides for specific switch ports.
        :param pulumi.Input[str] site: The name of the site to associate the device with.
        """
        if allow_adoption is not None:
            pulumi.set(__self__, "allow_adoption", allow_adoption)
        if forget_on_destroy is not None:
            pulumi.set(__self__, "forget_on_destroy", forget_on_destroy)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port_overrides is not None:
            pulumi.set(__self__, "port_overrides", port_overrides)
        if site is not None:
            pulumi.set(__self__, "site", site)

    @property
    @pulumi.getter(name="allowAdoption")
    def allow_adoption(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether this resource should tell the controller to adopt the device on create. Defaults to `true`.
        """
        return pulumi.get(self, "allow_adoption")

    @allow_adoption.setter
    def allow_adoption(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_adoption", value)

    @property
    @pulumi.getter(name="forgetOnDestroy")
    def forget_on_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether this resource should tell the controller to forget the device on destroy. Defaults to `true`.
        """
        return pulumi.get(self, "forget_on_destroy")

    @forget_on_destroy.setter
    def forget_on_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "forget_on_destroy", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        The MAC address of the device. This can be specified so that the provider can take control of a device (since devices are created through adoption).
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the device.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="portOverrides")
    def port_overrides(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DevicePortOverrideArgs']]]]:
        """
        Settings overrides for specific switch ports.
        """
        return pulumi.get(self, "port_overrides")

    @port_overrides.setter
    def port_overrides(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DevicePortOverrideArgs']]]]):
        pulumi.set(self, "port_overrides", value)

    @property
    @pulumi.getter
    def site(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the site to associate the device with.
        """
        return pulumi.get(self, "site")

    @site.setter
    def site(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "site", value)


@pulumi.input_type
class _DeviceState:
    def __init__(__self__, *,
                 allow_adoption: Optional[pulumi.Input[bool]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 forget_on_destroy: Optional[pulumi.Input[bool]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port_overrides: Optional[pulumi.Input[Sequence[pulumi.Input['DevicePortOverrideArgs']]]] = None,
                 site: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Device resources.
        :param pulumi.Input[bool] allow_adoption: Specifies whether this resource should tell the controller to adopt the device on create. Defaults to `true`.
        :param pulumi.Input[bool] disabled: Specifies whether this device should be disabled.
        :param pulumi.Input[bool] forget_on_destroy: Specifies whether this resource should tell the controller to forget the device on destroy. Defaults to `true`.
        :param pulumi.Input[str] mac: The MAC address of the device. This can be specified so that the provider can take control of a device (since devices are created through adoption).
        :param pulumi.Input[str] name: The name of the device.
        :param pulumi.Input[Sequence[pulumi.Input['DevicePortOverrideArgs']]] port_overrides: Settings overrides for specific switch ports.
        :param pulumi.Input[str] site: The name of the site to associate the device with.
        """
        if allow_adoption is not None:
            pulumi.set(__self__, "allow_adoption", allow_adoption)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if forget_on_destroy is not None:
            pulumi.set(__self__, "forget_on_destroy", forget_on_destroy)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port_overrides is not None:
            pulumi.set(__self__, "port_overrides", port_overrides)
        if site is not None:
            pulumi.set(__self__, "site", site)

    @property
    @pulumi.getter(name="allowAdoption")
    def allow_adoption(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether this resource should tell the controller to adopt the device on create. Defaults to `true`.
        """
        return pulumi.get(self, "allow_adoption")

    @allow_adoption.setter
    def allow_adoption(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_adoption", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether this device should be disabled.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="forgetOnDestroy")
    def forget_on_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether this resource should tell the controller to forget the device on destroy. Defaults to `true`.
        """
        return pulumi.get(self, "forget_on_destroy")

    @forget_on_destroy.setter
    def forget_on_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "forget_on_destroy", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        The MAC address of the device. This can be specified so that the provider can take control of a device (since devices are created through adoption).
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the device.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="portOverrides")
    def port_overrides(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DevicePortOverrideArgs']]]]:
        """
        Settings overrides for specific switch ports.
        """
        return pulumi.get(self, "port_overrides")

    @port_overrides.setter
    def port_overrides(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DevicePortOverrideArgs']]]]):
        pulumi.set(self, "port_overrides", value)

    @property
    @pulumi.getter
    def site(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the site to associate the device with.
        """
        return pulumi.get(self, "site")

    @site.setter
    def site(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "site", value)


class Device(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_adoption: Optional[pulumi.Input[bool]] = None,
                 forget_on_destroy: Optional[pulumi.Input[bool]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DevicePortOverrideArgs', 'DevicePortOverrideArgsDict']]]]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_unifi as unifi
        import pulumiverse_unifi as unifi

        disabled = unifi.port.get_profile(name="Disabled")
        poe = unifi.port.Profile("poe",
            name="poe",
            forward="customize",
            native_networkconf_id=native_network_id,
            tagged_networkconf_ids=[some_vlan_network_id],
            poe_mode="auto")
        us24_poe = unifi.Device("us_24_poe",
            mac="01:23:45:67:89:AB",
            name="Switch with POE",
            port_overrides=[
                {
                    "number": 1,
                    "name": "port w/ poe",
                    "port_profile_id": poe.id,
                },
                {
                    "number": 2,
                    "name": "disabled",
                    "port_profile_id": disabled.id,
                },
                {
                    "number": 11,
                    "op_mode": "aggregate",
                    "aggregate_num_ports": 2,
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_adoption: Specifies whether this resource should tell the controller to adopt the device on create. Defaults to `true`.
        :param pulumi.Input[bool] forget_on_destroy: Specifies whether this resource should tell the controller to forget the device on destroy. Defaults to `true`.
        :param pulumi.Input[str] mac: The MAC address of the device. This can be specified so that the provider can take control of a device (since devices are created through adoption).
        :param pulumi.Input[str] name: The name of the device.
        :param pulumi.Input[Sequence[pulumi.Input[Union['DevicePortOverrideArgs', 'DevicePortOverrideArgsDict']]]] port_overrides: Settings overrides for specific switch ports.
        :param pulumi.Input[str] site: The name of the site to associate the device with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DeviceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_unifi as unifi
        import pulumiverse_unifi as unifi

        disabled = unifi.port.get_profile(name="Disabled")
        poe = unifi.port.Profile("poe",
            name="poe",
            forward="customize",
            native_networkconf_id=native_network_id,
            tagged_networkconf_ids=[some_vlan_network_id],
            poe_mode="auto")
        us24_poe = unifi.Device("us_24_poe",
            mac="01:23:45:67:89:AB",
            name="Switch with POE",
            port_overrides=[
                {
                    "number": 1,
                    "name": "port w/ poe",
                    "port_profile_id": poe.id,
                },
                {
                    "number": 2,
                    "name": "disabled",
                    "port_profile_id": disabled.id,
                },
                {
                    "number": 11,
                    "op_mode": "aggregate",
                    "aggregate_num_ports": 2,
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param DeviceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_adoption: Optional[pulumi.Input[bool]] = None,
                 forget_on_destroy: Optional[pulumi.Input[bool]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DevicePortOverrideArgs', 'DevicePortOverrideArgsDict']]]]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceArgs.__new__(DeviceArgs)

            __props__.__dict__["allow_adoption"] = allow_adoption
            __props__.__dict__["forget_on_destroy"] = forget_on_destroy
            __props__.__dict__["mac"] = mac
            __props__.__dict__["name"] = name
            __props__.__dict__["port_overrides"] = port_overrides
            __props__.__dict__["site"] = site
            __props__.__dict__["disabled"] = None
        super(Device, __self__).__init__(
            'unifi:index/device:Device',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_adoption: Optional[pulumi.Input[bool]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            forget_on_destroy: Optional[pulumi.Input[bool]] = None,
            mac: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DevicePortOverrideArgs', 'DevicePortOverrideArgsDict']]]]] = None,
            site: Optional[pulumi.Input[str]] = None) -> 'Device':
        """
        Get an existing Device resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_adoption: Specifies whether this resource should tell the controller to adopt the device on create. Defaults to `true`.
        :param pulumi.Input[bool] disabled: Specifies whether this device should be disabled.
        :param pulumi.Input[bool] forget_on_destroy: Specifies whether this resource should tell the controller to forget the device on destroy. Defaults to `true`.
        :param pulumi.Input[str] mac: The MAC address of the device. This can be specified so that the provider can take control of a device (since devices are created through adoption).
        :param pulumi.Input[str] name: The name of the device.
        :param pulumi.Input[Sequence[pulumi.Input[Union['DevicePortOverrideArgs', 'DevicePortOverrideArgsDict']]]] port_overrides: Settings overrides for specific switch ports.
        :param pulumi.Input[str] site: The name of the site to associate the device with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeviceState.__new__(_DeviceState)

        __props__.__dict__["allow_adoption"] = allow_adoption
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["forget_on_destroy"] = forget_on_destroy
        __props__.__dict__["mac"] = mac
        __props__.__dict__["name"] = name
        __props__.__dict__["port_overrides"] = port_overrides
        __props__.__dict__["site"] = site
        return Device(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowAdoption")
    def allow_adoption(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether this resource should tell the controller to adopt the device on create. Defaults to `true`.
        """
        return pulumi.get(self, "allow_adoption")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Specifies whether this device should be disabled.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="forgetOnDestroy")
    def forget_on_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether this resource should tell the controller to forget the device on destroy. Defaults to `true`.
        """
        return pulumi.get(self, "forget_on_destroy")

    @property
    @pulumi.getter
    def mac(self) -> pulumi.Output[str]:
        """
        The MAC address of the device. This can be specified so that the provider can take control of a device (since devices are created through adoption).
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the device.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="portOverrides")
    def port_overrides(self) -> pulumi.Output[Optional[Sequence['outputs.DevicePortOverride']]]:
        """
        Settings overrides for specific switch ports.
        """
        return pulumi.get(self, "port_overrides")

    @property
    @pulumi.getter
    def site(self) -> pulumi.Output[str]:
        """
        The name of the site to associate the device with.
        """
        return pulumi.get(self, "site")

