// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/auction/v1/events.proto

package auction

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// EventRewardsAuctionResult is emitted at the end of each auction that has at least one bidder.
type EventRewardsAuctionResult struct {
	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Bidder string `protobuf:"bytes,2,opt,name=bidder,proto3" json:"bidder,omitempty"`
}

func (m *EventRewardsAuctionResult) Reset()         { *m = EventRewardsAuctionResult{} }
func (m *EventRewardsAuctionResult) String() string { return proto.CompactTextString(m) }
func (*EventRewardsAuctionResult) ProtoMessage()    {}
func (*EventRewardsAuctionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5998c755af8caa1, []int{0}
}
func (m *EventRewardsAuctionResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRewardsAuctionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRewardsAuctionResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRewardsAuctionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRewardsAuctionResult.Merge(m, src)
}
func (m *EventRewardsAuctionResult) XXX_Size() int {
	return m.Size()
}
func (m *EventRewardsAuctionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRewardsAuctionResult.DiscardUnknown(m)
}

var xxx_messageInfo_EventRewardsAuctionResult proto.InternalMessageInfo

// EventFundRewardsAuction is emitted when sending rewards to auction module
type EventFundRewardsAuction struct {
	Assets []types.Coin `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets"`
}

func (m *EventFundRewardsAuction) Reset()         { *m = EventFundRewardsAuction{} }
func (m *EventFundRewardsAuction) String() string { return proto.CompactTextString(m) }
func (*EventFundRewardsAuction) ProtoMessage()    {}
func (*EventFundRewardsAuction) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5998c755af8caa1, []int{1}
}
func (m *EventFundRewardsAuction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFundRewardsAuction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFundRewardsAuction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventFundRewardsAuction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFundRewardsAuction.Merge(m, src)
}
func (m *EventFundRewardsAuction) XXX_Size() int {
	return m.Size()
}
func (m *EventFundRewardsAuction) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFundRewardsAuction.DiscardUnknown(m)
}

var xxx_messageInfo_EventFundRewardsAuction proto.InternalMessageInfo

// EventRewardsBid is emitted when the user bid the auction rewward.
type EventRewardsBid struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// the current auction ID being bid on. Fails if the ID is not an ID of the current auction.
	Id uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// amount of the bid in the bond base tokens (uumee)
	Amount types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
}

func (m *EventRewardsBid) Reset()         { *m = EventRewardsBid{} }
func (m *EventRewardsBid) String() string { return proto.CompactTextString(m) }
func (*EventRewardsBid) ProtoMessage()    {}
func (*EventRewardsBid) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5998c755af8caa1, []int{2}
}
func (m *EventRewardsBid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRewardsBid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRewardsBid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRewardsBid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRewardsBid.Merge(m, src)
}
func (m *EventRewardsBid) XXX_Size() int {
	return m.Size()
}
func (m *EventRewardsBid) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRewardsBid.DiscardUnknown(m)
}

var xxx_messageInfo_EventRewardsBid proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventRewardsAuctionResult)(nil), "umee.auction.v1.EventRewardsAuctionResult")
	proto.RegisterType((*EventFundRewardsAuction)(nil), "umee.auction.v1.EventFundRewardsAuction")
	proto.RegisterType((*EventRewardsBid)(nil), "umee.auction.v1.EventRewardsBid")
}

func init() { proto.RegisterFile("umee/auction/v1/events.proto", fileDescriptor_b5998c755af8caa1) }

var fileDescriptor_b5998c755af8caa1 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xec, 0x30,
	0x18, 0x85, 0x9b, 0xce, 0xa5, 0x30, 0xb9, 0xdc, 0x3b, 0x50, 0x06, 0xed, 0x0c, 0x12, 0xcb, 0xac,
	0xc6, 0xc5, 0x34, 0x56, 0x41, 0xd7, 0x53, 0x51, 0xf7, 0x75, 0x27, 0x88, 0xb4, 0x4d, 0xa8, 0x41,
	0x9b, 0x48, 0x92, 0x76, 0x7c, 0x0c, 0x1f, 0xc6, 0x87, 0x98, 0xe5, 0xe0, 0xca, 0x95, 0xe8, 0xcc,
	0x8b, 0x48, 0xda, 0x28, 0xba, 0x73, 0x97, 0xc3, 0x39, 0xfc, 0xdf, 0xf9, 0xf3, 0xc3, 0x9d, 0xba,
	0xa2, 0x14, 0x67, 0x75, 0xa1, 0x99, 0xe0, 0xb8, 0x89, 0x31, 0x6d, 0x28, 0xd7, 0x2a, 0xba, 0x97,
	0x42, 0x0b, 0x7f, 0x60, 0xdc, 0xc8, 0xba, 0x51, 0x13, 0x8f, 0x87, 0xa5, 0x28, 0x45, 0xeb, 0x61,
	0xf3, 0xea, 0x62, 0xe3, 0x51, 0x21, 0x54, 0x25, 0xd4, 0x75, 0x67, 0x74, 0xc2, 0x5a, 0xa8, 0x53,
	0x38, 0xcf, 0x14, 0xc5, 0x4d, 0x9c, 0x53, 0x9d, 0xc5, 0xb8, 0x10, 0x8c, 0x77, 0xfe, 0xe4, 0x0a,
	0x8e, 0x4e, 0x0d, 0x31, 0xa5, 0x8b, 0x4c, 0x12, 0x35, 0xef, 0x50, 0x29, 0x55, 0xf5, 0x9d, 0xf6,
	0xff, 0x43, 0x97, 0x91, 0x00, 0x84, 0x60, 0xfa, 0x2f, 0x75, 0x19, 0xf1, 0xf7, 0xa1, 0x97, 0x33,
	0x42, 0xa8, 0x0c, 0xdc, 0x10, 0x4c, 0xfb, 0x49, 0xf0, 0xfc, 0x34, 0x1b, 0x5a, 0xdc, 0x9c, 0x10,
	0x49, 0x95, 0xba, 0xd0, 0x92, 0xf1, 0x32, 0xb5, 0xb9, 0x49, 0x0a, 0xb7, 0xdb, 0xf1, 0x67, 0x35,
	0x27, 0x3f, 0x11, 0xfe, 0x31, 0xf4, 0x32, 0xa5, 0xa8, 0x56, 0x01, 0x08, 0x7b, 0xd3, 0xbf, 0x07,
	0xa3, 0xc8, 0x4e, 0x32, 0x55, 0x23, 0x5b, 0x35, 0x3a, 0x11, 0x8c, 0x27, 0x7f, 0x96, 0xaf, 0xbb,
	0x4e, 0x6a, 0xe3, 0x13, 0x09, 0x07, 0xdf, 0x2b, 0x27, 0x8c, 0xf8, 0x5b, 0xd0, 0x53, 0x94, 0x9b,
	0x62, 0xa6, 0x6c, 0x3f, 0xb5, 0xca, 0x2e, 0xe0, 0x7e, 0x2d, 0x60, 0x98, 0x95, 0xa8, 0xb9, 0x0e,
	0x7a, 0x21, 0xf8, 0x1d, 0xb3, 0x8d, 0x27, 0xe7, 0xcb, 0x77, 0xe4, 0x2c, 0xd7, 0x08, 0xac, 0xd6,
	0x08, 0xbc, 0xad, 0x11, 0x78, 0xdc, 0x20, 0x67, 0xb5, 0x41, 0xce, 0xcb, 0x06, 0x39, 0x97, 0x7b,
	0x25, 0xd3, 0x37, 0x75, 0x1e, 0x15, 0xa2, 0xc2, 0xe6, 0x62, 0x33, 0x4e, 0xf5, 0x42, 0xc8, 0xdb,
	0x56, 0xe0, 0xe6, 0x08, 0x3f, 0x7c, 0x5e, 0x38, 0xf7, 0xda, 0x6f, 0x3f, 0xfc, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xef, 0x82, 0xa9, 0x6a, 0xf8, 0x01, 0x00, 0x00,
}

func (m *EventRewardsAuctionResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRewardsAuctionResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRewardsAuctionResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventFundRewardsAuction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFundRewardsAuction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventFundRewardsAuction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for iNdEx := len(m.Assets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Assets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *EventRewardsBid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRewardsBid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRewardsBid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Id != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventRewardsAuctionResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventFundRewardsAuction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for _, e := range m.Assets {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventRewardsBid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovEvents(uint64(m.Id))
	}
	l = m.Amount.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventRewardsAuctionResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRewardsAuctionResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRewardsAuctionResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventFundRewardsAuction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventFundRewardsAuction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFundRewardsAuction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Assets = append(m.Assets, types.Coin{})
			if err := m.Assets[len(m.Assets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventRewardsBid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRewardsBid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRewardsBid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
