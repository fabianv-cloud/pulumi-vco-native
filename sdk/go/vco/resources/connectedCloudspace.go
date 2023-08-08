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

type ConnectedCloudspace struct {
	pulumi.CustomResourceState

	Cloudspace_id           pulumi.StringOutput `pulumi:"cloudspace_id"`
	Connected_cloudspace_id pulumi.StringOutput `pulumi:"connected_cloudspace_id"`
	CustomerID              pulumi.StringOutput `pulumi:"customerID"`
	Local_cloudspace_ip     pulumi.StringOutput `pulumi:"local_cloudspace_ip"`
	Remote_cloudspace_ip    pulumi.StringOutput `pulumi:"remote_cloudspace_ip"`
	Status                  pulumi.StringOutput `pulumi:"status"`
	Token                   pulumi.StringOutput `pulumi:"token"`
	Url                     pulumi.StringOutput `pulumi:"url"`
}

// NewConnectedCloudspace registers a new resource with the given unique name, arguments, and options.
func NewConnectedCloudspace(ctx *pulumi.Context,
	name string, args *ConnectedCloudspaceArgs, opts ...pulumi.ResourceOption) (*ConnectedCloudspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cloudspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Cloudspace_id'")
	}
	if args.Connected_cloudspace_id == nil {
		return nil, errors.New("invalid value for required argument 'Connected_cloudspace_id'")
	}
	if args.CustomerID == nil {
		return nil, errors.New("invalid value for required argument 'CustomerID'")
	}
	if args.Local_cloudspace_ip == nil {
		return nil, errors.New("invalid value for required argument 'Local_cloudspace_ip '")
	}
	if args.Remote_cloudspace_ip == nil {
		return nil, errors.New("invalid value for required argument 'Remote_cloudspace_ip '")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConnectedCloudspace
	err := ctx.RegisterResource("vco:resources:ConnectedCloudspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectedCloudspace gets an existing ConnectedCloudspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectedCloudspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectedCloudspaceState, opts ...pulumi.ResourceOption) (*ConnectedCloudspace, error) {
	var resource ConnectedCloudspace
	err := ctx.ReadResource("vco:resources:ConnectedCloudspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectedCloudspace resources.
type connectedCloudspaceState struct {
}

type ConnectedCloudspaceState struct {
}

func (ConnectedCloudspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedCloudspaceState)(nil)).Elem()
}

type connectedCloudspaceArgs struct {
	Cloudspace_id           string `pulumi:"cloudspace_id"`
	Connected_cloudspace_id string `pulumi:"connected_cloudspace_id"`
	CustomerID              string `pulumi:"customerID"`
	Local_cloudspace_ip     string `pulumi:"local_cloudspace_ip "`
	Remote_cloudspace_ip    string `pulumi:"remote_cloudspace_ip "`
	Token                   string `pulumi:"token"`
	Url                     string `pulumi:"url"`
}

// The set of arguments for constructing a ConnectedCloudspace resource.
type ConnectedCloudspaceArgs struct {
	Cloudspace_id           pulumi.StringInput
	Connected_cloudspace_id pulumi.StringInput
	CustomerID              pulumi.StringInput
	Local_cloudspace_ip     pulumi.StringInput
	Remote_cloudspace_ip    pulumi.StringInput
	Token                   pulumi.StringInput
	Url                     pulumi.StringInput
}

func (ConnectedCloudspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedCloudspaceArgs)(nil)).Elem()
}

type ConnectedCloudspaceInput interface {
	pulumi.Input

	ToConnectedCloudspaceOutput() ConnectedCloudspaceOutput
	ToConnectedCloudspaceOutputWithContext(ctx context.Context) ConnectedCloudspaceOutput
}

func (*ConnectedCloudspace) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedCloudspace)(nil)).Elem()
}

func (i *ConnectedCloudspace) ToConnectedCloudspaceOutput() ConnectedCloudspaceOutput {
	return i.ToConnectedCloudspaceOutputWithContext(context.Background())
}

func (i *ConnectedCloudspace) ToConnectedCloudspaceOutputWithContext(ctx context.Context) ConnectedCloudspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedCloudspaceOutput)
}

type ConnectedCloudspaceOutput struct{ *pulumi.OutputState }

func (ConnectedCloudspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedCloudspace)(nil)).Elem()
}

func (o ConnectedCloudspaceOutput) ToConnectedCloudspaceOutput() ConnectedCloudspaceOutput {
	return o
}

func (o ConnectedCloudspaceOutput) ToConnectedCloudspaceOutputWithContext(ctx context.Context) ConnectedCloudspaceOutput {
	return o
}

func (o ConnectedCloudspaceOutput) Cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCloudspace) pulumi.StringOutput { return v.Cloudspace_id }).(pulumi.StringOutput)
}

func (o ConnectedCloudspaceOutput) Connected_cloudspace_id() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCloudspace) pulumi.StringOutput { return v.Connected_cloudspace_id }).(pulumi.StringOutput)
}

func (o ConnectedCloudspaceOutput) CustomerID() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCloudspace) pulumi.StringOutput { return v.CustomerID }).(pulumi.StringOutput)
}

func (o ConnectedCloudspaceOutput) Local_cloudspace_ip() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCloudspace) pulumi.StringOutput { return v.Local_cloudspace_ip }).(pulumi.StringOutput)
}

func (o ConnectedCloudspaceOutput) Remote_cloudspace_ip() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCloudspace) pulumi.StringOutput { return v.Remote_cloudspace_ip }).(pulumi.StringOutput)
}

func (o ConnectedCloudspaceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCloudspace) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o ConnectedCloudspaceOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCloudspace) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func (o ConnectedCloudspaceOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedCloudspace) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectedCloudspaceInput)(nil)).Elem(), &ConnectedCloudspace{})
	pulumi.RegisterOutputType(ConnectedCloudspaceOutput{})
}
