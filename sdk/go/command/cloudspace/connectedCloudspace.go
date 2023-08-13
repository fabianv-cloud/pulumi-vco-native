// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudspace

import (
	"context"
	"reflect"

	"errors"
	"github.com/fabianv-cloud/pulumi-vco-native/sdk/go/command/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
	var resource ConnectedCloudspace
	err := ctx.RegisterResource("vco:cloudspace:ConnectedCloudspace", name, args, &resource, opts...)
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
	err := ctx.ReadResource("vco:cloudspace:ConnectedCloudspace", name, id, state, &resource, opts...)
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

// ConnectedCloudspaceArrayInput is an input type that accepts ConnectedCloudspaceArray and ConnectedCloudspaceArrayOutput values.
// You can construct a concrete instance of `ConnectedCloudspaceArrayInput` via:
//
//	ConnectedCloudspaceArray{ ConnectedCloudspaceArgs{...} }
type ConnectedCloudspaceArrayInput interface {
	pulumi.Input

	ToConnectedCloudspaceArrayOutput() ConnectedCloudspaceArrayOutput
	ToConnectedCloudspaceArrayOutputWithContext(context.Context) ConnectedCloudspaceArrayOutput
}

type ConnectedCloudspaceArray []ConnectedCloudspaceInput

func (ConnectedCloudspaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConnectedCloudspace)(nil)).Elem()
}

func (i ConnectedCloudspaceArray) ToConnectedCloudspaceArrayOutput() ConnectedCloudspaceArrayOutput {
	return i.ToConnectedCloudspaceArrayOutputWithContext(context.Background())
}

func (i ConnectedCloudspaceArray) ToConnectedCloudspaceArrayOutputWithContext(ctx context.Context) ConnectedCloudspaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedCloudspaceArrayOutput)
}

// ConnectedCloudspaceMapInput is an input type that accepts ConnectedCloudspaceMap and ConnectedCloudspaceMapOutput values.
// You can construct a concrete instance of `ConnectedCloudspaceMapInput` via:
//
//	ConnectedCloudspaceMap{ "key": ConnectedCloudspaceArgs{...} }
type ConnectedCloudspaceMapInput interface {
	pulumi.Input

	ToConnectedCloudspaceMapOutput() ConnectedCloudspaceMapOutput
	ToConnectedCloudspaceMapOutputWithContext(context.Context) ConnectedCloudspaceMapOutput
}

type ConnectedCloudspaceMap map[string]ConnectedCloudspaceInput

func (ConnectedCloudspaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConnectedCloudspace)(nil)).Elem()
}

func (i ConnectedCloudspaceMap) ToConnectedCloudspaceMapOutput() ConnectedCloudspaceMapOutput {
	return i.ToConnectedCloudspaceMapOutputWithContext(context.Background())
}

func (i ConnectedCloudspaceMap) ToConnectedCloudspaceMapOutputWithContext(ctx context.Context) ConnectedCloudspaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedCloudspaceMapOutput)
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

type ConnectedCloudspaceArrayOutput struct{ *pulumi.OutputState }

func (ConnectedCloudspaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConnectedCloudspace)(nil)).Elem()
}

func (o ConnectedCloudspaceArrayOutput) ToConnectedCloudspaceArrayOutput() ConnectedCloudspaceArrayOutput {
	return o
}

func (o ConnectedCloudspaceArrayOutput) ToConnectedCloudspaceArrayOutputWithContext(ctx context.Context) ConnectedCloudspaceArrayOutput {
	return o
}

func (o ConnectedCloudspaceArrayOutput) Index(i pulumi.IntInput) ConnectedCloudspaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConnectedCloudspace {
		return vs[0].([]*ConnectedCloudspace)[vs[1].(int)]
	}).(ConnectedCloudspaceOutput)
}

type ConnectedCloudspaceMapOutput struct{ *pulumi.OutputState }

func (ConnectedCloudspaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConnectedCloudspace)(nil)).Elem()
}

func (o ConnectedCloudspaceMapOutput) ToConnectedCloudspaceMapOutput() ConnectedCloudspaceMapOutput {
	return o
}

func (o ConnectedCloudspaceMapOutput) ToConnectedCloudspaceMapOutputWithContext(ctx context.Context) ConnectedCloudspaceMapOutput {
	return o
}

func (o ConnectedCloudspaceMapOutput) MapIndex(k pulumi.StringInput) ConnectedCloudspaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConnectedCloudspace {
		return vs[0].(map[string]*ConnectedCloudspace)[vs[1].(string)]
	}).(ConnectedCloudspaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectedCloudspaceInput)(nil)).Elem(), &ConnectedCloudspace{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectedCloudspaceArrayInput)(nil)).Elem(), ConnectedCloudspaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectedCloudspaceMapInput)(nil)).Elem(), ConnectedCloudspaceMap{})
	pulumi.RegisterOutputType(ConnectedCloudspaceOutput{})
	pulumi.RegisterOutputType(ConnectedCloudspaceArrayOutput{})
	pulumi.RegisterOutputType(ConnectedCloudspaceMapOutput{})
}
