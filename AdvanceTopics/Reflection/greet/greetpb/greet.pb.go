// Code generated by protoc-gen-go. DO NOT EDIT.
// source: AdvanceTopics/Reflection/greet/greetpb/greet.proto

package greetpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac518c3d4f222987, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac518c3d4f222987, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac518c3d4f222987, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
}

func init() {
	proto.RegisterFile("AdvanceTopics/Reflection/greet/greetpb/greet.proto", fileDescriptor_ac518c3d4f222987)
}

var fileDescriptor_ac518c3d4f222987 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4b, 0xc5, 0x30,
	0x14, 0x85, 0x7d, 0xc2, 0x7b, 0x36, 0x57, 0x45, 0x88, 0x22, 0xa2, 0x08, 0x92, 0x45, 0x41, 0x68,
	0xa1, 0xba, 0x39, 0x29, 0xa2, 0x9b, 0x43, 0x74, 0x72, 0x91, 0x34, 0xde, 0x96, 0x40, 0x9b, 0xc4,
	0x24, 0xed, 0xef, 0x97, 0x26, 0xa9, 0xf8, 0x96, 0x36, 0xe7, 0x9c, 0xe4, 0x3b, 0x97, 0x0b, 0xf5,
	0xe3, 0xf7, 0x24, 0xb4, 0xc4, 0x0f, 0x63, 0x95, 0xf4, 0x15, 0xc7, 0xb6, 0x47, 0x19, 0x94, 0xd1,
	0x55, 0xe7, 0x10, 0x43, 0xfa, 0xda, 0x26, 0xfd, 0x4b, 0xeb, 0x4c, 0x30, 0x74, 0x1d, 0x05, 0x7b,
	0x81, 0xe2, 0x75, 0x3e, 0x28, 0xdd, 0xd1, 0x4b, 0x80, 0x56, 0x39, 0x1f, 0xbe, 0xb4, 0x18, 0xf0,
	0x6c, 0x75, 0xb5, 0xba, 0x21, 0x9c, 0x44, 0xe7, 0x4d, 0x0c, 0x48, 0x2f, 0x80, 0xf4, 0x62, 0x49,
	0x77, 0x63, 0x5a, 0xcc, 0xc6, 0x1c, 0xb2, 0x07, 0x38, 0x88, 0x1c, 0x8e, 0x3f, 0x23, 0xfa, 0x40,
	0x6f, 0xa1, 0xe8, 0x32, 0x37, 0x92, 0xf6, 0xeb, 0xa3, 0x32, 0xd5, 0x2f, 0x75, 0xfc, 0xef, 0x02,
	0xbb, 0x86, 0xc3, 0xfc, 0xd8, 0x5b, 0xa3, 0x3d, 0xd2, 0x53, 0xd8, 0x38, 0xf4, 0x63, 0x1f, 0xf2,
	0x14, 0x59, 0xd5, 0xcf, 0xb9, 0xe5, 0x1d, 0xdd, 0xa4, 0x24, 0xd2, 0x7b, 0x58, 0x47, 0x4d, 0x8f,
	0xff, 0xc3, 0xf3, 0x0c, 0xe7, 0x27, 0xdb, 0x66, 0x62, 0xb3, 0x9d, 0x27, 0xf2, 0xb9, 0x97, 0x37,
	0xd2, 0x6c, 0xe2, 0x32, 0xee, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xcf, 0x21, 0xcd, 0x42,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	// Unary API
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// Unary API
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AdvanceTopics/Reflection/greet/greetpb/greet.proto",
}
