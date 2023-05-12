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

// QuotaStatus status of ibc-transfer quota check for inflow and outflow
type IBCTransferQuotaStatus int32

const (
	// UNSPECIFIED  defines a no-op status.
	IBCTransferQuotaStatus_IBC_TRANSFER_QUOTA_STATUS_UNSPECIFIED IBCTransferQuotaStatus = 0
	// DISABLED: all quota checks are disabled.
	IBCTransferQuotaStatus_IBC_TRANSFER_QUOTA_STATUS_DISABLED IBCTransferQuotaStatus = 1
	// ENABLED: all quota checks are enabled.
	IBCTransferQuotaStatus_IBC_TRANSFER_QUOTA_STATUS_ENABLED IBCTransferQuotaStatus = 2
	// DISABLED OUT: quota outflow checks are disabled and inflows are enabled.
	IBCTransferQuotaStatus_IBC_TRANSFER_QUOTA_STATUS_OUT_DISABLED IBCTransferQuotaStatus = 3
	// DISABLED IN: quota inflow checks are disabled and outflows are enabled.
	IBCTransferQuotaStatus_IBC_TRANSFER_QUOTA_STATUS_IN_DISABLED IBCTransferQuotaStatus = 4
)

var IBCTransferQuotaStatus_name = map[int32]string{
	0: "IBC_TRANSFER_QUOTA_STATUS_UNSPECIFIED",
	1: "IBC_TRANSFER_QUOTA_STATUS_DISABLED",
	2: "IBC_TRANSFER_QUOTA_STATUS_ENABLED",
	3: "IBC_TRANSFER_QUOTA_STATUS_OUT_DISABLED",
	4: "IBC_TRANSFER_QUOTA_STATUS_IN_DISABLED",
}

var IBCTransferQuotaStatus_value = map[string]int32{
	"IBC_TRANSFER_QUOTA_STATUS_UNSPECIFIED":  0,
	"IBC_TRANSFER_QUOTA_STATUS_DISABLED":     1,
	"IBC_TRANSFER_QUOTA_STATUS_ENABLED":      2,
	"IBC_TRANSFER_QUOTA_STATUS_OUT_DISABLED": 3,
	"IBC_TRANSFER_QUOTA_STATUS_IN_DISABLED":  4,
}

func (x IBCTransferQuotaStatus) String() string {
	return proto.EnumName(IBCTransferQuotaStatus_name, int32(x))
}

func (IBCTransferQuotaStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_651be1a0280abcb6, []int{0}
}

// Params of x/uibc module
type Params struct {
	// quota_status defines the wethever quota check enabled, disbaled for inflow and outlfow.
	QuotaStatus IBCTransferQuotaStatus `protobuf:"varint,1,opt,name=quota_status,json=quotaStatus,proto3,enum=umee.uibc.v1.IBCTransferQuotaStatus" json:"quota_status,omitempty"`
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

func (m *Params) GetQuotaStatus() IBCTransferQuotaStatus {
	if m != nil {
		return m.QuotaStatus
	}
	return IBCTransferQuotaStatus_IBC_TRANSFER_QUOTA_STATUS_UNSPECIFIED
}

func (m *Params) GetQuotaDuration() time.Duration {
	if m != nil {
		return m.QuotaDuration
	}
	return 0
}

func init() {
	proto.RegisterEnum("umee.uibc.v1.IBCTransferQuotaStatus", IBCTransferQuotaStatus_name, IBCTransferQuotaStatus_value)
	proto.RegisterType((*Params)(nil), "umee.uibc.v1.Params")
}

func init() { proto.RegisterFile("umee/uibc/v1/quota.proto", fileDescriptor_651be1a0280abcb6) }

var fileDescriptor_651be1a0280abcb6 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0xcd, 0x4d, 0x4d,
	0xd5, 0x2f, 0xcd, 0x4c, 0x4a, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0x2c, 0xcd, 0x2f, 0x49, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0xc9, 0xe8, 0x81, 0x64, 0xf4, 0xca, 0x0c, 0xa5, 0x44,
	0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x12, 0xfa, 0x20, 0x16, 0x44, 0x8d, 0x94, 0x5c, 0x7a, 0x7e, 0x7e,
	0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x9f, 0x52, 0x5a, 0x94, 0x58, 0x92, 0x99,
	0x9f, 0x07, 0x95, 0x97, 0x4c, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x87, 0x68, 0x84, 0x70, 0x20,
	0x52, 0x4a, 0x13, 0x98, 0xb9, 0xd8, 0x02, 0x12, 0x8b, 0x12, 0x73, 0x8b, 0x85, 0xdc, 0xb9, 0x78,
	0xc0, 0x16, 0xc7, 0x17, 0x97, 0x24, 0x96, 0x94, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19,
	0xa9, 0xe8, 0x21, 0x3b, 0x40, 0xcf, 0xd3, 0xc9, 0x39, 0xa4, 0x28, 0x31, 0xaf, 0x38, 0x2d, 0xb5,
	0x28, 0x10, 0xa4, 0x38, 0x18, 0xac, 0x36, 0x88, 0xbb, 0x10, 0xc1, 0x11, 0x8a, 0xe5, 0xe2, 0x2e,
	0xc9, 0x2f, 0x49, 0xcc, 0x89, 0x07, 0x0b, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x3a, 0xd9, 0x9c,
	0xb8, 0x27, 0xcf, 0x70, 0xeb, 0x9e, 0xbc, 0x5a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0x2e, 0xd4, 0x25, 0x50, 0x4a, 0xb7, 0x38, 0x25, 0x5b, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x58,
	0xcf, 0x25, 0x35, 0xf9, 0xd2, 0x16, 0x5d, 0x2e, 0xa8, 0x43, 0x5d, 0x52, 0x93, 0x83, 0xb8, 0xc0,
	0x06, 0x82, 0x6d, 0x84, 0x18, 0x9f, 0x9d, 0x9a, 0x07, 0x35, 0x9e, 0x99, 0x3a, 0xc6, 0x67, 0xa7,
	0xe6, 0x41, 0x8c, 0xaf, 0xe7, 0xe2, 0x83, 0x04, 0x03, 0x2c, 0x10, 0x25, 0x58, 0x14, 0x18, 0x35,
	0xb8, 0x8d, 0x24, 0xf5, 0x20, 0xa1, 0xac, 0x07, 0x0b, 0x65, 0x3d, 0x17, 0xa8, 0x02, 0x27, 0x5b,
	0x90, 0xe5, 0xaf, 0xee, 0xc9, 0x4b, 0xa0, 0x6a, 0xd4, 0xc9, 0xcf, 0xcd, 0x2c, 0x49, 0xcd, 0x2d,
	0x28, 0xa9, 0xfc, 0x74, 0x4f, 0x5e, 0xb4, 0x32, 0x31, 0x37, 0xc7, 0x4a, 0x09, 0x55, 0x85, 0xd2,
	0x8c, 0xfb, 0xf2, 0x8c, 0x41, 0xbc, 0x60, 0x41, 0x98, 0x69, 0x5a, 0x2f, 0x19, 0xb9, 0xc4, 0xb0,
	0x07, 0xb3, 0x90, 0x26, 0x97, 0xaa, 0xa7, 0x93, 0x73, 0x7c, 0x48, 0x90, 0xa3, 0x5f, 0xb0, 0x9b,
	0x6b, 0x50, 0x7c, 0x60, 0xa8, 0x7f, 0x88, 0x63, 0x7c, 0x70, 0x88, 0x63, 0x48, 0x68, 0x70, 0x7c,
	0xa8, 0x5f, 0x70, 0x80, 0xab, 0xb3, 0xa7, 0x9b, 0xa7, 0xab, 0x8b, 0x00, 0x83, 0x90, 0x1a, 0x97,
	0x12, 0x6e, 0xa5, 0x2e, 0x9e, 0xc1, 0x8e, 0x4e, 0x3e, 0xae, 0x2e, 0x02, 0x8c, 0x42, 0xaa, 0x5c,
	0x8a, 0xb8, 0xd5, 0xb9, 0xfa, 0x41, 0x94, 0x31, 0x09, 0x69, 0x71, 0xa9, 0xe1, 0x56, 0xe6, 0x1f,
	0x1a, 0x82, 0x30, 0x92, 0x19, 0xbf, 0x2b, 0x3d, 0xfd, 0x10, 0x4a, 0x59, 0x9c, 0x1c, 0x4e, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x0a, 0x39, 0x22, 0x41, 0x29, 0x50, 0x37, 0x2f,
	0xb5, 0xa4, 0x3c, 0xbf, 0x28, 0x1b, 0xcc, 0xd1, 0x2f, 0x33, 0xd1, 0xaf, 0x00, 0x67, 0x97, 0x24,
	0x36, 0x70, 0x74, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x6d, 0x67, 0xd1, 0x42, 0x03,
	0x00, 0x00,
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
	if m.QuotaStatus != 0 {
		i = encodeVarintQuota(dAtA, i, uint64(m.QuotaStatus))
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
	if m.QuotaStatus != 0 {
		n += 1 + sovQuota(uint64(m.QuotaStatus))
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
				return fmt.Errorf("proto: wrong wireType = %d for field QuotaStatus", wireType)
			}
			m.QuotaStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuota
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
