# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['MgmtArgs', 'Mgmt']

@pulumi.input_type
class MgmtArgs:
    def __init__(__self__, *,
                 auto_upgrade: Optional[pulumi.Input[bool]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 ssh_enabled: Optional[pulumi.Input[bool]] = None,
                 ssh_keys: Optional[pulumi.Input[Sequence[pulumi.Input['MgmtSshKeyArgs']]]] = None):
        """
        The set of arguments for constructing a Mgmt resource.
        :param pulumi.Input[bool] auto_upgrade: Automatically upgrade device firmware.
        :param pulumi.Input[str] site: The name of the site to associate the settings with.
        :param pulumi.Input[bool] ssh_enabled: Enable SSH authentication.
        :param pulumi.Input[Sequence[pulumi.Input['MgmtSshKeyArgs']]] ssh_keys: SSH key.
        """
        if auto_upgrade is not None:
            pulumi.set(__self__, "auto_upgrade", auto_upgrade)
        if site is not None:
            pulumi.set(__self__, "site", site)
        if ssh_enabled is not None:
            pulumi.set(__self__, "ssh_enabled", ssh_enabled)
        if ssh_keys is not None:
            pulumi.set(__self__, "ssh_keys", ssh_keys)

    @property
    @pulumi.getter(name="autoUpgrade")
    def auto_upgrade(self) -> Optional[pulumi.Input[bool]]:
        """
        Automatically upgrade device firmware.
        """
        return pulumi.get(self, "auto_upgrade")

    @auto_upgrade.setter
    def auto_upgrade(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_upgrade", value)

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
    @pulumi.getter(name="sshEnabled")
    def ssh_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable SSH authentication.
        """
        return pulumi.get(self, "ssh_enabled")

    @ssh_enabled.setter
    def ssh_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ssh_enabled", value)

    @property
    @pulumi.getter(name="sshKeys")
    def ssh_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MgmtSshKeyArgs']]]]:
        """
        SSH key.
        """
        return pulumi.get(self, "ssh_keys")

    @ssh_keys.setter
    def ssh_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MgmtSshKeyArgs']]]]):
        pulumi.set(self, "ssh_keys", value)


@pulumi.input_type
class _MgmtState:
    def __init__(__self__, *,
                 auto_upgrade: Optional[pulumi.Input[bool]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 ssh_enabled: Optional[pulumi.Input[bool]] = None,
                 ssh_keys: Optional[pulumi.Input[Sequence[pulumi.Input['MgmtSshKeyArgs']]]] = None):
        """
        Input properties used for looking up and filtering Mgmt resources.
        :param pulumi.Input[bool] auto_upgrade: Automatically upgrade device firmware.
        :param pulumi.Input[str] site: The name of the site to associate the settings with.
        :param pulumi.Input[bool] ssh_enabled: Enable SSH authentication.
        :param pulumi.Input[Sequence[pulumi.Input['MgmtSshKeyArgs']]] ssh_keys: SSH key.
        """
        if auto_upgrade is not None:
            pulumi.set(__self__, "auto_upgrade", auto_upgrade)
        if site is not None:
            pulumi.set(__self__, "site", site)
        if ssh_enabled is not None:
            pulumi.set(__self__, "ssh_enabled", ssh_enabled)
        if ssh_keys is not None:
            pulumi.set(__self__, "ssh_keys", ssh_keys)

    @property
    @pulumi.getter(name="autoUpgrade")
    def auto_upgrade(self) -> Optional[pulumi.Input[bool]]:
        """
        Automatically upgrade device firmware.
        """
        return pulumi.get(self, "auto_upgrade")

    @auto_upgrade.setter
    def auto_upgrade(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_upgrade", value)

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
    @pulumi.getter(name="sshEnabled")
    def ssh_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable SSH authentication.
        """
        return pulumi.get(self, "ssh_enabled")

    @ssh_enabled.setter
    def ssh_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ssh_enabled", value)

    @property
    @pulumi.getter(name="sshKeys")
    def ssh_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MgmtSshKeyArgs']]]]:
        """
        SSH key.
        """
        return pulumi.get(self, "ssh_keys")

    @ssh_keys.setter
    def ssh_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MgmtSshKeyArgs']]]]):
        pulumi.set(self, "ssh_keys", value)


class Mgmt(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_upgrade: Optional[pulumi.Input[bool]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 ssh_enabled: Optional[pulumi.Input[bool]] = None,
                 ssh_keys: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MgmtSshKeyArgs', 'MgmtSshKeyArgsDict']]]]] = None,
                 __props__=None):
        """
        `setting.Mgmt` manages settings for a unifi site.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_unifi as unifi

        example_site = unifi.Site("exampleSite", description="example")
        example_mgmt = unifi.setting.Mgmt("exampleMgmt",
            site=example_site.name,
            auto_upgrade=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_upgrade: Automatically upgrade device firmware.
        :param pulumi.Input[str] site: The name of the site to associate the settings with.
        :param pulumi.Input[bool] ssh_enabled: Enable SSH authentication.
        :param pulumi.Input[Sequence[pulumi.Input[Union['MgmtSshKeyArgs', 'MgmtSshKeyArgsDict']]]] ssh_keys: SSH key.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[MgmtArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `setting.Mgmt` manages settings for a unifi site.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_unifi as unifi

        example_site = unifi.Site("exampleSite", description="example")
        example_mgmt = unifi.setting.Mgmt("exampleMgmt",
            site=example_site.name,
            auto_upgrade=True)
        ```

        :param str resource_name: The name of the resource.
        :param MgmtArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MgmtArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_upgrade: Optional[pulumi.Input[bool]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 ssh_enabled: Optional[pulumi.Input[bool]] = None,
                 ssh_keys: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MgmtSshKeyArgs', 'MgmtSshKeyArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MgmtArgs.__new__(MgmtArgs)

            __props__.__dict__["auto_upgrade"] = auto_upgrade
            __props__.__dict__["site"] = site
            __props__.__dict__["ssh_enabled"] = ssh_enabled
            __props__.__dict__["ssh_keys"] = ssh_keys
        super(Mgmt, __self__).__init__(
            'unifi:setting/mgmt:Mgmt',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_upgrade: Optional[pulumi.Input[bool]] = None,
            site: Optional[pulumi.Input[str]] = None,
            ssh_enabled: Optional[pulumi.Input[bool]] = None,
            ssh_keys: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MgmtSshKeyArgs', 'MgmtSshKeyArgsDict']]]]] = None) -> 'Mgmt':
        """
        Get an existing Mgmt resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_upgrade: Automatically upgrade device firmware.
        :param pulumi.Input[str] site: The name of the site to associate the settings with.
        :param pulumi.Input[bool] ssh_enabled: Enable SSH authentication.
        :param pulumi.Input[Sequence[pulumi.Input[Union['MgmtSshKeyArgs', 'MgmtSshKeyArgsDict']]]] ssh_keys: SSH key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MgmtState.__new__(_MgmtState)

        __props__.__dict__["auto_upgrade"] = auto_upgrade
        __props__.__dict__["site"] = site
        __props__.__dict__["ssh_enabled"] = ssh_enabled
        __props__.__dict__["ssh_keys"] = ssh_keys
        return Mgmt(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoUpgrade")
    def auto_upgrade(self) -> pulumi.Output[Optional[bool]]:
        """
        Automatically upgrade device firmware.
        """
        return pulumi.get(self, "auto_upgrade")

    @property
    @pulumi.getter
    def site(self) -> pulumi.Output[str]:
        """
        The name of the site to associate the settings with.
        """
        return pulumi.get(self, "site")

    @property
    @pulumi.getter(name="sshEnabled")
    def ssh_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable SSH authentication.
        """
        return pulumi.get(self, "ssh_enabled")

    @property
    @pulumi.getter(name="sshKeys")
    def ssh_keys(self) -> pulumi.Output[Optional[Sequence['outputs.MgmtSshKey']]]:
        """
        SSH key.
        """
        return pulumi.get(self, "ssh_keys")

