// Code generated by protoc-gen-go.
// source: admin.proto
// DO NOT EDIT!

/*
Package tokenserver is a generated protocol buffer package.

It is generated from these files:
	admin.proto
	config.proto

It has these top-level messages:
	ImportConfigResponse
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
*/
package tokenserver

import prpccommon "github.com/luci/luci-go/common/prpc"
import prpc "github.com/luci/luci-go/server/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"
import google_protobuf1 "github.com/luci/luci-go/common/proto/google"

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
const _ = proto.ProtoPackageIsVersion1

// ImportConfigResponse is returned by ImportConfig on success.
type ImportConfigResponse struct {
	Revision string `protobuf:"bytes,1,opt,name=revision" json:"revision,omitempty"`
}

func (m *ImportConfigResponse) Reset()                    { *m = ImportConfigResponse{} }
func (m *ImportConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ImportConfigResponse) ProtoMessage()               {}
func (*ImportConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// FetchCRLRequest identifies a name of CA to fetch CRL for.
type FetchCRLRequest struct {
	Cn    string `protobuf:"bytes,1,opt,name=cn" json:"cn,omitempty"`
	Force bool   `protobuf:"varint,2,opt,name=force" json:"force,omitempty"`
}

func (m *FetchCRLRequest) Reset()                    { *m = FetchCRLRequest{} }
func (m *FetchCRLRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchCRLRequest) ProtoMessage()               {}
func (*FetchCRLRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// FetchCRLResponse is returned by FetchCRL.
type FetchCRLResponse struct {
	CrlStatus *CRLStatus `protobuf:"bytes,1,opt,name=crl_status" json:"crl_status,omitempty"`
}

func (m *FetchCRLResponse) Reset()                    { *m = FetchCRLResponse{} }
func (m *FetchCRLResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchCRLResponse) ProtoMessage()               {}
func (*FetchCRLResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FetchCRLResponse) GetCrlStatus() *CRLStatus {
	if m != nil {
		return m.CrlStatus
	}
	return nil
}

// ListCAsResponse is returned by ListCAs.
type ListCAsResponse struct {
	Cn []string `protobuf:"bytes,1,rep,name=cn" json:"cn,omitempty"`
}

func (m *ListCAsResponse) Reset()                    { *m = ListCAsResponse{} }
func (m *ListCAsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListCAsResponse) ProtoMessage()               {}
func (*ListCAsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// GetCAStatusRequest identifies a name of CA to fetch.
type GetCAStatusRequest struct {
	Cn string `protobuf:"bytes,1,opt,name=cn" json:"cn,omitempty"`
}

func (m *GetCAStatusRequest) Reset()                    { *m = GetCAStatusRequest{} }
func (m *GetCAStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCAStatusRequest) ProtoMessage()               {}
func (*GetCAStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// GetCAStatusResponse is returned by GetCAStatus method.
//
// If requested CA doesn't exist, all fields are empty.
type GetCAStatusResponse struct {
	Config     *CertificateAuthorityConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	Cert       string                      `protobuf:"bytes,2,opt,name=cert" json:"cert,omitempty"`
	Removed    bool                        `protobuf:"varint,3,opt,name=removed" json:"removed,omitempty"`
	Ready      bool                        `protobuf:"varint,4,opt,name=ready" json:"ready,omitempty"`
	AddedRev   string                      `protobuf:"bytes,5,opt,name=added_rev" json:"added_rev,omitempty"`
	UpdatedRev string                      `protobuf:"bytes,6,opt,name=updated_rev" json:"updated_rev,omitempty"`
	RemovedRev string                      `protobuf:"bytes,7,opt,name=removed_rev" json:"removed_rev,omitempty"`
	CrlStatus  *CRLStatus                  `protobuf:"bytes,8,opt,name=crl_status" json:"crl_status,omitempty"`
}

func (m *GetCAStatusResponse) Reset()                    { *m = GetCAStatusResponse{} }
func (m *GetCAStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*GetCAStatusResponse) ProtoMessage()               {}
func (*GetCAStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetCAStatusResponse) GetConfig() *CertificateAuthorityConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *GetCAStatusResponse) GetCrlStatus() *CRLStatus {
	if m != nil {
		return m.CrlStatus
	}
	return nil
}

// IsRevokedCertRequest contains a name of the CA and a cert serial number.
type IsRevokedCertRequest struct {
	Ca string `protobuf:"bytes,1,opt,name=ca" json:"ca,omitempty"`
	Sn string `protobuf:"bytes,2,opt,name=sn" json:"sn,omitempty"`
}

func (m *IsRevokedCertRequest) Reset()                    { *m = IsRevokedCertRequest{} }
func (m *IsRevokedCertRequest) String() string            { return proto.CompactTextString(m) }
func (*IsRevokedCertRequest) ProtoMessage()               {}
func (*IsRevokedCertRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// IsRevokedCertResponse is returned by IsRevokedCert
type IsRevokedCertResponse struct {
	Revoked bool `protobuf:"varint,1,opt,name=revoked" json:"revoked,omitempty"`
}

func (m *IsRevokedCertResponse) Reset()                    { *m = IsRevokedCertResponse{} }
func (m *IsRevokedCertResponse) String() string            { return proto.CompactTextString(m) }
func (*IsRevokedCertResponse) ProtoMessage()               {}
func (*IsRevokedCertResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// CheckCertificateRequest contains a pem encoded certificate to check.
type CheckCertificateRequest struct {
	CertPem string `protobuf:"bytes,1,opt,name=cert_pem" json:"cert_pem,omitempty"`
}

func (m *CheckCertificateRequest) Reset()                    { *m = CheckCertificateRequest{} }
func (m *CheckCertificateRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckCertificateRequest) ProtoMessage()               {}
func (*CheckCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// CheckCertificateResponse is returned by CheckCertificate.
type CheckCertificateResponse struct {
	IsValid       bool   `protobuf:"varint,1,opt,name=is_valid" json:"is_valid,omitempty"`
	InvalidReason string `protobuf:"bytes,2,opt,name=invalid_reason" json:"invalid_reason,omitempty"`
}

func (m *CheckCertificateResponse) Reset()                    { *m = CheckCertificateResponse{} }
func (m *CheckCertificateResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckCertificateResponse) ProtoMessage()               {}
func (*CheckCertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// CRLStatus describes the latest known state of imported CRL.
type CRLStatus struct {
	LastUpdateTime    *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=last_update_time" json:"last_update_time,omitempty"`
	LastFetchTime     *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=last_fetch_time" json:"last_fetch_time,omitempty"`
	LastFetchEtag     string                      `protobuf:"bytes,3,opt,name=last_fetch_etag" json:"last_fetch_etag,omitempty"`
	RevokedCertsCount int64                       `protobuf:"varint,4,opt,name=revoked_certs_count" json:"revoked_certs_count,omitempty"`
}

func (m *CRLStatus) Reset()                    { *m = CRLStatus{} }
func (m *CRLStatus) String() string            { return proto.CompactTextString(m) }
func (*CRLStatus) ProtoMessage()               {}
func (*CRLStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CRLStatus) GetLastUpdateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *CRLStatus) GetLastFetchTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastFetchTime
	}
	return nil
}

func init() {
	proto.RegisterType((*ImportConfigResponse)(nil), "tokenserver.ImportConfigResponse")
	proto.RegisterType((*FetchCRLRequest)(nil), "tokenserver.FetchCRLRequest")
	proto.RegisterType((*FetchCRLResponse)(nil), "tokenserver.FetchCRLResponse")
	proto.RegisterType((*ListCAsResponse)(nil), "tokenserver.ListCAsResponse")
	proto.RegisterType((*GetCAStatusRequest)(nil), "tokenserver.GetCAStatusRequest")
	proto.RegisterType((*GetCAStatusResponse)(nil), "tokenserver.GetCAStatusResponse")
	proto.RegisterType((*IsRevokedCertRequest)(nil), "tokenserver.IsRevokedCertRequest")
	proto.RegisterType((*IsRevokedCertResponse)(nil), "tokenserver.IsRevokedCertResponse")
	proto.RegisterType((*CheckCertificateRequest)(nil), "tokenserver.CheckCertificateRequest")
	proto.RegisterType((*CheckCertificateResponse)(nil), "tokenserver.CheckCertificateResponse")
	proto.RegisterType((*CRLStatus)(nil), "tokenserver.CRLStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for Admin service

type AdminClient interface {
	// ImportConfig makes the server read its config from luci-config right now.
	//
	// Note that regularly configs are read in background each 5 min. ImportConfig
	// can be used to force config reread immediately. It will block until configs
	// are read.
	ImportConfig(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ImportConfigResponse, error)
	// FetchCRL makes the server fetch a CRL for some CA.
	FetchCRL(ctx context.Context, in *FetchCRLRequest, opts ...grpc.CallOption) (*FetchCRLResponse, error)
	// ListCAs returns a list of Common Names of registered CAs.
	ListCAs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListCAsResponse, error)
	// GetCAStatus returns configuration of some CA defined in the config.
	GetCAStatus(ctx context.Context, in *GetCAStatusRequest, opts ...grpc.CallOption) (*GetCAStatusResponse, error)
	// IsRevokedCert says whether a certificate serial number is in the CRL.
	IsRevokedCert(ctx context.Context, in *IsRevokedCertRequest, opts ...grpc.CallOption) (*IsRevokedCertResponse, error)
	// CheckCertificate says whether a certificate is valid or not.
	CheckCertificate(ctx context.Context, in *CheckCertificateRequest, opts ...grpc.CallOption) (*CheckCertificateResponse, error)
}
type adminPRPCClient struct {
	client *prpccommon.Client
}

func NewAdminPRPCClient(client *prpccommon.Client) AdminClient {
	return &adminPRPCClient{client}
}

func (c *adminPRPCClient) ImportConfig(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ImportConfigResponse, error) {
	out := new(ImportConfigResponse)
	err := c.client.Call(ctx, "tokenserver.Admin", "ImportConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminPRPCClient) FetchCRL(ctx context.Context, in *FetchCRLRequest, opts ...grpc.CallOption) (*FetchCRLResponse, error) {
	out := new(FetchCRLResponse)
	err := c.client.Call(ctx, "tokenserver.Admin", "FetchCRL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminPRPCClient) ListCAs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListCAsResponse, error) {
	out := new(ListCAsResponse)
	err := c.client.Call(ctx, "tokenserver.Admin", "ListCAs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminPRPCClient) GetCAStatus(ctx context.Context, in *GetCAStatusRequest, opts ...grpc.CallOption) (*GetCAStatusResponse, error) {
	out := new(GetCAStatusResponse)
	err := c.client.Call(ctx, "tokenserver.Admin", "GetCAStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminPRPCClient) IsRevokedCert(ctx context.Context, in *IsRevokedCertRequest, opts ...grpc.CallOption) (*IsRevokedCertResponse, error) {
	out := new(IsRevokedCertResponse)
	err := c.client.Call(ctx, "tokenserver.Admin", "IsRevokedCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminPRPCClient) CheckCertificate(ctx context.Context, in *CheckCertificateRequest, opts ...grpc.CallOption) (*CheckCertificateResponse, error) {
	out := new(CheckCertificateResponse)
	err := c.client.Call(ctx, "tokenserver.Admin", "CheckCertificate", in, out, opts...)
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

func (c *adminClient) ImportConfig(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ImportConfigResponse, error) {
	out := new(ImportConfigResponse)
	err := grpc.Invoke(ctx, "/tokenserver.Admin/ImportConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) FetchCRL(ctx context.Context, in *FetchCRLRequest, opts ...grpc.CallOption) (*FetchCRLResponse, error) {
	out := new(FetchCRLResponse)
	err := grpc.Invoke(ctx, "/tokenserver.Admin/FetchCRL", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) ListCAs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListCAsResponse, error) {
	out := new(ListCAsResponse)
	err := grpc.Invoke(ctx, "/tokenserver.Admin/ListCAs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetCAStatus(ctx context.Context, in *GetCAStatusRequest, opts ...grpc.CallOption) (*GetCAStatusResponse, error) {
	out := new(GetCAStatusResponse)
	err := grpc.Invoke(ctx, "/tokenserver.Admin/GetCAStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) IsRevokedCert(ctx context.Context, in *IsRevokedCertRequest, opts ...grpc.CallOption) (*IsRevokedCertResponse, error) {
	out := new(IsRevokedCertResponse)
	err := grpc.Invoke(ctx, "/tokenserver.Admin/IsRevokedCert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CheckCertificate(ctx context.Context, in *CheckCertificateRequest, opts ...grpc.CallOption) (*CheckCertificateResponse, error) {
	out := new(CheckCertificateResponse)
	err := grpc.Invoke(ctx, "/tokenserver.Admin/CheckCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Admin service

type AdminServer interface {
	// ImportConfig makes the server read its config from luci-config right now.
	//
	// Note that regularly configs are read in background each 5 min. ImportConfig
	// can be used to force config reread immediately. It will block until configs
	// are read.
	ImportConfig(context.Context, *google_protobuf.Empty) (*ImportConfigResponse, error)
	// FetchCRL makes the server fetch a CRL for some CA.
	FetchCRL(context.Context, *FetchCRLRequest) (*FetchCRLResponse, error)
	// ListCAs returns a list of Common Names of registered CAs.
	ListCAs(context.Context, *google_protobuf.Empty) (*ListCAsResponse, error)
	// GetCAStatus returns configuration of some CA defined in the config.
	GetCAStatus(context.Context, *GetCAStatusRequest) (*GetCAStatusResponse, error)
	// IsRevokedCert says whether a certificate serial number is in the CRL.
	IsRevokedCert(context.Context, *IsRevokedCertRequest) (*IsRevokedCertResponse, error)
	// CheckCertificate says whether a certificate is valid or not.
	CheckCertificate(context.Context, *CheckCertificateRequest) (*CheckCertificateResponse, error)
}

func RegisterAdminServer(s prpc.Registrar, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_ImportConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AdminServer).ImportConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Admin_FetchCRL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(FetchCRLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AdminServer).FetchCRL(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Admin_ListCAs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AdminServer).ListCAs(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Admin_GetCAStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetCAStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AdminServer).GetCAStatus(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Admin_IsRevokedCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(IsRevokedCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AdminServer).IsRevokedCert(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Admin_CheckCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CheckCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AdminServer).CheckCertificate(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tokenserver.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportConfig",
			Handler:    _Admin_ImportConfig_Handler,
		},
		{
			MethodName: "FetchCRL",
			Handler:    _Admin_FetchCRL_Handler,
		},
		{
			MethodName: "ListCAs",
			Handler:    _Admin_ListCAs_Handler,
		},
		{
			MethodName: "GetCAStatus",
			Handler:    _Admin_GetCAStatus_Handler,
		},
		{
			MethodName: "IsRevokedCert",
			Handler:    _Admin_IsRevokedCert_Handler,
		},
		{
			MethodName: "CheckCertificate",
			Handler:    _Admin_CheckCertificate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0x56, 0xd7, 0xf5, 0xd7, 0xb5, 0xa3, 0xc5, 0x85, 0x2e, 0xca, 0x98, 0x56, 0x22, 0x10, 0x13,
	0xa0, 0x4c, 0xda, 0x90, 0x78, 0x03, 0x55, 0x05, 0xa6, 0x49, 0x7b, 0x40, 0x05, 0xf1, 0x1a, 0x65,
	0x89, 0xdb, 0x5a, 0x6b, 0xe2, 0x60, 0x3b, 0x95, 0xfa, 0x3f, 0xf1, 0x6f, 0x21, 0xfe, 0x0d, 0x1c,
	0xdb, 0xd9, 0x9a, 0xb4, 0xdd, 0xde, 0x92, 0xbb, 0xef, 0x3e, 0xdf, 0x7d, 0x77, 0x1f, 0xb4, 0xfd,
	0x30, 0x22, 0xb1, 0x9b, 0x30, 0x2a, 0x28, 0x6a, 0x0b, 0x7a, 0x8b, 0x63, 0x8e, 0xd9, 0x12, 0x33,
	0xbb, 0x13, 0xd0, 0x78, 0x4a, 0x66, 0x3a, 0x65, 0x1f, 0xcd, 0x28, 0x9d, 0x2d, 0xf0, 0x99, 0xfa,
	0xbb, 0x49, 0xa7, 0x67, 0x38, 0x4a, 0xc4, 0xca, 0x24, 0x4f, 0xca, 0x49, 0x41, 0x22, 0xcc, 0x85,
	0x1f, 0x25, 0x1a, 0xe0, 0x9c, 0xc2, 0xb3, 0xab, 0x28, 0xa1, 0x4c, 0x8c, 0x15, 0xe7, 0x04, 0xf3,
	0x84, 0xca, 0x77, 0x50, 0x0f, 0x9a, 0x0c, 0x2f, 0x09, 0x27, 0x34, 0xb6, 0x2a, 0xc3, 0xca, 0x69,
	0xcb, 0x79, 0x0f, 0xdd, 0x6f, 0x58, 0x04, 0xf3, 0xf1, 0xe4, 0x7a, 0x82, 0x7f, 0xa7, 0x92, 0x05,
	0x01, 0xec, 0x05, 0x26, 0x8d, 0x0e, 0xa0, 0x36, 0xa5, 0x2c, 0xc0, 0xd6, 0x9e, 0xfc, 0x6d, 0x3a,
	0x9f, 0xa0, 0x77, 0x8f, 0x36, 0x9c, 0x6f, 0x01, 0x02, 0xb6, 0xf0, 0xe4, 0xf3, 0x22, 0xe5, 0xaa,
	0xac, 0x7d, 0x3e, 0x70, 0xd7, 0x26, 0x73, 0x25, 0xfa, 0x87, 0xca, 0x3a, 0xc7, 0xd0, 0xbd, 0x26,
	0x5c, 0x8c, 0x47, 0xfc, 0xae, 0x3c, 0x7f, 0xad, 0x2a, 0x9b, 0x19, 0x02, 0xba, 0xc4, 0x32, 0xab,
	0xd1, 0x5b, 0xfa, 0x71, 0xfe, 0x56, 0xa0, 0x5f, 0x80, 0x18, 0x96, 0x8f, 0x50, 0xd7, 0xf2, 0x99,
	0x06, 0xde, 0x14, 0x1b, 0xc0, 0x4c, 0x90, 0x29, 0x09, 0x7c, 0x81, 0x47, 0xa9, 0x98, 0x53, 0x46,
	0xc4, 0x4a, 0x2b, 0x83, 0x3a, 0xb0, 0x1f, 0xc8, 0xac, 0x9a, 0xaf, 0x85, 0xba, 0xd0, 0x60, 0x38,
	0xa2, 0x4b, 0x1c, 0x5a, 0xd5, 0x6c, 0xe0, 0x6c, 0x7e, 0x86, 0xfd, 0x70, 0x65, 0xed, 0xab, 0xdf,
	0xa7, 0xd0, 0xf2, 0xc3, 0x10, 0x87, 0x9e, 0x54, 0xd1, 0xaa, 0xa9, 0x92, 0x3e, 0xb4, 0xd3, 0x24,
	0x94, 0xcc, 0x3a, 0x58, 0xcf, 0x83, 0x86, 0x47, 0x05, 0x1b, 0x2a, 0x58, 0x14, 0xaa, 0xf9, 0xa0,
	0x50, 0xae, 0x5c, 0xa0, 0x9c, 0x6e, 0x29, 0x93, 0x61, 0xd6, 0xfd, 0xba, 0x16, 0xbe, 0xd9, 0x8d,
	0xfc, 0xe6, 0xb1, 0x6e, 0x5c, 0x2e, 0xfc, 0x79, 0x09, 0x6f, 0x84, 0x51, 0x13, 0xa9, 0xb0, 0xaa,
	0x6a, 0x3a, 0xef, 0xe0, 0x70, 0x3c, 0xc7, 0xc1, 0xed, 0x9a, 0x26, 0x39, 0xb9, 0xbc, 0x8e, 0x4c,
	0x0b, 0x2f, 0xc1, 0x91, 0x91, 0xfb, 0x0b, 0x58, 0x9b, 0xe0, 0xfb, 0x5b, 0x22, 0xdc, 0x5b, 0xfa,
	0x0b, 0x62, 0xa8, 0xd1, 0x00, 0x9e, 0x90, 0x58, 0x05, 0xe4, 0xd4, 0x3e, 0xa7, 0x79, 0x73, 0x7f,
	0x2a, 0xd0, 0xba, 0x1b, 0x0d, 0x7d, 0x80, 0xde, 0xc2, 0xe7, 0xc2, 0xd3, 0xaa, 0x79, 0xd9, 0xe9,
	0x9a, 0xa5, 0xd9, 0xae, 0xbe, 0x6b, 0x37, 0xbf, 0x6b, 0xf7, 0x67, 0x7e, 0xd7, 0xe8, 0x02, 0xba,
	0xaa, 0x6a, 0x9a, 0x9d, 0x9f, 0x2e, 0xda, 0x7b, 0xb4, 0xe8, 0xb0, 0x50, 0x84, 0x85, 0x3f, 0x53,
	0x6b, 0x6d, 0xa1, 0x23, 0xe8, 0x1b, 0x55, 0xbc, 0x6c, 0x62, 0xee, 0x05, 0x34, 0x8d, 0x85, 0x5a,
	0x72, 0xf5, 0xfc, 0x5f, 0x15, 0x6a, 0xa3, 0xcc, 0xa5, 0xe8, 0x0a, 0x3a, 0xeb, 0x36, 0x42, 0x83,
	0x8d, 0xb7, 0xbe, 0x66, 0xae, 0xb4, 0x5f, 0x16, 0xb6, 0xb8, 0xd5, 0x79, 0x97, 0xd0, 0xcc, 0x9d,
	0x83, 0x5e, 0x14, 0xe0, 0x25, 0xfb, 0xd9, 0xc7, 0x3b, 0xb2, 0x86, 0xe8, 0x33, 0x34, 0x8c, 0x85,
	0x76, 0xb6, 0x53, 0xe4, 0x2f, 0x1b, 0xee, 0x3b, 0xb4, 0xd7, 0x1c, 0x84, 0x4e, 0x0a, 0xe0, 0x4d,
	0xfb, 0xd9, 0xc3, 0xdd, 0x00, 0xc3, 0xf8, 0x0b, 0x0e, 0x0a, 0xc7, 0x87, 0x4a, 0x7a, 0x6c, 0x39,
	0x64, 0xdb, 0x79, 0x08, 0x62, 0x78, 0x3d, 0xe8, 0x95, 0xaf, 0x0f, 0xbd, 0x2a, 0x1a, 0x66, 0xfb,
	0x25, 0xdb, 0xaf, 0x1f, 0x41, 0xe9, 0x07, 0x6e, 0xea, 0x4a, 0xb8, 0x8b, 0xff, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x8e, 0x8a, 0x49, 0xfe, 0x95, 0x05, 0x00, 0x00,
}
