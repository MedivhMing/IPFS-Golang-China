
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//由原基因gogo生成的代码。不要编辑。
//来源：pin/internal/pb/header.proto

package pb

import proto "gx/ipfs/QmdxUuburamoF6zF9qjeQC4WYcWGbWuRmdLacMEsW8ioD8/gogo-protobuf/proto"
import fmt "fmt"
import math "math"

import encoding_binary "encoding/binary"

import io "io"

//引用导入以禁止错误（如果未使用）。
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//这是一个编译时断言，以确保生成的文件
//与正在编译的proto包兼容。
//此行的编译错误可能意味着您的
//需要更新proto包。
const _ = proto.GoGoProtoPackageIsVersion2 //请升级proto包

type Set struct {
//1目前，库将拒绝处理版本无法识别的条目。
	Version uint32 `protobuf:"varint,1,opt,name=version" json:"version"`
//有多少链接是子树
	Fanout uint32 `protobuf:"varint,2,opt,name=fanout" json:"fanout"`
//用于子树选择的哈希种子，随机数
	Seed                 uint32   `protobuf:"fixed32,3,opt,name=seed" json:"seed"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Set) Reset()         { *m = Set{} }
func (m *Set) String() string { return proto.CompactTextString(m) }
func (*Set) ProtoMessage()    {}
func (*Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_header_778100e52d428560, []int{0}
}
func (m *Set) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Set.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Set.Merge(dst, src)
}
func (m *Set) XXX_Size() int {
	return m.Size()
}
func (m *Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Set proto.InternalMessageInfo

func (m *Set) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Set) GetFanout() uint32 {
	if m != nil {
		return m.Fanout
	}
	return 0
}

func (m *Set) GetSeed() uint32 {
	if m != nil {
		return m.Seed
	}
	return 0
}

func init() {
	proto.RegisterType((*Set)(nil), "ipfs.pin.Set")
}
func (m *Set) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Set) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintHeader(dAtA, i, uint64(m.Version))
	dAtA[i] = 0x10
	i++
	i = encodeVarintHeader(dAtA, i, uint64(m.Fanout))
	dAtA[i] = 0x1d
	i++
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Seed))
	i += 4
	return i, nil
}

func encodeVarintHeader(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Set) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovHeader(uint64(m.Version))
	n += 1 + sovHeader(uint64(m.Fanout))
	n += 5
	return n
}

func sovHeader(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHeader(x uint64) (n int) {
	return sovHeader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Set) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeader
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Set: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Set: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fanout", wireType)
			}
			m.Fanout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fanout |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			m.Seed = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Seed = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		default:
			iNdEx = preIndex
			skippy, err := skipHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeader
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
func skipHeader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeader
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
					return 0, ErrIntOverflowHeader
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
					return 0, ErrIntOverflowHeader
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHeader
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHeader
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
				next, err := skipHeader(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthHeader = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeader   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("pin/internal/pb/header.proto", fileDescriptor_header_778100e52d428560)
}

var fileDescriptor_header_778100e52d428560 = []byte{
//gzip文件描述符或协议的154字节
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0xc8, 0xcc, 0xd3,
	0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x48, 0xd2, 0xcf, 0x48, 0x4d, 0x4c,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x2c, 0x48, 0x2b, 0xd6, 0x2b,
	0xc8, 0xcc, 0x53, 0x8a, 0xe5, 0x62, 0x0e, 0x4e, 0x2d, 0x11, 0x92, 0xe3, 0x62, 0x2f, 0x4b, 0x2d,
	0x2a, 0xce, 0xcc, 0xcf, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x75, 0x62, 0x39, 0x71, 0x4f, 0x9e,
	0x21, 0x08, 0x26, 0x28, 0x24, 0xc3, 0xc5, 0x96, 0x96, 0x98, 0x97, 0x5f, 0x5a, 0x22, 0xc1, 0x84,
	0x24, 0x0d, 0x15, 0x13, 0x92, 0xe0, 0x62, 0x29, 0x4e, 0x4d, 0x4d, 0x91, 0x60, 0x56, 0x60, 0xd4,
	0x60, 0x87, 0xca, 0x81, 0x45, 0x9c, 0x44, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0xa2, 0x98, 0x0a, 0x92, 0x00, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xc3, 0xf9, 0x7f, 0x24, 0x9d, 0x00, 0x00, 0x00,
}
