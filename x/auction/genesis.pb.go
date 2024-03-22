// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/auction/v1/genesis.proto

package auction

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the x/auction module's genesis state.
type GenesisState struct {
	RewardsParams RewardsParams `protobuf:"bytes,1,opt,name=rewards_params,json=rewardsParams,proto3" json:"rewards_params"`
	// latest active (in bid phase) reward auction.
	RewardAuctionId uint32 `protobuf:"varint,2,opt,name=reward_auction_id,json=rewardAuctionId,proto3" json:"reward_auction_id,omitempty"`
	// latest highest bid
	HighestBidder string       `protobuf:"bytes,3,opt,name=highest_bidder,json=highestBidder,proto3" json:"highest_bidder,omitempty"`
	RewardTokens  []types.Coin `protobuf:"bytes,4,rep,name=reward_tokens,json=rewardTokens,proto3" json:"reward_tokens"`
	// tokens collected for the next auction, while the current reward auction is still in the
	// bid phase.
	NextRewardTokens []types.Coin `protobuf:"bytes,5,rep,name=next_reward_tokens,json=nextRewardTokens,proto3" json:"next_reward_tokens"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e83c50dcf6ac7b, []int{0}
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
	proto.RegisterType((*GenesisState)(nil), "umee.auction.v1.GenesisState")
}

func init() { proto.RegisterFile("umee/auction/v1/genesis.proto", fileDescriptor_15e83c50dcf6ac7b) }

var fileDescriptor_15e83c50dcf6ac7b = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x4b, 0xf3, 0x40,
	0x18, 0x4c, 0xda, 0xbe, 0x2f, 0x98, 0x7e, 0x69, 0xf0, 0x10, 0x0b, 0xae, 0x41, 0x10, 0xa2, 0xe0,
	0x2e, 0xa9, 0xe0, 0xdd, 0x2a, 0x14, 0x11, 0x41, 0xa2, 0x27, 0x2f, 0x21, 0x1f, 0x4b, 0xba, 0x94,
	0x64, 0xcb, 0xee, 0x36, 0xed, 0xcf, 0xf0, 0x3f, 0x79, 0xe9, 0xb1, 0x47, 0x4f, 0xa2, 0xed, 0x1f,
	0x91, 0xec, 0x6e, 0xa1, 0xea, 0xc5, 0xdb, 0xee, 0xcc, 0x33, 0x33, 0x0f, 0xf3, 0x58, 0x87, 0xd3,
	0x1c, 0x63, 0x14, 0x4d, 0x13, 0x41, 0x68, 0x81, 0x4a, 0x1f, 0x65, 0xb8, 0xc0, 0x9c, 0x70, 0x38,
	0x61, 0x54, 0x50, 0xbb, 0x5b, 0xd1, 0x50, 0xd3, 0xb0, 0xf4, 0x7b, 0x20, 0xa1, 0x3c, 0xa7, 0x1c,
	0xc5, 0x11, 0xc7, 0xa8, 0xf4, 0x63, 0x2c, 0x22, 0x1f, 0x25, 0x94, 0x14, 0x4a, 0xd0, 0xdb, 0xcf,
	0x68, 0x46, 0xe5, 0x13, 0x55, 0x2f, 0x8d, 0xfe, 0x4a, 0xd9, 0x38, 0x4a, 0xfa, 0xf8, 0xb5, 0x66,
	0xb5, 0x86, 0x2a, 0xf7, 0x51, 0x44, 0x02, 0xdb, 0x77, 0x56, 0x87, 0xe1, 0x59, 0xc4, 0x52, 0x1e,
	0x4e, 0x22, 0x16, 0xe5, 0xdc, 0x31, 0x5d, 0xd3, 0x6b, 0xf6, 0x01, 0xfc, 0xb1, 0x0f, 0x0c, 0xd4,
	0xd8, 0x83, 0x9c, 0x1a, 0x34, 0x16, 0xef, 0x47, 0x46, 0xd0, 0x66, 0xdb, 0xa0, 0x7d, 0x66, 0xed,
	0x29, 0x20, 0xd4, 0xba, 0x90, 0xa4, 0x4e, 0xcd, 0x35, 0xbd, 0x76, 0xd0, 0x55, 0xc4, 0x95, 0xc2,
	0x6f, 0x53, 0xfb, 0xc4, 0xea, 0x8c, 0x48, 0x36, 0xc2, 0x5c, 0x84, 0x31, 0x49, 0x53, 0xcc, 0x9c,
	0xba, 0x6b, 0x7a, 0x3b, 0x41, 0x5b, 0xa3, 0x03, 0x09, 0xda, 0x37, 0x96, 0xce, 0x08, 0x05, 0x1d,
	0xe3, 0x82, 0x3b, 0x0d, 0xb7, 0xee, 0x35, 0xfb, 0x07, 0x50, 0xb5, 0x03, 0xab, 0x76, 0xa0, 0x6e,
	0x07, 0x5e, 0x53, 0x52, 0xe8, 0xcd, 0x5a, 0x4a, 0xf5, 0x24, 0x45, 0xf6, 0xbd, 0x65, 0x17, 0x78,
	0x2e, 0xc2, 0xef, 0x56, 0xff, 0xfe, 0x66, 0xb5, 0x5b, 0x49, 0x83, 0x2d, 0xbb, 0xc1, 0x70, 0xf1,
	0x09, 0x8c, 0xc5, 0x0a, 0x98, 0xcb, 0x15, 0x30, 0x3f, 0x56, 0xc0, 0x7c, 0x59, 0x03, 0x63, 0xb9,
	0x06, 0xc6, 0xdb, 0x1a, 0x18, 0xcf, 0xa7, 0x19, 0x11, 0xa3, 0x69, 0x0c, 0x13, 0x9a, 0xa3, 0xaa,
	0xc4, 0xf3, 0x02, 0x8b, 0x19, 0x65, 0x63, 0xf9, 0x41, 0xe5, 0x25, 0x9a, 0x6f, 0x8e, 0x12, 0xff,
	0x97, 0x57, 0xb9, 0xf8, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x86, 0x77, 0xd1, 0x1c, 0x02, 0x00,
	0x00,
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
	if len(m.NextRewardTokens) > 0 {
		for iNdEx := len(m.NextRewardTokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NextRewardTokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.RewardTokens) > 0 {
		for iNdEx := len(m.RewardTokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RewardTokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.HighestBidder) > 0 {
		i -= len(m.HighestBidder)
		copy(dAtA[i:], m.HighestBidder)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.HighestBidder)))
		i--
		dAtA[i] = 0x1a
	}
	if m.RewardAuctionId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RewardAuctionId))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.RewardsParams.MarshalToSizedBuffer(dAtA[:i])
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
	l = m.RewardsParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.RewardAuctionId != 0 {
		n += 1 + sovGenesis(uint64(m.RewardAuctionId))
	}
	l = len(m.HighestBidder)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.RewardTokens) > 0 {
		for _, e := range m.RewardTokens {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NextRewardTokens) > 0 {
		for _, e := range m.NextRewardTokens {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsParams", wireType)
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
			if err := m.RewardsParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAuctionId", wireType)
			}
			m.RewardAuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardAuctionId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HighestBidder", wireType)
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
			m.HighestBidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardTokens", wireType)
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
			m.RewardTokens = append(m.RewardTokens, types.Coin{})
			if err := m.RewardTokens[len(m.RewardTokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextRewardTokens", wireType)
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
			m.NextRewardTokens = append(m.NextRewardTokens, types.Coin{})
			if err := m.NextRewardTokens[len(m.NextRewardTokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
