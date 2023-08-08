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

type PortForward struct {
	pulumi.CustomResourceState

	Cloudspace_id    pulumi.StringOutput `pulumi:"cloudspace_id"`
	CustomerID       pulumi.StringOutput `pulumi:"customerID"`
	Local_port       pulumi.IntOutput    `pulumi:"local_port"`
	Nested_cs_id     pulumi.StringOutput `pulumi:"nested_cs_id"`
	Portforward_id   pulumi.StringOutput `pulumi:"portforward_id"`
	Protocol         pulumi.StringOutput `pulumi:"protocol"`
	Public_ip        pulumi.StringOutput `pulumi:"public_ip"`
	Public_port      pulumi.IntOutput    `pulumi:"public_port"`
	Till_public_port pulumi.IntOutput    `pulumi:"till_public_port"`
	Token            pulumi.StringOutput `pulumi:"token"`
	Url              pulumi.StringOutput `pulumi:"url"`
	Vm_id            pulumi.IntOutput    `pulumi:"vm_id"`
}

// NewPortForward registers a new resource with the given unique name, arguments, and options.
func NewPortForward(ctx *pulumi.Context,
	name string, args *PortForwardArgs, opts ...pulumi.ResourceOption) (*PortForward, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cloudspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Cloudspace_id'")
	}
	if args.CustomerID == nil {
		return nil, errors.New("invalid value for required argument 'CustomerID'")
	}
	if args.Local_port == nil {
		return nil, errors.New("invalid value for required argument 'Local_port'")
	}
	if args.Nested_cs_id == nil {
		return nil, errors.New("invalid value for required argument 'Nested_cs_id'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.Public_ip == nil {
		return nil, errors.New("invalid value for required argument 'Public_ip'")
	}
	if args.Public_port == nil {
		return nil, errors.New("invalid value for required argument 'Public_port'")
	}
	if args.Till_public_port == nil {
		return nil, errors.New("invalid value for required argument 'Till_public_port'")
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
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PortForward
	err := ctx.RegisterResource("vco:resources:PortForward", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPortForward gets an existing PortForward resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPortForward(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortForwardState, opts ...pulumi.ResourceOption) (*PortForward, error) {
	var resource PortForward
	err := ctx.ReadResource("vco:resources:PortForward", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PortForward resources.
type portForwardState struct {
}

type PortForwardState struct {
}

func (PortForwardState) ElementType() reflect.Type {
	return reflect.TypeOf((*portForwardState)(nil)).Elem()
}

type portForwardArgs struct {
	Cloudspace_id    string `pulumi:"cloudspace_id"`
	CustomerID       string `pulumi:"customerID"`
	Local_port       int    `pulumi:"local_port"`
	Nested_cs_id     string `pulumi:"nested_cs_id"`
	Protocol         string `pulumi:"protocol"`
	Public_ip        string `pulumi:"public_ip"`
	Public_port      int    `pulumi:"public_port"`
	Till_public_port int    `pulumi:"till_public_port"`
	Token            string `pulumi:"token"`
	Url              string `pulumi:"url"`
	Vm_id            int    `pulumi:"vm_id"`
}

// The set of arguments for constructing a PortForward resource.
type PortForwardArgs struct {
	Cloudspace_id    pulumi.StringInput
	CustomerID       pulumi.StringInput
	Local_port       pulumi.IntInput
	Nested_cs_id     pulumi.StringInput
	Protocol         pulumi.StringInput
	Public_ip        pulumi.StringInput
	Public_port      pulumi.IntInput
	Till_public_port pulumi.IntInput
	Token            pulumi.StringInput
	Url              pulumi.StringInput
	Vm_id            pulumi.IntInput
}

func (PortForwardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portForwardArgs)(nil)).Elem()
}

type PortForwardInput interface {
	pulumi.Input

	ToPortForwardOutput() PortForwardOutput
	ToPortForwardOutputWithContext(ctx context.Context) PortForwardOutput
}

func (*PortForward) ElementType() reflect.Type {
	return reflect.TypeOf((**PortForward)(nil)).Elem()
}

func (i *PortForward) ToPortForwardOutput() PortForwardOutput {
	return i.ToPortForwardOutputWithContext(context.Background())
}

func (i *PortForward) ToPortForwardOutputWithContext(ctx context.Context) PortForwardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortForwardOutput)
}

type PortForwardOutput struct{ *pulumi.OutputState }

func (PortForwardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PortForward)(nil)).Elem()
}

func (o PortForwardOutput) ToPortForwardOutput() PortForwardOutput {
	return o
}

func (o PortForwardOutput) ToPortForwardOutputWithContext(ctx context.Context) PortForwardOutput {
	return o
}

func (o PortForwardOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *PortForward) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o PortForwardOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *PortForward) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o PortForwardOutput) Local_port() pulumi.IntOutput {
	return o.ApplyT(func(v *PortForward) pulumi.IntOutput { return v.Local_port }).(pulumi.IntOutput)
}

func (o PortForwardOutput) Nested_cs_id() pulumi.StringOutput {
	return o.ApplyT(func(v *PortForward) pulumi.StringOutput { return v.Nested_cs_id }).(pulumi.StringOutput)
}

func (o PortForwardOutput) Portforward_id() pulumi.StringOutput {
	return o.ApplyT(func(v *PortForward) pulumi.StringOutput { return v.Portforward_id }).(pulumi.StringOutput)
}

func (o PortForwardOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *PortForward) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o PortForwardOutput) Public_ip() pulumi.StringOutput {
	return o.ApplyT(func(v *PortForward) pulumi.StringOutput { return v.Public_ip }).(pulumi.StringOutput)
}

func (o PortForwardOutput) Public_port() pulumi.IntOutput {
	return o.ApplyT(func(v *PortForward) pulumi.IntOutput { return v.Public_port }).(pulumi.IntOutput)
}

func (o PortForwardOutput) Till_public_port() pulumi.IntOutput {
	return o.ApplyT(func(v *PortForward) pulumi.IntOutput { return v.Till_public_port }).(pulumi.IntOutput)
}

func (o PortForwardOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *PortForward) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o PortForwardOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *PortForward) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func (o PortForwardOutput) Vm_id() pulumi.IntOutput {
	return o.ApplyT(func(v *PortForward) pulumi.IntOutput { return v.Vm_id }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PortForwardInput)(nil)).Elem(), &PortForward{})
	pulumi.RegisterOutputType(PortForwardOutput{})
}
