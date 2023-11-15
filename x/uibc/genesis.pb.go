// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/uibc/v1/genesis.proto

package uibc

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the uibc module's genesis state.
type GenesisState struct {
	Params   Params                                      `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Outflows github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=outflows,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"outflows"`
	// total_outflow_sum defines the total outflow sum of ibc-transfer in USD.
	TotalOutflowSum github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=total_outflow_sum,json=totalOutflowSum,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"total_outflow_sum"`
	// quota_expires defines quota expire time (as unix timestamp) for ibc-transfer denom.
	QuotaExpires time.Time `protobuf:"bytes,4,opt,name=quota_expires,json=quotaExpires,proto3,stdtime" json:"quota_duration,omitempty" yaml:"quota_expires"`
	// inflows tracks IBC inflow transfers (in USD) for each denom during quota period.
	Inflows github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,5,rep,name=inflows,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"inflows"`
	// total_inflow_sum defines tracks total sum of IBC inflow transfers (in USD) during quota period.
	TotalInflowSum github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=total_inflow_sum,json=totalInflowSum,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"total_inflow_sum"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0196ecf2d08401fb, []int{0}
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
	proto.RegisterType((*GenesisState)(nil), "umee.uibc.v1.GenesisState")
}

func init() { proto.RegisterFile("umee/uibc/v1/genesis.proto", fileDescriptor_0196ecf2d08401fb) }

var fileDescriptor_0196ecf2d08401fb = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xbf, 0x8e, 0x13, 0x3d,
	0x14, 0xc5, 0x67, 0xbe, 0xcd, 0x17, 0xc0, 0x1b, 0xfe, 0x8d, 0x52, 0x0c, 0x11, 0x9a, 0x89, 0x52,
	0xa0, 0x48, 0x10, 0x5b, 0xc9, 0x4a, 0x14, 0x88, 0x2a, 0x04, 0x21, 0x2a, 0x50, 0x96, 0x8a, 0x26,
	0xf2, 0xcc, 0x3a, 0xb3, 0x56, 0xe2, 0xb9, 0xc3, 0xd8, 0xce, 0x6e, 0x0a, 0xde, 0x61, 0x9f, 0x83,
	0x9a, 0x87, 0x48, 0xb9, 0xa2, 0x42, 0x14, 0x59, 0x48, 0x3a, 0x1a, 0x24, 0x9e, 0x00, 0x8d, 0xed,
	0xa0, 0xdd, 0x8e, 0x02, 0xaa, 0x99, 0x7b, 0x8f, 0xef, 0xb9, 0xc7, 0x3f, 0x19, 0xb5, 0xb4, 0x60,
	0x8c, 0x68, 0x9e, 0xa4, 0x64, 0xd1, 0x27, 0x19, 0xcb, 0x99, 0xe4, 0x12, 0x17, 0x25, 0x28, 0x08,
	0x1a, 0x95, 0x86, 0x2b, 0x0d, 0x2f, 0xfa, 0xad, 0x66, 0x06, 0x19, 0x18, 0x81, 0x54, 0x7f, 0xf6,
	0x4c, 0xeb, 0x5e, 0x0a, 0x52, 0x80, 0x9c, 0x58, 0xc1, 0x16, 0x4e, 0x8a, 0x6c, 0x45, 0x12, 0x2a,
	0x19, 0x59, 0xf4, 0x13, 0xa6, 0x68, 0x9f, 0xa4, 0xc0, 0x73, 0xa7, 0xc7, 0x19, 0x40, 0x36, 0x67,
	0xc4, 0x54, 0x89, 0x9e, 0x12, 0xc5, 0x05, 0x93, 0x8a, 0x8a, 0xc2, 0x1d, 0x08, 0xaf, 0x64, 0x7b,
	0xa7, 0x41, 0x51, 0xab, 0x74, 0x7e, 0xd4, 0x50, 0xe3, 0x85, 0xcd, 0x7a, 0xa8, 0xa8, 0x62, 0xc1,
	0x00, 0xd5, 0x0b, 0x5a, 0x52, 0x21, 0x43, 0xbf, 0xed, 0x77, 0xf7, 0x07, 0x4d, 0x7c, 0x39, 0x3b,
	0x7e, 0x6d, 0xb4, 0x61, 0x6d, 0xb5, 0x8e, 0xbd, 0xb1, 0x3b, 0x19, 0x08, 0x74, 0x1d, 0xb4, 0x9a,
	0xce, 0xe1, 0x44, 0x86, 0xff, 0xb5, 0xf7, 0xba, 0xfb, 0x83, 0xfb, 0xd8, 0x5d, 0xa0, 0x8a, 0x8c,
	0x5d, 0x64, 0x3c, 0x62, 0xe9, 0x33, 0xe0, 0xf9, 0xf0, 0xa0, 0x9a, 0xfe, 0x70, 0x11, 0x3f, 0xcc,
	0xb8, 0x3a, 0xd6, 0x09, 0x4e, 0x41, 0xb8, 0x0b, 0xbb, 0x4f, 0x4f, 0x1e, 0xcd, 0x88, 0x5a, 0x16,
	0x4c, 0xee, 0x66, 0xe4, 0xf8, 0xf7, 0x8a, 0xe0, 0x18, 0xdd, 0x55, 0xa0, 0xe8, 0x7c, 0xe2, 0x3a,
	0x13, 0xa9, 0x45, 0xb8, 0xd7, 0xf6, 0xbb, 0x37, 0x86, 0x4f, 0x2b, 0xe7, 0x2f, 0xeb, 0xf8, 0xc1,
	0x9f, 0x39, 0x7f, 0xfa, 0xd8, 0x43, 0x2e, 0xe8, 0x88, 0xa5, 0xe3, 0xdb, 0xc6, 0xf6, 0x95, 0x75,
	0x3d, 0xd4, 0x22, 0x78, 0x8f, 0x6e, 0x1a, 0x58, 0x13, 0x76, 0x5a, 0xf0, 0x92, 0xc9, 0xb0, 0x66,
	0x98, 0xb4, 0xb0, 0x05, 0x8e, 0x77, 0xc0, 0xf1, 0x9b, 0x1d, 0x70, 0x9b, 0xe0, 0xfb, 0x3a, 0x0e,
	0xed, 0xe0, 0x91, 0x2e, 0xa9, 0xe2, 0x90, 0x3f, 0x02, 0xc1, 0x15, 0x13, 0x85, 0x5a, 0xfe, 0x5c,
	0xc7, 0xcd, 0x25, 0x15, 0xf3, 0x27, 0x9d, 0x2b, 0xd6, 0x9d, 0xb3, 0x8b, 0xd8, 0x1f, 0x37, 0x4c,
	0xef, 0xb9, 0x6d, 0x05, 0x33, 0x74, 0x8d, 0xe7, 0x16, 0xeb, 0xff, 0xff, 0x0a, 0xeb, 0x6e, 0x43,
	0x30, 0x45, 0x77, 0x2c, 0x55, 0xdb, 0x30, 0x50, 0xeb, 0x7f, 0x01, 0xea, 0x2d, 0xe3, 0xfa, 0x32,
	0x77, 0x4c, 0x87, 0xa3, 0xd5, 0xb7, 0xc8, 0x5b, 0x6d, 0x22, 0xff, 0x7c, 0x13, 0xf9, 0x5f, 0x37,
	0x91, 0x7f, 0xb6, 0x8d, 0xbc, 0xf3, 0x6d, 0xe4, 0x7d, 0xde, 0x46, 0xde, 0xdb, 0xcb, 0x3b, 0xaa,
	0x87, 0xd7, 0xcb, 0x99, 0x3a, 0x81, 0x72, 0x66, 0x0a, 0xb2, 0x78, 0x4c, 0x4e, 0xcd, 0x33, 0x4e,
	0xea, 0x06, 0xfd, 0xc1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0x4f, 0x6e, 0xd5, 0x76, 0x03,
	0x00, 0x00,
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
	{
		size := m.TotalInflowSum.Size()
		i -= size
		if _, err := m.TotalInflowSum.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Inflows) > 0 {
		for iNdEx := len(m.Inflows) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Inflows[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.QuotaExpires, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.QuotaExpires):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size := m.TotalOutflowSum.Size()
		i -= size
		if _, err := m.TotalOutflowSum.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Outflows) > 0 {
		for iNdEx := len(m.Outflows) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Outflows[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Outflows) > 0 {
		for _, e := range m.Outflows {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.TotalOutflowSum.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.QuotaExpires)
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Inflows) > 0 {
		for _, e := range m.Inflows {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.TotalInflowSum.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Outflows", wireType)
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
			m.Outflows = append(m.Outflows, types.DecCoin{})
			if err := m.Outflows[len(m.Outflows)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalOutflowSum", wireType)
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
			if err := m.TotalOutflowSum.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuotaExpires", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.QuotaExpires, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflows", wireType)
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
			m.Inflows = append(m.Inflows, types.DecCoin{})
			if err := m.Inflows[len(m.Inflows)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalInflowSum", wireType)
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
			if err := m.TotalInflowSum.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
