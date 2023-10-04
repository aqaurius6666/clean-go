// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/entitypb/v1/entity.proto

package entitypb

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5a5d1e4f97a33e4, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	CreatorId            string   `protobuf:"bytes,3,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Creator              *User    `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5a5d1e4f97a33e4, []int{1}
}
func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetCreatorId() string {
	if m != nil {
		return m.CreatorId
	}
	return ""
}

func (m *Post) GetCreator() *User {
	if m != nil {
		return m.Creator
	}
	return nil
}

type Pagination struct {
	Offset               int      `protobuf:"varint,1,opt,name=offset,proto3,casttype=int" json:"offset" form:"offset"`
	Limit                int      `protobuf:"varint,2,opt,name=limit,proto3,casttype=int" json:"limit" form:"limit"`
	Total                int64    `protobuf:"varint,3,opt,name=total,proto3" json:"total" form:"-"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5a5d1e4f97a33e4, []int{2}
}
func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (m *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(m, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetOffset() int {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Pagination) GetLimit() int {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Pagination) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "proto.entitypb.v1.User")
	proto.RegisterType((*Post)(nil), "proto.entitypb.v1.Post")
	proto.RegisterType((*Pagination)(nil), "proto.entitypb.v1.Pagination")
}

func init() { proto.RegisterFile("proto/entitypb/v1/entity.proto", fileDescriptor_f5a5d1e4f97a33e4) }

var fileDescriptor_f5a5d1e4f97a33e4 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x6e, 0x9b, 0x40,
	0x1c, 0xc5, 0x3d, 0x7c, 0xd5, 0x9e, 0x7e, 0xc8, 0x1d, 0xb5, 0x32, 0xb2, 0xd4, 0x62, 0xb1, 0xa9,
	0x2b, 0xcb, 0x20, 0xbb, 0x2a, 0x8b, 0x7a, 0xe1, 0x76, 0xba, 0xea, 0xce, 0x42, 0xea, 0xa6, 0x9b,
	0x6a, 0x6c, 0x30, 0x19, 0x05, 0x18, 0x67, 0x18, 0x5b, 0xca, 0x15, 0x72, 0x91, 0x1c, 0x23, 0x57,
	0xc8, 0x09, 0x38, 0x40, 0x96, 0x5e, 0xb2, 0x8a, 0x98, 0x81, 0x24, 0x8a, 0xb3, 0xe2, 0xfd, 0xdf,
	0x8f, 0xf7, 0xc4, 0x93, 0x80, 0x9f, 0x77, 0x9c, 0x09, 0xe6, 0xc7, 0xb9, 0xa0, 0xe2, 0x72, 0xb7,
	0xf6, 0x0f, 0xb3, 0x46, 0x7b, 0x12, 0xa0, 0xf7, 0xf2, 0xe1, 0xb5, 0xdc, 0x3b, 0xcc, 0x86, 0x83,
	0x03, 0x49, 0x69, 0x44, 0x44, 0xec, 0xb7, 0x42, 0xbd, 0x3b, 0xfc, 0x90, 0xb0, 0x84, 0xa9, 0xbe,
	0x5a, 0x29, 0xd7, 0xfd, 0x0e, 0x8d, 0xbf, 0x45, 0xcc, 0xd1, 0x3b, 0xa8, 0xd1, 0xc8, 0x06, 0x23,
	0x30, 0xee, 0x85, 0x1a, 0x8d, 0xd0, 0x27, 0x68, 0xe4, 0x24, 0x8b, 0x6d, 0xad, 0x76, 0x70, 0xaf,
	0xc2, 0x16, 0x37, 0xfa, 0xc0, 0x9e, 0x87, 0xd2, 0x76, 0xaf, 0x01, 0x34, 0x56, 0xac, 0x10, 0xc8,
	0x7e, 0xcc, 0xe1, 0x6e, 0x85, 0x4d, 0xae, 0xdf, 0x02, 0x20, 0x1b, 0x1c, 0x68, 0x0a, 0x2a, 0xd2,
	0x17, 0x2a, 0x94, 0x8f, 0xbe, 0x40, 0xb8, 0xe1, 0x31, 0x11, 0x8c, 0xff, 0xa7, 0x91, 0xad, 0x3f,
	0xab, 0xe8, 0x35, 0xec, 0x4f, 0x84, 0x16, 0xf0, 0x55, 0x73, 0xd8, 0xc6, 0x08, 0x8c, 0x5f, 0xcf,
	0x07, 0xde, 0xc9, 0x6e, 0xaf, 0x5e, 0x21, 0xe3, 0x57, 0x40, 0xeb, 0x77, 0xc2, 0x36, 0xe1, 0xde,
	0x00, 0x08, 0x57, 0x24, 0xa1, 0x39, 0x11, 0x94, 0xe5, 0xe8, 0x37, 0xb4, 0xd8, 0x76, 0x5b, 0xc4,
	0x42, 0x7e, 0xb3, 0x89, 0x27, 0x15, 0xb6, 0x86, 0xc6, 0xb8, 0xf3, 0x13, 0xdc, 0x95, 0x4e, 0x83,
	0x8e, 0xa5, 0xf3, 0x76, 0xcb, 0x78, 0xf6, 0xc3, 0x55, 0xb7, 0x5b, 0x95, 0x8e, 0x4e, 0x73, 0x11,
	0x36, 0x1c, 0x2d, 0xa1, 0x99, 0xd2, 0x8c, 0x0a, 0x39, 0xcd, 0xc4, 0x5f, 0x9f, 0x76, 0x28, 0x72,
	0x2c, 0x9d, 0x37, 0xaa, 0x42, 0x9e, 0x0f, 0x0d, 0x8a, 0xa2, 0x09, 0x34, 0x05, 0x13, 0x24, 0x95,
	0xab, 0x75, 0xfc, 0xb1, 0xce, 0x49, 0xe3, 0x58, 0x3a, 0x5d, 0x95, 0x9b, 0xba, 0xa1, 0xb2, 0xf0,
	0xaf, 0x7f, 0xcb, 0x84, 0x8a, 0xb3, 0xfd, 0xda, 0xdb, 0xb0, 0xcc, 0x27, 0x17, 0x64, 0xcf, 0xe9,
	0xbe, 0x08, 0x82, 0x20, 0xf0, 0x37, 0x69, 0x4c, 0xf2, 0x69, 0xc2, 0xfc, 0xdd, 0x79, 0xe2, 0x9f,
	0xfc, 0x2b, 0x8b, 0x56, 0xaf, 0x2d, 0xc9, 0xbe, 0xdd, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x3b,
	0x78, 0xa8, 0x50, 0x02, 0x00, 0x00,
}