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

// MsgGovSetIBCStatus defines the request type for setting the IBC quota status.
type MsgGovSetIBCStatus struct {
	// authority is the address of the governance account.
	Authority   string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// ibc_status defines ibc transfer quota status
	IbcStatus IBCTransferStatus `protobuf:"varint,4,opt,name=ibc_status,json=ibcStatus,proto3,enum=umee.uibc.v1.IBCTransferStatus" json:"ibc_status,omitempty"`
}

func (m *MsgGovSetIBCStatus) Reset()      { *m = MsgGovSetIBCStatus{} }
func (*MsgGovSetIBCStatus) ProtoMessage() {}
func (*MsgGovSetIBCStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_1982abc7d531f4dc, []int{2}
}
func (m *MsgGovSetIBCStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovSetIBCStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovSetIBCStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovSetIBCStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovSetIBCStatus.Merge(m, src)
}
func (m *MsgGovSetIBCStatus) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovSetIBCStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovSetIBCStatus.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovSetIBCStatus proto.InternalMessageInfo

func (*MsgGovSetIBCStatus) XXX_MessageName() string {
	return "umee.uibc.v1.MsgGovSetIBCStatus"
}

// MsgGovSetIBCStatusResponse define the response type for Msg/MsgGovSetIBCStatus with x/gov proposals.
type MsgGovSetIBCStatusResponse struct {
}

func (m *MsgGovSetIBCStatusResponse) Reset()         { *m = MsgGovSetIBCStatusResponse{} }
func (m *MsgGovSetIBCStatusResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGovSetIBCStatusResponse) ProtoMessage()    {}
func (*MsgGovSetIBCStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1982abc7d531f4dc, []int{3}
}
func (m *MsgGovSetIBCStatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovSetIBCStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovSetIBCStatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovSetIBCStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovSetIBCStatusResponse.Merge(m, src)
}
func (m *MsgGovSetIBCStatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovSetIBCStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovSetIBCStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovSetIBCStatusResponse proto.InternalMessageInfo

func (*MsgGovSetIBCStatusResponse) XXX_MessageName() string {
	return "umee.uibc.v1.MsgGovSetIBCStatusResponse"
}
func init() {
	proto.RegisterType((*MsgGovUpdateQuota)(nil), "umee.uibc.v1.MsgGovUpdateQuota")
	proto.RegisterType((*MsgGovUpdateQuotaResponse)(nil), "umee.uibc.v1.MsgGovUpdateQuotaResponse")
	proto.RegisterType((*MsgGovSetIBCStatus)(nil), "umee.uibc.v1.MsgGovSetIBCStatus")
	proto.RegisterType((*MsgGovSetIBCStatusResponse)(nil), "umee.uibc.v1.MsgGovSetIBCStatusResponse")
}

func init() { proto.RegisterFile("umee/uibc/v1/tx.proto", fileDescriptor_1982abc7d531f4dc) }

var fileDescriptor_1982abc7d531f4dc = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xb6, 0x09, 0xa9, 0xc8, 0x15, 0x82, 0xb0, 0x52, 0x70, 0x02, 0xb2, 0xa3, 0x0c, 0x10, 0x21,
	0x62, 0xab, 0x01, 0x31, 0x54, 0x80, 0x44, 0x1a, 0x09, 0x3a, 0x74, 0xc0, 0x81, 0x81, 0x4a, 0x28,
	0xf2, 0x8f, 0xab, 0x6b, 0x35, 0xf6, 0x99, 0xbb, 0x73, 0x68, 0x26, 0x24, 0x26, 0x46, 0xc6, 0x8e,
	0xfd, 0x13, 0x18, 0xfa, 0x07, 0x30, 0x46, 0x62, 0xa9, 0x3a, 0x21, 0x84, 0x02, 0x24, 0x03, 0x12,
	0x23, 0x7f, 0x01, 0xba, 0x3b, 0x5b, 0xcd, 0x0f, 0x50, 0x17, 0xc4, 0x94, 0x7b, 0xef, 0xfb, 0xee,
	0x7b, 0xf7, 0xde, 0xf7, 0x62, 0xb0, 0x92, 0x84, 0x10, 0x9a, 0x49, 0xe0, 0xb8, 0x66, 0x7f, 0xd5,
	0xa4, 0x7b, 0x46, 0x8c, 0x11, 0x45, 0xca, 0x79, 0x96, 0x36, 0x58, 0xda, 0xe8, 0xaf, 0x56, 0x4a,
	0x3e, 0xf2, 0x11, 0x07, 0x4c, 0x76, 0x12, 0x9c, 0x8a, 0xe6, 0x23, 0xe4, 0xf7, 0xa0, 0xc9, 0x23,
	0x27, 0xd9, 0x36, 0xbd, 0x04, 0xdb, 0x34, 0x40, 0x51, 0x8a, 0x5f, 0x71, 0x11, 0x09, 0x11, 0x31,
	0x43, 0xe2, 0x33, 0xed, 0x90, 0xf8, 0x29, 0x50, 0x16, 0x40, 0x57, 0x28, 0x8a, 0x20, 0x85, 0xd4,
	0x99, 0xe7, 0xbc, 0x4c, 0x10, 0xb5, 0x05, 0x52, 0xfb, 0x98, 0x03, 0x97, 0x36, 0x89, 0xff, 0x08,
	0xf5, 0x9f, 0xc5, 0x9e, 0x4d, 0xe1, 0x13, 0x86, 0x29, 0x77, 0x41, 0xc1, 0x4e, 0xe8, 0x0e, 0xc2,
	0x01, 0x1d, 0xa8, 0x72, 0x55, 0xae, 0x17, 0x5a, 0xea, 0xf1, 0x61, 0xa3, 0x94, 0x8a, 0x3e, 0xf4,
	0x3c, 0x0c, 0x09, 0xe9, 0x50, 0x1c, 0x44, 0xbe, 0x75, 0x42, 0x55, 0x4a, 0x20, 0x4f, 0x03, 0xda,
	0x83, 0xea, 0x19, 0x76, 0xc7, 0x12, 0x81, 0x52, 0x05, 0xcb, 0x1e, 0x24, 0x2e, 0x0e, 0x62, 0xd6,
	0x86, 0x9a, 0xe3, 0xd8, 0x74, 0x4a, 0xb1, 0x40, 0x9e, 0x22, 0x6a, 0xf7, 0xd4, 0xb3, 0xbc, 0xd6,
	0xbd, 0xe1, 0x48, 0x97, 0x3e, 0x8f, 0xf4, 0xeb, 0x7e, 0x40, 0x77, 0x12, 0xc7, 0x70, 0x51, 0x98,
	0xf6, 0x93, 0xfe, 0x34, 0x88, 0xb7, 0x6b, 0xd2, 0x41, 0x0c, 0x89, 0xd1, 0x86, 0xee, 0xf1, 0x61,
	0x03, 0xa4, 0x2f, 0x6b, 0x43, 0xd7, 0x12, 0x52, 0xca, 0x73, 0x50, 0x88, 0x21, 0xee, 0x7a, 0x30,
	0x42, 0xa1, 0x9a, 0xff, 0x07, 0xba, 0xe7, 0x62, 0x88, 0xdb, 0x4c, 0x4d, 0x79, 0x0d, 0x8a, 0x7c,
	0x86, 0xdd, 0xcc, 0x1a, 0x75, 0xa9, 0x2a, 0xd7, 0x97, 0x9b, 0x65, 0x43, 0x78, 0x67, 0x64, 0xde,
	0x19, 0xed, 0x94, 0xd0, 0xba, 0xcf, 0x4a, 0xff, 0x1c, 0xe9, 0xea, 0xec, 0xc5, 0x5b, 0x28, 0x0c,
	0x28, 0x0c, 0x63, 0x3a, 0xf8, 0x35, 0xd2, 0x57, 0x06, 0x76, 0xd8, 0x5b, 0xab, 0xcd, 0x32, 0x6a,
	0xfb, 0x5f, 0x75, 0xd9, 0xba, 0xc0, 0x93, 0x99, 0xda, 0xda, 0xe5, 0xb7, 0x07, 0xba, 0xb4, 0x7f,
	0xa0, 0x4b, 0x6f, 0x7e, 0xbc, 0xbf, 0x79, 0x32, 0xff, 0xda, 0x55, 0x50, 0x5e, 0x30, 0xd3, 0x82,
	0x24, 0x46, 0x11, 0x81, 0xb5, 0x2f, 0x32, 0x50, 0x04, 0xda, 0x81, 0x74, 0xa3, 0xb5, 0xde, 0xa1,
	0x36, 0x4d, 0xc8, 0x7f, 0xf7, 0xfa, 0x01, 0x00, 0x81, 0xe3, 0x76, 0x09, 0xaf, 0xce, 0x0d, 0x2f,
	0x36, 0x75, 0x63, 0xfa, 0x8f, 0x61, 0x6c, 0xb4, 0xd6, 0x9f, 0x62, 0x3b, 0x22, 0xdb, 0x10, 0x8b,
	0x47, 0x5a, 0x85, 0xc0, 0x71, 0xc5, 0xf1, 0xaf, 0xbd, 0x5f, 0x03, 0x95, 0xc5, 0xee, 0xb2, 0xe6,
	0x9b, 0x1f, 0x64, 0x90, 0xdb, 0x24, 0xbe, 0xb2, 0x05, 0x8a, 0x73, 0xbb, 0x3e, 0x57, 0x7b, 0x61,
	0x7e, 0x95, 0x1b, 0xa7, 0x10, 0xb2, 0x1a, 0xca, 0x0b, 0x70, 0x71, 0x7e, 0xb8, 0xd5, 0x3f, 0xdd,
	0x9d, 0x66, 0x54, 0xea, 0xa7, 0x31, 0x32, 0xf9, 0xd6, 0xe3, 0xe1, 0x77, 0x4d, 0x1a, 0x8e, 0x35,
	0xf9, 0x68, 0xac, 0xc9, 0xdf, 0xc6, 0x9a, 0xfc, 0x6e, 0xa2, 0x49, 0xc3, 0x89, 0x26, 0x1f, 0x4d,
	0x34, 0xe9, 0xd3, 0x44, 0x93, 0xb6, 0xa6, 0xf7, 0x9a, 0xa9, 0x36, 0x22, 0x48, 0x5f, 0x21, 0xbc,
	0xcb, 0x03, 0xb3, 0x7f, 0xc7, 0xdc, 0xe3, 0xdf, 0x00, 0x67, 0x89, 0xef, 0xe7, 0xed, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xa9, 0xc6, 0xdd, 0x61, 0xa6, 0x04, 0x00, 0x00,
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
	// GovSetIBCStatus sets IBC ICS20 status. Must be called by x/gov.
	GovSetIBCStatus(ctx context.Context, in *MsgGovSetIBCStatus, opts ...grpc.CallOption) (*MsgGovSetIBCStatusResponse, error)
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

func (c *msgClient) GovSetIBCStatus(ctx context.Context, in *MsgGovSetIBCStatus, opts ...grpc.CallOption) (*MsgGovSetIBCStatusResponse, error) {
	out := new(MsgGovSetIBCStatusResponse)
	err := c.cc.Invoke(ctx, "/umee.uibc.v1.Msg/GovSetIBCStatus", in, out, opts...)
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
	// GovSetIBCStatus sets IBC ICS20 status. Must be called by x/gov.
	GovSetIBCStatus(context.Context, *MsgGovSetIBCStatus) (*MsgGovSetIBCStatusResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) GovUpdateQuota(ctx context.Context, req *MsgGovUpdateQuota) (*MsgGovUpdateQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovUpdateQuota not implemented")
}
func (*UnimplementedMsgServer) GovSetIBCStatus(ctx context.Context, req *MsgGovSetIBCStatus) (*MsgGovSetIBCStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovSetIBCStatus not implemented")
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

func _Msg_GovSetIBCStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovSetIBCStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovSetIBCStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.uibc.v1.Msg/GovSetIBCStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovSetIBCStatus(ctx, req.(*MsgGovSetIBCStatus))
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
			MethodName: "GovSetIBCStatus",
			Handler:    _Msg_GovSetIBCStatus_Handler,
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

func (m *MsgGovSetIBCStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovSetIBCStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovSetIBCStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IbcStatus != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.IbcStatus))
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

func (m *MsgGovSetIBCStatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovSetIBCStatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovSetIBCStatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *MsgGovSetIBCStatus) Size() (n int) {
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
	if m.IbcStatus != 0 {
		n += 1 + sovTx(uint64(m.IbcStatus))
	}
	return n
}

func (m *MsgGovSetIBCStatusResponse) Size() (n int) {
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
func (m *MsgGovSetIBCStatus) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGovSetIBCStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovSetIBCStatus: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field IbcStatus", wireType)
			}
			m.IbcStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IbcStatus |= IBCTransferStatus(b&0x7F) << shift
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
func (m *MsgGovSetIBCStatusResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGovSetIBCStatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovSetIBCStatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
