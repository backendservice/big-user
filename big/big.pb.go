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
	return fileDescriptor_big_1e8c0a456bcf5878, []int{0}
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
	return fileDescriptor_big_1e8c0a456bcf5878, []int{1}
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
	return fileDescriptor_big_1e8c0a456bcf5878, []int{2}
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

func (m *FindResponse) Reset()         { *m = FindResponse{} }
func (m *FindResponse) String() string { return proto.CompactTextString(m) }
func (*FindResponse) ProtoMessage()    {}
func (*FindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_big_1e8c0a456bcf5878, []int{3}
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

func (m *FindResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FindResponse) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *FindResponse) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *FindResponse) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *FindResponse) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *FindResponse) GetReligion() string {
	if m != nil {
		return m.Religion
	}
	return ""
}

func (m *FindResponse) GetNationality() string {
	if m != nil {
		return m.Nationality
	}
	return ""
}

func init() {
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
	RegistUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	FindUser(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
}

type bigClient struct {
	cc *grpc.ClientConn
}

func NewBigClient(cc *grpc.ClientConn) BigClient {
	return &bigClient{cc}
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
	RegistUser(context.Context, *UserRequest) (*UserResponse, error)
	FindUser(context.Context, *FindRequest) (*FindResponse, error)
}

func RegisterBigServer(s *grpc.Server, srv BigServer) {
	s.RegisterService(&_Big_serviceDesc, srv)
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

func init() { proto.RegisterFile("big.proto", fileDescriptor_big_1e8c0a456bcf5878) }

var fileDescriptor_big_1e8c0a456bcf5878 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x97, 0x75, 0xeb, 0xda, 0x77, 0x53, 0x66, 0x0e, 0x52, 0x86, 0x42, 0x29, 0xa2, 0x3b,
	0x15, 0x74, 0x57, 0x4f, 0x3b, 0x78, 0x96, 0x82, 0x1f, 0xa0, 0x7f, 0x5e, 0x62, 0xb0, 0x4d, 0x66,
	0x93, 0x0a, 0x7e, 0x26, 0x3f, 0x86, 0x7e, 0x30, 0x49, 0xba, 0x75, 0xdd, 0x44, 0xef, 0xde, 0xf2,
	0x3e, 0x79, 0x9e, 0xc0, 0xf3, 0x23, 0x2f, 0xf8, 0x19, 0x67, 0xf1, 0xa6, 0x96, 0x5a, 0x52, 0x27,
	0xe3, 0x2c, 0xfa, 0x24, 0x30, 0x7d, 0x52, 0x58, 0x27, 0xf8, 0xda, 0xa0, 0xd2, 0x94, 0xc2, 0x48,
	0xa4, 0x15, 0x06, 0x24, 0x24, 0x4b, 0x3f, 0xb1, 0x67, 0x3a, 0x07, 0x27, 0x65, 0x18, 0x0c, 0x43,
	0xb2, 0x1c, 0x27, 0xe6, 0x48, 0xcf, 0xc1, 0x65, 0x28, 0x0a, 0xac, 0x03, 0xc7, 0xfa, 0xb6, 0x13,
	0x5d, 0x80, 0x57, 0xa6, 0x9a, 0xeb, 0xa6, 0xc0, 0x60, 0x14, 0x92, 0x25, 0x49, 0xba, 0x99, 0x5e,
	0x80, 0x5f, 0x4a, 0xc1, 0xda, 0xcb, 0xb1, 0xbd, 0xdc, 0x0b, 0x26, 0x59, 0x63, 0xc9, 0x19, 0x97,
	0x22, 0x70, 0xed, 0x9b, 0xdd, 0x4c, 0x43, 0x98, 0x8a, 0x54, 0x73, 0x29, 0xd2, 0x92, 0xeb, 0xf7,
	0x60, 0x62, 0xaf, 0xfb, 0x52, 0x74, 0x0f, 0xb3, 0xb6, 0x84, 0xda, 0x48, 0xa1, 0xd0, 0xb4, 0xc8,
	0x65, 0xd1, 0xb5, 0x30, 0x67, 0x1a, 0xc0, 0xa4, 0x42, 0xa5, 0x76, 0x4d, 0xfc, 0x64, 0x37, 0x46,
	0x1f, 0x43, 0x98, 0x3e, 0x70, 0x51, 0xec, 0x18, 0x2c, 0xc0, 0x53, 0x3a, 0xad, 0xb5, 0xb1, 0x12,
	0x5b, 0xba, 0x9b, 0x4d, 0x73, 0x14, 0xc5, 0x1e, 0xc7, 0x76, 0xfa, 0x95, 0xc8, 0x15, 0x9c, 0xd8,
	0xec, 0x11, 0x96, 0x43, 0xd1, 0x34, 0x44, 0x51, 0x74, 0x9e, 0x96, 0x4e, 0x5f, 0xa2, 0xd7, 0x70,
	0xda, 0x46, 0x3a, 0x84, 0xae, 0x35, 0x1d, 0xa9, 0x34, 0x82, 0x99, 0x89, 0x75, 0xae, 0x89, 0x75,
	0x1d, 0x68, 0x07, 0xac, 0xbd, 0xbf, 0x59, 0xfb, 0x3f, 0x59, 0x7f, 0x11, 0x98, 0xb5, 0xb4, 0xf6,
	0xb0, 0xff, 0xdf, 0x97, 0xb9, 0xab, 0xc0, 0x59, 0x73, 0x46, 0x57, 0x00, 0x09, 0x32, 0xae, 0xb4,
	0xf9, 0x3f, 0x74, 0x1e, 0x9b, 0xf5, 0xe8, 0xed, 0xc3, 0xe2, 0xac, 0xa7, 0xb4, 0x7d, 0xa3, 0x01,
	0xbd, 0x05, 0xcf, 0x10, 0xe8, 0x45, 0x7a, 0xdf, 0x67, 0x1b, 0xe9, 0x23, 0x8a, 0x06, 0xeb, 0x1b,
	0xb8, 0xcc, 0x65, 0x15, 0x33, 0xae, 0x9f, 0x9b, 0x2c, 0xce, 0xd2, 0xfc, 0x05, 0x45, 0xa1, 0xb0,
	0x7e, 0xe3, 0x39, 0x1a, 0xff, 0x7a, 0xd4, 0x28, 0xac, 0x1f, 0x49, 0xe6, 0xda, 0xe5, 0x5c, 0x7d,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x93, 0xaf, 0xb7, 0xd8, 0xa9, 0x03, 0x00, 0x00,
}
