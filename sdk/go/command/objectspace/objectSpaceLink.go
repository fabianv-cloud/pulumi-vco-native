// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package objectspace

import (
	"context"
	"reflect"

	"errors"
	"github.com/fabianv-cloud/pulumi-vco-native/sdk/go/command/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ObjectSpaceLink struct {
	pulumi.CustomResourceState

	Cloudspace_id  pulumi.StringOutput `pulumi:"cloudspace_id"`
	CustomerID     pulumi.StringOutput `pulumi:"customerID"`
	Objectspace_id pulumi.StringOutput `pulumi:"objectspace_id"`
	Success        pulumi.BoolOutput   `pulumi:"success"`
	Token          pulumi.StringOutput `pulumi:"token"`
	Url            pulumi.StringOutput `pulumi:"url"`
}

// NewObjectSpaceLink registers a new resource with the given unique name, arguments, and options.
func NewObjectSpaceLink(ctx *pulumi.Context,
	name string, args *ObjectSpaceLinkArgs, opts ...pulumi.ResourceOption) (*ObjectSpaceLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cloudspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Cloudspace_id'")
	}
	if args.CustomerID == nil {
		return nil, errors.New("invalid value for required argument 'CustomerID'")
	}
	if args.Objectspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Objectspace_id'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ObjectSpaceLink
	err := ctx.RegisterResource("vco:objectspace:ObjectSpaceLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectSpaceLink gets an existing ObjectSpaceLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectSpaceLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectSpaceLinkState, opts ...pulumi.ResourceOption) (*ObjectSpaceLink, error) {
	var resource ObjectSpaceLink
	err := ctx.ReadResource("vco:objectspace:ObjectSpaceLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectSpaceLink resources.
type objectSpaceLinkState struct {
}

type ObjectSpaceLinkState struct {
}

func (ObjectSpaceLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectSpaceLinkState)(nil)).Elem()
}

type objectSpaceLinkArgs struct {
	Cloudspace_id  string `pulumi:"cloudspace_id"`
	CustomerID     string `pulumi:"customerID"`
	Objectspace_id string `pulumi:"objectspace_id"`
	Token          string `pulumi:"token"`
	Url            string `pulumi:"url"`
}

// The set of arguments for constructing a ObjectSpaceLink resource.
type ObjectSpaceLinkArgs struct {
	Cloudspace_id  pulumi.StringInput
	CustomerID     pulumi.StringInput
	Objectspace_id pulumi.StringInput
	Token          pulumi.StringInput
	Url            pulumi.StringInput
}

func (ObjectSpaceLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectSpaceLinkArgs)(nil)).Elem()
}

type ObjectSpaceLinkInput interface {
	pulumi.Input

	ToObjectSpaceLinkOutput() ObjectSpaceLinkOutput
	ToObjectSpaceLinkOutputWithContext(ctx context.Context) ObjectSpaceLinkOutput
}

func (*ObjectSpaceLink) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectSpaceLink)(nil)).Elem()
}

func (i *ObjectSpaceLink) ToObjectSpaceLinkOutput() ObjectSpaceLinkOutput {
	return i.ToObjectSpaceLinkOutputWithContext(context.Background())
}

func (i *ObjectSpaceLink) ToObjectSpaceLinkOutputWithContext(ctx context.Context) ObjectSpaceLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectSpaceLinkOutput)
}

// ObjectSpaceLinkArrayInput is an input type that accepts ObjectSpaceLinkArray and ObjectSpaceLinkArrayOutput values.
// You can construct a concrete instance of `ObjectSpaceLinkArrayInput` via:
//
//	ObjectSpaceLinkArray{ ObjectSpaceLinkArgs{...} }
type ObjectSpaceLinkArrayInput interface {
	pulumi.Input

	ToObjectSpaceLinkArrayOutput() ObjectSpaceLinkArrayOutput
	ToObjectSpaceLinkArrayOutputWithContext(context.Context) ObjectSpaceLinkArrayOutput
}

type ObjectSpaceLinkArray []ObjectSpaceLinkInput

func (ObjectSpaceLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectSpaceLink)(nil)).Elem()
}

func (i ObjectSpaceLinkArray) ToObjectSpaceLinkArrayOutput() ObjectSpaceLinkArrayOutput {
	return i.ToObjectSpaceLinkArrayOutputWithContext(context.Background())
}

func (i ObjectSpaceLinkArray) ToObjectSpaceLinkArrayOutputWithContext(ctx context.Context) ObjectSpaceLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectSpaceLinkArrayOutput)
}

// ObjectSpaceLinkMapInput is an input type that accepts ObjectSpaceLinkMap and ObjectSpaceLinkMapOutput values.
// You can construct a concrete instance of `ObjectSpaceLinkMapInput` via:
//
//	ObjectSpaceLinkMap{ "key": ObjectSpaceLinkArgs{...} }
type ObjectSpaceLinkMapInput interface {
	pulumi.Input

	ToObjectSpaceLinkMapOutput() ObjectSpaceLinkMapOutput
	ToObjectSpaceLinkMapOutputWithContext(context.Context) ObjectSpaceLinkMapOutput
}

type ObjectSpaceLinkMap map[string]ObjectSpaceLinkInput

func (ObjectSpaceLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectSpaceLink)(nil)).Elem()
}

func (i ObjectSpaceLinkMap) ToObjectSpaceLinkMapOutput() ObjectSpaceLinkMapOutput {
	return i.ToObjectSpaceLinkMapOutputWithContext(context.Background())
}

func (i ObjectSpaceLinkMap) ToObjectSpaceLinkMapOutputWithContext(ctx context.Context) ObjectSpaceLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectSpaceLinkMapOutput)
}

type ObjectSpaceLinkOutput struct{ *pulumi.OutputState }

func (ObjectSpaceLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectSpaceLink)(nil)).Elem()
}

func (o ObjectSpaceLinkOutput) ToObjectSpaceLinkOutput() ObjectSpaceLinkOutput {
	return o
}

func (o ObjectSpaceLinkOutput) ToObjectSpaceLinkOutputWithContext(ctx context.Context) ObjectSpaceLinkOutput {
	return o
}

func (o ObjectSpaceLinkOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpaceLink) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o ObjectSpaceLinkOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpaceLink) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o ObjectSpaceLinkOutput) Objectspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpaceLink) pulumi.StringOutput { return v.Objectspace_id }).(pulumi.StringOutput)
}

func (o ObjectSpaceLinkOutput) Success() pulumi.BoolOutput {
	return o.ApplyT(func(v *ObjectSpaceLink) pulumi.BoolOutput { return v.Success }).(pulumi.BoolOutput)
}

func (o ObjectSpaceLinkOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpaceLink) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o ObjectSpaceLinkOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectSpaceLink) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ObjectSpaceLinkArrayOutput struct{ *pulumi.OutputState }

func (ObjectSpaceLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectSpaceLink)(nil)).Elem()
}

func (o ObjectSpaceLinkArrayOutput) ToObjectSpaceLinkArrayOutput() ObjectSpaceLinkArrayOutput {
	return o
}

func (o ObjectSpaceLinkArrayOutput) ToObjectSpaceLinkArrayOutputWithContext(ctx context.Context) ObjectSpaceLinkArrayOutput {
	return o
}

func (o ObjectSpaceLinkArrayOutput) Index(i pulumi.IntInput) ObjectSpaceLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectSpaceLink {
		return vs[0].([]*ObjectSpaceLink)[vs[1].(int)]
	}).(ObjectSpaceLinkOutput)
}

type ObjectSpaceLinkMapOutput struct{ *pulumi.OutputState }

func (ObjectSpaceLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectSpaceLink)(nil)).Elem()
}

func (o ObjectSpaceLinkMapOutput) ToObjectSpaceLinkMapOutput() ObjectSpaceLinkMapOutput {
	return o
}

func (o ObjectSpaceLinkMapOutput) ToObjectSpaceLinkMapOutputWithContext(ctx context.Context) ObjectSpaceLinkMapOutput {
	return o
}

func (o ObjectSpaceLinkMapOutput) MapIndex(k pulumi.StringInput) ObjectSpaceLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectSpaceLink {
		return vs[0].(map[string]*ObjectSpaceLink)[vs[1].(string)]
	}).(ObjectSpaceLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectSpaceLinkInput)(nil)).Elem(), &ObjectSpaceLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectSpaceLinkArrayInput)(nil)).Elem(), ObjectSpaceLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectSpaceLinkMapInput)(nil)).Elem(), ObjectSpaceLinkMap{})
	pulumi.RegisterOutputType(ObjectSpaceLinkOutput{})
	pulumi.RegisterOutputType(ObjectSpaceLinkArrayOutput{})
	pulumi.RegisterOutputType(ObjectSpaceLinkMapOutput{})
}
