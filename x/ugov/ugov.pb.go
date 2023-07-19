// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/ugov/v1/ugov.proto

package ugov

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// InflationParams params for changing the inflation min rate and max rate of mint module params.
type InflationParams struct {
	// max_supply is the maximum supply for liquidation.
	MaxSupply types.Coin `protobuf:"bytes,1,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply"`
	// inflation_cycle duration after which inflation rates are changed.
	InflationCycle time.Duration `protobuf:"bytes,2,opt,name=inflation_cycle,json=inflationCycle,proto3,stdduration" json:"inflation_cycle,omitempty" yaml:"inflation_cycle"`
	// inflation_reduction_rate for every inflation cycle.
	InflationReductionRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=inflation_reduction_rate,json=inflationReductionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_reduction_rate"`
}

func (m *InflationParams) Reset()         { *m = InflationParams{} }
func (m *InflationParams) String() string { return proto.CompactTextString(m) }
func (*InflationParams) ProtoMessage()    {}
func (*InflationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_b75ef21394c8e122, []int{0}
}
func (m *InflationParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InflationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InflationParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InflationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InflationParams.Merge(m, src)
}
func (m *InflationParams) XXX_Size() int {
	return m.Size()
}
func (m *InflationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_InflationParams.DiscardUnknown(m)
}

var xxx_messageInfo_InflationParams proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InflationParams)(nil), "umee.ugov.v1.InflationParams")
}

func init() { proto.RegisterFile("umee/ugov/v1/ugov.proto", fileDescriptor_b75ef21394c8e122) }

var fileDescriptor_b75ef21394c8e122 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0x3f, 0x8f, 0xd3, 0x30,
	0x14, 0x8f, 0x0b, 0x42, 0x6a, 0x40, 0x54, 0x8a, 0x50, 0x49, 0x3b, 0x38, 0x55, 0x07, 0xd4, 0x81,
	0xda, 0x0a, 0x88, 0x05, 0x21, 0x90, 0xd2, 0x2c, 0x6c, 0x28, 0x6c, 0x2c, 0x95, 0xe3, 0xba, 0x21,
	0x6a, 0x1c, 0x47, 0x89, 0x13, 0x9a, 0x99, 0x2f, 0xc0, 0xc8, 0x07, 0xe1, 0x43, 0x74, 0xa3, 0x62,
	0x42, 0x37, 0xe4, 0xee, 0xda, 0xed, 0xc6, 0xfb, 0x04, 0xa7, 0x38, 0x49, 0xef, 0xd4, 0xc9, 0xef,
	0xbd, 0xdf, 0x9f, 0xf7, 0xd3, 0x4b, 0xf4, 0x97, 0x39, 0x67, 0x0c, 0xe7, 0x81, 0x28, 0x70, 0x61,
	0xab, 0x17, 0x25, 0xa9, 0x90, 0xc2, 0x78, 0x56, 0x03, 0x48, 0x0d, 0x0a, 0x7b, 0xfc, 0x22, 0x10,
	0x81, 0x50, 0x00, 0xae, 0xab, 0x86, 0x33, 0x86, 0x54, 0x64, 0x5c, 0x64, 0xd8, 0x27, 0x19, 0xc3,
	0x85, 0xed, 0x33, 0x49, 0x6c, 0x4c, 0x45, 0x18, 0x77, 0x78, 0x20, 0x44, 0x10, 0x31, 0xac, 0x3a,
	0x3f, 0x5f, 0xe3, 0x55, 0x9e, 0x12, 0x19, 0x8a, 0x0e, 0x1f, 0x35, 0xfa, 0x65, 0x63, 0xdc, 0x34,
	0x0d, 0x34, 0xfd, 0xdb, 0xd3, 0x07, 0x9f, 0xe3, 0x75, 0xa4, 0xe8, 0x5f, 0x48, 0x4a, 0x78, 0x66,
	0x7c, 0xd4, 0x75, 0x4e, 0xb6, 0xcb, 0x2c, 0x4f, 0x92, 0xa8, 0x34, 0xc1, 0x04, 0xcc, 0x9e, 0xbe,
	0x19, 0xa1, 0x56, 0x56, 0x67, 0x40, 0x6d, 0x06, 0xb4, 0x10, 0x61, 0xec, 0x3c, 0xde, 0x55, 0x96,
	0xe6, 0xf5, 0x39, 0xd9, 0x7e, 0x55, 0x0a, 0xe3, 0x27, 0xd0, 0x07, 0x61, 0xe7, 0xb9, 0xa4, 0x25,
	0x8d, 0x98, 0xd9, 0x6b, 0x5d, 0x9a, 0xa4, 0xa8, 0x4b, 0x8a, 0xdc, 0x36, 0xa9, 0xf3, 0xa9, 0x76,
	0xb9, 0xa9, 0xac, 0xd1, 0x99, 0xf2, 0xb5, 0xe0, 0xa1, 0x64, 0x3c, 0x91, 0xe5, 0x6d, 0x65, 0x0d,
	0x4b, 0xc2, 0xa3, 0xf7, 0xd3, 0x33, 0xca, 0xf4, 0xf7, 0xa5, 0x05, 0xbc, 0xe7, 0xa7, 0xe9, 0xa2,
	0x1e, 0x1a, 0x85, 0x6e, 0xde, 0xf3, 0x52, 0xb6, 0xca, 0x69, 0x53, 0x11, 0xc9, 0xcc, 0x47, 0x13,
	0x30, 0xeb, 0x3b, 0x1f, 0xea, 0x95, 0x17, 0x95, 0xf5, 0x2a, 0x08, 0xe5, 0xf7, 0xdc, 0x47, 0x54,
	0xf0, 0xf6, 0x38, 0xed, 0x33, 0xcf, 0x56, 0x1b, 0x2c, 0xcb, 0x84, 0x65, 0xc8, 0x65, 0xf4, 0xdf,
	0x9f, 0xb9, 0xde, 0x1e, 0xc1, 0x65, 0xd4, 0x1b, 0x9e, 0xdc, 0xbd, 0xce, 0xdc, 0x23, 0x92, 0x39,
	0xee, 0xee, 0x1a, 0x6a, 0xbb, 0x03, 0x04, 0xfb, 0x03, 0x04, 0x57, 0x07, 0x08, 0x7e, 0x1d, 0xa1,
	0xb6, 0x3f, 0x42, 0xed, 0xff, 0x11, 0x6a, 0xdf, 0x1e, 0xee, 0xaa, 0xbf, 0xfc, 0x3c, 0x66, 0xf2,
	0x87, 0x48, 0x37, 0xaa, 0xc1, 0xc5, 0x3b, 0xbc, 0x55, 0x3f, 0x87, 0xff, 0x44, 0x5d, 0xe8, 0xed,
	0x5d, 0x00, 0x00, 0x00, 0xff, 0xff, 0x62, 0x1a, 0x05, 0x7c, 0x38, 0x02, 0x00, 0x00,
}

func (m *InflationParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InflationParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InflationParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.InflationReductionRate.Size()
		i -= size
		if _, err := m.InflationReductionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintUgov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.InflationCycle, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.InflationCycle):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintUgov(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.MaxSupply.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintUgov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintUgov(dAtA []byte, offset int, v uint64) int {
	offset -= sovUgov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InflationParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MaxSupply.Size()
	n += 1 + l + sovUgov(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.InflationCycle)
	n += 1 + l + sovUgov(uint64(l))
	l = m.InflationReductionRate.Size()
	n += 1 + l + sovUgov(uint64(l))
	return n
}

func sovUgov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUgov(x uint64) (n int) {
	return sovUgov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InflationParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUgov
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
			return fmt.Errorf("proto: InflationParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InflationParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUgov
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
				return ErrInvalidLengthUgov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationCycle", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUgov
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
				return ErrInvalidLengthUgov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.InflationCycle, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationReductionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUgov
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
				return ErrInvalidLengthUgov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUgov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationReductionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUgov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUgov
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
func skipUgov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUgov
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
					return 0, ErrIntOverflowUgov
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
					return 0, ErrIntOverflowUgov
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
				return 0, ErrInvalidLengthUgov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUgov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUgov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUgov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUgov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUgov = fmt.Errorf("proto: unexpected end of group")
)
