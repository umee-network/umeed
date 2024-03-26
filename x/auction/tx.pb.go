// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/auction/v1/tx.proto

package auction

import (
	context "context"
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// MsgGovSetRewardsParams updates rewards auction parameters.
type MsgGovSetRewardsParams struct {
	// authority must be the address of the governance account.
	Authority string        `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	Params    RewardsParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgGovSetRewardsParams) Reset()         { *m = MsgGovSetRewardsParams{} }
func (m *MsgGovSetRewardsParams) String() string { return proto.CompactTextString(m) }
func (*MsgGovSetRewardsParams) ProtoMessage()    {}
func (*MsgGovSetRewardsParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_44a5dea2889d94ea, []int{0}
}
func (m *MsgGovSetRewardsParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovSetRewardsParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovSetRewardsParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovSetRewardsParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovSetRewardsParams.Merge(m, src)
}
func (m *MsgGovSetRewardsParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovSetRewardsParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovSetRewardsParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovSetRewardsParams proto.InternalMessageInfo

func (*MsgGovSetRewardsParams) XXX_MessageName() string {
	return "umee.auction.v1.MsgGovSetRewardsParams"
}

// MsgGovSetRewardsParamsResponse defines the Msg/GovSetRewardsParams response type.
type MsgGovSetRewardsParamsResponse struct {
}

func (m *MsgGovSetRewardsParamsResponse) Reset()         { *m = MsgGovSetRewardsParamsResponse{} }
func (m *MsgGovSetRewardsParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGovSetRewardsParamsResponse) ProtoMessage()    {}
func (*MsgGovSetRewardsParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44a5dea2889d94ea, []int{1}
}
func (m *MsgGovSetRewardsParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovSetRewardsParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovSetRewardsParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovSetRewardsParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovSetRewardsParamsResponse.Merge(m, src)
}
func (m *MsgGovSetRewardsParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovSetRewardsParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovSetRewardsParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovSetRewardsParamsResponse proto.InternalMessageInfo

func (*MsgGovSetRewardsParamsResponse) XXX_MessageName() string {
	return "umee.auction.v1.MsgGovSetRewardsParamsResponse"
}

// MsgRewardsBid places a bid for a reword auction.
type MsgRewardsBid struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// the current auction ID being bid on. Fails if the ID is not an ID of the current auction.
	Id uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// amount of the bid in the base tokens
	Amount cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
}

func (m *MsgRewardsBid) Reset()         { *m = MsgRewardsBid{} }
func (m *MsgRewardsBid) String() string { return proto.CompactTextString(m) }
func (*MsgRewardsBid) ProtoMessage()    {}
func (*MsgRewardsBid) Descriptor() ([]byte, []int) {
	return fileDescriptor_44a5dea2889d94ea, []int{2}
}
func (m *MsgRewardsBid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRewardsBid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRewardsBid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRewardsBid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRewardsBid.Merge(m, src)
}
func (m *MsgRewardsBid) XXX_Size() int {
	return m.Size()
}
func (m *MsgRewardsBid) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRewardsBid.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRewardsBid proto.InternalMessageInfo

func (*MsgRewardsBid) XXX_MessageName() string {
	return "umee.auction.v1.MsgRewardsBid"
}

// MsgRewardsBidResponse response type for Msg/RewardsBid
type MsgRewardsBidResponse struct {
}

func (m *MsgRewardsBidResponse) Reset()         { *m = MsgRewardsBidResponse{} }
func (m *MsgRewardsBidResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRewardsBidResponse) ProtoMessage()    {}
func (*MsgRewardsBidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_44a5dea2889d94ea, []int{3}
}
func (m *MsgRewardsBidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRewardsBidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRewardsBidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRewardsBidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRewardsBidResponse.Merge(m, src)
}
func (m *MsgRewardsBidResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRewardsBidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRewardsBidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRewardsBidResponse proto.InternalMessageInfo

func (*MsgRewardsBidResponse) XXX_MessageName() string {
	return "umee.auction.v1.MsgRewardsBidResponse"
}
func init() {
	proto.RegisterType((*MsgGovSetRewardsParams)(nil), "umee.auction.v1.MsgGovSetRewardsParams")
	proto.RegisterType((*MsgGovSetRewardsParamsResponse)(nil), "umee.auction.v1.MsgGovSetRewardsParamsResponse")
	proto.RegisterType((*MsgRewardsBid)(nil), "umee.auction.v1.MsgRewardsBid")
	proto.RegisterType((*MsgRewardsBidResponse)(nil), "umee.auction.v1.MsgRewardsBidResponse")
}

func init() { proto.RegisterFile("umee/auction/v1/tx.proto", fileDescriptor_44a5dea2889d94ea) }

var fileDescriptor_44a5dea2889d94ea = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x33, 0x5b, 0x59, 0xe8, 0x94, 0x56, 0x18, 0xfb, 0x27, 0x2e, 0x74, 0xba, 0xec, 0x41,
	0xab, 0xd0, 0x19, 0x5a, 0xb1, 0x87, 0xe2, 0xc5, 0x5c, 0x44, 0x64, 0x41, 0x52, 0x4f, 0x5e, 0x24,
	0xdd, 0x19, 0x66, 0x87, 0x32, 0x99, 0x65, 0x66, 0x92, 0x56, 0xbc, 0xf9, 0x09, 0xfc, 0x04, 0x7e,
	0x06, 0x0f, 0x7e, 0x88, 0x1c, 0x3c, 0x14, 0x4f, 0xe2, 0xa1, 0xe8, 0xe6, 0xe0, 0xd7, 0x90, 0x24,
	0x13, 0xd6, 0xae, 0x41, 0xbc, 0xe5, 0xcd, 0xef, 0x7d, 0x9f, 0xe7, 0x79, 0x5f, 0x06, 0x86, 0x99,
	0xe2, 0x9c, 0x26, 0xd9, 0xc4, 0x49, 0x9d, 0xd2, 0xfc, 0x90, 0xba, 0x4b, 0x32, 0x33, 0xda, 0x69,
	0x74, 0xbb, 0x22, 0xc4, 0x13, 0x92, 0x1f, 0x0e, 0xee, 0x4e, 0xb4, 0x55, 0xda, 0xbe, 0xa9, 0x31,
	0x6d, 0x8a, 0xa6, 0x77, 0xb0, 0xd3, 0x54, 0x54, 0x59, 0x51, 0x69, 0x28, 0x2b, 0x3c, 0xd8, 0x14,
	0x5a, 0xe8, 0x66, 0xa0, 0xfa, 0xf2, 0x7f, 0x77, 0x97, 0x4d, 0x5b, 0x97, 0x1a, 0x8f, 0x3e, 0x02,
	0xb8, 0x3d, 0xb6, 0xe2, 0x99, 0xce, 0x4f, 0xb9, 0x8b, 0xf9, 0x45, 0x62, 0x98, 0x7d, 0x99, 0x98,
	0x44, 0x59, 0x74, 0x0c, 0x57, 0x93, 0xcc, 0x4d, 0xb5, 0x91, 0xee, 0x6d, 0x08, 0x86, 0x60, 0x7f,
	0x35, 0x0a, 0xbf, 0x7e, 0x3e, 0xd8, 0xf4, 0x69, 0x9e, 0x32, 0x66, 0xb8, 0xb5, 0xa7, 0xce, 0xc8,
	0x54, 0xc4, 0x8b, 0x56, 0xf4, 0x04, 0xf6, 0x67, 0xb5, 0x42, 0xd8, 0x1b, 0x82, 0xfd, 0xb5, 0x23,
	0x4c, 0x96, 0xb6, 0x23, 0x37, 0x7c, 0xa2, 0x5b, 0xc5, 0xf5, 0x5e, 0x10, 0xfb, 0x99, 0x93, 0x8d,
	0xf7, 0xbf, 0x3e, 0x3d, 0x5c, 0xa8, 0x8d, 0x86, 0x10, 0x77, 0xe7, 0x8b, 0xb9, 0x9d, 0xe9, 0xd4,
	0xf2, 0xd1, 0x3b, 0xb8, 0x3e, 0xb6, 0xc2, 0xb3, 0x48, 0x32, 0xb4, 0x0d, 0xfb, 0x96, 0xa7, 0x8c,
	0x9b, 0x26, 0x75, 0xec, 0x2b, 0xb4, 0x01, 0x7b, 0x92, 0xd5, 0xa1, 0xd6, 0xe3, 0x9e, 0x64, 0xe8,
	0x31, 0xec, 0x27, 0x4a, 0x67, 0xa9, 0x0b, 0x57, 0xea, 0xed, 0x76, 0xab, 0x20, 0xdf, 0xaf, 0xf7,
	0xb6, 0x9a, 0x0d, 0x2d, 0x3b, 0x27, 0x52, 0x53, 0x95, 0xb8, 0x29, 0x79, 0x9e, 0xba, 0xd8, 0x37,
	0x9f, 0xac, 0x55, 0x09, 0xbd, 0xe6, 0x68, 0x07, 0x6e, 0xdd, 0x30, 0x6f, 0x53, 0x1d, 0x7d, 0x01,
	0x70, 0x65, 0x6c, 0x05, 0xd2, 0xf0, 0x4e, 0xd7, 0x71, 0xef, 0xff, 0x75, 0x94, 0xee, 0x2d, 0x07,
	0xf4, 0x3f, 0x1b, 0x5b, 0x63, 0xf4, 0x0a, 0xc2, 0x3f, 0x6e, 0x81, 0xbb, 0xc6, 0x17, 0x7c, 0x70,
	0xef, 0xdf, 0xbc, 0x55, 0x8d, 0x5e, 0x14, 0x3f, 0x71, 0x50, 0xcc, 0x31, 0xb8, 0x9a, 0x63, 0xf0,
	0x63, 0x8e, 0xc1, 0x87, 0x12, 0x07, 0x45, 0x89, 0xc1, 0x55, 0x89, 0x83, 0x6f, 0x25, 0x0e, 0x5e,
	0x3f, 0x10, 0xd2, 0x4d, 0xb3, 0x33, 0x32, 0xd1, 0x8a, 0x56, 0x9a, 0x07, 0x29, 0x77, 0x17, 0xda,
	0x9c, 0xd7, 0x05, 0xcd, 0x8f, 0xe9, 0x65, 0xfb, 0xf4, 0xce, 0xfa, 0xf5, 0xdb, 0x7b, 0xf4, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x3c, 0x09, 0x52, 0x39, 0x11, 0x03, 0x00, 0x00,
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
	// Allows x/gov to update rewards auction parameters.
	GovSetRewardsParams(ctx context.Context, in *MsgGovSetRewardsParams, opts ...grpc.CallOption) (*MsgGovSetRewardsParamsResponse, error)
	// Places a bid for a reword auction. Must be higher than the previous bid by at least
	// RewardParams.RewardsParams.
	RewardsBid(ctx context.Context, in *MsgRewardsBid, opts ...grpc.CallOption) (*MsgRewardsBidResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) GovSetRewardsParams(ctx context.Context, in *MsgGovSetRewardsParams, opts ...grpc.CallOption) (*MsgGovSetRewardsParamsResponse, error) {
	out := new(MsgGovSetRewardsParamsResponse)
	err := c.cc.Invoke(ctx, "/umee.auction.v1.Msg/GovSetRewardsParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RewardsBid(ctx context.Context, in *MsgRewardsBid, opts ...grpc.CallOption) (*MsgRewardsBidResponse, error) {
	out := new(MsgRewardsBidResponse)
	err := c.cc.Invoke(ctx, "/umee.auction.v1.Msg/RewardsBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Allows x/gov to update rewards auction parameters.
	GovSetRewardsParams(context.Context, *MsgGovSetRewardsParams) (*MsgGovSetRewardsParamsResponse, error)
	// Places a bid for a reword auction. Must be higher than the previous bid by at least
	// RewardParams.RewardsParams.
	RewardsBid(context.Context, *MsgRewardsBid) (*MsgRewardsBidResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) GovSetRewardsParams(ctx context.Context, req *MsgGovSetRewardsParams) (*MsgGovSetRewardsParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovSetRewardsParams not implemented")
}
func (*UnimplementedMsgServer) RewardsBid(ctx context.Context, req *MsgRewardsBid) (*MsgRewardsBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RewardsBid not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_GovSetRewardsParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovSetRewardsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovSetRewardsParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.auction.v1.Msg/GovSetRewardsParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovSetRewardsParams(ctx, req.(*MsgGovSetRewardsParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RewardsBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRewardsBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RewardsBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.auction.v1.Msg/RewardsBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RewardsBid(ctx, req.(*MsgRewardsBid))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "umee.auction.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GovSetRewardsParams",
			Handler:    _Msg_GovSetRewardsParams_Handler,
		},
		{
			MethodName: "RewardsBid",
			Handler:    _Msg_RewardsBid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umee/auction/v1/tx.proto",
}

func (m *MsgGovSetRewardsParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovSetRewardsParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovSetRewardsParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *MsgGovSetRewardsParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovSetRewardsParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovSetRewardsParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRewardsBid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRewardsBid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRewardsBid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRewardsBidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRewardsBidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRewardsBidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgGovSetRewardsParams) Size() (n int) {
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

func (m *MsgGovSetRewardsParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRewardsBid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgRewardsBidResponse) Size() (n int) {
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
func (m *MsgGovSetRewardsParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGovSetRewardsParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovSetRewardsParams: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgGovSetRewardsParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGovSetRewardsParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovSetRewardsParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgRewardsBid) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRewardsBid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRewardsBid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgRewardsBidResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRewardsBidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRewardsBidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
