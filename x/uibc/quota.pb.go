// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/uibc/v1/quota.proto

package uibc

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

// IBCTransferStatus status of ibc-transfer
type IBCTransferStatus int32

const (
	// UNSPECIFIED  defines a no-op status.
	IBCTransferStatus_IBC_TRANSFER_STATUS_UNSPECIFIED IBCTransferStatus = 0
	// DISABLED defines the quota checking diabled for ibc-transfer.
	IBCTransferStatus_IBC_TRANSFER_STATUS_DISABLED IBCTransferStatus = 1
	// ENABLED defines the enable quota checking for ibc-transfer.
	IBCTransferStatus_IBC_TRANSFER_STATUS_ENABLED IBCTransferStatus = 2
	// PAUSED defines pause the ibc-transfer from app.
	IBCTransferStatus_IBC_TRANSFER_STATUS_PAUSED IBCTransferStatus = 3
)

var IBCTransferStatus_name = map[int32]string{
	0: "IBC_TRANSFER_STATUS_UNSPECIFIED",
	1: "IBC_TRANSFER_STATUS_DISABLED",
	2: "IBC_TRANSFER_STATUS_ENABLED",
	3: "IBC_TRANSFER_STATUS_PAUSED",
}

var IBCTransferStatus_value = map[string]int32{
	"IBC_TRANSFER_STATUS_UNSPECIFIED": 0,
	"IBC_TRANSFER_STATUS_DISABLED":    1,
	"IBC_TRANSFER_STATUS_ENABLED":     2,
	"IBC_TRANSFER_STATUS_PAUSED":      3,
}

func (x IBCTransferStatus) String() string {
	return proto.EnumName(IBCTransferStatus_name, int32(x))
}

func (IBCTransferStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_651be1a0280abcb6, []int{0}
}

// Quota stores current sum of IBC outflow transfers of IBC Denom.
type Quota struct {
	// ibc_denom defines the ibc denom
	IbcDenom string `protobuf:"bytes,1,opt,name=ibc_denom,json=ibcDenom,proto3" json:"ibc_denom,omitempty"`
	// outflow_sum defines the sum of price (USD) value of outflow tokens through ibc-transfer
	OutflowSum github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=outflow_sum,json=outflowSum,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"outflow_sum"`
}

func (m *Quota) Reset()         { *m = Quota{} }
func (m *Quota) String() string { return proto.CompactTextString(m) }
func (*Quota) ProtoMessage()    {}
func (*Quota) Descriptor() ([]byte, []int) {
	return fileDescriptor_651be1a0280abcb6, []int{0}
}
func (m *Quota) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Quota) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Quota.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Quota) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quota.Merge(m, src)
}
func (m *Quota) XXX_Size() int {
	return m.Size()
}
func (m *Quota) XXX_DiscardUnknown() {
	xxx_messageInfo_Quota.DiscardUnknown(m)
}

var xxx_messageInfo_Quota proto.InternalMessageInfo

func (m *Quota) GetIbcDenom() string {
	if m != nil {
		return m.IbcDenom
	}
	return ""
}

// Params of x/uibc module
type Params struct {
	// ibc_status defines the wethever ibc-transfer enabled, disbaled or paused
	IbcPause IBCTransferStatus `protobuf:"varint,1,opt,name=ibc_pause,json=ibcPause,proto3,enum=umee.uibc.v1.IBCTransferStatus" json:"ibc_pause,omitempty"`
	// total_quota defines the total outflow limit of ibc-transfer in USD
	TotalQuota github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=total_quota,json=totalQuota,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"total_quota"`
	// token_quota defines the outflow limit per token in USD
	TokenQuota github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=token_quota,json=tokenQuota,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"token_quota"`
	// quota_duration defines quota expires for each ibc-transfer denom in seconds
	QuotaDuration time.Duration `protobuf:"bytes,4,opt,name=quota_duration,json=quotaDuration,proto3,stdduration" json:"quota_duration,omitempty" yaml:"quota_duration"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_651be1a0280abcb6, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetIbcPause() IBCTransferStatus {
	if m != nil {
		return m.IbcPause
	}
	return IBCTransferStatus_IBC_TRANSFER_STATUS_UNSPECIFIED
}

func (m *Params) GetQuotaDuration() time.Duration {
	if m != nil {
		return m.QuotaDuration
	}
	return 0
}

func init() {
	proto.RegisterEnum("umee.uibc.v1.IBCTransferStatus", IBCTransferStatus_name, IBCTransferStatus_value)
	proto.RegisterType((*Quota)(nil), "umee.uibc.v1.Quota")
	proto.RegisterType((*Params)(nil), "umee.uibc.v1.Params")
}

func init() { proto.RegisterFile("umee/uibc/v1/quota.proto", fileDescriptor_651be1a0280abcb6) }

var fileDescriptor_651be1a0280abcb6 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x8b, 0xd3, 0x4e,
	0x18, 0xed, 0x6c, 0x7f, 0xbf, 0xc5, 0x9d, 0xd5, 0xa5, 0x06, 0x85, 0x6c, 0x57, 0x92, 0x52, 0x41,
	0x16, 0xb1, 0x13, 0x76, 0xf5, 0x24, 0x2b, 0xd8, 0x34, 0x59, 0x08, 0x48, 0xa9, 0x49, 0x7b, 0x11,
	0x24, 0x24, 0xe9, 0x34, 0x86, 0x76, 0x32, 0x35, 0x99, 0xe9, 0xda, 0x93, 0x07, 0xf1, 0xee, 0x51,
	0xf0, 0xdf, 0xf0, 0x8f, 0xd8, 0xe3, 0xe2, 0x49, 0x3c, 0x54, 0x69, 0x6f, 0x1e, 0xfd, 0x0b, 0x64,
	0x26, 0x29, 0xec, 0x62, 0x8f, 0x7b, 0x4a, 0xbe, 0x79, 0xef, 0x7b, 0xdf, 0xe3, 0x7d, 0x33, 0x50,
	0xe5, 0x04, 0x63, 0x83, 0x27, 0x61, 0x64, 0xcc, 0x8e, 0x8c, 0xb7, 0x9c, 0xb2, 0x00, 0x4d, 0x33,
	0xca, 0xa8, 0x72, 0x53, 0x20, 0x48, 0x20, 0x68, 0x76, 0x54, 0xbf, 0x13, 0xd3, 0x98, 0x4a, 0xc0,
	0x10, 0x7f, 0x05, 0xa7, 0xae, 0xc5, 0x94, 0xc6, 0x13, 0x6c, 0xc8, 0x2a, 0xe4, 0x23, 0x63, 0xc8,
	0xb3, 0x80, 0x25, 0x34, 0x2d, 0xf1, 0xfd, 0x88, 0xe6, 0x84, 0xe6, 0x7e, 0xd1, 0x58, 0x14, 0x05,
	0xd4, 0xfc, 0x00, 0xe0, 0xff, 0x2f, 0xc5, 0x38, 0xe5, 0x00, 0xee, 0x24, 0x61, 0xe4, 0x0f, 0x71,
	0x4a, 0x89, 0x0a, 0x1a, 0xe0, 0x70, 0xc7, 0xbd, 0x91, 0x84, 0x91, 0x25, 0x6a, 0xe5, 0x35, 0xdc,
	0xa5, 0x9c, 0x8d, 0x26, 0xf4, 0xcc, 0xcf, 0x39, 0x51, 0xab, 0x02, 0x36, 0x4f, 0xce, 0x17, 0x7a,
	0xe5, 0xc7, 0x42, 0x7f, 0x10, 0x27, 0xec, 0x0d, 0x0f, 0x51, 0x44, 0x49, 0x29, 0x5e, 0x7e, 0x5a,
	0xf9, 0x70, 0x6c, 0xb0, 0xf9, 0x14, 0xe7, 0xc8, 0xc2, 0xd1, 0xb7, 0xaf, 0x2d, 0x58, 0xce, 0xb6,
	0x70, 0xe4, 0xc2, 0x52, 0xd0, 0xe3, 0xa4, 0xf9, 0xb1, 0x0a, 0xb7, 0x7b, 0x41, 0x16, 0x90, 0x5c,
	0x39, 0x29, 0x6c, 0x4c, 0x03, 0x9e, 0x63, 0x69, 0x63, 0xef, 0x58, 0x47, 0x97, 0x33, 0x40, 0x8e,
	0xd9, 0xe9, 0x67, 0x41, 0x9a, 0x8f, 0x70, 0xe6, 0xb1, 0x80, 0xf1, 0x5c, 0xfa, 0xec, 0x89, 0x06,
	0xe1, 0x93, 0x51, 0x16, 0x4c, 0x7c, 0x19, 0xa1, 0xba, 0x75, 0x1d, 0x3e, 0xa5, 0x60, 0x91, 0x91,
	0x94, 0x1f, 0xe3, 0xb4, 0x94, 0xaf, 0x5e, 0x8f, 0xfc, 0x18, 0xa7, 0x85, 0xfc, 0x7b, 0xb8, 0x27,
	0x85, 0xfd, 0xf5, 0xfe, 0xd4, 0xff, 0x1a, 0xe0, 0x70, 0xf7, 0x78, 0x1f, 0x15, 0x0b, 0x46, 0xeb,
	0x05, 0x23, 0xab, 0x24, 0x98, 0xcf, 0xc4, 0xf0, 0xdf, 0x0b, 0x5d, 0xbd, 0xda, 0xf8, 0x88, 0x92,
	0x84, 0x61, 0x32, 0x65, 0xf3, 0x3f, 0x0b, 0xfd, 0xee, 0x3c, 0x20, 0x93, 0xa7, 0xcd, 0xab, 0x8c,
	0xe6, 0xe7, 0x9f, 0x3a, 0x70, 0x6f, 0xc9, 0xc3, 0xb5, 0xda, 0xc3, 0x2f, 0x00, 0xde, 0xfe, 0x27,
	0x5e, 0xe5, 0x3e, 0xd4, 0x1d, 0xb3, 0xe3, 0xf7, 0xdd, 0x76, 0xd7, 0x3b, 0xb5, 0x5d, 0xdf, 0xeb,
	0xb7, 0xfb, 0x03, 0xcf, 0x1f, 0x74, 0xbd, 0x9e, 0xdd, 0x71, 0x4e, 0x1d, 0xdb, 0xaa, 0x55, 0x94,
	0x06, 0xbc, 0xb7, 0x89, 0x64, 0x39, 0x5e, 0xdb, 0x7c, 0x61, 0x5b, 0x35, 0xa0, 0xe8, 0xf0, 0x60,
	0x13, 0xc3, 0xee, 0x16, 0x84, 0x2d, 0x45, 0x83, 0xf5, 0x4d, 0x84, 0x5e, 0x7b, 0xe0, 0xd9, 0x56,
	0xad, 0x6a, 0x3e, 0x3f, 0x5f, 0x6a, 0xe0, 0x62, 0xa9, 0x81, 0x5f, 0x4b, 0x0d, 0x7c, 0x5a, 0x69,
	0x95, 0x8b, 0x95, 0x56, 0xf9, 0xbe, 0xd2, 0x2a, 0xaf, 0x2e, 0x47, 0x2f, 0xee, 0x4a, 0x2b, 0xc5,
	0xec, 0x8c, 0x66, 0x63, 0x59, 0x18, 0xb3, 0x27, 0xc6, 0x3b, 0xf9, 0xb6, 0xc2, 0x6d, 0x19, 0xe0,
	0xe3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x41, 0x84, 0x5e, 0x6f, 0x03, 0x00, 0x00,
}

func (m *Quota) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Quota) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Quota) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.OutflowSum.Size()
		i -= size
		if _, err := m.OutflowSum.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuota(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.IbcDenom) > 0 {
		i -= len(m.IbcDenom)
		copy(dAtA[i:], m.IbcDenom)
		i = encodeVarintQuota(dAtA, i, uint64(len(m.IbcDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.QuotaDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.QuotaDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintQuota(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size := m.TokenQuota.Size()
		i -= size
		if _, err := m.TokenQuota.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuota(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.TotalQuota.Size()
		i -= size
		if _, err := m.TotalQuota.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuota(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.IbcPause != 0 {
		i = encodeVarintQuota(dAtA, i, uint64(m.IbcPause))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuota(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuota(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Quota) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IbcDenom)
	if l > 0 {
		n += 1 + l + sovQuota(uint64(l))
	}
	l = m.OutflowSum.Size()
	n += 1 + l + sovQuota(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IbcPause != 0 {
		n += 1 + sovQuota(uint64(m.IbcPause))
	}
	l = m.TotalQuota.Size()
	n += 1 + l + sovQuota(uint64(l))
	l = m.TokenQuota.Size()
	n += 1 + l + sovQuota(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.QuotaDuration)
	n += 1 + l + sovQuota(uint64(l))
	return n
}

func sovQuota(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuota(x uint64) (n int) {
	return sovQuota(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Quota) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuota
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
			return fmt.Errorf("proto: Quota: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Quota: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuota
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutflowSum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuota
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OutflowSum.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuota(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuota
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuota
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcPause", wireType)
			}
			m.IbcPause = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IbcPause |= IBCTransferStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalQuota", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuota
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalQuota.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenQuota", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuota
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenQuota.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuotaDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
				return ErrInvalidLengthQuota
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuota
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
			skippy, err := skipQuota(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuota
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
func skipQuota(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuota
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
					return 0, ErrIntOverflowQuota
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
					return 0, ErrIntOverflowQuota
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
				return 0, ErrInvalidLengthQuota
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuota
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuota
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuota        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuota          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuota = fmt.Errorf("proto: unexpected end of group")
)
