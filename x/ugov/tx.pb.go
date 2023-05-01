// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/ugov/v1/tx.proto

package ugov

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgGovUpdateMinFees is a request type for the Msg/GovUpdateMinFees.
type MsgGovUpdateMinFees struct {
	// authority is the address of the governance account.
	Authority   string                                   `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	Title       string                                   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	MinTxFees   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=min_tx_fees,json=minTxFees,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_tx_fees"`
}

func (m *MsgGovUpdateMinFees) Reset()      { *m = MsgGovUpdateMinFees{} }
func (*MsgGovUpdateMinFees) ProtoMessage() {}
func (*MsgGovUpdateMinFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffc07de1c6ee91b, []int{0}
}
func (m *MsgGovUpdateMinFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovUpdateMinFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovUpdateMinFees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovUpdateMinFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovUpdateMinFees.Merge(m, src)
}
func (m *MsgGovUpdateMinFees) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovUpdateMinFees) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovUpdateMinFees.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovUpdateMinFees proto.InternalMessageInfo

// MsgGovUpdateMinFeesResponse is a response type for the Msg/GovUpdateMinFees.
type MsgGovUpdateMinFeesResponse struct {
}

func (m *MsgGovUpdateMinFeesResponse) Reset()         { *m = MsgGovUpdateMinFeesResponse{} }
func (m *MsgGovUpdateMinFeesResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGovUpdateMinFeesResponse) ProtoMessage()    {}
func (*MsgGovUpdateMinFeesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ffc07de1c6ee91b, []int{1}
}
func (m *MsgGovUpdateMinFeesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGovUpdateMinFeesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGovUpdateMinFeesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGovUpdateMinFeesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGovUpdateMinFeesResponse.Merge(m, src)
}
func (m *MsgGovUpdateMinFeesResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGovUpdateMinFeesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGovUpdateMinFeesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGovUpdateMinFeesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgGovUpdateMinFees)(nil), "umee.ugov.v1.MsgGovUpdateMinFees")
	proto.RegisterType((*MsgGovUpdateMinFeesResponse)(nil), "umee.ugov.v1.MsgGovUpdateMinFeesResponse")
}

func init() { proto.RegisterFile("umee/ugov/v1/tx.proto", fileDescriptor_9ffc07de1c6ee91b) }

var fileDescriptor_9ffc07de1c6ee91b = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3d, 0x8f, 0xd3, 0x40,
	0x10, 0xb5, 0x2f, 0x80, 0x94, 0x0d, 0x95, 0x09, 0xe0, 0x0b, 0xc2, 0x09, 0x57, 0xa0, 0x80, 0x94,
	0x5d, 0x72, 0x20, 0x0a, 0x3a, 0x02, 0xe2, 0xaa, 0x34, 0x01, 0x9a, 0x6b, 0x22, 0x27, 0x1e, 0xf6,
	0x56, 0xc1, 0x3b, 0x96, 0x67, 0x63, 0x7c, 0x2d, 0x15, 0xa2, 0xa2, 0xa4, 0xbc, 0x9a, 0x8a, 0x82,
	0x1f, 0x91, 0xf2, 0x44, 0x45, 0xc5, 0x47, 0x52, 0xf0, 0x37, 0xd0, 0xda, 0x7b, 0xba, 0x14, 0x27,
	0x51, 0xed, 0xbe, 0xf7, 0xe6, 0x63, 0xe7, 0xcd, 0xb2, 0xeb, 0xcb, 0x14, 0x40, 0x2c, 0x25, 0x16,
	0xa2, 0x18, 0x0a, 0x53, 0xf2, 0x2c, 0x47, 0x83, 0xc1, 0x55, 0x4b, 0x73, 0x4b, 0xf3, 0x62, 0xd8,
	0x89, 0xe6, 0x48, 0x29, 0x92, 0x98, 0xc5, 0x04, 0xa2, 0x18, 0xce, 0xc0, 0xc4, 0x43, 0x31, 0x47,
	0xa5, 0xeb, 0xe8, 0xce, 0x4d, 0xa7, 0xa7, 0x24, 0x6d, 0x95, 0x94, 0xa4, 0x13, 0x76, 0x6b, 0x61,
	0x5a, 0x21, 0x51, 0x03, 0x27, 0xb5, 0x25, 0x4a, 0xac, 0x79, 0x7b, 0xab, 0xd9, 0xbd, 0x8f, 0x3b,
	0xec, 0xda, 0x98, 0xe4, 0x01, 0x16, 0xaf, 0xb3, 0x24, 0x36, 0x30, 0x56, 0xfa, 0x05, 0x00, 0x05,
	0x8f, 0x59, 0x33, 0x5e, 0x9a, 0x23, 0xcc, 0x95, 0x39, 0x0e, 0xfd, 0x9e, 0xdf, 0x6f, 0x8e, 0xc2,
	0xef, 0xdf, 0x06, 0x6d, 0x57, 0xf2, 0x69, 0x92, 0xe4, 0x40, 0xf4, 0xd2, 0xe4, 0x4a, 0xcb, 0xc9,
	0x79, 0x68, 0xd0, 0x66, 0x97, 0x8d, 0x32, 0x6f, 0x21, 0xdc, 0xb1, 0x39, 0x93, 0x1a, 0x04, 0x3d,
	0xd6, 0x4a, 0x80, 0xe6, 0xb9, 0xca, 0x8c, 0x42, 0x1d, 0x36, 0x2a, 0x6d, 0x9b, 0x0a, 0x16, 0xac,
	0x95, 0x2a, 0x3d, 0x35, 0xe5, 0xf4, 0x0d, 0x00, 0x85, 0x97, 0x7a, 0x8d, 0x7e, 0x6b, 0x7f, 0x97,
	0xbb, 0x76, 0xd6, 0x07, 0xee, 0x7c, 0xe0, 0xcf, 0x50, 0xe9, 0xd1, 0x83, 0xd5, 0xcf, 0xae, 0xf7,
	0xe5, 0x57, 0xb7, 0x2f, 0x95, 0x39, 0x5a, 0xce, 0xf8, 0x1c, 0x53, 0x37, 0xae, 0x3b, 0x06, 0x94,
	0x2c, 0x84, 0x39, 0xce, 0x80, 0xaa, 0x04, 0x9a, 0x34, 0x53, 0xa5, 0x5f, 0x95, 0x76, 0xb8, 0x27,
	0x37, 0x3e, 0x9c, 0x74, 0xbd, 0xcf, 0x27, 0x5d, 0xef, 0xfd, 0xdf, 0xaf, 0xf7, 0xcf, 0x1f, 0xbf,
	0x77, 0x9b, 0xdd, 0xba, 0xc0, 0x8b, 0x09, 0x50, 0x86, 0x9a, 0x60, 0x3f, 0x66, 0x8d, 0x31, 0xc9,
	0xe0, 0x90, 0xb1, 0x03, 0x2c, 0xce, 0x8c, 0xba, 0xc3, 0xb7, 0x37, 0xc7, 0x2f, 0xc8, 0xef, 0xdc,
	0xfb, 0x6f, 0xc8, 0x59, 0x8b, 0xd1, 0xf3, 0xd5, 0x9f, 0xc8, 0x5b, 0xad, 0x23, 0xff, 0x74, 0x1d,
	0xf9, 0xbf, 0xd7, 0x91, 0xff, 0x69, 0x13, 0x79, 0xa7, 0x9b, 0xc8, 0xfb, 0xb1, 0x89, 0xbc, 0xc3,
	0xbb, 0x5b, 0xc3, 0xda, 0x92, 0x03, 0x0d, 0xe6, 0x1d, 0xe6, 0x8b, 0x0a, 0x88, 0xe2, 0x91, 0x28,
	0xab, 0x8f, 0x35, 0xbb, 0x52, 0xed, 0xf6, 0xe1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x2b,
	0xa5, 0x1c, 0x6c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// GovUpdateMinFees sets protocol controlled tx min fees.
	GovMinFees(ctx context.Context, in *MsgGovUpdateMinFees, opts ...grpc.CallOption) (*MsgGovUpdateMinFeesResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) GovMinFees(ctx context.Context, in *MsgGovUpdateMinFees, opts ...grpc.CallOption) (*MsgGovUpdateMinFeesResponse, error) {
	out := new(MsgGovUpdateMinFeesResponse)
	err := c.cc.Invoke(ctx, "/umee.ugov.v1.Msg/GovMinFees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// GovUpdateMinFees sets protocol controlled tx min fees.
	GovMinFees(context.Context, *MsgGovUpdateMinFees) (*MsgGovUpdateMinFeesResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) GovMinFees(ctx context.Context, req *MsgGovUpdateMinFees) (*MsgGovUpdateMinFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovMinFees not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_GovMinFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovUpdateMinFees)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovMinFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/umee.ugov.v1.Msg/GovMinFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovMinFees(ctx, req.(*MsgGovUpdateMinFees))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "umee.ugov.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GovMinFees",
			Handler:    _Msg_GovMinFees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umee/ugov/v1/tx.proto",
}

func (m *MsgGovUpdateMinFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovUpdateMinFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovUpdateMinFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgGovUpdateMinFeesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGovUpdateMinFeesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGovUpdateMinFeesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgGovUpdateMinFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.MinTxFees) > 0 {
		for _, e := range m.MinTxFees {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgGovUpdateMinFeesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgGovUpdateMinFees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgGovUpdateMinFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovUpdateMinFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTxFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinTxFees = append(m.MinTxFees, types.Coin{})
			if err := m.MinTxFees[len(m.MinTxFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgGovUpdateMinFeesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgGovUpdateMinFeesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGovUpdateMinFeesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
