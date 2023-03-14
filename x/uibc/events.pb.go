// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/uibc/v1/events.proto

package uibc

import (
	fmt "fmt"
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

// EventBadRevert is emitted on failure of ibc-transfer quota.
type EventBadRevert struct {
	// failure event type
	FailureType string `protobuf:"bytes,1,opt,name=failure_type,json=failureType,proto3" json:"failure_type,omitempty"`
	// ibc packet data
	Packet string `protobuf:"bytes,2,opt,name=packet,proto3" json:"packet,omitempty"`
}

func (m *EventBadRevert) Reset()         { *m = EventBadRevert{} }
func (m *EventBadRevert) String() string { return proto.CompactTextString(m) }
func (*EventBadRevert) ProtoMessage()    {}
func (*EventBadRevert) Descriptor() ([]byte, []int) {
	return fileDescriptor_c64e60b79cebf048, []int{0}
}
func (m *EventBadRevert) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBadRevert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBadRevert.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBadRevert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBadRevert.Merge(m, src)
}
func (m *EventBadRevert) XXX_Size() int {
	return m.Size()
}
func (m *EventBadRevert) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBadRevert.DiscardUnknown(m)
}

var xxx_messageInfo_EventBadRevert proto.InternalMessageInfo

// EventIBCTransferStatus is emitted on quota tracking pause status change.
type EventIBCTransferStatus struct {
	Status IBCTransferStatus `protobuf:"varint,1,opt,name=status,proto3,enum=umee.uibc.v1.IBCTransferStatus" json:"status,omitempty"`
}

func (m *EventIBCTransferStatus) Reset()         { *m = EventIBCTransferStatus{} }
func (m *EventIBCTransferStatus) String() string { return proto.CompactTextString(m) }
func (*EventIBCTransferStatus) ProtoMessage()    {}
func (*EventIBCTransferStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c64e60b79cebf048, []int{1}
}
func (m *EventIBCTransferStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventIBCTransferStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventIBCTransferStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventIBCTransferStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventIBCTransferStatus.Merge(m, src)
}
func (m *EventIBCTransferStatus) XXX_Size() int {
	return m.Size()
}
func (m *EventIBCTransferStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EventIBCTransferStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EventIBCTransferStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventBadRevert)(nil), "umee.uibc.v1.EventBadRevert")
	proto.RegisterType((*EventIBCTransferStatus)(nil), "umee.uibc.v1.EventIBCTransferStatus")
}

func init() { proto.RegisterFile("umee/uibc/v1/events.proto", fileDescriptor_c64e60b79cebf048) }

var fileDescriptor_c64e60b79cebf048 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0xcd, 0x4d, 0x4d,
	0xd5, 0x2f, 0xcd, 0x4c, 0x4a, 0xd6, 0x2f, 0x33, 0xd4, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0x49, 0xe9, 0x81, 0xa4, 0xf4, 0xca, 0x0c, 0xa5,
	0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x12, 0xfa, 0x20, 0x16, 0x44, 0x8d, 0x94, 0x04, 0x8a, 0xf6,
	0xc2, 0xd2, 0xfc, 0x92, 0x44, 0x88, 0x8c, 0x92, 0x37, 0x17, 0x9f, 0x2b, 0xc8, 0x34, 0xa7, 0xc4,
	0x94, 0xa0, 0xd4, 0xb2, 0xd4, 0xa2, 0x12, 0x21, 0x45, 0x2e, 0x9e, 0xb4, 0xc4, 0xcc, 0x9c, 0xd2,
	0xa2, 0xd4, 0xf8, 0x92, 0xca, 0x82, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x6e, 0xa8,
	0x58, 0x48, 0x65, 0x41, 0xaa, 0x90, 0x18, 0x17, 0x5b, 0x41, 0x62, 0x72, 0x76, 0x6a, 0x89, 0x04,
	0x13, 0x58, 0x12, 0xca, 0x53, 0x0a, 0xe4, 0x12, 0x03, 0x1b, 0xe6, 0xe9, 0xe4, 0x1c, 0x52, 0x94,
	0x98, 0x57, 0x9c, 0x96, 0x5a, 0x14, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x2c, 0x64, 0xce, 0xc5, 0x56,
	0x0c, 0x66, 0x81, 0x8d, 0xe3, 0x33, 0x92, 0xd7, 0x43, 0x76, 0xb5, 0x1e, 0x86, 0x86, 0x20, 0xa8,
	0x72, 0x27, 0x97, 0x13, 0x0f, 0xe5, 0x18, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e,
	0x21, 0x4a, 0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x64, 0xa0,
	0x6e, 0x5e, 0x6a, 0x49, 0x79, 0x7e, 0x51, 0x36, 0x98, 0xa3, 0x5f, 0x66, 0xa2, 0x5f, 0x01, 0xf6,
	0x74, 0x12, 0x1b, 0xd8, 0xb3, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xc0, 0x80, 0xb6,
	0x47, 0x01, 0x00, 0x00,
}

func (m *EventBadRevert) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBadRevert) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBadRevert) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Packet) > 0 {
		i -= len(m.Packet)
		copy(dAtA[i:], m.Packet)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Packet)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FailureType) > 0 {
		i -= len(m.FailureType)
		copy(dAtA[i:], m.FailureType)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.FailureType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventIBCTransferStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventIBCTransferStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventIBCTransferStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
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
func (m *EventBadRevert) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FailureType)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Packet)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventIBCTransferStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventBadRevert) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventBadRevert: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBadRevert: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureType", wireType)
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
			m.FailureType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Packet", wireType)
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
			m.Packet = string(dAtA[iNdEx:postIndex])
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
func (m *EventIBCTransferStatus) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventIBCTransferStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventIBCTransferStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= IBCTransferStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
