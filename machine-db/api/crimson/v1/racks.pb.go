// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/racks.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ListRacksRequest is a request to retrieve racks.
type ListRacksRequest struct {
	// The names of racks to retrieve.
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
	// The datacenters to filter retrieved racks on.
	Datacenters []string `protobuf:"bytes,2,rep,name=datacenters" json:"datacenters,omitempty"`
}

func (m *ListRacksRequest) Reset()                    { *m = ListRacksRequest{} }
func (m *ListRacksRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRacksRequest) ProtoMessage()               {}
func (*ListRacksRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ListRacksRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ListRacksRequest) GetDatacenters() []string {
	if m != nil {
		return m.Datacenters
	}
	return nil
}

// Rack describes a rack.
type Rack struct {
	// The name of this rack. Uniquely identifies this rack.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this rack.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// The datacenter this rack belongs to.
	Datacenter string `protobuf:"bytes,3,opt,name=datacenter" json:"datacenter,omitempty"`
}

func (m *Rack) Reset()                    { *m = Rack{} }
func (m *Rack) String() string            { return proto.CompactTextString(m) }
func (*Rack) ProtoMessage()               {}
func (*Rack) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *Rack) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rack) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Rack) GetDatacenter() string {
	if m != nil {
		return m.Datacenter
	}
	return ""
}

// ListRacksResponse is a response to a request to retrieve racks.
type ListRacksResponse struct {
	// The racks matching the request.
	Racks []*Rack `protobuf:"bytes,1,rep,name=racks" json:"racks,omitempty"`
}

func (m *ListRacksResponse) Reset()                    { *m = ListRacksResponse{} }
func (m *ListRacksResponse) String() string            { return proto.CompactTextString(m) }
func (*ListRacksResponse) ProtoMessage()               {}
func (*ListRacksResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *ListRacksResponse) GetRacks() []*Rack {
	if m != nil {
		return m.Racks
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRacksRequest)(nil), "crimson.ListRacksRequest")
	proto.RegisterType((*Rack)(nil), "crimson.Rack")
	proto.RegisterType((*ListRacksResponse)(nil), "crimson.ListRacksResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/racks.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0x7e, 0x28, 0x9d, 0x22, 0x68, 0xf0, 0x90, 0x93, 0x2c, 0xeb, 0xa5, 0x17, 0x13,
	0xd4, 0x8b, 0xf8, 0x13, 0xc4, 0x53, 0xce, 0x5e, 0xd2, 0x34, 0xb4, 0x83, 0x26, 0xb3, 0x66, 0xb2,
	0xfe, 0x7e, 0xd9, 0xe9, 0x42, 0xf7, 0x96, 0x3c, 0xcf, 0xcc, 0xf0, 0xbe, 0xf0, 0x7e, 0x24, 0x13,
	0x4e, 0x85, 0x12, 0x0e, 0xc9, 0x50, 0x39, 0xda, 0x9f, 0x21, 0xa0, 0x4d, 0x3e, 0x9c, 0x30, 0xc7,
	0xa7, 0xc3, 0xde, 0xfa, 0x1e, 0x6d, 0x28, 0x98, 0x98, 0xb2, 0xfd, 0x7b, 0xb6, 0xc5, 0x87, 0x6f,
	0x36, 0x7d, 0xa1, 0x4a, 0xea, 0x7a, 0xe2, 0xdd, 0x07, 0xdc, 0x7e, 0x22, 0x57, 0x37, 0x3a, 0x17,
	0x7f, 0x87, 0xc8, 0x55, 0xdd, 0xc3, 0x3a, 0xfb, 0x14, 0x59, 0x37, 0xed, 0x72, 0xb7, 0x71, 0xe7,
	0x8f, 0x6a, 0x61, 0x7b, 0xf0, 0xd5, 0x87, 0x98, 0x6b, 0x2c, 0xac, 0x17, 0xe2, 0xe6, 0xa8, 0xfb,
	0x82, 0xd5, 0x78, 0x47, 0x29, 0x58, 0x8d, 0x2b, 0xba, 0x69, 0x9b, 0xdd, 0xc6, 0xc9, 0x5b, 0xb6,
	0x23, 0x87, 0x82, 0x7d, 0x45, 0xca, 0x7a, 0x21, 0x6a, 0x8e, 0xd4, 0x03, 0xc0, 0xe5, 0x98, 0x5e,
	0xca, 0xc0, 0x8c, 0x74, 0x6f, 0x70, 0x37, 0x4b, 0xca, 0x3d, 0x65, 0x8e, 0xea, 0x11, 0xd6, 0x52,
	0x4b, 0xa2, 0x6e, 0x5f, 0x6e, 0xcc, 0xd4, 0xcb, 0x8c, 0x63, 0xee, 0xec, 0xf6, 0x57, 0xd2, 0xf9,
	0xf5, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xa9, 0x57, 0x7d, 0x31, 0x01, 0x00, 0x00,
}
