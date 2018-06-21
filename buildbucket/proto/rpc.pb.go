// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/rpc.proto

package buildbucketpb

import // import "go.chromium.org/luci/buildbucket/proto"
prpc "go.chromium.org/luci/grpc/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import status "google.golang.org/genproto/googleapis/rpc/status"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

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

// A request message for GetBuild rpc.
type GetBuildRequest struct {
	// Build id.
	// Mutually exclusive with builder and number.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Builder of the build.
	// Requires number. Mutually exclusive with id.
	Builder *BuilderID `protobuf:"bytes,2,opt,name=builder" json:"builder,omitempty"`
	// Build number.
	// Requires builder. Mutually exclusive with id.
	BuildNumber int32 `protobuf:"varint,3,opt,name=build_number,json=buildNumber" json:"build_number,omitempty"`
	// Fields to include in the response.
	// If not set, the default mask is used, see Build message comments for the
	// list of fields returned by default.
	//
	// Supports advanced semantics, see
	// https://chromium.googlesource.com/infra/luci/luci-py/+/f9ae69a37c4bdd0e08a8b0f7e123f6e403e774eb/appengine/components/components/protoutil/field_masks.py#7
	// In particular, if the client needs only some output properties, they
	// can be requested with paths "output.properties.fields.foo".
	Fields               *field_mask.FieldMask `protobuf:"bytes,100,opt,name=fields" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetBuildRequest) Reset()         { *m = GetBuildRequest{} }
func (m *GetBuildRequest) String() string { return proto.CompactTextString(m) }
func (*GetBuildRequest) ProtoMessage()    {}
func (*GetBuildRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ed8cc71e9cf76a07, []int{0}
}
func (m *GetBuildRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBuildRequest.Unmarshal(m, b)
}
func (m *GetBuildRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBuildRequest.Marshal(b, m, deterministic)
}
func (dst *GetBuildRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBuildRequest.Merge(dst, src)
}
func (m *GetBuildRequest) XXX_Size() int {
	return xxx_messageInfo_GetBuildRequest.Size(m)
}
func (m *GetBuildRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBuildRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBuildRequest proto.InternalMessageInfo

func (m *GetBuildRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetBuildRequest) GetBuilder() *BuilderID {
	if m != nil {
		return m.Builder
	}
	return nil
}

func (m *GetBuildRequest) GetBuildNumber() int32 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

func (m *GetBuildRequest) GetFields() *field_mask.FieldMask {
	if m != nil {
		return m.Fields
	}
	return nil
}

// A request message for SearchBuilds rpc.
type SearchBuildsRequest struct {
	// Returned builds must satisfy this predicate. Required.
	Predicate *BuildPredicate `protobuf:"bytes,1,opt,name=predicate" json:"predicate,omitempty"`
	// Fields to include in the response, see GetBuildRequest.fields.
	// Note that this applies to the response, not each build, so e.g. steps must
	// be requested with a path "builds.*.steps".
	Fields *field_mask.FieldMask `protobuf:"bytes,100,opt,name=fields" json:"fields,omitempty"`
	// Number of builds to return.
	// Any value >100 is interpreted as 100.
	PageSize int32 `protobuf:"varint,101,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Value of SearchBuildsResponse.next_page_token from the previous response.
	// Use it to continue searching.
	PageToken            string   `protobuf:"bytes,102,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchBuildsRequest) Reset()         { *m = SearchBuildsRequest{} }
func (m *SearchBuildsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchBuildsRequest) ProtoMessage()    {}
func (*SearchBuildsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ed8cc71e9cf76a07, []int{1}
}
func (m *SearchBuildsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchBuildsRequest.Unmarshal(m, b)
}
func (m *SearchBuildsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchBuildsRequest.Marshal(b, m, deterministic)
}
func (dst *SearchBuildsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchBuildsRequest.Merge(dst, src)
}
func (m *SearchBuildsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchBuildsRequest.Size(m)
}
func (m *SearchBuildsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchBuildsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchBuildsRequest proto.InternalMessageInfo

func (m *SearchBuildsRequest) GetPredicate() *BuildPredicate {
	if m != nil {
		return m.Predicate
	}
	return nil
}

func (m *SearchBuildsRequest) GetFields() *field_mask.FieldMask {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *SearchBuildsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchBuildsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// A response message for SearchBuilds rpc.
type SearchBuildsResponse struct {
	// Search results.
	//
	// Ordered by build id, descending. Ids are monotonically decreasing, so in
	// other words the order is newest-to-oldest.
	Builds []*Build `protobuf:"bytes,1,rep,name=builds" json:"builds,omitempty"`
	// Value for SearchBuildsRequest.page_token to continue searching.
	NextPageToken        string   `protobuf:"bytes,100,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchBuildsResponse) Reset()         { *m = SearchBuildsResponse{} }
func (m *SearchBuildsResponse) String() string { return proto.CompactTextString(m) }
func (*SearchBuildsResponse) ProtoMessage()    {}
func (*SearchBuildsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ed8cc71e9cf76a07, []int{2}
}
func (m *SearchBuildsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchBuildsResponse.Unmarshal(m, b)
}
func (m *SearchBuildsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchBuildsResponse.Marshal(b, m, deterministic)
}
func (dst *SearchBuildsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchBuildsResponse.Merge(dst, src)
}
func (m *SearchBuildsResponse) XXX_Size() int {
	return xxx_messageInfo_SearchBuildsResponse.Size(m)
}
func (m *SearchBuildsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchBuildsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchBuildsResponse proto.InternalMessageInfo

func (m *SearchBuildsResponse) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

func (m *SearchBuildsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// A request message for Batch rpc.
type BatchRequest struct {
	// Requests to execute in a single batch.
	//
	// Requests related to same build are coupled.
	// Mutation requests are executed transactionally, before read-only requests.
	Requests             []*BatchRequest_Request `protobuf:"bytes,1,rep,name=requests" json:"requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *BatchRequest) Reset()         { *m = BatchRequest{} }
func (m *BatchRequest) String() string { return proto.CompactTextString(m) }
func (*BatchRequest) ProtoMessage()    {}
func (*BatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ed8cc71e9cf76a07, []int{3}
}
func (m *BatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchRequest.Unmarshal(m, b)
}
func (m *BatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchRequest.Marshal(b, m, deterministic)
}
func (dst *BatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchRequest.Merge(dst, src)
}
func (m *BatchRequest) XXX_Size() int {
	return xxx_messageInfo_BatchRequest.Size(m)
}
func (m *BatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchRequest proto.InternalMessageInfo

func (m *BatchRequest) GetRequests() []*BatchRequest_Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

// One request in a batch.
type BatchRequest_Request struct {
	// Types that are valid to be assigned to Request:
	//	*BatchRequest_Request_GetBuild
	//	*BatchRequest_Request_SearchBuilds
	Request              isBatchRequest_Request_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *BatchRequest_Request) Reset()         { *m = BatchRequest_Request{} }
func (m *BatchRequest_Request) String() string { return proto.CompactTextString(m) }
func (*BatchRequest_Request) ProtoMessage()    {}
func (*BatchRequest_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ed8cc71e9cf76a07, []int{3, 0}
}
func (m *BatchRequest_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchRequest_Request.Unmarshal(m, b)
}
func (m *BatchRequest_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchRequest_Request.Marshal(b, m, deterministic)
}
func (dst *BatchRequest_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchRequest_Request.Merge(dst, src)
}
func (m *BatchRequest_Request) XXX_Size() int {
	return xxx_messageInfo_BatchRequest_Request.Size(m)
}
func (m *BatchRequest_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchRequest_Request.DiscardUnknown(m)
}

var xxx_messageInfo_BatchRequest_Request proto.InternalMessageInfo

type isBatchRequest_Request_Request interface {
	isBatchRequest_Request_Request()
}

type BatchRequest_Request_GetBuild struct {
	GetBuild *GetBuildRequest `protobuf:"bytes,1,opt,name=get_build,json=getBuild,oneof"`
}
type BatchRequest_Request_SearchBuilds struct {
	SearchBuilds *SearchBuildsRequest `protobuf:"bytes,2,opt,name=search_builds,json=searchBuilds,oneof"`
}

func (*BatchRequest_Request_GetBuild) isBatchRequest_Request_Request()     {}
func (*BatchRequest_Request_SearchBuilds) isBatchRequest_Request_Request() {}

func (m *BatchRequest_Request) GetRequest() isBatchRequest_Request_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *BatchRequest_Request) GetGetBuild() *GetBuildRequest {
	if x, ok := m.GetRequest().(*BatchRequest_Request_GetBuild); ok {
		return x.GetBuild
	}
	return nil
}

func (m *BatchRequest_Request) GetSearchBuilds() *SearchBuildsRequest {
	if x, ok := m.GetRequest().(*BatchRequest_Request_SearchBuilds); ok {
		return x.SearchBuilds
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BatchRequest_Request) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BatchRequest_Request_OneofMarshaler, _BatchRequest_Request_OneofUnmarshaler, _BatchRequest_Request_OneofSizer, []interface{}{
		(*BatchRequest_Request_GetBuild)(nil),
		(*BatchRequest_Request_SearchBuilds)(nil),
	}
}

func _BatchRequest_Request_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BatchRequest_Request)
	// request
	switch x := m.Request.(type) {
	case *BatchRequest_Request_GetBuild:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GetBuild); err != nil {
			return err
		}
	case *BatchRequest_Request_SearchBuilds:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SearchBuilds); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BatchRequest_Request.Request has unexpected type %T", x)
	}
	return nil
}

func _BatchRequest_Request_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BatchRequest_Request)
	switch tag {
	case 1: // request.get_build
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GetBuildRequest)
		err := b.DecodeMessage(msg)
		m.Request = &BatchRequest_Request_GetBuild{msg}
		return true, err
	case 2: // request.search_builds
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SearchBuildsRequest)
		err := b.DecodeMessage(msg)
		m.Request = &BatchRequest_Request_SearchBuilds{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BatchRequest_Request_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BatchRequest_Request)
	// request
	switch x := m.Request.(type) {
	case *BatchRequest_Request_GetBuild:
		s := proto.Size(x.GetBuild)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BatchRequest_Request_SearchBuilds:
		s := proto.Size(x.SearchBuilds)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A response message for Batch rpc.
type BatchResponse struct {
	// Responses in the same order as BatchRequest.requests.
	Responses            []*BatchResponse_Response `protobuf:"bytes,1,rep,name=responses" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BatchResponse) Reset()         { *m = BatchResponse{} }
func (m *BatchResponse) String() string { return proto.CompactTextString(m) }
func (*BatchResponse) ProtoMessage()    {}
func (*BatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ed8cc71e9cf76a07, []int{4}
}
func (m *BatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchResponse.Unmarshal(m, b)
}
func (m *BatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchResponse.Marshal(b, m, deterministic)
}
func (dst *BatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchResponse.Merge(dst, src)
}
func (m *BatchResponse) XXX_Size() int {
	return xxx_messageInfo_BatchResponse.Size(m)
}
func (m *BatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchResponse proto.InternalMessageInfo

func (m *BatchResponse) GetResponses() []*BatchResponse_Response {
	if m != nil {
		return m.Responses
	}
	return nil
}

// Response a BatchRequest.Response.
type BatchResponse_Response struct {
	// Types that are valid to be assigned to Response:
	//	*BatchResponse_Response_GetBuild
	//	*BatchResponse_Response_SearchBuilds
	//	*BatchResponse_Response_Error
	Response             isBatchResponse_Response_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *BatchResponse_Response) Reset()         { *m = BatchResponse_Response{} }
func (m *BatchResponse_Response) String() string { return proto.CompactTextString(m) }
func (*BatchResponse_Response) ProtoMessage()    {}
func (*BatchResponse_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ed8cc71e9cf76a07, []int{4, 0}
}
func (m *BatchResponse_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchResponse_Response.Unmarshal(m, b)
}
func (m *BatchResponse_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchResponse_Response.Marshal(b, m, deterministic)
}
func (dst *BatchResponse_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchResponse_Response.Merge(dst, src)
}
func (m *BatchResponse_Response) XXX_Size() int {
	return xxx_messageInfo_BatchResponse_Response.Size(m)
}
func (m *BatchResponse_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchResponse_Response.DiscardUnknown(m)
}

var xxx_messageInfo_BatchResponse_Response proto.InternalMessageInfo

type isBatchResponse_Response_Response interface {
	isBatchResponse_Response_Response()
}

type BatchResponse_Response_GetBuild struct {
	GetBuild *Build `protobuf:"bytes,1,opt,name=get_build,json=getBuild,oneof"`
}
type BatchResponse_Response_SearchBuilds struct {
	SearchBuilds *SearchBuildsResponse `protobuf:"bytes,2,opt,name=search_builds,json=searchBuilds,oneof"`
}
type BatchResponse_Response_Error struct {
	Error *status.Status `protobuf:"bytes,100,opt,name=error,oneof"`
}

func (*BatchResponse_Response_GetBuild) isBatchResponse_Response_Response()     {}
func (*BatchResponse_Response_SearchBuilds) isBatchResponse_Response_Response() {}
func (*BatchResponse_Response_Error) isBatchResponse_Response_Response()        {}

func (m *BatchResponse_Response) GetResponse() isBatchResponse_Response_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *BatchResponse_Response) GetGetBuild() *Build {
	if x, ok := m.GetResponse().(*BatchResponse_Response_GetBuild); ok {
		return x.GetBuild
	}
	return nil
}

func (m *BatchResponse_Response) GetSearchBuilds() *SearchBuildsResponse {
	if x, ok := m.GetResponse().(*BatchResponse_Response_SearchBuilds); ok {
		return x.SearchBuilds
	}
	return nil
}

func (m *BatchResponse_Response) GetError() *status.Status {
	if x, ok := m.GetResponse().(*BatchResponse_Response_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BatchResponse_Response) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BatchResponse_Response_OneofMarshaler, _BatchResponse_Response_OneofUnmarshaler, _BatchResponse_Response_OneofSizer, []interface{}{
		(*BatchResponse_Response_GetBuild)(nil),
		(*BatchResponse_Response_SearchBuilds)(nil),
		(*BatchResponse_Response_Error)(nil),
	}
}

func _BatchResponse_Response_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BatchResponse_Response)
	// response
	switch x := m.Response.(type) {
	case *BatchResponse_Response_GetBuild:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GetBuild); err != nil {
			return err
		}
	case *BatchResponse_Response_SearchBuilds:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SearchBuilds); err != nil {
			return err
		}
	case *BatchResponse_Response_Error:
		b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BatchResponse_Response.Response has unexpected type %T", x)
	}
	return nil
}

func _BatchResponse_Response_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BatchResponse_Response)
	switch tag {
	case 1: // response.get_build
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Build)
		err := b.DecodeMessage(msg)
		m.Response = &BatchResponse_Response_GetBuild{msg}
		return true, err
	case 2: // response.search_builds
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SearchBuildsResponse)
		err := b.DecodeMessage(msg)
		m.Response = &BatchResponse_Response_SearchBuilds{msg}
		return true, err
	case 100: // response.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(status.Status)
		err := b.DecodeMessage(msg)
		m.Response = &BatchResponse_Response_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BatchResponse_Response_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BatchResponse_Response)
	// response
	switch x := m.Response.(type) {
	case *BatchResponse_Response_GetBuild:
		s := proto.Size(x.GetBuild)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BatchResponse_Response_SearchBuilds:
		s := proto.Size(x.SearchBuilds)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BatchResponse_Response_Error:
		s := proto.Size(x.Error)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A build predicate.
//
// At least one of the following fields is required: builder, gerrit_changes and
// git_commits..
// If a field value is empty, it is ignored, unless stated otherwise.
type BuildPredicate struct {
	// A build must be in this builder.
	Builder *BuilderID `protobuf:"bytes,1,opt,name=builder" json:"builder,omitempty"`
	// A build must have this status.
	Status Status `protobuf:"varint,2,opt,name=status,enum=buildbucket.v2.Status" json:"status,omitempty"`
	// A build's Build.Input.gerrit_changes must include ALL of these changes.
	GerritChanges []*GerritChange `protobuf:"bytes,3,rep,name=gerrit_changes,json=gerritChanges" json:"gerrit_changes,omitempty"`
	// A build must be created by this identity.
	CreatedBy string `protobuf:"bytes,5,opt,name=created_by,json=createdBy" json:"created_by,omitempty"`
	// A build must have ALL of these tags.
	// For "ANY of these tags" make separate RPCs.
	Tags []*StringPair `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	// A build must have been created within the specified range.
	// Both boundaries are optional.
	CreateTime *TimeRange `protobuf:"bytes,7,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// If false (default), a build must be non-experimental.
	// Otherwise it may be experimental or non-experimental.
	IncludeExperimental  bool     `protobuf:"varint,8,opt,name=include_experimental,json=includeExperimental" json:"include_experimental,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildPredicate) Reset()         { *m = BuildPredicate{} }
func (m *BuildPredicate) String() string { return proto.CompactTextString(m) }
func (*BuildPredicate) ProtoMessage()    {}
func (*BuildPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ed8cc71e9cf76a07, []int{5}
}
func (m *BuildPredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildPredicate.Unmarshal(m, b)
}
func (m *BuildPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildPredicate.Marshal(b, m, deterministic)
}
func (dst *BuildPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildPredicate.Merge(dst, src)
}
func (m *BuildPredicate) XXX_Size() int {
	return xxx_messageInfo_BuildPredicate.Size(m)
}
func (m *BuildPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_BuildPredicate proto.InternalMessageInfo

func (m *BuildPredicate) GetBuilder() *BuilderID {
	if m != nil {
		return m.Builder
	}
	return nil
}

func (m *BuildPredicate) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (m *BuildPredicate) GetGerritChanges() []*GerritChange {
	if m != nil {
		return m.GerritChanges
	}
	return nil
}

func (m *BuildPredicate) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *BuildPredicate) GetTags() []*StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *BuildPredicate) GetCreateTime() *TimeRange {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *BuildPredicate) GetIncludeExperimental() bool {
	if m != nil {
		return m.IncludeExperimental
	}
	return false
}

func init() {
	proto.RegisterType((*GetBuildRequest)(nil), "buildbucket.v2.GetBuildRequest")
	proto.RegisterType((*SearchBuildsRequest)(nil), "buildbucket.v2.SearchBuildsRequest")
	proto.RegisterType((*SearchBuildsResponse)(nil), "buildbucket.v2.SearchBuildsResponse")
	proto.RegisterType((*BatchRequest)(nil), "buildbucket.v2.BatchRequest")
	proto.RegisterType((*BatchRequest_Request)(nil), "buildbucket.v2.BatchRequest.Request")
	proto.RegisterType((*BatchResponse)(nil), "buildbucket.v2.BatchResponse")
	proto.RegisterType((*BatchResponse_Response)(nil), "buildbucket.v2.BatchResponse.Response")
	proto.RegisterType((*BuildPredicate)(nil), "buildbucket.v2.BuildPredicate")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BuildsClient is the client API for Builds service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildsClient interface {
	// Gets a build.
	//
	// By default the returned build does not include all fields.
	// See GetBuildRequest.fields.
	//
	// Buildbot: if the specified build is a buildbot build, converts it to Build
	// message with the following rules:
	// * bucket names are full, e.g. "master.chromium". Note that LUCI buckets
	//   in v2 are shortened, e.g. "ci".
	// * if a v2 Build field does not make sense in V1, it is unset/empty.
	// * step support is not implemented for Buildbot builds.
	// Note that it does not support getting a buildbot build by build number.
	GetBuild(ctx context.Context, in *GetBuildRequest, opts ...grpc.CallOption) (*Build, error)
	// Searches for builds.
	//
	// Buildbot: can return Buildbot builds, see GetBuild for conversion rules.
	// For example, response may include a mix of LUCI and Buildbot builds if the
	// predicate is a CL.
	// Cannot search in a buildbot bucket or buildbot builder, e.g.
	// {
	//   "predicate": {
	//     "builder": {
	//       "project": "chromium",
	//       "bucket": "master.chromium",
	//       "builder": "linux-rel"
	//     }
	//   }
	// }
	// will look for builds in "master.chromium" LUCI bucket which probably does
	// not exist.
	SearchBuilds(ctx context.Context, in *SearchBuildsRequest, opts ...grpc.CallOption) (*SearchBuildsResponse, error)
	// Executes multiple requests in a batch.
	// The response code is always OK.
	Batch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponse, error)
}
type buildsPRPCClient struct {
	client *prpc.Client
}

func NewBuildsPRPCClient(client *prpc.Client) BuildsClient {
	return &buildsPRPCClient{client}
}

func (c *buildsPRPCClient) GetBuild(ctx context.Context, in *GetBuildRequest, opts ...grpc.CallOption) (*Build, error) {
	out := new(Build)
	err := c.client.Call(ctx, "buildbucket.v2.Builds", "GetBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsPRPCClient) SearchBuilds(ctx context.Context, in *SearchBuildsRequest, opts ...grpc.CallOption) (*SearchBuildsResponse, error) {
	out := new(SearchBuildsResponse)
	err := c.client.Call(ctx, "buildbucket.v2.Builds", "SearchBuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsPRPCClient) Batch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponse, error) {
	out := new(BatchResponse)
	err := c.client.Call(ctx, "buildbucket.v2.Builds", "Batch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type buildsClient struct {
	cc *grpc.ClientConn
}

func NewBuildsClient(cc *grpc.ClientConn) BuildsClient {
	return &buildsClient{cc}
}

func (c *buildsClient) GetBuild(ctx context.Context, in *GetBuildRequest, opts ...grpc.CallOption) (*Build, error) {
	out := new(Build)
	err := c.cc.Invoke(ctx, "/buildbucket.v2.Builds/GetBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsClient) SearchBuilds(ctx context.Context, in *SearchBuildsRequest, opts ...grpc.CallOption) (*SearchBuildsResponse, error) {
	out := new(SearchBuildsResponse)
	err := c.cc.Invoke(ctx, "/buildbucket.v2.Builds/SearchBuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsClient) Batch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponse, error) {
	out := new(BatchResponse)
	err := c.cc.Invoke(ctx, "/buildbucket.v2.Builds/Batch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildsServer is the server API for Builds service.
type BuildsServer interface {
	// Gets a build.
	//
	// By default the returned build does not include all fields.
	// See GetBuildRequest.fields.
	//
	// Buildbot: if the specified build is a buildbot build, converts it to Build
	// message with the following rules:
	// * bucket names are full, e.g. "master.chromium". Note that LUCI buckets
	//   in v2 are shortened, e.g. "ci".
	// * if a v2 Build field does not make sense in V1, it is unset/empty.
	// * step support is not implemented for Buildbot builds.
	// Note that it does not support getting a buildbot build by build number.
	GetBuild(context.Context, *GetBuildRequest) (*Build, error)
	// Searches for builds.
	//
	// Buildbot: can return Buildbot builds, see GetBuild for conversion rules.
	// For example, response may include a mix of LUCI and Buildbot builds if the
	// predicate is a CL.
	// Cannot search in a buildbot bucket or buildbot builder, e.g.
	// {
	//   "predicate": {
	//     "builder": {
	//       "project": "chromium",
	//       "bucket": "master.chromium",
	//       "builder": "linux-rel"
	//     }
	//   }
	// }
	// will look for builds in "master.chromium" LUCI bucket which probably does
	// not exist.
	SearchBuilds(context.Context, *SearchBuildsRequest) (*SearchBuildsResponse, error)
	// Executes multiple requests in a batch.
	// The response code is always OK.
	Batch(context.Context, *BatchRequest) (*BatchResponse, error)
}

func RegisterBuildsServer(s prpc.Registrar, srv BuildsServer) {
	s.RegisterService(&_Builds_serviceDesc, srv)
}

func _Builds_GetBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).GetBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbucket.v2.Builds/GetBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).GetBuild(ctx, req.(*GetBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builds_SearchBuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBuildsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).SearchBuilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbucket.v2.Builds/SearchBuilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).SearchBuilds(ctx, req.(*SearchBuildsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builds_Batch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).Batch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbucket.v2.Builds/Batch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).Batch(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Builds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buildbucket.v2.Builds",
	HandlerType: (*BuildsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBuild",
			Handler:    _Builds_GetBuild_Handler,
		},
		{
			MethodName: "SearchBuilds",
			Handler:    _Builds_SearchBuilds_Handler,
		},
		{
			MethodName: "Batch",
			Handler:    _Builds_Batch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/buildbucket/proto/rpc.proto",
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/buildbucket/proto/rpc.proto", fileDescriptor_rpc_ed8cc71e9cf76a07)
}

var fileDescriptor_rpc_ed8cc71e9cf76a07 = []byte{
	// 768 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x15, 0xa5, 0xea, 0x75, 0xf5, 0x30, 0x30, 0x76, 0x5b, 0x96, 0xb5, 0x5b, 0x95, 0x35, 0x0c,
	0xa1, 0x40, 0xa9, 0x96, 0x36, 0xba, 0x68, 0x8b, 0xa2, 0x90, 0x5d, 0xd7, 0x6d, 0xd1, 0xc0, 0x18,
	0x7b, 0x95, 0x2c, 0x08, 0x8a, 0x1c, 0x53, 0x03, 0x89, 0x8f, 0xcc, 0x0c, 0x03, 0xdb, 0xff, 0x12,
	0x20, 0x3f, 0x90, 0xaf, 0x08, 0xf2, 0x45, 0xde, 0x64, 0x19, 0x70, 0x38, 0xb4, 0xa9, 0x47, 0x1c,
	0x21, 0x2b, 0x8d, 0xce, 0x3d, 0xf7, 0x71, 0xee, 0xdc, 0x3b, 0x84, 0x9f, 0x82, 0xd8, 0xf2, 0xa6,
	0x2c, 0x0e, 0x69, 0x1a, 0x5a, 0x31, 0x0b, 0x46, 0xf3, 0xd4, 0xa3, 0xa3, 0x49, 0x4a, 0xe7, 0xfe,
	0x24, 0xf5, 0x66, 0x44, 0x8c, 0x12, 0x16, 0x8b, 0x78, 0xc4, 0x12, 0xcf, 0x92, 0x27, 0xd4, 0x2f,
	0x19, 0xad, 0x17, 0xb6, 0x31, 0x08, 0xe2, 0x38, 0x98, 0x93, 0x9c, 0x37, 0x49, 0xaf, 0x46, 0x57,
	0x94, 0xcc, 0x7d, 0x27, 0x74, 0xf9, 0x2c, 0xf7, 0x30, 0xbe, 0x54, 0x0c, 0x96, 0x78, 0x23, 0x2e,
	0x5c, 0x91, 0x72, 0x65, 0x38, 0xdc, 0x30, 0xb9, 0x17, 0x87, 0x61, 0x1c, 0x29, 0x27, 0x7b, 0x43,
	0x27, 0x89, 0xe4, 0x3e, 0xe6, 0x6b, 0x0d, 0xb6, 0xfe, 0x26, 0x62, 0x9c, 0x41, 0x98, 0x3c, 0x4f,
	0x09, 0x17, 0xa8, 0x0f, 0x55, 0xea, 0xeb, 0xda, 0x40, 0x1b, 0xd6, 0x70, 0x95, 0xfa, 0xe8, 0x10,
	0x9a, 0xd2, 0x85, 0x30, 0xbd, 0x3a, 0xd0, 0x86, 0x1d, 0xfb, 0x2b, 0x6b, 0x51, 0xa9, 0x35, 0xce,
	0xcd, 0xff, 0x9c, 0xe0, 0x82, 0x89, 0xbe, 0x83, 0xae, 0x3c, 0x3a, 0x51, 0x1a, 0x4e, 0x08, 0xd3,
	0x6b, 0x03, 0x6d, 0x58, 0xc7, 0x1d, 0x89, 0x3d, 0x91, 0x10, 0xb2, 0xa1, 0x21, 0x3b, 0xc2, 0x75,
	0x5f, 0x86, 0x35, 0xac, 0xbc, 0x1d, 0x56, 0xd1, 0x30, 0xeb, 0x34, 0x33, 0xff, 0xef, 0xf2, 0x19,
	0x56, 0x4c, 0xf3, 0x8d, 0x06, 0xdb, 0x17, 0xc4, 0x65, 0xde, 0x54, 0xe6, 0xe4, 0x45, 0xcd, 0xbf,
	0x43, 0x3b, 0x61, 0xc4, 0xa7, 0x9e, 0x2b, 0x88, 0x2c, 0xbd, 0x63, 0x7f, 0xb3, 0xb6, 0xca, 0xf3,
	0x82, 0x85, 0x1f, 0x1c, 0x3e, 0xa5, 0x12, 0xf4, 0x35, 0xb4, 0x13, 0x37, 0x20, 0x0e, 0xa7, 0xb7,
	0x44, 0x27, 0x52, 0x5d, 0x2b, 0x03, 0x2e, 0xe8, 0x2d, 0x41, 0x7b, 0x00, 0xd2, 0x28, 0xe2, 0x19,
	0x89, 0xf4, 0xab, 0x81, 0x36, 0x6c, 0x63, 0x49, 0xbf, 0xcc, 0x00, 0x33, 0x84, 0x9d, 0x45, 0x11,
	0x3c, 0x89, 0x23, 0x4e, 0xd0, 0x8f, 0xd0, 0x90, 0x35, 0x73, 0x5d, 0x1b, 0xd4, 0x86, 0x1d, 0xfb,
	0xf3, 0xb5, 0x12, 0xb0, 0x22, 0xa1, 0x03, 0xd8, 0x8a, 0xc8, 0xb5, 0x70, 0x4a, 0xa9, 0x7c, 0x99,
	0xaa, 0x97, 0xc1, 0xe7, 0xf7, 0xe9, 0xee, 0x34, 0xe8, 0x8e, 0x5d, 0xe1, 0x4d, 0x8b, 0x6e, 0xfd,
	0x09, 0x2d, 0x96, 0x1f, 0x8b, 0x4c, 0xfb, 0x2b, 0x99, 0x4a, 0x7c, 0x4b, 0xfd, 0xe2, 0x7b, 0x2f,
	0xe3, 0x95, 0x06, 0xcd, 0x22, 0xda, 0x1f, 0xd0, 0x0e, 0x88, 0x70, 0x64, 0x00, 0xd5, 0xfb, 0x6f,
	0x97, 0xc3, 0x2d, 0xcd, 0xd8, 0x59, 0x05, 0xb7, 0x02, 0x05, 0xa1, 0x7f, 0xa1, 0xc7, 0x65, 0x37,
	0x1c, 0x25, 0x3e, 0x9f, 0xb2, 0xef, 0x97, 0x63, 0xac, 0xb9, 0xf7, 0xb3, 0x0a, 0xee, 0xf2, 0x12,
	0x3c, 0x6e, 0x43, 0x53, 0xd5, 0x68, 0xbe, 0xac, 0x42, 0x4f, 0xa9, 0x50, 0xed, 0x3d, 0x81, 0x36,
	0x53, 0xe7, 0x42, 0xf7, 0xc1, 0x07, 0x74, 0xe7, 0x2c, 0xab, 0x38, 0xe0, 0x07, 0x47, 0xe3, 0xad,
	0x06, 0xad, 0xfb, 0x90, 0x47, 0xab, 0xda, 0xd7, 0x5f, 0xda, 0x82, 0xe2, 0xff, 0xd6, 0x2b, 0xde,
	0x7f, 0x5c, 0x71, 0x9e, 0x72, 0x59, 0x32, 0xfa, 0x01, 0xea, 0x84, 0xb1, 0x98, 0xa9, 0xd9, 0x45,
	0xc5, 0xec, 0x66, 0x0f, 0xd3, 0x85, 0x7c, 0x54, 0xce, 0x2a, 0x38, 0xa7, 0x8c, 0x21, 0xbb, 0xf8,
	0x3c, 0x8e, 0xf9, 0xae, 0x0a, 0xfd, 0xc5, 0x95, 0x28, 0x6f, 0xba, 0xb6, 0xf1, 0xa6, 0x5b, 0xd0,
	0xc8, 0xdf, 0x2e, 0xa9, 0xa2, 0x6f, 0x7f, 0xb1, 0xa2, 0x42, 0x5a, 0xb1, 0x62, 0xa1, 0x63, 0xe8,
	0x07, 0x84, 0x31, 0x2a, 0x1c, 0x6f, 0xea, 0x46, 0x01, 0xe1, 0x7a, 0x4d, 0x5e, 0xc5, 0xee, 0xea,
	0xcc, 0x64, 0xac, 0x63, 0x49, 0xc2, 0xbd, 0xa0, 0xf4, 0x8f, 0x67, 0x0b, 0xe6, 0x31, 0xe2, 0x0a,
	0xe2, 0x3b, 0x93, 0x1b, 0xbd, 0x9e, 0x2f, 0x98, 0x42, 0xc6, 0x37, 0xc8, 0x82, 0xcf, 0x84, 0x1b,
	0x70, 0xbd, 0x21, 0x23, 0x1b, 0xab, 0x15, 0x31, 0x1a, 0x05, 0xe7, 0x2e, 0x65, 0x58, 0xf2, 0xd0,
	0xaf, 0xd0, 0xc9, 0x9d, 0x1d, 0x41, 0x43, 0xa2, 0x37, 0xd7, 0x8b, 0xbf, 0xa4, 0x21, 0xc1, 0xb2,
	0x1a, 0x95, 0x3c, 0x03, 0xd0, 0xcf, 0xb0, 0x43, 0x23, 0x6f, 0x9e, 0xfa, 0xc4, 0x21, 0xd7, 0x09,
	0x61, 0x34, 0x24, 0x91, 0x70, 0xe7, 0x7a, 0x6b, 0xa0, 0x0d, 0x5b, 0x78, 0x5b, 0xd9, 0xfe, 0x2a,
	0x99, 0xec, 0x3b, 0x0d, 0x1a, 0xea, 0xf6, 0x4e, 0xa0, 0x55, 0xec, 0x06, 0xfa, 0xd8, 0xd6, 0x18,
	0xeb, 0x47, 0xcb, 0xac, 0xa0, 0x67, 0xd0, 0x2d, 0xcf, 0x0a, 0xda, 0x64, 0x77, 0x8c, 0x8d, 0xc6,
	0xcd, 0xac, 0xa0, 0x53, 0xa8, 0xcb, 0xad, 0x40, 0xbb, 0x8f, 0x3d, 0x12, 0xc6, 0xde, 0xa3, 0xab,
	0x64, 0x56, 0xc6, 0xbf, 0x3c, 0x3d, 0xda, 0xec, 0x0b, 0xf5, 0x5b, 0x09, 0x49, 0x26, 0x93, 0x86,
	0x04, 0x0f, 0xdf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x64, 0x3b, 0x09, 0x92, 0x07, 0x00, 0x00,
}
