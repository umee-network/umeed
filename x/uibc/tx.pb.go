// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/uibc/v1/tx.proto

package uibc

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgGovUpdateQuota defines the Msg/GovUpdateQuota request type.
type MsgGovUpdateQuota struct {
	// authority is the address of the governance account.
	Authority   string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// total quota defines the total outflow of ibc-transfer in USD
	Total github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=total,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"total"`
	// per_denom quota for outflows per denom. All denoms have the same quota size.
	PerDenom github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=per_denom,json=perDenom,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"per_denom"`
	// quota_duration defines quota expires per denom, All denoms have the same expire time.
	QuotaDuration time.Duration `protobuf:"bytes,6,opt,name=quota_duration,json=quotaDuration,proto3,stdduration" json:"quota_duration,omitempty" yaml:"quota_duration"`
}

func (m *MsgGovUpdateQuota) Reset()      { *m = MsgGovUpdateQuota{} }
func (*MsgGovUpdateQuota) ProtoMessage() {}
func (*MsgGovUpdateQuota) Descriptor() ([]byte, []int) {
	return fileDescriptor_1982abc7d531f4dc, []int{0}
}
func (m *MsgGovUpdateQuota) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovUpdateQuota) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovUpdateQuota.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovUpdateQuota) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovUpdateQuota.Merge(m, src)
}
func (m *MsgGovUpdateQuota) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovUpdateQuota) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovUpdateQuota.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovUpdateQuota proto.InternalMessageInfo

func (*MsgGovUpdateQuota) XXX_MessageName() string {
	return "umee.uibc.v1.MsgGovUpdateQuota"
}

// MsgGovUpdateQuotaResponse defines response type for the Msg/GovUpdateQuota for with x/gov proposals.
type MsgGovUpdateQuotaResponse struct {
}

func (m *MsgGovUpdateQuotaResponse) Reset()         { *m = MsgGovUpdateQuotaResponse{} }
func (m *MsgGovUpdateQuotaResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGovUpdateQuotaResponse) ProtoMessage()    {}
func (*MsgGovUpdateQuotaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1982abc7d531f4dc, []int{1}
}
func (m *MsgGovUpdateQuotaResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovUpdateQuotaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovUpdateQuotaResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovUpdateQuotaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovUpdateQuotaResponse.Merge(m, src)
}
func (m *MsgGovUpdateQuotaResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovUpdateQuotaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovUpdateQuotaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovUpdateQuotaResponse proto.InternalMessageInfo

func (*MsgGovUpdateQuotaResponse) XXX_MessageName() string {
	return "umee.uibc.v1.MsgGovUpdateQuotaResponse"
}

// MsgGovSetIBCSQuotaStatus defines the request type for setting the IBC quota status.
type MsgGovSetIBCSQuotaStatus struct {
	// authority is the address of the governance account.
	Authority   string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// quota_status defines ibc transfer quota status
	QuotaStatus IBCTransferQuotaStatus `protobuf:"varint,4,opt,name=quota_status,json=quotaStatus,proto3,enum=umee.uibc.v1.IBCTransferQuotaStatus" json:"quota_status,omitempty"`
}

func (m *MsgGovSetIBCSQuotaStatus) Reset()      { *m = MsgGovSetIBCSQuotaStatus{} }
func (*MsgGovSetIBCSQuotaStatus) ProtoMessage() {}
func (*MsgGovSetIBCSQuotaStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_1982abc7d531f4dc, []int{2}
}
func (m *MsgGovSetIBCSQuotaStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovSetIBCSQuotaStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovSetIBCSQuotaStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovSetIBCSQuotaStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovSetIBCSQuotaStatus.Merge(m, src)
}
func (m *MsgGovSetIBCSQuotaStatus) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovSetIBCSQuotaStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovSetIBCSQuotaStatus.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovSetIBCSQuotaStatus proto.InternalMessageInfo

func (*MsgGovSetIBCSQuotaStatus) XXX_MessageName() string {
	return "umee.uibc.v1.MsgGovSetIBCSQuotaStatus"
}

// MsgGovSetIBCQuotaStatusResponse define the response type for Msg/MsgGovSetIBCSQuotaStatus with x/gov proposals.
type MsgGovSetIBCQuotaStatusResponse struct {
}

func (m *MsgGovSetIBCQuotaStatusResponse) Reset()         { *m = MsgGovSetIBCQuotaStatusResponse{} }
func (m *MsgGovSetIBCQuotaStatusResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGovSetIBCQuotaStatusResponse) ProtoMessage()    {}
func (*MsgGovSetIBCQuotaStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1982abc7d531f4dc, []int{3}
}
func (m *MsgGovSetIBCQuotaStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovSetIBCQuotaStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovSetIBCQuotaStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovSetIBCQuotaStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovSetIBCQuotaStatusResponse.Merge(m, src)
}
func (m *MsgGovSetIBCQuotaStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovSetIBCQuotaStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovSetIBCQuotaStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovSetIBCQuotaStatusResponse proto.InternalMessageInfo

func (*MsgGovSetIBCQuotaStatusResponse) XXX_MessageName() string {
	return "umee.uibc.v1.MsgGovSetIBCQuotaStatusResponse"
}
func init() {
	proto.RegisterType((*MsgGovUpdateQuota)(nil), "umee.uibc.v1.MsgGovUpdateQuota")
	proto.RegisterType((*MsgGovUpdateQuotaResponse)(nil), "umee.uibc.v1.MsgGovUpdateQuotaResponse")
	proto.RegisterType((*MsgGovSetIBCSQuotaStatus)(nil), "umee.uibc.v1.MsgGovSetIBCSQuotaStatus")
	proto.RegisterType((*MsgGovSetIBCQuotaStatusResponse)(nil), "umee.uibc.v1.MsgGovSetIBCQuotaStatusResponse")
}

func init() { proto.RegisterFile("umee/uibc/v1/tx.proto", fileDescriptor_1982abc7d531f4dc) }

var fileDescriptor_1982abc7d531f4dc = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xb6, 0x7f, 0xf9, 0xa5, 0x22, 0x97, 0x12, 0x09, 0x2b, 0x05, 0x27, 0x48, 0x76, 0x88, 0x50,
	0xa9, 0x10, 0xb1, 0xd5, 0x80, 0x18, 0x2a, 0x18, 0x48, 0x23, 0x95, 0x0e, 0x1d, 0x70, 0x60, 0xa0,
	0x4b, 0xe4, 0xd8, 0x57, 0xd7, 0x6a, 0xec, 0x73, 0xef, 0x5e, 0x87, 0x66, 0x42, 0x62, 0x62, 0x64,
	0xec, 0xd8, 0x8f, 0xc0, 0xd0, 0x0f, 0x11, 0x89, 0x25, 0xea, 0x84, 0x18, 0x02, 0x24, 0x03, 0x12,
	0x6c, 0x7c, 0x02, 0xe4, 0xb3, 0xa3, 0x38, 0x24, 0x15, 0x0b, 0x62, 0xb2, 0xdf, 0xf7, 0x79, 0xee,
	0x79, 0xff, 0x3c, 0x67, 0xa3, 0xb5, 0xd0, 0xc3, 0x58, 0x0f, 0xdd, 0x8e, 0xa5, 0xf7, 0x36, 0x75,
	0x38, 0xd1, 0x02, 0x4a, 0x80, 0x48, 0xab, 0x51, 0x5a, 0x8b, 0xd2, 0x5a, 0x6f, 0xb3, 0x5c, 0x74,
	0x88, 0x43, 0x38, 0xa0, 0x47, 0x6f, 0x31, 0xa7, 0xac, 0x38, 0x84, 0x38, 0x5d, 0xac, 0xf3, 0xa8,
	0x13, 0x1e, 0xe8, 0x76, 0x48, 0x4d, 0x70, 0x89, 0x9f, 0xe0, 0x37, 0x2c, 0xc2, 0x3c, 0xc2, 0x74,
	0x8f, 0x39, 0x91, 0xb6, 0xc7, 0x9c, 0x04, 0x28, 0xc5, 0x40, 0x3b, 0x56, 0x8c, 0x83, 0x04, 0x92,
	0xe7, 0xda, 0x39, 0x0e, 0x09, 0x98, 0x31, 0x52, 0xfd, 0x90, 0x41, 0xd7, 0xf6, 0x98, 0xb3, 0x43,
	0x7a, 0x2f, 0x02, 0xdb, 0x04, 0xfc, 0x2c, 0xc2, 0xa4, 0x87, 0x28, 0x67, 0x86, 0x70, 0x48, 0xa8,
	0x0b, 0x7d, 0x59, 0xac, 0x88, 0x1b, 0xb9, 0x86, 0x7c, 0x71, 0x5e, 0x2b, 0x26, 0xa2, 0x4f, 0x6c,
	0x9b, 0x62, 0xc6, 0x5a, 0x40, 0x5d, 0xdf, 0x31, 0x66, 0x54, 0xa9, 0x88, 0xb2, 0xe0, 0x42, 0x17,
	0xcb, 0xff, 0x45, 0x67, 0x8c, 0x38, 0x90, 0x2a, 0x28, 0x6f, 0x63, 0x66, 0x51, 0x37, 0x88, 0xc6,
	0x90, 0x33, 0x1c, 0x4b, 0xa7, 0x24, 0x03, 0x65, 0x81, 0x80, 0xd9, 0x95, 0xff, 0xe7, 0xb5, 0x1e,
	0x0d, 0x46, 0xaa, 0xf0, 0x69, 0xa4, 0xae, 0x3b, 0x2e, 0x1c, 0x86, 0x1d, 0xcd, 0x22, 0x5e, 0x32,
	0x4f, 0xf2, 0xa8, 0x31, 0xfb, 0x48, 0x87, 0x7e, 0x80, 0x99, 0xd6, 0xc4, 0xd6, 0xc5, 0x79, 0x0d,
	0x25, 0x9d, 0x35, 0xb1, 0x65, 0xc4, 0x52, 0xd2, 0x4b, 0x94, 0x0b, 0x30, 0x6d, 0xdb, 0xd8, 0x27,
	0x9e, 0x9c, 0xfd, 0x0b, 0xba, 0x57, 0x02, 0x4c, 0x9b, 0x91, 0x9a, 0xf4, 0x1a, 0x15, 0xf8, 0x0e,
	0xdb, 0x53, 0x6b, 0xe4, 0x95, 0x8a, 0xb8, 0x91, 0xaf, 0x97, 0xb4, 0xd8, 0x3b, 0x6d, 0xea, 0x9d,
	0xd6, 0x4c, 0x08, 0x8d, 0xc7, 0x51, 0xe9, 0xef, 0x23, 0x55, 0x9e, 0x3f, 0x78, 0x8f, 0x78, 0x2e,
	0x60, 0x2f, 0x80, 0xfe, 0xcf, 0x91, 0xba, 0xd6, 0x37, 0xbd, 0xee, 0x56, 0x75, 0x9e, 0x51, 0x3d,
	0xfd, 0xac, 0x8a, 0xc6, 0x55, 0x9e, 0x9c, 0xaa, 0x6d, 0x5d, 0x7f, 0x7b, 0xa6, 0x0a, 0xa7, 0x67,
	0xaa, 0xf0, 0xe6, 0xdb, 0xfb, 0xbb, 0xb3, 0xfd, 0x57, 0x6f, 0xa2, 0xd2, 0x82, 0x99, 0x06, 0x66,
	0x01, 0xf1, 0x19, 0xae, 0xfe, 0x10, 0x91, 0x1c, 0xa3, 0x2d, 0x0c, 0xbb, 0x8d, 0xed, 0x16, 0x87,
	0x5b, 0x60, 0x42, 0xc8, 0xfe, 0xb9, 0xe3, 0x3b, 0x68, 0x35, 0x9e, 0x93, 0xf1, 0xfa, 0xdc, 0xf8,
	0x42, 0xfd, 0xb6, 0x96, 0xfe, 0x40, 0xb4, 0xdd, 0xc6, 0xf6, 0x73, 0x6a, 0xfa, 0xec, 0x00, 0xd3,
	0x54, 0xaf, 0x46, 0xfe, 0x78, 0x16, 0x5c, 0xba, 0x8a, 0x5b, 0x48, 0x4d, 0x0f, 0x9b, 0x3e, 0x9f,
	0x2c, 0xa4, 0x3e, 0x14, 0x51, 0x66, 0x8f, 0x39, 0xd2, 0x3e, 0x2a, 0xfc, 0x76, 0xff, 0xd5, 0xf9,
	0x3e, 0x16, 0x76, 0x5a, 0xbe, 0xf3, 0x07, 0xc2, 0xb4, 0x86, 0xe4, 0xa1, 0xe2, 0xb2, 0x1e, 0xa4,
	0xf5, 0x65, 0x02, 0x8b, 0xbe, 0x94, 0x6b, 0x97, 0xf3, 0x96, 0x8c, 0xd4, 0x78, 0x3a, 0xf8, 0xaa,
	0x08, 0x83, 0xb1, 0x22, 0x0e, 0xc7, 0x8a, 0xf8, 0x65, 0xac, 0x88, 0xef, 0x26, 0x8a, 0x30, 0x98,
	0x28, 0xe2, 0x70, 0xa2, 0x08, 0x1f, 0x27, 0x8a, 0xb0, 0x9f, 0xbe, 0xfb, 0x91, 0x74, 0xcd, 0xc7,
	0xf0, 0x8a, 0xd0, 0x23, 0x1e, 0xe8, 0xbd, 0x07, 0xfa, 0x09, 0xff, 0x4f, 0x74, 0x56, 0xf8, 0x1d,
	0xbe, 0xff, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xae, 0x17, 0x81, 0x54, 0xca, 0x04, 0x00, 0x00,
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
	// GovUpdateQuota adds new quota for ibc denoms or
	// updates the quota for existed ibc denoms.
	GovUpdateQuota(ctx context.Context, in *MsgGovUpdateQuota, opts ...grpc.CallOption) (*MsgGovUpdateQuotaResponse, error)
	// GovSetIBCQuotaStatus updates the status of the IBC transfer quota check for both inflow and outflow.
	GovSetIBCQuotaStatus(ctx context.Context, in *MsgGovSetIBCSQuotaStatus, opts ...grpc.CallOption) (*MsgGovSetIBCQuotaStatusResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) GovUpdateQuota(ctx context.Context, in *MsgGovUpdateQuota, opts ...grpc.CallOption) (*MsgGovUpdateQuotaResponse, error) {
	out := new(MsgGovUpdateQuotaResponse)
	err := c.cc.Invoke(ctx, "/umee.uibc.v1.Msg/GovUpdateQuota", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GovSetIBCQuotaStatus(ctx context.Context, in *MsgGovSetIBCSQuotaStatus, opts ...grpc.CallOption) (*MsgGovSetIBCQuotaStatusResponse, error) {
	out := new(MsgGovSetIBCQuotaStatusResponse)
	err := c.cc.Invoke(ctx, "/umee.uibc.v1.Msg/GovSetIBCQuotaStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// GovUpdateQuota adds new quota for ibc denoms or
	// updates the quota for existed ibc denoms.
	GovUpdateQuota(context.Context, *MsgGovUpdateQuota) (*MsgGovUpdateQuotaResponse, error)
	// GovSetIBCQuotaStatus updates the status of the IBC transfer quota check for both inflow and outflow.
	GovSetIBCQuotaStatus(context.Context, *MsgGovSetIBCSQuotaStatus) (*MsgGovSetIBCQuotaStatusResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) GovUpdateQuota(ctx context.Context, req *MsgGovUpdateQuota) (*MsgGovUpdateQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovUpdateQuota not implemented")
}
func (*UnimplementedMsgServer) GovSetIBCQuotaStatus(ctx context.Context, req *MsgGovSetIBCSQuotaStatus) (*MsgGovSetIBCQuotaStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovSetIBCQuotaStatus not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_GovUpdateQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovUpdateQuota)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovUpdateQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.uibc.v1.Msg/GovUpdateQuota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovUpdateQuota(ctx, req.(*MsgGovUpdateQuota))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GovSetIBCQuotaStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovSetIBCSQuotaStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovSetIBCQuotaStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.uibc.v1.Msg/GovSetIBCQuotaStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovSetIBCQuotaStatus(ctx, req.(*MsgGovSetIBCSQuotaStatus))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "umee.uibc.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GovUpdateQuota",
			Handler:    _Msg_GovUpdateQuota_Handler,
		},
		{
			MethodName: "GovSetIBCQuotaStatus",
			Handler:    _Msg_GovSetIBCQuotaStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umee/uibc/v1/tx.proto",
}

func (m *MsgGovUpdateQuota) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovUpdateQuota) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovUpdateQuota) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.QuotaDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.QuotaDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTx(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	{
		size := m.PerDenom.Size()
		i -= size
		if _, err := m.PerDenom.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Total.Size()
		i -= size
		if _, err := m.Total.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
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

func (m *MsgGovUpdateQuotaResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovUpdateQuotaResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovUpdateQuotaResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgGovSetIBCSQuotaStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovSetIBCSQuotaStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovSetIBCSQuotaStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.QuotaStatus != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.QuotaStatus))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
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

func (m *MsgGovSetIBCQuotaStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovSetIBCQuotaStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovSetIBCQuotaStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgGovUpdateQuota) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Total.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.PerDenom.Size()
	n += 1 + l + sovTx(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.QuotaDuration)
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgGovUpdateQuotaResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgGovSetIBCSQuotaStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.QuotaStatus != 0 {
		n += 1 + sovTx(uint64(m.QuotaStatus))
	}
	return n
}

func (m *MsgGovSetIBCQuotaStatusResponse) Size() (n int) {
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
func (m *MsgGovUpdateQuota) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGovUpdateQuota: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovUpdateQuota: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
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
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
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
			if err := m.Total.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerDenom", wireType)
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
			if err := m.PerDenom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuotaDuration", wireType)
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.QuotaDuration, dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgGovUpdateQuotaResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGovUpdateQuotaResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovUpdateQuotaResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgGovSetIBCSQuotaStatus) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGovSetIBCSQuotaStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovSetIBCSQuotaStatus: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
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
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuotaStatus", wireType)
			}
			m.QuotaStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QuotaStatus |= IBCTransferQuotaStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgGovSetIBCQuotaStatusResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGovSetIBCQuotaStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovSetIBCQuotaStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
