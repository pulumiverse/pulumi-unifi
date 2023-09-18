# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RadiusProfileArgs', 'RadiusProfile']

@pulumi.input_type
class RadiusProfileArgs:
    def __init__(__self__, *,
                 accounting_enabled: Optional[pulumi.Input[bool]] = None,
                 acct_servers: Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAcctServerArgs']]]] = None,
                 auth_servers: Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAuthServerArgs']]]] = None,
                 interim_update_enabled: Optional[pulumi.Input[bool]] = None,
                 interim_update_interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 use_usg_acct_server: Optional[pulumi.Input[bool]] = None,
                 use_usg_auth_server: Optional[pulumi.Input[bool]] = None,
                 vlan_enabled: Optional[pulumi.Input[bool]] = None,
                 vlan_wlan_mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RadiusProfile resource.
        :param pulumi.Input[bool] accounting_enabled: Specifies whether to use RADIUS accounting. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input['RadiusProfileAcctServerArgs']]] acct_servers: RADIUS accounting servers.
        :param pulumi.Input[Sequence[pulumi.Input['RadiusProfileAuthServerArgs']]] auth_servers: RADIUS authentication servers.
        :param pulumi.Input[bool] interim_update_enabled: Specifies whether to use interim_update. Defaults to `false`.
        :param pulumi.Input[int] interim_update_interval: Specifies interim_update interval. Defaults to `3600`.
        :param pulumi.Input[str] name: The name of the profile.
        :param pulumi.Input[str] site: The name of the site to associate the settings with.
        :param pulumi.Input[bool] use_usg_acct_server: Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
        :param pulumi.Input[bool] use_usg_auth_server: Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
        :param pulumi.Input[bool] vlan_enabled: Specifies whether to use vlan on wired connections. Defaults to `false`.
        :param pulumi.Input[str] vlan_wlan_mode: Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
        """
        RadiusProfileArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            accounting_enabled=accounting_enabled,
            acct_servers=acct_servers,
            auth_servers=auth_servers,
            interim_update_enabled=interim_update_enabled,
            interim_update_interval=interim_update_interval,
            name=name,
            site=site,
            use_usg_acct_server=use_usg_acct_server,
            use_usg_auth_server=use_usg_auth_server,
            vlan_enabled=vlan_enabled,
            vlan_wlan_mode=vlan_wlan_mode,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             accounting_enabled: Optional[pulumi.Input[bool]] = None,
             acct_servers: Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAcctServerArgs']]]] = None,
             auth_servers: Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAuthServerArgs']]]] = None,
             interim_update_enabled: Optional[pulumi.Input[bool]] = None,
             interim_update_interval: Optional[pulumi.Input[int]] = None,
             name: Optional[pulumi.Input[str]] = None,
             site: Optional[pulumi.Input[str]] = None,
             use_usg_acct_server: Optional[pulumi.Input[bool]] = None,
             use_usg_auth_server: Optional[pulumi.Input[bool]] = None,
             vlan_enabled: Optional[pulumi.Input[bool]] = None,
             vlan_wlan_mode: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if accounting_enabled is not None:
            _setter("accounting_enabled", accounting_enabled)
        if acct_servers is not None:
            _setter("acct_servers", acct_servers)
        if auth_servers is not None:
            _setter("auth_servers", auth_servers)
        if interim_update_enabled is not None:
            _setter("interim_update_enabled", interim_update_enabled)
        if interim_update_interval is not None:
            _setter("interim_update_interval", interim_update_interval)
        if name is not None:
            _setter("name", name)
        if site is not None:
            _setter("site", site)
        if use_usg_acct_server is not None:
            _setter("use_usg_acct_server", use_usg_acct_server)
        if use_usg_auth_server is not None:
            _setter("use_usg_auth_server", use_usg_auth_server)
        if vlan_enabled is not None:
            _setter("vlan_enabled", vlan_enabled)
        if vlan_wlan_mode is not None:
            _setter("vlan_wlan_mode", vlan_wlan_mode)

    @property
    @pulumi.getter(name="accountingEnabled")
    def accounting_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to use RADIUS accounting. Defaults to `false`.
        """
        return pulumi.get(self, "accounting_enabled")

    @accounting_enabled.setter
    def accounting_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "accounting_enabled", value)

    @property
    @pulumi.getter(name="acctServers")
    def acct_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAcctServerArgs']]]]:
        """
        RADIUS accounting servers.
        """
        return pulumi.get(self, "acct_servers")

    @acct_servers.setter
    def acct_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAcctServerArgs']]]]):
        pulumi.set(self, "acct_servers", value)

    @property
    @pulumi.getter(name="authServers")
    def auth_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAuthServerArgs']]]]:
        """
        RADIUS authentication servers.
        """
        return pulumi.get(self, "auth_servers")

    @auth_servers.setter
    def auth_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAuthServerArgs']]]]):
        pulumi.set(self, "auth_servers", value)

    @property
    @pulumi.getter(name="interimUpdateEnabled")
    def interim_update_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to use interim_update. Defaults to `false`.
        """
        return pulumi.get(self, "interim_update_enabled")

    @interim_update_enabled.setter
    def interim_update_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "interim_update_enabled", value)

    @property
    @pulumi.getter(name="interimUpdateInterval")
    def interim_update_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies interim_update interval. Defaults to `3600`.
        """
        return pulumi.get(self, "interim_update_interval")

    @interim_update_interval.setter
    def interim_update_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interim_update_interval", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the profile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def site(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the site to associate the settings with.
        """
        return pulumi.get(self, "site")

    @site.setter
    def site(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "site", value)

    @property
    @pulumi.getter(name="useUsgAcctServer")
    def use_usg_acct_server(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
        """
        return pulumi.get(self, "use_usg_acct_server")

    @use_usg_acct_server.setter
    def use_usg_acct_server(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_usg_acct_server", value)

    @property
    @pulumi.getter(name="useUsgAuthServer")
    def use_usg_auth_server(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
        """
        return pulumi.get(self, "use_usg_auth_server")

    @use_usg_auth_server.setter
    def use_usg_auth_server(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_usg_auth_server", value)

    @property
    @pulumi.getter(name="vlanEnabled")
    def vlan_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to use vlan on wired connections. Defaults to `false`.
        """
        return pulumi.get(self, "vlan_enabled")

    @vlan_enabled.setter
    def vlan_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "vlan_enabled", value)

    @property
    @pulumi.getter(name="vlanWlanMode")
    def vlan_wlan_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
        """
        return pulumi.get(self, "vlan_wlan_mode")

    @vlan_wlan_mode.setter
    def vlan_wlan_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vlan_wlan_mode", value)


@pulumi.input_type
class _RadiusProfileState:
    def __init__(__self__, *,
                 accounting_enabled: Optional[pulumi.Input[bool]] = None,
                 acct_servers: Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAcctServerArgs']]]] = None,
                 auth_servers: Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAuthServerArgs']]]] = None,
                 interim_update_enabled: Optional[pulumi.Input[bool]] = None,
                 interim_update_interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 use_usg_acct_server: Optional[pulumi.Input[bool]] = None,
                 use_usg_auth_server: Optional[pulumi.Input[bool]] = None,
                 vlan_enabled: Optional[pulumi.Input[bool]] = None,
                 vlan_wlan_mode: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RadiusProfile resources.
        :param pulumi.Input[bool] accounting_enabled: Specifies whether to use RADIUS accounting. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input['RadiusProfileAcctServerArgs']]] acct_servers: RADIUS accounting servers.
        :param pulumi.Input[Sequence[pulumi.Input['RadiusProfileAuthServerArgs']]] auth_servers: RADIUS authentication servers.
        :param pulumi.Input[bool] interim_update_enabled: Specifies whether to use interim_update. Defaults to `false`.
        :param pulumi.Input[int] interim_update_interval: Specifies interim_update interval. Defaults to `3600`.
        :param pulumi.Input[str] name: The name of the profile.
        :param pulumi.Input[str] site: The name of the site to associate the settings with.
        :param pulumi.Input[bool] use_usg_acct_server: Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
        :param pulumi.Input[bool] use_usg_auth_server: Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
        :param pulumi.Input[bool] vlan_enabled: Specifies whether to use vlan on wired connections. Defaults to `false`.
        :param pulumi.Input[str] vlan_wlan_mode: Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
        """
        _RadiusProfileState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            accounting_enabled=accounting_enabled,
            acct_servers=acct_servers,
            auth_servers=auth_servers,
            interim_update_enabled=interim_update_enabled,
            interim_update_interval=interim_update_interval,
            name=name,
            site=site,
            use_usg_acct_server=use_usg_acct_server,
            use_usg_auth_server=use_usg_auth_server,
            vlan_enabled=vlan_enabled,
            vlan_wlan_mode=vlan_wlan_mode,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             accounting_enabled: Optional[pulumi.Input[bool]] = None,
             acct_servers: Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAcctServerArgs']]]] = None,
             auth_servers: Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAuthServerArgs']]]] = None,
             interim_update_enabled: Optional[pulumi.Input[bool]] = None,
             interim_update_interval: Optional[pulumi.Input[int]] = None,
             name: Optional[pulumi.Input[str]] = None,
             site: Optional[pulumi.Input[str]] = None,
             use_usg_acct_server: Optional[pulumi.Input[bool]] = None,
             use_usg_auth_server: Optional[pulumi.Input[bool]] = None,
             vlan_enabled: Optional[pulumi.Input[bool]] = None,
             vlan_wlan_mode: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if accounting_enabled is not None:
            _setter("accounting_enabled", accounting_enabled)
        if acct_servers is not None:
            _setter("acct_servers", acct_servers)
        if auth_servers is not None:
            _setter("auth_servers", auth_servers)
        if interim_update_enabled is not None:
            _setter("interim_update_enabled", interim_update_enabled)
        if interim_update_interval is not None:
            _setter("interim_update_interval", interim_update_interval)
        if name is not None:
            _setter("name", name)
        if site is not None:
            _setter("site", site)
        if use_usg_acct_server is not None:
            _setter("use_usg_acct_server", use_usg_acct_server)
        if use_usg_auth_server is not None:
            _setter("use_usg_auth_server", use_usg_auth_server)
        if vlan_enabled is not None:
            _setter("vlan_enabled", vlan_enabled)
        if vlan_wlan_mode is not None:
            _setter("vlan_wlan_mode", vlan_wlan_mode)

    @property
    @pulumi.getter(name="accountingEnabled")
    def accounting_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to use RADIUS accounting. Defaults to `false`.
        """
        return pulumi.get(self, "accounting_enabled")

    @accounting_enabled.setter
    def accounting_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "accounting_enabled", value)

    @property
    @pulumi.getter(name="acctServers")
    def acct_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAcctServerArgs']]]]:
        """
        RADIUS accounting servers.
        """
        return pulumi.get(self, "acct_servers")

    @acct_servers.setter
    def acct_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAcctServerArgs']]]]):
        pulumi.set(self, "acct_servers", value)

    @property
    @pulumi.getter(name="authServers")
    def auth_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAuthServerArgs']]]]:
        """
        RADIUS authentication servers.
        """
        return pulumi.get(self, "auth_servers")

    @auth_servers.setter
    def auth_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RadiusProfileAuthServerArgs']]]]):
        pulumi.set(self, "auth_servers", value)

    @property
    @pulumi.getter(name="interimUpdateEnabled")
    def interim_update_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to use interim_update. Defaults to `false`.
        """
        return pulumi.get(self, "interim_update_enabled")

    @interim_update_enabled.setter
    def interim_update_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "interim_update_enabled", value)

    @property
    @pulumi.getter(name="interimUpdateInterval")
    def interim_update_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies interim_update interval. Defaults to `3600`.
        """
        return pulumi.get(self, "interim_update_interval")

    @interim_update_interval.setter
    def interim_update_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interim_update_interval", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the profile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def site(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the site to associate the settings with.
        """
        return pulumi.get(self, "site")

    @site.setter
    def site(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "site", value)

    @property
    @pulumi.getter(name="useUsgAcctServer")
    def use_usg_acct_server(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
        """
        return pulumi.get(self, "use_usg_acct_server")

    @use_usg_acct_server.setter
    def use_usg_acct_server(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_usg_acct_server", value)

    @property
    @pulumi.getter(name="useUsgAuthServer")
    def use_usg_auth_server(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
        """
        return pulumi.get(self, "use_usg_auth_server")

    @use_usg_auth_server.setter
    def use_usg_auth_server(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_usg_auth_server", value)

    @property
    @pulumi.getter(name="vlanEnabled")
    def vlan_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to use vlan on wired connections. Defaults to `false`.
        """
        return pulumi.get(self, "vlan_enabled")

    @vlan_enabled.setter
    def vlan_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "vlan_enabled", value)

    @property
    @pulumi.getter(name="vlanWlanMode")
    def vlan_wlan_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
        """
        return pulumi.get(self, "vlan_wlan_mode")

    @vlan_wlan_mode.setter
    def vlan_wlan_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vlan_wlan_mode", value)


class RadiusProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accounting_enabled: Optional[pulumi.Input[bool]] = None,
                 acct_servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RadiusProfileAcctServerArgs']]]]] = None,
                 auth_servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RadiusProfileAuthServerArgs']]]]] = None,
                 interim_update_enabled: Optional[pulumi.Input[bool]] = None,
                 interim_update_interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 use_usg_acct_server: Optional[pulumi.Input[bool]] = None,
                 use_usg_auth_server: Optional[pulumi.Input[bool]] = None,
                 vlan_enabled: Optional[pulumi.Input[bool]] = None,
                 vlan_wlan_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `RadiusProfile` manages RADIUS profiles.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] accounting_enabled: Specifies whether to use RADIUS accounting. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RadiusProfileAcctServerArgs']]]] acct_servers: RADIUS accounting servers.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RadiusProfileAuthServerArgs']]]] auth_servers: RADIUS authentication servers.
        :param pulumi.Input[bool] interim_update_enabled: Specifies whether to use interim_update. Defaults to `false`.
        :param pulumi.Input[int] interim_update_interval: Specifies interim_update interval. Defaults to `3600`.
        :param pulumi.Input[str] name: The name of the profile.
        :param pulumi.Input[str] site: The name of the site to associate the settings with.
        :param pulumi.Input[bool] use_usg_acct_server: Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
        :param pulumi.Input[bool] use_usg_auth_server: Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
        :param pulumi.Input[bool] vlan_enabled: Specifies whether to use vlan on wired connections. Defaults to `false`.
        :param pulumi.Input[str] vlan_wlan_mode: Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RadiusProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `RadiusProfile` manages RADIUS profiles.

        :param str resource_name: The name of the resource.
        :param RadiusProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RadiusProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            RadiusProfileArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accounting_enabled: Optional[pulumi.Input[bool]] = None,
                 acct_servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RadiusProfileAcctServerArgs']]]]] = None,
                 auth_servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RadiusProfileAuthServerArgs']]]]] = None,
                 interim_update_enabled: Optional[pulumi.Input[bool]] = None,
                 interim_update_interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 use_usg_acct_server: Optional[pulumi.Input[bool]] = None,
                 use_usg_auth_server: Optional[pulumi.Input[bool]] = None,
                 vlan_enabled: Optional[pulumi.Input[bool]] = None,
                 vlan_wlan_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RadiusProfileArgs.__new__(RadiusProfileArgs)

            __props__.__dict__["accounting_enabled"] = accounting_enabled
            __props__.__dict__["acct_servers"] = acct_servers
            __props__.__dict__["auth_servers"] = auth_servers
            __props__.__dict__["interim_update_enabled"] = interim_update_enabled
            __props__.__dict__["interim_update_interval"] = interim_update_interval
            __props__.__dict__["name"] = name
            __props__.__dict__["site"] = site
            __props__.__dict__["use_usg_acct_server"] = use_usg_acct_server
            __props__.__dict__["use_usg_auth_server"] = use_usg_auth_server
            __props__.__dict__["vlan_enabled"] = vlan_enabled
            __props__.__dict__["vlan_wlan_mode"] = vlan_wlan_mode
        super(RadiusProfile, __self__).__init__(
            'unifi:index/radiusProfile:RadiusProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accounting_enabled: Optional[pulumi.Input[bool]] = None,
            acct_servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RadiusProfileAcctServerArgs']]]]] = None,
            auth_servers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RadiusProfileAuthServerArgs']]]]] = None,
            interim_update_enabled: Optional[pulumi.Input[bool]] = None,
            interim_update_interval: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            site: Optional[pulumi.Input[str]] = None,
            use_usg_acct_server: Optional[pulumi.Input[bool]] = None,
            use_usg_auth_server: Optional[pulumi.Input[bool]] = None,
            vlan_enabled: Optional[pulumi.Input[bool]] = None,
            vlan_wlan_mode: Optional[pulumi.Input[str]] = None) -> 'RadiusProfile':
        """
        Get an existing RadiusProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] accounting_enabled: Specifies whether to use RADIUS accounting. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RadiusProfileAcctServerArgs']]]] acct_servers: RADIUS accounting servers.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RadiusProfileAuthServerArgs']]]] auth_servers: RADIUS authentication servers.
        :param pulumi.Input[bool] interim_update_enabled: Specifies whether to use interim_update. Defaults to `false`.
        :param pulumi.Input[int] interim_update_interval: Specifies interim_update interval. Defaults to `3600`.
        :param pulumi.Input[str] name: The name of the profile.
        :param pulumi.Input[str] site: The name of the site to associate the settings with.
        :param pulumi.Input[bool] use_usg_acct_server: Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
        :param pulumi.Input[bool] use_usg_auth_server: Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
        :param pulumi.Input[bool] vlan_enabled: Specifies whether to use vlan on wired connections. Defaults to `false`.
        :param pulumi.Input[str] vlan_wlan_mode: Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RadiusProfileState.__new__(_RadiusProfileState)

        __props__.__dict__["accounting_enabled"] = accounting_enabled
        __props__.__dict__["acct_servers"] = acct_servers
        __props__.__dict__["auth_servers"] = auth_servers
        __props__.__dict__["interim_update_enabled"] = interim_update_enabled
        __props__.__dict__["interim_update_interval"] = interim_update_interval
        __props__.__dict__["name"] = name
        __props__.__dict__["site"] = site
        __props__.__dict__["use_usg_acct_server"] = use_usg_acct_server
        __props__.__dict__["use_usg_auth_server"] = use_usg_auth_server
        __props__.__dict__["vlan_enabled"] = vlan_enabled
        __props__.__dict__["vlan_wlan_mode"] = vlan_wlan_mode
        return RadiusProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountingEnabled")
    def accounting_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to use RADIUS accounting. Defaults to `false`.
        """
        return pulumi.get(self, "accounting_enabled")

    @property
    @pulumi.getter(name="acctServers")
    def acct_servers(self) -> pulumi.Output[Optional[Sequence['outputs.RadiusProfileAcctServer']]]:
        """
        RADIUS accounting servers.
        """
        return pulumi.get(self, "acct_servers")

    @property
    @pulumi.getter(name="authServers")
    def auth_servers(self) -> pulumi.Output[Optional[Sequence['outputs.RadiusProfileAuthServer']]]:
        """
        RADIUS authentication servers.
        """
        return pulumi.get(self, "auth_servers")

    @property
    @pulumi.getter(name="interimUpdateEnabled")
    def interim_update_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to use interim_update. Defaults to `false`.
        """
        return pulumi.get(self, "interim_update_enabled")

    @property
    @pulumi.getter(name="interimUpdateInterval")
    def interim_update_interval(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies interim_update interval. Defaults to `3600`.
        """
        return pulumi.get(self, "interim_update_interval")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the profile.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def site(self) -> pulumi.Output[str]:
        """
        The name of the site to associate the settings with.
        """
        return pulumi.get(self, "site")

    @property
    @pulumi.getter(name="useUsgAcctServer")
    def use_usg_acct_server(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to use usg as a RADIUS accounting server. Defaults to `false`.
        """
        return pulumi.get(self, "use_usg_acct_server")

    @property
    @pulumi.getter(name="useUsgAuthServer")
    def use_usg_auth_server(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to use usg as a RADIUS authentication server. Defaults to `false`.
        """
        return pulumi.get(self, "use_usg_auth_server")

    @property
    @pulumi.getter(name="vlanEnabled")
    def vlan_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to use vlan on wired connections. Defaults to `false`.
        """
        return pulumi.get(self, "vlan_enabled")

    @property
    @pulumi.getter(name="vlanWlanMode")
    def vlan_wlan_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether to use vlan on wireless connections. Must be one of `disabled`, `optional`, or `required`. Defaults to ``.
        """
        return pulumi.get(self, "vlan_wlan_mode")

