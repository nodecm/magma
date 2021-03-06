// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feg/protos/swx_proxy.proto

package protos // import "magma/feg/cloud/go/protos"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// SwxErrorCode reflects Experimental-Result values which are 3GPP failures
// to be processed by EPC. Diameter Base Protocol errors are reflected in gRPC status code
type SwxErrorCode int32

const (
	SwxErrorCode_ERROR_UNDEFINED               SwxErrorCode = 0
	SwxErrorCode_IDENTITY_ALREADY_REGISTERED   SwxErrorCode = 5005
	SwxErrorCode_USER_NO_NON_3GPP_SUBSCRIPTION SwxErrorCode = 5450
)

var SwxErrorCode_name = map[int32]string{
	0:    "ERROR_UNDEFINED",
	5005: "IDENTITY_ALREADY_REGISTERED",
	5450: "USER_NO_NON_3GPP_SUBSCRIPTION",
}
var SwxErrorCode_value = map[string]int32{
	"ERROR_UNDEFINED":               0,
	"IDENTITY_ALREADY_REGISTERED":   5005,
	"USER_NO_NON_3GPP_SUBSCRIPTION": 5450,
}

func (x SwxErrorCode) String() string {
	return proto.EnumName(SwxErrorCode_name, int32(x))
}
func (SwxErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_8f915e667fe054f4, []int{0}
}

type AuthenticationScheme int32

const (
	AuthenticationScheme_EAP_AKA       AuthenticationScheme = 0
	AuthenticationScheme_EAP_AKA_PRIME AuthenticationScheme = 1
)

var AuthenticationScheme_name = map[int32]string{
	0: "EAP_AKA",
	1: "EAP_AKA_PRIME",
}
var AuthenticationScheme_value = map[string]int32{
	"EAP_AKA":       0,
	"EAP_AKA_PRIME": 1,
}

func (x AuthenticationScheme) String() string {
	return proto.EnumName(AuthenticationScheme_name, int32(x))
}
func (AuthenticationScheme) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_8f915e667fe054f4, []int{1}
}

// AuthenticationRequest (Section 8.2.2.1)
type AuthenticationRequest struct {
	// Subscriber identifier
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// Number of authentication vectors requested
	SipNumAuthVectors uint32 `protobuf:"varint,2,opt,name=sip_num_auth_vectors,json=sipNumAuthVectors,proto3" json:"sip_num_auth_vectors,omitempty"`
	// EAP-AKA or EAP-AKA'
	AuthenticationScheme AuthenticationScheme `protobuf:"varint,3,opt,name=authentication_scheme,json=authenticationScheme,proto3,enum=magma.feg.AuthenticationScheme" json:"authentication_scheme,omitempty"`
	// Concatenation of RAND and AUTS in the case of resync
	ResyncInfo []byte `protobuf:"bytes,4,opt,name=resync_info,json=resyncInfo,proto3" json:"resync_info,omitempty"`
	// Send an additional SAR message to the HSS to retrieve user profile params
	RetrieveUserProfile  bool     `protobuf:"varint,5,opt,name=retrieve_user_profile,json=retrieveUserProfile,proto3" json:"retrieve_user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationRequest) Reset()         { *m = AuthenticationRequest{} }
func (m *AuthenticationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticationRequest) ProtoMessage()    {}
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_8f915e667fe054f4, []int{0}
}
func (m *AuthenticationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationRequest.Unmarshal(m, b)
}
func (m *AuthenticationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationRequest.Marshal(b, m, deterministic)
}
func (dst *AuthenticationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationRequest.Merge(dst, src)
}
func (m *AuthenticationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticationRequest.Size(m)
}
func (m *AuthenticationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationRequest proto.InternalMessageInfo

func (m *AuthenticationRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AuthenticationRequest) GetSipNumAuthVectors() uint32 {
	if m != nil {
		return m.SipNumAuthVectors
	}
	return 0
}

func (m *AuthenticationRequest) GetAuthenticationScheme() AuthenticationScheme {
	if m != nil {
		return m.AuthenticationScheme
	}
	return AuthenticationScheme_EAP_AKA
}

func (m *AuthenticationRequest) GetResyncInfo() []byte {
	if m != nil {
		return m.ResyncInfo
	}
	return nil
}

func (m *AuthenticationRequest) GetRetrieveUserProfile() bool {
	if m != nil {
		return m.RetrieveUserProfile
	}
	return false
}

// MultimediaAuthenticationAnswer (Section 8.2.2.1)
type AuthenticationAnswer struct {
	// Subscriber identifier
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// For details about fields read 3GPP 29.273
	SipAuthVectors       []*AuthenticationAnswer_SIPAuthVector `protobuf:"bytes,2,rep,name=sip_auth_vectors,json=sipAuthVectors,proto3" json:"sip_auth_vectors,omitempty"`
	UserProfile          *AuthenticationAnswer_UserProfile     `protobuf:"bytes,3,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *AuthenticationAnswer) Reset()         { *m = AuthenticationAnswer{} }
func (m *AuthenticationAnswer) String() string { return proto.CompactTextString(m) }
func (*AuthenticationAnswer) ProtoMessage()    {}
func (*AuthenticationAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_8f915e667fe054f4, []int{1}
}
func (m *AuthenticationAnswer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationAnswer.Unmarshal(m, b)
}
func (m *AuthenticationAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationAnswer.Marshal(b, m, deterministic)
}
func (dst *AuthenticationAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationAnswer.Merge(dst, src)
}
func (m *AuthenticationAnswer) XXX_Size() int {
	return xxx_messageInfo_AuthenticationAnswer.Size(m)
}
func (m *AuthenticationAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationAnswer proto.InternalMessageInfo

func (m *AuthenticationAnswer) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AuthenticationAnswer) GetSipAuthVectors() []*AuthenticationAnswer_SIPAuthVector {
	if m != nil {
		return m.SipAuthVectors
	}
	return nil
}

func (m *AuthenticationAnswer) GetUserProfile() *AuthenticationAnswer_UserProfile {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

// Only for EAP-AKA/EAP-AKA'
type AuthenticationAnswer_SIPAuthVector struct {
	// Contains one of EAP-AKA or EAP-AKA'
	AuthenticationScheme AuthenticationScheme `protobuf:"varint,1,opt,name=authentication_scheme,json=authenticationScheme,proto3,enum=magma.feg.AuthenticationScheme" json:"authentication_scheme,omitempty"`
	// Concatenation of challenge RAND and token AUTN
	RandAutn []byte `protobuf:"bytes,2,opt,name=rand_autn,json=randAutn,proto3" json:"rand_autn,omitempty"`
	// Expected response
	Xres []byte `protobuf:"bytes,3,opt,name=xres,proto3" json:"xres,omitempty"`
	// Confidentiality Key
	ConfidentialityKey []byte `protobuf:"bytes,4,opt,name=confidentiality_key,json=confidentialityKey,proto3" json:"confidentiality_key,omitempty"`
	// Integrity Key
	IntegrityKey         []byte   `protobuf:"bytes,5,opt,name=integrity_key,json=integrityKey,proto3" json:"integrity_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationAnswer_SIPAuthVector) Reset()         { *m = AuthenticationAnswer_SIPAuthVector{} }
func (m *AuthenticationAnswer_SIPAuthVector) String() string { return proto.CompactTextString(m) }
func (*AuthenticationAnswer_SIPAuthVector) ProtoMessage()    {}
func (*AuthenticationAnswer_SIPAuthVector) Descriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_8f915e667fe054f4, []int{1, 0}
}
func (m *AuthenticationAnswer_SIPAuthVector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationAnswer_SIPAuthVector.Unmarshal(m, b)
}
func (m *AuthenticationAnswer_SIPAuthVector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationAnswer_SIPAuthVector.Marshal(b, m, deterministic)
}
func (dst *AuthenticationAnswer_SIPAuthVector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationAnswer_SIPAuthVector.Merge(dst, src)
}
func (m *AuthenticationAnswer_SIPAuthVector) XXX_Size() int {
	return xxx_messageInfo_AuthenticationAnswer_SIPAuthVector.Size(m)
}
func (m *AuthenticationAnswer_SIPAuthVector) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationAnswer_SIPAuthVector.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationAnswer_SIPAuthVector proto.InternalMessageInfo

func (m *AuthenticationAnswer_SIPAuthVector) GetAuthenticationScheme() AuthenticationScheme {
	if m != nil {
		return m.AuthenticationScheme
	}
	return AuthenticationScheme_EAP_AKA
}

func (m *AuthenticationAnswer_SIPAuthVector) GetRandAutn() []byte {
	if m != nil {
		return m.RandAutn
	}
	return nil
}

func (m *AuthenticationAnswer_SIPAuthVector) GetXres() []byte {
	if m != nil {
		return m.Xres
	}
	return nil
}

func (m *AuthenticationAnswer_SIPAuthVector) GetConfidentialityKey() []byte {
	if m != nil {
		return m.ConfidentialityKey
	}
	return nil
}

func (m *AuthenticationAnswer_SIPAuthVector) GetIntegrityKey() []byte {
	if m != nil {
		return m.IntegrityKey
	}
	return nil
}

type AuthenticationAnswer_UserProfile struct {
	// MSISDN from HSS
	Msisdn               string   `protobuf:"bytes,1,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationAnswer_UserProfile) Reset()         { *m = AuthenticationAnswer_UserProfile{} }
func (m *AuthenticationAnswer_UserProfile) String() string { return proto.CompactTextString(m) }
func (*AuthenticationAnswer_UserProfile) ProtoMessage()    {}
func (*AuthenticationAnswer_UserProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_8f915e667fe054f4, []int{1, 1}
}
func (m *AuthenticationAnswer_UserProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationAnswer_UserProfile.Unmarshal(m, b)
}
func (m *AuthenticationAnswer_UserProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationAnswer_UserProfile.Marshal(b, m, deterministic)
}
func (dst *AuthenticationAnswer_UserProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationAnswer_UserProfile.Merge(dst, src)
}
func (m *AuthenticationAnswer_UserProfile) XXX_Size() int {
	return xxx_messageInfo_AuthenticationAnswer_UserProfile.Size(m)
}
func (m *AuthenticationAnswer_UserProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationAnswer_UserProfile.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationAnswer_UserProfile proto.InternalMessageInfo

func (m *AuthenticationAnswer_UserProfile) GetMsisdn() string {
	if m != nil {
		return m.Msisdn
	}
	return ""
}

// RegistrationRequest:
// ServerAssignmentRequest with ServerAssignmentType set to REGISTRATION (Section 8.2.2.3)
type RegistrationRequest struct {
	// Subscriber identifier
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrationRequest) Reset()         { *m = RegistrationRequest{} }
func (m *RegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*RegistrationRequest) ProtoMessage()    {}
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_8f915e667fe054f4, []int{2}
}
func (m *RegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationRequest.Unmarshal(m, b)
}
func (m *RegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationRequest.Marshal(b, m, deterministic)
}
func (dst *RegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationRequest.Merge(dst, src)
}
func (m *RegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_RegistrationRequest.Size(m)
}
func (m *RegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationRequest proto.InternalMessageInfo

func (m *RegistrationRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// RegistrationAnswer:
// ServerAssignmentAnswer with ServerAssignmentType set to REGISTRATION (Section 8.2.2.3)
type RegistrationAnswer struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrationAnswer) Reset()         { *m = RegistrationAnswer{} }
func (m *RegistrationAnswer) String() string { return proto.CompactTextString(m) }
func (*RegistrationAnswer) ProtoMessage()    {}
func (*RegistrationAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_swx_proxy_8f915e667fe054f4, []int{3}
}
func (m *RegistrationAnswer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationAnswer.Unmarshal(m, b)
}
func (m *RegistrationAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationAnswer.Marshal(b, m, deterministic)
}
func (dst *RegistrationAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationAnswer.Merge(dst, src)
}
func (m *RegistrationAnswer) XXX_Size() int {
	return xxx_messageInfo_RegistrationAnswer.Size(m)
}
func (m *RegistrationAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationAnswer proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AuthenticationRequest)(nil), "magma.feg.AuthenticationRequest")
	proto.RegisterType((*AuthenticationAnswer)(nil), "magma.feg.AuthenticationAnswer")
	proto.RegisterType((*AuthenticationAnswer_SIPAuthVector)(nil), "magma.feg.AuthenticationAnswer.SIPAuthVector")
	proto.RegisterType((*AuthenticationAnswer_UserProfile)(nil), "magma.feg.AuthenticationAnswer.UserProfile")
	proto.RegisterType((*RegistrationRequest)(nil), "magma.feg.RegistrationRequest")
	proto.RegisterType((*RegistrationAnswer)(nil), "magma.feg.RegistrationAnswer")
	proto.RegisterEnum("magma.feg.SwxErrorCode", SwxErrorCode_name, SwxErrorCode_value)
	proto.RegisterEnum("magma.feg.AuthenticationScheme", AuthenticationScheme_name, AuthenticationScheme_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SwxProxyClient is the client API for SwxProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SwxProxyClient interface {
	// Retrieve authentication vectors from the HSS using MAR/MAA
	Authenticate(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationAnswer, error)
	// Register the AAA server serving a user to the HSS using SAR/SAA
	Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationAnswer, error)
}

type swxProxyClient struct {
	cc *grpc.ClientConn
}

func NewSwxProxyClient(cc *grpc.ClientConn) SwxProxyClient {
	return &swxProxyClient{cc}
}

func (c *swxProxyClient) Authenticate(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationAnswer, error) {
	out := new(AuthenticationAnswer)
	err := c.cc.Invoke(ctx, "/magma.feg.SwxProxy/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swxProxyClient) Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationAnswer, error) {
	out := new(RegistrationAnswer)
	err := c.cc.Invoke(ctx, "/magma.feg.SwxProxy/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SwxProxyServer is the server API for SwxProxy service.
type SwxProxyServer interface {
	// Retrieve authentication vectors from the HSS using MAR/MAA
	Authenticate(context.Context, *AuthenticationRequest) (*AuthenticationAnswer, error)
	// Register the AAA server serving a user to the HSS using SAR/SAA
	Register(context.Context, *RegistrationRequest) (*RegistrationAnswer, error)
}

func RegisterSwxProxyServer(s *grpc.Server, srv SwxProxyServer) {
	s.RegisterService(&_SwxProxy_serviceDesc, srv)
}

func _SwxProxy_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwxProxyServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.SwxProxy/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwxProxyServer).Authenticate(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwxProxy_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwxProxyServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.feg.SwxProxy/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwxProxyServer).Register(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SwxProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.feg.SwxProxy",
	HandlerType: (*SwxProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _SwxProxy_Authenticate_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _SwxProxy_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feg/protos/swx_proxy.proto",
}

func init() {
	proto.RegisterFile("feg/protos/swx_proxy.proto", fileDescriptor_swx_proxy_8f915e667fe054f4)
}

var fileDescriptor_swx_proxy_8f915e667fe054f4 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdf, 0x6e, 0xda, 0x4a,
	0x10, 0xc6, 0xe3, 0xfc, 0x3b, 0x30, 0x40, 0x0e, 0x59, 0xc8, 0x11, 0x07, 0x94, 0x06, 0x51, 0x55,
	0x45, 0xa9, 0x0a, 0x12, 0x91, 0x7a, 0xef, 0x04, 0x37, 0xb2, 0x68, 0x8d, 0xb5, 0x86, 0x56, 0xe9,
	0xcd, 0xca, 0x85, 0x81, 0xac, 0x1a, 0xaf, 0xe9, 0xae, 0x1d, 0xe0, 0x21, 0x2a, 0xf5, 0x41, 0xfa,
	0x24, 0x7d, 0x96, 0x5e, 0xf5, 0x09, 0x2a, 0x1b, 0x27, 0x81, 0x28, 0x34, 0x95, 0x7a, 0x65, 0xef,
	0x37, 0xe3, 0x6f, 0x67, 0x7e, 0xb3, 0x5e, 0x28, 0x8f, 0x70, 0xdc, 0x9c, 0x48, 0x3f, 0xf0, 0x55,
	0x53, 0x4d, 0x67, 0x6c, 0x22, 0xfd, 0xd9, 0xbc, 0x11, 0x0b, 0x24, 0xed, 0xb9, 0x63, 0xcf, 0x6d,
	0x8c, 0x70, 0x5c, 0xfb, 0xba, 0x09, 0x07, 0x7a, 0x18, 0x5c, 0xa2, 0x08, 0xf8, 0xc0, 0x0d, 0xb8,
	0x2f, 0x28, 0x7e, 0x0e, 0x51, 0x05, 0xa4, 0x02, 0xe9, 0x50, 0xa1, 0x64, 0xc2, 0xf5, 0xb0, 0xa4,
	0x55, 0xb5, 0x7a, 0x9a, 0xa6, 0x22, 0xc1, 0x72, 0x3d, 0x24, 0x4d, 0x28, 0x2a, 0x3e, 0x61, 0x22,
	0xf4, 0x98, 0x1b, 0x06, 0x97, 0xec, 0x1a, 0x07, 0x81, 0x2f, 0x55, 0x69, 0xb3, 0xaa, 0xd5, 0x73,
	0x74, 0x5f, 0xf1, 0x89, 0x15, 0x7a, 0x91, 0xef, 0xbb, 0x45, 0x80, 0xf4, 0xe0, 0xc0, 0x5d, 0xd9,
	0x86, 0xa9, 0xc1, 0x25, 0x7a, 0x58, 0xda, 0xaa, 0x6a, 0xf5, 0xbd, 0xd6, 0x51, 0xe3, 0xb6, 0xa4,
	0xc6, 0x6a, 0x39, 0x4e, 0x9c, 0x46, 0x8b, 0xee, 0x03, 0x2a, 0x39, 0x82, 0x8c, 0x44, 0x35, 0x17,
	0x03, 0xc6, 0xc5, 0xc8, 0x2f, 0x6d, 0x57, 0xb5, 0x7a, 0x96, 0xc2, 0x42, 0x32, 0xc5, 0xc8, 0x27,
	0x2d, 0x38, 0x90, 0x18, 0x48, 0x8e, 0xd7, 0xc8, 0xe2, 0x6e, 0x26, 0xd2, 0x1f, 0xf1, 0x2b, 0x2c,
	0xed, 0x54, 0xb5, 0x7a, 0x8a, 0x16, 0x6e, 0x82, 0x7d, 0x85, 0xd2, 0x5e, 0x84, 0x6a, 0x3f, 0xb7,
	0xa0, 0xb8, 0x5a, 0x83, 0x2e, 0xd4, 0x14, 0xe5, 0xef, 0x89, 0xbc, 0x87, 0x7c, 0x44, 0xe4, 0x1e,
	0x8d, 0xad, 0x7a, 0xa6, 0xf5, 0x72, 0x6d, 0x6f, 0x0b, 0xdf, 0x86, 0x63, 0xda, 0x77, 0xa8, 0xe8,
	0x9e, 0xe2, 0x93, 0x65, 0x72, 0x16, 0x64, 0x57, 0x2a, 0x8f, 0x80, 0x65, 0x5a, 0x2f, 0x1e, 0x33,
	0x5d, 0xea, 0x88, 0x66, 0xc2, 0xbb, 0x45, 0xf9, 0x87, 0x06, 0xb9, 0x95, 0x1d, 0xd7, 0xcf, 0x46,
	0xfb, 0x9b, 0xd9, 0x54, 0x20, 0x2d, 0x5d, 0x31, 0x8c, 0x88, 0x88, 0xf8, 0x5c, 0x64, 0x69, 0x2a,
	0x12, 0xf4, 0x30, 0x10, 0x84, 0xc0, 0xf6, 0x4c, 0xa2, 0x8a, 0x9b, 0xc9, 0xd2, 0xf8, 0x9d, 0x34,
	0xa1, 0x30, 0xf0, 0xc5, 0x88, 0x0f, 0x23, 0x2b, 0xf7, 0x8a, 0x07, 0x73, 0xf6, 0x09, 0xe7, 0xc9,
	0x50, 0xc9, 0xbd, 0x50, 0x07, 0xe7, 0xe4, 0x29, 0xe4, 0xb8, 0x08, 0x70, 0x2c, 0x6f, 0x52, 0x77,
	0xe2, 0xd4, 0xec, 0xad, 0xd8, 0xc1, 0x79, 0xf9, 0x19, 0x64, 0x96, 0x50, 0x90, 0xff, 0x60, 0xd7,
	0x53, 0x5c, 0x0d, 0x45, 0x32, 0xc0, 0x64, 0x55, 0x6b, 0x41, 0x81, 0xe2, 0x98, 0xab, 0x40, 0xfe,
	0xf1, 0x4f, 0x50, 0x2b, 0x02, 0x59, 0xfe, 0x66, 0x01, 0xfe, 0x98, 0x43, 0xd6, 0x99, 0xce, 0x0c,
	0x29, 0x7d, 0x79, 0xe6, 0x0f, 0x91, 0x14, 0xe0, 0x5f, 0x83, 0xd2, 0x2e, 0x65, 0x7d, 0xab, 0x6d,
	0xbc, 0x36, 0x2d, 0xa3, 0x9d, 0xdf, 0x20, 0x55, 0xa8, 0x98, 0x6d, 0xc3, 0xea, 0x99, 0xbd, 0x0b,
	0xa6, 0xbf, 0xa1, 0x86, 0xde, 0xbe, 0x60, 0xd4, 0x38, 0x37, 0x9d, 0x9e, 0x41, 0x8d, 0x76, 0xfe,
	0xcb, 0x73, 0x52, 0x83, 0xc3, 0xbe, 0x63, 0x50, 0x66, 0x75, 0x99, 0xd5, 0xb5, 0xd8, 0xc9, 0xb9,
	0x6d, 0x33, 0xa7, 0x7f, 0xea, 0x9c, 0x51, 0xd3, 0xee, 0x99, 0x5d, 0x2b, 0xff, 0xfd, 0xf8, 0xf8,
	0xd5, 0xfd, 0x83, 0x9a, 0xa0, 0xcf, 0xc0, 0x3f, 0x86, 0x6e, 0x33, 0xbd, 0xa3, 0xe7, 0x37, 0xc8,
	0x3e, 0xe4, 0x92, 0x05, 0xb3, 0xa9, 0xf9, 0xd6, 0xc8, 0x6b, 0xad, 0x6f, 0x1a, 0xa4, 0x9c, 0xe9,
	0xcc, 0x8e, 0xae, 0x04, 0xe2, 0x40, 0x76, 0xc9, 0x04, 0x49, 0x75, 0xed, 0xb8, 0x13, 0x28, 0xe5,
	0xa3, 0x47, 0xce, 0x5e, 0x6d, 0x83, 0x74, 0x20, 0xb5, 0x40, 0x83, 0x92, 0x3c, 0x59, 0x4a, 0x7f,
	0x80, 0x71, 0xf9, 0x70, 0x4d, 0xfc, 0xc6, 0xec, 0xb4, 0xf2, 0xe1, 0xff, 0x38, 0xa3, 0x19, 0x5d,
	0x69, 0x83, 0x2b, 0x3f, 0x1c, 0x36, 0xc7, 0x7e, 0x72, 0xb7, 0x7d, 0xdc, 0x8d, 0x9f, 0x27, 0xbf,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xf3, 0x09, 0xd1, 0xf0, 0x04, 0x00, 0x00,
}
