# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GroupArgs', 'Group']

@pulumi.input_type
class GroupArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 qos_rate_max_down: Optional[pulumi.Input[int]] = None,
                 qos_rate_max_up: Optional[pulumi.Input[int]] = None,
                 site: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Group resource.
        :param pulumi.Input[str] name: The name of the user group.
        :param pulumi.Input[int] qos_rate_max_down: The QOS maximum download rate. Defaults to `-1`.
        :param pulumi.Input[int] qos_rate_max_up: The QOS maximum upload rate. Defaults to `-1`.
        :param pulumi.Input[str] site: The name of the site to associate the user group with.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if qos_rate_max_down is not None:
            pulumi.set(__self__, "qos_rate_max_down", qos_rate_max_down)
        if qos_rate_max_up is not None:
            pulumi.set(__self__, "qos_rate_max_up", qos_rate_max_up)
        if site is not None:
            pulumi.set(__self__, "site", site)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the user group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="qosRateMaxDown")
    def qos_rate_max_down(self) -> Optional[pulumi.Input[int]]:
        """
        The QOS maximum download rate. Defaults to `-1`.
        """
        return pulumi.get(self, "qos_rate_max_down")

    @qos_rate_max_down.setter
    def qos_rate_max_down(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "qos_rate_max_down", value)

    @property
    @pulumi.getter(name="qosRateMaxUp")
    def qos_rate_max_up(self) -> Optional[pulumi.Input[int]]:
        """
        The QOS maximum upload rate. Defaults to `-1`.
        """
        return pulumi.get(self, "qos_rate_max_up")

    @qos_rate_max_up.setter
    def qos_rate_max_up(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "qos_rate_max_up", value)

    @property
    @pulumi.getter
    def site(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the site to associate the user group with.
        """
        return pulumi.get(self, "site")

    @site.setter
    def site(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "site", value)


@pulumi.input_type
class _GroupState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 qos_rate_max_down: Optional[pulumi.Input[int]] = None,
                 qos_rate_max_up: Optional[pulumi.Input[int]] = None,
                 site: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Group resources.
        :param pulumi.Input[str] name: The name of the user group.
        :param pulumi.Input[int] qos_rate_max_down: The QOS maximum download rate. Defaults to `-1`.
        :param pulumi.Input[int] qos_rate_max_up: The QOS maximum upload rate. Defaults to `-1`.
        :param pulumi.Input[str] site: The name of the site to associate the user group with.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if qos_rate_max_down is not None:
            pulumi.set(__self__, "qos_rate_max_down", qos_rate_max_down)
        if qos_rate_max_up is not None:
            pulumi.set(__self__, "qos_rate_max_up", qos_rate_max_up)
        if site is not None:
            pulumi.set(__self__, "site", site)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the user group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="qosRateMaxDown")
    def qos_rate_max_down(self) -> Optional[pulumi.Input[int]]:
        """
        The QOS maximum download rate. Defaults to `-1`.
        """
        return pulumi.get(self, "qos_rate_max_down")

    @qos_rate_max_down.setter
    def qos_rate_max_down(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "qos_rate_max_down", value)

    @property
    @pulumi.getter(name="qosRateMaxUp")
    def qos_rate_max_up(self) -> Optional[pulumi.Input[int]]:
        """
        The QOS maximum upload rate. Defaults to `-1`.
        """
        return pulumi.get(self, "qos_rate_max_up")

    @qos_rate_max_up.setter
    def qos_rate_max_up(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "qos_rate_max_up", value)

    @property
    @pulumi.getter
    def site(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the site to associate the user group with.
        """
        return pulumi.get(self, "site")

    @site.setter
    def site(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "site", value)


class Group(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 qos_rate_max_down: Optional[pulumi.Input[int]] = None,
                 qos_rate_max_up: Optional[pulumi.Input[int]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `iam.Group` manages a user group (called "client group" in the UI), which can be used to limit bandwidth for groups of users.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_unifi as unifi

        wifi = unifi.iam.Group("wifi",
            name="wifi",
            qos_rate_max_down=2000,
            qos_rate_max_up=10)
        ```

        ## Import

        import using the ID

        ```sh
        $ pulumi import unifi:iam/group:Group wifi 5fe6261995fe130013456a36
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the user group.
        :param pulumi.Input[int] qos_rate_max_down: The QOS maximum download rate. Defaults to `-1`.
        :param pulumi.Input[int] qos_rate_max_up: The QOS maximum upload rate. Defaults to `-1`.
        :param pulumi.Input[str] site: The name of the site to associate the user group with.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[GroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `iam.Group` manages a user group (called "client group" in the UI), which can be used to limit bandwidth for groups of users.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_unifi as unifi

        wifi = unifi.iam.Group("wifi",
            name="wifi",
            qos_rate_max_down=2000,
            qos_rate_max_up=10)
        ```

        ## Import

        import using the ID

        ```sh
        $ pulumi import unifi:iam/group:Group wifi 5fe6261995fe130013456a36
        ```

        :param str resource_name: The name of the resource.
        :param GroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 qos_rate_max_down: Optional[pulumi.Input[int]] = None,
                 qos_rate_max_up: Optional[pulumi.Input[int]] = None,
                 site: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupArgs.__new__(GroupArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["qos_rate_max_down"] = qos_rate_max_down
            __props__.__dict__["qos_rate_max_up"] = qos_rate_max_up
            __props__.__dict__["site"] = site
        super(Group, __self__).__init__(
            'unifi:iam/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            qos_rate_max_down: Optional[pulumi.Input[int]] = None,
            qos_rate_max_up: Optional[pulumi.Input[int]] = None,
            site: Optional[pulumi.Input[str]] = None) -> 'Group':
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the user group.
        :param pulumi.Input[int] qos_rate_max_down: The QOS maximum download rate. Defaults to `-1`.
        :param pulumi.Input[int] qos_rate_max_up: The QOS maximum upload rate. Defaults to `-1`.
        :param pulumi.Input[str] site: The name of the site to associate the user group with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupState.__new__(_GroupState)

        __props__.__dict__["name"] = name
        __props__.__dict__["qos_rate_max_down"] = qos_rate_max_down
        __props__.__dict__["qos_rate_max_up"] = qos_rate_max_up
        __props__.__dict__["site"] = site
        return Group(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the user group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="qosRateMaxDown")
    def qos_rate_max_down(self) -> pulumi.Output[Optional[int]]:
        """
        The QOS maximum download rate. Defaults to `-1`.
        """
        return pulumi.get(self, "qos_rate_max_down")

    @property
    @pulumi.getter(name="qosRateMaxUp")
    def qos_rate_max_up(self) -> pulumi.Output[Optional[int]]:
        """
        The QOS maximum upload rate. Defaults to `-1`.
        """
        return pulumi.get(self, "qos_rate_max_up")

    @property
    @pulumi.getter
    def site(self) -> pulumi.Output[str]:
        """
        The name of the site to associate the user group with.
        """
        return pulumi.get(self, "site")

