// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/appengine/mapper/messages.proto

package mapper

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// State of a job or one of its shards.
type State int32

const (
	State_STATE_UNSPECIFIED State = 0
	State_STARTING          State = 1
	State_RUNNING           State = 2
	State_SUCCESS           State = 3
	State_FAIL              State = 4
)

var State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "STARTING",
	2: "RUNNING",
	3: "SUCCESS",
	4: "FAIL",
}

var State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"STARTING":          1,
	"RUNNING":           2,
	"SUCCESS":           3,
	"FAIL":              4,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_009da4539d0ef669, []int{0}
}

func init() {
	proto.RegisterEnum("appengine.mapper.messages.State", State_name, State_value)
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/appengine/mapper/messages.proto", fileDescriptor_009da4539d0ef669)
}

var fileDescriptor_009da4539d0ef669 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x4f, 0x2c, 0x28, 0x48, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0xd5, 0xcf, 0x05, 0x31, 0x8b, 0xf4,
	0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x8b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x24,
	0xe1, 0x0a, 0xf4, 0x20, 0x0a, 0xf4, 0x60, 0x0a, 0xb4, 0x02, 0xb8, 0x58, 0x83, 0x4b, 0x12, 0x4b,
	0x52, 0x85, 0x44, 0xb9, 0x04, 0x83, 0x43, 0x1c, 0x43, 0x5c, 0xe3, 0x43, 0xfd, 0x82, 0x03, 0x5c,
	0x9d, 0x3d, 0xdd, 0x3c, 0x5d, 0x5d, 0x04, 0x18, 0x84, 0x78, 0xb8, 0x38, 0x82, 0x43, 0x1c, 0x83,
	0x42, 0x3c, 0xfd, 0xdc, 0x05, 0x18, 0x85, 0xb8, 0xb9, 0xd8, 0x83, 0x42, 0xfd, 0xfc, 0x40, 0x1c,
	0x26, 0x10, 0x27, 0x38, 0xd4, 0xd9, 0xd9, 0x35, 0x38, 0x58, 0x80, 0x59, 0x88, 0x83, 0x8b, 0xc5,
	0xcd, 0xd1, 0xd3, 0x47, 0x80, 0xc5, 0x89, 0x23, 0x8a, 0x0d, 0x62, 0x49, 0x12, 0x1b, 0xd8, 0x76,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0x2b, 0x0b, 0xc2, 0xb5, 0x00, 0x00, 0x00,
}
