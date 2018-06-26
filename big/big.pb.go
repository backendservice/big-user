// Code generated by protoc-gen-go. DO NOT EDIT.
// source: big.proto

package big

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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_big_ccc6a785c59c7e91, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_big_ccc6a785c59c7e91, []int{1}
}
func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (dst *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(dst, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Gender               string   `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Latitude             float64  `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Religion             string   `protobuf:"bytes,6,opt,name=religion,proto3" json:"religion,omitempty"`
	Nationality          string   `protobuf:"bytes,7,opt,name=nationality,proto3" json:"nationality,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_big_ccc6a785c59c7e91, []int{2}
}
func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (dst *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(dst, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserRequest) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UserRequest) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *UserRequest) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *UserRequest) GetReligion() string {
	if m != nil {
		return m.Religion
	}
	return ""
}

func (m *UserRequest) GetNationality() string {
	if m != nil {
		return m.Nationality
	}
	return ""
}

type UserResponse struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_big_ccc6a785c59c7e91, []int{3}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *UserResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type FindRequest struct {
	Startage             int32    `protobuf:"varint,1,opt,name=startage,proto3" json:"startage,omitempty"`
	Endage               int32    `protobuf:"varint,2,opt,name=endage,proto3" json:"endage,omitempty"`
	Gender               string   `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Startlatitude        float64  `protobuf:"fixed64,4,opt,name=startlatitude,proto3" json:"startlatitude,omitempty"`
	Endlatitude          float64  `protobuf:"fixed64,5,opt,name=endlatitude,proto3" json:"endlatitude,omitempty"`
	Startlongitude       float64  `protobuf:"fixed64,6,opt,name=startlongitude,proto3" json:"startlongitude,omitempty"`
	Endlongitude         float64  `protobuf:"fixed64,7,opt,name=endlongitude,proto3" json:"endlongitude,omitempty"`
	Religion             string   `protobuf:"bytes,8,opt,name=religion,proto3" json:"religion,omitempty"`
	Nationality          string   `protobuf:"bytes,9,opt,name=nationality,proto3" json:"nationality,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_big_ccc6a785c59c7e91, []int{4}
}
func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (dst *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(dst, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetStartage() int32 {
	if m != nil {
		return m.Startage
	}
	return 0
}

func (m *FindRequest) GetEndage() int32 {
	if m != nil {
		return m.Endage
	}
	return 0
}

func (m *FindRequest) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *FindRequest) GetStartlatitude() float64 {
	if m != nil {
		return m.Startlatitude
	}
	return 0
}

func (m *FindRequest) GetEndlatitude() float64 {
	if m != nil {
		return m.Endlatitude
	}
	return 0
}

func (m *FindRequest) GetStartlongitude() float64 {
	if m != nil {
		return m.Startlongitude
	}
	return 0
}

func (m *FindRequest) GetEndlongitude() float64 {
	if m != nil {
		return m.Endlongitude
	}
	return 0
}

func (m *FindRequest) GetReligion() string {
	if m != nil {
		return m.Religion
	}
	return ""
}

func (m *FindRequest) GetNationality() string {
	if m != nil {
		return m.Nationality
	}
	return ""
}

type FindResponse struct {
	Result               []*UserRequest `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	Count                int32          `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Code                 string         `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Message              string         `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FindResponse) Reset()         { *m = FindResponse{} }
func (m *FindResponse) String() string { return proto.CompactTextString(m) }
func (*FindResponse) ProtoMessage()    {}
func (*FindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_big_ccc6a785c59c7e91, []int{5}
}
func (m *FindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResponse.Unmarshal(m, b)
}
func (m *FindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResponse.Marshal(b, m, deterministic)
}
func (dst *FindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResponse.Merge(dst, src)
}
func (m *FindResponse) XXX_Size() int {
	return xxx_messageInfo_FindResponse.Size(m)
}
func (m *FindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindResponse proto.InternalMessageInfo

func (m *FindResponse) GetResult() []*UserRequest {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *FindResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *FindResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *FindResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "big.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "big.HelloResponse")
	proto.RegisterType((*UserRequest)(nil), "big.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "big.UserResponse")
	proto.RegisterType((*FindRequest)(nil), "big.FindRequest")
	proto.RegisterType((*FindResponse)(nil), "big.FindResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BigClient is the client API for Big service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BigClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	RegistUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	FindUser(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
}

type bigClient struct {
	cc *grpc.ClientConn
}

func NewBigClient(cc *grpc.ClientConn) BigClient {
	return &bigClient{cc}
}

func (c *bigClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/big.Big/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigClient) RegistUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/big.Big/RegistUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigClient) FindUser(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/big.Big/FindUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BigServer is the server API for Big service.
type BigServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	RegistUser(context.Context, *UserRequest) (*UserResponse, error)
	FindUser(context.Context, *FindRequest) (*FindResponse, error)
}

func RegisterBigServer(s *grpc.Server, srv BigServer) {
	s.RegisterService(&_Big_serviceDesc, srv)
}

func _Big_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/big.Big/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Big_RegistUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigServer).RegistUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/big.Big/RegistUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigServer).RegistUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Big_FindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigServer).FindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/big.Big/FindUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigServer).FindUser(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Big_serviceDesc = grpc.ServiceDesc{
	ServiceName: "big.Big",
	HandlerType: (*BigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Big_Hello_Handler,
		},
		{
			MethodName: "RegistUser",
			Handler:    _Big_RegistUser_Handler,
		},
		{
			MethodName: "FindUser",
			Handler:    _Big_FindUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "big.proto",
}

func init() { proto.RegisterFile("big.proto", fileDescriptor_big_ccc6a785c59c7e91) }

var fileDescriptor_big_ccc6a785c59c7e91 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x37, 0xcd, 0x9f, 0xb6, 0xd3, 0x2e, 0xda, 0x1d, 0x21, 0x14, 0x55, 0x20, 0x45, 0x11,
	0x82, 0x70, 0xa9, 0x60, 0xf7, 0xca, 0xa9, 0x07, 0xc4, 0x11, 0x45, 0xe2, 0x01, 0xf2, 0x67, 0x64,
	0x2c, 0x52, 0x7b, 0x89, 0x1d, 0x24, 0x24, 0xde, 0x85, 0x07, 0xe0, 0x31, 0x78, 0x31, 0x64, 0x27,
	0xf5, 0x3a, 0xbb, 0xd0, 0x9b, 0xe7, 0xf3, 0xf7, 0x25, 0x9e, 0x9f, 0xc7, 0xb0, 0xae, 0x39, 0xdb,
	0xdf, 0xf5, 0x52, 0x4b, 0x0c, 0x6b, 0xce, 0xf2, 0x1c, 0xb6, 0x1f, 0xa9, 0xeb, 0x64, 0x49, 0xdf,
	0x06, 0x52, 0x1a, 0x11, 0x22, 0x51, 0x1d, 0x29, 0x0d, 0xb2, 0xa0, 0x58, 0x97, 0x76, 0x9d, 0xbf,
	0x81, 0xcb, 0xc9, 0xa3, 0xee, 0xa4, 0x50, 0x84, 0x29, 0x2c, 0x8f, 0xa4, 0x54, 0xc5, 0x4e, 0xbe,
	0x53, 0x99, 0xff, 0x09, 0x60, 0xf3, 0x59, 0x51, 0x7f, 0xe6, 0x73, 0x78, 0x05, 0xa1, 0x49, 0x2e,
	0xb2, 0xa0, 0x88, 0x4b, 0xb3, 0xc4, 0x67, 0x90, 0x30, 0x12, 0x2d, 0xf5, 0x69, 0x68, 0x7d, 0x53,
	0x85, 0x3b, 0x58, 0x75, 0x95, 0xe6, 0x7a, 0x68, 0x29, 0x8d, 0xb2, 0xa0, 0x08, 0x4a, 0x57, 0xe3,
	0x73, 0x58, 0x77, 0x52, 0xb0, 0x71, 0x33, 0xb6, 0x9b, 0xf7, 0x82, 0x49, 0xf6, 0xd4, 0x71, 0xc6,
	0xa5, 0x48, 0x13, 0xfb, 0x4d, 0x57, 0x63, 0x06, 0x1b, 0x51, 0x69, 0x2e, 0x45, 0xd5, 0x71, 0xfd,
	0x23, 0x5d, 0xda, 0x6d, 0x5f, 0xca, 0xdf, 0xc3, 0x76, 0x6c, 0x62, 0xea, 0x17, 0x21, 0x6a, 0x64,
	0xeb, 0xba, 0x30, 0x6b, 0x9f, 0xc1, 0x62, 0xce, 0xe0, 0xf7, 0x02, 0x36, 0x1f, 0xb8, 0x68, 0x4f,
	0x0c, 0x76, 0xb0, 0x52, 0xba, 0xea, 0xf5, 0x09, 0x57, 0x5c, 0xba, 0xda, 0x74, 0x4e, 0xa2, 0xbd,
	0xc7, 0x31, 0x55, 0xff, 0x25, 0xf2, 0x12, 0x2e, 0x6d, 0xf6, 0x01, 0x96, 0xb9, 0x68, 0x3a, 0x24,
	0xd1, 0x3a, 0xcf, 0x48, 0xc7, 0x97, 0xf0, 0x15, 0x3c, 0x19, 0x23, 0x0e, 0x61, 0x62, 0x4d, 0x0f,
	0x54, 0xcc, 0x61, 0x6b, 0x62, 0xce, 0xb5, 0xb4, 0xae, 0x99, 0x36, 0x63, 0xbd, 0x3a, 0xcf, 0x7a,
	0xfd, 0x98, 0xf5, 0x4f, 0xd8, 0x8e, 0xb0, 0x26, 0xd6, 0x05, 0x24, 0x3d, 0xa9, 0xa1, 0xd3, 0x69,
	0x90, 0x85, 0xc5, 0xe6, 0xe6, 0x6a, 0x6f, 0x26, 0xd6, 0x9b, 0xa9, 0x72, 0xda, 0xc7, 0xa7, 0x10,
	0x37, 0x72, 0x10, 0x7a, 0x42, 0x37, 0x16, 0xee, 0xae, 0xc2, 0x7f, 0xdf, 0x55, 0x34, 0xbb, 0xab,
	0x9b, 0x5f, 0x01, 0x84, 0x07, 0xce, 0xf0, 0x2d, 0xc4, 0x76, 0xc4, 0xf1, 0xda, 0xfe, 0xce, 0x7f,
	0x12, 0x3b, 0xf4, 0xa5, 0xf1, 0x94, 0xf9, 0x05, 0xde, 0x02, 0x94, 0xc4, 0xb8, 0xd2, 0xe6, 0x68,
	0xf8, 0xe8, 0x94, 0xbb, 0x6b, 0x4f, 0x71, 0xa1, 0x77, 0xb0, 0x32, 0xcd, 0x7a, 0x11, 0x6f, 0x50,
	0xa6, 0x88, 0x4f, 0x23, 0xbf, 0x38, 0xbc, 0x86, 0x17, 0x8d, 0x3c, 0xee, 0x19, 0xd7, 0x5f, 0x86,
	0x7a, 0x5f, 0x57, 0xcd, 0x57, 0x12, 0xad, 0xa2, 0xfe, 0x3b, 0x6f, 0xc8, 0xf8, 0x0f, 0xd1, 0xa0,
	0xa8, 0xff, 0x14, 0xd4, 0x89, 0x7d, 0xd5, 0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x8e,
	0x51, 0x66, 0xe2, 0x03, 0x00, 0x00,
}
