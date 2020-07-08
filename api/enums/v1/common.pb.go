// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/enums/v1/common.proto

package enums

import (
	fmt "fmt"
	math "math"
	strconv "strconv"

	proto "github.com/gogo/protobuf/proto"
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

type DeadLetterQueueType int32

const (
	DEAD_LETTER_QUEUE_TYPE_UNSPECIFIED DeadLetterQueueType = 0
	DEAD_LETTER_QUEUE_TYPE_REPLICATION DeadLetterQueueType = 1
	DEAD_LETTER_QUEUE_TYPE_NAMESPACE   DeadLetterQueueType = 2
)

var DeadLetterQueueType_name = map[int32]string{
	0: "Unspecified",
	1: "Replication",
	2: "Namespace",
}

var DeadLetterQueueType_value = map[string]int32{
	"Unspecified": 0,
	"Replication": 1,
	"Namespace":   2,
}

func (DeadLetterQueueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4a3bfa9c01eff6e4, []int{0}
}

type ChecksumFlavor int32

const (
	CHECKSUM_FLAVOR_UNSPECIFIED                   ChecksumFlavor = 0
	CHECKSUM_FLAVOR_IEEE_CRC32_OVER_PROTO3_BINARY ChecksumFlavor = 1
)

var ChecksumFlavor_name = map[int32]string{
	0: "Unspecified",
	1: "IeeeCrc32OverProto3Binary",
}

var ChecksumFlavor_value = map[string]int32{
	"Unspecified":               0,
	"IeeeCrc32OverProto3Binary": 1,
}

func (ChecksumFlavor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4a3bfa9c01eff6e4, []int{1}
}

func init() {
	proto.RegisterEnum("temporal.server.api.enums.v1.DeadLetterQueueType", DeadLetterQueueType_name, DeadLetterQueueType_value)
	proto.RegisterEnum("temporal.server.api.enums.v1.ChecksumFlavor", ChecksumFlavor_name, ChecksumFlavor_value)
}

func init() {
	proto.RegisterFile("temporal/server/api/enums/v1/common.proto", fileDescriptor_4a3bfa9c01eff6e4)
}

var fileDescriptor_4a3bfa9c01eff6e4 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd0, 0x31, 0x4b, 0xf3, 0x40,
	0x1c, 0xc7, 0xf1, 0xdc, 0x33, 0x3c, 0xc3, 0x0d, 0x12, 0xe2, 0xa8, 0x9c, 0x22, 0x22, 0x5a, 0x30,
	0xa1, 0x76, 0x74, 0x4a, 0x2f, 0xff, 0x62, 0x30, 0x4d, 0xd2, 0x6b, 0x52, 0xa8, 0x83, 0x47, 0x6c,
	0x0f, 0x2d, 0x36, 0xbd, 0x90, 0x26, 0x01, 0x37, 0x77, 0x17, 0x5f, 0x86, 0x2f, 0xc5, 0xb1, 0x63,
	0x47, 0x7b, 0x5d, 0x1c, 0xfb, 0x12, 0x84, 0x8a, 0x0e, 0x45, 0xdd, 0x7e, 0xc3, 0x67, 0xf8, 0xf1,
	0xc5, 0x27, 0x85, 0x48, 0x33, 0x99, 0x27, 0x63, 0x6b, 0x2a, 0xf2, 0x4a, 0xe4, 0x56, 0x92, 0x8d,
	0x2c, 0x31, 0x29, 0xd3, 0xa9, 0x55, 0xd5, 0xad, 0x81, 0x4c, 0x53, 0x39, 0x31, 0xb3, 0x5c, 0x16,
	0xd2, 0xd8, 0xfd, 0xa2, 0xe6, 0x27, 0x35, 0x93, 0x6c, 0x64, 0xae, 0xa9, 0x59, 0xd5, 0x6b, 0x4f,
	0x08, 0x6f, 0x3b, 0x22, 0x19, 0x7a, 0xa2, 0x28, 0x44, 0xde, 0x29, 0x45, 0x29, 0xa2, 0x87, 0x4c,
	0x18, 0x47, 0xf8, 0xc0, 0x01, 0xdb, 0xe1, 0x1e, 0x44, 0x11, 0x30, 0xde, 0x89, 0x21, 0x06, 0x1e,
	0xf5, 0x43, 0xe0, 0xb1, 0xdf, 0x0d, 0x81, 0xba, 0x2d, 0x17, 0x1c, 0x5d, 0xfb, 0xc3, 0x31, 0x08,
	0x3d, 0x97, 0xda, 0x91, 0x1b, 0xf8, 0x3a, 0x32, 0x0e, 0xf1, 0xfe, 0x2f, 0xce, 0xb7, 0xdb, 0xd0,
	0x0d, 0x6d, 0x0a, 0xfa, 0xbf, 0xda, 0x10, 0x6f, 0xd1, 0x3b, 0x31, 0xb8, 0x9f, 0x96, 0x69, 0x6b,
	0x9c, 0x54, 0x32, 0x37, 0xf6, 0xf0, 0x0e, 0xbd, 0x00, 0x7a, 0xd9, 0x8d, 0xdb, 0xbc, 0xe5, 0xd9,
	0xbd, 0x80, 0x6d, 0x1c, 0xa8, 0xe3, 0xd3, 0x4d, 0xe0, 0x02, 0x00, 0xa7, 0x8c, 0x36, 0xce, 0x78,
	0xd0, 0x03, 0xc6, 0x43, 0x16, 0x44, 0x41, 0x83, 0x37, 0x5d, 0xdf, 0x66, 0x7d, 0x1d, 0x35, 0xaf,
	0x67, 0x0b, 0xa2, 0xcd, 0x17, 0x44, 0x5b, 0x2d, 0x08, 0x7a, 0x54, 0x04, 0xbd, 0x28, 0x82, 0x5e,
	0x15, 0x41, 0x33, 0x45, 0xd0, 0x9b, 0x22, 0xe8, 0x5d, 0x11, 0x6d, 0xa5, 0x08, 0x7a, 0x5e, 0x12,
	0x6d, 0xb6, 0x24, 0xda, 0x7c, 0x49, 0xb4, 0xab, 0xe3, 0x5b, 0x69, 0x7e, 0xa7, 0x1c, 0xc9, 0x9f,
	0xc2, 0x9f, 0xaf, 0xc7, 0xcd, 0xff, 0x75, 0xf8, 0xc6, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18,
	0x8a, 0xe1, 0xa4, 0xa5, 0x01, 0x00, 0x00,
}

func (x DeadLetterQueueType) String() string {
	s, ok := DeadLetterQueueType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ChecksumFlavor) String() string {
	s, ok := ChecksumFlavor_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
