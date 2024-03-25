// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/auction/v1/genesis.proto

package auction

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

// GenesisState defines the x/auction module's genesis state.
type GenesisState struct {
	RewardsParams RewardsParams `protobuf:"bytes,1,opt,name=rewards_params,json=rewardsParams,proto3" json:"rewards_params"`
	// Latest active (in bid phase) reward auction.
	RewardAuctionId uint32 `protobuf:"varint,2,opt,name=reward_auction_id,json=rewardAuctionId,proto3" json:"reward_auction_id,omitempty"`
	// Latest highest bid.
	HighestBidder string `protobuf:"bytes,3,opt,name=highest_bidder,json=highestBidder,proto3" json:"highest_bidder,omitempty"`
	// Tokens collected for the current auction.
	CurrentRewards []types.Coin `protobuf:"bytes,4,rep,name=current_rewards,json=currentRewards,proto3" json:"current_rewards"`
	// Tokens collected for the next auction, while the current reward auction is still in the
	// bid phase.
	NextRewards              []types.Coin `protobuf:"bytes,5,rep,name=next_rewards,json=nextRewards,proto3" json:"next_rewards"`
	CurrentRewardsAuctionEnd time.Time    `protobuf:"bytes,6,opt,name=current_rewards_auction_end,json=currentRewardsAuctionEnd,proto3,stdtime" json:"current_rewards_auction_end"`
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
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x6d, 0x52, 0x2a, 0x70, 0x9a, 0x44, 0x58, 0x3d, 0xb8, 0x41, 0x6c, 0x22, 0x4e, 0x01,
	0xa9, 0xbb, 0x4a, 0x91, 0xb8, 0xa2, 0x1a, 0xa1, 0x82, 0xb8, 0x20, 0x97, 0x13, 0x17, 0xcb, 0xf6,
	0x0e, 0x9b, 0x15, 0x78, 0x37, 0xda, 0x5d, 0xa7, 0x7d, 0x8c, 0x3e, 0x0c, 0x37, 0x5e, 0x20, 0xc7,
	0x8a, 0x13, 0x27, 0xfe, 0x24, 0x2f, 0x82, 0xbc, 0xbb, 0xe6, 0x4f, 0xb8, 0xf4, 0xe6, 0x99, 0x6f,
	0xbe, 0xd1, 0x6f, 0x3f, 0x4f, 0xf4, 0xa0, 0xa9, 0x01, 0x48, 0xd1, 0x54, 0x86, 0x4b, 0x41, 0x56,
	0x73, 0xc2, 0x40, 0x80, 0xe6, 0x1a, 0x2f, 0x95, 0x34, 0x32, 0x1e, 0xb5, 0x32, 0xf6, 0x32, 0x5e,
	0xcd, 0xc7, 0x13, 0x26, 0x25, 0xfb, 0x08, 0xc4, 0xca, 0x65, 0xf3, 0x9e, 0x18, 0x5e, 0x83, 0x36,
	0x45, 0xbd, 0x74, 0x8e, 0xf1, 0x51, 0x25, 0x75, 0x2d, 0x75, 0x6e, 0x2b, 0xe2, 0x0a, 0x2f, 0x21,
	0x57, 0x91, 0xb2, 0xd0, 0x40, 0x56, 0xf3, 0x12, 0x4c, 0x31, 0x27, 0x95, 0xe4, 0xc2, 0xeb, 0x87,
	0x4c, 0x32, 0xe9, 0x7c, 0xed, 0x97, 0xef, 0xfe, 0x47, 0xd8, 0xd1, 0x58, 0xf9, 0xe1, 0xe7, 0x5e,
	0x74, 0x70, 0xe6, 0x98, 0xcf, 0x4d, 0x61, 0x20, 0x7e, 0x1d, 0x0d, 0x15, 0x5c, 0x14, 0x8a, 0xea,
	0x7c, 0x59, 0xa8, 0xa2, 0xd6, 0x49, 0x38, 0x0d, 0x67, 0xfd, 0x13, 0x84, 0x77, 0xde, 0x82, 0x33,
	0x37, 0xf6, 0xc6, 0x4e, 0xa5, 0x7b, 0xeb, 0x6f, 0x93, 0x20, 0x1b, 0xa8, 0xbf, 0x9b, 0xf1, 0xe3,
	0xe8, 0x9e, 0x6b, 0xe4, 0xde, 0x97, 0x73, 0x9a, 0xdc, 0x9a, 0x86, 0xb3, 0x41, 0x36, 0x72, 0xc2,
	0xa9, 0xeb, 0xbf, 0xa2, 0xf1, 0xb3, 0x68, 0xb8, 0xe0, 0x6c, 0x01, 0xda, 0xe4, 0x25, 0xa7, 0x14,
	0x54, 0xd2, 0x9b, 0x86, 0xb3, 0xbb, 0x69, 0xf2, 0xe5, 0xd3, 0xf1, 0xa1, 0x0f, 0xe2, 0x94, 0x52,
	0x05, 0x5a, 0x9f, 0x1b, 0xc5, 0x05, 0xcb, 0x06, 0x7e, 0x3e, 0xb5, 0xe3, 0xf1, 0xcb, 0x68, 0x54,
	0x35, 0x4a, 0x81, 0x30, 0xb9, 0xa7, 0x48, 0xf6, 0xa6, 0xbd, 0x59, 0xff, 0xe4, 0x08, 0x7b, 0x7b,
	0x9b, 0x1c, 0xf6, 0xc9, 0xe1, 0xe7, 0x92, 0x0b, 0x4f, 0x3d, 0xf4, 0x3e, 0xff, 0xa2, 0x38, 0x8d,
	0x0e, 0x04, 0x5c, 0xfe, 0x59, 0x73, 0xfb, 0x66, 0x6b, 0xfa, 0xad, 0xa9, 0xdb, 0x51, 0x45, 0xf7,
	0x77, 0x68, 0x7e, 0x67, 0x00, 0x82, 0x26, 0xfb, 0x36, 0xd4, 0x31, 0x76, 0xf7, 0x80, 0xbb, 0x7b,
	0xc0, 0x6f, 0xbb, 0x7b, 0x48, 0xef, 0xb4, 0x3b, 0xaf, 0xbe, 0x4f, 0xc2, 0x2c, 0xf9, 0x17, 0xcf,
	0x47, 0xf6, 0x42, 0xd0, 0xf4, 0x6c, 0xfd, 0x13, 0x05, 0xeb, 0x0d, 0x0a, 0xaf, 0x37, 0x28, 0xfc,
	0xb1, 0x41, 0xe1, 0xd5, 0x16, 0x05, 0xd7, 0x5b, 0x14, 0x7c, 0xdd, 0xa2, 0xe0, 0xdd, 0x23, 0xc6,
	0xcd, 0xa2, 0x29, 0x71, 0x25, 0x6b, 0xd2, 0xfe, 0xbc, 0x63, 0x01, 0xe6, 0x42, 0xaa, 0x0f, 0xb6,
	0x20, 0xab, 0xa7, 0xe4, 0xb2, 0x3b, 0x86, 0x72, 0xdf, 0x02, 0x3c, 0xf9, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x68, 0x1c, 0x46, 0x71, 0xd0, 0x02, 0x00, 0x00,
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
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.CurrentRewardsAuctionEnd, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CurrentRewardsAuctionEnd):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	if len(m.NextRewards) > 0 {
		for iNdEx := len(m.NextRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NextRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.CurrentRewards) > 0 {
		for iNdEx := len(m.CurrentRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CurrentRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.CurrentRewards) > 0 {
		for _, e := range m.CurrentRewards {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NextRewards) > 0 {
		for _, e := range m.NextRewards {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CurrentRewardsAuctionEnd)
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
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentRewards", wireType)
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
			m.CurrentRewards = append(m.CurrentRewards, types.Coin{})
			if err := m.CurrentRewards[len(m.CurrentRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextRewards", wireType)
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
			m.NextRewards = append(m.NextRewards, types.Coin{})
			if err := m.NextRewards[len(m.NextRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentRewardsAuctionEnd", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.CurrentRewardsAuctionEnd, dAtA[iNdEx:postIndex]); err != nil {
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
