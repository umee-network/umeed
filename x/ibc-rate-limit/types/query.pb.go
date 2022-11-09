// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/ibc_rate_limit/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParams defines the request structure for the Params gRPC service
// handler.
type QueryParams struct {
}

func (m *QueryParams) Reset()         { *m = QueryParams{} }
func (m *QueryParams) String() string { return proto.CompactTextString(m) }
func (*QueryParams) ProtoMessage()    {}
func (*QueryParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1706951c856d331f, []int{0}
}
func (m *QueryParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParams.Merge(m, src)
}
func (m *QueryParams) XXX_Size() int {
	return m.Size()
}
func (m *QueryParams) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParams.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParams proto.InternalMessageInfo

// QueryParamsResponse defines the response structure for the Params gRPC
// service handler.
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1706951c856d331f, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

// QueryRateLimitsOfIBCDenoms defines request type
type QueryRateLimitsOfIBCDenoms struct {
}

func (m *QueryRateLimitsOfIBCDenoms) Reset()         { *m = QueryRateLimitsOfIBCDenoms{} }
func (m *QueryRateLimitsOfIBCDenoms) String() string { return proto.CompactTextString(m) }
func (*QueryRateLimitsOfIBCDenoms) ProtoMessage()    {}
func (*QueryRateLimitsOfIBCDenoms) Descriptor() ([]byte, []int) {
	return fileDescriptor_1706951c856d331f, []int{2}
}
func (m *QueryRateLimitsOfIBCDenoms) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRateLimitsOfIBCDenoms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRateLimitsOfIBCDenoms.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRateLimitsOfIBCDenoms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRateLimitsOfIBCDenoms.Merge(m, src)
}
func (m *QueryRateLimitsOfIBCDenoms) XXX_Size() int {
	return m.Size()
}
func (m *QueryRateLimitsOfIBCDenoms) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRateLimitsOfIBCDenoms.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRateLimitsOfIBCDenoms proto.InternalMessageInfo

// QueryRateLimitsOfIBCDenomsResponse defines response type
type QueryRateLimitsOfIBCDenomsResponse struct {
	RateLimits []RateLimit `protobuf:"bytes,1,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits"`
}

func (m *QueryRateLimitsOfIBCDenomsResponse) Reset()         { *m = QueryRateLimitsOfIBCDenomsResponse{} }
func (m *QueryRateLimitsOfIBCDenomsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRateLimitsOfIBCDenomsResponse) ProtoMessage()    {}
func (*QueryRateLimitsOfIBCDenomsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1706951c856d331f, []int{3}
}
func (m *QueryRateLimitsOfIBCDenomsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRateLimitsOfIBCDenomsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRateLimitsOfIBCDenomsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRateLimitsOfIBCDenomsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRateLimitsOfIBCDenomsResponse.Merge(m, src)
}
func (m *QueryRateLimitsOfIBCDenomsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRateLimitsOfIBCDenomsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRateLimitsOfIBCDenomsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRateLimitsOfIBCDenomsResponse proto.InternalMessageInfo

// QueryRateLimitsOfIBCDenom defines request type
type QueryRateLimitsOfIBCDenom struct {
	IbcDenom string `protobuf:"bytes,1,opt,name=ibc_denom,json=ibcDenom,proto3" json:"ibc_denom,omitempty"`
}

func (m *QueryRateLimitsOfIBCDenom) Reset()         { *m = QueryRateLimitsOfIBCDenom{} }
func (m *QueryRateLimitsOfIBCDenom) String() string { return proto.CompactTextString(m) }
func (*QueryRateLimitsOfIBCDenom) ProtoMessage()    {}
func (*QueryRateLimitsOfIBCDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_1706951c856d331f, []int{4}
}
func (m *QueryRateLimitsOfIBCDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRateLimitsOfIBCDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRateLimitsOfIBCDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRateLimitsOfIBCDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRateLimitsOfIBCDenom.Merge(m, src)
}
func (m *QueryRateLimitsOfIBCDenom) XXX_Size() int {
	return m.Size()
}
func (m *QueryRateLimitsOfIBCDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRateLimitsOfIBCDenom.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRateLimitsOfIBCDenom proto.InternalMessageInfo

// QueryRateLimitsOfIBCDenomResponse defines response type of QueryRateLimitsOfIBCDenom
type QueryRateLimitsOfIBCDenomResponse struct {
	RateLimit RateLimit `protobuf:"bytes,1,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit"`
}

func (m *QueryRateLimitsOfIBCDenomResponse) Reset()         { *m = QueryRateLimitsOfIBCDenomResponse{} }
func (m *QueryRateLimitsOfIBCDenomResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRateLimitsOfIBCDenomResponse) ProtoMessage()    {}
func (*QueryRateLimitsOfIBCDenomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1706951c856d331f, []int{5}
}
func (m *QueryRateLimitsOfIBCDenomResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRateLimitsOfIBCDenomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRateLimitsOfIBCDenomResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRateLimitsOfIBCDenomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRateLimitsOfIBCDenomResponse.Merge(m, src)
}
func (m *QueryRateLimitsOfIBCDenomResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRateLimitsOfIBCDenomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRateLimitsOfIBCDenomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRateLimitsOfIBCDenomResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryParams)(nil), "umee.ibcratelimit.v1beta1.QueryParams")
	proto.RegisterType((*QueryParamsResponse)(nil), "umee.ibcratelimit.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryRateLimitsOfIBCDenoms)(nil), "umee.ibcratelimit.v1beta1.QueryRateLimitsOfIBCDenoms")
	proto.RegisterType((*QueryRateLimitsOfIBCDenomsResponse)(nil), "umee.ibcratelimit.v1beta1.QueryRateLimitsOfIBCDenomsResponse")
	proto.RegisterType((*QueryRateLimitsOfIBCDenom)(nil), "umee.ibcratelimit.v1beta1.QueryRateLimitsOfIBCDenom")
	proto.RegisterType((*QueryRateLimitsOfIBCDenomResponse)(nil), "umee.ibcratelimit.v1beta1.QueryRateLimitsOfIBCDenomResponse")
}

func init() {
	proto.RegisterFile("umee/ibc_rate_limit/v1beta1/query.proto", fileDescriptor_1706951c856d331f)
}

var fileDescriptor_1706951c856d331f = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x63, 0x18, 0x15, 0x7d, 0x2b, 0x2e, 0x66, 0x48, 0x2c, 0x4c, 0x66, 0x33, 0x30, 0x26,
	0xa1, 0xc6, 0xb4, 0xe3, 0xdf, 0x01, 0x84, 0x54, 0xb8, 0x4c, 0x20, 0x01, 0x15, 0xe2, 0xc0, 0x65,
	0x72, 0x8a, 0x09, 0x16, 0x4b, 0x9c, 0x25, 0xee, 0x60, 0x42, 0x5c, 0x38, 0x71, 0x44, 0x42, 0x7c,
	0x19, 0xae, 0x5c, 0x7a, 0x9c, 0xc4, 0x85, 0x13, 0x82, 0x96, 0x0f, 0x82, 0xec, 0x04, 0xaf, 0x48,
	0x49, 0x19, 0xbd, 0xf9, 0xcf, 0xfb, 0x3c, 0xcf, 0x2f, 0x7e, 0x5f, 0x05, 0x2e, 0x0e, 0x63, 0x21,
	0x98, 0x0c, 0x07, 0x5b, 0x19, 0xd7, 0x62, 0x6b, 0x5b, 0xc6, 0x52, 0xb3, 0xdd, 0x4e, 0x28, 0x34,
	0xef, 0xb0, 0x9d, 0xa1, 0xc8, 0xf6, 0x82, 0x34, 0x53, 0x5a, 0xe1, 0x25, 0x53, 0x18, 0xc8, 0x70,
	0x60, 0xea, 0x6c, 0x59, 0x50, 0x96, 0xf9, 0xcb, 0x91, 0x52, 0xd1, 0xb6, 0x60, 0x3c, 0x95, 0x8c,
	0x27, 0x89, 0xd2, 0x5c, 0x4b, 0x95, 0xe4, 0x85, 0xd0, 0x5f, 0x8c, 0x54, 0xa4, 0xec, 0x92, 0x99,
	0x55, 0x79, 0x7a, 0x79, 0x56, 0xee, 0xdf, 0xc7, 0x85, 0x82, 0x9e, 0x80, 0xd6, 0x23, 0xc3, 0xf3,
	0x90, 0x67, 0x3c, 0xce, 0xe9, 0x13, 0x38, 0x39, 0xb5, 0xed, 0x8b, 0x3c, 0x55, 0x49, 0x2e, 0xf0,
	0x6d, 0x68, 0xa4, 0xf6, 0xe4, 0x34, 0x5a, 0x41, 0xeb, 0xad, 0xee, 0x6a, 0x50, 0xcb, 0x1d, 0x14,
	0xd2, 0xde, 0xc2, 0xe8, 0xfb, 0x59, 0xaf, 0x5f, 0xca, 0xe8, 0x32, 0xf8, 0xd6, 0xb7, 0xcf, 0xb5,
	0xb8, 0x6f, 0xca, 0xf3, 0x07, 0xcf, 0x37, 0x7b, 0x77, 0xee, 0x8a, 0x44, 0xc5, 0x39, 0xdd, 0x01,
	0x5a, 0x7f, 0xeb, 0x20, 0xee, 0x41, 0xeb, 0x00, 0xdf, 0x90, 0x1c, 0x5d, 0x6f, 0x75, 0xcf, 0xcf,
	0x20, 0x71, 0x76, 0x25, 0x0c, 0x64, 0xce, 0x9f, 0xde, 0x80, 0xa5, 0xda, 0x48, 0x7c, 0x06, 0x9a,
	0xe6, 0xb1, 0x9e, 0x99, 0x8d, 0xfd, 0xe2, 0x66, 0xff, 0xb8, 0x0c, 0x07, 0xf6, 0x92, 0x26, 0xb0,
	0x5a, 0xab, 0x74, 0xac, 0x9b, 0x00, 0x07, 0xac, 0xe5, 0xa3, 0xfd, 0x0f, 0x6a, 0xd3, 0xa1, 0x76,
	0x3f, 0x2d, 0xc0, 0x31, 0x1b, 0x88, 0xdf, 0x23, 0x68, 0x14, 0xaf, 0x8b, 0xd7, 0x66, 0x78, 0x4d,
	0x35, 0xd0, 0x0f, 0x0e, 0x57, 0xf7, 0x87, 0x9b, 0xae, 0xbd, 0xfb, 0xfa, 0xeb, 0xe3, 0x91, 0x15,
	0x4c, 0x58, 0xf5, 0x24, 0xb1, 0xa2, 0x9f, 0xf8, 0x33, 0x82, 0x53, 0x95, 0xdd, 0xc2, 0x57, 0xff,
	0x95, 0x58, 0x29, 0xf3, 0x6f, 0xcd, 0x25, 0x73, 0xdc, 0x97, 0x2c, 0xf7, 0x05, 0x7c, 0xae, 0x8e,
	0x7b, 0x6a, 0x72, 0xf0, 0x17, 0x04, 0x8b, 0x95, 0x7d, 0xbf, 0x32, 0x0f, 0x84, 0x7f, 0x73, 0x1e,
	0x95, 0x23, 0xbf, 0x6e, 0xc9, 0x3b, 0x98, 0x1d, 0x82, 0x9c, 0xbd, 0x71, 0x63, 0xf9, 0xb6, 0xf7,
	0x78, 0xf4, 0x93, 0x78, 0xa3, 0x31, 0x41, 0xfb, 0x63, 0x82, 0x7e, 0x8c, 0x09, 0xfa, 0x30, 0x21,
	0xde, 0xfe, 0x84, 0x78, 0xdf, 0x26, 0xc4, 0x7b, 0x7a, 0x2d, 0x92, 0xfa, 0xc5, 0x30, 0x0c, 0x06,
	0x2a, 0xb6, 0xc6, 0xed, 0x44, 0xe8, 0x57, 0x2a, 0x7b, 0x59, 0xa4, 0xec, 0x6e, 0xb0, 0xd7, 0x26,
	0xaa, 0x6d, 0xcc, 0xdb, 0x45, 0x94, 0xde, 0x4b, 0x45, 0x1e, 0x36, 0xec, 0x6f, 0x61, 0xe3, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x53, 0x22, 0x5a, 0xc2, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries the parameters of the x/ibc-rate-limit module.
	Params(ctx context.Context, in *QueryParams, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// RateLimitsOfIBCDenoms queries the rate limits of ibc denoms.
	RateLimitsOfIBCDenoms(ctx context.Context, in *QueryRateLimitsOfIBCDenoms, opts ...grpc.CallOption) (*QueryRateLimitsOfIBCDenomsResponse, error)
	// RateLimitsOfIBCDenom queries the rate limits of ibc denom.
	RateLimitsOfIBCDenom(ctx context.Context, in *QueryRateLimitsOfIBCDenom, opts ...grpc.CallOption) (*QueryRateLimitsOfIBCDenomResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParams, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/umee.ibcratelimit.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RateLimitsOfIBCDenoms(ctx context.Context, in *QueryRateLimitsOfIBCDenoms, opts ...grpc.CallOption) (*QueryRateLimitsOfIBCDenomsResponse, error) {
	out := new(QueryRateLimitsOfIBCDenomsResponse)
	err := c.cc.Invoke(ctx, "/umee.ibcratelimit.v1beta1.Query/RateLimitsOfIBCDenoms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RateLimitsOfIBCDenom(ctx context.Context, in *QueryRateLimitsOfIBCDenom, opts ...grpc.CallOption) (*QueryRateLimitsOfIBCDenomResponse, error) {
	out := new(QueryRateLimitsOfIBCDenomResponse)
	err := c.cc.Invoke(ctx, "/umee.ibcratelimit.v1beta1.Query/RateLimitsOfIBCDenom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries the parameters of the x/ibc-rate-limit module.
	Params(context.Context, *QueryParams) (*QueryParamsResponse, error)
	// RateLimitsOfIBCDenoms queries the rate limits of ibc denoms.
	RateLimitsOfIBCDenoms(context.Context, *QueryRateLimitsOfIBCDenoms) (*QueryRateLimitsOfIBCDenomsResponse, error)
	// RateLimitsOfIBCDenom queries the rate limits of ibc denom.
	RateLimitsOfIBCDenom(context.Context, *QueryRateLimitsOfIBCDenom) (*QueryRateLimitsOfIBCDenomResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParams) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) RateLimitsOfIBCDenoms(ctx context.Context, req *QueryRateLimitsOfIBCDenoms) (*QueryRateLimitsOfIBCDenomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateLimitsOfIBCDenoms not implemented")
}
func (*UnimplementedQueryServer) RateLimitsOfIBCDenom(ctx context.Context, req *QueryRateLimitsOfIBCDenom) (*QueryRateLimitsOfIBCDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateLimitsOfIBCDenom not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.ibcratelimit.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RateLimitsOfIBCDenoms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRateLimitsOfIBCDenoms)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RateLimitsOfIBCDenoms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.ibcratelimit.v1beta1.Query/RateLimitsOfIBCDenoms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RateLimitsOfIBCDenoms(ctx, req.(*QueryRateLimitsOfIBCDenoms))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RateLimitsOfIBCDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRateLimitsOfIBCDenom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RateLimitsOfIBCDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.ibcratelimit.v1beta1.Query/RateLimitsOfIBCDenom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RateLimitsOfIBCDenom(ctx, req.(*QueryRateLimitsOfIBCDenom))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "umee.ibcratelimit.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "RateLimitsOfIBCDenoms",
			Handler:    _Query_RateLimitsOfIBCDenoms_Handler,
		},
		{
			MethodName: "RateLimitsOfIBCDenom",
			Handler:    _Query_RateLimitsOfIBCDenom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umee/ibc_rate_limit/v1beta1/query.proto",
}

func (m *QueryParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryRateLimitsOfIBCDenoms) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRateLimitsOfIBCDenoms) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRateLimitsOfIBCDenoms) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryRateLimitsOfIBCDenomsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRateLimitsOfIBCDenomsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRateLimitsOfIBCDenomsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RateLimits) > 0 {
		for iNdEx := len(m.RateLimits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RateLimits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryRateLimitsOfIBCDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRateLimitsOfIBCDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRateLimitsOfIBCDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IbcDenom) > 0 {
		i -= len(m.IbcDenom)
		copy(dAtA[i:], m.IbcDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.IbcDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRateLimitsOfIBCDenomResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRateLimitsOfIBCDenomResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRateLimitsOfIBCDenomResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.RateLimit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryRateLimitsOfIBCDenoms) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryRateLimitsOfIBCDenomsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RateLimits) > 0 {
		for _, e := range m.RateLimits {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryRateLimitsOfIBCDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IbcDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRateLimitsOfIBCDenomResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.RateLimit.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryRateLimitsOfIBCDenoms) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRateLimitsOfIBCDenoms: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRateLimitsOfIBCDenoms: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryRateLimitsOfIBCDenomsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRateLimitsOfIBCDenomsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRateLimitsOfIBCDenomsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RateLimits = append(m.RateLimits, RateLimit{})
			if err := m.RateLimits[len(m.RateLimits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryRateLimitsOfIBCDenom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRateLimitsOfIBCDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRateLimitsOfIBCDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryRateLimitsOfIBCDenomResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRateLimitsOfIBCDenomResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRateLimitsOfIBCDenomResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RateLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
