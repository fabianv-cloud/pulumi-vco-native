// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package virtual_machine

import (
	"context"
	"reflect"

	"errors"
	"github.com/fabianv-cloud/pulumi-vco-native/sdk/go/command/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineDisk struct {
	pulumi.CustomResourceState

	Cloudspace_id pulumi.StringOutput `pulumi:"cloudspace_id"`
	CustomerID    pulumi.StringOutput `pulumi:"customerID"`
	Disk_id       pulumi.IntOutput    `pulumi:"disk_id"`
	Token         pulumi.StringOutput `pulumi:"token"`
	Url           pulumi.StringOutput `pulumi:"url"`
	Vm_id         pulumi.IntOutput    `pulumi:"vm_id"`
}

// NewVirtualMachineDisk registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineDisk(ctx *pulumi.Context,
	name string, args *VirtualMachineDiskArgs, opts ...pulumi.ResourceOption) (*VirtualMachineDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cloudspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Cloudspace_id'")
	}
	if args.CustomerID == nil {
		return nil, errors.New("invalid value for required argument 'CustomerID'")
	}
	if args.Disk_id == nil {
		return nil, errors.New("invalid value for required argument 'Disk_id'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Vm_id == nil {
		return nil, errors.New("invalid value for required argument 'Vm_id'")
	}
	if args.CustomerID != nil {
		args.CustomerID = pulumi.ToSecret(args.CustomerID).(pulumi.StringInput)
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringInput)
	}
	if args.Url != nil {
		args.Url = pulumi.ToSecret(args.Url).(pulumi.StringInput)
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualMachineDisk
	err := ctx.RegisterResource("vco:virtual_machine:VirtualMachineDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineDisk gets an existing VirtualMachineDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineDiskState, opts ...pulumi.ResourceOption) (*VirtualMachineDisk, error) {
	var resource VirtualMachineDisk
	err := ctx.ReadResource("vco:virtual_machine:VirtualMachineDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineDisk resources.
type virtualMachineDiskState struct {
}

type VirtualMachineDiskState struct {
}

func (VirtualMachineDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineDiskState)(nil)).Elem()
}

type virtualMachineDiskArgs struct {
	Cloudspace_id string `pulumi:"cloudspace_id"`
	CustomerID    string `pulumi:"customerID"`
	Disk_id       int    `pulumi:"disk_id"`
	Token         string `pulumi:"token"`
	Url           string `pulumi:"url"`
	Vm_id         int    `pulumi:"vm_id"`
}

// The set of arguments for constructing a VirtualMachineDisk resource.
type VirtualMachineDiskArgs struct {
	Cloudspace_id pulumi.StringInput
	CustomerID    pulumi.StringInput
	Disk_id       pulumi.IntInput
	Token         pulumi.StringInput
	Url           pulumi.StringInput
	Vm_id         pulumi.IntInput
}

func (VirtualMachineDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineDiskArgs)(nil)).Elem()
}

type VirtualMachineDiskInput interface {
	pulumi.Input

	ToVirtualMachineDiskOutput() VirtualMachineDiskOutput
	ToVirtualMachineDiskOutputWithContext(ctx context.Context) VirtualMachineDiskOutput
}

func (*VirtualMachineDisk) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineDisk)(nil)).Elem()
}

func (i *VirtualMachineDisk) ToVirtualMachineDiskOutput() VirtualMachineDiskOutput {
	return i.ToVirtualMachineDiskOutputWithContext(context.Background())
}

func (i *VirtualMachineDisk) ToVirtualMachineDiskOutputWithContext(ctx context.Context) VirtualMachineDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineDiskOutput)
}

// VirtualMachineDiskArrayInput is an input type that accepts VirtualMachineDiskArray and VirtualMachineDiskArrayOutput values.
// You can construct a concrete instance of `VirtualMachineDiskArrayInput` via:
//
//	VirtualMachineDiskArray{ VirtualMachineDiskArgs{...} }
type VirtualMachineDiskArrayInput interface {
	pulumi.Input

	ToVirtualMachineDiskArrayOutput() VirtualMachineDiskArrayOutput
	ToVirtualMachineDiskArrayOutputWithContext(context.Context) VirtualMachineDiskArrayOutput
}

type VirtualMachineDiskArray []VirtualMachineDiskInput

func (VirtualMachineDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualMachineDisk)(nil)).Elem()
}

func (i VirtualMachineDiskArray) ToVirtualMachineDiskArrayOutput() VirtualMachineDiskArrayOutput {
	return i.ToVirtualMachineDiskArrayOutputWithContext(context.Background())
}

func (i VirtualMachineDiskArray) ToVirtualMachineDiskArrayOutputWithContext(ctx context.Context) VirtualMachineDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineDiskArrayOutput)
}

// VirtualMachineDiskMapInput is an input type that accepts VirtualMachineDiskMap and VirtualMachineDiskMapOutput values.
// You can construct a concrete instance of `VirtualMachineDiskMapInput` via:
//
//	VirtualMachineDiskMap{ "key": VirtualMachineDiskArgs{...} }
type VirtualMachineDiskMapInput interface {
	pulumi.Input

	ToVirtualMachineDiskMapOutput() VirtualMachineDiskMapOutput
	ToVirtualMachineDiskMapOutputWithContext(context.Context) VirtualMachineDiskMapOutput
}

type VirtualMachineDiskMap map[string]VirtualMachineDiskInput

func (VirtualMachineDiskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualMachineDisk)(nil)).Elem()
}

func (i VirtualMachineDiskMap) ToVirtualMachineDiskMapOutput() VirtualMachineDiskMapOutput {
	return i.ToVirtualMachineDiskMapOutputWithContext(context.Background())
}

func (i VirtualMachineDiskMap) ToVirtualMachineDiskMapOutputWithContext(ctx context.Context) VirtualMachineDiskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineDiskMapOutput)
}

type VirtualMachineDiskOutput struct{ *pulumi.OutputState }

func (VirtualMachineDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineDisk)(nil)).Elem()
}

func (o VirtualMachineDiskOutput) ToVirtualMachineDiskOutput() VirtualMachineDiskOutput {
	return o
}

func (o VirtualMachineDiskOutput) ToVirtualMachineDiskOutputWithContext(ctx context.Context) VirtualMachineDiskOutput {
	return o
}

func (o VirtualMachineDiskOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineDisk) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o VirtualMachineDiskOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineDisk) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o VirtualMachineDiskOutput) Disk_id() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineDisk) pulumi.IntOutput { return v.Disk_id }).(pulumi.IntOutput)
}

func (o VirtualMachineDiskOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineDisk) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o VirtualMachineDiskOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineDisk) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func (o VirtualMachineDiskOutput) Vm_id() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineDisk) pulumi.IntOutput { return v.Vm_id }).(pulumi.IntOutput)
}

type VirtualMachineDiskArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualMachineDisk)(nil)).Elem()
}

func (o VirtualMachineDiskArrayOutput) ToVirtualMachineDiskArrayOutput() VirtualMachineDiskArrayOutput {
	return o
}

func (o VirtualMachineDiskArrayOutput) ToVirtualMachineDiskArrayOutputWithContext(ctx context.Context) VirtualMachineDiskArrayOutput {
	return o
}

func (o VirtualMachineDiskArrayOutput) Index(i pulumi.IntInput) VirtualMachineDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualMachineDisk {
		return vs[0].([]*VirtualMachineDisk)[vs[1].(int)]
	}).(VirtualMachineDiskOutput)
}

type VirtualMachineDiskMapOutput struct{ *pulumi.OutputState }

func (VirtualMachineDiskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualMachineDisk)(nil)).Elem()
}

func (o VirtualMachineDiskMapOutput) ToVirtualMachineDiskMapOutput() VirtualMachineDiskMapOutput {
	return o
}

func (o VirtualMachineDiskMapOutput) ToVirtualMachineDiskMapOutputWithContext(ctx context.Context) VirtualMachineDiskMapOutput {
	return o
}

func (o VirtualMachineDiskMapOutput) MapIndex(k pulumi.StringInput) VirtualMachineDiskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualMachineDisk {
		return vs[0].(map[string]*VirtualMachineDisk)[vs[1].(string)]
	}).(VirtualMachineDiskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineDiskInput)(nil)).Elem(), &VirtualMachineDisk{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineDiskArrayInput)(nil)).Elem(), VirtualMachineDiskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineDiskMapInput)(nil)).Elem(), VirtualMachineDiskMap{})
	pulumi.RegisterOutputType(VirtualMachineDiskOutput{})
	pulumi.RegisterOutputType(VirtualMachineDiskArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineDiskMapOutput{})
}
