// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ProjectBloggingServiceUsingMongoAndGrpc/blogpb/blog.proto

package blogpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AuthorId             string   `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Contact              string   `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_841e1cc5d38fc1b8, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetContact() string {
	if m != nil {
		return m.Contact
	}
	return ""
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
}

func init() {
	proto.RegisterFile("ProjectBloggingServiceUsingMongoAndGrpc/blogpb/blog.proto", fileDescriptor_841e1cc5d38fc1b8)
}

var fileDescriptor_841e1cc5d38fc1b8 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x0c, 0x28, 0xca, 0xcf,
	0x4a, 0x4d, 0x2e, 0x71, 0xca, 0xc9, 0x4f, 0x4f, 0xcf, 0xcc, 0x4b, 0x0f, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0x0d, 0x2d, 0xce, 0xcc, 0x4b, 0xf7, 0xcd, 0xcf, 0x4b, 0xcf, 0x77, 0xcc, 0x4b, 0x71,
	0x2f, 0x2a, 0x48, 0xd6, 0x4f, 0xca, 0xc9, 0x4f, 0x2f, 0x48, 0x02, 0x53, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x52, 0x32, 0x17, 0x0b, 0x48, 0xaf, 0x10, 0x1f, 0x17, 0x93,
	0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x93, 0x67, 0x8a, 0x90, 0x34, 0x17, 0x67,
	0x62, 0x69, 0x49, 0x46, 0x7e, 0x51, 0x7c, 0x66, 0x8a, 0x04, 0x13, 0x58, 0x98, 0x03, 0x22, 0xe0,
	0x99, 0x22, 0x24, 0xc2, 0xc5, 0x5a, 0x92, 0x59, 0x92, 0x93, 0x2a, 0xc1, 0x0c, 0x96, 0x80, 0x70,
	0x84, 0x24, 0xb8, 0xd8, 0x93, 0xf3, 0xf3, 0x4a, 0x12, 0x93, 0x4b, 0x24, 0x58, 0xc0, 0xe2, 0x30,
	0xae, 0x11, 0x2f, 0x17, 0x37, 0xc8, 0x12, 0xa8, 0xe3, 0x9c, 0x38, 0xa2, 0xd8, 0x20, 0xce, 0x49,
	0x62, 0x03, 0x3b, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xad, 0x48, 0xc9, 0xc7, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
}

// UnimplementedBlogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "ProjectBloggingServiceUsingMongoAndGrpc/blogpb/blog.proto",
}
