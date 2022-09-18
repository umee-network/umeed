// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/oracle/v1/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// EventDelegateFeedConsent is emitted on Msg/DelegateFeedConsent
type EventDelegateFeedConsent struct {
	// Operator bech32 address who delegates his feed consent
	Operator string `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	// Delegate bech32 address
	Delegate string `protobuf:"bytes,2,opt,name=delegate,proto3" json:"delegate,omitempty"`
}

func (m *EventDelegateFeedConsent) Reset()         { *m = EventDelegateFeedConsent{} }
func (m *EventDelegateFeedConsent) String() string { return proto.CompactTextString(m) }
func (*EventDelegateFeedConsent) ProtoMessage()    {}
func (*EventDelegateFeedConsent) Descriptor() ([]byte, []int) {
	return fileDescriptor_6380f28dac582975, []int{0}
}
func (m *EventDelegateFeedConsent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventDelegateFeedConsent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventDelegateFeedConsent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventDelegateFeedConsent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDelegateFeedConsent.Merge(m, src)
}
func (m *EventDelegateFeedConsent) XXX_Size() int {
	return m.Size()
}
func (m *EventDelegateFeedConsent) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDelegateFeedConsent.DiscardUnknown(m)
}

var xxx_messageInfo_EventDelegateFeedConsent proto.InternalMessageInfo

// EventSetFxRate is emitted on exchange rate update
type EventSetFxRate struct {
	// uToken denom
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// Exchange rate (based to USD)
	Rate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate"`
}

func (m *EventSetFxRate) Reset()         { *m = EventSetFxRate{} }
func (m *EventSetFxRate) String() string { return proto.CompactTextString(m) }
func (*EventSetFxRate) ProtoMessage()    {}
func (*EventSetFxRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_6380f28dac582975, []int{1}
}
func (m *EventSetFxRate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetFxRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetFxRate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetFxRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetFxRate.Merge(m, src)
}
func (m *EventSetFxRate) XXX_Size() int {
	return m.Size()
}
func (m *EventSetFxRate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetFxRate.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetFxRate proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventDelegateFeedConsent)(nil), "umee.oracle.v1.EventDelegateFeedConsent")
	proto.RegisterType((*EventSetFxRate)(nil), "umee.oracle.v1.EventSetFxRate")
}

func init() { proto.RegisterFile("umee/oracle/v1/events.proto", fileDescriptor_6380f28dac582975) }

var fileDescriptor_6380f28dac582975 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x13, 0x04, 0x08, 0x3c, 0x74, 0x88, 0x3a, 0x84, 0x22, 0xb9, 0xa8, 0x03, 0x62, 0x49,
	0xac, 0xaa, 0x8c, 0x2c, 0x94, 0xd2, 0x89, 0x01, 0xa5, 0x1b, 0x0b, 0x4a, 0xe3, 0x53, 0xa8, 0xda,
	0xf8, 0x2a, 0xdb, 0x0d, 0xe5, 0x05, 0x98, 0x79, 0x98, 0x3e, 0x44, 0xc7, 0xaa, 0x13, 0x62, 0xa8,
	0xa0, 0x79, 0x11, 0x14, 0x3b, 0x05, 0x36, 0x26, 0xfb, 0xee, 0xff, 0xef, 0xbb, 0xf3, 0x99, 0x9c,
	0xce, 0x32, 0x00, 0x86, 0x32, 0x4e, 0x26, 0xc0, 0xf2, 0x36, 0x83, 0x1c, 0x84, 0x56, 0xe1, 0x54,
	0xa2, 0x46, 0xaf, 0x56, 0x8a, 0xa1, 0x15, 0xc3, 0xbc, 0xdd, 0x38, 0x49, 0x50, 0x65, 0xa8, 0x1e,
	0x8d, 0xca, 0x6c, 0x60, 0xad, 0x8d, 0x7a, 0x8a, 0x29, 0xda, 0x7c, 0x79, 0xb3, 0xd9, 0xd6, 0xab,
	0x4b, 0xfc, 0xdb, 0x92, 0xd8, 0x83, 0x09, 0xa4, 0xb1, 0x86, 0x3e, 0x00, 0xbf, 0x41, 0xa1, 0x40,
	0x68, 0xef, 0x92, 0x1c, 0xe1, 0x14, 0x64, 0xac, 0x51, 0xfa, 0xee, 0x99, 0x7b, 0x71, 0xdc, 0xf5,
	0xd7, 0x8b, 0xa0, 0x5e, 0x61, 0xaf, 0x39, 0x97, 0xa0, 0xd4, 0x40, 0xcb, 0x91, 0x48, 0xa3, 0x1f,
	0x67, 0x59, 0xc5, 0x2b, 0x98, 0xbf, 0xf7, 0x5f, 0xd5, 0xce, 0xd9, 0x9a, 0x93, 0x9a, 0x99, 0x63,
	0x00, 0xba, 0x3f, 0x8f, 0x62, 0x0d, 0x5e, 0x9d, 0x1c, 0x70, 0x10, 0x98, 0xd9, 0xd6, 0x91, 0x0d,
	0xbc, 0x7b, 0xb2, 0x2f, 0x7f, 0xc9, 0x57, 0xcb, 0x4d, 0xd3, 0xf9, 0xd8, 0x34, 0xcf, 0xd3, 0x91,
	0x7e, 0x9a, 0x0d, 0xc3, 0x04, 0xb3, 0xea, 0xd5, 0xd5, 0x11, 0x28, 0x3e, 0x66, 0xfa, 0x65, 0x0a,
	0x2a, 0xec, 0x41, 0xb2, 0x5e, 0x04, 0xa4, 0x9a, 0xa3, 0x07, 0x49, 0x64, 0x48, 0xdd, 0xbb, 0xe5,
	0x17, 0x75, 0x96, 0x5b, 0xea, 0xae, 0xb6, 0xd4, 0xfd, 0xdc, 0x52, 0xf7, 0xad, 0xa0, 0xce, 0xaa,
	0xa0, 0xce, 0x7b, 0x41, 0x9d, 0x87, 0xf0, 0x0f, 0xb9, 0x5c, 0x76, 0x20, 0x40, 0x3f, 0xa3, 0x1c,
	0x9b, 0x80, 0xe5, 0x1d, 0x36, 0xdf, 0xfd, 0x8d, 0xe9, 0x32, 0x3c, 0x34, 0x7b, 0xed, 0x7c, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xfd, 0xea, 0x39, 0x53, 0xb7, 0x01, 0x00, 0x00,
}

func (m *EventDelegateFeedConsent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventDelegateFeedConsent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventDelegateFeedConsent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Delegate) > 0 {
		i -= len(m.Delegate)
		copy(dAtA[i:], m.Delegate)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Delegate)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSetFxRate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetFxRate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetFxRate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
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
func (m *EventDelegateFeedConsent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Delegate)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventSetFxRate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Rate.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventDelegateFeedConsent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventDelegateFeedConsent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventDelegateFeedConsent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
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
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegate", wireType)
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
			m.Delegate = string(dAtA[iNdEx:postIndex])
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
func (m *EventSetFxRate) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventSetFxRate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetFxRate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
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
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
