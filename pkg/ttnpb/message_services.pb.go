// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/message_services.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ProcessUplinkMessageRequest struct {
	EndDeviceIdentifiers `protobuf:"bytes,1,opt,name=ids,proto3,embedded=ids" json:"ids"`
	EndDeviceVersionIDs  EndDeviceVersionIdentifiers `protobuf:"bytes,2,opt,name=end_device_version_ids,json=endDeviceVersionIds,proto3" json:"end_device_version_ids"`
	Message              ApplicationUplink           `protobuf:"bytes,3,opt,name=message,proto3" json:"message"`
	Parameter            string                      `protobuf:"bytes,4,opt,name=parameter,proto3" json:"parameter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ProcessUplinkMessageRequest) Reset()      { *m = ProcessUplinkMessageRequest{} }
func (*ProcessUplinkMessageRequest) ProtoMessage() {}
func (*ProcessUplinkMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa32647f6f069ac, []int{0}
}
func (m *ProcessUplinkMessageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProcessUplinkMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProcessUplinkMessageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProcessUplinkMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessUplinkMessageRequest.Merge(m, src)
}
func (m *ProcessUplinkMessageRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProcessUplinkMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessUplinkMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessUplinkMessageRequest proto.InternalMessageInfo

func (m *ProcessUplinkMessageRequest) GetEndDeviceVersionIDs() EndDeviceVersionIdentifiers {
	if m != nil {
		return m.EndDeviceVersionIDs
	}
	return EndDeviceVersionIdentifiers{}
}

func (m *ProcessUplinkMessageRequest) GetMessage() ApplicationUplink {
	if m != nil {
		return m.Message
	}
	return ApplicationUplink{}
}

func (m *ProcessUplinkMessageRequest) GetParameter() string {
	if m != nil {
		return m.Parameter
	}
	return ""
}

type ProcessDownlinkMessageRequest struct {
	EndDeviceIdentifiers `protobuf:"bytes,1,opt,name=ids,proto3,embedded=ids" json:"ids"`
	EndDeviceVersionIDs  EndDeviceVersionIdentifiers `protobuf:"bytes,2,opt,name=end_device_version_ids,json=endDeviceVersionIds,proto3" json:"end_device_version_ids"`
	Message              ApplicationDownlink         `protobuf:"bytes,3,opt,name=message,proto3" json:"message"`
	Parameter            string                      `protobuf:"bytes,4,opt,name=parameter,proto3" json:"parameter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ProcessDownlinkMessageRequest) Reset()      { *m = ProcessDownlinkMessageRequest{} }
func (*ProcessDownlinkMessageRequest) ProtoMessage() {}
func (*ProcessDownlinkMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fa32647f6f069ac, []int{1}
}
func (m *ProcessDownlinkMessageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProcessDownlinkMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProcessDownlinkMessageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProcessDownlinkMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessDownlinkMessageRequest.Merge(m, src)
}
func (m *ProcessDownlinkMessageRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProcessDownlinkMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessDownlinkMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessDownlinkMessageRequest proto.InternalMessageInfo

func (m *ProcessDownlinkMessageRequest) GetEndDeviceVersionIDs() EndDeviceVersionIdentifiers {
	if m != nil {
		return m.EndDeviceVersionIDs
	}
	return EndDeviceVersionIdentifiers{}
}

func (m *ProcessDownlinkMessageRequest) GetMessage() ApplicationDownlink {
	if m != nil {
		return m.Message
	}
	return ApplicationDownlink{}
}

func (m *ProcessDownlinkMessageRequest) GetParameter() string {
	if m != nil {
		return m.Parameter
	}
	return ""
}

func init() {
	proto.RegisterType((*ProcessUplinkMessageRequest)(nil), "ttn.lorawan.v3.ProcessUplinkMessageRequest")
	golang_proto.RegisterType((*ProcessUplinkMessageRequest)(nil), "ttn.lorawan.v3.ProcessUplinkMessageRequest")
	proto.RegisterType((*ProcessDownlinkMessageRequest)(nil), "ttn.lorawan.v3.ProcessDownlinkMessageRequest")
	golang_proto.RegisterType((*ProcessDownlinkMessageRequest)(nil), "ttn.lorawan.v3.ProcessDownlinkMessageRequest")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/message_services.proto", fileDescriptor_6fa32647f6f069ac)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/message_services.proto", fileDescriptor_6fa32647f6f069ac)
}

var fileDescriptor_6fa32647f6f069ac = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x94, 0xcf, 0x6b, 0x13, 0x4f,
	0x18, 0xc6, 0x67, 0xd2, 0xef, 0xd7, 0xd2, 0x55, 0xa4, 0x6c, 0xa1, 0x84, 0xa8, 0x6f, 0x63, 0xe2,
	0x21, 0xa0, 0xd9, 0x85, 0xf4, 0x2f, 0x68, 0x88, 0x60, 0x11, 0x41, 0x02, 0x0a, 0x0a, 0x12, 0x36,
	0xbb, 0xd3, 0xcd, 0x90, 0x64, 0x66, 0xdd, 0x99, 0x24, 0xe6, 0xd6, 0x8b, 0x50, 0x3c, 0x79, 0xf4,
	0x28, 0x9e, 0x7a, 0xec, 0xb1, 0xc7, 0x1e, 0x73, 0xcc, 0x31, 0x07, 0x09, 0xdd, 0xd9, 0x4b, 0x8f,
	0x3d, 0x16, 0x4f, 0x92, 0xcd, 0x96, 0xfc, 0x58, 0x53, 0xc5, 0xab, 0xb7, 0x79, 0x93, 0x67, 0x9e,
	0xe7, 0x7d, 0x3f, 0xfb, 0x32, 0x5a, 0xa1, 0xc5, 0x7d, 0xab, 0x67, 0xb1, 0xa2, 0x90, 0x96, 0xdd,
	0x34, 0x2d, 0x8f, 0x9a, 0x6d, 0x22, 0x84, 0xe5, 0x92, 0x9a, 0x20, 0x7e, 0x97, 0xda, 0x44, 0x18,
	0x9e, 0xcf, 0x25, 0xd7, 0xef, 0x4a, 0xc9, 0x8c, 0x58, 0x6d, 0x74, 0x77, 0x33, 0x7b, 0x2e, 0x95,
	0x8d, 0x4e, 0xdd, 0xb0, 0x79, 0xdb, 0x24, 0xac, 0xcb, 0xfb, 0x9e, 0xcf, 0x3f, 0xf4, 0xcd, 0x48,
	0x6c, 0x17, 0x5d, 0xc2, 0x8a, 0x5d, 0xab, 0x45, 0x1d, 0x4b, 0x12, 0x33, 0x71, 0x98, 0x5a, 0x66,
	0x8a, 0x73, 0x16, 0x2e, 0x77, 0xf9, 0xf4, 0x72, 0xbd, 0x73, 0x10, 0x55, 0x51, 0x11, 0x9d, 0x62,
	0x79, 0x2e, 0xd9, 0x2b, 0x61, 0x4e, 0xcd, 0x21, 0x93, 0x36, 0x63, 0x4d, 0x3e, 0xa9, 0xa1, 0x0e,
	0x61, 0x92, 0x1e, 0x50, 0xe2, 0xc7, 0xa3, 0x64, 0xb2, 0x2b, 0x87, 0x8e, 0x15, 0xb9, 0x51, 0x4a,
	0xbb, 0xf7, 0xd2, 0xe7, 0x36, 0x11, 0xe2, 0x95, 0xd7, 0xa2, 0xac, 0xf9, 0x62, 0xfa, 0x7f, 0x95,
	0xbc, 0xef, 0x10, 0x21, 0xf5, 0x67, 0xda, 0x1a, 0x75, 0x44, 0x1a, 0x67, 0x71, 0xe1, 0x76, 0xe9,
	0x91, 0xb1, 0x88, 0xc6, 0x78, 0xca, 0x9c, 0x4a, 0xd4, 0xd4, 0xfe, 0x2c, 0xba, 0xbc, 0xf9, 0xa3,
	0xfc, 0xff, 0x27, 0x9c, 0xda, 0xc4, 0x83, 0xf1, 0x0e, 0x1a, 0x8e, 0x77, 0x70, 0x75, 0x62, 0xa1,
	0x7f, 0xc4, 0xda, 0xf6, 0x6c, 0x8a, 0x5a, 0x97, 0xf8, 0x82, 0x72, 0x56, 0x9b, 0xb8, 0xa7, 0x22,
	0xf7, 0xc7, 0x2b, 0xdd, 0x5f, 0x4f, 0xb5, 0xf3, 0x21, 0xf9, 0xf9, 0x10, 0x35, 0xde, 0xd9, 0x4a,
	0x88, 0x2b, 0xa2, 0xba, 0x45, 0x12, 0x0e, 0x42, 0xdf, 0xd7, 0xd6, 0x63, 0x06, 0xe9, 0xb5, 0x28,
	0xf7, 0xe1, 0x72, 0xee, 0x9e, 0xe7, 0xb5, 0xa8, 0x6d, 0x49, 0xca, 0xd9, 0x94, 0x49, 0xf9, 0xce,
	0x7c, 0x5a, 0xf5, 0xfa, 0xbe, 0x7e, 0x5f, 0xdb, 0xf0, 0x2c, 0xdf, 0x6a, 0x13, 0x49, 0xfc, 0xf4,
	0x7f, 0x59, 0x5c, 0xd8, 0xa8, 0xce, 0x7e, 0xc8, 0x7d, 0x4f, 0x69, 0x0f, 0x62, 0xb4, 0x15, 0xde,
	0x63, 0xff, 0x04, 0xdc, 0xe7, 0xcb, 0x70, 0xf3, 0x37, 0xc0, 0xbd, 0xa6, 0xf2, 0x57, 0x78, 0x4b,
	0x42, 0xdb, 0x5e, 0xd8, 0xd8, 0x18, 0x35, 0xf7, 0xf5, 0x37, 0xda, 0x7a, 0x5c, 0xe8, 0x89, 0xb1,
	0x6f, 0xd8, 0xf5, 0xcc, 0xef, 0x17, 0xa1, 0xd4, 0xd7, 0xd2, 0x4b, 0xdf, 0x72, 0x16, 0xfb, 0x6e,
	0x16, 0x5b, 0x5c, 0x11, 0xfb, 0xeb, 0x3d, 0xc8, 0xfc, 0x11, 0xa4, 0x6f, 0x78, 0x10, 0x00, 0x1e,
	0x06, 0x80, 0x47, 0x01, 0xa0, 0xf3, 0x00, 0xd0, 0x45, 0x00, 0xe8, 0x32, 0x00, 0x74, 0x15, 0x00,
	0x3e, 0x54, 0x80, 0x8f, 0x14, 0xa0, 0x63, 0x05, 0xf8, 0x44, 0x01, 0x3a, 0x55, 0x80, 0xce, 0x14,
	0xa0, 0x81, 0x02, 0x3c, 0x54, 0x80, 0x47, 0x0a, 0xd0, 0xb9, 0x02, 0x7c, 0xa1, 0x00, 0x5d, 0x2a,
	0xc0, 0x57, 0x0a, 0xd0, 0x61, 0x08, 0xe8, 0x28, 0x04, 0xfc, 0x39, 0x04, 0xf4, 0x25, 0x04, 0xfc,
	0x35, 0x04, 0x74, 0x1c, 0x02, 0x3a, 0x09, 0x01, 0x9f, 0x86, 0x80, 0xcf, 0x42, 0xc0, 0x6f, 0x9f,
	0xb8, 0xdc, 0x90, 0x0d, 0x22, 0x1b, 0x94, 0xb9, 0xc2, 0x60, 0x44, 0xf6, 0xb8, 0xdf, 0x34, 0x17,
	0xdf, 0x15, 0xaf, 0xe9, 0x9a, 0x52, 0x32, 0xaf, 0x5e, 0xbf, 0x15, 0xbd, 0x2a, 0xbb, 0x3f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xb7, 0x2a, 0x85, 0x19, 0x6e, 0x05, 0x00, 0x00,
}

func (this *ProcessUplinkMessageRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProcessUplinkMessageRequest)
	if !ok {
		that2, ok := that.(ProcessUplinkMessageRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.EndDeviceIdentifiers.Equal(&that1.EndDeviceIdentifiers) {
		return false
	}
	if !this.EndDeviceVersionIDs.Equal(&that1.EndDeviceVersionIDs) {
		return false
	}
	if !this.Message.Equal(&that1.Message) {
		return false
	}
	if this.Parameter != that1.Parameter {
		return false
	}
	return true
}
func (this *ProcessDownlinkMessageRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProcessDownlinkMessageRequest)
	if !ok {
		that2, ok := that.(ProcessDownlinkMessageRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.EndDeviceIdentifiers.Equal(&that1.EndDeviceIdentifiers) {
		return false
	}
	if !this.EndDeviceVersionIDs.Equal(&that1.EndDeviceVersionIDs) {
		return false
	}
	if !this.Message.Equal(&that1.Message) {
		return false
	}
	if this.Parameter != that1.Parameter {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UplinkMessageProcessorClient is the client API for UplinkMessageProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UplinkMessageProcessorClient interface {
	Process(ctx context.Context, in *ProcessUplinkMessageRequest, opts ...grpc.CallOption) (*ApplicationUplink, error)
}

type uplinkMessageProcessorClient struct {
	cc *grpc.ClientConn
}

func NewUplinkMessageProcessorClient(cc *grpc.ClientConn) UplinkMessageProcessorClient {
	return &uplinkMessageProcessorClient{cc}
}

func (c *uplinkMessageProcessorClient) Process(ctx context.Context, in *ProcessUplinkMessageRequest, opts ...grpc.CallOption) (*ApplicationUplink, error) {
	out := new(ApplicationUplink)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UplinkMessageProcessor/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UplinkMessageProcessorServer is the server API for UplinkMessageProcessor service.
type UplinkMessageProcessorServer interface {
	Process(context.Context, *ProcessUplinkMessageRequest) (*ApplicationUplink, error)
}

// UnimplementedUplinkMessageProcessorServer can be embedded to have forward compatible implementations.
type UnimplementedUplinkMessageProcessorServer struct {
}

func (*UnimplementedUplinkMessageProcessorServer) Process(ctx context.Context, req *ProcessUplinkMessageRequest) (*ApplicationUplink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}

func RegisterUplinkMessageProcessorServer(s *grpc.Server, srv UplinkMessageProcessorServer) {
	s.RegisterService(&_UplinkMessageProcessor_serviceDesc, srv)
}

func _UplinkMessageProcessor_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessUplinkMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UplinkMessageProcessorServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UplinkMessageProcessor/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UplinkMessageProcessorServer).Process(ctx, req.(*ProcessUplinkMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UplinkMessageProcessor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.UplinkMessageProcessor",
	HandlerType: (*UplinkMessageProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _UplinkMessageProcessor_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/message_services.proto",
}

// DownlinkMessageProcessorClient is the client API for DownlinkMessageProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DownlinkMessageProcessorClient interface {
	Process(ctx context.Context, in *ProcessDownlinkMessageRequest, opts ...grpc.CallOption) (*ApplicationDownlink, error)
}

type downlinkMessageProcessorClient struct {
	cc *grpc.ClientConn
}

func NewDownlinkMessageProcessorClient(cc *grpc.ClientConn) DownlinkMessageProcessorClient {
	return &downlinkMessageProcessorClient{cc}
}

func (c *downlinkMessageProcessorClient) Process(ctx context.Context, in *ProcessDownlinkMessageRequest, opts ...grpc.CallOption) (*ApplicationDownlink, error) {
	out := new(ApplicationDownlink)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.DownlinkMessageProcessor/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DownlinkMessageProcessorServer is the server API for DownlinkMessageProcessor service.
type DownlinkMessageProcessorServer interface {
	Process(context.Context, *ProcessDownlinkMessageRequest) (*ApplicationDownlink, error)
}

// UnimplementedDownlinkMessageProcessorServer can be embedded to have forward compatible implementations.
type UnimplementedDownlinkMessageProcessorServer struct {
}

func (*UnimplementedDownlinkMessageProcessorServer) Process(ctx context.Context, req *ProcessDownlinkMessageRequest) (*ApplicationDownlink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}

func RegisterDownlinkMessageProcessorServer(s *grpc.Server, srv DownlinkMessageProcessorServer) {
	s.RegisterService(&_DownlinkMessageProcessor_serviceDesc, srv)
}

func _DownlinkMessageProcessor_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessDownlinkMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownlinkMessageProcessorServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.DownlinkMessageProcessor/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownlinkMessageProcessorServer).Process(ctx, req.(*ProcessDownlinkMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DownlinkMessageProcessor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.DownlinkMessageProcessor",
	HandlerType: (*DownlinkMessageProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _DownlinkMessageProcessor_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/message_services.proto",
}

func (m *ProcessUplinkMessageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessUplinkMessageRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProcessUplinkMessageRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Parameter) > 0 {
		i -= len(m.Parameter)
		copy(dAtA[i:], m.Parameter)
		i = encodeVarintMessageServices(dAtA, i, uint64(len(m.Parameter)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Message.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMessageServices(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.EndDeviceVersionIDs.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMessageServices(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.EndDeviceIdentifiers.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMessageServices(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ProcessDownlinkMessageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessDownlinkMessageRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProcessDownlinkMessageRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Parameter) > 0 {
		i -= len(m.Parameter)
		copy(dAtA[i:], m.Parameter)
		i = encodeVarintMessageServices(dAtA, i, uint64(len(m.Parameter)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Message.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMessageServices(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.EndDeviceVersionIDs.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMessageServices(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.EndDeviceIdentifiers.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMessageServices(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMessageServices(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessageServices(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedProcessUplinkMessageRequest(r randyMessageServices, easy bool) *ProcessUplinkMessageRequest {
	this := &ProcessUplinkMessageRequest{}
	v1 := NewPopulatedEndDeviceIdentifiers(r, easy)
	this.EndDeviceIdentifiers = *v1
	v2 := NewPopulatedEndDeviceVersionIdentifiers(r, easy)
	this.EndDeviceVersionIDs = *v2
	v3 := NewPopulatedApplicationUplink(r, easy)
	this.Message = *v3
	this.Parameter = randStringMessageServices(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedProcessDownlinkMessageRequest(r randyMessageServices, easy bool) *ProcessDownlinkMessageRequest {
	this := &ProcessDownlinkMessageRequest{}
	v4 := NewPopulatedEndDeviceIdentifiers(r, easy)
	this.EndDeviceIdentifiers = *v4
	v5 := NewPopulatedEndDeviceVersionIdentifiers(r, easy)
	this.EndDeviceVersionIDs = *v5
	v6 := NewPopulatedApplicationDownlink(r, easy)
	this.Message = *v6
	this.Parameter = randStringMessageServices(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyMessageServices interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMessageServices(r randyMessageServices) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMessageServices(r randyMessageServices) string {
	v7 := r.Intn(100)
	tmps := make([]rune, v7)
	for i := 0; i < v7; i++ {
		tmps[i] = randUTF8RuneMessageServices(r)
	}
	return string(tmps)
}
func randUnrecognizedMessageServices(r randyMessageServices, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMessageServices(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMessageServices(dAtA []byte, r randyMessageServices, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMessageServices(dAtA, uint64(key))
		v8 := r.Int63()
		if r.Intn(2) == 0 {
			v8 *= -1
		}
		dAtA = encodeVarintPopulateMessageServices(dAtA, uint64(v8))
	case 1:
		dAtA = encodeVarintPopulateMessageServices(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMessageServices(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMessageServices(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMessageServices(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMessageServices(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ProcessUplinkMessageRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EndDeviceIdentifiers.Size()
	n += 1 + l + sovMessageServices(uint64(l))
	l = m.EndDeviceVersionIDs.Size()
	n += 1 + l + sovMessageServices(uint64(l))
	l = m.Message.Size()
	n += 1 + l + sovMessageServices(uint64(l))
	l = len(m.Parameter)
	if l > 0 {
		n += 1 + l + sovMessageServices(uint64(l))
	}
	return n
}

func (m *ProcessDownlinkMessageRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EndDeviceIdentifiers.Size()
	n += 1 + l + sovMessageServices(uint64(l))
	l = m.EndDeviceVersionIDs.Size()
	n += 1 + l + sovMessageServices(uint64(l))
	l = m.Message.Size()
	n += 1 + l + sovMessageServices(uint64(l))
	l = len(m.Parameter)
	if l > 0 {
		n += 1 + l + sovMessageServices(uint64(l))
	}
	return n
}

func sovMessageServices(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessageServices(x uint64) (n int) {
	return sovMessageServices((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *ProcessUplinkMessageRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessUplinkMessageRequest{`,
		`EndDeviceIdentifiers:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.EndDeviceIdentifiers), "EndDeviceIdentifiers", "EndDeviceIdentifiers", 1), `&`, ``, 1) + `,`,
		`EndDeviceVersionIDs:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.EndDeviceVersionIDs), "EndDeviceVersionIdentifiers", "EndDeviceVersionIdentifiers", 1), `&`, ``, 1) + `,`,
		`Message:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Message), "ApplicationUplink", "ApplicationUplink", 1), `&`, ``, 1) + `,`,
		`Parameter:` + fmt.Sprintf("%v", this.Parameter) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProcessDownlinkMessageRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessDownlinkMessageRequest{`,
		`EndDeviceIdentifiers:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.EndDeviceIdentifiers), "EndDeviceIdentifiers", "EndDeviceIdentifiers", 1), `&`, ``, 1) + `,`,
		`EndDeviceVersionIDs:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.EndDeviceVersionIDs), "EndDeviceVersionIdentifiers", "EndDeviceVersionIdentifiers", 1), `&`, ``, 1) + `,`,
		`Message:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Message), "ApplicationDownlink", "ApplicationDownlink", 1), `&`, ``, 1) + `,`,
		`Parameter:` + fmt.Sprintf("%v", this.Parameter) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessageServices(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ProcessUplinkMessageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessageServices
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProcessUplinkMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessUplinkMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDeviceIdentifiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessageServices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EndDeviceIdentifiers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDeviceVersionIDs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessageServices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EndDeviceVersionIDs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessageServices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parameter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageServices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parameter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessageServices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessageServices
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessageServices
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProcessDownlinkMessageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessageServices
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProcessDownlinkMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessDownlinkMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDeviceIdentifiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessageServices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EndDeviceIdentifiers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDeviceVersionIDs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessageServices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EndDeviceVersionIDs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessageServices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parameter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessageServices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parameter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessageServices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessageServices
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessageServices
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessageServices(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessageServices
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMessageServices
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMessageServices
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessageServices
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMessageServices(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMessageServices
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMessageServices = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessageServices   = fmt.Errorf("proto: integer overflow")
)