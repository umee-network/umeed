// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/ugov/v1/tx.proto

package ugov

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgGovUpdateMinGasPrice request type.
type MsgGovUpdateMinGasPrice struct {
	// authority must be the address of the governance account.
	Authority   string        `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	MinGasPrice types.DecCoin `protobuf:"bytes,2,opt,name=min_gas_price,json=minGasPrice,proto3" json:"min_gas_price"`
}

func (m *MsgGovUpdateMinGasPrice) Reset()      { *m = MsgGovUpdateMinGasPrice{} }
func (*MsgGovUpdateMinGasPrice) ProtoMessage() {}
func (*MsgGovUpdateMinGasPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffc07de1c6ee91b, []int{0}
}
func (m *MsgGovUpdateMinGasPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovUpdateMinGasPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovUpdateMinGasPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovUpdateMinGasPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovUpdateMinGasPrice.Merge(m, src)
}
func (m *MsgGovUpdateMinGasPrice) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovUpdateMinGasPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovUpdateMinGasPrice.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovUpdateMinGasPrice proto.InternalMessageInfo

// MsgGovUpdateMinGasPriceResponse response type.
type MsgGovUpdateMinGasPriceResponse struct {
}

func (m *MsgGovUpdateMinGasPriceResponse) Reset()         { *m = MsgGovUpdateMinGasPriceResponse{} }
func (m *MsgGovUpdateMinGasPriceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGovUpdateMinGasPriceResponse) ProtoMessage()    {}
func (*MsgGovUpdateMinGasPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffc07de1c6ee91b, []int{1}
}
func (m *MsgGovUpdateMinGasPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovUpdateMinGasPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovUpdateMinGasPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovUpdateMinGasPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovUpdateMinGasPriceResponse.Merge(m, src)
}
func (m *MsgGovUpdateMinGasPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovUpdateMinGasPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovUpdateMinGasPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovUpdateMinGasPriceResponse proto.InternalMessageInfo

// MsgGovSetEmergencyGroup request type.
type MsgGovSetEmergencyGroup struct {
	// authority must be the address of the governance account.
	Authority      string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	EmergencyGroup string `protobuf:"bytes,2,opt,name=emergency_group,json=emergencyGroup,proto3" json:"emergency_group,omitempty"`
}

func (m *MsgGovSetEmergencyGroup) Reset()         { *m = MsgGovSetEmergencyGroup{} }
func (m *MsgGovSetEmergencyGroup) String() string { return proto.CompactTextString(m) }
func (*MsgGovSetEmergencyGroup) ProtoMessage()    {}
func (*MsgGovSetEmergencyGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffc07de1c6ee91b, []int{2}
}
func (m *MsgGovSetEmergencyGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovSetEmergencyGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovSetEmergencyGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovSetEmergencyGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovSetEmergencyGroup.Merge(m, src)
}
func (m *MsgGovSetEmergencyGroup) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovSetEmergencyGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovSetEmergencyGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovSetEmergencyGroup proto.InternalMessageInfo

// MsgGovSetEmergencyGroupResponse response type.
type MsgGovSetEmergencyGroupResponse struct {
}

func (m *MsgGovSetEmergencyGroupResponse) Reset()         { *m = MsgGovSetEmergencyGroupResponse{} }
func (m *MsgGovSetEmergencyGroupResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGovSetEmergencyGroupResponse) ProtoMessage()    {}
func (*MsgGovSetEmergencyGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffc07de1c6ee91b, []int{3}
}
func (m *MsgGovSetEmergencyGroupResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovSetEmergencyGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovSetEmergencyGroupResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovSetEmergencyGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovSetEmergencyGroupResponse.Merge(m, src)
}
func (m *MsgGovSetEmergencyGroupResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovSetEmergencyGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovSetEmergencyGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovSetEmergencyGroupResponse proto.InternalMessageInfo

// MsgGovUpdateInflationParams request type.
type MsgGovUpdateInflationParams struct {
	// authority must be the address of the governance account.
	Authority string          `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	Params    InflationParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgGovUpdateInflationParams) Reset()         { *m = MsgGovUpdateInflationParams{} }
func (m *MsgGovUpdateInflationParams) String() string { return proto.CompactTextString(m) }
func (*MsgGovUpdateInflationParams) ProtoMessage()    {}
func (*MsgGovUpdateInflationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffc07de1c6ee91b, []int{4}
}
func (m *MsgGovUpdateInflationParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovUpdateInflationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovUpdateInflationParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovUpdateInflationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovUpdateInflationParams.Merge(m, src)
}
func (m *MsgGovUpdateInflationParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovUpdateInflationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovUpdateInflationParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovUpdateInflationParams proto.InternalMessageInfo

// GovUpdateInflationParamsResponse response type.
type GovUpdateInflationParamsResponse struct {
}

func (m *GovUpdateInflationParamsResponse) Reset()         { *m = GovUpdateInflationParamsResponse{} }
func (m *GovUpdateInflationParamsResponse) String() string { return proto.CompactTextString(m) }
func (*GovUpdateInflationParamsResponse) ProtoMessage()    {}
func (*GovUpdateInflationParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffc07de1c6ee91b, []int{5}
}
func (m *GovUpdateInflationParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GovUpdateInflationParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GovUpdateInflationParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GovUpdateInflationParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GovUpdateInflationParamsResponse.Merge(m, src)
}
func (m *GovUpdateInflationParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GovUpdateInflationParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GovUpdateInflationParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GovUpdateInflationParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgGovUpdateMinGasPrice)(nil), "umee.ugov.v1.MsgGovUpdateMinGasPrice")
	proto.RegisterType((*MsgGovUpdateMinGasPriceResponse)(nil), "umee.ugov.v1.MsgGovUpdateMinGasPriceResponse")
	proto.RegisterType((*MsgGovSetEmergencyGroup)(nil), "umee.ugov.v1.MsgGovSetEmergencyGroup")
	proto.RegisterType((*MsgGovSetEmergencyGroupResponse)(nil), "umee.ugov.v1.MsgGovSetEmergencyGroupResponse")
	proto.RegisterType((*MsgGovUpdateInflationParams)(nil), "umee.ugov.v1.MsgGovUpdateInflationParams")
	proto.RegisterType((*GovUpdateInflationParamsResponse)(nil), "umee.ugov.v1.GovUpdateInflationParamsResponse")
}

func init() { proto.RegisterFile("umee/ugov/v1/tx.proto", fileDescriptor_9ffc07de1c6ee91b) }

var fileDescriptor_9ffc07de1c6ee91b = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0xdf, 0xad, 0x52, 0xe8, 0x54, 0x2b, 0x2c, 0x91, 0xc6, 0xa8, 0x9b, 0x18, 0x50, 0xaa, 0x90,
	0x19, 0x52, 0xa1, 0x87, 0x7a, 0x6a, 0xac, 0x06, 0x0f, 0x81, 0x92, 0xe2, 0xc5, 0x4b, 0x98, 0x6c,
	0x9e, 0xd3, 0xc1, 0xce, 0x4c, 0x98, 0x99, 0x5d, 0xdb, 0xab, 0x9f, 0xc0, 0xa3, 0x47, 0x11, 0xaf,
	0x82, 0x07, 0x3f, 0x44, 0x8e, 0xc5, 0x53, 0x4f, 0xa2, 0xc9, 0xc1, 0xaf, 0x21, 0xbb, 0x99, 0x34,
	0x4d, 0x93, 0x55, 0xe9, 0x6d, 0x67, 0xde, 0xef, 0xfd, 0xfe, 0xec, 0x7b, 0x0c, 0xba, 0x19, 0x0b,
	0x00, 0x12, 0x33, 0x95, 0x90, 0xa4, 0x4e, 0xec, 0x11, 0xee, 0x6b, 0x65, 0x55, 0x70, 0x2d, 0xbd,
	0xc6, 0xe9, 0x35, 0x4e, 0xea, 0xa5, 0x30, 0x52, 0x46, 0x28, 0x43, 0xba, 0xd4, 0x00, 0x49, 0xea,
	0x5d, 0xb0, 0xb4, 0x4e, 0x22, 0xc5, 0xe5, 0x18, 0x5d, 0x5a, 0x77, 0x75, 0x61, 0x58, 0xca, 0x22,
	0x0c, 0x73, 0x85, 0x5b, 0xe3, 0x42, 0x27, 0x3b, 0x91, 0xf1, 0xc1, 0x95, 0x0a, 0x4c, 0x31, 0x35,
	0xbe, 0x4f, 0xbf, 0x26, 0x4c, 0x33, 0x76, 0x32, 0xfd, 0xac, 0x50, 0xfd, 0xe2, 0xa3, 0xf5, 0x96,
	0x61, 0x4d, 0x95, 0xbc, 0xec, 0xf7, 0xa8, 0x85, 0x16, 0x97, 0x4d, 0x6a, 0xf6, 0x34, 0x8f, 0x20,
	0xd8, 0x42, 0x2b, 0x34, 0xb6, 0x07, 0x4a, 0x73, 0x7b, 0x5c, 0xf4, 0x2b, 0xfe, 0xc6, 0x4a, 0xa3,
	0xf8, 0xfd, 0x5b, 0xad, 0xe0, 0xf4, 0x76, 0x7a, 0x3d, 0x0d, 0xc6, 0xec, 0x5b, 0xcd, 0x25, 0x6b,
	0x4f, 0xa1, 0xc1, 0x73, 0x74, 0x5d, 0x70, 0xd9, 0x61, 0x34, 0x35, 0xc8, 0x23, 0x28, 0x2e, 0x55,
	0xfc, 0x8d, 0xd5, 0xcd, 0x3b, 0xd8, 0x35, 0xa6, 0x71, 0xb1, 0x8b, 0x8b, 0x77, 0x21, 0x7a, 0xaa,
	0xb8, 0x6c, 0x5c, 0x1d, 0xfc, 0x28, 0x7b, 0xed, 0x55, 0x31, 0xd5, 0xdf, 0x0e, 0x3e, 0x7c, 0x2c,
	0x7b, 0xef, 0x7e, 0x7f, 0x7d, 0x34, 0xe5, 0xae, 0xde, 0x43, 0xe5, 0x1c, 0xbb, 0x6d, 0x30, 0x7d,
	0x25, 0x0d, 0x54, 0x3f, 0x9f, 0x45, 0xda, 0x07, 0xfb, 0x4c, 0x80, 0x66, 0x20, 0xa3, 0xe3, 0xa6,
	0x56, 0x71, 0xff, 0xd2, 0x91, 0x76, 0xd0, 0x0d, 0x98, 0x30, 0x75, 0x58, 0x4a, 0x95, 0x85, 0xfa,
	0x5b, 0xf7, 0x1a, 0xcc, 0x48, 0x6f, 0xaf, 0xe5, 0x25, 0x99, 0x73, 0x79, 0x96, 0xe4, 0x93, 0x8f,
	0x6e, 0x9f, 0x4f, 0xfb, 0x42, 0xbe, 0x3e, 0xa4, 0x96, 0x2b, 0xb9, 0x47, 0x35, 0x15, 0xe6, 0xd2,
	0x69, 0x9e, 0xa0, 0xe5, 0x7e, 0xc6, 0xe0, 0x26, 0x73, 0x17, 0x9f, 0x5f, 0x4b, 0x7c, 0x41, 0xc6,
	0x8d, 0xc6, 0xb5, 0xcc, 0xe5, 0xa8, 0xa2, 0x4a, 0x9e, 0xc1, 0x49, 0x90, 0xcd, 0xd3, 0x25, 0x74,
	0xa5, 0x65, 0x58, 0x70, 0x88, 0x0a, 0x0b, 0x37, 0xed, 0xfe, 0xac, 0x81, 0x9c, 0x09, 0x97, 0x6a,
	0xff, 0x05, 0x9b, 0xa8, 0x3a, 0xb5, 0xf9, 0x25, 0x58, 0xa8, 0x36, 0x07, 0x5b, 0xac, 0x96, 0x3b,
	0xac, 0x20, 0x46, 0xc5, 0xdc, 0x41, 0x3d, 0xcc, 0x37, 0x7e, 0x01, 0x5a, 0xc2, 0xb3, 0xd0, 0x7f,
	0xfd, 0xda, 0xc6, 0xee, 0xe0, 0x57, 0xe8, 0x0d, 0x86, 0xa1, 0x7f, 0x32, 0x0c, 0xfd, 0x9f, 0xc3,
	0xd0, 0x7f, 0x3f, 0x0a, 0xbd, 0x93, 0x51, 0xe8, 0x9d, 0x8e, 0x42, 0xef, 0xd5, 0x03, 0xc6, 0xed,
	0x41, 0xdc, 0xc5, 0x91, 0x12, 0x24, 0xe5, 0xad, 0x49, 0xb0, 0x6f, 0x95, 0x7e, 0x93, 0x1d, 0x48,
	0xb2, 0x45, 0x8e, 0xb2, 0xc7, 0xa0, 0xbb, 0x9c, 0xbd, 0x06, 0x8f, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x72, 0xab, 0x12, 0x18, 0xb7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// GovUpdateMinGasPrice sets protocol controlled tx min fees.
	GovUpdateMinGasPrice(ctx context.Context, in *MsgGovUpdateMinGasPrice, opts ...grpc.CallOption) (*MsgGovUpdateMinGasPriceResponse, error)
	// GovSetEmergencyGroup sets emergency group address.
	GovSetEmergencyGroup(ctx context.Context, in *MsgGovSetEmergencyGroup, opts ...grpc.CallOption) (*MsgGovSetEmergencyGroupResponse, error)
	// GovUpdateInflationParams sets new params for inflation rate change.
	GovUpdateInflationParams(ctx context.Context, in *MsgGovUpdateInflationParams, opts ...grpc.CallOption) (*GovUpdateInflationParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) GovUpdateMinGasPrice(ctx context.Context, in *MsgGovUpdateMinGasPrice, opts ...grpc.CallOption) (*MsgGovUpdateMinGasPriceResponse, error) {
	out := new(MsgGovUpdateMinGasPriceResponse)
	err := c.cc.Invoke(ctx, "/umee.ugov.v1.Msg/GovUpdateMinGasPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GovSetEmergencyGroup(ctx context.Context, in *MsgGovSetEmergencyGroup, opts ...grpc.CallOption) (*MsgGovSetEmergencyGroupResponse, error) {
	out := new(MsgGovSetEmergencyGroupResponse)
	err := c.cc.Invoke(ctx, "/umee.ugov.v1.Msg/GovSetEmergencyGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GovUpdateInflationParams(ctx context.Context, in *MsgGovUpdateInflationParams, opts ...grpc.CallOption) (*GovUpdateInflationParamsResponse, error) {
	out := new(GovUpdateInflationParamsResponse)
	err := c.cc.Invoke(ctx, "/umee.ugov.v1.Msg/GovUpdateInflationParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// GovUpdateMinGasPrice sets protocol controlled tx min fees.
	GovUpdateMinGasPrice(context.Context, *MsgGovUpdateMinGasPrice) (*MsgGovUpdateMinGasPriceResponse, error)
	// GovSetEmergencyGroup sets emergency group address.
	GovSetEmergencyGroup(context.Context, *MsgGovSetEmergencyGroup) (*MsgGovSetEmergencyGroupResponse, error)
	// GovUpdateInflationParams sets new params for inflation rate change.
	GovUpdateInflationParams(context.Context, *MsgGovUpdateInflationParams) (*GovUpdateInflationParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) GovUpdateMinGasPrice(ctx context.Context, req *MsgGovUpdateMinGasPrice) (*MsgGovUpdateMinGasPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovUpdateMinGasPrice not implemented")
}
func (*UnimplementedMsgServer) GovSetEmergencyGroup(ctx context.Context, req *MsgGovSetEmergencyGroup) (*MsgGovSetEmergencyGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovSetEmergencyGroup not implemented")
}
func (*UnimplementedMsgServer) GovUpdateInflationParams(ctx context.Context, req *MsgGovUpdateInflationParams) (*GovUpdateInflationParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovUpdateInflationParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_GovUpdateMinGasPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovUpdateMinGasPrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovUpdateMinGasPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.ugov.v1.Msg/GovUpdateMinGasPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovUpdateMinGasPrice(ctx, req.(*MsgGovUpdateMinGasPrice))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GovSetEmergencyGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovSetEmergencyGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovSetEmergencyGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.ugov.v1.Msg/GovSetEmergencyGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovSetEmergencyGroup(ctx, req.(*MsgGovSetEmergencyGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GovUpdateInflationParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovUpdateInflationParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovUpdateInflationParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.ugov.v1.Msg/GovUpdateInflationParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovUpdateInflationParams(ctx, req.(*MsgGovUpdateInflationParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "umee.ugov.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GovUpdateMinGasPrice",
			Handler:    _Msg_GovUpdateMinGasPrice_Handler,
		},
		{
			MethodName: "GovSetEmergencyGroup",
			Handler:    _Msg_GovSetEmergencyGroup_Handler,
		},
		{
			MethodName: "GovUpdateInflationParams",
			Handler:    _Msg_GovUpdateInflationParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umee/ugov/v1/tx.proto",
}

func (m *MsgGovUpdateMinGasPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovUpdateMinGasPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovUpdateMinGasPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgGovUpdateMinGasPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovUpdateMinGasPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovUpdateMinGasPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgGovSetEmergencyGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovSetEmergencyGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovSetEmergencyGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EmergencyGroup) > 0 {
		i -= len(m.EmergencyGroup)
		copy(dAtA[i:], m.EmergencyGroup)
		i = encodeVarintTx(dAtA, i, uint64(len(m.EmergencyGroup)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgGovSetEmergencyGroupResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovSetEmergencyGroupResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovSetEmergencyGroupResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgGovUpdateInflationParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovUpdateInflationParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovUpdateInflationParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GovUpdateInflationParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GovUpdateInflationParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GovUpdateInflationParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgGovUpdateMinGasPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.MinGasPrice.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgGovUpdateMinGasPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgGovSetEmergencyGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.EmergencyGroup)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgGovSetEmergencyGroupResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgGovUpdateInflationParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *GovUpdateInflationParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgGovUpdateMinGasPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgGovUpdateMinGasPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovUpdateMinGasPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGasPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgGovUpdateMinGasPriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgGovUpdateMinGasPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovUpdateMinGasPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgGovSetEmergencyGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgGovSetEmergencyGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovSetEmergencyGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmergencyGroup", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EmergencyGroup = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgGovSetEmergencyGroupResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgGovSetEmergencyGroupResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovSetEmergencyGroupResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgGovUpdateInflationParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgGovUpdateInflationParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovUpdateInflationParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *GovUpdateInflationParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: GovUpdateInflationParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GovUpdateInflationParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
