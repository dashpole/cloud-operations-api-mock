// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mock_trace_service.proto

package mocktrace

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetNumSpansResponse struct {
	NumSpans             int32    `protobuf:"varint,1,opt,name=num_spans,json=numSpans,proto3" json:"num_spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNumSpansResponse) Reset()         { *m = GetNumSpansResponse{} }
func (m *GetNumSpansResponse) String() string { return proto.CompactTextString(m) }
func (*GetNumSpansResponse) ProtoMessage()    {}
func (*GetNumSpansResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_819c37506e2b5275, []int{0}
}

func (m *GetNumSpansResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNumSpansResponse.Unmarshal(m, b)
}
func (m *GetNumSpansResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNumSpansResponse.Marshal(b, m, deterministic)
}
func (m *GetNumSpansResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNumSpansResponse.Merge(m, src)
}
func (m *GetNumSpansResponse) XXX_Size() int {
	return xxx_messageInfo_GetNumSpansResponse.Size(m)
}
func (m *GetNumSpansResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNumSpansResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNumSpansResponse proto.InternalMessageInfo

func (m *GetNumSpansResponse) GetNumSpans() int32 {
	if m != nil {
		return m.NumSpans
	}
	return 0
}

func init() {
	proto.RegisterType((*GetNumSpansResponse)(nil), "api.GetNumSpansResponse")
}

func init() {
	proto.RegisterFile("mock_trace_service.proto", fileDescriptor_819c37506e2b5275)
}

var fileDescriptor_819c37506e2b5275 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xcf, 0x3f, 0x4b, 0x04, 0x31,
	0x10, 0x05, 0x70, 0x0f, 0x51, 0x34, 0x36, 0x12, 0x41, 0x96, 0xbb, 0x46, 0xae, 0xb2, 0xd9, 0x04,
	0xce, 0xd2, 0x46, 0x0f, 0xc4, 0x4a, 0x8b, 0x3b, 0x6d, 0x6c, 0x96, 0x6c, 0x1c, 0xd7, 0xe1, 0x36,
	0x99, 0x21, 0x7f, 0x04, 0xbf, 0xbd, 0x64, 0xa3, 0x60, 0x71, 0x5d, 0x78, 0xbc, 0xcc, 0x8f, 0x27,
	0x1a, 0x47, 0x76, 0xd7, 0xa5, 0x60, 0x2c, 0x74, 0x11, 0xc2, 0x17, 0x5a, 0x50, 0x1c, 0x28, 0x91,
	0x3c, 0x34, 0x8c, 0xf3, 0xc5, 0x40, 0x34, 0x8c, 0xa0, 0xa7, 0xa8, 0xcf, 0x1f, 0x1a, 0x1c, 0xa7,
	0xef, 0xda, 0x58, 0xae, 0xc4, 0xc5, 0x23, 0xa4, 0xe7, 0xec, 0xb6, 0x6c, 0x7c, 0xdc, 0x40, 0x64,
	0xf2, 0x11, 0xe4, 0x42, 0x9c, 0xfa, 0xec, 0xba, 0x58, 0xc2, 0x66, 0x76, 0x35, 0xbb, 0x3e, 0xda,
	0x9c, 0xf8, 0xdf, 0xd2, 0xea, 0x55, 0x9c, 0x3f, 0x91, 0xdd, 0xbd, 0x14, 0x70, 0x5b, 0x3d, 0x79,
	0x2f, 0xce, 0xfe, 0xdd, 0x91, 0x97, 0xaa, 0xa2, 0xea, 0x0f, 0x55, 0x0f, 0x05, 0x9d, 0x37, 0xca,
	0x30, 0xaa, 0x3d, 0xe2, 0xf2, 0x60, 0xbd, 0x7e, 0xbb, 0x1b, 0x30, 0x7d, 0xe6, 0x5e, 0x59, 0x72,
	0xba, 0xfe, 0x47, 0x9f, 0x20, 0xf8, 0xa8, 0xed, 0x48, 0xf9, 0xbd, 0x25, 0x86, 0x60, 0x12, 0x92,
	0x8f, 0xad, 0x61, 0x6c, 0xcb, 0x6c, 0x6d, 0x18, 0x6f, 0xcb, 0x63, 0x9a, 0xdf, 0x1f, 0x4f, 0xde,
	0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0x28, 0x29, 0x9f, 0x13, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MockTraceServiceClient is the client API for MockTraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MockTraceServiceClient interface {
	GetNumSpans(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetNumSpansResponse, error)
}

type mockTraceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMockTraceServiceClient(cc grpc.ClientConnInterface) MockTraceServiceClient {
	return &mockTraceServiceClient{cc}
}

func (c *mockTraceServiceClient) GetNumSpans(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetNumSpansResponse, error) {
	out := new(GetNumSpansResponse)
	err := c.cc.Invoke(ctx, "/api.MockTraceService/GetNumSpans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MockTraceServiceServer is the server API for MockTraceService service.
type MockTraceServiceServer interface {
	GetNumSpans(context.Context, *empty.Empty) (*GetNumSpansResponse, error)
}

// UnimplementedMockTraceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMockTraceServiceServer struct {
}

func (*UnimplementedMockTraceServiceServer) GetNumSpans(ctx context.Context, req *empty.Empty) (*GetNumSpansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumSpans not implemented")
}

func RegisterMockTraceServiceServer(s *grpc.Server, srv MockTraceServiceServer) {
	s.RegisterService(&_MockTraceService_serviceDesc, srv)
}

func _MockTraceService_GetNumSpans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockTraceServiceServer).GetNumSpans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MockTraceService/GetNumSpans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockTraceServiceServer).GetNumSpans(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _MockTraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.MockTraceService",
	HandlerType: (*MockTraceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNumSpans",
			Handler:    _MockTraceService_GetNumSpans_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mock_trace_service.proto",
}
