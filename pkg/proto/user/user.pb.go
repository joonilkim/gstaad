// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type MutationReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutationReply) Reset()         { *m = MutationReply{} }
func (m *MutationReply) String() string { return proto.CompactTextString(m) }
func (*MutationReply) ProtoMessage()    {}
func (*MutationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *MutationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutationReply.Unmarshal(m, b)
}
func (m *MutationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutationReply.Marshal(b, m, deterministic)
}
func (m *MutationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutationReply.Merge(m, src)
}
func (m *MutationReply) XXX_Size() int {
	return xxx_messageInfo_MutationReply.Size(m)
}
func (m *MutationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MutationReply.DiscardUnknown(m)
}

var xxx_messageInfo_MutationReply proto.InternalMessageInfo

func (m *MutationReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            int64    `protobuf:"varint,15,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type Token struct {
	AccessToken          string               `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	ExpiresIn            *timestamp.Timestamp `protobuf:"bytes,15,opt,name=expiresIn,proto3" json:"expiresIn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Token) GetExpiresIn() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiresIn
	}
	return nil
}

type LoginRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LoginReply struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *LoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReply.Unmarshal(m, b)
}
func (m *LoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReply.Marshal(b, m, deterministic)
}
func (m *LoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReply.Merge(m, src)
}
func (m *LoginReply) XXX_Size() int {
	return xxx_messageInfo_LoginReply.Size(m)
}
func (m *LoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReply proto.InternalMessageInfo

func (m *LoginReply) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type ProfileReply struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	PostCount            int32    `protobuf:"varint,2,opt,name=postCount,proto3" json:"postCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileReply) Reset()         { *m = ProfileReply{} }
func (m *ProfileReply) String() string { return proto.CompactTextString(m) }
func (*ProfileReply) ProtoMessage()    {}
func (*ProfileReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *ProfileReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileReply.Unmarshal(m, b)
}
func (m *ProfileReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileReply.Marshal(b, m, deterministic)
}
func (m *ProfileReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileReply.Merge(m, src)
}
func (m *ProfileReply) XXX_Size() int {
	return xxx_messageInfo_ProfileReply.Size(m)
}
func (m *ProfileReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileReply.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileReply proto.InternalMessageInfo

func (m *ProfileReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ProfileReply) GetPostCount() int32 {
	if m != nil {
		return m.PostCount
	}
	return 0
}

func init() {
	proto.RegisterType((*MutationReply)(nil), "user.MutationReply")
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Token)(nil), "user.Token")
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "user.LoginReply")
	proto.RegisterType((*ProfileReply)(nil), "user.ProfileReply")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xb5, 0x69, 0xb6, 0x25, 0xb3, 0xa1, 0x20, 0x1f, 0xaa, 0x68, 0x29, 0x10, 0x7c, 0xa1,
	0xe2, 0xb0, 0x2b, 0x85, 0x0b, 0xe2, 0xd6, 0x20, 0x24, 0x90, 0x8a, 0x14, 0x99, 0xf2, 0x00, 0xee,
	0x66, 0x1a, 0x59, 0xec, 0xae, 0x8d, 0xed, 0x45, 0xe4, 0xca, 0x2b, 0xc0, 0x9b, 0xf1, 0x0a, 0x3c,
	0x08, 0xf2, 0xd8, 0x4d, 0x02, 0xdc, 0xec, 0x7f, 0xfe, 0xf9, 0xf5, 0xcd, 0xd8, 0x00, 0x83, 0x43,
	0x5b, 0x19, 0xab, 0xbd, 0x66, 0xe3, 0x70, 0x2e, 0x1f, 0x6d, 0xb4, 0xde, 0xb4, 0x58, 0x93, 0x76,
	0x33, 0xdc, 0xd6, 0xd8, 0x19, 0xbf, 0x8d, 0x96, 0xf2, 0xe9, 0xbf, 0x45, 0xaf, 0x3a, 0x74, 0x5e,
	0x76, 0x26, 0x19, 0xce, 0x93, 0x41, 0x1a, 0x55, 0xcb, 0xbe, 0xd7, 0x5e, 0x7a, 0xa5, 0x7b, 0x17,
	0xab, 0xfc, 0x39, 0xdc, 0xff, 0x30, 0x44, 0x49, 0xa0, 0x69, 0xb7, 0xec, 0x0c, 0x8e, 0x2d, 0xba,
	0xa1, 0xf5, 0xb3, 0x6c, 0x9e, 0x5d, 0xdc, 0x13, 0xe9, 0xc6, 0xdf, 0xc1, 0xf8, 0x93, 0x43, 0xcb,
	0x4e, 0x61, 0xa4, 0xd6, 0x54, 0x9b, 0x88, 0x91, 0x5a, 0x33, 0x06, 0xe3, 0x5e, 0x76, 0x38, 0x1b,
	0x91, 0x42, 0x67, 0x76, 0x0e, 0x93, 0xc6, 0xa2, 0xf4, 0xb8, 0xbe, 0xf4, 0xb3, 0x07, 0xf3, 0xec,
	0xe2, 0x48, 0xec, 0x05, 0xde, 0x40, 0x7e, 0xad, 0x3f, 0x63, 0xcf, 0xe6, 0x50, 0xc8, 0xa6, 0x41,
	0xe7, 0xe8, 0x9a, 0x32, 0x0f, 0x25, 0xf6, 0x0a, 0x26, 0xf8, 0xcd, 0x28, 0x8b, 0xee, 0x7d, 0x4f,
	0x41, 0xc5, 0xa2, 0xac, 0xe2, 0x3c, 0xd5, 0xdd, 0xc0, 0xd5, 0xf5, 0xdd, 0xc0, 0x62, 0x6f, 0xe6,
	0x1c, 0xa6, 0x57, 0x7a, 0xa3, 0x7a, 0x81, 0x5f, 0x06, 0x74, 0x7e, 0x87, 0x99, 0xed, 0x31, 0x79,
	0x0d, 0x90, 0x3c, 0x61, 0xf0, 0x67, 0x90, 0xfb, 0x1d, 0x47, 0xb1, 0x28, 0x2a, 0x7a, 0x07, 0xe2,
	0x10, 0xb1, 0xc2, 0xaf, 0x60, 0xba, 0xb2, 0xfa, 0x56, 0xb5, 0x18, 0x5b, 0x9e, 0x00, 0x3d, 0x50,
	0xea, 0x80, 0xd8, 0x11, 0xb6, 0x24, 0x48, 0x0f, 0x7b, 0x30, 0xda, 0xf9, 0x37, 0x7a, 0xe8, 0x3d,
	0x2d, 0x28, 0x17, 0x7b, 0x61, 0xf1, 0x33, 0x83, 0x22, 0x98, 0x3f, 0xa2, 0xfd, 0xaa, 0x1a, 0x64,
	0x97, 0x90, 0x13, 0x0e, 0x63, 0x31, 0xe8, 0x90, 0xbf, 0x7c, 0xf8, 0x97, 0x66, 0xda, 0x2d, 0x67,
	0xdf, 0x7f, 0xfd, 0xfe, 0x31, 0x9a, 0xf2, 0x93, 0x9a, 0xe0, 0xdc, 0xeb, 0xec, 0x05, 0x5b, 0xc2,
	0x49, 0x02, 0x64, 0x67, 0xff, 0xed, 0xe9, 0x6d, 0xf8, 0x35, 0x65, 0x0a, 0x3f, 0x9c, 0x83, 0x17,
	0x14, 0x95, 0xb3, 0xa3, 0xba, 0xc3, 0xe5, 0x63, 0x38, 0x6d, 0x74, 0xb5, 0x71, 0x5e, 0xca, 0x35,
	0x79, 0x97, 0x93, 0x40, 0xb9, 0x0a, 0x31, 0xab, 0xec, 0xe6, 0x98, 0xf2, 0x5e, 0xfe, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x3e, 0x0b, 0x89, 0xd9, 0xa7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	Profile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProfileReply, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/user.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Profile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/user.UserService/Profile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Profile(context.Context, *empty.Empty) (*ProfileReply, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Profile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Profile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Profile(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Profile",
			Handler:    _UserService_Profile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
