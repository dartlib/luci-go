// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/proto/rpc/v1/invocation.proto

package rpcpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_type "go.chromium.org/luci/resultdb/proto/type"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Invocation_State int32

const (
	// The default value. This value is used if the state is omitted.
	Invocation_STATE_UNSPECIFIED Invocation_State = 0
	// The invocation was created and accepts new results.
	Invocation_ACTIVE Invocation_State = 1
	// The invocation is immutable and no longer accepts new results nor
	// inclusions directly or indirectly.
	Invocation_FINALIZED Invocation_State = 3
)

var Invocation_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "ACTIVE",
	3: "FINALIZED",
}

var Invocation_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"ACTIVE":            1,
	"FINALIZED":         3,
}

func (x Invocation_State) String() string {
	return proto.EnumName(Invocation_State_name, int32(x))
}

func (Invocation_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4005c8951497aaef, []int{0, 0}
}

// A conceptual container of results. Immutable once finalized.
// It represents all results of some computation; examples: swarming task,
// buildbucket build, CQ attempt.
// Composable: can include other invocations, see inclusion.proto.
type Invocation struct {
	// Can be used to refer to this invocation, e.g. in ResultDB.GetInvocation
	// RPC.
	// Format: invocations/{INVOCATION_ID}
	// See also https://aip.dev/122.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Current state of the invocation.
	State Invocation_State `protobuf:"varint,2,opt,name=state,proto3,enum=luci.resultdb.rpc.v1.Invocation_State" json:"state,omitempty"`
	// True if the invocation is inactive and does NOT contain all the results
	// that the associated computation was expected to compute.
	//  * The computation was interrupted prematurely.
	//  * Such invocation should be discarded.
	//  * Often the associated computation is retried.
	//
	// False could mean 2 things:
	// * the invocation is still ACTIVE;
	// * the invocation is inactive and contains all the results that the
	//   associated computation was expected to compute.
	//
	// Use this field with state above.
	Interrupted bool `protobuf:"varint,3,opt,name=interrupted,proto3" json:"interrupted,omitempty"`
	// When the invocation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Invocation-level string key-value pairs.
	// A key can be repeated.
	Tags []*_type.StringPair `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// When the invocation was finalized, i.e. transitioned to FINALIZED state.
	// If this field is set, implies that the invocation is finalized.
	FinalizeTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=finalize_time,json=finalizeTime,proto3" json:"finalize_time,omitempty"`
	// Timestamp when the invocation will be forcefully finalized.
	// Can be extended with UpdateInvocation until finalized.
	Deadline *timestamp.Timestamp `protobuf:"bytes,7,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// Names of invocations included into this one. Overall results of this
	// invocation is a UNION of results directly included into this invocation
	// and results from the included invocations, recursively.
	// For example, a Buildbucket build invocation may include invocations of its
	// child swarming tasks and represent overall result of the build,
	// encapsulating the internal structure of the build.
	//
	// The graph is directed.
	// There can be at most one edge between a given pair of invocations.
	// The shape of the graph does not matter. What matters is only the set of
	// reachable invocations. Thus cycles are allowed and are noop.
	//
	// QueryTestResults returns test results from the transitive closure of
	// invocations.
	//
	// Use Recorder.Include RPC to modify this field.
	IncludedInvocations  []string `protobuf:"bytes,8,rep,name=included_invocations,json=includedInvocations,proto3" json:"included_invocations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Invocation) Reset()         { *m = Invocation{} }
func (m *Invocation) String() string { return proto.CompactTextString(m) }
func (*Invocation) ProtoMessage()    {}
func (*Invocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4005c8951497aaef, []int{0}
}

func (m *Invocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invocation.Unmarshal(m, b)
}
func (m *Invocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invocation.Marshal(b, m, deterministic)
}
func (m *Invocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invocation.Merge(m, src)
}
func (m *Invocation) XXX_Size() int {
	return xxx_messageInfo_Invocation.Size(m)
}
func (m *Invocation) XXX_DiscardUnknown() {
	xxx_messageInfo_Invocation.DiscardUnknown(m)
}

var xxx_messageInfo_Invocation proto.InternalMessageInfo

func (m *Invocation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Invocation) GetState() Invocation_State {
	if m != nil {
		return m.State
	}
	return Invocation_STATE_UNSPECIFIED
}

func (m *Invocation) GetInterrupted() bool {
	if m != nil {
		return m.Interrupted
	}
	return false
}

func (m *Invocation) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Invocation) GetTags() []*_type.StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Invocation) GetFinalizeTime() *timestamp.Timestamp {
	if m != nil {
		return m.FinalizeTime
	}
	return nil
}

func (m *Invocation) GetDeadline() *timestamp.Timestamp {
	if m != nil {
		return m.Deadline
	}
	return nil
}

func (m *Invocation) GetIncludedInvocations() []string {
	if m != nil {
		return m.IncludedInvocations
	}
	return nil
}

// BigQueryExport indicates that results in this invocation should be exported
// to BigQuery after finalization.
type BigQueryExport struct {
	// Name of the BigQuery project.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Name of the BigQuery Dataset.
	Dataset string `protobuf:"bytes,2,opt,name=dataset,proto3" json:"dataset,omitempty"`
	// Name of the BigQuery Table.
	Table                string                      `protobuf:"bytes,3,opt,name=table,proto3" json:"table,omitempty"`
	TestResults          *BigQueryExport_TestResults `protobuf:"bytes,4,opt,name=test_results,json=testResults,proto3" json:"test_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *BigQueryExport) Reset()         { *m = BigQueryExport{} }
func (m *BigQueryExport) String() string { return proto.CompactTextString(m) }
func (*BigQueryExport) ProtoMessage()    {}
func (*BigQueryExport) Descriptor() ([]byte, []int) {
	return fileDescriptor_4005c8951497aaef, []int{1}
}

func (m *BigQueryExport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigQueryExport.Unmarshal(m, b)
}
func (m *BigQueryExport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigQueryExport.Marshal(b, m, deterministic)
}
func (m *BigQueryExport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigQueryExport.Merge(m, src)
}
func (m *BigQueryExport) XXX_Size() int {
	return xxx_messageInfo_BigQueryExport.Size(m)
}
func (m *BigQueryExport) XXX_DiscardUnknown() {
	xxx_messageInfo_BigQueryExport.DiscardUnknown(m)
}

var xxx_messageInfo_BigQueryExport proto.InternalMessageInfo

func (m *BigQueryExport) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *BigQueryExport) GetDataset() string {
	if m != nil {
		return m.Dataset
	}
	return ""
}

func (m *BigQueryExport) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *BigQueryExport) GetTestResults() *BigQueryExport_TestResults {
	if m != nil {
		return m.TestResults
	}
	return nil
}

// TestResultExport indicates that test results should be exported.
type BigQueryExport_TestResults struct {
	// Use predicate to query test results that should be exported to
	// BigQuery table.
	Predicate            *TestResultPredicate `protobuf:"bytes,1,opt,name=predicate,proto3" json:"predicate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BigQueryExport_TestResults) Reset()         { *m = BigQueryExport_TestResults{} }
func (m *BigQueryExport_TestResults) String() string { return proto.CompactTextString(m) }
func (*BigQueryExport_TestResults) ProtoMessage()    {}
func (*BigQueryExport_TestResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_4005c8951497aaef, []int{1, 0}
}

func (m *BigQueryExport_TestResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigQueryExport_TestResults.Unmarshal(m, b)
}
func (m *BigQueryExport_TestResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigQueryExport_TestResults.Marshal(b, m, deterministic)
}
func (m *BigQueryExport_TestResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigQueryExport_TestResults.Merge(m, src)
}
func (m *BigQueryExport_TestResults) XXX_Size() int {
	return xxx_messageInfo_BigQueryExport_TestResults.Size(m)
}
func (m *BigQueryExport_TestResults) XXX_DiscardUnknown() {
	xxx_messageInfo_BigQueryExport_TestResults.DiscardUnknown(m)
}

var xxx_messageInfo_BigQueryExport_TestResults proto.InternalMessageInfo

func (m *BigQueryExport_TestResults) GetPredicate() *TestResultPredicate {
	if m != nil {
		return m.Predicate
	}
	return nil
}

func init() {
	proto.RegisterEnum("luci.resultdb.rpc.v1.Invocation_State", Invocation_State_name, Invocation_State_value)
	proto.RegisterType((*Invocation)(nil), "luci.resultdb.rpc.v1.Invocation")
	proto.RegisterType((*BigQueryExport)(nil), "luci.resultdb.rpc.v1.BigQueryExport")
	proto.RegisterType((*BigQueryExport_TestResults)(nil), "luci.resultdb.rpc.v1.BigQueryExport.TestResults")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/proto/rpc/v1/invocation.proto", fileDescriptor_4005c8951497aaef)
}

var fileDescriptor_4005c8951497aaef = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x25, 0xcb, 0xda, 0xad, 0xce, 0x36, 0x0d, 0x33, 0xa4, 0x50, 0x09, 0x88, 0x26, 0x81, 0xc2,
	0x8b, 0xb3, 0x05, 0x31, 0x09, 0xf6, 0x94, 0xae, 0x19, 0x8a, 0x84, 0xa6, 0x92, 0x96, 0x3d, 0xec,
	0xa5, 0x72, 0x12, 0x37, 0x33, 0x4a, 0x62, 0xcb, 0x71, 0x2a, 0xc6, 0x0f, 0xe1, 0xf7, 0xf5, 0x87,
	0xf0, 0x80, 0xf2, 0xd5, 0x0c, 0x54, 0x69, 0xe3, 0xd1, 0xe7, 0x9e, 0xe3, 0x7b, 0x7c, 0xee, 0x35,
	0x38, 0x8f, 0x19, 0x0a, 0x6f, 0x05, 0x4b, 0x69, 0x91, 0x22, 0x26, 0x62, 0x2b, 0x29, 0x42, 0x6a,
	0x09, 0x92, 0x17, 0x89, 0x8c, 0x02, 0x8b, 0x0b, 0x26, 0x99, 0x25, 0x78, 0x68, 0x2d, 0x4f, 0x2d,
	0x9a, 0x2d, 0x59, 0x88, 0x25, 0x65, 0x19, 0xaa, 0x70, 0x78, 0x54, 0x92, 0x51, 0x4b, 0x46, 0x82,
	0x87, 0x68, 0x79, 0x3a, 0x7c, 0x1d, 0x33, 0x16, 0x27, 0xc4, 0xc2, 0x9c, 0x5a, 0x0b, 0x4a, 0x92,
	0x68, 0x1e, 0x90, 0x5b, 0xbc, 0xa4, 0x4c, 0xd4, 0xb2, 0x35, 0xa1, 0x3a, 0x05, 0xc5, 0xc2, 0x92,
	0x34, 0x25, 0xb9, 0xc4, 0x29, 0x6f, 0x08, 0x9f, 0xfe, 0xc3, 0x14, 0x17, 0x24, 0xa2, 0x21, 0x96,
	0xa4, 0xd1, 0x7e, 0x78, 0x8c, 0x56, 0xde, 0x71, 0x62, 0x85, 0x2c, 0x4d, 0xdb, 0xa7, 0x1c, 0xff,
	0x56, 0x01, 0xf0, 0xd6, 0xef, 0x83, 0x43, 0xb0, 0x9d, 0xe1, 0x94, 0xe8, 0x8a, 0xa1, 0x98, 0x83,
	0x51, 0x7f, 0xe5, 0xa8, 0x2b, 0xa7, 0xe7, 0x57, 0x18, 0x74, 0x40, 0x2f, 0x97, 0x58, 0x12, 0x7d,
	0xcb, 0x50, 0xcc, 0x03, 0xfb, 0x2d, 0xda, 0x94, 0x02, 0xea, 0x2e, 0x43, 0xd3, 0x92, 0x3d, 0x52,
	0x57, 0x8e, 0xea, 0xd7, 0x4a, 0xf8, 0x06, 0x68, 0x34, 0x93, 0x44, 0x88, 0x82, 0x4b, 0x12, 0xe9,
	0xaa, 0xa1, 0x98, 0xbb, 0x35, 0xe1, 0x3e, 0x0e, 0x2f, 0x80, 0x16, 0x0a, 0x82, 0x25, 0x99, 0x97,
	0x09, 0xe9, 0xdb, 0x86, 0x62, 0x6a, 0xf6, 0x10, 0xd5, 0xf1, 0xa1, 0x36, 0x3e, 0x34, 0x6b, 0xe3,
	0x5b, 0x1b, 0x05, 0xb5, 0xac, 0x2c, 0x40, 0x1b, 0x6c, 0x4b, 0x1c, 0xe7, 0x7a, 0xcf, 0x50, 0x4d,
	0xcd, 0x7e, 0xf5, 0x8f, 0xdb, 0x32, 0x09, 0x34, 0x95, 0x82, 0x66, 0xf1, 0x04, 0x53, 0xe1, 0x57,
	0x5c, 0x38, 0x06, 0xfb, 0x0b, 0x9a, 0xe1, 0x84, 0xfe, 0x6c, 0x5a, 0xf7, 0x1f, 0x6c, 0x5d, 0xb9,
	0xdf, 0x6b, 0x55, 0x55, 0xe7, 0x33, 0xb0, 0x1b, 0x11, 0x1c, 0x25, 0x34, 0x23, 0xfa, 0xce, 0x43,
	0x17, 0xf8, 0x6b, 0x2e, 0x3c, 0x03, 0x47, 0x34, 0x0b, 0x93, 0x22, 0x22, 0xd1, 0xbc, 0xdb, 0xb9,
	0x5c, 0xdf, 0x35, 0x54, 0x73, 0x50, 0x37, 0x7a, 0xd6, 0x12, 0xba, 0x98, 0xf3, 0xe3, 0x8f, 0xa0,
	0x57, 0x45, 0x0d, 0x9f, 0x83, 0xa7, 0xd3, 0x99, 0x33, 0x73, 0xe7, 0xdf, 0xae, 0xa6, 0x13, 0xf7,
	0xc2, 0xbb, 0xf4, 0xdc, 0xf1, 0xe1, 0x13, 0x08, 0x40, 0xdf, 0xb9, 0x98, 0x79, 0xd7, 0xee, 0xa1,
	0x02, 0xf7, 0xc1, 0xe0, 0xd2, 0xbb, 0x72, 0xbe, 0x78, 0x37, 0xee, 0xf8, 0x50, 0x3d, 0xfe, 0xb5,
	0x05, 0x0e, 0x46, 0x34, 0xfe, 0x5a, 0x10, 0x71, 0xe7, 0xfe, 0xe0, 0x4c, 0x48, 0xf8, 0x12, 0xec,
	0x70, 0xc1, 0xbe, 0x93, 0x50, 0x36, 0x5b, 0xa0, 0xae, 0x9c, 0x2d, 0xbf, 0xc5, 0xca, 0x72, 0x84,
	0x25, 0xce, 0x89, 0xac, 0xf6, 0xa0, 0x2d, 0x37, 0x18, 0x7c, 0x01, 0x7a, 0x12, 0x07, 0x09, 0xa9,
	0x66, 0xdb, 0x14, 0x6b, 0x04, 0x4e, 0xc1, 0x9e, 0x24, 0xb9, 0x9c, 0xd7, 0x33, 0xc8, 0x9b, 0xb1,
	0x9e, 0x6c, 0x5e, 0xa3, 0xbf, 0x4d, 0xa1, 0x19, 0xc9, 0xa5, 0x5f, 0xeb, 0x7c, 0x4d, 0x76, 0x87,
	0xe1, 0x35, 0xd0, 0xee, 0xd5, 0xe0, 0x67, 0x30, 0x58, 0x7f, 0x8c, 0xca, 0xbe, 0x66, 0xbf, 0xdb,
	0xdc, 0xa0, 0x53, 0x4d, 0x5a, 0x81, 0xdf, 0x69, 0x47, 0xf6, 0xcd, 0xc9, 0xe3, 0x3f, 0xe3, 0xb9,
	0xe0, 0x21, 0x0f, 0x82, 0x7e, 0x85, 0xbd, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x69, 0x8c, 0xff,
	0x66, 0x5c, 0x04, 0x00, 0x00,
}
