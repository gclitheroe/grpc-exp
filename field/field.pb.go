// Code generated by protoc-gen-go.
// source: field.proto
// DO NOT EDIT!

/*
Package field is a generated protocol buffer package.

It is generated from these files:
	field.proto

It has these top-level messages:
	Result
	Device
*/
package field

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Result struct {
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Device - metrics are collected for a device.
type Device struct {
	// DeviceID e.g., idu-birchfarm.
	DeviceID string `protobuf:"bytes,1,opt,name=device_iD,json=deviceID" json:"device_iD,omitempty"`
	// ModelID for the device e.g., Trimble-NetR9.
	ModelID string `protobuf:"bytes,2,opt,name=model_iD,json=modelID" json:"model_iD,omitempty"`
	// Latitude and Longitude, only uses three digits of precision after decimal.
	Latitude  float32 `protobuf:"fixed32,3,opt,name=latitude" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,4,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Result)(nil), "field.Result")
	proto.RegisterType((*Device)(nil), "field.Device")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Field service

type FieldClient interface {
	// DeviceSave creates or updates the Device.
	DeviceSave(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Result, error)
}

type fieldClient struct {
	cc *grpc.ClientConn
}

func NewFieldClient(cc *grpc.ClientConn) FieldClient {
	return &fieldClient{cc}
}

func (c *fieldClient) DeviceSave(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/field.Field/DeviceSave", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Field service

type FieldServer interface {
	// DeviceSave creates or updates the Device.
	DeviceSave(context.Context, *Device) (*Result, error)
}

func RegisterFieldServer(s *grpc.Server, srv FieldServer) {
	s.RegisterService(&_Field_serviceDesc, srv)
}

func _Field_DeviceSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FieldServer).DeviceSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/field.Field/DeviceSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FieldServer).DeviceSave(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

var _Field_serviceDesc = grpc.ServiceDesc{
	ServiceName: "field.Field",
	HandlerType: (*FieldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeviceSave",
			Handler:    _Field_DeviceSave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("field.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcb, 0x4c, 0xcd,
	0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x38, 0xb8, 0xd8, 0x82,
	0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x94, 0xaa, 0xb8, 0xd8, 0x5c, 0x52, 0xcb, 0x32, 0x93, 0x53, 0x85,
	0xa4, 0xb9, 0x38, 0x53, 0xc0, 0xac, 0xf8, 0x4c, 0x17, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x0e, 0x88, 0x80, 0xa7, 0x8b, 0x90, 0x24, 0x17, 0x47, 0x6e, 0x7e, 0x4a, 0x6a, 0x0e, 0x48, 0x8e,
	0x09, 0x2c, 0xc7, 0x0e, 0xe6, 0x03, 0xa5, 0xa4, 0xb8, 0x38, 0x72, 0x12, 0x4b, 0x32, 0x4b, 0x4a,
	0x53, 0x52, 0x25, 0x98, 0x81, 0x52, 0x4c, 0x41, 0x70, 0xbe, 0x90, 0x0c, 0x17, 0x67, 0x4e, 0x7e,
	0x5e, 0x3a, 0x44, 0x92, 0x05, 0x2c, 0x89, 0x10, 0x30, 0x32, 0xe6, 0x62, 0x75, 0x03, 0x39, 0x47,
	0x48, 0x8b, 0x8b, 0x0b, 0xe2, 0x88, 0xe0, 0xc4, 0xb2, 0x54, 0x21, 0x5e, 0x3d, 0x88, 0x8b, 0x21,
	0x42, 0x52, 0x30, 0x2e, 0xc4, 0xc1, 0x4e, 0xec, 0x51, 0x10, 0x3f, 0x24, 0xb1, 0x81, 0x7d, 0x64,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x01, 0x70, 0x34, 0xd5, 0xe0, 0x00, 0x00, 0x00,
}
