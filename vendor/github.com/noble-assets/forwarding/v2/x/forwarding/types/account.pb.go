// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: noble/forwarding/v1/account.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
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

type ForwardingAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty"`
	Channel            string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Recipient          string `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	CreatedAt          int64  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Fallback           string `protobuf:"bytes,5,opt,name=fallback,proto3" json:"fallback,omitempty"`
}

func (m *ForwardingAccount) Reset()         { *m = ForwardingAccount{} }
func (m *ForwardingAccount) String() string { return proto.CompactTextString(m) }
func (*ForwardingAccount) ProtoMessage()    {}
func (*ForwardingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d86db85ab0c667b, []int{0}
}
func (m *ForwardingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ForwardingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ForwardingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ForwardingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardingAccount.Merge(m, src)
}
func (m *ForwardingAccount) XXX_Size() int {
	return m.Size()
}
func (m *ForwardingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardingAccount proto.InternalMessageInfo

func (m *ForwardingAccount) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *ForwardingAccount) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *ForwardingAccount) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *ForwardingAccount) GetFallback() string {
	if m != nil {
		return m.Fallback
	}
	return ""
}

func init() {
	proto.RegisterType((*ForwardingAccount)(nil), "noble.forwarding.v1.ForwardingAccount")
}

func init() { proto.RegisterFile("noble/forwarding/v1/account.proto", fileDescriptor_9d86db85ab0c667b) }

var fileDescriptor_9d86db85ab0c667b = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4e, 0x32, 0x31,
	0x14, 0x85, 0xe9, 0x0f, 0xbf, 0x4a, 0x71, 0xe3, 0xe8, 0x62, 0x24, 0x58, 0x47, 0x57, 0x6c, 0x98,
	0x66, 0x70, 0x63, 0xdc, 0xc1, 0xc2, 0x84, 0xed, 0x2c, 0xdd, 0x90, 0xdb, 0x52, 0x86, 0x89, 0x43,
	0x4b, 0xa6, 0x05, 0xf5, 0x2d, 0x7c, 0x18, 0x1f, 0xc2, 0xb8, 0x62, 0xe9, 0xca, 0x18, 0x78, 0x04,
	0x5f, 0xc0, 0xd8, 0x56, 0xd0, 0xc4, 0xdd, 0x7c, 0xe7, 0xde, 0x33, 0xe7, 0xa4, 0x17, 0x9f, 0x49,
	0xc5, 0x0a, 0x41, 0xc7, 0xaa, 0xbc, 0x83, 0x72, 0x94, 0xcb, 0x8c, 0x2e, 0x12, 0x0a, 0x9c, 0xab,
	0xb9, 0x34, 0xf1, 0xac, 0x54, 0x46, 0x05, 0x87, 0x76, 0x25, 0xde, 0xae, 0xc4, 0x8b, 0xa4, 0x49,
	0xb8, 0xd2, 0x53, 0xa5, 0x29, 0xcc, 0xcd, 0x84, 0x2e, 0x12, 0x26, 0x0c, 0x24, 0x16, 0x9c, 0xa9,
	0x79, 0xec, 0xe6, 0x43, 0x4b, 0xd4, 0x81, 0x1f, 0x1d, 0x65, 0x2a, 0x53, 0x4e, 0xff, 0xfa, 0x72,
	0xea, 0xf9, 0x07, 0xc2, 0x07, 0xd7, 0x9b, 0x88, 0x9e, 0x6b, 0x10, 0x0c, 0xf0, 0x3e, 0x03, 0x2d,
	0x86, 0xbe, 0x51, 0x88, 0x22, 0xd4, 0x6e, 0x74, 0xa3, 0xd8, 0xff, 0xd0, 0x06, 0xfa, 0xf4, 0xb8,
	0x0f, 0x5a, 0x78, 0x5f, 0xbf, 0xb6, 0x7c, 0x3b, 0x45, 0x69, 0x83, 0x6d, 0xa5, 0x20, 0xc4, 0xbb,
	0x7c, 0x02, 0x52, 0x8a, 0x22, 0xfc, 0x17, 0xa1, 0x76, 0x3d, 0xfd, 0xc6, 0xa0, 0x85, 0xeb, 0xa5,
	0xe0, 0xf9, 0x2c, 0x17, 0xd2, 0x84, 0x55, 0x3b, 0xdb, 0x0a, 0xc1, 0x09, 0xc6, 0xbc, 0x14, 0x60,
	0xc4, 0x68, 0x08, 0x26, 0xac, 0x45, 0xa8, 0x5d, 0x4d, 0xeb, 0x5e, 0xe9, 0x99, 0xa0, 0x89, 0xf7,
	0xc6, 0x50, 0x14, 0x0c, 0xf8, 0x6d, 0xf8, 0xdf, 0x7a, 0x37, 0x7c, 0x15, 0xbd, 0x3c, 0x75, 0x5a,
	0x7f, 0x55, 0xf5, 0x9d, 0x06, 0xfd, 0xf4, 0x79, 0x45, 0xd0, 0x72, 0x45, 0xd0, 0xfb, 0x8a, 0xa0,
	0xc7, 0x35, 0xa9, 0x2c, 0xd7, 0xa4, 0xf2, 0xba, 0x26, 0x95, 0x9b, 0xcb, 0x2c, 0x37, 0x93, 0x39,
	0x8b, 0xb9, 0x9a, 0x52, 0x7b, 0x80, 0x0e, 0x68, 0x2d, 0x8c, 0xfe, 0x75, 0xaa, 0x2e, 0xbd, 0xff,
	0xc9, 0xe6, 0x61, 0x26, 0x34, 0xdb, 0xb1, 0x0f, 0x7a, 0xf1, 0x19, 0x00, 0x00, 0xff, 0xff, 0x0a,
	0xd4, 0x7c, 0xf3, 0xdb, 0x01, 0x00, 0x00,
}

func (m *ForwardingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForwardingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ForwardingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fallback) > 0 {
		i -= len(m.Fallback)
		copy(dAtA[i:], m.Fallback)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Fallback)))
		i--
		dAtA[i] = 0x2a
	}
	if m.CreatedAt != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAccount(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ForwardingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovAccount(uint64(m.CreatedAt))
	}
	l = len(m.Fallback)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ForwardingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
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
			return fmt.Errorf("proto: ForwardingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForwardingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fallback", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fallback = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
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
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
				return 0, ErrInvalidLengthAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccount = fmt.Errorf("proto: unexpected end of group")
)