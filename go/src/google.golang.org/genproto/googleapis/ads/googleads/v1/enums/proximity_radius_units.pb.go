// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/proximity_radius_units.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The unit of radius distance in proximity (e.g. MILES)
type ProximityRadiusUnitsEnum_ProximityRadiusUnits int32

const (
	// Not specified.
	ProximityRadiusUnitsEnum_UNSPECIFIED ProximityRadiusUnitsEnum_ProximityRadiusUnits = 0
	// Used for return value only. Represents value unknown in this version.
	ProximityRadiusUnitsEnum_UNKNOWN ProximityRadiusUnitsEnum_ProximityRadiusUnits = 1
	// Miles
	ProximityRadiusUnitsEnum_MILES ProximityRadiusUnitsEnum_ProximityRadiusUnits = 2
	// Kilometers
	ProximityRadiusUnitsEnum_KILOMETERS ProximityRadiusUnitsEnum_ProximityRadiusUnits = 3
)

var ProximityRadiusUnitsEnum_ProximityRadiusUnits_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MILES",
	3: "KILOMETERS",
}
var ProximityRadiusUnitsEnum_ProximityRadiusUnits_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"MILES":       2,
	"KILOMETERS":  3,
}

func (x ProximityRadiusUnitsEnum_ProximityRadiusUnits) String() string {
	return proto.EnumName(ProximityRadiusUnitsEnum_ProximityRadiusUnits_name, int32(x))
}
func (ProximityRadiusUnitsEnum_ProximityRadiusUnits) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_proximity_radius_units_9767cab966548058, []int{0, 0}
}

// Container for enum describing unit of radius in proximity.
type ProximityRadiusUnitsEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProximityRadiusUnitsEnum) Reset()         { *m = ProximityRadiusUnitsEnum{} }
func (m *ProximityRadiusUnitsEnum) String() string { return proto.CompactTextString(m) }
func (*ProximityRadiusUnitsEnum) ProtoMessage()    {}
func (*ProximityRadiusUnitsEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_proximity_radius_units_9767cab966548058, []int{0}
}
func (m *ProximityRadiusUnitsEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProximityRadiusUnitsEnum.Unmarshal(m, b)
}
func (m *ProximityRadiusUnitsEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProximityRadiusUnitsEnum.Marshal(b, m, deterministic)
}
func (dst *ProximityRadiusUnitsEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProximityRadiusUnitsEnum.Merge(dst, src)
}
func (m *ProximityRadiusUnitsEnum) XXX_Size() int {
	return xxx_messageInfo_ProximityRadiusUnitsEnum.Size(m)
}
func (m *ProximityRadiusUnitsEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ProximityRadiusUnitsEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ProximityRadiusUnitsEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProximityRadiusUnitsEnum)(nil), "google.ads.googleads.v1.enums.ProximityRadiusUnitsEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.ProximityRadiusUnitsEnum_ProximityRadiusUnits", ProximityRadiusUnitsEnum_ProximityRadiusUnits_name, ProximityRadiusUnitsEnum_ProximityRadiusUnits_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/proximity_radius_units.proto", fileDescriptor_proximity_radius_units_9767cab966548058)
}

var fileDescriptor_proximity_radius_units_9767cab966548058 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xfb, 0x30,
	0x18, 0xc5, 0xff, 0xeb, 0xf8, 0x2b, 0x66, 0xa0, 0xa5, 0x78, 0xa1, 0xe2, 0x2e, 0xb6, 0x07, 0x48,
	0x28, 0xde, 0xc5, 0xab, 0x4e, 0xe3, 0x28, 0xdb, 0xba, 0xb2, 0xb9, 0x0a, 0x52, 0x18, 0xd1, 0x94,
	0x10, 0x5c, 0x93, 0xd2, 0xb4, 0x43, 0x5f, 0xc7, 0x4b, 0x1f, 0xc5, 0x47, 0xf1, 0xc2, 0x67, 0x90,
	0xa6, 0x6b, 0xaf, 0xa6, 0x37, 0xe5, 0xd0, 0xf3, 0x9d, 0x5f, 0xce, 0xf7, 0x01, 0xcc, 0x95, 0xe2,
	0x9b, 0x04, 0x51, 0xa6, 0x51, 0x2d, 0x2b, 0xb5, 0x75, 0x51, 0x22, 0xcb, 0x54, 0xa3, 0x2c, 0x57,
	0xaf, 0x22, 0x15, 0xc5, 0xdb, 0x3a, 0xa7, 0x4c, 0x94, 0x7a, 0x5d, 0x4a, 0x51, 0x68, 0x98, 0xe5,
	0xaa, 0x50, 0x4e, 0xbf, 0x0e, 0x40, 0xca, 0x34, 0x6c, 0xb3, 0x70, 0xeb, 0x42, 0x93, 0xbd, 0xb8,
	0x6c, 0xd0, 0x99, 0x40, 0x54, 0x4a, 0x55, 0xd0, 0x42, 0x28, 0xb9, 0x0b, 0x0f, 0x5f, 0xc0, 0x59,
	0xd8, 0xc0, 0x17, 0x86, 0xbd, 0xaa, 0xd0, 0x44, 0x96, 0xe9, 0x70, 0x0e, 0x4e, 0xf7, 0x79, 0xce,
	0x09, 0xe8, 0xad, 0x82, 0x65, 0x48, 0x6e, 0xfc, 0x3b, 0x9f, 0xdc, 0xda, 0xff, 0x9c, 0x1e, 0x38,
	0x5c, 0x05, 0x93, 0x60, 0xfe, 0x10, 0xd8, 0x1d, 0xe7, 0x08, 0xfc, 0x9f, 0xf9, 0x53, 0xb2, 0xb4,
	0x2d, 0xe7, 0x18, 0x80, 0x89, 0x3f, 0x9d, 0xcf, 0xc8, 0x3d, 0x59, 0x2c, 0xed, 0xee, 0xe8, 0xbb,
	0x03, 0x06, 0xcf, 0x2a, 0x85, 0x7f, 0x16, 0x1e, 0x9d, 0xef, 0x7b, 0x34, 0xac, 0xda, 0x86, 0x9d,
	0xc7, 0xd1, 0x2e, 0xcb, 0xd5, 0x86, 0x4a, 0x0e, 0x55, 0xce, 0x11, 0x4f, 0xa4, 0xd9, 0xa5, 0x39,
	0x5c, 0x26, 0xf4, 0x2f, 0x77, 0xbc, 0x36, 0xdf, 0x77, 0xab, 0x3b, 0xf6, 0xbc, 0x0f, 0xab, 0x3f,
	0xae, 0x51, 0x1e, 0xd3, 0xb0, 0x96, 0x95, 0x8a, 0x5c, 0x58, 0x2d, 0xaf, 0x3f, 0x1b, 0x3f, 0xf6,
	0x98, 0x8e, 0x5b, 0x3f, 0x8e, 0xdc, 0xd8, 0xf8, 0x5f, 0xd6, 0xa0, 0xfe, 0x89, 0xb1, 0xc7, 0x34,
	0xc6, 0xed, 0x04, 0xc6, 0x91, 0x8b, 0xb1, 0x99, 0x79, 0x3a, 0x30, 0xc5, 0xae, 0x7e, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x05, 0xc7, 0x92, 0xa6, 0xdf, 0x01, 0x00, 0x00,
}
