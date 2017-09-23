// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/milo/api/proto/buildbot.proto

/*
Package milo is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/milo/api/proto/buildbot.proto
	go.chromium.org/luci/milo/api/proto/buildinfo.proto
	go.chromium.org/luci/milo/api/proto/console_git_info.proto

It has these top-level messages:
	MasterRequest
	CompressedMasterJSON
	BuildbotBuildRequest
	BuildbotBuildJSON
	BuildbotBuildsRequest
	BuildbotBuildsJSON
	BuildInfoRequest
	BuildInfoResponse
	ConsoleGitInfo
*/
package milo

import prpc "go.chromium.org/luci/grpc/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// The request containing the name of the master.
type MasterRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// if true, exclude response data that the foundation team is actively trying
	// to deprecate:
	// - slave info
	ExcludeDeprecated bool `protobuf:"varint,10,opt,name=exclude_deprecated,json=excludeDeprecated" json:"exclude_deprecated,omitempty"`
}

func (m *MasterRequest) Reset()                    { *m = MasterRequest{} }
func (m *MasterRequest) String() string            { return proto.CompactTextString(m) }
func (*MasterRequest) ProtoMessage()               {}
func (*MasterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MasterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MasterRequest) GetExcludeDeprecated() bool {
	if m != nil {
		return m.ExcludeDeprecated
	}
	return false
}

// The response message containing master information.
type CompressedMasterJSON struct {
	// Whether the master is internal or not.
	Internal bool `protobuf:"varint,1,opt,name=internal" json:"internal,omitempty"`
	// Timestamp of the freshness of the master data.
	Modified *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=modified" json:"modified,omitempty"`
	// Gzipped json data of the master.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *CompressedMasterJSON) Reset()                    { *m = CompressedMasterJSON{} }
func (m *CompressedMasterJSON) String() string            { return proto.CompactTextString(m) }
func (*CompressedMasterJSON) ProtoMessage()               {}
func (*CompressedMasterJSON) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CompressedMasterJSON) GetInternal() bool {
	if m != nil {
		return m.Internal
	}
	return false
}

func (m *CompressedMasterJSON) GetModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

func (m *CompressedMasterJSON) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// The request for a specific build.
type BuildbotBuildRequest struct {
	Master   string `protobuf:"bytes,1,opt,name=master" json:"master,omitempty"`
	Builder  string `protobuf:"bytes,2,opt,name=builder" json:"builder,omitempty"`
	BuildNum int64  `protobuf:"varint,3,opt,name=build_num,json=buildNum" json:"build_num,omitempty"`
	// if true, exclude response data that the foundation team is actively trying
	// to deprecate:
	// - slave info
	ExcludeDeprecated bool `protobuf:"varint,10,opt,name=exclude_deprecated,json=excludeDeprecated" json:"exclude_deprecated,omitempty"`
}

func (m *BuildbotBuildRequest) Reset()                    { *m = BuildbotBuildRequest{} }
func (m *BuildbotBuildRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildbotBuildRequest) ProtoMessage()               {}
func (*BuildbotBuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BuildbotBuildRequest) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *BuildbotBuildRequest) GetBuilder() string {
	if m != nil {
		return m.Builder
	}
	return ""
}

func (m *BuildbotBuildRequest) GetBuildNum() int64 {
	if m != nil {
		return m.BuildNum
	}
	return 0
}

func (m *BuildbotBuildRequest) GetExcludeDeprecated() bool {
	if m != nil {
		return m.ExcludeDeprecated
	}
	return false
}

// The response message for a specific build.
type BuildbotBuildJSON struct {
	// Json data of the build.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BuildbotBuildJSON) Reset()                    { *m = BuildbotBuildJSON{} }
func (m *BuildbotBuildJSON) String() string            { return proto.CompactTextString(m) }
func (*BuildbotBuildJSON) ProtoMessage()               {}
func (*BuildbotBuildJSON) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BuildbotBuildJSON) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// The request for multiple build on a builder.
type BuildbotBuildsRequest struct {
	Master  string `protobuf:"bytes,1,opt,name=master" json:"master,omitempty"`
	Builder string `protobuf:"bytes,2,opt,name=builder" json:"builder,omitempty"`
	// Limit to the number of builds to return (default: 20).
	Limit int32 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	// Include ongoing builds (default: false).
	IncludeCurrent bool `protobuf:"varint,4,opt,name=include_current,json=includeCurrent" json:"include_current,omitempty"`
	// Return builds starting from this cursor.
	Cursor string `protobuf:"bytes,5,opt,name=cursor" json:"cursor,omitempty"`
}

func (m *BuildbotBuildsRequest) Reset()                    { *m = BuildbotBuildsRequest{} }
func (m *BuildbotBuildsRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildbotBuildsRequest) ProtoMessage()               {}
func (*BuildbotBuildsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BuildbotBuildsRequest) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *BuildbotBuildsRequest) GetBuilder() string {
	if m != nil {
		return m.Builder
	}
	return ""
}

func (m *BuildbotBuildsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *BuildbotBuildsRequest) GetIncludeCurrent() bool {
	if m != nil {
		return m.IncludeCurrent
	}
	return false
}

func (m *BuildbotBuildsRequest) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

// The response message for multiple builds in a builder.
type BuildbotBuildsJSON struct {
	// builds is the list of builds resulting from the builds request.
	Builds []*BuildbotBuildJSON `protobuf:"bytes,1,rep,name=builds" json:"builds,omitempty"`
	// The cursor to the next request from, if the exact same query was given.
	Cursor string `protobuf:"bytes,2,opt,name=cursor" json:"cursor,omitempty"`
}

func (m *BuildbotBuildsJSON) Reset()                    { *m = BuildbotBuildsJSON{} }
func (m *BuildbotBuildsJSON) String() string            { return proto.CompactTextString(m) }
func (*BuildbotBuildsJSON) ProtoMessage()               {}
func (*BuildbotBuildsJSON) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BuildbotBuildsJSON) GetBuilds() []*BuildbotBuildJSON {
	if m != nil {
		return m.Builds
	}
	return nil
}

func (m *BuildbotBuildsJSON) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func init() {
	proto.RegisterType((*MasterRequest)(nil), "milo.MasterRequest")
	proto.RegisterType((*CompressedMasterJSON)(nil), "milo.CompressedMasterJSON")
	proto.RegisterType((*BuildbotBuildRequest)(nil), "milo.BuildbotBuildRequest")
	proto.RegisterType((*BuildbotBuildJSON)(nil), "milo.BuildbotBuildJSON")
	proto.RegisterType((*BuildbotBuildsRequest)(nil), "milo.BuildbotBuildsRequest")
	proto.RegisterType((*BuildbotBuildsJSON)(nil), "milo.BuildbotBuildsJSON")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Buildbot service

type BuildbotClient interface {
	GetCompressedMasterJSON(ctx context.Context, in *MasterRequest, opts ...grpc.CallOption) (*CompressedMasterJSON, error)
	GetBuildbotBuildJSON(ctx context.Context, in *BuildbotBuildRequest, opts ...grpc.CallOption) (*BuildbotBuildJSON, error)
	GetBuildbotBuildsJSON(ctx context.Context, in *BuildbotBuildsRequest, opts ...grpc.CallOption) (*BuildbotBuildsJSON, error)
}
type buildbotPRPCClient struct {
	client *prpc.Client
}

func NewBuildbotPRPCClient(client *prpc.Client) BuildbotClient {
	return &buildbotPRPCClient{client}
}

func (c *buildbotPRPCClient) GetCompressedMasterJSON(ctx context.Context, in *MasterRequest, opts ...grpc.CallOption) (*CompressedMasterJSON, error) {
	out := new(CompressedMasterJSON)
	err := c.client.Call(ctx, "milo.Buildbot", "GetCompressedMasterJSON", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildbotPRPCClient) GetBuildbotBuildJSON(ctx context.Context, in *BuildbotBuildRequest, opts ...grpc.CallOption) (*BuildbotBuildJSON, error) {
	out := new(BuildbotBuildJSON)
	err := c.client.Call(ctx, "milo.Buildbot", "GetBuildbotBuildJSON", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildbotPRPCClient) GetBuildbotBuildsJSON(ctx context.Context, in *BuildbotBuildsRequest, opts ...grpc.CallOption) (*BuildbotBuildsJSON, error) {
	out := new(BuildbotBuildsJSON)
	err := c.client.Call(ctx, "milo.Buildbot", "GetBuildbotBuildsJSON", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type buildbotClient struct {
	cc *grpc.ClientConn
}

func NewBuildbotClient(cc *grpc.ClientConn) BuildbotClient {
	return &buildbotClient{cc}
}

func (c *buildbotClient) GetCompressedMasterJSON(ctx context.Context, in *MasterRequest, opts ...grpc.CallOption) (*CompressedMasterJSON, error) {
	out := new(CompressedMasterJSON)
	err := grpc.Invoke(ctx, "/milo.Buildbot/GetCompressedMasterJSON", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildbotClient) GetBuildbotBuildJSON(ctx context.Context, in *BuildbotBuildRequest, opts ...grpc.CallOption) (*BuildbotBuildJSON, error) {
	out := new(BuildbotBuildJSON)
	err := grpc.Invoke(ctx, "/milo.Buildbot/GetBuildbotBuildJSON", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildbotClient) GetBuildbotBuildsJSON(ctx context.Context, in *BuildbotBuildsRequest, opts ...grpc.CallOption) (*BuildbotBuildsJSON, error) {
	out := new(BuildbotBuildsJSON)
	err := grpc.Invoke(ctx, "/milo.Buildbot/GetBuildbotBuildsJSON", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Buildbot service

type BuildbotServer interface {
	GetCompressedMasterJSON(context.Context, *MasterRequest) (*CompressedMasterJSON, error)
	GetBuildbotBuildJSON(context.Context, *BuildbotBuildRequest) (*BuildbotBuildJSON, error)
	GetBuildbotBuildsJSON(context.Context, *BuildbotBuildsRequest) (*BuildbotBuildsJSON, error)
}

func RegisterBuildbotServer(s prpc.Registrar, srv BuildbotServer) {
	s.RegisterService(&_Buildbot_serviceDesc, srv)
}

func _Buildbot_GetCompressedMasterJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildbotServer).GetCompressedMasterJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milo.Buildbot/GetCompressedMasterJSON",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildbotServer).GetCompressedMasterJSON(ctx, req.(*MasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Buildbot_GetBuildbotBuildJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildbotBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildbotServer).GetBuildbotBuildJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milo.Buildbot/GetBuildbotBuildJSON",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildbotServer).GetBuildbotBuildJSON(ctx, req.(*BuildbotBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Buildbot_GetBuildbotBuildsJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildbotBuildsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildbotServer).GetBuildbotBuildsJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milo.Buildbot/GetBuildbotBuildsJSON",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildbotServer).GetBuildbotBuildsJSON(ctx, req.(*BuildbotBuildsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Buildbot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milo.Buildbot",
	HandlerType: (*BuildbotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompressedMasterJSON",
			Handler:    _Buildbot_GetCompressedMasterJSON_Handler,
		},
		{
			MethodName: "GetBuildbotBuildJSON",
			Handler:    _Buildbot_GetBuildbotBuildJSON_Handler,
		},
		{
			MethodName: "GetBuildbotBuildsJSON",
			Handler:    _Buildbot_GetBuildbotBuildsJSON_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/milo/api/proto/buildbot.proto",
}

func init() { proto.RegisterFile("go.chromium.org/luci/milo/api/proto/buildbot.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0xc7, 0xeb, 0xee, 0x07, 0xd9, 0x29, 0x1f, 0xaa, 0xd9, 0x52, 0x2b, 0x3d, 0x10, 0xe5, 0xd2,
	0x5c, 0x48, 0xa4, 0x45, 0xe2, 0x01, 0x28, 0x52, 0x25, 0x44, 0x0b, 0x32, 0x5c, 0x51, 0x95, 0x4d,
	0xa6, 0x8b, 0xa5, 0x38, 0x0e, 0x8e, 0x2d, 0x71, 0xe2, 0x29, 0x78, 0x00, 0x5e, 0x93, 0x1b, 0x5a,
	0x3b, 0x59, 0x6d, 0xd8, 0x70, 0x80, 0x53, 0x3c, 0x5f, 0xff, 0xf9, 0xcd, 0x64, 0x60, 0xb5, 0x51,
	0x69, 0xf1, 0x45, 0x2b, 0x29, 0xac, 0x4c, 0x95, 0xde, 0x64, 0x95, 0x2d, 0x44, 0x26, 0x45, 0xa5,
	0xb2, 0xbc, 0x11, 0x59, 0xa3, 0x95, 0x51, 0xd9, 0xda, 0x8a, 0xaa, 0x5c, 0x2b, 0x93, 0x3a, 0x93,
	0x4e, 0xb7, 0xe1, 0xf0, 0xf9, 0x46, 0xa9, 0x4d, 0x85, 0x3e, 0x65, 0x6d, 0xef, 0x33, 0x23, 0x24,
	0xb6, 0x26, 0x97, 0x8d, 0x4f, 0x8b, 0x39, 0x3c, 0xba, 0xc9, 0x5b, 0x83, 0x9a, 0xe3, 0x57, 0x8b,
	0xad, 0xa1, 0x14, 0xa6, 0x75, 0x2e, 0x91, 0x91, 0x88, 0x24, 0x0b, 0xee, 0xde, 0xf4, 0x05, 0x50,
	0xfc, 0x56, 0x54, 0xb6, 0xc4, 0xbb, 0x12, 0x1b, 0x8d, 0x45, 0x6e, 0xb0, 0x64, 0x10, 0x91, 0x24,
	0xe0, 0xa7, 0x5d, 0xe4, 0xcd, 0x2e, 0x10, 0x7f, 0x87, 0xe5, 0x95, 0x92, 0x8d, 0xc6, 0xb6, 0xc5,
	0xd2, 0xab, 0xbf, 0xfd, 0xf8, 0xfe, 0x96, 0x86, 0x10, 0x88, 0xda, 0xa0, 0xae, 0xf3, 0xca, 0xc9,
	0x07, 0x7c, 0x67, 0xd3, 0x57, 0x10, 0x48, 0x55, 0x8a, 0x7b, 0x81, 0x25, 0x3b, 0x8e, 0x48, 0x72,
	0xb2, 0x0a, 0x53, 0xcf, 0x9e, 0xf6, 0xec, 0xe9, 0xa7, 0x9e, 0x9d, 0xef, 0x72, 0xb7, 0xb8, 0x65,
	0x6e, 0x72, 0x36, 0x89, 0x48, 0xf2, 0x90, 0xbb, 0x77, 0xfc, 0x83, 0xc0, 0xf2, 0x75, 0xb7, 0x0d,
	0xf7, 0xed, 0x67, 0x7b, 0x06, 0x73, 0xe9, 0x70, 0xba, 0xe9, 0x3a, 0x8b, 0x32, 0x78, 0xe0, 0xb6,
	0x87, 0xda, 0xf5, 0x5e, 0xf0, 0xde, 0xa4, 0x17, 0xb0, 0x70, 0xcf, 0xbb, 0xda, 0x4a, 0xd7, 0x63,
	0xc2, 0x03, 0xe7, 0xb8, 0xb5, 0xf2, 0x5f, 0xd7, 0x72, 0x09, 0xa7, 0x03, 0x2a, 0xb7, 0x93, 0x9e,
	0x9f, 0xec, 0xf1, 0xff, 0x24, 0x70, 0x36, 0xc8, 0x6c, 0xff, 0x7f, 0x80, 0x25, 0xcc, 0x2a, 0x21,
	0x85, 0x71, 0xf0, 0x33, 0xee, 0x0d, 0x7a, 0x09, 0x4f, 0x44, 0xed, 0xc9, 0x0b, 0xab, 0x35, 0xd6,
	0x86, 0x4d, 0x1d, 0xf6, 0xe3, 0xce, 0x7d, 0xe5, 0xbd, 0xdb, 0x86, 0x85, 0xd5, 0xad, 0xd2, 0x6c,
	0xe6, 0x1b, 0x7a, 0x2b, 0xfe, 0x0c, 0x74, 0x48, 0xe8, 0x86, 0xc9, 0x60, 0xee, 0xfa, 0xb6, 0x8c,
	0x44, 0x93, 0xe4, 0x64, 0x75, 0x9e, 0x6e, 0x8f, 0x30, 0x3d, 0x98, 0x9a, 0x77, 0x69, 0x7b, 0xf2,
	0xc7, 0xfb, 0xf2, 0xab, 0x5f, 0x04, 0x82, 0xbe, 0x8a, 0xbe, 0x83, 0xf3, 0x6b, 0x34, 0xa3, 0x17,
	0xf5, 0xd4, 0x37, 0x18, 0x5c, 0x70, 0x18, 0x7a, 0xe7, 0x58, 0x41, 0x7c, 0x44, 0x6f, 0x60, 0x79,
	0x8d, 0xe6, 0xf0, 0x47, 0x84, 0x23, 0xac, 0xbd, 0xe2, 0xdf, 0xe6, 0x88, 0x8f, 0xe8, 0x07, 0x38,
	0xfb, 0x53, 0xce, 0xef, 0xe2, 0x62, 0xa4, 0xa6, 0xff, 0x8f, 0x21, 0x1b, 0x0b, 0x7a, 0xc5, 0xf5,
	0xdc, 0xdd, 0xfb, 0xcb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x18, 0x66, 0x3b, 0xf9, 0xf5, 0x03,
	0x00, 0x00,
}
