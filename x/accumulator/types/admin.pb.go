// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accumulator/admin.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Admin struct {
	Address             string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	VestingPeriod       int64      `protobuf:"varint,2,opt,name=vesting_period,json=vestingPeriod,proto3" json:"vesting_period,omitempty"`
	RewardPerPeriod     types.Coin `protobuf:"bytes,3,opt,name=reward_per_period,json=rewardPerPeriod,proto3" json:"reward_per_period"`
	VestingPeriodsCount int64      `protobuf:"varint,4,opt,name=vesting_periods_count,json=vestingPeriodsCount,proto3" json:"vesting_periods_count,omitempty"`
	LastVestingTime     int64      `protobuf:"varint,5,opt,name=last_vesting_time,json=lastVestingTime,proto3" json:"last_vesting_time,omitempty"`
	VestingCounter      int64      `protobuf:"varint,6,opt,name=vesting_counter,json=vestingCounter,proto3" json:"vesting_counter,omitempty"`
	Denom               string     `protobuf:"bytes,7,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *Admin) Reset()         { *m = Admin{} }
func (m *Admin) String() string { return proto.CompactTextString(m) }
func (*Admin) ProtoMessage()    {}
func (*Admin) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea173fc964753617, []int{0}
}
func (m *Admin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Admin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Admin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Admin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Admin.Merge(m, src)
}
func (m *Admin) XXX_Size() int {
	return m.Size()
}
func (m *Admin) XXX_DiscardUnknown() {
	xxx_messageInfo_Admin.DiscardUnknown(m)
}

var xxx_messageInfo_Admin proto.InternalMessageInfo

func (m *Admin) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Admin) GetVestingPeriod() int64 {
	if m != nil {
		return m.VestingPeriod
	}
	return 0
}

func (m *Admin) GetRewardPerPeriod() types.Coin {
	if m != nil {
		return m.RewardPerPeriod
	}
	return types.Coin{}
}

func (m *Admin) GetVestingPeriodsCount() int64 {
	if m != nil {
		return m.VestingPeriodsCount
	}
	return 0
}

func (m *Admin) GetLastVestingTime() int64 {
	if m != nil {
		return m.LastVestingTime
	}
	return 0
}

func (m *Admin) GetVestingCounter() int64 {
	if m != nil {
		return m.VestingCounter
	}
	return 0
}

func (m *Admin) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*Admin)(nil), "cosmos.accumulator.Admin")
}

func init() { proto.RegisterFile("cosmos/accumulator/admin.proto", fileDescriptor_ea173fc964753617) }

var fileDescriptor_ea173fc964753617 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0xcd, 0x8e, 0xda, 0x30,
	0x10, 0x4e, 0xf8, 0x55, 0x5d, 0xb5, 0x08, 0x97, 0x4a, 0x01, 0x55, 0x2e, 0xaa, 0x54, 0x15, 0x21,
	0x35, 0x2e, 0xf4, 0x09, 0x0a, 0xb7, 0xf6, 0x82, 0x50, 0xd5, 0x43, 0x2f, 0xc8, 0x49, 0xdc, 0xd4,
	0x2a, 0xf1, 0xa4, 0xb6, 0xc3, 0x2e, 0x6f, 0xb1, 0x0f, 0xb3, 0x0f, 0xc1, 0x91, 0xe3, 0x9e, 0x56,
	0x2b, 0x78, 0x91, 0x55, 0x1c, 0x47, 0x82, 0x53, 0x66, 0xbe, 0xef, 0x9b, 0x6f, 0xbe, 0x4c, 0x82,
	0x48, 0x0c, 0x3a, 0x03, 0x4d, 0x59, 0x1c, 0x17, 0x59, 0xb1, 0x65, 0x06, 0x14, 0x65, 0x49, 0x26,
	0x64, 0x98, 0x2b, 0x30, 0x80, 0x71, 0xc5, 0x87, 0x17, 0xfc, 0x68, 0x98, 0x02, 0xa4, 0x5b, 0x4e,
	0xad, 0x22, 0x2a, 0xfe, 0x50, 0x26, 0xf7, 0x95, 0x7c, 0x34, 0x75, 0x76, 0x11, 0xd3, 0x9c, 0xfe,
	0x2f, 0xb8, 0xda, 0xd3, 0xdd, 0x2c, 0xe2, 0x86, 0xcd, 0x68, 0xce, 0x52, 0x21, 0x99, 0x11, 0xe0,
	0xac, 0x47, 0xef, 0x9c, 0x0d, 0xcb, 0x05, 0x65, 0x52, 0x82, 0xb1, 0xa4, 0x76, 0xec, 0x20, 0x85,
	0x14, 0x6c, 0x49, 0xcb, 0xca, 0xa1, 0xe4, 0xd2, 0xbf, 0x76, 0x8e, 0xa1, 0x8e, 0xfb, 0xe1, 0xbe,
	0x81, 0xda, 0xdf, 0xca, 0xf8, 0x38, 0x40, 0x5d, 0x96, 0x24, 0x8a, 0x6b, 0x1d, 0xf8, 0x63, 0x7f,
	0xf2, 0x62, 0x5d, 0xb7, 0xf8, 0x23, 0x7a, 0xbd, 0xe3, 0xda, 0x08, 0x99, 0x6e, 0x72, 0xae, 0x04,
	0x24, 0x41, 0x63, 0xec, 0x4f, 0x9a, 0xeb, 0x57, 0x0e, 0x5d, 0x59, 0x10, 0xff, 0x40, 0x7d, 0xc5,
	0x6f, 0x98, 0x4a, 0x4a, 0x55, 0xad, 0x6c, 0x8e, 0xfd, 0xc9, 0xcb, 0xf9, 0x30, 0x74, 0x57, 0x29,
	0x63, 0x84, 0x2e, 0x46, 0xb8, 0x04, 0x21, 0x17, 0xad, 0xc3, 0xe3, 0x7b, 0x6f, 0xdd, 0xab, 0x26,
	0x57, 0x5c, 0x39, 0xb3, 0x39, 0x7a, 0x7b, 0xbd, 0x53, 0x6f, 0x62, 0x28, 0xa4, 0x09, 0x5a, 0x76,
	0xf5, 0x9b, 0xab, 0xd5, 0x7a, 0x59, 0x52, 0x78, 0x8a, 0xfa, 0x5b, 0xa6, 0xcd, 0xa6, 0x1e, 0x34,
	0x22, 0xe3, 0x41, 0xdb, 0xea, 0x7b, 0x25, 0xf1, 0xab, 0xc2, 0x7f, 0x8a, 0x8c, 0xe3, 0x4f, 0xa8,
	0x57, 0xcb, 0xac, 0x2f, 0x57, 0x41, 0xc7, 0x2a, 0xeb, 0x57, 0x5d, 0x56, 0x28, 0x1e, 0xa0, 0x76,
	0xc2, 0x25, 0x64, 0x41, 0xd7, 0x1e, 0xa5, 0x6a, 0x16, 0xdf, 0x0f, 0x27, 0xe2, 0x1f, 0x4f, 0xc4,
	0x7f, 0x3a, 0x11, 0xff, 0xee, 0x4c, 0xbc, 0xe3, 0x99, 0x78, 0x0f, 0x67, 0xe2, 0xfd, 0xfe, 0x92,
	0x0a, 0xf3, 0xb7, 0x88, 0xc2, 0x18, 0x32, 0xea, 0x6e, 0x5f, 0x3d, 0x3e, 0xeb, 0xe4, 0x1f, 0xbd,
	0xbd, 0xfa, 0x6f, 0xcc, 0x3e, 0xe7, 0x3a, 0xea, 0xd8, 0x2f, 0xf1, 0xf5, 0x39, 0x00, 0x00, 0xff,
	0xff, 0xef, 0x01, 0x2a, 0xaa, 0x5a, 0x02, 0x00, 0x00,
}

func (m *Admin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Admin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Admin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x3a
	}
	if m.VestingCounter != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.VestingCounter))
		i--
		dAtA[i] = 0x30
	}
	if m.LastVestingTime != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.LastVestingTime))
		i--
		dAtA[i] = 0x28
	}
	if m.VestingPeriodsCount != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.VestingPeriodsCount))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.RewardPerPeriod.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAdmin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.VestingPeriod != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.VestingPeriod))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAdmin(dAtA []byte, offset int, v uint64) int {
	offset -= sovAdmin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Admin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	if m.VestingPeriod != 0 {
		n += 1 + sovAdmin(uint64(m.VestingPeriod))
	}
	l = m.RewardPerPeriod.Size()
	n += 1 + l + sovAdmin(uint64(l))
	if m.VestingPeriodsCount != 0 {
		n += 1 + sovAdmin(uint64(m.VestingPeriodsCount))
	}
	if m.LastVestingTime != 0 {
		n += 1 + sovAdmin(uint64(m.LastVestingTime))
	}
	if m.VestingCounter != 0 {
		n += 1 + sovAdmin(uint64(m.VestingCounter))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	return n
}

func sovAdmin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAdmin(x uint64) (n int) {
	return sovAdmin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Admin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: Admin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Admin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriod", wireType)
			}
			m.VestingPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VestingPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPerPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardPerPeriod.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriodsCount", wireType)
			}
			m.VestingPeriodsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VestingPeriodsCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastVestingTime", wireType)
			}
			m.LastVestingTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastVestingTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingCounter", wireType)
			}
			m.VestingCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VestingCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmin
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
func skipAdmin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
				return 0, ErrInvalidLengthAdmin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAdmin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAdmin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAdmin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdmin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAdmin = fmt.Errorf("proto: unexpected end of group")
)
