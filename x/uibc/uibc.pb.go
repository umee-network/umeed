// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/uibc/v1/uibc.proto

package uibc

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

type ICS20Memo struct {
	// messages is a list of `sdk.Msg`s that will be executed when handling ICS20 transfer.
	Messages []*types.Any `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (m *ICS20Memo) Reset()         { *m = ICS20Memo{} }
func (m *ICS20Memo) String() string { return proto.CompactTextString(m) }
func (*ICS20Memo) ProtoMessage()    {}
func (*ICS20Memo) Descriptor() ([]byte, []int) {
	return fileDescriptor_963b2b690b6cd9dd, []int{0}
}
func (m *ICS20Memo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ICS20Memo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ICS20Memo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ICS20Memo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ICS20Memo.Merge(m, src)
}
func (m *ICS20Memo) XXX_Size() int {
	return m.Size()
}
func (m *ICS20Memo) XXX_DiscardUnknown() {
	xxx_messageInfo_ICS20Memo.DiscardUnknown(m)
}

var xxx_messageInfo_ICS20Memo proto.InternalMessageInfo

func (*ICS20Memo) XXX_MessageName() string {
	return "umee.uibc.v1.ICS20Memo"
}

type DecCoinSymbol struct {
	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount Dec    `protobuf:"bytes,2,opt,name=amount,proto3,customtype=Dec" json:"amount"`
	// token symbol name
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (m *DecCoinSymbol) Reset()         { *m = DecCoinSymbol{} }
func (m *DecCoinSymbol) String() string { return proto.CompactTextString(m) }
func (*DecCoinSymbol) ProtoMessage()    {}
func (*DecCoinSymbol) Descriptor() ([]byte, []int) {
	return fileDescriptor_963b2b690b6cd9dd, []int{1}
}
func (m *DecCoinSymbol) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DecCoinSymbol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DecCoinSymbol.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DecCoinSymbol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecCoinSymbol.Merge(m, src)
}
func (m *DecCoinSymbol) XXX_Size() int {
	return m.Size()
}
func (m *DecCoinSymbol) XXX_DiscardUnknown() {
	xxx_messageInfo_DecCoinSymbol.DiscardUnknown(m)
}

var xxx_messageInfo_DecCoinSymbol proto.InternalMessageInfo

func (*DecCoinSymbol) XXX_MessageName() string {
	return "umee.uibc.v1.DecCoinSymbol"
}
func init() {
	proto.RegisterType((*ICS20Memo)(nil), "umee.uibc.v1.ICS20Memo")
	proto.RegisterType((*DecCoinSymbol)(nil), "umee.uibc.v1.DecCoinSymbol")
}

func init() { proto.RegisterFile("umee/uibc/v1/uibc.proto", fileDescriptor_963b2b690b6cd9dd) }

var fileDescriptor_963b2b690b6cd9dd = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x31, 0x4f, 0x02, 0x31,
	0x18, 0x86, 0xaf, 0x12, 0x89, 0x54, 0x5d, 0x2e, 0x44, 0x0f, 0x86, 0x42, 0x18, 0x0c, 0x0b, 0x2d,
	0x60, 0xe2, 0xe6, 0x20, 0x30, 0xe8, 0xe0, 0x02, 0x9b, 0x8b, 0xe1, 0x6a, 0xad, 0x44, 0xda, 0x8f,
	0xd0, 0x3b, 0xf4, 0xfe, 0x85, 0x3f, 0xc6, 0x1f, 0x71, 0x23, 0x71, 0x32, 0x0e, 0x44, 0xef, 0xfe,
	0x88, 0xb9, 0xf6, 0x74, 0x6a, 0x9f, 0xf7, 0x7d, 0xbf, 0x2f, 0xf9, 0x5e, 0x7c, 0x1a, 0x2b, 0x21,
	0x58, 0xbc, 0x08, 0x39, 0xdb, 0x0c, 0xec, 0x4b, 0x57, 0x6b, 0x88, 0xc0, 0x3f, 0x2a, 0x0c, 0x6a,
	0x85, 0xcd, 0xa0, 0xd9, 0xe0, 0x60, 0x14, 0x98, 0x7b, 0xeb, 0x31, 0x07, 0x2e, 0xd8, 0xac, 0x4b,
	0x90, 0xe0, 0xf4, 0xe2, 0x57, 0xaa, 0x0d, 0x09, 0x20, 0x97, 0x82, 0x59, 0x0a, 0xe3, 0x47, 0x36,
	0xd7, 0x89, 0xb3, 0x3a, 0x97, 0xb8, 0x76, 0x33, 0x9e, 0x0d, 0xfb, 0xb7, 0x42, 0x81, 0xdf, 0xc7,
	0x07, 0x4a, 0x18, 0x33, 0x97, 0xc2, 0x04, 0xa8, 0x5d, 0xe9, 0x1e, 0x0e, 0xeb, 0xd4, 0x8d, 0xd2,
	0xbf, 0x51, 0x7a, 0xa5, 0x93, 0xe9, 0x7f, 0xaa, 0xb3, 0xc2, 0xc7, 0x13, 0xc1, 0xc7, 0xb0, 0xd0,
	0xb3, 0x44, 0x85, 0xb0, 0xf4, 0xeb, 0x78, 0xff, 0x41, 0x68, 0x50, 0x01, 0x6a, 0xa3, 0x6e, 0x6d,
	0xea, 0xc0, 0x1f, 0xe0, 0xea, 0x5c, 0x41, 0xac, 0xa3, 0x60, 0xaf, 0x90, 0x47, 0x8d, 0x74, 0xd7,
	0xf2, 0xbe, 0x76, 0xad, 0xca, 0x44, 0xf0, 0x8f, 0xf7, 0x1e, 0x2e, 0x6f, 0x98, 0x08, 0x3e, 0x2d,
	0x83, 0xfe, 0x09, 0xae, 0x1a, 0xbb, 0x32, 0xa8, 0xd8, 0x4d, 0x25, 0x8d, 0xae, 0xd3, 0x1f, 0xe2,
	0xa5, 0x19, 0x41, 0xdb, 0x8c, 0xa0, 0xef, 0x8c, 0xa0, 0xb7, 0x9c, 0x78, 0x69, 0x4e, 0xd0, 0x36,
	0x27, 0xde, 0x67, 0x4e, 0xbc, 0xbb, 0x33, 0xb9, 0x88, 0x9e, 0xe2, 0x90, 0x72, 0x50, 0xac, 0xe8,
	0xad, 0xa7, 0x45, 0xf4, 0x02, 0xeb, 0x67, 0x0b, 0x6c, 0x73, 0xc1, 0x5e, 0x6d, 0xb5, 0x61, 0xd5,
	0xde, 0x74, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x8d, 0x4e, 0xe3, 0x76, 0x01, 0x00, 0x00,
}

func (m *ICS20Memo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ICS20Memo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ICS20Memo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUibc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DecCoinSymbol) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DecCoinSymbol) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DecCoinSymbol) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintUibc(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintUibc(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintUibc(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUibc(dAtA []byte, offset int, v uint64) int {
	offset -= sovUibc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ICS20Memo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovUibc(uint64(l))
		}
	}
	return n
}

func (m *DecCoinSymbol) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovUibc(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovUibc(uint64(l))
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovUibc(uint64(l))
	}
	return n
}

func sovUibc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUibc(x uint64) (n int) {
	return sovUibc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ICS20Memo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUibc
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
			return fmt.Errorf("proto: ICS20Memo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ICS20Memo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUibc
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
				return ErrInvalidLengthUibc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUibc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, &types.Any{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUibc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUibc
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
func (m *DecCoinSymbol) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUibc
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
			return fmt.Errorf("proto: DecCoinSymbol: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DecCoinSymbol: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUibc
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
				return ErrInvalidLengthUibc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUibc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUibc
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
				return ErrInvalidLengthUibc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUibc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUibc
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
				return ErrInvalidLengthUibc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUibc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUibc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUibc
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
func skipUibc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUibc
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
					return 0, ErrIntOverflowUibc
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
					return 0, ErrIntOverflowUibc
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
				return 0, ErrInvalidLengthUibc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUibc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUibc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUibc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUibc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUibc = fmt.Errorf("proto: unexpected end of group")
)
