// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/ugov/v1/query.proto

package ugov

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
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

// QueryMinGasPrice is a request type.
type QueryMinGasPrice struct {
}

func (m *QueryMinGasPrice) Reset()         { *m = QueryMinGasPrice{} }
func (m *QueryMinGasPrice) String() string { return proto.CompactTextString(m) }
func (*QueryMinGasPrice) ProtoMessage()    {}
func (*QueryMinGasPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_25fa04679024a47d, []int{0}
}
func (m *QueryMinGasPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMinGasPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMinGasPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMinGasPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMinGasPrice.Merge(m, src)
}
func (m *QueryMinGasPrice) XXX_Size() int {
	return m.Size()
}
func (m *QueryMinGasPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMinGasPrice.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMinGasPrice proto.InternalMessageInfo

// QueryMinGasPriceResponse response type.
type QueryMinGasPriceResponse struct {
	MinGasPrice types.DecCoin `protobuf:"bytes,1,opt,name=min_gas_price,json=minGasPrice,proto3" json:"min_gas_price"`
}

func (m *QueryMinGasPriceResponse) Reset()         { *m = QueryMinGasPriceResponse{} }
func (m *QueryMinGasPriceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMinGasPriceResponse) ProtoMessage()    {}
func (*QueryMinGasPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_25fa04679024a47d, []int{1}
}
func (m *QueryMinGasPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMinGasPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMinGasPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMinGasPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMinGasPriceResponse.Merge(m, src)
}
func (m *QueryMinGasPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryMinGasPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMinGasPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMinGasPriceResponse proto.InternalMessageInfo

// QueryEmergencyGroup request type.
type QueryEmergencyGroup struct {
}

func (m *QueryEmergencyGroup) Reset()         { *m = QueryEmergencyGroup{} }
func (m *QueryEmergencyGroup) String() string { return proto.CompactTextString(m) }
func (*QueryEmergencyGroup) ProtoMessage()    {}
func (*QueryEmergencyGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_25fa04679024a47d, []int{2}
}
func (m *QueryEmergencyGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEmergencyGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEmergencyGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEmergencyGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEmergencyGroup.Merge(m, src)
}
func (m *QueryEmergencyGroup) XXX_Size() int {
	return m.Size()
}
func (m *QueryEmergencyGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEmergencyGroup.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEmergencyGroup proto.InternalMessageInfo

// QueryEmergencyGroupResponse response type.
type QueryEmergencyGroupResponse struct {
	EmergencyGroup string `protobuf:"bytes,1,opt,name=emergency_group,json=emergencyGroup,proto3" json:"emergency_group,omitempty"`
}

func (m *QueryEmergencyGroupResponse) Reset()         { *m = QueryEmergencyGroupResponse{} }
func (m *QueryEmergencyGroupResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEmergencyGroupResponse) ProtoMessage()    {}
func (*QueryEmergencyGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_25fa04679024a47d, []int{3}
}
func (m *QueryEmergencyGroupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEmergencyGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEmergencyGroupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEmergencyGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEmergencyGroupResponse.Merge(m, src)
}
func (m *QueryEmergencyGroupResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEmergencyGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEmergencyGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEmergencyGroupResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryMinGasPrice)(nil), "umee.ugov.v1.QueryMinGasPrice")
	proto.RegisterType((*QueryMinGasPriceResponse)(nil), "umee.ugov.v1.QueryMinGasPriceResponse")
	proto.RegisterType((*QueryEmergencyGroup)(nil), "umee.ugov.v1.QueryEmergencyGroup")
	proto.RegisterType((*QueryEmergencyGroupResponse)(nil), "umee.ugov.v1.QueryEmergencyGroupResponse")
}

func init() { proto.RegisterFile("umee/ugov/v1/query.proto", fileDescriptor_25fa04679024a47d) }

var fileDescriptor_25fa04679024a47d = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x93, 0x0a, 0x90, 0xf0, 0x42, 0x41, 0xa6, 0x48, 0x21, 0x6d, 0x5d, 0x08, 0xa2, 0x82,
	0x43, 0x6d, 0xa5, 0x88, 0x07, 0xe8, 0x52, 0xe8, 0x09, 0x09, 0x96, 0x1b, 0x97, 0x25, 0x49, 0x47,
	0xc6, 0x82, 0x78, 0x82, 0x9d, 0x04, 0xca, 0x11, 0x89, 0x13, 0x17, 0x24, 0x5e, 0x85, 0x87, 0xd8,
	0x63, 0x05, 0x17, 0x4e, 0x08, 0x76, 0x79, 0x10, 0x14, 0x27, 0xad, 0x36, 0xab, 0x15, 0xbd, 0xc5,
	0xf3, 0x4f, 0xfe, 0xff, 0xf3, 0x8c, 0x49, 0x50, 0xe5, 0x00, 0xa2, 0x92, 0x58, 0x8b, 0x3a, 0x16,
	0x6f, 0x2b, 0x30, 0x47, 0xbc, 0x30, 0x58, 0x22, 0xbd, 0xd4, 0x28, 0xbc, 0x51, 0x78, 0x1d, 0x87,
	0x2c, 0x43, 0x9b, 0xa3, 0x15, 0x69, 0x62, 0x41, 0xd4, 0x71, 0x0a, 0x65, 0x12, 0x8b, 0x0c, 0x95,
	0x6e, 0xbb, 0xc3, 0x1b, 0xad, 0x3e, 0x76, 0x27, 0xd1, 0x1e, 0x3a, 0x69, 0x4d, 0xa2, 0xc4, 0xb6,
	0xde, 0x7c, 0x75, 0xd5, 0x0d, 0x89, 0x28, 0xdf, 0x80, 0x48, 0x0a, 0x25, 0x12, 0xad, 0xb1, 0x4c,
	0x4a, 0x85, 0xba, 0xfb, 0x27, 0xa2, 0xe4, 0xea, 0xb3, 0x86, 0xe5, 0x89, 0xd2, 0x07, 0x89, 0x7d,
	0x6a, 0x54, 0x06, 0x51, 0x4a, 0x82, 0xc5, 0xda, 0x08, 0x6c, 0x81, 0xda, 0x02, 0x7d, 0x4c, 0x2e,
	0xe7, 0x4a, 0x8f, 0x65, 0xd2, 0x10, 0xa8, 0x0c, 0x02, 0xff, 0xa6, 0x7f, 0x77, 0xb0, 0xbb, 0xc1,
	0x3b, 0x92, 0x06, 0x9b, 0x77, 0xd8, 0x7c, 0x1f, 0xb2, 0x87, 0xa8, 0xf4, 0xf0, 0xdc, 0xe4, 0xd7,
	0x96, 0x37, 0x1a, 0xe4, 0x73, 0x19, 0xd7, 0xc9, 0x35, 0x97, 0xf1, 0x28, 0x07, 0x23, 0x41, 0x67,
	0x47, 0x07, 0x06, 0xab, 0x22, 0x7a, 0x49, 0xd6, 0x97, 0x94, 0x4f, 0xd3, 0xf7, 0xc8, 0x15, 0x38,
	0x51, 0xc6, 0xb2, 0x91, 0x5c, 0xfe, 0xc5, 0x61, 0xf0, 0xfd, 0xdb, 0xce, 0x5a, 0x87, 0xb0, 0x77,
	0x78, 0x68, 0xc0, 0xda, 0xe7, 0xa5, 0x51, 0x5a, 0x8e, 0x56, 0xa1, 0x67, 0xb5, 0xfb, 0x79, 0x85,
	0x9c, 0x77, 0x11, 0xf4, 0x03, 0x19, 0xcc, 0xdd, 0x90, 0x32, 0x3e, 0xbf, 0x07, 0xbe, 0x38, 0x81,
	0x70, 0xfb, 0xff, 0xfa, 0x09, 0x63, 0x74, 0xfb, 0xe3, 0x8f, 0xbf, 0x5f, 0x57, 0x36, 0xe9, 0xba,
	0xe8, 0x6d, 0xbc, 0x37, 0x35, 0xfa, 0xc9, 0x27, 0xab, 0xfd, 0x3b, 0xd2, 0x5b, 0x4b, 0xfc, 0xfb,
	0x2d, 0xe1, 0xbd, 0x33, 0x5b, 0x4e, 0x29, 0xee, 0x38, 0x8a, 0x2d, 0xba, 0xd9, 0xa7, 0x58, 0x98,
	0xde, 0x70, 0x7f, 0xf2, 0x87, 0x79, 0x93, 0x29, 0xf3, 0x8f, 0xa7, 0xcc, 0xff, 0x3d, 0x65, 0xfe,
	0x97, 0x19, 0xf3, 0x8e, 0x67, 0xcc, 0xfb, 0x39, 0x63, 0xde, 0x8b, 0x6d, 0xa9, 0xca, 0x57, 0x55,
	0xca, 0x33, 0xcc, 0x9d, 0xcd, 0x8e, 0x86, 0xf2, 0x1d, 0x9a, 0xd7, 0xad, 0x67, 0xfd, 0x40, 0xbc,
	0x77, 0xc6, 0xe9, 0x05, 0xf7, 0x96, 0xee, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xce, 0xec, 0x9a,
	0x05, 0xe4, 0x02, 0x00, 0x00,
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
	// MinGasPrice returns minimum transaction fees.
	MinGasPrice(ctx context.Context, in *QueryMinGasPrice, opts ...grpc.CallOption) (*QueryMinGasPriceResponse, error)
	// EmergencyGroup returns emergency group address
	EmergencyGroup(ctx context.Context, in *QueryEmergencyGroup, opts ...grpc.CallOption) (*QueryEmergencyGroupResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) MinGasPrice(ctx context.Context, in *QueryMinGasPrice, opts ...grpc.CallOption) (*QueryMinGasPriceResponse, error) {
	out := new(QueryMinGasPriceResponse)
	err := c.cc.Invoke(ctx, "/umee.ugov.v1.Query/MinGasPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EmergencyGroup(ctx context.Context, in *QueryEmergencyGroup, opts ...grpc.CallOption) (*QueryEmergencyGroupResponse, error) {
	out := new(QueryEmergencyGroupResponse)
	err := c.cc.Invoke(ctx, "/umee.ugov.v1.Query/EmergencyGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// MinGasPrice returns minimum transaction fees.
	MinGasPrice(context.Context, *QueryMinGasPrice) (*QueryMinGasPriceResponse, error)
	// EmergencyGroup returns emergency group address
	EmergencyGroup(context.Context, *QueryEmergencyGroup) (*QueryEmergencyGroupResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) MinGasPrice(ctx context.Context, req *QueryMinGasPrice) (*QueryMinGasPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MinGasPrice not implemented")
}
func (*UnimplementedQueryServer) EmergencyGroup(ctx context.Context, req *QueryEmergencyGroup) (*QueryEmergencyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmergencyGroup not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_MinGasPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMinGasPrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MinGasPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.ugov.v1.Query/MinGasPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MinGasPrice(ctx, req.(*QueryMinGasPrice))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EmergencyGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEmergencyGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EmergencyGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.ugov.v1.Query/EmergencyGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EmergencyGroup(ctx, req.(*QueryEmergencyGroup))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "umee.ugov.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MinGasPrice",
			Handler:    _Query_MinGasPrice_Handler,
		},
		{
			MethodName: "EmergencyGroup",
			Handler:    _Query_EmergencyGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umee/ugov/v1/query.proto",
}

func (m *QueryMinGasPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMinGasPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMinGasPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryMinGasPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMinGasPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMinGasPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MinGasPrice.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryEmergencyGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEmergencyGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEmergencyGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryEmergencyGroupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEmergencyGroupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEmergencyGroupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EmergencyGroup) > 0 {
		i -= len(m.EmergencyGroup)
		copy(dAtA[i:], m.EmergencyGroup)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EmergencyGroup)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *QueryMinGasPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryMinGasPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinGasPrice.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryEmergencyGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryEmergencyGroupResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EmergencyGroup)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryMinGasPrice) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMinGasPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMinGasPrice: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryMinGasPriceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMinGasPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMinGasPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGasPrice", wireType)
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
			if err := m.MinGasPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryEmergencyGroup) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEmergencyGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEmergencyGroup: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryEmergencyGroupResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryEmergencyGroupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEmergencyGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmergencyGroup", wireType)
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
			m.EmergencyGroup = string(dAtA[iNdEx:postIndex])
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
