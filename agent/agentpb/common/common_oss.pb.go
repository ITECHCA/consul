// +build !consulent

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: agent/agentpb/common/common_oss.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EnterpriseMeta struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterpriseMeta) Reset()         { *m = EnterpriseMeta{} }
func (m *EnterpriseMeta) String() string { return proto.CompactTextString(m) }
func (*EnterpriseMeta) ProtoMessage()    {}
func (*EnterpriseMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6fd603a9adc5dd5, []int{0}
}
func (m *EnterpriseMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EnterpriseMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EnterpriseMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EnterpriseMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterpriseMeta.Merge(m, src)
}
func (m *EnterpriseMeta) XXX_Size() int {
	return m.Size()
}
func (m *EnterpriseMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterpriseMeta.DiscardUnknown(m)
}

var xxx_messageInfo_EnterpriseMeta proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EnterpriseMeta)(nil), "common.EnterpriseMeta")
}

func init() {
	proto.RegisterFile("agent/agentpb/common/common_oss.proto", fileDescriptor_b6fd603a9adc5dd5)
}

var fileDescriptor_b6fd603a9adc5dd5 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0x4c, 0x4f, 0xcd,
	0x2b, 0xd1, 0x07, 0x93, 0x05, 0x49, 0xfa, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0x50, 0x2a, 0x3e,
	0xbf, 0xb8, 0x58, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x0d, 0x22, 0xa2, 0x24, 0xc0, 0xc5,
	0xe7, 0x9a, 0x57, 0x92, 0x5a, 0x54, 0x50, 0x94, 0x59, 0x9c, 0xea, 0x9b, 0x5a, 0x92, 0xe8, 0xe4,
	0x74, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c,
	0xc7, 0x10, 0x65, 0x90, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x91,
	0x58, 0x9c, 0x91, 0x99, 0x9c, 0x5f, 0x54, 0xa0, 0x9f, 0x9c, 0x9f, 0x57, 0x5c, 0x9a, 0xa3, 0x8f,
	0xcd, 0xba, 0x24, 0x36, 0xb0, 0x25, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x73, 0x1c,
	0x7c, 0x8d, 0x00, 0x00, 0x00,
}

func (m *EnterpriseMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnterpriseMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintCommonOss(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EnterpriseMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCommonOss(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCommonOss(x uint64) (n int) {
	return sovCommonOss(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EnterpriseMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommonOss
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
			return fmt.Errorf("proto: EnterpriseMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnterpriseMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCommonOss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommonOss
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommonOss
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCommonOss(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommonOss
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
					return 0, ErrIntOverflowCommonOss
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommonOss
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
				return 0, ErrInvalidLengthCommonOss
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCommonOss
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCommonOss
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCommonOss(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCommonOss
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCommonOss = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommonOss   = fmt.Errorf("proto: integer overflow")
)