// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

/*
Package protobeans is a generated protocol buffer package.

It is generated from these files:
	account.proto
	friends.proto
	message.proto

It has these top-level messages:
	Login
	LoginResponse
	Regist
	RegistResponse
	SearchUser
	SearchUserResp
	AddFriend
	AddFriendResp
	AddFriendConfirm
	NewFriend
	Friend
	FriendsList
	RemoveFriend
	TextMessage
	TextMessgeResponse
	ReceiveTextMessage
*/
package protobeans

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Login struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *Login) Reset()                    { *m = Login{} }
func (m *Login) String() string            { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()               {}
func (*Login) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Login) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Login) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Ok    bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *LoginResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Regist struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *Regist) Reset()                    { *m = Regist{} }
func (m *Regist) String() string            { return proto.CompactTextString(m) }
func (*Regist) ProtoMessage()               {}
func (*Regist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Regist) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Regist) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegistResponse struct {
	Ok       bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error    string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	UserId   int32  `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	UserName string `protobuf:"bytes,4,opt,name=userName" json:"userName,omitempty"`
}

func (m *RegistResponse) Reset()                    { *m = RegistResponse{} }
func (m *RegistResponse) String() string            { return proto.CompactTextString(m) }
func (*RegistResponse) ProtoMessage()               {}
func (*RegistResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RegistResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *RegistResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *RegistResponse) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RegistResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*Login)(nil), "protobeans.Login")
	proto.RegisterType((*LoginResponse)(nil), "protobeans.LoginResponse")
	proto.RegisterType((*Regist)(nil), "protobeans.Regist")
	proto.RegisterType((*RegistResponse)(nil), "protobeans.RegistResponse")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x41, 0x0b, 0x82, 0x40,
	0x10, 0x85, 0xd1, 0x52, 0x6c, 0x40, 0x0f, 0x4b, 0x84, 0x74, 0x12, 0x4f, 0x9e, 0xba, 0x44, 0xe7,
	0xba, 0x06, 0xd1, 0x61, 0xff, 0xc1, 0xaa, 0x83, 0x98, 0xb4, 0x23, 0x3b, 0x4a, 0x7f, 0x3f, 0x76,
	0xd5, 0x3a, 0xd7, 0x69, 0xf9, 0x78, 0xec, 0xf7, 0x1e, 0x03, 0xb1, 0xaa, 0x2a, 0x1a, 0xf5, 0x70,
	0xe8, 0x0d, 0x0d, 0x24, 0xc0, 0x3d, 0x25, 0x2a, 0xcd, 0xf9, 0x19, 0x82, 0x1b, 0x35, 0xad, 0x16,
	0x7b, 0x88, 0x46, 0x46, 0xa3, 0xd5, 0x13, 0x53, 0x2f, 0xf3, 0x8a, 0x8d, 0xfc, 0xb0, 0xcd, 0x7a,
	0xc5, 0xfc, 0x22, 0x53, 0xa7, 0xfe, 0x94, 0x2d, 0x9c, 0x9f, 0x20, 0x76, 0x02, 0x89, 0xdc, 0x93,
	0x66, 0x14, 0x09, 0xf8, 0xd4, 0x39, 0x45, 0x24, 0x7d, 0xea, 0xc4, 0x16, 0x02, 0x34, 0x86, 0xcc,
	0xfc, 0x73, 0x82, 0xfc, 0x02, 0xa1, 0xc4, 0xa6, 0xe5, 0xe1, 0xef, 0xe2, 0x07, 0x24, 0x93, 0xe1,
	0xb7, 0x66, 0xb1, 0x83, 0xd0, 0xfa, 0xaf, 0x75, 0xba, 0xca, 0xbc, 0x22, 0x90, 0x33, 0x2d, 0x3b,
	0xee, 0x76, 0xc7, 0xfa, 0xbb, 0xc3, 0x72, 0x19, 0xba, 0x8b, 0x1d, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xa4, 0x2b, 0x43, 0x59, 0x49, 0x01, 0x00, 0x00,
}