// Code generated by protoc-gen-go. DO NOT EDIT.
// source: team.proto

package pbs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Team struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b4e9e93d7b2c6bb, []int{0}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*Team)(nil), "events.Team")
}

func init() { proto.RegisterFile("team.proto", fileDescriptor_8b4e9e93d7b2c6bb) }

var fileDescriptor_8b4e9e93d7b2c6bb = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0x8e, 0xb1, 0xca, 0xc2, 0x30,
	0x14, 0x85, 0x69, 0xff, 0x52, 0xf8, 0xaf, 0xe2, 0x90, 0xa9, 0x38, 0x15, 0xa7, 0x0e, 0x9a, 0x80,
	0xbe, 0x81, 0xb3, 0x53, 0xd1, 0xc5, 0x2d, 0x49, 0x2f, 0x36, 0x60, 0x9a, 0x90, 0x7b, 0xe3, 0xf3,
	0x8b, 0xe9, 0x76, 0xf8, 0x0e, 0xe7, 0xf0, 0x01, 0x30, 0x6a, 0x2f, 0x63, 0x0a, 0x1c, 0x44, 0x8b,
	0x1f, 0x5c, 0x98, 0xf6, 0x90, 0x09, 0xd3, 0xca, 0x0e, 0x37, 0x68, 0xee, 0xa8, 0xbd, 0xd8, 0x41,
	0xed, 0xa6, 0xae, 0xea, 0xab, 0xe1, 0x7f, 0xac, 0xdd, 0x24, 0x04, 0x34, 0x8b, 0xf6, 0xd8, 0xd5,
	0x85, 0x94, 0x2c, 0x7a, 0x68, 0x7e, 0xcb, 0xee, 0xaf, 0xaf, 0x86, 0xcd, 0x79, 0x2b, 0xd7, 0x3b,
	0xf9, 0x20, 0x4c, 0x63, 0x69, 0xae, 0xf2, 0x79, 0x7c, 0x39, 0x9e, 0xb3, 0x91, 0x36, 0x78, 0x65,
	0x34, 0xdb, 0xd9, 0x86, 0x14, 0x55, 0x7c, 0x67, 0x6f, 0x30, 0x29, 0x46, 0xe2, 0x93, 0x26, 0x42,
	0x26, 0x15, 0x0d, 0x99, 0xb6, 0x48, 0x5c, 0xbe, 0x01, 0x00, 0x00, 0xff, 0xff, 0x72, 0x09, 0xf1,
	0xc6, 0xa6, 0x00, 0x00, 0x00,
}
