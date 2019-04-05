// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/email/email.proto

package email

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

type Message struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Subject              string   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_840f7b8354a62044, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Message) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "email.Message")
}

func init() { proto.RegisterFile("proto/email/email.proto", fileDescriptor_840f7b8354a62044) }

var fileDescriptor_840f7b8354a62044 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xcd, 0x4d, 0xcc, 0xcc, 0x81, 0x90, 0x7a, 0x60, 0x11, 0x21, 0x56, 0x30, 0x47,
	0x29, 0x9c, 0x8b, 0xdd, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x55, 0x48, 0x82, 0x8b, 0x3d, 0x31,
	0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x05, 0xc9,
	0x14, 0x97, 0x26, 0x65, 0xa5, 0x26, 0x97, 0x48, 0x30, 0x41, 0x64, 0xa0, 0x5c, 0x90, 0x4c, 0x72,
	0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x89, 0x04, 0x33, 0x44, 0x06, 0xca, 0x4d, 0x62, 0x03, 0x5b, 0x63,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x9b, 0xab, 0x08, 0x81, 0x00, 0x00, 0x00,
}
