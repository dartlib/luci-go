// Copyright 2018 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/cipd/api/cipd/v1/events.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EventKind int32

const (
	EventKind_EVENT_KIND_UNSPECIFIED EventKind = 0
	// Prefix events: relate to some CIPD prefix.
	EventKind_PREFIX_ACL_CHANGED EventKind = 100
	// Package events: relate to a package (as a whole).
	EventKind_PACKAGE_CREATED  EventKind = 200
	EventKind_PACKAGE_DELETED  EventKind = 201
	EventKind_PACKAGE_HIDDEN   EventKind = 202
	EventKind_PACKAGE_UNHIDDEN EventKind = 203
	// Instance events: relate to a particular package instance.
	EventKind_INSTANCE_CREATED      EventKind = 300
	EventKind_INSTANCE_DELETED      EventKind = 301
	EventKind_INSTANCE_REF_SET      EventKind = 302
	EventKind_INSTANCE_REF_UNSET    EventKind = 303
	EventKind_INSTANCE_TAG_ATTACHED EventKind = 304
	EventKind_INSTANCE_TAG_DETACHED EventKind = 305
)

// Enum value maps for EventKind.
var (
	EventKind_name = map[int32]string{
		0:   "EVENT_KIND_UNSPECIFIED",
		100: "PREFIX_ACL_CHANGED",
		200: "PACKAGE_CREATED",
		201: "PACKAGE_DELETED",
		202: "PACKAGE_HIDDEN",
		203: "PACKAGE_UNHIDDEN",
		300: "INSTANCE_CREATED",
		301: "INSTANCE_DELETED",
		302: "INSTANCE_REF_SET",
		303: "INSTANCE_REF_UNSET",
		304: "INSTANCE_TAG_ATTACHED",
		305: "INSTANCE_TAG_DETACHED",
	}
	EventKind_value = map[string]int32{
		"EVENT_KIND_UNSPECIFIED": 0,
		"PREFIX_ACL_CHANGED":     100,
		"PACKAGE_CREATED":        200,
		"PACKAGE_DELETED":        201,
		"PACKAGE_HIDDEN":         202,
		"PACKAGE_UNHIDDEN":       203,
		"INSTANCE_CREATED":       300,
		"INSTANCE_DELETED":       301,
		"INSTANCE_REF_SET":       302,
		"INSTANCE_REF_UNSET":     303,
		"INSTANCE_TAG_ATTACHED":  304,
		"INSTANCE_TAG_DETACHED":  305,
	}
)

func (x EventKind) Enum() *EventKind {
	p := new(EventKind)
	*p = x
	return p
}

func (x EventKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventKind) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_enumTypes[0].Descriptor()
}

func (EventKind) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_enumTypes[0]
}

func (x EventKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventKind.Descriptor instead.
func (EventKind) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_rawDescGZIP(), []int{0}
}

// Event in a global structured event log.
//
// It exists in both BigQuery (for adhoc queries) and in Datastore (for showing
// in web UI, e.g. for "recent tags" feature).
//
// Datastore entities contains serialized Event as is, plus a copy of some of
// its fields for indexing.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind EventKind `protobuf:"varint,1,opt,name=kind,proto3,enum=cipd.EventKind" json:"kind,omitempty"`
	Who  string    `protobuf:"bytes,2,opt,name=who,proto3" json:"who,omitempty"` // an identity string, e.g. "user:<email>"
	// Real time is used only for up to millisecond precisions. Nanoseconds are
	// abused to order events emitted by a single transaction.
	When     *timestamp.Timestamp `protobuf:"bytes,3,opt,name=when,proto3" json:"when,omitempty"`
	Package  string               `protobuf:"bytes,4,opt,name=package,proto3" json:"package,omitempty"`   // a package name or a prefix (for PREFIX_* events)
	Instance string               `protobuf:"bytes,5,opt,name=instance,proto3" json:"instance,omitempty"` // an instance ID for INSTANCE_*
	Ref      string               `protobuf:"bytes,6,opt,name=ref,proto3" json:"ref,omitempty"`           // a ref name for INSTANCE_REF_*
	Tag      string               `protobuf:"bytes,7,opt,name=tag,proto3" json:"tag,omitempty"`           // a tag (in 'k:v' form) for INSTANCE_TAG_*
	// An ACL diff for PREFIX_ACL_CHANGED.
	GrantedRole []*PrefixMetadata_ACL `protobuf:"bytes,8,rep,name=granted_role,json=grantedRole,proto3" json:"granted_role,omitempty"`
	RevokedRole []*PrefixMetadata_ACL `protobuf:"bytes,9,rep,name=revoked_role,json=revokedRole,proto3" json:"revoked_role,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetKind() EventKind {
	if x != nil {
		return x.Kind
	}
	return EventKind_EVENT_KIND_UNSPECIFIED
}

func (x *Event) GetWho() string {
	if x != nil {
		return x.Who
	}
	return ""
}

func (x *Event) GetWhen() *timestamp.Timestamp {
	if x != nil {
		return x.When
	}
	return nil
}

func (x *Event) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *Event) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *Event) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *Event) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Event) GetGrantedRole() []*PrefixMetadata_ACL {
	if x != nil {
		return x.GrantedRole
	}
	return nil
}

func (x *Event) GetRevokedRole() []*PrefixMetadata_ACL {
	if x != nil {
		return x.RevokedRole
	}
	return nil
}

var File_go_chromium_org_luci_cipd_api_cipd_v1_events_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x69, 0x70, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x69, 0x70, 0x64, 0x1a, 0x30, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x02,
	0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x77, 0x68, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x68, 0x6f, 0x12, 0x2e,
	0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x3b, 0x0a, 0x0c, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x43, 0x4c, 0x52, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x64, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64,
	0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x69,
	0x70, 0x64, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x41, 0x43, 0x4c, 0x52, 0x0b, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x52, 0x6f,
	0x6c, 0x65, 0x2a, 0xad, 0x02, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x1a, 0x0a, 0x16, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12,
	0x50, 0x52, 0x45, 0x46, 0x49, 0x58, 0x5f, 0x41, 0x43, 0x4c, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x44, 0x10, 0x64, 0x12, 0x14, 0x0a, 0x0f, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x5f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0xc8, 0x01, 0x12, 0x14, 0x0a, 0x0f, 0x50, 0x41,
	0x43, 0x4b, 0x41, 0x47, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0xc9, 0x01,
	0x12, 0x13, 0x0a, 0x0e, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x5f, 0x48, 0x49, 0x44, 0x44,
	0x45, 0x4e, 0x10, 0xca, 0x01, 0x12, 0x15, 0x0a, 0x10, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47, 0x45,
	0x5f, 0x55, 0x4e, 0x48, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0xcb, 0x01, 0x12, 0x15, 0x0a, 0x10,
	0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44,
	0x10, 0xac, 0x02, 0x12, 0x15, 0x0a, 0x10, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0xad, 0x02, 0x12, 0x15, 0x0a, 0x10, 0x49, 0x4e,
	0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x46, 0x5f, 0x53, 0x45, 0x54, 0x10, 0xae,
	0x02, 0x12, 0x17, 0x0a, 0x12, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x52, 0x45,
	0x46, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0xaf, 0x02, 0x12, 0x1a, 0x0a, 0x15, 0x49, 0x4e,
	0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43,
	0x48, 0x45, 0x44, 0x10, 0xb0, 0x02, 0x12, 0x1a, 0x0a, 0x15, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e,
	0x43, 0x45, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x43, 0x48, 0x45, 0x44, 0x10,
	0xb1, 0x02, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_rawDescData = file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_rawDesc
)

func file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_rawDescData
}

var file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_goTypes = []interface{}{
	(EventKind)(0),              // 0: cipd.EventKind
	(*Event)(nil),               // 1: cipd.Event
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*PrefixMetadata_ACL)(nil),  // 3: cipd.PrefixMetadata.ACL
}
var file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_depIdxs = []int32{
	0, // 0: cipd.Event.kind:type_name -> cipd.EventKind
	2, // 1: cipd.Event.when:type_name -> google.protobuf.Timestamp
	3, // 2: cipd.Event.granted_role:type_name -> cipd.PrefixMetadata.ACL
	3, // 3: cipd.Event.revoked_role:type_name -> cipd.PrefixMetadata.ACL
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_init() }
func file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_init() {
	if File_go_chromium_org_luci_cipd_api_cipd_v1_events_proto != nil {
		return
	}
	file_go_chromium_org_luci_cipd_api_cipd_v1_repo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cipd_api_cipd_v1_events_proto = out.File
	file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_rawDesc = nil
	file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_goTypes = nil
	file_go_chromium_org_luci_cipd_api_cipd_v1_events_proto_depIdxs = nil
}
