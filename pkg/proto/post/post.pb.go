// Code generated by protoc-gen-go. DO NOT EDIT.
// source: post.proto

package post

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PostRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostRequest) Reset()         { *m = PostRequest{} }
func (m *PostRequest) String() string { return proto.CompactTextString(m) }
func (*PostRequest) ProtoMessage()    {}
func (*PostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{0}
}

func (m *PostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostRequest.Unmarshal(m, b)
}
func (m *PostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostRequest.Marshal(b, m, deterministic)
}
func (m *PostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostRequest.Merge(m, src)
}
func (m *PostRequest) XXX_Size() int {
	return xxx_messageInfo_PostRequest.Size(m)
}
func (m *PostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostRequest proto.InternalMessageInfo

func (m *PostRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PostReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostReply) Reset()         { *m = PostReply{} }
func (m *PostReply) String() string { return proto.CompactTextString(m) }
func (*PostReply) ProtoMessage()    {}
func (*PostReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{1}
}

func (m *PostReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostReply.Unmarshal(m, b)
}
func (m *PostReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostReply.Marshal(b, m, deterministic)
}
func (m *PostReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostReply.Merge(m, src)
}
func (m *PostReply) XXX_Size() int {
	return xxx_messageInfo_PostReply.Size(m)
}
func (m *PostReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PostReply.DiscardUnknown(m)
}

var xxx_messageInfo_PostReply proto.InternalMessageInfo

func (m *PostReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PostRequest)(nil), "post.PostRequest")
	proto.RegisterType((*PostReply)(nil), "post.PostReply")
}

func init() { proto.RegisterFile("post.proto", fileDescriptor_e114ad14deab1dd1) }

var fileDescriptor_e114ad14deab1dd1 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc8, 0x2f, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3,
	0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3,
	0xf3, 0x8a, 0x21, 0x6a, 0x94, 0x14, 0xb9, 0xb8, 0x03, 0xf2, 0x8b, 0x4b, 0x82, 0x52, 0x0b, 0x4b,
	0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xc0, 0x6c, 0x25, 0x55, 0x2e, 0x4e, 0x88, 0x92, 0x82, 0x9c, 0x4a, 0x21, 0x09, 0x2e,
	0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0x74, 0x98, 0x1a, 0x18, 0xd7, 0xc8, 0x93, 0x8b, 0x05, 0xa4,
	0x4c, 0xc8, 0x91, 0x8b, 0xcd, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x55, 0x48, 0x50, 0x0f, 0xec, 0x18,
	0x24, 0xf3, 0xa5, 0xf8, 0x91, 0x85, 0x0a, 0x72, 0x2a, 0x95, 0x84, 0x9b, 0x2e, 0x3f, 0x99, 0xcc,
	0xc4, 0xab, 0xc4, 0xa1, 0x5f, 0x66, 0xa8, 0x0f, 0x92, 0xb3, 0x62, 0xd4, 0x72, 0x92, 0xe5, 0xe2,
	0x4b, 0xce, 0xd7, 0x4b, 0x2f, 0x2e, 0x49, 0x4c, 0x4c, 0x01, 0x6b, 0x70, 0x02, 0xbb, 0x20, 0x00,
	0xe4, 0xe2, 0x00, 0xc6, 0x24, 0x36, 0xb0, 0xd3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x96,
	0x8d, 0x28, 0x11, 0xec, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostClient is the client API for Post service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostClient interface {
	Create(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostReply, error)
}

type postClient struct {
	cc *grpc.ClientConn
}

func NewPostClient(cc *grpc.ClientConn) PostClient {
	return &postClient{cc}
}

func (c *postClient) Create(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostReply, error) {
	out := new(PostReply)
	err := c.cc.Invoke(ctx, "/post.Post/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServer is the server API for Post service.
type PostServer interface {
	Create(context.Context, *PostRequest) (*PostReply, error)
}

func RegisterPostServer(s *grpc.Server, srv PostServer) {
	s.RegisterService(&_Post_serviceDesc, srv)
}

func _Post_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.Post/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).Create(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Post_serviceDesc = grpc.ServiceDesc{
	ServiceName: "post.Post",
	HandlerType: (*PostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Post_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post.proto",
}