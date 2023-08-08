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

type Host struct {
	pulumi.CustomResourceState

	Address       pulumi.StringOutput `pulumi:"address"`
	Cloudspace_id pulumi.StringOutput `pulumi:"cloudspace_id"`
	CustomerID    pulumi.StringOutput `pulumi:"customerID"`
	Host_id       pulumi.StringOutput `pulumi:"host_id"`
	Serverpool_id pulumi.StringOutput `pulumi:"serverpool_id"`
	Token         pulumi.StringOutput `pulumi:"token"`
	Url           pulumi.StringOutput `pulumi:"url"`
}

// NewHost registers a new resource with the given unique name, arguments, and options.
func NewHost(ctx *pulumi.Context,
	name string, args *HostArgs, opts ...pulumi.ResourceOption) (*Host, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.Cloudspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Cloudspace_id'")
	}
	if args.CustomerID == nil {
		return nil, errors.New("invalid value for required argument 'CustomerID'")
	}
	if args.Serverpool_id == nil {
		return nil, errors.New("invalid value for required argument 'Serverpool_id'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Host
	err := ctx.RegisterResource("vco:resources:Host", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHost gets an existing Host resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostState, opts ...pulumi.ResourceOption) (*Host, error) {
	var resource Host
	err := ctx.ReadResource("vco:resources:Host", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Host resources.
type hostState struct {
}

type HostState struct {
}

func (HostState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostState)(nil)).Elem()
}

type hostArgs struct {
	Address       string `pulumi:"address"`
	Cloudspace_id string `pulumi:"cloudspace_id"`
	CustomerID    string `pulumi:"customerID"`
	Serverpool_id string `pulumi:"serverpool_id"`
	Token         string `pulumi:"token"`
	Url           string `pulumi:"url"`
}

// The set of arguments for constructing a Host resource.
type HostArgs struct {
	Address       pulumi.StringInput
	Cloudspace_id pulumi.StringInput
	CustomerID    pulumi.StringInput
	Serverpool_id pulumi.StringInput
	Token         pulumi.StringInput
	Url           pulumi.StringInput
}

func (HostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostArgs)(nil)).Elem()
}

type HostInput interface {
	pulumi.Input

	ToHostOutput() HostOutput
	ToHostOutputWithContext(ctx context.Context) HostOutput
}

func (*Host) ElementType() reflect.Type {
	return reflect.TypeOf((**Host)(nil)).Elem()
}

func (i *Host) ToHostOutput() HostOutput {
	return i.ToHostOutputWithContext(context.Background())
}

func (i *Host) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostOutput)
}

type HostOutput struct{ *pulumi.OutputState }

func (HostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Host)(nil)).Elem()
}

func (o HostOutput) ToHostOutput() HostOutput {
	return o
}

func (o HostOutput) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return o
}

func (o HostOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

func (o HostOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o HostOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o HostOutput) Host_id() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Host_id }).(pulumi.StringOutput)
}

func (o HostOutput) Serverpool_id() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Serverpool_id }).(pulumi.StringOutput)
}

func (o HostOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o HostOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Host) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostInput)(nil)).Elem(), &Host{})
	pulumi.RegisterOutputType(HostOutput{})
}
