// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet/greetpb/greet.proto

package ClientStreamingpb

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
	return fileDescriptor_fe6f881da19a2871, []int{0}
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

type LongGreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LongGreetRequest) Reset()         { *m = LongGreetRequest{} }
func (m *LongGreetRequest) String() string { return proto.CompactTextString(m) }
func (*LongGreetRequest) ProtoMessage()    {}
func (*LongGreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{1}
}

func (m *LongGreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreetRequest.Unmarshal(m, b)
}
func (m *LongGreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreetRequest.Marshal(b, m, deterministic)
}
func (m *LongGreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreetRequest.Merge(m, src)
}
func (m *LongGreetRequest) XXX_Size() int {
	return xxx_messageInfo_LongGreetRequest.Size(m)
}
func (m *LongGreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreetRequest proto.InternalMessageInfo

func (m *LongGreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type LongGreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongGreetResponse) Reset()         { *m = LongGreetResponse{} }
func (m *LongGreetResponse) String() string { return proto.CompactTextString(m) }
func (*LongGreetResponse) ProtoMessage()    {}
func (*LongGreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{2}
}

func (m *LongGreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreetResponse.Unmarshal(m, b)
}
func (m *LongGreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreetResponse.Marshal(b, m, deterministic)
}
func (m *LongGreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreetResponse.Merge(m, src)
}
func (m *LongGreetResponse) XXX_Size() int {
	return xxx_messageInfo_LongGreetResponse.Size(m)
}
func (m *LongGreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreetResponse proto.InternalMessageInfo

func (m *LongGreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "ClientStreaming.Greeting")
	proto.RegisterType((*LongGreetRequest)(nil), "ClientStreaming.LongGreetRequest")
	proto.RegisterType((*LongGreetResponse)(nil), "ClientStreaming.LongGreetResponse")
}

func init() { proto.RegisterFile("greet/greetpb/greet.proto", fileDescriptor_fe6f881da19a2871) }

var fileDescriptor_fe6f881da19a2871 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x07, 0x93, 0x05, 0x49, 0x10, 0x5a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0xdf,
	0x39, 0x27, 0x33, 0x35, 0xaf, 0x24, 0xb8, 0xa4, 0x28, 0x35, 0x31, 0x37, 0x33, 0x2f, 0x5d, 0xc9,
	0x8d, 0x8b, 0xc3, 0x1d, 0x24, 0x9f, 0x99, 0x97, 0x2e, 0x24, 0xcb, 0xc5, 0x95, 0x96, 0x59, 0x54,
	0x5c, 0x12, 0x9f, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x09, 0x16,
	0xf1, 0x4b, 0xcc, 0x4d, 0x15, 0x92, 0xe6, 0xe2, 0xcc, 0x49, 0x84, 0xc9, 0x32, 0x81, 0x65, 0x39,
	0x40, 0x02, 0x20, 0x49, 0x25, 0x4f, 0x2e, 0x01, 0x9f, 0xfc, 0xbc, 0x74, 0xb0, 0x59, 0x41, 0xa9,
	0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xa6, 0x5c, 0x1c, 0xe9, 0x50, 0xb3, 0xc1, 0xa6, 0x71, 0x1b,
	0x49, 0xea, 0xa1, 0xd9, 0xaf, 0x07, 0xb3, 0x3c, 0x08, 0xae, 0x54, 0x49, 0x9b, 0x4b, 0x10, 0xc9,
	0xa8, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c,
	0x12, 0xa8, 0xbb, 0xa0, 0x3c, 0xa3, 0x2c, 0x24, 0x7b, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53,
	0x85, 0xc2, 0xb8, 0x38, 0xe1, 0x62, 0x42, 0x8a, 0x18, 0x56, 0xa2, 0xbb, 0x53, 0x4a, 0x09, 0x9f,
	0x12, 0x88, 0xfd, 0x4a, 0x0c, 0x1a, 0x8c, 0x4e, 0xc2, 0x51, 0x82, 0x68, 0x0a, 0x0b, 0x92, 0x92,
	0xd8, 0xc0, 0x01, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x43, 0x52, 0xd0, 0xb4, 0x75, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LongGreetServiceClient is the client API for LongGreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LongGreetServiceClient interface {
	LongGreet(ctx context.Context, opts ...grpc.CallOption) (LongGreetService_LongGreetClient, error)
}

type longGreetServiceClient struct {
	cc *grpc.ClientConn
}

func NewLongGreetServiceClient(cc *grpc.ClientConn) LongGreetServiceClient {
	return &longGreetServiceClient{cc}
}

func (c *longGreetServiceClient) LongGreet(ctx context.Context, opts ...grpc.CallOption) (LongGreetService_LongGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LongGreetService_serviceDesc.Streams[0], "/ClientStreaming.LongGreetService/LongGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &longGreetServiceLongGreetClient{stream}
	return x, nil
}

type LongGreetService_LongGreetClient interface {
	Send(*LongGreetRequest) error
	CloseAndRecv() (*LongGreetResponse, error)
	grpc.ClientStream
}

type longGreetServiceLongGreetClient struct {
	grpc.ClientStream
}

func (x *longGreetServiceLongGreetClient) Send(m *LongGreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *longGreetServiceLongGreetClient) CloseAndRecv() (*LongGreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LongGreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LongGreetServiceServer is the server API for LongGreetService service.
type LongGreetServiceServer interface {
	LongGreet(LongGreetService_LongGreetServer) error
}

// UnimplementedLongGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLongGreetServiceServer struct {
}

func (*UnimplementedLongGreetServiceServer) LongGreet(srv LongGreetService_LongGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method LongGreet not implemented")
}

func RegisterLongGreetServiceServer(s *grpc.Server, srv LongGreetServiceServer) {
	s.RegisterService(&_LongGreetService_serviceDesc, srv)
}

func _LongGreetService_LongGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LongGreetServiceServer).LongGreet(&longGreetServiceLongGreetServer{stream})
}

type LongGreetService_LongGreetServer interface {
	SendAndClose(*LongGreetResponse) error
	Recv() (*LongGreetRequest, error)
	grpc.ServerStream
}

type longGreetServiceLongGreetServer struct {
	grpc.ServerStream
}

func (x *longGreetServiceLongGreetServer) SendAndClose(m *LongGreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *longGreetServiceLongGreetServer) Recv() (*LongGreetRequest, error) {
	m := new(LongGreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LongGreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ClientStreaming.LongGreetService",
	HandlerType: (*LongGreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LongGreet",
			Handler:       _LongGreetService_LongGreet_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "greet/greetpb/greet.proto",
}
