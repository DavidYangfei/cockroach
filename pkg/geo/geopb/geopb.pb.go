// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: geo/geopb/geopb.proto

package geopb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// ShapeType is the type of a spatial shape. Each of these corresponds to a
// different representation and serialization format. For example, a Point is a
// pair of doubles (or more than that for geometries with Z or N), a LineString
// is an ordered series of Points, etc.
type ShapeType int32

const (
	ShapeType_Unset           ShapeType = 0
	ShapeType_Point           ShapeType = 1
	ShapeType_LineString      ShapeType = 2
	ShapeType_Polygon         ShapeType = 3
	ShapeType_MultiPoint      ShapeType = 4
	ShapeType_MultiLineString ShapeType = 5
	ShapeType_MultiPolygon    ShapeType = 6
	// Geometry can contain any type.
	ShapeType_Geometry ShapeType = 7
	// GeometryCollection can contain a list of any above type except for Geometry.
	ShapeType_GeometryCollection ShapeType = 8
)

var ShapeType_name = map[int32]string{
	0: "Unset",
	1: "Point",
	2: "LineString",
	3: "Polygon",
	4: "MultiPoint",
	5: "MultiLineString",
	6: "MultiPolygon",
	7: "Geometry",
	8: "GeometryCollection",
}
var ShapeType_value = map[string]int32{
	"Unset":              0,
	"Point":              1,
	"LineString":         2,
	"Polygon":            3,
	"MultiPoint":         4,
	"MultiLineString":    5,
	"MultiPolygon":       6,
	"Geometry":           7,
	"GeometryCollection": 8,
}

func (x ShapeType) String() string {
	return proto.EnumName(ShapeType_name, int32(x))
}
func (ShapeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_geopb_d79593e6d8b39035, []int{0}
}

// SpatialObject represents a serialization of a Geospatial type.
type SpatialObject struct {
	// EWKB is the EWKB representation of the spatial object.
	EWKB EWKB `protobuf:"bytes,1,opt,name=ewkb,proto3,casttype=EWKB" json:"ewkb,omitempty"`
	// SRID is the denormalized SRID derived from the EWKB.
	SRID SRID `protobuf:"varint,2,opt,name=srid,proto3,casttype=SRID" json:"srid,omitempty"`
	// ShapeType is denormalized ShapeType derived from the EWKB.
	ShapeType ShapeType `protobuf:"varint,3,opt,name=shape_type,json=shapeType,proto3,enum=cockroach.geopb.ShapeType" json:"shape_type,omitempty"`
	// BoundingBox is the bounding box of the SpatialObject.
	BoundingBox *BoundingBox `protobuf:"bytes,4,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
}

func (m *SpatialObject) Reset()         { *m = SpatialObject{} }
func (m *SpatialObject) String() string { return proto.CompactTextString(m) }
func (*SpatialObject) ProtoMessage()    {}
func (*SpatialObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_geopb_d79593e6d8b39035, []int{0}
}
func (m *SpatialObject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpatialObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *SpatialObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpatialObject.Merge(dst, src)
}
func (m *SpatialObject) XXX_Size() int {
	return m.Size()
}
func (m *SpatialObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SpatialObject.DiscardUnknown(m)
}

var xxx_messageInfo_SpatialObject proto.InternalMessageInfo

// BoundingBox represents the bounding box of a Geospatial type.
type BoundingBox struct {
	MinX float64 `protobuf:"fixed64,1,opt,name=min_x,json=minX,proto3" json:"min_x,omitempty"`
	MinY float64 `protobuf:"fixed64,2,opt,name=min_y,json=minY,proto3" json:"min_y,omitempty"`
	MaxX float64 `protobuf:"fixed64,3,opt,name=max_x,json=maxX,proto3" json:"max_x,omitempty"`
	MaxY float64 `protobuf:"fixed64,4,opt,name=max_y,json=maxY,proto3" json:"max_y,omitempty"`
}

func (m *BoundingBox) Reset()         { *m = BoundingBox{} }
func (m *BoundingBox) String() string { return proto.CompactTextString(m) }
func (*BoundingBox) ProtoMessage()    {}
func (*BoundingBox) Descriptor() ([]byte, []int) {
	return fileDescriptor_geopb_d79593e6d8b39035, []int{1}
}
func (m *BoundingBox) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BoundingBox) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *BoundingBox) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingBox.Merge(dst, src)
}
func (m *BoundingBox) XXX_Size() int {
	return m.Size()
}
func (m *BoundingBox) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingBox.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingBox proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SpatialObject)(nil), "cockroach.geopb.SpatialObject")
	proto.RegisterType((*BoundingBox)(nil), "cockroach.geopb.BoundingBox")
	proto.RegisterEnum("cockroach.geopb.ShapeType", ShapeType_name, ShapeType_value)
}
func (m *SpatialObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpatialObject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.EWKB) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGeopb(dAtA, i, uint64(len(m.EWKB)))
		i += copy(dAtA[i:], m.EWKB)
	}
	if m.SRID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGeopb(dAtA, i, uint64(m.SRID))
	}
	if m.ShapeType != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGeopb(dAtA, i, uint64(m.ShapeType))
	}
	if m.BoundingBox != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGeopb(dAtA, i, uint64(m.BoundingBox.Size()))
		n1, err := m.BoundingBox.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *BoundingBox) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BoundingBox) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MinX != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.MinX))))
		i += 8
	}
	if m.MinY != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.MinY))))
		i += 8
	}
	if m.MaxX != 0 {
		dAtA[i] = 0x19
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.MaxX))))
		i += 8
	}
	if m.MaxY != 0 {
		dAtA[i] = 0x21
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.MaxY))))
		i += 8
	}
	return i, nil
}

func encodeVarintGeopb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SpatialObject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EWKB)
	if l > 0 {
		n += 1 + l + sovGeopb(uint64(l))
	}
	if m.SRID != 0 {
		n += 1 + sovGeopb(uint64(m.SRID))
	}
	if m.ShapeType != 0 {
		n += 1 + sovGeopb(uint64(m.ShapeType))
	}
	if m.BoundingBox != nil {
		l = m.BoundingBox.Size()
		n += 1 + l + sovGeopb(uint64(l))
	}
	return n
}

func (m *BoundingBox) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinX != 0 {
		n += 9
	}
	if m.MinY != 0 {
		n += 9
	}
	if m.MaxX != 0 {
		n += 9
	}
	if m.MaxY != 0 {
		n += 9
	}
	return n
}

func sovGeopb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGeopb(x uint64) (n int) {
	return sovGeopb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SpatialObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGeopb
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
			return fmt.Errorf("proto: SpatialObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpatialObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EWKB", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeopb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGeopb
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EWKB = append(m.EWKB[:0], dAtA[iNdEx:postIndex]...)
			if m.EWKB == nil {
				m.EWKB = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SRID", wireType)
			}
			m.SRID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeopb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SRID |= (SRID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShapeType", wireType)
			}
			m.ShapeType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeopb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShapeType |= (ShapeType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoundingBox", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeopb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGeopb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BoundingBox == nil {
				m.BoundingBox = &BoundingBox{}
			}
			if err := m.BoundingBox.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGeopb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGeopb
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
func (m *BoundingBox) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGeopb
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
			return fmt.Errorf("proto: BoundingBox: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BoundingBox: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinX", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MinX = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinY", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MinY = float64(math.Float64frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxX", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MaxX = float64(math.Float64frombits(v))
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxY", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MaxY = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipGeopb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGeopb
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
func skipGeopb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGeopb
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
					return 0, ErrIntOverflowGeopb
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
					return 0, ErrIntOverflowGeopb
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
				return 0, ErrInvalidLengthGeopb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGeopb
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
				next, err := skipGeopb(dAtA[start:])
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
	ErrInvalidLengthGeopb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGeopb   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("geo/geopb/geopb.proto", fileDescriptor_geopb_d79593e6d8b39035) }

var fileDescriptor_geopb_d79593e6d8b39035 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x33, 0x6d, 0xd2, 0x76, 0xdf, 0xae, 0xed, 0x30, 0x55, 0x59, 0x8a, 0x8c, 0x4b, 0x11,
	0x5c, 0x3c, 0xa4, 0x50, 0x4f, 0x9e, 0x84, 0xa8, 0x88, 0xa8, 0x58, 0x12, 0xc5, 0xd6, 0x4b, 0x48,
	0xd2, 0x21, 0x1d, 0x9b, 0x9d, 0x17, 0x92, 0x29, 0x26, 0xdf, 0xc2, 0xbb, 0x5f, 0xa8, 0xc7, 0xde,
	0xec, 0x49, 0x34, 0xfb, 0x2d, 0x7a, 0x92, 0x4c, 0x12, 0x5d, 0xf4, 0x92, 0xf7, 0xde, 0xff, 0xff,
	0x7b, 0xf0, 0x7f, 0x61, 0xe0, 0x4e, 0x2a, 0xf0, 0x20, 0x15, 0x98, 0xc7, 0xdd, 0xd7, 0xcd, 0x0b,
	0xd4, 0xc8, 0x76, 0x12, 0x4c, 0xce, 0x0b, 0x8c, 0x92, 0x33, 0xd7, 0xc8, 0x7b, 0xb7, 0x53, 0x4c,
	0xd1, 0x78, 0x07, 0x6d, 0xd7, 0x61, 0xfb, 0xdf, 0x09, 0xdc, 0x0a, 0xf2, 0x48, 0xcb, 0x28, 0x7b,
	0x17, 0x7f, 0x16, 0x89, 0x66, 0x0f, 0xc0, 0x16, 0x5f, 0xce, 0xe3, 0x29, 0x99, 0x91, 0xf9, 0xc4,
	0xa3, 0xcd, 0x8f, 0xfb, 0xf6, 0x8b, 0x8f, 0xaf, 0xbd, 0x9b, 0xbe, 0xfa, 0xc6, 0x6d, 0xa9, 0xb2,
	0x90, 0xa7, 0xd3, 0xb5, 0x19, 0x99, 0x3b, 0x1d, 0x15, 0xf8, 0xaf, 0x9e, 0xdf, 0xf4, 0xd5, 0x37,
	0x2e, 0x7b, 0x02, 0x50, 0x9e, 0x45, 0xb9, 0x08, 0x75, 0x9d, 0x8b, 0xe9, 0xfa, 0x8c, 0xcc, 0xb7,
	0x0f, 0xf7, 0xdc, 0x7f, 0x92, 0xb9, 0x41, 0x8b, 0xbc, 0xaf, 0x73, 0xe1, 0x8f, 0xca, 0xa1, 0x65,
	0x4f, 0x61, 0x12, 0xe3, 0x85, 0x3a, 0x95, 0x2a, 0x0d, 0x63, 0xac, 0xa6, 0xf6, 0x8c, 0xcc, 0xc7,
	0x87, 0xf7, 0xfe, 0x5b, 0xf6, 0x7a, 0xc8, 0xc3, 0xca, 0x1f, 0xc7, 0x7f, 0x87, 0xfd, 0x08, 0xc6,
	0x2b, 0x1e, 0xdb, 0x05, 0x67, 0x21, 0x55, 0x58, 0x99, 0xbb, 0x88, 0x6f, 0x2f, 0xa4, 0x3a, 0x1e,
	0xc4, 0xda, 0x9c, 0xd1, 0x89, 0x27, 0x46, 0x8c, 0xaa, 0xb0, 0x32, 0x79, 0x5b, 0x31, 0xaa, 0x8e,
	0x07, 0xb1, 0x36, 0x39, 0x3a, 0xf1, 0xe4, 0xd1, 0x37, 0x02, 0xa3, 0x3f, 0xe1, 0xd9, 0x08, 0x9c,
	0x0f, 0xaa, 0x14, 0x9a, 0x5a, 0x6d, 0x7b, 0x84, 0x52, 0x69, 0x4a, 0xd8, 0x36, 0xc0, 0x1b, 0xa9,
	0x44, 0xa0, 0x0b, 0xa9, 0x52, 0xba, 0xc6, 0xc6, 0xb0, 0x79, 0x84, 0x59, 0x9d, 0xa2, 0xa2, 0xeb,
	0xad, 0xf9, 0xf6, 0x22, 0xd3, 0xb2, 0x83, 0x6d, 0xb6, 0x0b, 0x3b, 0x66, 0x5e, 0xd9, 0x70, 0x18,
	0x85, 0x49, 0x0f, 0x75, 0x6b, 0x1b, 0x6c, 0x02, 0x5b, 0x2f, 0x05, 0x2e, 0x84, 0x2e, 0x6a, 0xba,
	0xc9, 0xee, 0x02, 0x1b, 0xa6, 0x67, 0x98, 0x65, 0x22, 0xd1, 0x12, 0x15, 0xdd, 0xf2, 0x1e, 0x5e,
	0xfe, 0xe2, 0xd6, 0x65, 0xc3, 0xc9, 0x55, 0xc3, 0xc9, 0x75, 0xc3, 0xc9, 0xcf, 0x86, 0x93, 0xaf,
	0x4b, 0x6e, 0x5d, 0x2d, 0xb9, 0x75, 0xbd, 0xe4, 0xd6, 0x27, 0xc7, 0xfc, 0xc2, 0x78, 0xc3, 0x3c,
	0x85, 0xc7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x0b, 0x89, 0x20, 0x4a, 0x02, 0x00, 0x00,
}
