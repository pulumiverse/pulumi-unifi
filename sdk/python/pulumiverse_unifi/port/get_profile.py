# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetProfileResult',
    'AwaitableGetProfileResult',
    'get_profile',
    'get_profile_output',
]

@pulumi.output_type
class GetProfileResult:
    """
    A collection of values returned by getProfile.
    """
    def __init__(__self__, id=None, name=None, site=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if site and not isinstance(site, str):
            raise TypeError("Expected argument 'site' to be a str")
        pulumi.set(__self__, "site", site)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this port profile.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the port profile to look up. Defaults to `All`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def site(self) -> str:
        """
        The name of the site the port profile is associated with.
        """
        return pulumi.get(self, "site")


class AwaitableGetProfileResult(GetProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProfileResult(
            id=self.id,
            name=self.name,
            site=self.site)


def get_profile(name: Optional[str] = None,
                site: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProfileResult:
    """
    `port.Profile` data source can be used to retrieve the ID for a port profile by name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_unifi as unifi

    all = unifi.port.get_profile()
    ```


    :param str name: The name of the port profile to look up. Defaults to `All`.
    :param str site: The name of the site the port profile is associated with.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['site'] = site
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('unifi:port/getProfile:getProfile', __args__, opts=opts, typ=GetProfileResult).value

    return AwaitableGetProfileResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        site=pulumi.get(__ret__, 'site'))


@_utilities.lift_output_func(get_profile)
def get_profile_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                       site: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProfileResult]:
    """
    `port.Profile` data source can be used to retrieve the ID for a port profile by name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_unifi as unifi

    all = unifi.port.get_profile()
    ```


    :param str name: The name of the port profile to look up. Defaults to `All`.
    :param str site: The name of the site the port profile is associated with.
    """
    ...
