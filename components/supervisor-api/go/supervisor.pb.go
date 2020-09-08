// Code generated by protoc-gen-go. DO NOT EDIT.
// source: supervisor.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ContentSource int32

const (
	ContentSource_from_other    ContentSource = 0
	ContentSource_from_backup   ContentSource = 1
	ContentSource_from_prebuild ContentSource = 2
)

var ContentSource_name = map[int32]string{
	0: "from_other",
	1: "from_backup",
	2: "from_prebuild",
}

var ContentSource_value = map[string]int32{
	"from_other":    0,
	"from_backup":   1,
	"from_prebuild": 2,
}

func (x ContentSource) String() string {
	return proto.EnumName(ContentSource_name, int32(x))
}

func (ContentSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{0}
}

type PrepareBackupRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareBackupRequest) Reset()         { *m = PrepareBackupRequest{} }
func (m *PrepareBackupRequest) String() string { return proto.CompactTextString(m) }
func (*PrepareBackupRequest) ProtoMessage()    {}
func (*PrepareBackupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{0}
}

func (m *PrepareBackupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareBackupRequest.Unmarshal(m, b)
}
func (m *PrepareBackupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareBackupRequest.Marshal(b, m, deterministic)
}
func (m *PrepareBackupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareBackupRequest.Merge(m, src)
}
func (m *PrepareBackupRequest) XXX_Size() int {
	return xxx_messageInfo_PrepareBackupRequest.Size(m)
}
func (m *PrepareBackupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareBackupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareBackupRequest proto.InternalMessageInfo

type PrepareBackupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareBackupResponse) Reset()         { *m = PrepareBackupResponse{} }
func (m *PrepareBackupResponse) String() string { return proto.CompactTextString(m) }
func (*PrepareBackupResponse) ProtoMessage()    {}
func (*PrepareBackupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{1}
}

func (m *PrepareBackupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareBackupResponse.Unmarshal(m, b)
}
func (m *PrepareBackupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareBackupResponse.Marshal(b, m, deterministic)
}
func (m *PrepareBackupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareBackupResponse.Merge(m, src)
}
func (m *PrepareBackupResponse) XXX_Size() int {
	return xxx_messageInfo_PrepareBackupResponse.Size(m)
}
func (m *PrepareBackupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareBackupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareBackupResponse proto.InternalMessageInfo

type SupervisorStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupervisorStatusRequest) Reset()         { *m = SupervisorStatusRequest{} }
func (m *SupervisorStatusRequest) String() string { return proto.CompactTextString(m) }
func (*SupervisorStatusRequest) ProtoMessage()    {}
func (*SupervisorStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{2}
}

func (m *SupervisorStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupervisorStatusRequest.Unmarshal(m, b)
}
func (m *SupervisorStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupervisorStatusRequest.Marshal(b, m, deterministic)
}
func (m *SupervisorStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupervisorStatusRequest.Merge(m, src)
}
func (m *SupervisorStatusRequest) XXX_Size() int {
	return xxx_messageInfo_SupervisorStatusRequest.Size(m)
}
func (m *SupervisorStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SupervisorStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SupervisorStatusRequest proto.InternalMessageInfo

type SupervisorStatusResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupervisorStatusResponse) Reset()         { *m = SupervisorStatusResponse{} }
func (m *SupervisorStatusResponse) String() string { return proto.CompactTextString(m) }
func (*SupervisorStatusResponse) ProtoMessage()    {}
func (*SupervisorStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{3}
}

func (m *SupervisorStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupervisorStatusResponse.Unmarshal(m, b)
}
func (m *SupervisorStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupervisorStatusResponse.Marshal(b, m, deterministic)
}
func (m *SupervisorStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupervisorStatusResponse.Merge(m, src)
}
func (m *SupervisorStatusResponse) XXX_Size() int {
	return xxx_messageInfo_SupervisorStatusResponse.Size(m)
}
func (m *SupervisorStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SupervisorStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SupervisorStatusResponse proto.InternalMessageInfo

func (m *SupervisorStatusResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type IDEStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDEStatusRequest) Reset()         { *m = IDEStatusRequest{} }
func (m *IDEStatusRequest) String() string { return proto.CompactTextString(m) }
func (*IDEStatusRequest) ProtoMessage()    {}
func (*IDEStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{4}
}

func (m *IDEStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDEStatusRequest.Unmarshal(m, b)
}
func (m *IDEStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDEStatusRequest.Marshal(b, m, deterministic)
}
func (m *IDEStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDEStatusRequest.Merge(m, src)
}
func (m *IDEStatusRequest) XXX_Size() int {
	return xxx_messageInfo_IDEStatusRequest.Size(m)
}
func (m *IDEStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IDEStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IDEStatusRequest proto.InternalMessageInfo

type IDEStatusResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDEStatusResponse) Reset()         { *m = IDEStatusResponse{} }
func (m *IDEStatusResponse) String() string { return proto.CompactTextString(m) }
func (*IDEStatusResponse) ProtoMessage()    {}
func (*IDEStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{5}
}

func (m *IDEStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDEStatusResponse.Unmarshal(m, b)
}
func (m *IDEStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDEStatusResponse.Marshal(b, m, deterministic)
}
func (m *IDEStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDEStatusResponse.Merge(m, src)
}
func (m *IDEStatusResponse) XXX_Size() int {
	return xxx_messageInfo_IDEStatusResponse.Size(m)
}
func (m *IDEStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IDEStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IDEStatusResponse proto.InternalMessageInfo

func (m *IDEStatusResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type BackupStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackupStatusRequest) Reset()         { *m = BackupStatusRequest{} }
func (m *BackupStatusRequest) String() string { return proto.CompactTextString(m) }
func (*BackupStatusRequest) ProtoMessage()    {}
func (*BackupStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{6}
}

func (m *BackupStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackupStatusRequest.Unmarshal(m, b)
}
func (m *BackupStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackupStatusRequest.Marshal(b, m, deterministic)
}
func (m *BackupStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupStatusRequest.Merge(m, src)
}
func (m *BackupStatusRequest) XXX_Size() int {
	return xxx_messageInfo_BackupStatusRequest.Size(m)
}
func (m *BackupStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BackupStatusRequest proto.InternalMessageInfo

type BackupStatusResponse struct {
	CanaryAvailable      bool     `protobuf:"varint,1,opt,name=canary_available,json=canaryAvailable,proto3" json:"canary_available,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackupStatusResponse) Reset()         { *m = BackupStatusResponse{} }
func (m *BackupStatusResponse) String() string { return proto.CompactTextString(m) }
func (*BackupStatusResponse) ProtoMessage()    {}
func (*BackupStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{7}
}

func (m *BackupStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackupStatusResponse.Unmarshal(m, b)
}
func (m *BackupStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackupStatusResponse.Marshal(b, m, deterministic)
}
func (m *BackupStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupStatusResponse.Merge(m, src)
}
func (m *BackupStatusResponse) XXX_Size() int {
	return xxx_messageInfo_BackupStatusResponse.Size(m)
}
func (m *BackupStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BackupStatusResponse proto.InternalMessageInfo

func (m *BackupStatusResponse) GetCanaryAvailable() bool {
	if m != nil {
		return m.CanaryAvailable
	}
	return false
}

type ContentStatusRequest struct {
	// if true this request will return either when it times out or when the workspace content
	// has become available.
	Wait                 bool     `protobuf:"varint,1,opt,name=wait,proto3" json:"wait,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentStatusRequest) Reset()         { *m = ContentStatusRequest{} }
func (m *ContentStatusRequest) String() string { return proto.CompactTextString(m) }
func (*ContentStatusRequest) ProtoMessage()    {}
func (*ContentStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{8}
}

func (m *ContentStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentStatusRequest.Unmarshal(m, b)
}
func (m *ContentStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentStatusRequest.Marshal(b, m, deterministic)
}
func (m *ContentStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentStatusRequest.Merge(m, src)
}
func (m *ContentStatusRequest) XXX_Size() int {
	return xxx_messageInfo_ContentStatusRequest.Size(m)
}
func (m *ContentStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContentStatusRequest proto.InternalMessageInfo

func (m *ContentStatusRequest) GetWait() bool {
	if m != nil {
		return m.Wait
	}
	return false
}

type ContentStatusResponse struct {
	// true if the workspace content is available
	Available bool `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	// source indicates where the workspace content came from
	Source               ContentSource `protobuf:"varint,2,opt,name=source,proto3,enum=supervisor.ContentSource" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ContentStatusResponse) Reset()         { *m = ContentStatusResponse{} }
func (m *ContentStatusResponse) String() string { return proto.CompactTextString(m) }
func (*ContentStatusResponse) ProtoMessage()    {}
func (*ContentStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{9}
}

func (m *ContentStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentStatusResponse.Unmarshal(m, b)
}
func (m *ContentStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentStatusResponse.Marshal(b, m, deterministic)
}
func (m *ContentStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentStatusResponse.Merge(m, src)
}
func (m *ContentStatusResponse) XXX_Size() int {
	return xxx_messageInfo_ContentStatusResponse.Size(m)
}
func (m *ContentStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContentStatusResponse proto.InternalMessageInfo

func (m *ContentStatusResponse) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *ContentStatusResponse) GetSource() ContentSource {
	if m != nil {
		return m.Source
	}
	return ContentSource_from_other
}

func init() {
	proto.RegisterEnum("supervisor.ContentSource", ContentSource_name, ContentSource_value)
	proto.RegisterType((*PrepareBackupRequest)(nil), "supervisor.PrepareBackupRequest")
	proto.RegisterType((*PrepareBackupResponse)(nil), "supervisor.PrepareBackupResponse")
	proto.RegisterType((*SupervisorStatusRequest)(nil), "supervisor.SupervisorStatusRequest")
	proto.RegisterType((*SupervisorStatusResponse)(nil), "supervisor.SupervisorStatusResponse")
	proto.RegisterType((*IDEStatusRequest)(nil), "supervisor.IDEStatusRequest")
	proto.RegisterType((*IDEStatusResponse)(nil), "supervisor.IDEStatusResponse")
	proto.RegisterType((*BackupStatusRequest)(nil), "supervisor.BackupStatusRequest")
	proto.RegisterType((*BackupStatusResponse)(nil), "supervisor.BackupStatusResponse")
	proto.RegisterType((*ContentStatusRequest)(nil), "supervisor.ContentStatusRequest")
	proto.RegisterType((*ContentStatusResponse)(nil), "supervisor.ContentStatusResponse")
}

func init() {
	proto.RegisterFile("supervisor.proto", fileDescriptor_b8b9452d77b1c7d2)
}

var fileDescriptor_b8b9452d77b1c7d2 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x41, 0x73, 0xd2, 0x40,
	0x14, 0xc7, 0x25, 0xd5, 0x6a, 0x9f, 0x42, 0xc3, 0x2b, 0x14, 0x9a, 0xa1, 0x63, 0x1a, 0x74, 0xa6,
	0x72, 0x68, 0xa6, 0x78, 0x74, 0x3c, 0xd0, 0xea, 0xc1, 0x9b, 0x43, 0x6f, 0xbd, 0xd4, 0x25, 0x5d,
	0xdb, 0x48, 0xcc, 0xc6, 0xdd, 0x0d, 0x8e, 0x53, 0xbd, 0xf8, 0x15, 0x3c, 0xf8, 0xc1, 0xfc, 0x0a,
	0x7e, 0x07, 0xaf, 0x0e, 0xbb, 0x9b, 0xb2, 0x49, 0x81, 0x0b, 0x03, 0xff, 0xf7, 0x7f, 0xef, 0xf7,
	0xd8, 0xfd, 0x27, 0xe0, 0x8a, 0x3c, 0xa3, 0x7c, 0x16, 0x0b, 0xc6, 0x8f, 0x32, 0xce, 0x24, 0x43,
	0x58, 0x28, 0x5e, 0xef, 0x8a, 0xb1, 0xab, 0x84, 0x86, 0x24, 0x8b, 0x43, 0x92, 0xa6, 0x4c, 0x12,
	0x19, 0xb3, 0x54, 0x68, 0x67, 0xb0, 0x0b, 0xad, 0xf7, 0x9c, 0x66, 0x84, 0xd3, 0x13, 0x12, 0x4d,
	0xf3, 0x6c, 0x4c, 0xbf, 0xe4, 0x54, 0xc8, 0xa0, 0x03, 0xed, 0x8a, 0x2e, 0x32, 0x96, 0x0a, 0x1a,
	0xec, 0x41, 0xe7, 0xec, 0x76, 0xf8, 0x99, 0x24, 0x32, 0x17, 0x45, 0xcf, 0x00, 0xba, 0x77, 0x4b,
	0xba, 0x0d, 0x1b, 0xe0, 0xb0, 0x69, 0xb7, 0xe6, 0xd7, 0x0e, 0x1f, 0x8d, 0x1d, 0x36, 0x0d, 0x10,
	0xdc, 0x77, 0x6f, 0xde, 0x96, 0xfb, 0xfb, 0xd0, 0xb4, 0xb4, 0x15, 0x8d, 0x6d, 0xd8, 0xd1, 0x1b,
	0x95, 0x7b, 0x47, 0xd0, 0x2a, 0xcb, 0xa6, 0xfd, 0x05, 0xb8, 0x11, 0x49, 0x09, 0xff, 0x76, 0x41,
	0x66, 0x24, 0x4e, 0xc8, 0x24, 0xa1, 0x66, 0xd8, 0xb6, 0xd6, 0x47, 0x85, 0x1c, 0x0c, 0xa0, 0x75,
	0xca, 0x52, 0x49, 0x53, 0x59, 0x1a, 0x8d, 0x08, 0xf7, 0xbf, 0x92, 0x58, 0x9a, 0x36, 0xf5, 0x3d,
	0xb8, 0x86, 0x76, 0xc5, 0x6b, 0x78, 0x3d, 0xd8, 0xaa, 0x82, 0x16, 0x02, 0x1e, 0xc3, 0xa6, 0x60,
	0x39, 0x8f, 0x68, 0xd7, 0xf1, 0x6b, 0x87, 0x8d, 0xe1, 0xde, 0x91, 0x75, 0x75, 0xc5, 0x40, 0x65,
	0x18, 0x1b, 0xe3, 0xe0, 0x14, 0xea, 0xa5, 0x02, 0x36, 0x00, 0x3e, 0x72, 0xf6, 0xf9, 0x82, 0xc9,
	0x6b, 0xca, 0xdd, 0x7b, 0xb8, 0x0d, 0x8f, 0xd5, 0xef, 0x89, 0xfa, 0xfb, 0x6e, 0x0d, 0x9b, 0x50,
	0x57, 0x42, 0xc6, 0xe9, 0x24, 0x8f, 0x93, 0x4b, 0xd7, 0x19, 0xde, 0x40, 0xdd, 0x9c, 0xce, 0x1c,
	0x16, 0x51, 0xfc, 0x04, 0x0f, 0xcd, 0xf5, 0xa2, 0x6f, 0xef, 0xb0, 0x2c, 0x0b, 0xde, 0xc1, 0x1a,
	0x87, 0x49, 0x85, 0xf7, 0xf3, 0xcf, 0xdf, 0x5f, 0x4e, 0x0b, 0x31, 0x9c, 0x1d, 0x87, 0x7a, 0x93,
	0x30, 0xd3, 0xce, 0xe1, 0xbf, 0x0d, 0xa8, 0xeb, 0x53, 0x2a, 0xe8, 0xdf, 0xc1, 0xad, 0x06, 0x05,
	0xfb, 0x36, 0x64, 0x45, 0xc2, 0xbc, 0x67, 0xeb, 0x4d, 0x66, 0x99, 0x7d, 0xb5, 0x4c, 0x07, 0xdb,
	0xf3, 0x65, 0x84, 0xaa, 0x85, 0x8b, 0x3e, 0xfc, 0x00, 0x5b, 0xb7, 0x31, 0xc3, 0x9e, 0x3d, 0xb1,
	0x9a, 0x48, 0x6f, 0x7f, 0x45, 0xd5, 0x80, 0x76, 0x15, 0xc8, 0xc5, 0x86, 0x05, 0x8a, 0x2f, 0x29,
	0x26, 0xf0, 0xc4, 0x0e, 0x23, 0x3e, 0xb5, 0xc7, 0x2c, 0x49, 0xaf, 0xe7, 0xaf, 0x36, 0x14, 0x8f,
	0x9d, 0x42, 0xed, 0x60, 0xd3, 0x42, 0xe9, 0x73, 0xc6, 0xdf, 0xb5, 0x45, 0x44, 0x34, 0xcf, 0x5f,
	0x16, 0xab, 0x12, 0xf0, 0x60, 0x8d, 0xc3, 0x10, 0x47, 0x8a, 0xf8, 0x4a, 0x5f, 0xa9, 0x21, 0x46,
	0xda, 0x79, 0xfe, 0x1c, 0xfb, 0x77, 0xd5, 0x70, 0xfe, 0x64, 0x84, 0x37, 0xf3, 0xcf, 0xd7, 0x92,
	0xe7, 0xf4, 0xc7, 0xc9, 0x83, 0xf3, 0x0d, 0x92, 0xc5, 0x93, 0x4d, 0xf5, 0xaa, 0x79, 0xf9, 0x3f,
	0x00, 0x00, 0xff, 0xff, 0x35, 0xd6, 0x09, 0x46, 0xa8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BackupServiceClient is the client API for BackupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackupServiceClient interface {
	// Prepare prepares a workspace backup and is intended to be called by the PreStop
	// hook of the container. It should hardly ever be neccesary to call this from any
	// other point.
	Prepare(ctx context.Context, in *PrepareBackupRequest, opts ...grpc.CallOption) (*PrepareBackupResponse, error)
}

type backupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupServiceClient(cc grpc.ClientConnInterface) BackupServiceClient {
	return &backupServiceClient{cc}
}

func (c *backupServiceClient) Prepare(ctx context.Context, in *PrepareBackupRequest, opts ...grpc.CallOption) (*PrepareBackupResponse, error) {
	out := new(PrepareBackupResponse)
	err := c.cc.Invoke(ctx, "/supervisor.BackupService/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupServiceServer is the server API for BackupService service.
type BackupServiceServer interface {
	// Prepare prepares a workspace backup and is intended to be called by the PreStop
	// hook of the container. It should hardly ever be neccesary to call this from any
	// other point.
	Prepare(context.Context, *PrepareBackupRequest) (*PrepareBackupResponse, error)
}

// UnimplementedBackupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBackupServiceServer struct {
}

func (*UnimplementedBackupServiceServer) Prepare(ctx context.Context, req *PrepareBackupRequest) (*PrepareBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}

func RegisterBackupServiceServer(s *grpc.Server, srv BackupServiceServer) {
	s.RegisterService(&_BackupService_serviceDesc, srv)
}

func _BackupService_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.BackupService/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).Prepare(ctx, req.(*PrepareBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "supervisor.BackupService",
	HandlerType: (*BackupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prepare",
			Handler:    _BackupService_Prepare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supervisor.proto",
}

// StatusServiceClient is the client API for StatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatusServiceClient interface {
	// SupervisorStatus returns once supervisor is running.
	SupervisorStatus(ctx context.Context, in *SupervisorStatusRequest, opts ...grpc.CallOption) (*SupervisorStatusResponse, error)
	// IDEStatus returns OK if the IDE can serve requests.
	IDEStatus(ctx context.Context, in *IDEStatusRequest, opts ...grpc.CallOption) (*IDEStatusResponse, error)
	// BackupStatus offers feedback on the workspace backup status. This status information can
	// be relayed to the user to provide transparency as to how "safe" their files/content
	// data are w.r.t. to being lost.
	BackupStatus(ctx context.Context, in *BackupStatusRequest, opts ...grpc.CallOption) (*BackupStatusResponse, error)
	// ContentStatus returns the status of the workspace content. When used with `wait`, the call
	// returns when the content has become available.
	ContentStatus(ctx context.Context, in *ContentStatusRequest, opts ...grpc.CallOption) (*ContentStatusResponse, error)
}

type statusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatusServiceClient(cc grpc.ClientConnInterface) StatusServiceClient {
	return &statusServiceClient{cc}
}

func (c *statusServiceClient) SupervisorStatus(ctx context.Context, in *SupervisorStatusRequest, opts ...grpc.CallOption) (*SupervisorStatusResponse, error) {
	out := new(SupervisorStatusResponse)
	err := c.cc.Invoke(ctx, "/supervisor.StatusService/SupervisorStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusServiceClient) IDEStatus(ctx context.Context, in *IDEStatusRequest, opts ...grpc.CallOption) (*IDEStatusResponse, error) {
	out := new(IDEStatusResponse)
	err := c.cc.Invoke(ctx, "/supervisor.StatusService/IDEStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusServiceClient) BackupStatus(ctx context.Context, in *BackupStatusRequest, opts ...grpc.CallOption) (*BackupStatusResponse, error) {
	out := new(BackupStatusResponse)
	err := c.cc.Invoke(ctx, "/supervisor.StatusService/BackupStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusServiceClient) ContentStatus(ctx context.Context, in *ContentStatusRequest, opts ...grpc.CallOption) (*ContentStatusResponse, error) {
	out := new(ContentStatusResponse)
	err := c.cc.Invoke(ctx, "/supervisor.StatusService/ContentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatusServiceServer is the server API for StatusService service.
type StatusServiceServer interface {
	// SupervisorStatus returns once supervisor is running.
	SupervisorStatus(context.Context, *SupervisorStatusRequest) (*SupervisorStatusResponse, error)
	// IDEStatus returns OK if the IDE can serve requests.
	IDEStatus(context.Context, *IDEStatusRequest) (*IDEStatusResponse, error)
	// BackupStatus offers feedback on the workspace backup status. This status information can
	// be relayed to the user to provide transparency as to how "safe" their files/content
	// data are w.r.t. to being lost.
	BackupStatus(context.Context, *BackupStatusRequest) (*BackupStatusResponse, error)
	// ContentStatus returns the status of the workspace content. When used with `wait`, the call
	// returns when the content has become available.
	ContentStatus(context.Context, *ContentStatusRequest) (*ContentStatusResponse, error)
}

// UnimplementedStatusServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStatusServiceServer struct {
}

func (*UnimplementedStatusServiceServer) SupervisorStatus(ctx context.Context, req *SupervisorStatusRequest) (*SupervisorStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupervisorStatus not implemented")
}
func (*UnimplementedStatusServiceServer) IDEStatus(ctx context.Context, req *IDEStatusRequest) (*IDEStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IDEStatus not implemented")
}
func (*UnimplementedStatusServiceServer) BackupStatus(ctx context.Context, req *BackupStatusRequest) (*BackupStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BackupStatus not implemented")
}
func (*UnimplementedStatusServiceServer) ContentStatus(ctx context.Context, req *ContentStatusRequest) (*ContentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContentStatus not implemented")
}

func RegisterStatusServiceServer(s *grpc.Server, srv StatusServiceServer) {
	s.RegisterService(&_StatusService_serviceDesc, srv)
}

func _StatusService_SupervisorStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupervisorStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).SupervisorStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.StatusService/SupervisorStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).SupervisorStatus(ctx, req.(*SupervisorStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatusService_IDEStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDEStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).IDEStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.StatusService/IDEStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).IDEStatus(ctx, req.(*IDEStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatusService_BackupStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).BackupStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.StatusService/BackupStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).BackupStatus(ctx, req.(*BackupStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatusService_ContentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).ContentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.StatusService/ContentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).ContentStatus(ctx, req.(*ContentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "supervisor.StatusService",
	HandlerType: (*StatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SupervisorStatus",
			Handler:    _StatusService_SupervisorStatus_Handler,
		},
		{
			MethodName: "IDEStatus",
			Handler:    _StatusService_IDEStatus_Handler,
		},
		{
			MethodName: "BackupStatus",
			Handler:    _StatusService_BackupStatus_Handler,
		},
		{
			MethodName: "ContentStatus",
			Handler:    _StatusService_ContentStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supervisor.proto",
}