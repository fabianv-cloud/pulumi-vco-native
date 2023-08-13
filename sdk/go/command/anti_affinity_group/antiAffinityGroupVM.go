// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package anti_affinity_group

import (
	"context"
	"reflect"

	"errors"
	"github.com/fabianv-cloud/pulumi-vco-native/sdk/go/command/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AntiAffinityGroupVM struct {
	pulumi.CustomResourceState

	Cloudspace_id pulumi.StringOutput `pulumi:"cloudspace_id"`
	CustomerID    pulumi.StringOutput `pulumi:"customerID"`
	Group_id      pulumi.StringOutput `pulumi:"group_id"`
	Status        pulumi.StringOutput `pulumi:"status"`
	Token         pulumi.StringOutput `pulumi:"token"`
	Url           pulumi.StringOutput `pulumi:"url"`
	Vm_id         pulumi.IntOutput    `pulumi:"vm_id"`
}

// NewAntiAffinityGroupVM registers a new resource with the given unique name, arguments, and options.
func NewAntiAffinityGroupVM(ctx *pulumi.Context,
	name string, args *AntiAffinityGroupVMArgs, opts ...pulumi.ResourceOption) (*AntiAffinityGroupVM, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cloudspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Cloudspace_id'")
	}
	if args.CustomerID == nil {
		return nil, errors.New("invalid value for required argument 'CustomerID'")
	}
	if args.Group_id == nil {
		return nil, errors.New("invalid value for required argument 'Group_id'")
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
	var resource AntiAffinityGroupVM
	err := ctx.RegisterResource("vco:anti_affinity_group:AntiAffinityGroupVM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAntiAffinityGroupVM gets an existing AntiAffinityGroupVM resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAntiAffinityGroupVM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AntiAffinityGroupVMState, opts ...pulumi.ResourceOption) (*AntiAffinityGroupVM, error) {
	var resource AntiAffinityGroupVM
	err := ctx.ReadResource("vco:anti_affinity_group:AntiAffinityGroupVM", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AntiAffinityGroupVM resources.
type antiAffinityGroupVMState struct {
}

type AntiAffinityGroupVMState struct {
}

func (AntiAffinityGroupVMState) ElementType() reflect.Type {
	return reflect.TypeOf((*antiAffinityGroupVMState)(nil)).Elem()
}

type antiAffinityGroupVMArgs struct {
	Cloudspace_id string `pulumi:"cloudspace_id"`
	CustomerID    string `pulumi:"customerID"`
	Group_id      string `pulumi:"group_id"`
	Token         string `pulumi:"token"`
	Url           string `pulumi:"url"`
	Vm_id         int    `pulumi:"vm_id"`
}

// The set of arguments for constructing a AntiAffinityGroupVM resource.
type AntiAffinityGroupVMArgs struct {
	Cloudspace_id pulumi.StringInput
	CustomerID    pulumi.StringInput
	Group_id      pulumi.StringInput
	Token         pulumi.StringInput
	Url           pulumi.StringInput
	Vm_id         pulumi.IntInput
}

func (AntiAffinityGroupVMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*antiAffinityGroupVMArgs)(nil)).Elem()
}

type AntiAffinityGroupVMInput interface {
	pulumi.Input

	ToAntiAffinityGroupVMOutput() AntiAffinityGroupVMOutput
	ToAntiAffinityGroupVMOutputWithContext(ctx context.Context) AntiAffinityGroupVMOutput
}

func (*AntiAffinityGroupVM) ElementType() reflect.Type {
	return reflect.TypeOf((**AntiAffinityGroupVM)(nil)).Elem()
}

func (i *AntiAffinityGroupVM) ToAntiAffinityGroupVMOutput() AntiAffinityGroupVMOutput {
	return i.ToAntiAffinityGroupVMOutputWithContext(context.Background())
}

func (i *AntiAffinityGroupVM) ToAntiAffinityGroupVMOutputWithContext(ctx context.Context) AntiAffinityGroupVMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AntiAffinityGroupVMOutput)
}

// AntiAffinityGroupVMArrayInput is an input type that accepts AntiAffinityGroupVMArray and AntiAffinityGroupVMArrayOutput values.
// You can construct a concrete instance of `AntiAffinityGroupVMArrayInput` via:
//
//	AntiAffinityGroupVMArray{ AntiAffinityGroupVMArgs{...} }
type AntiAffinityGroupVMArrayInput interface {
	pulumi.Input

	ToAntiAffinityGroupVMArrayOutput() AntiAffinityGroupVMArrayOutput
	ToAntiAffinityGroupVMArrayOutputWithContext(context.Context) AntiAffinityGroupVMArrayOutput
}

type AntiAffinityGroupVMArray []AntiAffinityGroupVMInput

func (AntiAffinityGroupVMArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AntiAffinityGroupVM)(nil)).Elem()
}

func (i AntiAffinityGroupVMArray) ToAntiAffinityGroupVMArrayOutput() AntiAffinityGroupVMArrayOutput {
	return i.ToAntiAffinityGroupVMArrayOutputWithContext(context.Background())
}

func (i AntiAffinityGroupVMArray) ToAntiAffinityGroupVMArrayOutputWithContext(ctx context.Context) AntiAffinityGroupVMArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AntiAffinityGroupVMArrayOutput)
}

// AntiAffinityGroupVMMapInput is an input type that accepts AntiAffinityGroupVMMap and AntiAffinityGroupVMMapOutput values.
// You can construct a concrete instance of `AntiAffinityGroupVMMapInput` via:
//
//	AntiAffinityGroupVMMap{ "key": AntiAffinityGroupVMArgs{...} }
type AntiAffinityGroupVMMapInput interface {
	pulumi.Input

	ToAntiAffinityGroupVMMapOutput() AntiAffinityGroupVMMapOutput
	ToAntiAffinityGroupVMMapOutputWithContext(context.Context) AntiAffinityGroupVMMapOutput
}

type AntiAffinityGroupVMMap map[string]AntiAffinityGroupVMInput

func (AntiAffinityGroupVMMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AntiAffinityGroupVM)(nil)).Elem()
}

func (i AntiAffinityGroupVMMap) ToAntiAffinityGroupVMMapOutput() AntiAffinityGroupVMMapOutput {
	return i.ToAntiAffinityGroupVMMapOutputWithContext(context.Background())
}

func (i AntiAffinityGroupVMMap) ToAntiAffinityGroupVMMapOutputWithContext(ctx context.Context) AntiAffinityGroupVMMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AntiAffinityGroupVMMapOutput)
}

type AntiAffinityGroupVMOutput struct{ *pulumi.OutputState }

func (AntiAffinityGroupVMOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AntiAffinityGroupVM)(nil)).Elem()
}

func (o AntiAffinityGroupVMOutput) ToAntiAffinityGroupVMOutput() AntiAffinityGroupVMOutput {
	return o
}

func (o AntiAffinityGroupVMOutput) ToAntiAffinityGroupVMOutputWithContext(ctx context.Context) AntiAffinityGroupVMOutput {
	return o
}

func (o AntiAffinityGroupVMOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiAffinityGroupVM) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o AntiAffinityGroupVMOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiAffinityGroupVM) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o AntiAffinityGroupVMOutput) Group_id() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiAffinityGroupVM) pulumi.StringOutput { return v.Group_id }).(pulumi.StringOutput)
}

func (o AntiAffinityGroupVMOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiAffinityGroupVM) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o AntiAffinityGroupVMOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiAffinityGroupVM) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o AntiAffinityGroupVMOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiAffinityGroupVM) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func (o AntiAffinityGroupVMOutput) Vm_id() pulumi.IntOutput {
	return o.ApplyT(func(v *AntiAffinityGroupVM) pulumi.IntOutput { return v.Vm_id }).(pulumi.IntOutput)
}

type AntiAffinityGroupVMArrayOutput struct{ *pulumi.OutputState }

func (AntiAffinityGroupVMArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AntiAffinityGroupVM)(nil)).Elem()
}

func (o AntiAffinityGroupVMArrayOutput) ToAntiAffinityGroupVMArrayOutput() AntiAffinityGroupVMArrayOutput {
	return o
}

func (o AntiAffinityGroupVMArrayOutput) ToAntiAffinityGroupVMArrayOutputWithContext(ctx context.Context) AntiAffinityGroupVMArrayOutput {
	return o
}

func (o AntiAffinityGroupVMArrayOutput) Index(i pulumi.IntInput) AntiAffinityGroupVMOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AntiAffinityGroupVM {
		return vs[0].([]*AntiAffinityGroupVM)[vs[1].(int)]
	}).(AntiAffinityGroupVMOutput)
}

type AntiAffinityGroupVMMapOutput struct{ *pulumi.OutputState }

func (AntiAffinityGroupVMMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AntiAffinityGroupVM)(nil)).Elem()
}

func (o AntiAffinityGroupVMMapOutput) ToAntiAffinityGroupVMMapOutput() AntiAffinityGroupVMMapOutput {
	return o
}

func (o AntiAffinityGroupVMMapOutput) ToAntiAffinityGroupVMMapOutputWithContext(ctx context.Context) AntiAffinityGroupVMMapOutput {
	return o
}

func (o AntiAffinityGroupVMMapOutput) MapIndex(k pulumi.StringInput) AntiAffinityGroupVMOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AntiAffinityGroupVM {
		return vs[0].(map[string]*AntiAffinityGroupVM)[vs[1].(string)]
	}).(AntiAffinityGroupVMOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AntiAffinityGroupVMInput)(nil)).Elem(), &AntiAffinityGroupVM{})
	pulumi.RegisterInputType(reflect.TypeOf((*AntiAffinityGroupVMArrayInput)(nil)).Elem(), AntiAffinityGroupVMArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AntiAffinityGroupVMMapInput)(nil)).Elem(), AntiAffinityGroupVMMap{})
	pulumi.RegisterOutputType(AntiAffinityGroupVMOutput{})
	pulumi.RegisterOutputType(AntiAffinityGroupVMArrayOutput{})
	pulumi.RegisterOutputType(AntiAffinityGroupVMMapOutput{})
}
