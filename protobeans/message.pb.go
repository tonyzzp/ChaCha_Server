// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package protobeans

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TextMessage struct {
	Sequence int64  `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Receiver int32  `protobuf:"varint,2,opt,name=receiver" json:"receiver,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
	SendTime int64  `protobuf:"varint,4,opt,name=sendTime" json:"sendTime,omitempty"`
}

func (m *TextMessage) Reset()                    { *m = TextMessage{} }
func (m *TextMessage) String() string            { return proto.CompactTextString(m) }
func (*TextMessage) ProtoMessage()               {}
func (*TextMessage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *TextMessage) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *TextMessage) GetReceiver() int32 {
	if m != nil {
		return m.Receiver
	}
	return 0
}

func (m *TextMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TextMessage) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

type TextMessgeResponse struct {
	Ok       bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error    string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Sequence int64  `protobuf:"varint,3,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *TextMessgeResponse) Reset()                    { *m = TextMessgeResponse{} }
func (m *TextMessgeResponse) String() string            { return proto.CompactTextString(m) }
func (*TextMessgeResponse) ProtoMessage()               {}
func (*TextMessgeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *TextMessgeResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *TextMessgeResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *TextMessgeResponse) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

type ReceiveTextMessage struct {
	Sender   int32  `protobuf:"varint,1,opt,name=sender" json:"sender,omitempty"`
	Content  string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	SendTime int64  `protobuf:"varint,3,opt,name=sendTime" json:"sendTime,omitempty"`
}

func (m *ReceiveTextMessage) Reset()                    { *m = ReceiveTextMessage{} }
func (m *ReceiveTextMessage) String() string            { return proto.CompactTextString(m) }
func (*ReceiveTextMessage) ProtoMessage()               {}
func (*ReceiveTextMessage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ReceiveTextMessage) GetSender() int32 {
	if m != nil {
		return m.Sender
	}
	return 0
}

func (m *ReceiveTextMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ReceiveTextMessage) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func init() {
	proto.RegisterType((*TextMessage)(nil), "protobeans.TextMessage")
	proto.RegisterType((*TextMessgeResponse)(nil), "protobeans.TextMessgeResponse")
	proto.RegisterType((*ReceiveTextMessage)(nil), "protobeans.ReceiveTextMessage")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x6a, 0xc3, 0x30,
	0x18, 0x84, 0xb1, 0x5d, 0xa7, 0xc9, 0x5f, 0xda, 0x41, 0x84, 0x20, 0x3a, 0x19, 0x4f, 0x9e, 0xbc,
	0xf4, 0x0d, 0xba, 0x77, 0xf9, 0x09, 0xdd, 0x65, 0xf5, 0x48, 0x4c, 0xb0, 0xe4, 0x4a, 0x6a, 0x29,
	0xe9, 0xcb, 0x97, 0x48, 0x75, 0xb0, 0x87, 0x4c, 0xc7, 0x87, 0xc4, 0xdd, 0x27, 0xd1, 0xe3, 0x00,
	0xef, 0xd5, 0x01, 0xed, 0xe8, 0x6c, 0xb0, 0x82, 0x62, 0x74, 0x50, 0xc6, 0xd7, 0xbf, 0xf4, 0xb0,
	0xc7, 0x4f, 0x78, 0x4b, 0x17, 0xc4, 0x33, 0xad, 0x3d, 0x3e, 0xbf, 0x60, 0x34, 0x64, 0x56, 0x65,
	0x4d, 0xc1, 0x57, 0xbe, 0x9c, 0x39, 0x68, 0xf4, 0xdf, 0x70, 0x32, 0xaf, 0xb2, 0xa6, 0xe4, 0x2b,
	0x0b, 0x49, 0xf7, 0xda, 0x9a, 0x00, 0x13, 0x64, 0x51, 0x65, 0xcd, 0x86, 0x27, 0x4c, 0x8d, 0xe6,
	0x63, 0xdf, 0x0f, 0x90, 0x77, 0x53, 0x63, 0xe2, 0xfa, 0x9d, 0xc4, 0x34, 0x7e, 0x00, 0xc3, 0x8f,
	0xd6, 0x78, 0x88, 0x27, 0xca, 0xed, 0x29, 0xae, 0xaf, 0x39, 0xb7, 0x27, 0xb1, 0xa5, 0x12, 0xce,
	0xd9, 0x34, 0xba, 0xe1, 0x04, 0x0b, 0xd3, 0x62, 0x69, 0x5a, 0x77, 0x24, 0x38, 0x99, 0xcd, 0xdf,
	0xb6, 0xa3, 0xd5, 0x65, 0x19, 0x2e, 0x76, 0x97, 0xfc, 0x4f, 0x73, 0xf7, 0xfc, 0xb6, 0x7b, 0xb1,
	0x74, 0x7f, 0xdd, 0xd1, 0x76, 0x40, 0xdb, 0x9f, 0xcf, 0x63, 0xab, 0x8f, 0x4a, 0x1f, 0x55, 0xfa,
	0xdc, 0x6e, 0x15, 0xe3, 0xe5, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x52, 0xbd, 0x29, 0x74, 0x01,
	0x00, 0x00,
}
