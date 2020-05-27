// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/sql/sql.proto

package sql // import "github.com/weiwolves/protoc-gen-sqlx/pb/sql"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ExtraField struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tags                 string   `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtraField) Reset()         { *m = ExtraField{} }
func (m *ExtraField) String() string { return proto.CompactTextString(m) }
func (*ExtraField) ProtoMessage()    {}
func (*ExtraField) Descriptor() ([]byte, []int) {
	return fileDescriptor_sql_541f48426f05c127, []int{0}
}
func (m *ExtraField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraField.Unmarshal(m, b)
}
func (m *ExtraField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraField.Marshal(b, m, deterministic)
}
func (dst *ExtraField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraField.Merge(dst, src)
}
func (m *ExtraField) XXX_Size() int {
	return xxx_messageInfo_ExtraField.Size(m)
}
func (m *ExtraField) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraField.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraField proto.InternalMessageInfo

func (m *ExtraField) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ExtraField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExtraField) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

type SqlxMessageOptions struct {
	Orm                  bool          `protobuf:"varint,1,opt,name=orm,proto3" json:"orm,omitempty"`
	Jsonb                bool          `protobuf:"varint,2,opt,name=jsonb,proto3" json:"jsonb,omitempty"`
	Gorm                 bool          `protobuf:"varint,3,opt,name=gorm,proto3" json:"gorm,omitempty"`
	Table                string        `protobuf:"bytes,4,opt,name=table,proto3" json:"table,omitempty"`
	Driver               string        `protobuf:"bytes,5,opt,name=driver,proto3" json:"driver,omitempty"`
	Include              []*ExtraField `protobuf:"bytes,6,rep,name=include" json:"include,omitempty"`
	Request              string        `protobuf:"bytes,7,opt,name=request,proto3" json:"request,omitempty"`
	User                 bool          `protobuf:"varint,8,opt,name=user,proto3" json:"user,omitempty"`
	Product              bool          `protobuf:"varint,9,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SqlxMessageOptions) Reset()         { *m = SqlxMessageOptions{} }
func (m *SqlxMessageOptions) String() string { return proto.CompactTextString(m) }
func (*SqlxMessageOptions) ProtoMessage()    {}
func (*SqlxMessageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_sql_541f48426f05c127, []int{1}
}
func (m *SqlxMessageOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqlxMessageOptions.Unmarshal(m, b)
}
func (m *SqlxMessageOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqlxMessageOptions.Marshal(b, m, deterministic)
}
func (dst *SqlxMessageOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqlxMessageOptions.Merge(dst, src)
}
func (m *SqlxMessageOptions) XXX_Size() int {
	return xxx_messageInfo_SqlxMessageOptions.Size(m)
}
func (m *SqlxMessageOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_SqlxMessageOptions.DiscardUnknown(m)
}

var xxx_messageInfo_SqlxMessageOptions proto.InternalMessageInfo

func (m *SqlxMessageOptions) GetOrm() bool {
	if m != nil {
		return m.Orm
	}
	return false
}

func (m *SqlxMessageOptions) GetJsonb() bool {
	if m != nil {
		return m.Jsonb
	}
	return false
}

func (m *SqlxMessageOptions) GetGorm() bool {
	if m != nil {
		return m.Gorm
	}
	return false
}

func (m *SqlxMessageOptions) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *SqlxMessageOptions) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *SqlxMessageOptions) GetInclude() []*ExtraField {
	if m != nil {
		return m.Include
	}
	return nil
}

func (m *SqlxMessageOptions) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

func (m *SqlxMessageOptions) GetUser() bool {
	if m != nil {
		return m.User
	}
	return false
}

func (m *SqlxMessageOptions) GetProduct() bool {
	if m != nil {
		return m.Product
	}
	return false
}

type SqlxFieldOptions struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Colname              string   `protobuf:"bytes,2,opt,name=colname,proto3" json:"colname,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Tags                 string   `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	Drop                 bool     `protobuf:"varint,5,opt,name=drop,proto3" json:"drop,omitempty"`
	Pk                   bool     `protobuf:"varint,6,opt,name=pk,proto3" json:"pk,omitempty"`
	Fk                   string   `protobuf:"bytes,7,opt,name=fk,proto3" json:"fk,omitempty"`
	Customname           string   `protobuf:"bytes,8,opt,name=customname,proto3" json:"customname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SqlxFieldOptions) Reset()         { *m = SqlxFieldOptions{} }
func (m *SqlxFieldOptions) String() string { return proto.CompactTextString(m) }
func (*SqlxFieldOptions) ProtoMessage()    {}
func (*SqlxFieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_sql_541f48426f05c127, []int{2}
}
func (m *SqlxFieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqlxFieldOptions.Unmarshal(m, b)
}
func (m *SqlxFieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqlxFieldOptions.Marshal(b, m, deterministic)
}
func (dst *SqlxFieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqlxFieldOptions.Merge(dst, src)
}
func (m *SqlxFieldOptions) XXX_Size() int {
	return xxx_messageInfo_SqlxFieldOptions.Size(m)
}
func (m *SqlxFieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_SqlxFieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_SqlxFieldOptions proto.InternalMessageInfo

func (m *SqlxFieldOptions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SqlxFieldOptions) GetColname() string {
	if m != nil {
		return m.Colname
	}
	return ""
}

func (m *SqlxFieldOptions) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SqlxFieldOptions) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *SqlxFieldOptions) GetDrop() bool {
	if m != nil {
		return m.Drop
	}
	return false
}

func (m *SqlxFieldOptions) GetPk() bool {
	if m != nil {
		return m.Pk
	}
	return false
}

func (m *SqlxFieldOptions) GetFk() string {
	if m != nil {
		return m.Fk
	}
	return ""
}

func (m *SqlxFieldOptions) GetCustomname() string {
	if m != nil {
		return m.Customname
	}
	return ""
}

type SqlxServiceOptions struct {
	Autogen              bool     `protobuf:"varint,1,opt,name=autogen,proto3" json:"autogen,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SqlxServiceOptions) Reset()         { *m = SqlxServiceOptions{} }
func (m *SqlxServiceOptions) String() string { return proto.CompactTextString(m) }
func (*SqlxServiceOptions) ProtoMessage()    {}
func (*SqlxServiceOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_sql_541f48426f05c127, []int{3}
}
func (m *SqlxServiceOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqlxServiceOptions.Unmarshal(m, b)
}
func (m *SqlxServiceOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqlxServiceOptions.Marshal(b, m, deterministic)
}
func (dst *SqlxServiceOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqlxServiceOptions.Merge(dst, src)
}
func (m *SqlxServiceOptions) XXX_Size() int {
	return xxx_messageInfo_SqlxServiceOptions.Size(m)
}
func (m *SqlxServiceOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_SqlxServiceOptions.DiscardUnknown(m)
}

var xxx_messageInfo_SqlxServiceOptions proto.InternalMessageInfo

func (m *SqlxServiceOptions) GetAutogen() bool {
	if m != nil {
		return m.Autogen
	}
	return false
}

type SqlxMethodOptions struct {
	Jsonb                bool     `protobuf:"varint,1,opt,name=jsonb,proto3" json:"jsonb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SqlxMethodOptions) Reset()         { *m = SqlxMethodOptions{} }
func (m *SqlxMethodOptions) String() string { return proto.CompactTextString(m) }
func (*SqlxMethodOptions) ProtoMessage()    {}
func (*SqlxMethodOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_sql_541f48426f05c127, []int{4}
}
func (m *SqlxMethodOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqlxMethodOptions.Unmarshal(m, b)
}
func (m *SqlxMethodOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqlxMethodOptions.Marshal(b, m, deterministic)
}
func (dst *SqlxMethodOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqlxMethodOptions.Merge(dst, src)
}
func (m *SqlxMethodOptions) XXX_Size() int {
	return xxx_messageInfo_SqlxMethodOptions.Size(m)
}
func (m *SqlxMethodOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_SqlxMethodOptions.DiscardUnknown(m)
}

var xxx_messageInfo_SqlxMethodOptions proto.InternalMessageInfo

func (m *SqlxMethodOptions) GetJsonb() bool {
	if m != nil {
		return m.Jsonb
	}
	return false
}

var E_Opts = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*SqlxMessageOptions)(nil),
	Field:         99901,
	Name:          "sql.opts",
	Tag:           "bytes,99901,opt,name=opts",
	Filename:      "pb/sql/sql.proto",
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*SqlxFieldOptions)(nil),
	Field:         99902,
	Name:          "sql.field",
	Tag:           "bytes,99902,opt,name=field",
	Filename:      "pb/sql/sql.proto",
}

var E_Server = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*SqlxServiceOptions)(nil),
	Field:         99903,
	Name:          "sql.server",
	Tag:           "bytes,99903,opt,name=server",
	Filename:      "pb/sql/sql.proto",
}

var E_Method = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*SqlxMethodOptions)(nil),
	Field:         99904,
	Name:          "sql.method",
	Tag:           "bytes,99904,opt,name=method",
	Filename:      "pb/sql/sql.proto",
}

func init() {
	proto.RegisterType((*ExtraField)(nil), "sql.ExtraField")
	proto.RegisterType((*SqlxMessageOptions)(nil), "sql.SqlxMessageOptions")
	proto.RegisterType((*SqlxFieldOptions)(nil), "sql.SqlxFieldOptions")
	proto.RegisterType((*SqlxServiceOptions)(nil), "sql.SqlxServiceOptions")
	proto.RegisterType((*SqlxMethodOptions)(nil), "sql.SqlxMethodOptions")
	proto.RegisterExtension(E_Opts)
	proto.RegisterExtension(E_Field)
	proto.RegisterExtension(E_Server)
	proto.RegisterExtension(E_Method)
}

func init() { proto.RegisterFile("pb/sql/sql.proto", fileDescriptor_sql_541f48426f05c127) }

var fileDescriptor_sql_541f48426f05c127 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcb, 0x8e, 0xd3, 0x4a,
	0x10, 0x95, 0xe3, 0xc4, 0xc9, 0xd4, 0x48, 0xf7, 0x86, 0x16, 0x0c, 0x2d, 0x24, 0x86, 0x28, 0xab,
	0x99, 0x45, 0x1c, 0x14, 0x76, 0x61, 0x87, 0x04, 0x62, 0xc1, 0x4b, 0x9e, 0x1d, 0x3b, 0x3f, 0x3a,
	0x1e, 0x93, 0xb6, 0xbb, 0xdd, 0xdd, 0x1e, 0x85, 0x7f, 0xe0, 0x93, 0x78, 0xfd, 0x12, 0x7f, 0x80,
	0xba, 0xda, 0xf6, 0x38, 0x0a, 0x0b, 0x4b, 0x55, 0xa7, 0x4a, 0xc7, 0x75, 0x4e, 0x55, 0xc3, 0x5c,
	0x26, 0x6b, 0x5d, 0x73, 0xfb, 0x85, 0x52, 0x09, 0x23, 0x88, 0xaf, 0x6b, 0xfe, 0x64, 0x91, 0x0b,
	0x91, 0x73, 0xb6, 0x46, 0x28, 0x69, 0x76, 0xeb, 0x8c, 0xe9, 0x54, 0x15, 0xd2, 0x08, 0xe5, 0xda,
	0x96, 0x6f, 0x01, 0x5e, 0x1f, 0x8c, 0x8a, 0xdf, 0x14, 0x8c, 0x67, 0x84, 0xc0, 0xd8, 0x7c, 0x95,
	0x8c, 0x7a, 0x0b, 0xef, 0xea, 0x2c, 0xc2, 0xd8, 0x62, 0x55, 0x5c, 0x32, 0x3a, 0x72, 0x98, 0x8d,
	0xb1, 0x2f, 0xce, 0x35, 0xf5, 0xdb, 0xbe, 0x38, 0xd7, 0xcb, 0x3f, 0x1e, 0x90, 0x9b, 0x9a, 0x1f,
	0xde, 0x33, 0xad, 0xe3, 0x9c, 0x7d, 0x94, 0xa6, 0x10, 0x95, 0x26, 0x73, 0xf0, 0x85, 0x2a, 0x91,
	0x71, 0x16, 0xd9, 0x90, 0x3c, 0x84, 0xc9, 0x17, 0x2d, 0xaa, 0x04, 0x19, 0x67, 0x91, 0x4b, 0x2c,
	0x65, 0x6e, 0x1b, 0x7d, 0x04, 0x31, 0xb6, 0x9d, 0x26, 0x4e, 0x38, 0xa3, 0x63, 0xfc, 0x8f, 0x4b,
	0xc8, 0x05, 0x04, 0x99, 0x2a, 0xee, 0x98, 0xa2, 0x13, 0x84, 0xdb, 0x8c, 0x5c, 0xc3, 0xb4, 0xa8,
	0x52, 0xde, 0x64, 0x8c, 0x06, 0x0b, 0xff, 0xea, 0x7c, 0xf3, 0x7f, 0x68, 0xed, 0xb8, 0x97, 0x17,
	0x75, 0x75, 0x42, 0x61, 0xaa, 0x58, 0xdd, 0x30, 0x6d, 0xe8, 0x14, 0x39, 0xba, 0xd4, 0x8e, 0xd1,
	0x68, 0xa6, 0xe8, 0xcc, 0x8d, 0x61, 0x63, 0xdb, 0x2d, 0x95, 0xc8, 0x9a, 0xd4, 0xd0, 0x33, 0x84,
	0xbb, 0x74, 0xf9, 0xdd, 0x83, 0xb9, 0xd5, 0x8c, 0xf4, 0x9d, 0xe2, 0xce, 0x30, 0x6f, 0x60, 0x18,
	0x85, 0x69, 0x2a, 0xf8, 0xc0, 0xc7, 0x2e, 0xed, 0x2d, 0xf7, 0x8f, 0x2d, 0x47, 0x7b, 0xc7, 0xf7,
	0xf6, 0x5a, 0x2c, 0x53, 0x42, 0xa2, 0xe6, 0x59, 0x84, 0x31, 0xf9, 0x0f, 0x46, 0x72, 0x4f, 0x03,
	0x44, 0x46, 0x72, 0x6f, 0xf3, 0xdd, 0xbe, 0x55, 0x34, 0xda, 0xed, 0xc9, 0x25, 0x40, 0xda, 0x68,
	0x23, 0x4a, 0xfc, 0xf1, 0x0c, 0xf1, 0x01, 0xb2, 0x0c, 0xdd, 0xc6, 0x6e, 0x98, 0xba, 0x2b, 0xd2,
	0x7e, 0x63, 0x14, 0xa6, 0x71, 0x63, 0x44, 0xce, 0xaa, 0x76, 0x6b, 0x5d, 0xba, 0xbc, 0x86, 0x07,
	0x6e, 0xc3, 0xe6, 0x56, 0xf4, 0x72, 0xfb, 0x75, 0x7a, 0x83, 0x75, 0x6e, 0x3f, 0xc0, 0x58, 0x48,
	0xa3, 0xc9, 0xb3, 0xd0, 0x9d, 0x60, 0xd8, 0x9d, 0x60, 0x78, 0x7c, 0x1f, 0xf4, 0xc7, 0x37, 0x3b,
	0xff, 0xf9, 0xe6, 0x31, 0x2e, 0xeb, 0xf4, 0x80, 0x22, 0xe4, 0xd9, 0xbe, 0x83, 0xc9, 0x0e, 0x4f,
	0xf4, 0xe9, 0x09, 0xe1, 0xd0, 0x7c, 0xfa, 0xb3, 0xa5, 0x7b, 0xd4, 0xd3, 0x0d, 0xcb, 0x91, 0x23,
	0xd9, 0x46, 0x10, 0x68, 0xa6, 0xec, 0xd1, 0x9c, 0xce, 0x77, 0xec, 0x06, 0xfd, 0x75, 0x32, 0xdf,
	0x71, 0x43, 0xd4, 0x32, 0x6d, 0x3f, 0x41, 0x50, 0xa2, 0x31, 0xe4, 0xf2, 0x1f, 0x9a, 0x07, 0x8e,
	0xd1, 0xdf, 0x2d, 0xe5, 0xc5, 0x40, 0xf2, 0xa0, 0x1e, 0xb5, 0x3c, 0xaf, 0x36, 0x9f, 0x9f, 0xe7,
	0x85, 0xb9, 0x6d, 0x92, 0x30, 0x15, 0xe5, 0xba, 0x2c, 0x52, 0x25, 0x56, 0xb9, 0x92, 0xa9, 0x7b,
	0xce, 0xe9, 0x2a, 0x67, 0xd5, 0x4a, 0xd7, 0xfc, 0xb0, 0x76, 0x8f, 0xff, 0xa5, 0xae, 0x79, 0x12,
	0x60, 0xed, 0xc5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0xc1, 0xcf, 0x8d, 0x11, 0x04, 0x00,
	0x00,
}
