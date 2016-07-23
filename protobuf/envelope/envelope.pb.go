// Code generated by protoc-gen-go.
// source: envelope.proto
// DO NOT EDIT!

/*
Package envelope is a generated protocol buffer package.

It is generated from these files:
	envelope.proto

It has these top-level messages:
	Request
	Response
	AuthTicket
*/
package envelope

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import enum "github.com/vquintin/pokemongo/protobuf/enum"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Direction        *enum.RpcDirection  `protobuf:"varint,1,req,name=direction,enum=enum.RpcDirection" json:"direction,omitempty"`
	RpcId            *int64              `protobuf:"varint,3,opt,name=rpc_id,json=rpcId" json:"rpc_id,omitempty"`
	Requests         []*Request_Requests `protobuf:"bytes,4,rep,name=requests" json:"requests,omitempty"`
	Unknown6         *Request_Unknown6   `protobuf:"bytes,6,opt,name=unknown6" json:"unknown6,omitempty"`
	Latitude         *uint64             `protobuf:"fixed64,7,opt,name=latitude" json:"latitude,omitempty"`
	Longitude        *uint64             `protobuf:"fixed64,8,opt,name=longitude" json:"longitude,omitempty"`
	Altitude         *uint64             `protobuf:"fixed64,9,opt,name=altitude" json:"altitude,omitempty"`
	Auth             *Request_AuthInfo   `protobuf:"bytes,10,opt,name=auth" json:"auth,omitempty"`
	AuthTicket       *AuthTicket         `protobuf:"bytes,11,opt,name=auth_ticket,json=authTicket" json:"auth_ticket,omitempty"`
	Unknown12        *int64              `protobuf:"varint,12,opt,name=unknown12" json:"unknown12,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetDirection() enum.RpcDirection {
	if m != nil && m.Direction != nil {
		return *m.Direction
	}
	return enum.RpcDirection_UNKNOWN
}

func (m *Request) GetRpcId() int64 {
	if m != nil && m.RpcId != nil {
		return *m.RpcId
	}
	return 0
}

func (m *Request) GetRequests() []*Request_Requests {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *Request) GetUnknown6() *Request_Unknown6 {
	if m != nil {
		return m.Unknown6
	}
	return nil
}

func (m *Request) GetLatitude() uint64 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

func (m *Request) GetLongitude() uint64 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *Request) GetAltitude() uint64 {
	if m != nil && m.Altitude != nil {
		return *m.Altitude
	}
	return 0
}

func (m *Request) GetAuth() *Request_AuthInfo {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *Request) GetAuthTicket() *AuthTicket {
	if m != nil {
		return m.AuthTicket
	}
	return nil
}

func (m *Request) GetUnknown12() int64 {
	if m != nil && m.Unknown12 != nil {
		return *m.Unknown12
	}
	return 0
}

type Request_Requests struct {
	Type             *enum.RequestMethod `protobuf:"varint,1,req,name=type,enum=enum.RequestMethod" json:"type,omitempty"`
	Parameters       []byte              `protobuf:"bytes,2,opt,name=parameters" json:"parameters,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Request_Requests) Reset()                    { *m = Request_Requests{} }
func (m *Request_Requests) String() string            { return proto.CompactTextString(m) }
func (*Request_Requests) ProtoMessage()               {}
func (*Request_Requests) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Request_Requests) GetType() enum.RequestMethod {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return enum.RequestMethod_METHOD_UNSET
}

func (m *Request_Requests) GetParameters() []byte {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type Request_AuthInfo struct {
	Provider         *string               `protobuf:"bytes,1,req,name=provider" json:"provider,omitempty"`
	Token            *Request_AuthInfo_JWT `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Request_AuthInfo) Reset()                    { *m = Request_AuthInfo{} }
func (m *Request_AuthInfo) String() string            { return proto.CompactTextString(m) }
func (*Request_AuthInfo) ProtoMessage()               {}
func (*Request_AuthInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *Request_AuthInfo) GetProvider() string {
	if m != nil && m.Provider != nil {
		return *m.Provider
	}
	return ""
}

func (m *Request_AuthInfo) GetToken() *Request_AuthInfo_JWT {
	if m != nil {
		return m.Token
	}
	return nil
}

type Request_AuthInfo_JWT struct {
	Contents         *string `protobuf:"bytes,1,req,name=contents" json:"contents,omitempty"`
	Unknown13        *int32  `protobuf:"varint,2,req,name=unknown13" json:"unknown13,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request_AuthInfo_JWT) Reset()                    { *m = Request_AuthInfo_JWT{} }
func (m *Request_AuthInfo_JWT) String() string            { return proto.CompactTextString(m) }
func (*Request_AuthInfo_JWT) ProtoMessage()               {}
func (*Request_AuthInfo_JWT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1, 0} }

func (m *Request_AuthInfo_JWT) GetContents() string {
	if m != nil && m.Contents != nil {
		return *m.Contents
	}
	return ""
}

func (m *Request_AuthInfo_JWT) GetUnknown13() int32 {
	if m != nil && m.Unknown13 != nil {
		return *m.Unknown13
	}
	return 0
}

type Request_Unknown3 struct {
	Unknown4         *string `protobuf:"bytes,1,req,name=unknown4" json:"unknown4,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request_Unknown3) Reset()                    { *m = Request_Unknown3{} }
func (m *Request_Unknown3) String() string            { return proto.CompactTextString(m) }
func (*Request_Unknown3) ProtoMessage()               {}
func (*Request_Unknown3) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func (m *Request_Unknown3) GetUnknown4() string {
	if m != nil && m.Unknown4 != nil {
		return *m.Unknown4
	}
	return ""
}

type Request_Unknown6 struct {
	Unknown1         *int32                     `protobuf:"varint,1,req,name=unknown1" json:"unknown1,omitempty"`
	Unknown2         *Request_Unknown6_Unknown2 `protobuf:"bytes,2,req,name=unknown2" json:"unknown2,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *Request_Unknown6) Reset()                    { *m = Request_Unknown6{} }
func (m *Request_Unknown6) String() string            { return proto.CompactTextString(m) }
func (*Request_Unknown6) ProtoMessage()               {}
func (*Request_Unknown6) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3} }

func (m *Request_Unknown6) GetUnknown1() int32 {
	if m != nil && m.Unknown1 != nil {
		return *m.Unknown1
	}
	return 0
}

func (m *Request_Unknown6) GetUnknown2() *Request_Unknown6_Unknown2 {
	if m != nil {
		return m.Unknown2
	}
	return nil
}

type Request_Unknown6_Unknown2 struct {
	Unknown1         []byte `protobuf:"bytes,1,req,name=unknown1" json:"unknown1,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Request_Unknown6_Unknown2) Reset()                    { *m = Request_Unknown6_Unknown2{} }
func (m *Request_Unknown6_Unknown2) String() string            { return proto.CompactTextString(m) }
func (*Request_Unknown6_Unknown2) ProtoMessage()               {}
func (*Request_Unknown6_Unknown2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3, 0} }

func (m *Request_Unknown6_Unknown2) GetUnknown1() []byte {
	if m != nil {
		return m.Unknown1
	}
	return nil
}

type Response struct {
	Direction        *enum.RpcDirection `protobuf:"varint,1,req,name=direction,enum=enum.RpcDirection" json:"direction,omitempty"`
	Unknown2         *int64             `protobuf:"varint,2,opt,name=unknown2" json:"unknown2,omitempty"`
	ApiUrl           *string            `protobuf:"bytes,3,opt,name=api_url,json=apiUrl" json:"api_url,omitempty"`
	Unknown6         *Response_Unknown6 `protobuf:"bytes,6,req,name=unknown6" json:"unknown6,omitempty"`
	AuthTicket       *AuthTicket        `protobuf:"bytes,7,opt,name=auth_ticket,json=authTicket" json:"auth_ticket,omitempty"`
	Responses        [][]byte           `protobuf:"bytes,100,rep,name=responses" json:"responses,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetDirection() enum.RpcDirection {
	if m != nil && m.Direction != nil {
		return *m.Direction
	}
	return enum.RpcDirection_UNKNOWN
}

func (m *Response) GetUnknown2() int64 {
	if m != nil && m.Unknown2 != nil {
		return *m.Unknown2
	}
	return 0
}

func (m *Response) GetApiUrl() string {
	if m != nil && m.ApiUrl != nil {
		return *m.ApiUrl
	}
	return ""
}

func (m *Response) GetUnknown6() *Response_Unknown6 {
	if m != nil {
		return m.Unknown6
	}
	return nil
}

func (m *Response) GetAuthTicket() *AuthTicket {
	if m != nil {
		return m.AuthTicket
	}
	return nil
}

func (m *Response) GetResponses() [][]byte {
	if m != nil {
		return m.Responses
	}
	return nil
}

type Response_Unknown6 struct {
	Unknown1         *int32                      `protobuf:"varint,1,req,name=unknown1" json:"unknown1,omitempty"`
	Unknown2         *Response_Unknown6_Unknown2 `protobuf:"bytes,2,req,name=unknown2" json:"unknown2,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *Response_Unknown6) Reset()                    { *m = Response_Unknown6{} }
func (m *Response_Unknown6) String() string            { return proto.CompactTextString(m) }
func (*Response_Unknown6) ProtoMessage()               {}
func (*Response_Unknown6) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Response_Unknown6) GetUnknown1() int32 {
	if m != nil && m.Unknown1 != nil {
		return *m.Unknown1
	}
	return 0
}

func (m *Response_Unknown6) GetUnknown2() *Response_Unknown6_Unknown2 {
	if m != nil {
		return m.Unknown2
	}
	return nil
}

type Response_Unknown6_Unknown2 struct {
	Unknown1         []byte `protobuf:"bytes,1,req,name=unknown1" json:"unknown1,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Response_Unknown6_Unknown2) Reset()         { *m = Response_Unknown6_Unknown2{} }
func (m *Response_Unknown6_Unknown2) String() string { return proto.CompactTextString(m) }
func (*Response_Unknown6_Unknown2) ProtoMessage()    {}
func (*Response_Unknown6_Unknown2) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0, 0}
}

func (m *Response_Unknown6_Unknown2) GetUnknown1() []byte {
	if m != nil {
		return m.Unknown1
	}
	return nil
}

type Response_Unknown7 struct {
	Unknown71        []byte `protobuf:"bytes,1,opt,name=unknown71" json:"unknown71,omitempty"`
	Unknown72        *int64 `protobuf:"varint,2,opt,name=unknown72" json:"unknown72,omitempty"`
	Unknown73        []byte `protobuf:"bytes,3,opt,name=unknown73" json:"unknown73,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Response_Unknown7) Reset()                    { *m = Response_Unknown7{} }
func (m *Response_Unknown7) String() string            { return proto.CompactTextString(m) }
func (*Response_Unknown7) ProtoMessage()               {}
func (*Response_Unknown7) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

func (m *Response_Unknown7) GetUnknown71() []byte {
	if m != nil {
		return m.Unknown71
	}
	return nil
}

func (m *Response_Unknown7) GetUnknown72() int64 {
	if m != nil && m.Unknown72 != nil {
		return *m.Unknown72
	}
	return 0
}

func (m *Response_Unknown7) GetUnknown73() []byte {
	if m != nil {
		return m.Unknown73
	}
	return nil
}

type AuthTicket struct {
	Start             []byte  `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	ExpireTimestampMs *uint64 `protobuf:"varint,2,opt,name=expire_timestamp_ms,json=expireTimestampMs" json:"expire_timestamp_ms,omitempty"`
	End               []byte  `protobuf:"bytes,3,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *AuthTicket) Reset()                    { *m = AuthTicket{} }
func (m *AuthTicket) String() string            { return proto.CompactTextString(m) }
func (*AuthTicket) ProtoMessage()               {}
func (*AuthTicket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthTicket) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *AuthTicket) GetExpireTimestampMs() uint64 {
	if m != nil && m.ExpireTimestampMs != nil {
		return *m.ExpireTimestampMs
	}
	return 0
}

func (m *AuthTicket) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Request_Requests)(nil), "Request.Requests")
	proto.RegisterType((*Request_AuthInfo)(nil), "Request.AuthInfo")
	proto.RegisterType((*Request_AuthInfo_JWT)(nil), "Request.AuthInfo.JWT")
	proto.RegisterType((*Request_Unknown3)(nil), "Request.Unknown3")
	proto.RegisterType((*Request_Unknown6)(nil), "Request.Unknown6")
	proto.RegisterType((*Request_Unknown6_Unknown2)(nil), "Request.Unknown6.Unknown2")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*Response_Unknown6)(nil), "Response.Unknown6")
	proto.RegisterType((*Response_Unknown6_Unknown2)(nil), "Response.Unknown6.Unknown2")
	proto.RegisterType((*Response_Unknown7)(nil), "Response.Unknown7")
	proto.RegisterType((*AuthTicket)(nil), "AuthTicket")
}

func init() { proto.RegisterFile("envelope.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0x97, 0xfe, 0xa4, 0xa7, 0xd5, 0xc4, 0x3c, 0x26, 0xa2, 0x80, 0xd0, 0x34, 0x09, 0x98,
	0x04, 0xa4, 0x34, 0x45, 0xed, 0x25, 0x42, 0xe2, 0x66, 0x48, 0xbb, 0x31, 0x9d, 0xb8, 0xac, 0xb2,
	0xc4, 0x5b, 0xa3, 0x26, 0x76, 0x96, 0x38, 0x05, 0x5e, 0x00, 0x5e, 0x00, 0xf1, 0x76, 0xbc, 0x0b,
	0xb6, 0x63, 0x27, 0xd9, 0x8a, 0x04, 0x82, 0x9b, 0x36, 0xe7, 0x7c, 0xdf, 0xf9, 0x7c, 0x7e, 0x61,
	0x9f, 0xd0, 0x2d, 0x49, 0x58, 0x46, 0xbc, 0x2c, 0x67, 0x9c, 0xb9, 0xf3, 0xeb, 0x98, 0xaf, 0xcb,
	0x4b, 0x2f, 0x64, 0xe9, 0x64, 0x7b, 0x53, 0xc6, 0x94, 0xc7, 0x74, 0x92, 0xb1, 0x0d, 0x49, 0x19,
	0xbd, 0x66, 0x13, 0x45, 0xba, 0x2c, 0xaf, 0x26, 0x84, 0x96, 0xa9, 0xfa, 0xa9, 0xe2, 0x4e, 0x7e,
	0xf4, 0x61, 0x80, 0xc9, 0x4d, 0x49, 0x0a, 0x8e, 0x5e, 0xc1, 0x30, 0x8a, 0x73, 0x12, 0xf2, 0x98,
	0x51, 0xa7, 0x73, 0xbc, 0x77, 0xba, 0xef, 0x23, 0x4f, 0x71, 0x71, 0x16, 0xbe, 0x33, 0x08, 0x6e,
	0x48, 0xe8, 0x08, 0xfa, 0x79, 0x16, 0xae, 0xe2, 0xc8, 0xb1, 0x8e, 0x3b, 0xa7, 0x16, 0xee, 0x09,
	0xeb, 0x2c, 0x42, 0x2f, 0xc1, 0xce, 0x2b, 0xcd, 0xc2, 0xe9, 0x1e, 0x5b, 0xa7, 0x23, 0xff, 0xc0,
	0xd3, 0x8f, 0x98, 0xff, 0x02, 0xd7, 0x14, 0x49, 0x2f, 0xe9, 0x86, 0xb2, 0x4f, 0x74, 0xee, 0xf4,
	0x85, 0x4e, 0x9b, 0x7e, 0xa1, 0x01, 0x5c, 0x53, 0x90, 0x0b, 0x76, 0x12, 0xf0, 0x98, 0x97, 0x11,
	0x71, 0x06, 0x82, 0xde, 0xc7, 0xb5, 0x8d, 0x1e, 0xc1, 0x30, 0x11, 0x05, 0x57, 0xa0, 0xad, 0xc0,
	0xc6, 0x21, 0x23, 0x83, 0x44, 0x47, 0x0e, 0xab, 0x48, 0x63, 0xa3, 0x27, 0xd0, 0x0d, 0x4a, 0xbe,
	0x76, 0xe0, 0x4e, 0x02, 0x6f, 0x85, 0xf3, 0x8c, 0x5e, 0x31, 0xac, 0x60, 0xf4, 0x02, 0x46, 0xf2,
	0x7f, 0xc5, 0xe3, 0x70, 0x43, 0xb8, 0x33, 0x52, 0xec, 0x91, 0x62, 0x2d, 0x95, 0x0b, 0x43, 0x50,
	0x7f, 0xcb, 0x74, 0x74, 0xda, 0x53, 0xdf, 0x19, 0xab, 0x16, 0x35, 0x0e, 0xf7, 0x03, 0xd8, 0xa6,
	0x1b, 0xe8, 0x19, 0x74, 0xf9, 0x97, 0x8c, 0xe8, 0xb6, 0x1f, 0xea, 0xb6, 0x57, 0xe8, 0x39, 0xe1,
	0x6b, 0x16, 0x61, 0x45, 0x40, 0x8f, 0x01, 0xb2, 0x20, 0x0f, 0x52, 0xc2, 0x49, 0x5e, 0x38, 0x7b,
	0x42, 0x73, 0x8c, 0x5b, 0x1e, 0xf7, 0x7b, 0x07, 0x6c, 0x93, 0xb3, 0x2c, 0x58, 0x8c, 0x79, 0x1b,
	0x47, 0x24, 0x57, 0xca, 0x43, 0x5c, 0xdb, 0xe8, 0x39, 0xf4, 0xb8, 0xd8, 0x0f, 0x2a, 0x34, 0xf6,
	0x44, 0x0d, 0x47, 0x3b, 0x15, 0x7b, 0xef, 0x3f, 0x2e, 0x71, 0xc5, 0x71, 0xdf, 0x80, 0x25, 0x2c,
	0xa9, 0x17, 0x32, 0xca, 0x09, 0x15, 0x83, 0xd5, 0x7a, 0xc6, 0x6e, 0xd7, 0x3a, 0x53, 0x9a, 0xbd,
	0xa6, 0xd6, 0x99, 0xfb, 0x14, 0x6c, 0x3d, 0xca, 0x99, 0x54, 0xd1, 0xc0, 0x6b, 0xa3, 0x62, 0x6c,
	0xf7, 0x6b, 0xa7, 0x26, 0xce, 0x5b, 0xc4, 0xa9, 0x22, 0xf6, 0x6a, 0xe2, 0x14, 0xcd, 0x6b, 0xcc,
	0xd7, 0x15, 0xb8, 0x3b, 0x4b, 0x63, 0x3e, 0xfc, 0x3a, 0xce, 0x6f, 0x25, 0xe2, 0xef, 0xe8, 0x8f,
	0x1b, 0xfd, 0x93, 0x9f, 0x96, 0x9c, 0x4e, 0x91, 0x31, 0x5a, 0x90, 0x7f, 0xb8, 0x0c, 0xf7, 0x56,
	0x7a, 0x72, 0xf0, 0xb5, 0x8d, 0x1e, 0xc0, 0x20, 0xc8, 0xe2, 0x55, 0x99, 0x27, 0xea, 0x6c, 0x86,
	0xb8, 0x2f, 0xcc, 0x8b, 0x3c, 0x41, 0xde, 0xad, 0x43, 0x90, 0x35, 0x21, 0xcf, 0xe4, 0xf0, 0xbb,
	0x4b, 0xb8, 0xb3, 0x8c, 0x83, 0x3f, 0x2e, 0x63, 0xae, 0xc5, 0x0a, 0x27, 0x12, 0x67, 0x39, 0xc6,
	0x8d, 0xc3, 0xfd, 0xf6, 0xb7, 0x8d, 0x5f, 0xec, 0x34, 0xfe, 0xe1, 0x6e, 0x92, 0xff, 0xd1, 0x79,
	0x37, 0xaa, 0x79, 0x8b, 0xd6, 0x52, 0x2d, 0x24, 0x51, 0x2e, 0x7b, 0xe3, 0x68, 0xa3, 0xa6, 0xcb,
	0x8d, 0xa3, 0x8d, 0xce, 0x54, 0xa3, 0x5b, 0xb1, 0xb3, 0x93, 0x08, 0xa0, 0xe9, 0x13, 0xba, 0x0f,
	0xbd, 0x82, 0x07, 0x39, 0xd7, 0x6f, 0x54, 0x86, 0x98, 0xc7, 0x21, 0xf9, 0x9c, 0x89, 0x99, 0x8a,
	0x0e, 0xa7, 0x62, 0xb1, 0x82, 0x34, 0x5b, 0xa5, 0xd5, 0xd1, 0x75, 0xf1, 0x41, 0x05, 0x2d, 0x0d,
	0x72, 0x5e, 0xa0, 0x7b, 0x60, 0x11, 0x1a, 0xe9, 0xb7, 0xe4, 0xe7, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xdf, 0x56, 0x2e, 0x77, 0xa7, 0x05, 0x00, 0x00,
}