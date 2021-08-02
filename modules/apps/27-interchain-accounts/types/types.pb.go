// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/interchain_accounts/v1/types.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// The different types of interchain account transactions
// EXECUTE_TX is used when sending a TX from the controller side to the host side. The host side will execute the tx on
// behalf of the interchain account.
type Type int32

const (
	// Execute message type
	EXECUTE_TX Type = 0
)

var Type_name = map[int32]string{
	0: "TYPE_EXECUTE_TX_UNSPECIFIED",
}

var Type_value = map[string]int32{
	"TYPE_EXECUTE_TX_UNSPECIFIED": 0,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39bab93e18d89799, []int{0}
}

// Raw tx body
type IBCTxRaw struct {
	BodyBytes []byte `protobuf:"bytes,1,opt,name=body_bytes,json=bodyBytes,proto3" json:"body_bytes,omitempty" yaml:"body_bytes"`
}

func (m *IBCTxRaw) Reset()         { *m = IBCTxRaw{} }
func (m *IBCTxRaw) String() string { return proto.CompactTextString(m) }
func (*IBCTxRaw) ProtoMessage()    {}
func (*IBCTxRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_39bab93e18d89799, []int{0}
}
func (m *IBCTxRaw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCTxRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCTxRaw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCTxRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCTxRaw.Merge(m, src)
}
func (m *IBCTxRaw) XXX_Size() int {
	return m.Size()
}
func (m *IBCTxRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCTxRaw.DiscardUnknown(m)
}

var xxx_messageInfo_IBCTxRaw proto.InternalMessageInfo

func (m *IBCTxRaw) GetBodyBytes() []byte {
	if m != nil {
		return m.BodyBytes
	}
	return nil
}

// Body of a tx for an ics27 IBC packet
type IBCTxBody struct {
	Messages []*types.Any `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (m *IBCTxBody) Reset()         { *m = IBCTxBody{} }
func (m *IBCTxBody) String() string { return proto.CompactTextString(m) }
func (*IBCTxBody) ProtoMessage()    {}
func (*IBCTxBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_39bab93e18d89799, []int{1}
}
func (m *IBCTxBody) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCTxBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCTxBody.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCTxBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCTxBody.Merge(m, src)
}
func (m *IBCTxBody) XXX_Size() int {
	return m.Size()
}
func (m *IBCTxBody) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCTxBody.DiscardUnknown(m)
}

var xxx_messageInfo_IBCTxBody proto.InternalMessageInfo

func (m *IBCTxBody) GetMessages() []*types.Any {
	if m != nil {
		return m.Messages
	}
	return nil
}

// Packet data is comprised of raw transaction & type of transaction
type IBCAccountPacketData struct {
	Type Type   `protobuf:"varint,1,opt,name=type,proto3,enum=ibc.applications.interchain_accounts.v1.Type" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *IBCAccountPacketData) Reset()         { *m = IBCAccountPacketData{} }
func (m *IBCAccountPacketData) String() string { return proto.CompactTextString(m) }
func (*IBCAccountPacketData) ProtoMessage()    {}
func (*IBCAccountPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_39bab93e18d89799, []int{2}
}
func (m *IBCAccountPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCAccountPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCAccountPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCAccountPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCAccountPacketData.Merge(m, src)
}
func (m *IBCAccountPacketData) XXX_Size() int {
	return m.Size()
}
func (m *IBCAccountPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCAccountPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_IBCAccountPacketData proto.InternalMessageInfo

func (m *IBCAccountPacketData) GetType() Type {
	if m != nil {
		return m.Type
	}
	return EXECUTE_TX
}

func (m *IBCAccountPacketData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("ibc.applications.interchain_accounts.v1.Type", Type_name, Type_value)
	proto.RegisterType((*IBCTxRaw)(nil), "ibc.applications.interchain_accounts.v1.IBCTxRaw")
	proto.RegisterType((*IBCTxBody)(nil), "ibc.applications.interchain_accounts.v1.IBCTxBody")
	proto.RegisterType((*IBCAccountPacketData)(nil), "ibc.applications.interchain_accounts.v1.IBCAccountPacketData")
}

func init() {
	proto.RegisterFile("ibc/applications/interchain_accounts/v1/types.proto", fileDescriptor_39bab93e18d89799)
}

var fileDescriptor_39bab93e18d89799 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x41, 0x6b, 0xd4, 0x40,
	0x18, 0x4d, 0x34, 0x48, 0x3b, 0x4a, 0xd1, 0xb0, 0x42, 0x8d, 0x10, 0x4b, 0x2e, 0x16, 0x21, 0x33,
	0x76, 0x2b, 0x08, 0x42, 0xc1, 0x4d, 0x1a, 0x21, 0x17, 0x59, 0x62, 0x0a, 0x55, 0x84, 0x30, 0x33,
	0x19, 0xd3, 0xc1, 0x24, 0x13, 0x3a, 0x93, 0xd5, 0xf9, 0x07, 0xe2, 0xc9, 0x3f, 0xe0, 0xc9, 0x3f,
	0xe3, 0x71, 0x8f, 0x9e, 0x44, 0x76, 0xff, 0x81, 0xbf, 0x40, 0x32, 0xc1, 0x5d, 0x0f, 0x1e, 0x7a,
	0x7b, 0x33, 0xdf, 0xf7, 0x1e, 0xef, 0x7b, 0x0f, 0x1c, 0x73, 0x42, 0x11, 0xee, 0xba, 0x9a, 0x53,
	0xac, 0xb8, 0x68, 0x25, 0xe2, 0xad, 0x62, 0x97, 0xf4, 0x02, 0xf3, 0xb6, 0xc0, 0x94, 0x8a, 0xbe,
	0x55, 0x12, 0x2d, 0x8e, 0x90, 0xd2, 0x1d, 0x93, 0xb0, 0xbb, 0x14, 0x4a, 0xb8, 0x0f, 0x39, 0xa1,
	0xf0, 0x5f, 0x12, 0xfc, 0x0f, 0x09, 0x2e, 0x8e, 0xbc, 0x7b, 0x95, 0x10, 0x55, 0xcd, 0x90, 0xa1,
	0x91, 0xfe, 0x1d, 0xc2, 0xad, 0x1e, 0x35, 0xbc, 0x49, 0x25, 0x2a, 0x61, 0x20, 0x1a, 0xd0, 0xf8,
	0x1b, 0x3c, 0x07, 0x3b, 0x69, 0x14, 0xe7, 0x1f, 0x33, 0xfc, 0xc1, 0x7d, 0x02, 0x00, 0x11, 0xa5,
	0x2e, 0x88, 0x56, 0x4c, 0xee, 0xdb, 0x07, 0xf6, 0xe1, 0xad, 0xe8, 0xee, 0xef, 0x9f, 0x0f, 0xee,
	0x68, 0xdc, 0xd4, 0xcf, 0x82, 0xed, 0x2c, 0xc8, 0x76, 0x87, 0x47, 0x64, 0xf0, 0x09, 0xd8, 0x35,
	0x0a, 0x91, 0x28, 0xb5, 0xfb, 0x18, 0xec, 0x34, 0x4c, 0x4a, 0x5c, 0x19, 0x81, 0xeb, 0x87, 0x37,
	0xa7, 0x13, 0x38, 0x5a, 0x82, 0x7f, 0x2d, 0xc1, 0x59, 0xab, 0xb3, 0xcd, 0x56, 0xd0, 0x80, 0x49,
	0x1a, 0xc5, 0xb3, 0xf1, 0x86, 0x39, 0xa6, 0xef, 0x99, 0x3a, 0xc5, 0x0a, 0xbb, 0x33, 0xe0, 0x0c,
	0x09, 0x18, 0x1b, 0x7b, 0xd3, 0x10, 0x5e, 0x31, 0x01, 0x98, 0xeb, 0x8e, 0x65, 0x86, 0xea, 0xba,
	0xc0, 0x29, 0xb1, 0xc2, 0xfb, 0xd7, 0x86, 0x4b, 0x32, 0x83, 0x1f, 0x9d, 0x00, 0x67, 0xd8, 0x70,
	0x11, 0xb8, 0x9f, 0xbf, 0x9e, 0x27, 0x45, 0x72, 0x9e, 0xc4, 0x67, 0x79, 0x52, 0xe4, 0xe7, 0xc5,
	0xd9, 0xcb, 0x57, 0xf3, 0x24, 0x4e, 0x5f, 0xa4, 0xc9, 0xe9, 0x6d, 0xcb, 0xdb, 0xfb, 0xfc, 0xf5,
	0x00, 0x6c, 0xa7, 0x9e, 0xf3, 0xe9, 0x9b, 0x6f, 0x45, 0x6f, 0xbf, 0xaf, 0x7c, 0x7b, 0xb9, 0xf2,
	0xed, 0x5f, 0x2b, 0xdf, 0xfe, 0xb2, 0xf6, 0xad, 0xe5, 0xda, 0xb7, 0x7e, 0xac, 0x7d, 0xeb, 0x4d,
	0x54, 0x71, 0x75, 0xd1, 0x13, 0x48, 0x45, 0x83, 0xa8, 0x90, 0x8d, 0x90, 0x88, 0x13, 0x1a, 0x56,
	0x02, 0x35, 0xa2, 0xec, 0x6b, 0x26, 0x87, 0xd2, 0x25, 0x9a, 0x3e, 0x0d, 0xb7, 0xc6, 0xc3, 0x4d,
	0xdf, 0xa6, 0x6c, 0x72, 0xc3, 0x64, 0x74, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0x70, 0xa3, 0x00,
	0x15, 0x24, 0x02, 0x00, 0x00,
}

func (m *IBCTxRaw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCTxRaw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCTxRaw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BodyBytes) > 0 {
		i -= len(m.BodyBytes)
		copy(dAtA[i:], m.BodyBytes)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BodyBytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IBCTxBody) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCTxBody) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCTxBody) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *IBCAccountPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCAccountPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCAccountPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IBCTxRaw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BodyBytes)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *IBCTxBody) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *IBCAccountPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IBCTxRaw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: IBCTxRaw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCTxRaw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BodyBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BodyBytes = append(m.BodyBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.BodyBytes == nil {
				m.BodyBytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *IBCTxBody) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: IBCTxBody: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCTxBody: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
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
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *IBCAccountPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: IBCAccountPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCAccountPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
