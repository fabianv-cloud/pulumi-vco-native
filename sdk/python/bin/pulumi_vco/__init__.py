# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_vco.anti_affinity_group as __anti_affinity_group
    anti_affinity_group = __anti_affinity_group
    import pulumi_vco.base as __base
    base = __base
    import pulumi_vco.cloudspace as __cloudspace
    cloudspace = __cloudspace
    import pulumi_vco.disk as __disk
    disk = __disk
    import pulumi_vco.ingress as __ingress
    ingress = __ingress
    import pulumi_vco.objectspace as __objectspace
    objectspace = __objectspace
    import pulumi_vco.virtual_machine as __virtual_machine
    virtual_machine = __virtual_machine
else:
    anti_affinity_group = _utilities.lazy_import('pulumi_vco.anti_affinity_group')
    base = _utilities.lazy_import('pulumi_vco.base')
    cloudspace = _utilities.lazy_import('pulumi_vco.cloudspace')
    disk = _utilities.lazy_import('pulumi_vco.disk')
    ingress = _utilities.lazy_import('pulumi_vco.ingress')
    objectspace = _utilities.lazy_import('pulumi_vco.objectspace')
    virtual_machine = _utilities.lazy_import('pulumi_vco.virtual_machine')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "vco",
  "mod": "anti_affinity_group",
  "fqn": "pulumi_vco.anti_affinity_group",
  "classes": {
   "vco:anti_affinity_group:AntiAffinityGroupVM": "AntiAffinityGroupVM"
  }
 },
 {
  "pkg": "vco",
  "mod": "base",
  "fqn": "pulumi_vco.base",
  "classes": {
   "vco:base:Cloudspace": "Cloudspace",
   "vco:base:Disk": "Disk"
  }
 },
 {
  "pkg": "vco",
  "mod": "cloudspace",
  "fqn": "pulumi_vco.cloudspace",
  "classes": {
   "vco:cloudspace:AntiAffinityGroup": "AntiAffinityGroup",
   "vco:cloudspace:ConnectedCloudspace": "ConnectedCloudspace",
   "vco:cloudspace:ExternalNetwork": "ExternalNetwork",
   "vco:cloudspace:PortForward": "PortForward",
   "vco:cloudspace:VirtualMachine": "VirtualMachine"
  }
 },
 {
  "pkg": "vco",
  "mod": "disk",
  "fqn": "pulumi_vco.disk",
  "classes": {
   "vco:disk:ExposedDisk": "ExposedDisk"
  }
 },
 {
  "pkg": "vco",
  "mod": "ingress",
  "fqn": "pulumi_vco.ingress",
  "classes": {
   "vco:ingress:Host": "Host",
   "vco:ingress:LoadBalancer": "LoadBalancer",
   "vco:ingress:ReverseProxy": "ReverseProxy",
   "vco:ingress:ServerPool": "ServerPool"
  }
 },
 {
  "pkg": "vco",
  "mod": "objectspace",
  "fqn": "pulumi_vco.objectspace",
  "classes": {
   "vco:objectspace:ObjectSpaceLink": "ObjectSpaceLink"
  }
 },
 {
  "pkg": "vco",
  "mod": "virtual_machine",
  "fqn": "pulumi_vco.virtual_machine",
  "classes": {
   "vco:virtual_machine:VirtualMachineCD": "VirtualMachineCD",
   "vco:virtual_machine:VirtualMachineDisk": "VirtualMachineDisk",
   "vco:virtual_machine:VirtualMachineNIC": "VirtualMachineNIC"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "vco",
  "token": "pulumi:providers:vco",
  "fqn": "pulumi_vco",
  "class": "Provider"
 }
]
"""
)
