// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/ugov/v1/events.proto

package ugov

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// EventMinTxFees is emitted when MsgGovUpdateMinFees is correctly executed.
type EventMinTxFees struct {
	MinTxFees []types.DecCoin `protobuf:"bytes,1,rep,name=min_tx_fees,json=minTxFees,proto3" json:"min_tx_fees"`
}

func (m *EventMinTxFees) Reset()         { *m = EventMinTxFees{} }
func (m *EventMinTxFees) String() string { return proto.CompactTextString(m) }
func (*EventMinTxFees) ProtoMessage()    {}
func (*EventMinTxFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_0885cdf0808da4ea, []int{0}
}
func (m *EventMinTxFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventMinTxFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventMinTxFees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventMinTxFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMinTxFees.Merge(m, src)
}
func (m *EventMinTxFees) XXX_Size() int {
	return m.Size()
}
func (m *EventMinTxFees) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMinTxFees.DiscardUnknown(m)
}

var xxx_messageInfo_EventMinTxFees proto.InternalMessageInfo

// EventEmergencyGroup is emitted when MsgGovSetEmergencyGroup is correctly executed.
type EventEmergencyGroup struct {
	EmergencyGroup string `protobuf:"bytes,1,opt,name=emergency_group,json=emergencyGroup,proto3" json:"emergency_group,omitempty"`
}

func (m *EventEmergencyGroup) Reset()         { *m = EventEmergencyGroup{} }
func (m *EventEmergencyGroup) String() string { return proto.CompactTextString(m) }
func (*EventEmergencyGroup) ProtoMessage()    {}
func (*EventEmergencyGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_0885cdf0808da4ea, []int{1}
}
func (m *EventEmergencyGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventEmergencyGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventEmergencyGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventEmergencyGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEmergencyGroup.Merge(m, src)
}
func (m *EventEmergencyGroup) XXX_Size() int {
	return m.Size()
}
func (m *EventEmergencyGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEmergencyGroup.DiscardUnknown(m)
}

var xxx_messageInfo_EventEmergencyGroup proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventMinTxFees)(nil), "umee.ugov.v1.EventMinTxFees")
	proto.RegisterType((*EventEmergencyGroup)(nil), "umee.ugov.v1.EventEmergencyGroup")
}

func init() { proto.RegisterFile("umee/ugov/v1/events.proto", fileDescriptor_0885cdf0808da4ea) }

var fileDescriptor_0885cdf0808da4ea = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4e, 0xf2, 0x40,
	0x14, 0x85, 0xdb, 0xfc, 0x7f, 0x4c, 0x28, 0x06, 0x13, 0x64, 0x01, 0xc4, 0x8c, 0x84, 0x85, 0x61,
	0xc3, 0x4c, 0xaa, 0xf1, 0x01, 0xa8, 0xa0, 0x2b, 0x37, 0xc8, 0xc2, 0xb8, 0x69, 0x68, 0xb9, 0x8e,
	0x8d, 0x99, 0xb9, 0x64, 0x66, 0x3a, 0xe2, 0x5b, 0xf8, 0x30, 0x3e, 0x04, 0x4b, 0xe2, 0xca, 0x95,
	0x51, 0x78, 0x11, 0x33, 0x6d, 0x59, 0xb8, 0x9b, 0x73, 0xce, 0xe4, 0x7c, 0xb9, 0x27, 0xe8, 0xe4,
	0x02, 0x80, 0xe5, 0x1c, 0x2d, 0xb3, 0x21, 0x03, 0x0b, 0xd2, 0x68, 0xba, 0x54, 0x68, 0xb0, 0x79,
	0xe8, 0x22, 0xea, 0x22, 0x6a, 0xc3, 0x2e, 0x49, 0x51, 0x0b, 0xd4, 0x2c, 0x99, 0x6b, 0x60, 0x36,
	0x4c, 0xc0, 0xcc, 0x43, 0x96, 0x62, 0x26, 0xcb, 0xdf, 0xdd, 0x4e, 0x99, 0xc7, 0x85, 0x62, 0xa5,
	0xa8, 0xa2, 0x16, 0x47, 0x8e, 0xa5, 0xef, 0x5e, 0xa5, 0xdb, 0x9f, 0x05, 0x8d, 0x89, 0xc3, 0xdd,
	0x66, 0x72, 0xb6, 0xba, 0x06, 0xd0, 0xcd, 0x28, 0xa8, 0x8b, 0x4c, 0xc6, 0x66, 0x15, 0x3f, 0x02,
	0xe8, 0xb6, 0xdf, 0xfb, 0x37, 0xa8, 0x9f, 0x9f, 0xd0, 0xaa, 0xcb, 0x81, 0x69, 0x05, 0xa6, 0x63,
	0x48, 0xaf, 0x30, 0x93, 0xd1, 0xff, 0xf5, 0xd7, 0xa9, 0x37, 0xad, 0x89, 0x7d, 0x47, 0xff, 0x3e,
	0x38, 0x2e, 0x5a, 0x27, 0x02, 0x14, 0x07, 0x99, 0xbe, 0xde, 0x28, 0xcc, 0x97, 0xcd, 0x51, 0x70,
	0x04, 0x7b, 0x27, 0xe6, 0xce, 0x6a, 0xfb, 0x3d, 0x7f, 0x50, 0x8b, 0xda, 0x1f, 0xef, 0xc3, 0x56,
	0x45, 0x18, 0x2d, 0x16, 0x0a, 0xb4, 0xbe, 0x33, 0x2a, 0x93, 0x7c, 0xda, 0x80, 0x3f, 0x15, 0xd1,
	0x78, 0xfd, 0x43, 0xbc, 0xf5, 0x96, 0xf8, 0x9b, 0x2d, 0xf1, 0xbf, 0xb7, 0xc4, 0x7f, 0xdb, 0x11,
	0x6f, 0xb3, 0x23, 0xde, 0xe7, 0x8e, 0x78, 0x0f, 0x67, 0x3c, 0x33, 0x4f, 0x79, 0x42, 0x53, 0x14,
	0xcc, 0xed, 0x36, 0x94, 0x60, 0x5e, 0x50, 0x3d, 0x17, 0x82, 0xd9, 0x4b, 0xb6, 0x2a, 0x46, 0x4e,
	0x0e, 0x8a, 0xe3, 0x2f, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x41, 0x70, 0xf3, 0x78, 0x01,
	0x00, 0x00,
}

func (m *EventMinTxFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventMinTxFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventMinTxFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinTxFees) > 0 {
		for iNdEx := len(m.MinTxFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinTxFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *EventEmergencyGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventEmergencyGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventEmergencyGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EmergencyGroup) > 0 {
		i -= len(m.EmergencyGroup)
		copy(dAtA[i:], m.EmergencyGroup)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.EmergencyGroup)))
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
func (m *EventMinTxFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MinTxFees) > 0 {
		for _, e := range m.MinTxFees {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventEmergencyGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EmergencyGroup)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventMinTxFees) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventMinTxFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventMinTxFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTxFees", wireType)
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
			m.MinTxFees = append(m.MinTxFees, types.DecCoin{})
			if err := m.MinTxFees[len(m.MinTxFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *EventEmergencyGroup) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventEmergencyGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventEmergencyGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmergencyGroup", wireType)
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
			m.EmergencyGroup = string(dAtA[iNdEx:postIndex])
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
