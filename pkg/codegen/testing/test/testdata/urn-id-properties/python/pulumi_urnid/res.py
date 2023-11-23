# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = ['ResArgs', 'Res']

@pulumi.input_type
class ResArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 urn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Res resource.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if urn is not None:
            pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def urn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "urn")

    @urn.setter
    def urn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "urn", value)


class Res(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 urn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        It's fine to use urn and id as input properties

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ResArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        It's fine to use urn and id as input properties

        :param str resource_name: The name of the resource.
        :param ResArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 urn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResArgs.__new__(ResArgs)

            __props__.__dict__["id"] = id
            __props__.__dict__["urn"] = urn
            __props__.__dict__["output"] = None
        super(Res, __self__).__init__(
            'urnid:index:Res',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Res':
        """
        Get an existing Res resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ResArgs.__new__(ResArgs)

        __props__.__dict__["output"] = None
        return Res(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def output(self) -> pulumi.Output[Optional['outputs.InnerType']]:
        return pulumi.get(self, "output")

