// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session.proto

package session

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SuccessMessage struct {
	Successfull          bool     `protobuf:"varint,1,opt,name=Successfull,proto3" json:"Successfull,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuccessMessage) Reset()         { *m = SuccessMessage{} }
func (m *SuccessMessage) String() string { return proto.CompactTextString(m) }
func (*SuccessMessage) ProtoMessage()    {}
func (*SuccessMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{0}
}

func (m *SuccessMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessMessage.Unmarshal(m, b)
}
func (m *SuccessMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessMessage.Marshal(b, m, deterministic)
}
func (m *SuccessMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessMessage.Merge(m, src)
}
func (m *SuccessMessage) XXX_Size() int {
	return xxx_messageInfo_SuccessMessage.Size(m)
}
func (m *SuccessMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessMessage proto.InternalMessageInfo

func (m *SuccessMessage) GetSuccessfull() bool {
	if m != nil {
		return m.Successfull
	}
	return false
}

type CreateSessionMessage struct {
	Ttl                  int64    `protobuf:"varint,1,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSessionMessage) Reset()         { *m = CreateSessionMessage{} }
func (m *CreateSessionMessage) String() string { return proto.CompactTextString(m) }
func (*CreateSessionMessage) ProtoMessage()    {}
func (*CreateSessionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{1}
}

func (m *CreateSessionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSessionMessage.Unmarshal(m, b)
}
func (m *CreateSessionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSessionMessage.Marshal(b, m, deterministic)
}
func (m *CreateSessionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSessionMessage.Merge(m, src)
}
func (m *CreateSessionMessage) XXX_Size() int {
	return xxx_messageInfo_CreateSessionMessage.Size(m)
}
func (m *CreateSessionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSessionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSessionMessage proto.InternalMessageInfo

func (m *CreateSessionMessage) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type GetSessionMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionMessage) Reset()         { *m = GetSessionMessage{} }
func (m *GetSessionMessage) String() string { return proto.CompactTextString(m) }
func (*GetSessionMessage) ProtoMessage()    {}
func (*GetSessionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{2}
}

func (m *GetSessionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionMessage.Unmarshal(m, b)
}
func (m *GetSessionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionMessage.Marshal(b, m, deterministic)
}
func (m *GetSessionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionMessage.Merge(m, src)
}
func (m *GetSessionMessage) XXX_Size() int {
	return xxx_messageInfo_GetSessionMessage.Size(m)
}
func (m *GetSessionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionMessage proto.InternalMessageInfo

func (m *GetSessionMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AddValueToSessionMessage struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string         `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                *_struct.Value `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddValueToSessionMessage) Reset()         { *m = AddValueToSessionMessage{} }
func (m *AddValueToSessionMessage) String() string { return proto.CompactTextString(m) }
func (*AddValueToSessionMessage) ProtoMessage()    {}
func (*AddValueToSessionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{3}
}

func (m *AddValueToSessionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddValueToSessionMessage.Unmarshal(m, b)
}
func (m *AddValueToSessionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddValueToSessionMessage.Marshal(b, m, deterministic)
}
func (m *AddValueToSessionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddValueToSessionMessage.Merge(m, src)
}
func (m *AddValueToSessionMessage) XXX_Size() int {
	return xxx_messageInfo_AddValueToSessionMessage.Size(m)
}
func (m *AddValueToSessionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddValueToSessionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddValueToSessionMessage proto.InternalMessageInfo

func (m *AddValueToSessionMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddValueToSessionMessage) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AddValueToSessionMessage) GetValue() *_struct.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type AddValuesToSessionMessage struct {
	Id                   string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Values               map[string]*_struct.Value `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AddValuesToSessionMessage) Reset()         { *m = AddValuesToSessionMessage{} }
func (m *AddValuesToSessionMessage) String() string { return proto.CompactTextString(m) }
func (*AddValuesToSessionMessage) ProtoMessage()    {}
func (*AddValuesToSessionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{4}
}

func (m *AddValuesToSessionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddValuesToSessionMessage.Unmarshal(m, b)
}
func (m *AddValuesToSessionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddValuesToSessionMessage.Marshal(b, m, deterministic)
}
func (m *AddValuesToSessionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddValuesToSessionMessage.Merge(m, src)
}
func (m *AddValuesToSessionMessage) XXX_Size() int {
	return xxx_messageInfo_AddValuesToSessionMessage.Size(m)
}
func (m *AddValuesToSessionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddValuesToSessionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddValuesToSessionMessage proto.InternalMessageInfo

func (m *AddValuesToSessionMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddValuesToSessionMessage) GetValues() map[string]*_struct.Value {
	if m != nil {
		return m.Values
	}
	return nil
}

type SessionResponse struct {
	Id                   string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Values               map[string]*_struct.Value `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *SessionResponse) Reset()         { *m = SessionResponse{} }
func (m *SessionResponse) String() string { return proto.CompactTextString(m) }
func (*SessionResponse) ProtoMessage()    {}
func (*SessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{5}
}

func (m *SessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionResponse.Unmarshal(m, b)
}
func (m *SessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionResponse.Marshal(b, m, deterministic)
}
func (m *SessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionResponse.Merge(m, src)
}
func (m *SessionResponse) XXX_Size() int {
	return xxx_messageInfo_SessionResponse.Size(m)
}
func (m *SessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SessionResponse proto.InternalMessageInfo

func (m *SessionResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SessionResponse) GetValues() map[string]*_struct.Value {
	if m != nil {
		return m.Values
	}
	return nil
}

type InvalidateSessionMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvalidateSessionMessage) Reset()         { *m = InvalidateSessionMessage{} }
func (m *InvalidateSessionMessage) String() string { return proto.CompactTextString(m) }
func (*InvalidateSessionMessage) ProtoMessage()    {}
func (*InvalidateSessionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{6}
}

func (m *InvalidateSessionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvalidateSessionMessage.Unmarshal(m, b)
}
func (m *InvalidateSessionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvalidateSessionMessage.Marshal(b, m, deterministic)
}
func (m *InvalidateSessionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvalidateSessionMessage.Merge(m, src)
}
func (m *InvalidateSessionMessage) XXX_Size() int {
	return xxx_messageInfo_InvalidateSessionMessage.Size(m)
}
func (m *InvalidateSessionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InvalidateSessionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InvalidateSessionMessage proto.InternalMessageInfo

func (m *InvalidateSessionMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type InvalidateSessionValueMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvalidateSessionValueMessage) Reset()         { *m = InvalidateSessionValueMessage{} }
func (m *InvalidateSessionValueMessage) String() string { return proto.CompactTextString(m) }
func (*InvalidateSessionValueMessage) ProtoMessage()    {}
func (*InvalidateSessionValueMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{7}
}

func (m *InvalidateSessionValueMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvalidateSessionValueMessage.Unmarshal(m, b)
}
func (m *InvalidateSessionValueMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvalidateSessionValueMessage.Marshal(b, m, deterministic)
}
func (m *InvalidateSessionValueMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvalidateSessionValueMessage.Merge(m, src)
}
func (m *InvalidateSessionValueMessage) XXX_Size() int {
	return xxx_messageInfo_InvalidateSessionValueMessage.Size(m)
}
func (m *InvalidateSessionValueMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InvalidateSessionValueMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InvalidateSessionValueMessage proto.InternalMessageInfo

func (m *InvalidateSessionValueMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InvalidateSessionValueMessage) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type InvalidateSessionValuesMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Keys                 []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvalidateSessionValuesMessage) Reset()         { *m = InvalidateSessionValuesMessage{} }
func (m *InvalidateSessionValuesMessage) String() string { return proto.CompactTextString(m) }
func (*InvalidateSessionValuesMessage) ProtoMessage()    {}
func (*InvalidateSessionValuesMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{8}
}

func (m *InvalidateSessionValuesMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvalidateSessionValuesMessage.Unmarshal(m, b)
}
func (m *InvalidateSessionValuesMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvalidateSessionValuesMessage.Marshal(b, m, deterministic)
}
func (m *InvalidateSessionValuesMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvalidateSessionValuesMessage.Merge(m, src)
}
func (m *InvalidateSessionValuesMessage) XXX_Size() int {
	return xxx_messageInfo_InvalidateSessionValuesMessage.Size(m)
}
func (m *InvalidateSessionValuesMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InvalidateSessionValuesMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InvalidateSessionValuesMessage proto.InternalMessageInfo

func (m *InvalidateSessionValuesMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InvalidateSessionValuesMessage) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterType((*SuccessMessage)(nil), "github.com.hobord.infra2.session.SuccessMessage")
	proto.RegisterType((*CreateSessionMessage)(nil), "github.com.hobord.infra2.session.CreateSessionMessage")
	proto.RegisterType((*GetSessionMessage)(nil), "github.com.hobord.infra2.session.GetSessionMessage")
	proto.RegisterType((*AddValueToSessionMessage)(nil), "github.com.hobord.infra2.session.AddValueToSessionMessage")
	proto.RegisterType((*AddValuesToSessionMessage)(nil), "github.com.hobord.infra2.session.AddValuesToSessionMessage")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "github.com.hobord.infra2.session.AddValuesToSessionMessage.ValuesEntry")
	proto.RegisterType((*SessionResponse)(nil), "github.com.hobord.infra2.session.SessionResponse")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "github.com.hobord.infra2.session.SessionResponse.ValuesEntry")
	proto.RegisterType((*InvalidateSessionMessage)(nil), "github.com.hobord.infra2.session.InvalidateSessionMessage")
	proto.RegisterType((*InvalidateSessionValueMessage)(nil), "github.com.hobord.infra2.session.InvalidateSessionValueMessage")
	proto.RegisterType((*InvalidateSessionValuesMessage)(nil), "github.com.hobord.infra2.session.InvalidateSessionValuesMessage")
}

func init() { proto.RegisterFile("session.proto", fileDescriptor_3a6be1b361fa6f14) }

var fileDescriptor_3a6be1b361fa6f14 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0xed, 0x24, 0xda, 0xba, 0x77, 0xe9, 0xea, 0x0e, 0x52, 0xe3, 0xa2, 0x12, 0xc6, 0x97, 0x45,
	0x64, 0xaa, 0x29, 0x88, 0x54, 0x44, 0xeb, 0x07, 0xc5, 0x07, 0x1f, 0xcc, 0xaa, 0x0f, 0xbe, 0x48,
	0x36, 0x99, 0xdd, 0xc6, 0xc6, 0x4c, 0xc9, 0x4c, 0x02, 0x0b, 0xbe, 0x29, 0xbe, 0x89, 0xff, 0xc8,
	0x3f, 0xe0, 0x8f, 0xf0, 0xaf, 0x48, 0x26, 0x89, 0xb6, 0xf9, 0x68, 0x62, 0x0a, 0xbe, 0x0d, 0x37,
	0xf7, 0x9e, 0x7b, 0xce, 0xe4, 0x9c, 0x81, 0x4d, 0xc1, 0x84, 0xf0, 0x79, 0x48, 0x8f, 0x22, 0x2e,
	0x39, 0x36, 0x97, 0xbe, 0x3c, 0x88, 0xe7, 0xd4, 0xe5, 0x1f, 0xe9, 0x01, 0x9f, 0xf3, 0xc8, 0xa3,
	0x7e, 0xb8, 0x88, 0x1c, 0x8b, 0xe6, 0x7d, 0x93, 0x6b, 0x4b, 0xce, 0x97, 0x01, 0xdb, 0x56, 0xfd,
	0xf3, 0x78, 0xb1, 0x2d, 0x64, 0x14, 0xbb, 0x32, 0x9b, 0x27, 0x16, 0x8c, 0x66, 0xb1, 0xeb, 0x32,
	0x21, 0x5e, 0x32, 0x21, 0x9c, 0x25, 0xc3, 0x26, 0x0c, 0xf3, 0xca, 0x22, 0x0e, 0x02, 0x03, 0x99,
	0x68, 0x7a, 0xc1, 0x3e, 0x5e, 0x22, 0x53, 0xb8, 0xfc, 0x34, 0x62, 0x8e, 0x64, 0xb3, 0x6c, 0x45,
	0x31, 0x79, 0x09, 0x74, 0x29, 0xb3, 0x09, 0xdd, 0x4e, 0x8f, 0xe4, 0x26, 0x8c, 0xf7, 0x99, 0x2c,
	0xb5, 0x8d, 0x40, 0xf3, 0x3d, 0xd5, 0x35, 0xb0, 0x35, 0xdf, 0x23, 0x1f, 0xc0, 0xd8, 0xf3, 0xbc,
	0xb7, 0x4e, 0x10, 0xb3, 0xd7, 0xfc, 0xf4, 0xde, 0x74, 0xc5, 0x21, 0x5b, 0x19, 0x9a, 0x2a, 0xa4,
	0x47, 0x7c, 0x1b, 0xce, 0x27, 0xe9, 0xa8, 0xa1, 0x9b, 0x68, 0x3a, 0xb4, 0xb6, 0x68, 0x26, 0x97,
	0x16, 0x72, 0xa9, 0x02, 0xb6, 0xb3, 0x26, 0xf2, 0x0b, 0xc1, 0xd5, 0x62, 0x99, 0x68, 0xdd, 0xf6,
	0x1e, 0xd6, 0xd5, 0x98, 0x30, 0x74, 0x53, 0x9f, 0x0e, 0xad, 0x7d, 0xda, 0x76, 0xdb, 0xb4, 0x11,
	0x3c, 0xe3, 0x21, 0x9e, 0x87, 0x32, 0x5a, 0xd9, 0x39, 0xec, 0xe4, 0x15, 0x0c, 0x8f, 0x95, 0x0b,
	0x75, 0xa8, 0x46, 0x9d, 0xd6, 0x41, 0xdd, 0xae, 0x76, 0x1f, 0x91, 0x9f, 0x08, 0x2e, 0xe6, 0x9b,
	0x6d, 0x26, 0x8e, 0x78, 0x28, 0xaa, 0xba, 0xde, 0xfc, 0xd1, 0xa5, 0x29, 0x5d, 0x0f, 0xdb, 0x75,
	0x95, 0x20, 0xff, 0x97, 0x9a, 0x5b, 0x60, 0xbc, 0x08, 0x13, 0x27, 0xf0, 0xbd, 0xaa, 0xdd, 0xca,
	0x3e, 0xda, 0x83, 0xeb, 0x95, 0x5e, 0x05, 0xd8, 0xd9, 0x4c, 0xe4, 0x19, 0xdc, 0xa8, 0x87, 0x10,
	0x4d, 0x18, 0x18, 0xce, 0x1d, 0xb2, 0x55, 0x76, 0x91, 0x03, 0x5b, 0x9d, 0xad, 0x1f, 0x1b, 0x30,
	0xca, 0x87, 0x67, 0x2c, 0x4a, 0x7c, 0x97, 0xe1, 0x04, 0xe0, 0x6f, 0x10, 0xf0, 0x4e, 0xfb, 0x7d,
	0x57, 0x62, 0x33, 0xb9, 0xfb, 0xcf, 0x3f, 0x89, 0xac, 0xe1, 0x4f, 0xb0, 0x79, 0x22, 0xaa, 0xf8,
	0x5e, 0x3b, 0x4a, 0x5d, 0xb6, 0xfb, 0x6d, 0xff, 0x82, 0x60, 0x5c, 0x89, 0x36, 0xde, 0xed, 0x9e,
	0xa2, 0x72, 0x88, 0xfa, 0xd1, 0xf8, 0x8a, 0x00, 0x57, 0x73, 0x89, 0x1f, 0x9c, 0x21, 0xcd, 0xfd,
	0x88, 0x7c, 0x43, 0xb0, 0x55, 0xef, 0x2f, 0xfc, 0xa8, 0x1d, 0xef, 0x54, 0x73, 0x4f, 0xee, 0x74,
	0x20, 0x74, 0xe2, 0xa1, 0x27, 0x6b, 0xf8, 0x3b, 0x82, 0x2b, 0x0d, 0x7e, 0xc7, 0x8f, 0xfb, 0x12,
	0x12, 0x67, 0x61, 0xf4, 0x19, 0xc1, 0xb8, 0x02, 0xdb, 0xc5, 0x31, 0x4d, 0xaf, 0x44, 0x1f, 0x16,
	0x4f, 0x06, 0xef, 0x36, 0xf2, 0x6f, 0xf3, 0x75, 0xf5, 0x36, 0xed, 0xfc, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x80, 0x2d, 0xbd, 0xc6, 0x77, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionServiceClient interface {
	GetSession(ctx context.Context, in *GetSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error)
	CreateSession(ctx context.Context, in *CreateSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error)
	AddValueToSession(ctx context.Context, in *AddValueToSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error)
	AddValuesToSession(ctx context.Context, in *AddValuesToSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error)
	InvalidateSessionValue(ctx context.Context, in *InvalidateSessionValueMessage, opts ...grpc.CallOption) (*SuccessMessage, error)
	InvalidateSessionValues(ctx context.Context, in *InvalidateSessionValuesMessage, opts ...grpc.CallOption) (*SuccessMessage, error)
	InvalidateSession(ctx context.Context, in *InvalidateSessionMessage, opts ...grpc.CallOption) (*SuccessMessage, error)
}

type sessionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSessionServiceClient(cc *grpc.ClientConn) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) GetSession(ctx context.Context, in *GetSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error) {
	out := new(SessionResponse)
	err := c.cc.Invoke(ctx, "/github.com.hobord.infra2.session.SessionService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) CreateSession(ctx context.Context, in *CreateSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error) {
	out := new(SessionResponse)
	err := c.cc.Invoke(ctx, "/github.com.hobord.infra2.session.SessionService/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) AddValueToSession(ctx context.Context, in *AddValueToSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error) {
	out := new(SessionResponse)
	err := c.cc.Invoke(ctx, "/github.com.hobord.infra2.session.SessionService/AddValueToSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) AddValuesToSession(ctx context.Context, in *AddValuesToSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error) {
	out := new(SessionResponse)
	err := c.cc.Invoke(ctx, "/github.com.hobord.infra2.session.SessionService/AddValuesToSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) InvalidateSessionValue(ctx context.Context, in *InvalidateSessionValueMessage, opts ...grpc.CallOption) (*SuccessMessage, error) {
	out := new(SuccessMessage)
	err := c.cc.Invoke(ctx, "/github.com.hobord.infra2.session.SessionService/InvalidateSessionValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) InvalidateSessionValues(ctx context.Context, in *InvalidateSessionValuesMessage, opts ...grpc.CallOption) (*SuccessMessage, error) {
	out := new(SuccessMessage)
	err := c.cc.Invoke(ctx, "/github.com.hobord.infra2.session.SessionService/InvalidateSessionValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) InvalidateSession(ctx context.Context, in *InvalidateSessionMessage, opts ...grpc.CallOption) (*SuccessMessage, error) {
	out := new(SuccessMessage)
	err := c.cc.Invoke(ctx, "/github.com.hobord.infra2.session.SessionService/InvalidateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
type SessionServiceServer interface {
	GetSession(context.Context, *GetSessionMessage) (*SessionResponse, error)
	CreateSession(context.Context, *CreateSessionMessage) (*SessionResponse, error)
	AddValueToSession(context.Context, *AddValueToSessionMessage) (*SessionResponse, error)
	AddValuesToSession(context.Context, *AddValuesToSessionMessage) (*SessionResponse, error)
	InvalidateSessionValue(context.Context, *InvalidateSessionValueMessage) (*SuccessMessage, error)
	InvalidateSessionValues(context.Context, *InvalidateSessionValuesMessage) (*SuccessMessage, error)
	InvalidateSession(context.Context, *InvalidateSessionMessage) (*SuccessMessage, error)
}

// UnimplementedSessionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSessionServiceServer struct {
}

func (*UnimplementedSessionServiceServer) GetSession(ctx context.Context, req *GetSessionMessage) (*SessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (*UnimplementedSessionServiceServer) CreateSession(ctx context.Context, req *CreateSessionMessage) (*SessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (*UnimplementedSessionServiceServer) AddValueToSession(ctx context.Context, req *AddValueToSessionMessage) (*SessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddValueToSession not implemented")
}
func (*UnimplementedSessionServiceServer) AddValuesToSession(ctx context.Context, req *AddValuesToSessionMessage) (*SessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddValuesToSession not implemented")
}
func (*UnimplementedSessionServiceServer) InvalidateSessionValue(ctx context.Context, req *InvalidateSessionValueMessage) (*SuccessMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateSessionValue not implemented")
}
func (*UnimplementedSessionServiceServer) InvalidateSessionValues(ctx context.Context, req *InvalidateSessionValuesMessage) (*SuccessMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateSessionValues not implemented")
}
func (*UnimplementedSessionServiceServer) InvalidateSession(ctx context.Context, req *InvalidateSessionMessage) (*SuccessMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateSession not implemented")
}

func RegisterSessionServiceServer(s *grpc.Server, srv SessionServiceServer) {
	s.RegisterService(&_SessionService_serviceDesc, srv)
}

func _SessionService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hobord.infra2.session.SessionService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetSession(ctx, req.(*GetSessionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hobord.infra2.session.SessionService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).CreateSession(ctx, req.(*CreateSessionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_AddValueToSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddValueToSessionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).AddValueToSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hobord.infra2.session.SessionService/AddValueToSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).AddValueToSession(ctx, req.(*AddValueToSessionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_AddValuesToSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddValuesToSessionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).AddValuesToSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hobord.infra2.session.SessionService/AddValuesToSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).AddValuesToSession(ctx, req.(*AddValuesToSessionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_InvalidateSessionValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateSessionValueMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).InvalidateSessionValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hobord.infra2.session.SessionService/InvalidateSessionValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).InvalidateSessionValue(ctx, req.(*InvalidateSessionValueMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_InvalidateSessionValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateSessionValuesMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).InvalidateSessionValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hobord.infra2.session.SessionService/InvalidateSessionValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).InvalidateSessionValues(ctx, req.(*InvalidateSessionValuesMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_InvalidateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateSessionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).InvalidateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hobord.infra2.session.SessionService/InvalidateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).InvalidateSession(ctx, req.(*InvalidateSessionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.hobord.infra2.session.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSession",
			Handler:    _SessionService_GetSession_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _SessionService_CreateSession_Handler,
		},
		{
			MethodName: "AddValueToSession",
			Handler:    _SessionService_AddValueToSession_Handler,
		},
		{
			MethodName: "AddValuesToSession",
			Handler:    _SessionService_AddValuesToSession_Handler,
		},
		{
			MethodName: "InvalidateSessionValue",
			Handler:    _SessionService_InvalidateSessionValue_Handler,
		},
		{
			MethodName: "InvalidateSessionValues",
			Handler:    _SessionService_InvalidateSessionValues_Handler,
		},
		{
			MethodName: "InvalidateSession",
			Handler:    _SessionService_InvalidateSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}