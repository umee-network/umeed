// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/uibc/v1/quota.proto

package uibc

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

// IBCTransferStatus status of ibc-transfer quota check for inflow and outflow
type IBCTransferStatus int32

const (
	// UNSPECIFIED  defines a no-op status.
	IBCTransferStatus_IBC_TRANSFER_STATUS_UNSPECIFIED IBCTransferStatus = 0
	// DISABLED: all inflow and outflow quota checks are disabled.
	IBCTransferStatus_IBC_TRANSFER_STATUS_QUOTA_DISABLED IBCTransferStatus = 1
	// ENABLED: all inflow and outflow quota checks are enabled.
	IBCTransferStatus_IBC_TRANSFER_STATUS_QUOTA_ENABLED IBCTransferStatus = 2
	// DISABLED OUT: outflow quota check is disabled, while the inflow quota check is enabled.
	IBCTransferStatus_IBC_TRANSFER_STATUS_QUOTA_OUT_DISABLED IBCTransferStatus = 3
	// DISABLED IN: inflow quota check is disabled, while the outflow quota check is enabled.
	IBCTransferStatus_IBC_TRANSFER_STATUS_QUOTA_IN_DISABLED IBCTransferStatus = 4
	// PAUSED: all IBC transfers are paused.
	IBCTransferStatus_IBC_TRANSFER_STATUS_TRANSFERS_PAUSED IBCTransferStatus = 5
)

var IBCTransferStatus_name = map[int32]string{
	0: "IBC_TRANSFER_STATUS_UNSPECIFIED",
	1: "IBC_TRANSFER_STATUS_QUOTA_DISABLED",
	2: "IBC_TRANSFER_STATUS_QUOTA_ENABLED",
	3: "IBC_TRANSFER_STATUS_QUOTA_OUT_DISABLED",
	4: "IBC_TRANSFER_STATUS_QUOTA_IN_DISABLED",
	5: "IBC_TRANSFER_STATUS_TRANSFERS_PAUSED",
}

var IBCTransferStatus_value = map[string]int32{
	"IBC_TRANSFER_STATUS_UNSPECIFIED":        0,
	"IBC_TRANSFER_STATUS_QUOTA_DISABLED":     1,
	"IBC_TRANSFER_STATUS_QUOTA_ENABLED":      2,
	"IBC_TRANSFER_STATUS_QUOTA_OUT_DISABLED": 3,
	"IBC_TRANSFER_STATUS_QUOTA_IN_DISABLED":  4,
	"IBC_TRANSFER_STATUS_TRANSFERS_PAUSED":   5,
}

func (x IBCTransferStatus) String() string {
	return proto.EnumName(IBCTransferStatus_name, int32(x))
}

func (IBCTransferStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_651be1a0280abcb6, []int{0}
}

// Params of x/uibc module
type Params struct {
	// ibc_status defines the IBC ICS20 status (transfer quota or transfers disabled).
	IbcStatus IBCTransferStatus `protobuf:"varint,1,opt,name=ibc_status,json=ibcStatus,proto3,enum=umee.uibc.v1.IBCTransferStatus" json:"ibc_status,omitempty"`
	// total_quota defines the total outflow limit of ibc-transfer in USD
	TotalQuota github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=total_quota,json=totalQuota,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"total_quota"`
	// token_quota defines the outflow limit per token in USD
	TokenQuota github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=token_quota,json=tokenQuota,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"token_quota"`
	// quota_duration defines quota expires for each ibc-transfer denom in seconds
	QuotaDuration time.Duration `protobuf:"bytes,4,opt,name=quota_duration,json=quotaDuration,proto3,stdduration" json:"quota_duration,omitempty" yaml:"quota_duration"`
	// total_inflow_quota defines the total inflow limit of ibc-transfer in USD
	TotalInflowQuota github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=total_inflow_quota,json=totalInflowQuota,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"total_inflow_quota"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_651be1a0280abcb6, []int{0}
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

func (m *Params) GetIbcStatus() IBCTransferStatus {
	if m != nil {
		return m.IbcStatus
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
	proto.RegisterType((*Params)(nil), "umee.uibc.v1.Params")
}

func init() { proto.RegisterFile("umee/uibc/v1/quota.proto", fileDescriptor_651be1a0280abcb6) }

var fileDescriptor_651be1a0280abcb6 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0x80, 0xe3, 0xa4, 0xad, 0xd4, 0x2b, 0x54, 0xe1, 0x04, 0x52, 0xda, 0xc1, 0x0e, 0x81, 0x46,
	0xa1, 0x22, 0x67, 0xb5, 0x48, 0x0c, 0x08, 0x10, 0x71, 0xec, 0x4a, 0x96, 0x50, 0x9a, 0xc6, 0xf6,
	0x82, 0x84, 0x2c, 0xdb, 0xbd, 0x18, 0x93, 0xd8, 0x17, 0xec, 0x73, 0x4a, 0x26, 0x56, 0x46, 0x46,
	0x7e, 0x08, 0x3f, 0xa2, 0x63, 0xc5, 0x84, 0x18, 0x02, 0x4a, 0x36, 0x46, 0xf8, 0x03, 0xc8, 0x77,
	0x8e, 0xda, 0xaa, 0xcd, 0xd6, 0xc9, 0xf7, 0xfc, 0xbe, 0xfb, 0xee, 0xbd, 0x67, 0x1f, 0xa8, 0xa4,
	0x21, 0xc6, 0x72, 0x1a, 0xb8, 0x9e, 0x3c, 0xde, 0x93, 0x3f, 0xa4, 0x84, 0x3a, 0x68, 0x14, 0x13,
	0x4a, 0xe0, 0xad, 0x2c, 0x83, 0xb2, 0x0c, 0x1a, 0xef, 0x6d, 0xdf, 0xf5, 0x89, 0x4f, 0x58, 0x42,
	0xce, 0x56, 0x9c, 0xd9, 0x16, 0x7d, 0x42, 0xfc, 0x21, 0x96, 0x59, 0xe4, 0xa6, 0x7d, 0xf9, 0x38,
	0x8d, 0x1d, 0x1a, 0x90, 0x28, 0xcf, 0x6f, 0x79, 0x24, 0x09, 0x49, 0x62, 0xf3, 0x8d, 0x3c, 0xe0,
	0xa9, 0xda, 0xbf, 0x12, 0x58, 0xeb, 0x3a, 0xb1, 0x13, 0x26, 0xf0, 0x25, 0x00, 0x81, 0xeb, 0xd9,
	0x09, 0x75, 0x68, 0x9a, 0x54, 0x84, 0xaa, 0xd0, 0xd8, 0xdc, 0x97, 0xd0, 0xc5, 0xe3, 0x91, 0xae,
	0xb4, 0xcd, 0xd8, 0x89, 0x92, 0x3e, 0x8e, 0x0d, 0x86, 0xf5, 0xd6, 0x03, 0xd7, 0xe3, 0x4b, 0xf8,
	0x16, 0x6c, 0x50, 0x42, 0x9d, 0xa1, 0xcd, 0xca, 0xaf, 0x14, 0xab, 0x42, 0x63, 0x5d, 0x79, 0x7e,
	0x3a, 0x95, 0x0a, 0x3f, 0xa7, 0x52, 0xdd, 0x0f, 0xe8, 0xbb, 0xd4, 0x45, 0x1e, 0x09, 0xf3, 0x02,
	0xf2, 0x47, 0x33, 0x39, 0x1e, 0xc8, 0x74, 0x32, 0xc2, 0x09, 0x52, 0xb1, 0xf7, 0xfd, 0x5b, 0x13,
	0xe4, 0xf5, 0xa9, 0xd8, 0xeb, 0x01, 0x26, 0x3c, 0xca, 0x7c, 0x5c, 0x3f, 0xc0, 0x51, 0xae, 0x2f,
	0xdd, 0x8c, 0x7e, 0x80, 0x23, 0xae, 0xff, 0x04, 0x36, 0x99, 0xd8, 0x5e, 0xcc, 0xae, 0xb2, 0x52,
	0x15, 0x1a, 0x1b, 0xfb, 0x5b, 0x88, 0x0f, 0x17, 0x2d, 0x86, 0x8b, 0xd4, 0x1c, 0x50, 0x5e, 0x64,
	0x87, 0xff, 0x99, 0x4a, 0x95, 0xcb, 0x1b, 0x1f, 0x93, 0x30, 0xa0, 0x38, 0x1c, 0xd1, 0xc9, 0xdf,
	0xa9, 0x74, 0x6f, 0xe2, 0x84, 0xc3, 0x67, 0xb5, 0xcb, 0x44, 0xed, 0xeb, 0x2f, 0x49, 0xe8, 0xdd,
	0x66, 0x2f, 0x17, 0x36, 0xf8, 0x1e, 0x40, 0x3e, 0xbe, 0x20, 0xea, 0x0f, 0xc9, 0x49, 0xde, 0xe6,
	0xea, 0x0d, 0xb4, 0x59, 0x66, 0x5e, 0x9d, 0x69, 0x59, 0xb3, 0xbb, 0x9f, 0x8b, 0xe0, 0xce, 0x95,
	0x6f, 0x09, 0x1f, 0x00, 0x49, 0x57, 0xda, 0xb6, 0xd9, 0x6b, 0x75, 0x8c, 0x03, 0xad, 0x67, 0x1b,
	0x66, 0xcb, 0xb4, 0x0c, 0xdb, 0xea, 0x18, 0x5d, 0xad, 0xad, 0x1f, 0xe8, 0x9a, 0x5a, 0x2e, 0xc0,
	0x3a, 0xa8, 0x5d, 0x07, 0x1d, 0x59, 0x87, 0x66, 0xcb, 0x56, 0x75, 0xa3, 0xa5, 0xbc, 0xd6, 0xd4,
	0xb2, 0x00, 0x77, 0xc0, 0xfd, 0xe5, 0x9c, 0xd6, 0xe1, 0x58, 0x11, 0xee, 0x82, 0xfa, 0x72, 0xec,
	0xd0, 0x32, 0xcf, 0x95, 0x25, 0xf8, 0x08, 0xec, 0x2c, 0x67, 0xf5, 0xce, 0x39, 0xba, 0x02, 0x1b,
	0xe0, 0xe1, 0x75, 0xe8, 0x22, 0x36, 0xec, 0x6e, 0xcb, 0x32, 0x34, 0xb5, 0xbc, 0xaa, 0xbc, 0x3a,
	0x9d, 0x89, 0xc2, 0xd9, 0x4c, 0x14, 0x7e, 0xcf, 0x44, 0xe1, 0xcb, 0x5c, 0x2c, 0x9c, 0xcd, 0xc5,
	0xc2, 0x8f, 0xb9, 0x58, 0x78, 0x73, 0x71, 0xd8, 0xd9, 0x2d, 0x68, 0x46, 0x98, 0x9e, 0x90, 0x78,
	0xc0, 0x02, 0x79, 0xfc, 0x54, 0xfe, 0xc8, 0x2e, 0xac, 0xbb, 0xc6, 0xfe, 0x8c, 0x27, 0xff, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x88, 0xbe, 0x35, 0x75, 0xc4, 0x03, 0x00, 0x00,
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
	{
		size := m.TotalInflowQuota.Size()
		i -= size
		if _, err := m.TotalInflowQuota.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuota(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.QuotaDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.QuotaDuration):])
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
	if m.IbcStatus != 0 {
		i = encodeVarintQuota(dAtA, i, uint64(m.IbcStatus))
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
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IbcStatus != 0 {
		n += 1 + sovQuota(uint64(m.IbcStatus))
	}
	l = m.TotalQuota.Size()
	n += 1 + l + sovQuota(uint64(l))
	l = m.TokenQuota.Size()
	n += 1 + l + sovQuota(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.QuotaDuration)
	n += 1 + l + sovQuota(uint64(l))
	l = m.TotalInflowQuota.Size()
	n += 1 + l + sovQuota(uint64(l))
	return n
}

func sovQuota(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuota(x uint64) (n int) {
	return sovQuota(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
				return fmt.Errorf("proto: wrong wireType = %d for field IbcStatus", wireType)
			}
			m.IbcStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.QuotaDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalInflowQuota", wireType)
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
			if err := m.TotalInflowQuota.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
