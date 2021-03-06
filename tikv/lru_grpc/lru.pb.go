// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lru.proto

package lru_gprc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type Key struct {
	Data                 string   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ffa8ae9705147b8, []int{0}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Keys struct {
	Data                 []string `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Keys) Reset()         { *m = Keys{} }
func (m *Keys) String() string { return proto.CompactTextString(m) }
func (*Keys) ProtoMessage()    {}
func (*Keys) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ffa8ae9705147b8, []int{1}
}

func (m *Keys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keys.Unmarshal(m, b)
}
func (m *Keys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keys.Marshal(b, m, deterministic)
}
func (m *Keys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keys.Merge(m, src)
}
func (m *Keys) XXX_Size() int {
	return xxx_messageInfo_Keys.Size(m)
}
func (m *Keys) XXX_DiscardUnknown() {
	xxx_messageInfo_Keys.DiscardUnknown(m)
}

var xxx_messageInfo_Keys proto.InternalMessageInfo

func (m *Keys) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

// The response message containing the greetings
type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ffa8ae9705147b8, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Key)(nil), "lru_gprc.Key")
	proto.RegisterType((*Keys)(nil), "lru_gprc.Keys")
	proto.RegisterType((*Response)(nil), "lru_gprc.Response")
}

func init() { proto.RegisterFile("lru.proto", fileDescriptor_5ffa8ae9705147b8) }

var fileDescriptor_5ffa8ae9705147b8 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x29, 0x2a, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x29, 0x2a, 0x8d, 0x4f, 0x2f, 0x28, 0x4a, 0x56,
	0x92, 0xe4, 0x62, 0xf6, 0x4e, 0xad, 0x14, 0x12, 0xe2, 0x62, 0x71, 0x49, 0x2c, 0x49, 0x94, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xa4, 0xb8, 0x58, 0xbc, 0x53, 0x2b, 0x8b, 0x91,
	0xe4, 0x98, 0xe1, 0x72, 0x2a, 0x5c, 0x1c, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42,
	0x12, 0x5c, 0xec, 0xc5, 0xa5, 0xc9, 0xc9, 0xa9, 0xc5, 0xc5, 0x60, 0xed, 0x1c, 0x41, 0x30, 0xae,
	0x51, 0x2e, 0x17, 0xbb, 0x4f, 0x51, 0xa9, 0x7b, 0x50, 0x41, 0xb2, 0x90, 0x2e, 0x17, 0x5b, 0x50,
	0x6a, 0x6e, 0x7e, 0x59, 0xaa, 0x10, 0xaf, 0x1e, 0xcc, 0x72, 0x3d, 0xef, 0xd4, 0x4a, 0x29, 0x21,
	0x04, 0x17, 0x66, 0xa2, 0x12, 0x83, 0x90, 0x31, 0x17, 0x37, 0x44, 0xb9, 0x53, 0x62, 0x49, 0x72,
	0x86, 0x10, 0x1f, 0x8a, 0x9e, 0x62, 0xec, 0x9a, 0x92, 0xd8, 0xc0, 0x9e, 0x33, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xf0, 0x5f, 0x51, 0xf3, 0xe9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LruGRpcClient is the client API for LruGRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LruGRpcClient interface {
	// Sends a greeting
	Remove(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Response, error)
	RemoveBatch(ctx context.Context, in *Keys, opts ...grpc.CallOption) (*Response, error)
}

type lruGRpcClient struct {
	cc *grpc.ClientConn
}

func NewLruGRpcClient(cc *grpc.ClientConn) LruGRpcClient {
	return &lruGRpcClient{cc}
}

func (c *lruGRpcClient) Remove(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/lru_gprc.LruGRpc/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lruGRpcClient) RemoveBatch(ctx context.Context, in *Keys, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/lru_gprc.LruGRpc/RemoveBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LruGRpcServer is the server API for LruGRpc service.
type LruGRpcServer interface {
	// Sends a greeting
	Remove(context.Context, *Key) (*Response, error)
	RemoveBatch(context.Context, *Keys) (*Response, error)
}

func RegisterLruGRpcServer(s *grpc.Server, srv LruGRpcServer) {
	s.RegisterService(&_LruGRpc_serviceDesc, srv)
}

func _LruGRpc_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LruGRpcServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lru_gprc.LruGRpc/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LruGRpcServer).Remove(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _LruGRpc_RemoveBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Keys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LruGRpcServer).RemoveBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lru_gprc.LruGRpc/RemoveBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LruGRpcServer).RemoveBatch(ctx, req.(*Keys))
	}
	return interceptor(ctx, in, info, handler)
}

var _LruGRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lru_gprc.LruGRpc",
	HandlerType: (*LruGRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Remove",
			Handler:    _LruGRpc_Remove_Handler,
		},
		{
			MethodName: "RemoveBatch",
			Handler:    _LruGRpc_RemoveBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lru.proto",
}
