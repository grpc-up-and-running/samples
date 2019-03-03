// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product_mgt.proto

package ecommerce

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

type ProductRequest struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductRequest) Reset()         { *m = ProductRequest{} }
func (m *ProductRequest) String() string { return proto.CompactTextString(m) }
func (*ProductRequest) ProtoMessage()    {}
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc353d63eb36ab2d, []int{0}
}

func (m *ProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductRequest.Unmarshal(m, b)
}
func (m *ProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductRequest.Marshal(b, m, deterministic)
}
func (m *ProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductRequest.Merge(m, src)
}
func (m *ProductRequest) XXX_Size() int {
	return xxx_messageInfo_ProductRequest.Size(m)
}
func (m *ProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProductRequest proto.InternalMessageInfo

func (m *ProductRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ProductResponse struct {
	ProductID            string   `protobuf:"bytes,1,opt,name=productID,proto3" json:"productID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductResponse) Reset()         { *m = ProductResponse{} }
func (m *ProductResponse) String() string { return proto.CompactTextString(m) }
func (*ProductResponse) ProtoMessage()    {}
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc353d63eb36ab2d, []int{1}
}

func (m *ProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductResponse.Unmarshal(m, b)
}
func (m *ProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductResponse.Marshal(b, m, deterministic)
}
func (m *ProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductResponse.Merge(m, src)
}
func (m *ProductResponse) XXX_Size() int {
	return xxx_messageInfo_ProductResponse.Size(m)
}
func (m *ProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductResponse proto.InternalMessageInfo

func (m *ProductResponse) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductRequest)(nil), "ecommerce.ProductRequest")
	proto.RegisterType((*ProductResponse)(nil), "ecommerce.ProductResponse")
}

func init() { proto.RegisterFile("product_mgt.proto", fileDescriptor_bc353d63eb36ab2d) }

var fileDescriptor_bc353d63eb36ab2d = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0x89, 0xcf, 0x4d, 0x2f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c,
	0x4d, 0xce, 0xcf, 0xcd, 0x4d, 0x2d, 0x4a, 0x4e, 0x55, 0x72, 0xe3, 0xe2, 0x0b, 0x80, 0xc8, 0x07,
	0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30,
	0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x0a, 0x5c, 0xdc, 0x29, 0xa9, 0xc5, 0xc9, 0x45,
	0x99, 0x05, 0x25, 0x99, 0xf9, 0x79, 0x12, 0xcc, 0x60, 0x29, 0x64, 0x21, 0x25, 0x7d, 0x2e, 0x7e,
	0xb8, 0x39, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x32, 0x5c, 0x9c, 0x50, 0xab, 0x3d, 0x5d,
	0x24, 0x18, 0xc1, 0x5a, 0x10, 0x02, 0x46, 0x81, 0x5c, 0x5c, 0x50, 0x0d, 0xbe, 0xe9, 0x25, 0x42,
	0xce, 0x5c, 0x5c, 0x89, 0x29, 0x29, 0x50, 0x01, 0x21, 0x49, 0x3d, 0xb8, 0x03, 0xf5, 0x50, 0x5d,
	0x27, 0x25, 0x85, 0x4d, 0x0a, 0x62, 0x61, 0x12, 0x1b, 0xd8, 0x77, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x40, 0x53, 0x73, 0x89, 0xf2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductMgtClient is the client API for ProductMgt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductMgtClient interface {
	AddProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
}

type productMgtClient struct {
	cc *grpc.ClientConn
}

func NewProductMgtClient(cc *grpc.ClientConn) ProductMgtClient {
	return &productMgtClient{cc}
}

func (c *productMgtClient) AddProduct(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/ecommerce.ProductMgt/addProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductMgtServer is the server API for ProductMgt service.
type ProductMgtServer interface {
	AddProduct(context.Context, *ProductRequest) (*ProductResponse, error)
}

func RegisterProductMgtServer(s *grpc.Server, srv ProductMgtServer) {
	s.RegisterService(&_ProductMgt_serviceDesc, srv)
}

func _ProductMgt_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductMgtServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.ProductMgt/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductMgtServer).AddProduct(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductMgt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.ProductMgt",
	HandlerType: (*ProductMgtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addProduct",
			Handler:    _ProductMgt_AddProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_mgt.proto",
}