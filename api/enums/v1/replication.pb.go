// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/enums/v1/replication.proto

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

type ReplicationTaskType int32

const (
	REPLICATION_TASK_TYPE_UNSPECIFIED            ReplicationTaskType = 0
	REPLICATION_TASK_TYPE_NAMESPACE_TASK         ReplicationTaskType = 1
	REPLICATION_TASK_TYPE_HISTORY_TASK           ReplicationTaskType = 2
	REPLICATION_TASK_TYPE_SYNC_SHARD_STATUS_TASK ReplicationTaskType = 3
	REPLICATION_TASK_TYPE_SYNC_ACTIVITY_TASK     ReplicationTaskType = 4
	REPLICATION_TASK_TYPE_HISTORY_METADATA_TASK  ReplicationTaskType = 5
	REPLICATION_TASK_TYPE_HISTORY_V2_TASK        ReplicationTaskType = 6
)

var ReplicationTaskType_name = map[int32]string{
	0: "Unspecified",
	1: "NamespaceTask",
	2: "HistoryTask",
	3: "SyncShardStatusTask",
	4: "SyncActivityTask",
	5: "HistoryMetadataTask",
	6: "HistoryV2Task",
}

var ReplicationTaskType_value = map[string]int32{
	"Unspecified":         0,
	"NamespaceTask":       1,
	"HistoryTask":         2,
	"SyncShardStatusTask": 3,
	"SyncActivityTask":    4,
	"HistoryMetadataTask": 5,
	"HistoryV2Task":       6,
}

func (ReplicationTaskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3f4df3039790445d, []int{0}
}

type NamespaceOperation int32

const (
	NAMESPACE_OPERATION_UNSPECIFIED NamespaceOperation = 0
	NAMESPACE_OPERATION_CREATE      NamespaceOperation = 1
	NAMESPACE_OPERATION_UPDATE      NamespaceOperation = 2
)

var NamespaceOperation_name = map[int32]string{
	0: "Unspecified",
	1: "Create",
	2: "Update",
}

var NamespaceOperation_value = map[string]int32{
	"Unspecified": 0,
	"Create":      1,
	"Update":      2,
}

func (NamespaceOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3f4df3039790445d, []int{1}
}

func init() {
	proto.RegisterEnum("temporal.server.api.enums.v1.ReplicationTaskType", ReplicationTaskType_name, ReplicationTaskType_value)
	proto.RegisterEnum("temporal.server.api.enums.v1.NamespaceOperation", NamespaceOperation_name, NamespaceOperation_value)
}

func init() {
	proto.RegisterFile("temporal/server/api/enums/v1/replication.proto", fileDescriptor_3f4df3039790445d)
}

var fileDescriptor_3f4df3039790445d = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x73, 0x55, 0x3b, 0xdc, 0x54, 0xce, 0x4d, 0xe4, 0xc4, 0x3f, 0x95, 0x5a, 0x4b, 0x62,
	0x75, 0x74, 0x3a, 0x93, 0x93, 0x06, 0x6d, 0x12, 0x72, 0xd7, 0x42, 0x1d, 0x0c, 0xb1, 0x1c, 0x12,
	0x6c, 0x9b, 0x23, 0xa9, 0x85, 0x6e, 0x7e, 0x04, 0x3f, 0x86, 0x1f, 0x45, 0x9c, 0x3a, 0x76, 0xb4,
	0xe9, 0xe2, 0xd8, 0x8f, 0x20, 0x26, 0x6a, 0x11, 0x62, 0xb7, 0xe3, 0xde, 0xdf, 0xef, 0xe5, 0xe5,
	0xe1, 0x81, 0xea, 0x50, 0xf4, 0x65, 0x18, 0xf9, 0x3d, 0x2d, 0x16, 0xd1, 0x48, 0x44, 0x9a, 0x2f,
	0x03, 0x4d, 0x0c, 0x1e, 0xfb, 0xb1, 0x36, 0xaa, 0x6b, 0x91, 0x90, 0xbd, 0xa0, 0xeb, 0x0f, 0x83,
	0x70, 0xa0, 0xca, 0x28, 0x1c, 0x86, 0x68, 0xfb, 0x87, 0x57, 0x33, 0x5e, 0xf5, 0x65, 0xa0, 0xa6,
	0xbc, 0x3a, 0xaa, 0x57, 0xdf, 0x0a, 0x70, 0xd3, 0x5d, 0x3a, 0xdc, 0x8f, 0x1f, 0xf8, 0x58, 0x0a,
	0x54, 0x86, 0xbb, 0x2e, 0x75, 0xae, 0x4d, 0x9d, 0x70, 0xd3, 0xb6, 0x3c, 0x4e, 0xd8, 0x95, 0xc7,
	0x3b, 0x0e, 0xf5, 0x5a, 0x16, 0x73, 0xa8, 0x6e, 0x5e, 0x9a, 0xd4, 0x28, 0x29, 0xa8, 0x02, 0x0f,
	0xf2, 0x31, 0x8b, 0x34, 0x29, 0x73, 0x88, 0x4e, 0xd3, 0xbf, 0x12, 0x40, 0x87, 0x70, 0x2f, 0x9f,
	0x6c, 0x98, 0x8c, 0xdb, 0x6e, 0x27, 0xe3, 0x0a, 0xe8, 0x04, 0xd6, 0xf2, 0x39, 0xd6, 0xb1, 0x74,
	0x8f, 0x35, 0x88, 0x6b, 0x78, 0x8c, 0x13, 0xde, 0x62, 0x99, 0xb1, 0x86, 0x6a, 0xb0, 0xb2, 0xc2,
	0x20, 0x3a, 0x37, 0xdb, 0x26, 0xff, 0xde, 0xbf, 0x8e, 0x34, 0x78, 0xbc, 0xfa, 0x8e, 0x26, 0xe5,
	0xc4, 0x20, 0x9c, 0x64, 0xc2, 0x06, 0x3a, 0x82, 0xe5, 0xd5, 0x42, 0xfb, 0x34, 0x43, 0x8b, 0xd5,
	0x31, 0x44, 0x96, 0xdf, 0x17, 0xb1, 0xf4, 0xbb, 0xc2, 0x96, 0x22, 0x4a, 0x23, 0x45, 0xfb, 0x70,
	0x67, 0x99, 0x86, 0xed, 0x50, 0x37, 0x5b, 0xf4, 0x37, 0x48, 0x0c, 0xb7, 0xf2, 0x20, 0xdd, 0xa5,
	0x84, 0xd3, 0x12, 0xf8, 0x6f, 0xde, 0x72, 0x8c, 0xaf, 0x79, 0xe1, 0xe2, 0x76, 0x32, 0xc3, 0xca,
	0x74, 0x86, 0x95, 0xc5, 0x0c, 0x83, 0xa7, 0x04, 0x83, 0x97, 0x04, 0x83, 0xd7, 0x04, 0x83, 0x49,
	0x82, 0xc1, 0x7b, 0x82, 0xc1, 0x47, 0x82, 0x95, 0x45, 0x82, 0xc1, 0xf3, 0x1c, 0x2b, 0x93, 0x39,
	0x56, 0xa6, 0x73, 0xac, 0xdc, 0x54, 0xee, 0xc3, 0xdf, 0x3a, 0xa9, 0x41, 0x98, 0xd7, 0xa8, 0xf3,
	0xf4, 0x71, 0x57, 0x4c, 0xcb, 0x74, 0xf6, 0x19, 0x00, 0x00, 0xff, 0xff, 0xab, 0x9f, 0x42, 0xee,
	0x7e, 0x02, 0x00, 0x00,
}

func (x ReplicationTaskType) String() string {
	s, ok := ReplicationTaskType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x NamespaceOperation) String() string {
	s, ok := NamespaceOperation_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
