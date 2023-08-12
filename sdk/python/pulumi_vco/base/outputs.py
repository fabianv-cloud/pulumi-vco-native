# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'IOTune',
]

@pulumi.output_type
class IOTune(dict):
    def __init__(__self__, *,
                 iops: float):
        pulumi.set(__self__, "iops", iops)

    @property
    @pulumi.getter
    def iops(self) -> float:
        return pulumi.get(self, "iops")


