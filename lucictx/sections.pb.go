// Copyright 2020 The LUCI Authors.
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
// source: go.chromium.org/luci/lucictx/sections.proto

package lucictx

import (
	proto "github.com/golang/protobuf/proto"
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

// LocalAuth is a struct that may be used with the "local_auth" section of
// LUCI_CONTEXT.
type LocalAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RPCPort and Secret define how to connect to the local auth server.
	RpcPort uint32 `protobuf:"varint,1,opt,name=rpc_port,proto3" json:"rpc_port,omitempty"`
	Secret  []byte `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	// Accounts and DefaultAccountID defines what access tokens are available.
	Accounts         []*LocalAuthAccount `protobuf:"bytes,3,rep,name=accounts,proto3" json:"accounts,omitempty"`
	DefaultAccountId string              `protobuf:"bytes,4,opt,name=default_account_id,proto3" json:"default_account_id,omitempty"`
}

func (x *LocalAuth) Reset() {
	*x = LocalAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalAuth) ProtoMessage() {}

func (x *LocalAuth) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalAuth.ProtoReflect.Descriptor instead.
func (*LocalAuth) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_lucictx_sections_proto_rawDescGZIP(), []int{0}
}

func (x *LocalAuth) GetRpcPort() uint32 {
	if x != nil {
		return x.RpcPort
	}
	return 0
}

func (x *LocalAuth) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *LocalAuth) GetAccounts() []*LocalAuthAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *LocalAuth) GetDefaultAccountId() string {
	if x != nil {
		return x.DefaultAccountId
	}
	return ""
}

// LocalAuthAccount contains information about a service account available
// through a local auth server.
type LocalAuthAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is logical identifier of the account, e.g. "system" or "task".
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Email is an account email or "-" if not available.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *LocalAuthAccount) Reset() {
	*x = LocalAuthAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalAuthAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalAuthAccount) ProtoMessage() {}

func (x *LocalAuthAccount) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalAuthAccount.ProtoReflect.Descriptor instead.
func (*LocalAuthAccount) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_lucictx_sections_proto_rawDescGZIP(), []int{1}
}

func (x *LocalAuthAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LocalAuthAccount) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// Swarming is a struct that may be used with the "swarming" section of
// LUCI_CONTEXT.
type Swarming struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user-supplied secret bytes specified for the task, if any. This can be
	// used to pass application or task-specific secret keys, JSON, etc. from the
	// task triggerer directly to the task. The bytes will not appear on any
	// swarming UI, or be visible to any users of the swarming service.
	SecretBytes []byte `protobuf:"bytes,1,opt,name=secret_bytes,proto3" json:"secret_bytes,omitempty"`
}

func (x *Swarming) Reset() {
	*x = Swarming{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Swarming) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Swarming) ProtoMessage() {}

func (x *Swarming) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Swarming.ProtoReflect.Descriptor instead.
func (*Swarming) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_lucictx_sections_proto_rawDescGZIP(), []int{2}
}

func (x *Swarming) GetSecretBytes() []byte {
	if x != nil {
		return x.SecretBytes
	}
	return nil
}

// LUCIExe is a struct that may be used with the "luciexe" section of
// LUCI_CONTEXT.
type LUCIExe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The absolute path of the base cache directory. This directory MAY be on the
	// same filesystem as CWD (but is not guaranteed to be). The available caches
	// are described in Buildbucket as CacheEntry messages.
	CacheDir string `protobuf:"bytes,1,opt,name=cache_dir,proto3" json:"cache_dir,omitempty"`
}

func (x *LUCIExe) Reset() {
	*x = LUCIExe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LUCIExe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LUCIExe) ProtoMessage() {}

func (x *LUCIExe) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LUCIExe.ProtoReflect.Descriptor instead.
func (*LUCIExe) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_lucictx_sections_proto_rawDescGZIP(), []int{3}
}

func (x *LUCIExe) GetCacheDir() string {
	if x != nil {
		return x.CacheDir
	}
	return ""
}

// ResultDB is a struct that may be used with the "resultdb" section of
// LUCI_CONTEXT.
type ResultDB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"` // e.g. results.api.cr.dev
	// The invocation in the current context.
	// For example, in a Buildbucket build context, it is the build's invocation.
	//
	// This is the recommended way to propagate invocation name and update token
	// to subprocesses.
	CurrentInvocation *ResultDBInvocation `protobuf:"bytes,2,opt,name=current_invocation,proto3" json:"current_invocation,omitempty"`
}

func (x *ResultDB) Reset() {
	*x = ResultDB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultDB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultDB) ProtoMessage() {}

func (x *ResultDB) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultDB.ProtoReflect.Descriptor instead.
func (*ResultDB) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_lucictx_sections_proto_rawDescGZIP(), []int{4}
}

func (x *ResultDB) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ResultDB) GetCurrentInvocation() *ResultDBInvocation {
	if x != nil {
		return x.CurrentInvocation
	}
	return nil
}

// ResultDBInvocation is a struct that contains the necessary info to update an
// invocation in the ResultDB service.
type ResultDBInvocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                 // e.g. "invocations/build:1234567890"
	UpdateToken string `protobuf:"bytes,2,opt,name=update_token,proto3" json:"update_token,omitempty"` // required in all mutation requests
}

func (x *ResultDBInvocation) Reset() {
	*x = ResultDBInvocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultDBInvocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultDBInvocation) ProtoMessage() {}

func (x *ResultDBInvocation) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultDBInvocation.ProtoReflect.Descriptor instead.
func (*ResultDBInvocation) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_lucictx_sections_proto_rawDescGZIP(), []int{5}
}

func (x *ResultDBInvocation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResultDBInvocation) GetUpdateToken() string {
	if x != nil {
		return x.UpdateToken
	}
	return ""
}

type ResultSink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TCP address (e.g. "localhost:62115") where a ResultSink pRPC server is hosted.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// secret string required in all ResultSink requests in HTTP header
	// `Authorization: ResultSink <auth-token>`
	AuthToken string `protobuf:"bytes,2,opt,name=auth_token,proto3" json:"auth_token,omitempty"`
}

func (x *ResultSink) Reset() {
	*x = ResultSink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultSink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultSink) ProtoMessage() {}

func (x *ResultSink) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultSink.ProtoReflect.Descriptor instead.
func (*ResultSink) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_lucictx_sections_proto_rawDescGZIP(), []int{6}
}

func (x *ResultSink) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ResultSink) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

var File_go_chromium_org_luci_lucictx_sections_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_lucictx_sections_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x63, 0x74, 0x78, 0x2f, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c,
	0x75, 0x63, 0x69, 0x63, 0x74, 0x78, 0x22, 0xa6, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x75, 0x63,
	0x69, 0x63, 0x74, 0x78, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12,
	0x2e, 0x0a, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22,
	0x38, 0x0a, 0x10, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2e, 0x0a, 0x08, 0x53, 0x77, 0x61,
	0x72, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x27, 0x0a, 0x07, 0x4c, 0x55, 0x43,
	0x49, 0x45, 0x78, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x64, 0x69,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x64,
	0x69, 0x72, 0x22, 0x73, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x42, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x12, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x63, 0x74, 0x78,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x42, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x76,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x44, 0x42, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53,
	0x69, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x63, 0x74, 0x78, 0x3b, 0x6c, 0x75,
	0x63, 0x69, 0x63, 0x74, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_lucictx_sections_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_lucictx_sections_proto_rawDescData = file_go_chromium_org_luci_lucictx_sections_proto_rawDesc
)

func file_go_chromium_org_luci_lucictx_sections_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_lucictx_sections_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_lucictx_sections_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_lucictx_sections_proto_rawDescData)
	})
	return file_go_chromium_org_luci_lucictx_sections_proto_rawDescData
}

var file_go_chromium_org_luci_lucictx_sections_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_go_chromium_org_luci_lucictx_sections_proto_goTypes = []interface{}{
	(*LocalAuth)(nil),          // 0: lucictx.LocalAuth
	(*LocalAuthAccount)(nil),   // 1: lucictx.LocalAuthAccount
	(*Swarming)(nil),           // 2: lucictx.Swarming
	(*LUCIExe)(nil),            // 3: lucictx.LUCIExe
	(*ResultDB)(nil),           // 4: lucictx.ResultDB
	(*ResultDBInvocation)(nil), // 5: lucictx.ResultDBInvocation
	(*ResultSink)(nil),         // 6: lucictx.ResultSink
}
var file_go_chromium_org_luci_lucictx_sections_proto_depIdxs = []int32{
	1, // 0: lucictx.LocalAuth.accounts:type_name -> lucictx.LocalAuthAccount
	5, // 1: lucictx.ResultDB.current_invocation:type_name -> lucictx.ResultDBInvocation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_lucictx_sections_proto_init() }
func file_go_chromium_org_luci_lucictx_sections_proto_init() {
	if File_go_chromium_org_luci_lucictx_sections_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalAuth); i {
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
		file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalAuthAccount); i {
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
		file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Swarming); i {
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
		file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LUCIExe); i {
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
		file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultDB); i {
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
		file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultDBInvocation); i {
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
		file_go_chromium_org_luci_lucictx_sections_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultSink); i {
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
			RawDescriptor: file_go_chromium_org_luci_lucictx_sections_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_lucictx_sections_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_lucictx_sections_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_lucictx_sections_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_lucictx_sections_proto = out.File
	file_go_chromium_org_luci_lucictx_sections_proto_rawDesc = nil
	file_go_chromium_org_luci_lucictx_sections_proto_goTypes = nil
	file_go_chromium_org_luci_lucictx_sections_proto_depIdxs = nil
}