// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc_query/v1/packet.proto

package types

import (
	fmt "fmt"
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

// IBCQueryPacketData defines a struct for the cross chain query packet payload
type IBCQueryPacketData struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path        string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	QueryHeight uint64 `protobuf:"varint,3,opt,name=query_height,json=queryHeight,proto3" json:"query_height,omitempty"`
}

func (m *IBCQueryPacketData) Reset()         { *m = IBCQueryPacketData{} }
func (m *IBCQueryPacketData) String() string { return proto.CompactTextString(m) }
func (*IBCQueryPacketData) ProtoMessage()    {}
func (*IBCQueryPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e14a8048b2ebdda4, []int{0}
}
func (m *IBCQueryPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCQueryPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCQueryPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCQueryPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCQueryPacketData.Merge(m, src)
}
func (m *IBCQueryPacketData) XXX_Size() int {
	return m.Size()
}
func (m *IBCQueryPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCQueryPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_IBCQueryPacketData proto.InternalMessageInfo

func (m *IBCQueryPacketData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *IBCQueryPacketData) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *IBCQueryPacketData) GetQueryHeight() uint64 {
	if m != nil {
		return m.QueryHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*IBCQueryPacketData)(nil), "ibc.applications.ibc_query.v1.IBCQueryPacketData")
}

func init() { proto.RegisterFile("ibc_query/v1/packet.proto", fileDescriptor_e14a8048b2ebdda4) }

var fileDescriptor_e14a8048b2ebdda4 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x4c, 0x4a, 0x8e,
	0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2f, 0x33, 0xd4, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xcd, 0x4c, 0x4a, 0xd6, 0x4b, 0x2c, 0x28, 0xc8, 0xc9,
	0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x83, 0xab, 0xd5, 0x2b, 0x33, 0x54, 0x8a, 0xe6,
	0x12, 0xf2, 0x74, 0x72, 0x0e, 0x04, 0x71, 0x03, 0xc0, 0xda, 0x5c, 0x12, 0x4b, 0x12, 0x85, 0xf8,
	0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x84, 0x84,
	0xb8, 0x58, 0x0a, 0x12, 0x4b, 0x32, 0x24, 0x98, 0xc0, 0x22, 0x60, 0xb6, 0x90, 0x22, 0x17, 0x0f,
	0xd8, 0x94, 0xf8, 0x8c, 0xd4, 0xcc, 0xf4, 0x8c, 0x12, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x96, 0x20,
	0x6e, 0xb0, 0x98, 0x07, 0x58, 0xc8, 0x29, 0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5,
	0x18, 0xa2, 0xcc, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x93, 0xf3,
	0x8b, 0x73, 0xf3, 0x8b, 0xf5, 0x33, 0xf3, 0x4a, 0x52, 0x8b, 0x92, 0x33, 0x12, 0x33, 0xf3, 0x74,
	0x41, 0x66, 0x64, 0xa6, 0x16, 0xeb, 0x57, 0xe8, 0x23, 0xfc, 0x55, 0x52, 0x59, 0x90, 0x5a, 0x9c,
	0xc4, 0x06, 0xf6, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x85, 0x6d, 0xc8, 0xcc, 0xf1, 0x00,
	0x00, 0x00,
}

func (m *IBCQueryPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCQueryPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCQueryPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.QueryHeight != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.QueryHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IBCQueryPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.QueryHeight != 0 {
		n += 1 + sovPacket(uint64(m.QueryHeight))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IBCQueryPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: IBCQueryPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCQueryPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryHeight", wireType)
			}
			m.QueryHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QueryHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
