# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetGroupResult',
    'AwaitableGetGroupResult',
    'get_group',
    'get_group_output',
]

@pulumi.output_type
class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, id=None, name=None, qos_rate_max_down=None, qos_rate_max_up=None, site=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if qos_rate_max_down and not isinstance(qos_rate_max_down, int):
            raise TypeError("Expected argument 'qos_rate_max_down' to be a int")
        pulumi.set(__self__, "qos_rate_max_down", qos_rate_max_down)
        if qos_rate_max_up and not isinstance(qos_rate_max_up, int):
            raise TypeError("Expected argument 'qos_rate_max_up' to be a int")
        pulumi.set(__self__, "qos_rate_max_up", qos_rate_max_up)
        if site and not isinstance(site, str):
            raise TypeError("Expected argument 'site' to be a str")
        pulumi.set(__self__, "site", site)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this AP group.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the user group to look up. Defaults to `Default`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="qosRateMaxDown")
    def qos_rate_max_down(self) -> int:
        return pulumi.get(self, "qos_rate_max_down")

    @property
    @pulumi.getter(name="qosRateMaxUp")
    def qos_rate_max_up(self) -> int:
        return pulumi.get(self, "qos_rate_max_up")

    @property
    @pulumi.getter
    def site(self) -> str:
        """
        The name of the site the user group is associated with.
        """
        return pulumi.get(self, "site")


class AwaitableGetGroupResult(GetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupResult(
            id=self.id,
            name=self.name,
            qos_rate_max_down=self.qos_rate_max_down,
            qos_rate_max_up=self.qos_rate_max_up,
            site=self.site)


def get_group(name: Optional[str] = None,
              site: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupResult:
    """
    `IAM.Group` data source can be used to retrieve the ID for a user group by name.


    :param str name: The name of the user group to look up. Defaults to `Default`.
    :param str site: The name of the site the user group is associated with.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['site'] = site
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('unifi:IAM/getGroup:getGroup', __args__, opts=opts, typ=GetGroupResult).value

    return AwaitableGetGroupResult(
        id=__ret__.id,
        name=__ret__.name,
        qos_rate_max_down=__ret__.qos_rate_max_down,
        qos_rate_max_up=__ret__.qos_rate_max_up,
        site=__ret__.site)


@_utilities.lift_output_func(get_group)
def get_group_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                     site: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupResult]:
    """
    `IAM.Group` data source can be used to retrieve the ID for a user group by name.


    :param str name: The name of the user group to look up. Defaults to `Default`.
    :param str site: The name of the site the user group is associated with.
    """
    ...
