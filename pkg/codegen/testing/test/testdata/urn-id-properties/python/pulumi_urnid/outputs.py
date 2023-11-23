# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'InnerType',
]

@pulumi.output_type
class InnerType(dict):
    """
    It's fine to use urn and id in nested objects
    """
    def __init__(__self__, *,
                 id: Optional[str] = None,
                 urn: Optional[str] = None):
        """
        It's fine to use urn and id in nested objects
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if urn is not None:
            pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def urn(self) -> Optional[str]:
        return pulumi.get(self, "urn")


