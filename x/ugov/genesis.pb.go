// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/ugov/v1/genesis.proto

package ugov

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// GenesisState of the ugov module.
type GenesisState struct {
	MinGasPrice types.DecCoin `protobuf:"bytes,1,opt,name=min_gas_price,json=minGasPrice,proto3" json:"min_gas_price"`
	// Emergency Group address
	EmergencyGroup string `protobuf:"bytes,2,opt,name=emergency_group,json=emergencyGroup,proto3" json:"emergency_group,omitempty"`
	// InflationParams is params for inflation rate changes
	InflationParams InflationParams `protobuf:"bytes,3,opt,name=inflation_params,json=inflationParams,proto3" json:"inflation_params"`
	// Time when the current inflation cycle will end
	InflationCycleEnd time.Time `protobuf:"bytes,4,opt,name=inflation_cycle_end,json=inflationCycleEnd,proto3,stdtime" json:"inflation_cycle_end"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_82f39cd8e8ede8c7, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "umee.ugov.v1.GenesisState")
}

func init() { proto.RegisterFile("umee/ugov/v1/genesis.proto", fileDescriptor_82f39cd8e8ede8c7) }

var fileDescriptor_82f39cd8e8ede8c7 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcd, 0x6e, 0xd4, 0x30,
	0x14, 0x85, 0x93, 0x52, 0x21, 0x48, 0x0b, 0x85, 0x50, 0x89, 0x10, 0x81, 0xa7, 0x62, 0x81, 0xba,
	0xa9, 0xad, 0x80, 0x78, 0x80, 0xa6, 0x85, 0x11, 0x1b, 0x54, 0x4d, 0xbb, 0x62, 0x13, 0x39, 0xce,
	0xad, 0xb1, 0x18, 0xdb, 0x51, 0xec, 0x04, 0xfa, 0x16, 0xf3, 0x30, 0xec, 0xd9, 0xce, 0x72, 0xc4,
	0x8a, 0x15, 0x3f, 0x33, 0x2f, 0x82, 0xec, 0x64, 0x06, 0x66, 0x97, 0x9b, 0x73, 0xee, 0x97, 0x4f,
	0x37, 0x51, 0xda, 0x4a, 0x00, 0xd2, 0x72, 0xdd, 0x91, 0x2e, 0x23, 0x1c, 0x14, 0x18, 0x61, 0x70,
	0xdd, 0x68, 0xab, 0xe3, 0x7d, 0x97, 0x61, 0x97, 0xe1, 0x2e, 0x4b, 0x0f, 0xb9, 0xe6, 0xda, 0x07,
	0xc4, 0x3d, 0xf5, 0x9d, 0x14, 0x31, 0x6d, 0xa4, 0x36, 0xa4, 0xa4, 0x06, 0x48, 0x97, 0x95, 0x60,
	0x69, 0x46, 0x98, 0x16, 0x6a, 0xc8, 0x47, 0x5c, 0x6b, 0x3e, 0x05, 0xe2, 0xa7, 0xb2, 0xbd, 0x26,
	0x56, 0x48, 0x30, 0x96, 0xca, 0x7a, 0x28, 0x3c, 0xe9, 0x01, 0x45, 0x4f, 0xee, 0x87, 0x21, 0x7a,
	0xbc, 0xe5, 0xe6, 0x3d, 0x7c, 0xf0, 0xfc, 0xdb, 0x4e, 0xb4, 0x3f, 0xee, 0x55, 0x2f, 0x2d, 0xb5,
	0x10, 0xbf, 0x8d, 0xee, 0x49, 0xa1, 0x0a, 0x4e, 0x1d, 0x47, 0x30, 0x48, 0xc2, 0xa3, 0xf0, 0x78,
	0xef, 0xe5, 0x53, 0x3c, 0xf0, 0x9c, 0x1d, 0x1e, 0xec, 0xf0, 0x39, 0xb0, 0x33, 0x2d, 0x54, 0xbe,
	0x3b, 0xff, 0x39, 0x0a, 0x26, 0x7b, 0x52, 0xa8, 0x31, 0x35, 0x17, 0x6e, 0x2d, 0x3e, 0x8d, 0x0e,
	0x40, 0x42, 0xc3, 0x41, 0xb1, 0x9b, 0x82, 0x37, 0xba, 0xad, 0x93, 0x9d, 0xa3, 0xf0, 0xf8, 0x6e,
	0x9e, 0x7c, 0xff, 0x7a, 0x72, 0x38, 0xc0, 0x4e, 0xab, 0xaa, 0x01, 0x63, 0x2e, 0x6d, 0x23, 0x14,
	0x9f, 0xdc, 0xdf, 0x2c, 0x8c, 0x5d, 0x3f, 0x7e, 0x1f, 0x3d, 0x10, 0xea, 0x7a, 0x4a, 0xad, 0xd0,
	0xaa, 0xa8, 0x69, 0x43, 0xa5, 0x49, 0x6e, 0x79, 0x9b, 0x67, 0xf8, 0xff, 0x7b, 0xe2, 0x77, 0xeb,
	0xd6, 0x85, 0x2f, 0x0d, 0x3a, 0x07, 0x62, 0xfb, 0x75, 0x7c, 0x15, 0x3d, 0xfa, 0xc7, 0x63, 0x37,
	0x6c, 0x0a, 0x05, 0xa8, 0x2a, 0xd9, 0xf5, 0xc8, 0x14, 0xf7, 0xe7, 0xc5, 0xeb, 0xf3, 0xe2, 0xab,
	0xf5, 0x79, 0xf3, 0x3b, 0x8e, 0x37, 0xfb, 0x35, 0x0a, 0x27, 0x0f, 0x37, 0x80, 0x33, 0xb7, 0xff,
	0x46, 0x55, 0xf9, 0xf9, 0xfc, 0x0f, 0x0a, 0xe6, 0x4b, 0x14, 0x2e, 0x96, 0x28, 0xfc, 0xbd, 0x44,
	0xe1, 0x6c, 0x85, 0x82, 0xc5, 0x0a, 0x05, 0x3f, 0x56, 0x28, 0xf8, 0xf0, 0x82, 0x0b, 0xfb, 0xb1,
	0x2d, 0x31, 0xd3, 0x92, 0x38, 0xe7, 0x13, 0x05, 0xf6, 0xb3, 0x6e, 0x3e, 0xf9, 0x81, 0x74, 0xaf,
	0xc9, 0x17, 0xff, 0x37, 0xca, 0xdb, 0xfe, 0xb3, 0xaf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf9,
	0xbe, 0x7e, 0xae, 0x45, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.InflationCycleEnd, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.InflationCycleEnd):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size, err := m.InflationParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.EmergencyGroup) > 0 {
		i -= len(m.EmergencyGroup)
		copy(dAtA[i:], m.EmergencyGroup)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.EmergencyGroup)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.MinGasPrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinGasPrice.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.EmergencyGroup)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.InflationParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.InflationCycleEnd)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGasPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinGasPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmergencyGroup", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EmergencyGroup = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationCycleEnd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.InflationCycleEnd, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
