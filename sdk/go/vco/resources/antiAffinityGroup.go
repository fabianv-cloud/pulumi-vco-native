// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

type AntiAffinityGroup struct {
	pulumi.CustomResourceState

	Cloudspace_id pulumi.StringOutput `pulumi:"cloudspace_id"`
	CustomerID    pulumi.StringOutput `pulumi:"customerID"`
	Group_id      pulumi.StringOutput `pulumi:"group_id"`
	Spread        pulumi.IntOutput    `pulumi:"spread"`
	Status        pulumi.StringOutput `pulumi:"status"`
	Token         pulumi.StringOutput `pulumi:"token"`
	Url           pulumi.StringOutput `pulumi:"url"`
}

// NewAntiAffinityGroup registers a new resource with the given unique name, arguments, and options.
func NewAntiAffinityGroup(ctx *pulumi.Context,
	name string, args *AntiAffinityGroupArgs, opts ...pulumi.ResourceOption) (*AntiAffinityGroup, error) {
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
	if args.Spread == nil {
		return nil, errors.New("invalid value for required argument 'Spread'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AntiAffinityGroup
	err := ctx.RegisterResource("vco:resources:AntiAffinityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAntiAffinityGroup gets an existing AntiAffinityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAntiAffinityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AntiAffinityGroupState, opts ...pulumi.ResourceOption) (*AntiAffinityGroup, error) {
	var resource AntiAffinityGroup
	err := ctx.ReadResource("vco:resources:AntiAffinityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AntiAffinityGroup resources.
type antiAffinityGroupState struct {
}

type AntiAffinityGroupState struct {
}

func (AntiAffinityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*antiAffinityGroupState)(nil)).Elem()
}

type antiAffinityGroupArgs struct {
	Cloudspace_id string `pulumi:"cloudspace_id"`
	CustomerID    string `pulumi:"customerID"`
	Group_id      string `pulumi:"group_id"`
	Spread        int    `pulumi:"spread"`
	Token         string `pulumi:"token"`
	Url           string `pulumi:"url"`
}

// The set of arguments for constructing a AntiAffinityGroup resource.
type AntiAffinityGroupArgs struct {
	Cloudspace_id pulumi.StringInput
	CustomerID    pulumi.StringInput
	Group_id      pulumi.StringInput
	Spread        pulumi.IntInput
	Token         pulumi.StringInput
	Url           pulumi.StringInput
}

func (AntiAffinityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*antiAffinityGroupArgs)(nil)).Elem()
}

type AntiAffinityGroupInput interface {
	pulumi.Input

	ToAntiAffinityGroupOutput() AntiAffinityGroupOutput
	ToAntiAffinityGroupOutputWithContext(ctx context.Context) AntiAffinityGroupOutput
}

func (*AntiAffinityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**AntiAffinityGroup)(nil)).Elem()
}

func (i *AntiAffinityGroup) ToAntiAffinityGroupOutput() AntiAffinityGroupOutput {
	return i.ToAntiAffinityGroupOutputWithContext(context.Background())
}

func (i *AntiAffinityGroup) ToAntiAffinityGroupOutputWithContext(ctx context.Context) AntiAffinityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AntiAffinityGroupOutput)
}

type AntiAffinityGroupOutput struct{ *pulumi.OutputState }

func (AntiAffinityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AntiAffinityGroup)(nil)).Elem()
}

func (o AntiAffinityGroupOutput) ToAntiAffinityGroupOutput() AntiAffinityGroupOutput {
	return o
}

func (o AntiAffinityGroupOutput) ToAntiAffinityGroupOutputWithContext(ctx context.Context) AntiAffinityGroupOutput {
	return o
}

func (o AntiAffinityGroupOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiAffinityGroup) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o AntiAffinityGroupOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiAffinityGroup) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o AntiAffinityGroupOutput) Group_id() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiAffinityGroup) pulumi.StringOutput { return v.Group_id }).(pulumi.StringOutput)
}

func (o AntiAffinityGroupOutput) Spread() pulumi.IntOutput {
	return o.ApplyT(func(v *AntiAffinityGroup) pulumi.IntOutput { return v.Spread }).(pulumi.IntOutput)
}

func (o AntiAffinityGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiAffinityGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o AntiAffinityGroupOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiAffinityGroup) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o AntiAffinityGroupOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *AntiAffinityGroup) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AntiAffinityGroupInput)(nil)).Elem(), &AntiAffinityGroup{})
	pulumi.RegisterOutputType(AntiAffinityGroupOutput{})
}
