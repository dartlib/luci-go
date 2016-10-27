// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/tokenserver/api/admin/v1/admin.proto
// DO NOT EDIT!

/*
Package admin is a generated protocol buffer package.

It is generated from these files:
	github.com/luci/luci-go/tokenserver/api/admin/v1/admin.proto
	github.com/luci/luci-go/tokenserver/api/admin/v1/certificate_authorities.proto
	github.com/luci/luci-go/tokenserver/api/admin/v1/config.proto

It has these top-level messages:
	ImportedConfigs
	InspectMachineTokenRequest
	InspectMachineTokenResponse
	FetchCRLRequest
	FetchCRLResponse
	ListCAsResponse
	GetCAStatusRequest
	GetCAStatusResponse
	IsRevokedCertRequest
	IsRevokedCertResponse
	CheckCertificateRequest
	CheckCertificateResponse
	CRLStatus
	TokenServerConfig
	CertificateAuthorityConfig
	DomainConfig
	DelegationPermissions
	DelegationRule
*/
package admin

import prpc "github.com/luci/luci-go/grpc/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"
import tokenserver "github.com/luci/luci-go/tokenserver/api"

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

// ImportedConfigs is returned by ImportCAConfigs methods on success.
type ImportedConfigs struct {
	// The list of imported config files with their revision numbers.
	ImportedConfigs []*ImportedConfigs_ConfigFile `protobuf:"bytes,1,rep,name=imported_configs,json=importedConfigs" json:"imported_configs,omitempty"`
}

func (m *ImportedConfigs) Reset()                    { *m = ImportedConfigs{} }
func (m *ImportedConfigs) String() string            { return proto.CompactTextString(m) }
func (*ImportedConfigs) ProtoMessage()               {}
func (*ImportedConfigs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ImportedConfigs) GetImportedConfigs() []*ImportedConfigs_ConfigFile {
	if m != nil {
		return m.ImportedConfigs
	}
	return nil
}

type ImportedConfigs_ConfigFile struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Revision string `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
}

func (m *ImportedConfigs_ConfigFile) Reset()                    { *m = ImportedConfigs_ConfigFile{} }
func (m *ImportedConfigs_ConfigFile) String() string            { return proto.CompactTextString(m) }
func (*ImportedConfigs_ConfigFile) ProtoMessage()               {}
func (*ImportedConfigs_ConfigFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// InspectMachineTokenRequest is body of InspectMachineToken RPC call.
//
// It contains machine token of some kind.
type InspectMachineTokenRequest struct {
	// The type of token being checked.
	//
	// Currently only LUCI_MACHINE_TOKEN is supported. This is also the default.
	TokenType tokenserver.MachineTokenType `protobuf:"varint,1,opt,name=token_type,json=tokenType,enum=tokenserver.MachineTokenType" json:"token_type,omitempty"`
	// The token body. Exact meaning depends on token_type.
	Token string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *InspectMachineTokenRequest) Reset()                    { *m = InspectMachineTokenRequest{} }
func (m *InspectMachineTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*InspectMachineTokenRequest) ProtoMessage()               {}
func (*InspectMachineTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// InspectMachineTokenResponse is return value of InspectMachineToken RPC call.
type InspectMachineTokenResponse struct {
	// True if the token is valid.
	//
	// A token is valid if its signature is correct, it hasn't expired yet and
	// the credentials it was built from (e.g. a certificate) wasn't revoked.
	Valid bool `protobuf:"varint,1,opt,name=valid" json:"valid,omitempty"`
	// Human readable summary of why token is invalid.
	//
	// Summarizes the rest of the fields of this struct. Set only if 'valid' is
	// false.
	InvalidityReason string `protobuf:"bytes,2,opt,name=invalidity_reason,json=invalidityReason" json:"invalidity_reason,omitempty"`
	// True if the token signature was verified.
	//
	// It means the token was generated by the trusted server and its body is not
	// a garbage. Note that a token can be correctly signed, but invalid (if it
	// has expired or was revoked).
	//
	// If 'signed' is false, token_type below may (or may not) be a garbage.
	// The token server uses private keys managed by Google Cloud Platform, they
	// are constantly being rotated and "old" signatures become invalid over time
	// (when corresponding keys are rotated out of existence).
	//
	// If 'signed' is false, use the rest of the response only as FYI, possibly
	// invalid or even maliciously constructed.
	Signed bool `protobuf:"varint,3,opt,name=signed" json:"signed,omitempty"`
	// True if the token signature was verified and token hasn't expired yet.
	//
	// We use "non_" prefix to make default 'false' value safer.
	NonExpired bool `protobuf:"varint,45,opt,name=non_expired,json=nonExpired" json:"non_expired,omitempty"`
	// True if the token signature was verified and the token wasn't revoked.
	//
	// It is possible for an expired token to be non revoked. They are independent
	// properties.
	//
	// We use "non_" prefix to make default 'false' value safer.
	NonRevoked bool `protobuf:"varint,5,opt,name=non_revoked,json=nonRevoked" json:"non_revoked,omitempty"`
	// Id of a private key used to sign this token, if applicable.
	SigningKeyId string `protobuf:"bytes,6,opt,name=signing_key_id,json=signingKeyId" json:"signing_key_id,omitempty"`
	// Name of a CA that issued the cert the token is based on, if applicable.
	//
	// Resolved from 'ca_id' field of the token body.
	CertCaName string `protobuf:"bytes,7,opt,name=cert_ca_name,json=certCaName" json:"cert_ca_name,omitempty"`
	// The decoded token body (depends on token_type request parameter). Empty if
	// token was malformed and couldn't be deserialized.
	//
	// Types that are valid to be assigned to TokenType:
	//	*InspectMachineTokenResponse_LuciMachineToken
	TokenType isInspectMachineTokenResponse_TokenType `protobuf_oneof:"token_type"`
}

func (m *InspectMachineTokenResponse) Reset()                    { *m = InspectMachineTokenResponse{} }
func (m *InspectMachineTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*InspectMachineTokenResponse) ProtoMessage()               {}
func (*InspectMachineTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isInspectMachineTokenResponse_TokenType interface {
	isInspectMachineTokenResponse_TokenType()
}

type InspectMachineTokenResponse_LuciMachineToken struct {
	LuciMachineToken *tokenserver.MachineTokenBody `protobuf:"bytes,20,opt,name=luci_machine_token,json=luciMachineToken,oneof"`
}

func (*InspectMachineTokenResponse_LuciMachineToken) isInspectMachineTokenResponse_TokenType() {}

func (m *InspectMachineTokenResponse) GetTokenType() isInspectMachineTokenResponse_TokenType {
	if m != nil {
		return m.TokenType
	}
	return nil
}

func (m *InspectMachineTokenResponse) GetLuciMachineToken() *tokenserver.MachineTokenBody {
	if x, ok := m.GetTokenType().(*InspectMachineTokenResponse_LuciMachineToken); ok {
		return x.LuciMachineToken
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*InspectMachineTokenResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _InspectMachineTokenResponse_OneofMarshaler, _InspectMachineTokenResponse_OneofUnmarshaler, _InspectMachineTokenResponse_OneofSizer, []interface{}{
		(*InspectMachineTokenResponse_LuciMachineToken)(nil),
	}
}

func _InspectMachineTokenResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*InspectMachineTokenResponse)
	// token_type
	switch x := m.TokenType.(type) {
	case *InspectMachineTokenResponse_LuciMachineToken:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LuciMachineToken); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("InspectMachineTokenResponse.TokenType has unexpected type %T", x)
	}
	return nil
}

func _InspectMachineTokenResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*InspectMachineTokenResponse)
	switch tag {
	case 20: // token_type.luci_machine_token
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(tokenserver.MachineTokenBody)
		err := b.DecodeMessage(msg)
		m.TokenType = &InspectMachineTokenResponse_LuciMachineToken{msg}
		return true, err
	default:
		return false, nil
	}
}

func _InspectMachineTokenResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*InspectMachineTokenResponse)
	// token_type
	switch x := m.TokenType.(type) {
	case *InspectMachineTokenResponse_LuciMachineToken:
		s := proto.Size(x.LuciMachineToken)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ImportedConfigs)(nil), "tokenserver.admin.ImportedConfigs")
	proto.RegisterType((*ImportedConfigs_ConfigFile)(nil), "tokenserver.admin.ImportedConfigs.ConfigFile")
	proto.RegisterType((*InspectMachineTokenRequest)(nil), "tokenserver.admin.InspectMachineTokenRequest")
	proto.RegisterType((*InspectMachineTokenResponse)(nil), "tokenserver.admin.InspectMachineTokenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Admin service

type AdminClient interface {
	// ImportCAConfigs makes the server read CA configs from luci-config.
	//
	// This reads 'tokenserver.cfg' file.
	//
	// Note that regularly configs are read in background each 5 min.
	// ImportCAConfigs can be used to force config reread immediately. It will
	// block until the configs are read.
	ImportCAConfigs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error)
	// ImportDelegationConfigs makes the server read 'delegation.cfg' config.
	//
	// Note that regularly configs are read in background each 5 min.
	// ImportDelegationConfigs can be used to force config reread immediately. It
	// will block until the configs are read.
	ImportDelegationConfigs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error)
	// InspectMachineToken decodes a machine token and verifies it is valid.
	//
	// It verifies the token was signed by a private key of the token server and
	// checks token's expiration time and revocation status.
	//
	// It tries to give as much information about the token and its status as
	// possible (e.g. it checks for revocation status even if token is already
	// expired).
	//
	// Administrators can use this call to debug issues with tokens.
	//
	// Returns:
	//   InspectMachineTokenResponse for tokens of supported kind.
	//   grpc.InvalidArgument error for unsupported token kind.
	//   grpc.Internal error for transient errors.
	InspectMachineToken(ctx context.Context, in *InspectMachineTokenRequest, opts ...grpc.CallOption) (*InspectMachineTokenResponse, error)
}
type adminPRPCClient struct {
	client *prpc.Client
}

func NewAdminPRPCClient(client *prpc.Client) AdminClient {
	return &adminPRPCClient{client}
}

func (c *adminPRPCClient) ImportCAConfigs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error) {
	out := new(ImportedConfigs)
	err := c.client.Call(ctx, "tokenserver.admin.Admin", "ImportCAConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminPRPCClient) ImportDelegationConfigs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error) {
	out := new(ImportedConfigs)
	err := c.client.Call(ctx, "tokenserver.admin.Admin", "ImportDelegationConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminPRPCClient) InspectMachineToken(ctx context.Context, in *InspectMachineTokenRequest, opts ...grpc.CallOption) (*InspectMachineTokenResponse, error) {
	out := new(InspectMachineTokenResponse)
	err := c.client.Call(ctx, "tokenserver.admin.Admin", "InspectMachineToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) ImportCAConfigs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error) {
	out := new(ImportedConfigs)
	err := grpc.Invoke(ctx, "/tokenserver.admin.Admin/ImportCAConfigs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ImportDelegationConfigs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ImportedConfigs, error) {
	out := new(ImportedConfigs)
	err := grpc.Invoke(ctx, "/tokenserver.admin.Admin/ImportDelegationConfigs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) InspectMachineToken(ctx context.Context, in *InspectMachineTokenRequest, opts ...grpc.CallOption) (*InspectMachineTokenResponse, error) {
	out := new(InspectMachineTokenResponse)
	err := grpc.Invoke(ctx, "/tokenserver.admin.Admin/InspectMachineToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Admin service

type AdminServer interface {
	// ImportCAConfigs makes the server read CA configs from luci-config.
	//
	// This reads 'tokenserver.cfg' file.
	//
	// Note that regularly configs are read in background each 5 min.
	// ImportCAConfigs can be used to force config reread immediately. It will
	// block until the configs are read.
	ImportCAConfigs(context.Context, *google_protobuf.Empty) (*ImportedConfigs, error)
	// ImportDelegationConfigs makes the server read 'delegation.cfg' config.
	//
	// Note that regularly configs are read in background each 5 min.
	// ImportDelegationConfigs can be used to force config reread immediately. It
	// will block until the configs are read.
	ImportDelegationConfigs(context.Context, *google_protobuf.Empty) (*ImportedConfigs, error)
	// InspectMachineToken decodes a machine token and verifies it is valid.
	//
	// It verifies the token was signed by a private key of the token server and
	// checks token's expiration time and revocation status.
	//
	// It tries to give as much information about the token and its status as
	// possible (e.g. it checks for revocation status even if token is already
	// expired).
	//
	// Administrators can use this call to debug issues with tokens.
	//
	// Returns:
	//   InspectMachineTokenResponse for tokens of supported kind.
	//   grpc.InvalidArgument error for unsupported token kind.
	//   grpc.Internal error for transient errors.
	InspectMachineToken(context.Context, *InspectMachineTokenRequest) (*InspectMachineTokenResponse, error)
}

func RegisterAdminServer(s prpc.Registrar, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_ImportCAConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ImportCAConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.Admin/ImportCAConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ImportCAConfigs(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_ImportDelegationConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ImportDelegationConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.Admin/ImportDelegationConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ImportDelegationConfigs(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_InspectMachineToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InspectMachineTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).InspectMachineToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.Admin/InspectMachineToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).InspectMachineToken(ctx, req.(*InspectMachineTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tokenserver.admin.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportCAConfigs",
			Handler:    _Admin_ImportCAConfigs_Handler,
		},
		{
			MethodName: "ImportDelegationConfigs",
			Handler:    _Admin_ImportDelegationConfigs_Handler,
		},
		{
			MethodName: "InspectMachineToken",
			Handler:    _Admin_InspectMachineToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/tokenserver/api/admin/v1/admin.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x25, 0x29, 0x09, 0xed, 0x24, 0x6a, 0xd3, 0xa5, 0x2a, 0x96, 0x2b, 0x44, 0x14, 0x71, 0x88,
	0x84, 0xb2, 0x16, 0xe1, 0x48, 0x2e, 0x25, 0x14, 0x11, 0xa1, 0x70, 0xb0, 0x8a, 0xc4, 0x6d, 0xe5,
	0xd8, 0x53, 0x77, 0x15, 0x7b, 0x77, 0xb1, 0x37, 0x16, 0xfe, 0x1e, 0x4e, 0x7c, 0x15, 0xbf, 0x82,
	0xbc, 0x1b, 0xa7, 0x09, 0x04, 0xe8, 0x81, 0x8b, 0x3d, 0xf3, 0xe6, 0xcd, 0xcc, 0x5b, 0xef, 0x33,
	0x4c, 0x62, 0xae, 0x6f, 0x57, 0x0b, 0x1a, 0xca, 0xd4, 0x4b, 0x56, 0x21, 0x37, 0x8f, 0x51, 0x2c,
	0x3d, 0x2d, 0x97, 0x28, 0x72, 0xcc, 0x0a, 0xcc, 0xbc, 0x40, 0x71, 0x2f, 0x88, 0x52, 0x2e, 0xbc,
	0xe2, 0xa5, 0x0d, 0xa8, 0xca, 0xa4, 0x96, 0xe4, 0x74, 0x8b, 0x45, 0x4d, 0xc1, 0xbd, 0x88, 0xa5,
	0x8c, 0x13, 0xf4, 0x0c, 0x61, 0xb1, 0xba, 0xf1, 0x30, 0x55, 0xba, 0xb4, 0x7c, 0xf7, 0xf5, 0x7d,
	0xb7, 0xa5, 0x41, 0x78, 0xcb, 0x05, 0x32, 0x83, 0xdb, 0xe6, 0xc1, 0xf7, 0x06, 0x9c, 0xcc, 0x52,
	0x25, 0x33, 0x8d, 0xd1, 0x54, 0x8a, 0x1b, 0x1e, 0xe7, 0xe4, 0x33, 0xf4, 0xf8, 0x1a, 0x62, 0xa1,
	0xc5, 0x9c, 0x46, 0xff, 0x60, 0xd8, 0x19, 0x8f, 0xe8, 0x6f, 0xda, 0xe8, 0x2f, 0xdd, 0xd4, 0xbe,
	0xdf, 0xf1, 0x04, 0xfd, 0x13, 0xbe, 0x5b, 0x73, 0x27, 0x00, 0x77, 0x65, 0x42, 0xe0, 0xa1, 0x08,
	0x52, 0x74, 0x1a, 0xfd, 0xc6, 0xf0, 0xc8, 0x37, 0x31, 0x71, 0xe1, 0x30, 0xc3, 0x82, 0xe7, 0x5c,
	0x0a, 0xa7, 0x69, 0xf0, 0x4d, 0x3e, 0x50, 0xe0, 0xce, 0x44, 0xae, 0x30, 0xd4, 0x73, 0x7b, 0x92,
	0xeb, 0x4a, 0x8c, 0x8f, 0x5f, 0x56, 0x98, 0x6b, 0x32, 0x01, 0x30, 0xe2, 0x98, 0x2e, 0x95, 0x9d,
	0x79, 0x3c, 0x7e, 0xba, 0xa3, 0x77, 0xbb, 0xeb, 0xba, 0x54, 0xe8, 0x1f, 0xe9, 0x3a, 0x24, 0x67,
	0xd0, 0x32, 0xc9, 0x7a, 0xa9, 0x4d, 0x06, 0x3f, 0x9a, 0x70, 0xb1, 0x77, 0x65, 0xae, 0xa4, 0xc8,
	0x4d, 0x57, 0x11, 0x24, 0x3c, 0x32, 0xeb, 0x0e, 0x7d, 0x9b, 0x90, 0x17, 0x70, 0xca, 0x85, 0x09,
	0xb9, 0x2e, 0x59, 0x86, 0x41, 0xbe, 0x39, 0x4c, 0xef, 0xae, 0xe0, 0x1b, 0x9c, 0x9c, 0x43, 0x3b,
	0xe7, 0xb1, 0xc0, 0xc8, 0x39, 0x30, 0x33, 0xd6, 0x19, 0x79, 0x06, 0x1d, 0x21, 0x05, 0xc3, 0xaf,
	0x8a, 0x67, 0x18, 0x39, 0x23, 0x53, 0x04, 0x21, 0xc5, 0x95, 0x45, 0x6a, 0x42, 0x86, 0x85, 0x5c,
	0x62, 0xe4, 0xb4, 0x36, 0x04, 0xdf, 0x22, 0xe4, 0x39, 0x1c, 0x57, 0xb3, 0xb8, 0x88, 0xd9, 0x12,
	0x4b, 0xc6, 0x23, 0xa7, 0x6d, 0x34, 0x74, 0xd7, 0xe8, 0x07, 0x2c, 0x67, 0x11, 0xe9, 0x43, 0x37,
	0xc4, 0x4c, 0xb3, 0x30, 0x60, 0xe6, 0x32, 0x1e, 0x19, 0x0e, 0x54, 0xd8, 0x34, 0xf8, 0x58, 0x5d,
	0xc9, 0x1c, 0x48, 0xe5, 0x28, 0xb6, 0x63, 0x1f, 0xe7, 0xac, 0xdf, 0x18, 0x76, 0xfe, 0xf2, 0x81,
	0xdf, 0xc8, 0xa8, 0x7c, 0xff, 0xc0, 0xef, 0x55, 0xad, 0x3b, 0x78, 0x77, 0xfb, 0x9e, 0xc6, 0xdf,
	0x9a, 0xd0, 0xba, 0xac, 0x7c, 0x44, 0xe6, 0xb5, 0x11, 0xa7, 0x97, 0xb5, 0x11, 0xcf, 0xa9, 0xf5,
	0x3d, 0xad, 0x7d, 0x4f, 0xaf, 0x2a, 0xdf, 0xbb, 0x83, 0x7f, 0xdb, 0x90, 0x7c, 0x82, 0x27, 0x16,
	0x7a, 0x8b, 0x09, 0xc6, 0x81, 0xe6, 0x52, 0xfc, 0x8f, 0xb1, 0x1a, 0x1e, 0xef, 0x31, 0x04, 0xd9,
	0xfb, 0x63, 0xfc, 0xd1, 0xab, 0x2e, 0xbd, 0x2f, 0xdd, 0xfa, 0x6c, 0xd1, 0x36, 0x4a, 0x5f, 0xfd,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x23, 0xa2, 0xeb, 0x59, 0x04, 0x00, 0x00,
}
