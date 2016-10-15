// Code generated by protoc-gen-go.
// source: v2ray.com/core/proxy/shadowsocks/config.proto
// DO NOT EDIT!

/*
Package shadowsocks is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/proxy/shadowsocks/config.proto

It has these top-level messages:
	Account
	ServerConfig
	ClientConfig
*/
package shadowsocks

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_protocol "v2ray.com/core/common/protocol"
import v2ray_core_common_protocol1 "v2ray.com/core/common/protocol"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CipherType int32

const (
	CipherType_UNKNOWN       CipherType = 0
	CipherType_AES_128_CFB   CipherType = 1
	CipherType_AES_256_CFB   CipherType = 2
	CipherType_CHACHA20      CipherType = 3
	CipherType_CHACHA20_IEFT CipherType = 4
)

var CipherType_name = map[int32]string{
	0: "UNKNOWN",
	1: "AES_128_CFB",
	2: "AES_256_CFB",
	3: "CHACHA20",
	4: "CHACHA20_IEFT",
}
var CipherType_value = map[string]int32{
	"UNKNOWN":       0,
	"AES_128_CFB":   1,
	"AES_256_CFB":   2,
	"CHACHA20":      3,
	"CHACHA20_IEFT": 4,
}

func (x CipherType) String() string {
	return proto.EnumName(CipherType_name, int32(x))
}
func (CipherType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Account struct {
	Password   string     `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	CipherType CipherType `protobuf:"varint,2,opt,name=cipher_type,json=cipherType,enum=v2ray.core.proxy.shadowsocks.CipherType" json:"cipher_type,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServerConfig struct {
	UdpEnabled bool                             `protobuf:"varint,1,opt,name=udp_enabled,json=udpEnabled" json:"udp_enabled,omitempty"`
	User       *v2ray_core_common_protocol.User `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServerConfig) GetUser() *v2ray_core_common_protocol.User {
	if m != nil {
		return m.User
	}
	return nil
}

type ClientConfig struct {
	Server []*v2ray_core_common_protocol1.ServerEndpoint `protobuf:"bytes,1,rep,name=server" json:"server,omitempty"`
}

func (m *ClientConfig) Reset()                    { *m = ClientConfig{} }
func (m *ClientConfig) String() string            { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()               {}
func (*ClientConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientConfig) GetServer() []*v2ray_core_common_protocol1.ServerEndpoint {
	if m != nil {
		return m.Server
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "v2ray.core.proxy.shadowsocks.Account")
	proto.RegisterType((*ServerConfig)(nil), "v2ray.core.proxy.shadowsocks.ServerConfig")
	proto.RegisterType((*ClientConfig)(nil), "v2ray.core.proxy.shadowsocks.ClientConfig")
	proto.RegisterEnum("v2ray.core.proxy.shadowsocks.CipherType", CipherType_name, CipherType_value)
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/shadowsocks/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x50, 0xdd, 0xab, 0xd3, 0x30,
	0x14, 0xb7, 0xf7, 0x5e, 0xee, 0x9d, 0x27, 0x53, 0x6b, 0x9e, 0xc6, 0x10, 0x2c, 0x7b, 0xaa, 0x17,
	0x4c, 0x67, 0xfd, 0xc0, 0x07, 0x5f, 0xda, 0xd2, 0xb1, 0x21, 0x4c, 0xe9, 0x36, 0x04, 0x11, 0x4a,
	0x97, 0x46, 0x57, 0x5c, 0x9b, 0x90, 0xb4, 0x9b, 0xfd, 0xef, 0x65, 0xc9, 0x3a, 0x87, 0x0f, 0xbb,
	0x6f, 0x39, 0x27, 0xbf, 0xcf, 0x03, 0xaf, 0x77, 0xbe, 0xcc, 0x5a, 0x42, 0x79, 0xe9, 0x51, 0x2e,
	0x99, 0x27, 0x24, 0xff, 0xd3, 0x7a, 0x6a, 0x93, 0xe5, 0x7c, 0xaf, 0x38, 0xfd, 0xad, 0x3c, 0xca,
	0xab, 0x9f, 0xc5, 0x2f, 0x22, 0x24, 0xaf, 0x39, 0x7e, 0xd1, 0xc1, 0x25, 0x23, 0x1a, 0x4a, 0xce,
	0xa0, 0xc3, 0x57, 0xff, 0x89, 0x51, 0x5e, 0x96, 0xbc, 0xf2, 0x34, 0x95, 0xf2, 0xad, 0xd7, 0x28,
	0x26, 0x8d, 0xd0, 0x70, 0xfc, 0x00, 0x54, 0x31, 0xb9, 0x63, 0x32, 0x55, 0x82, 0x51, 0xc3, 0x18,
	0x09, 0xb8, 0x0b, 0x28, 0xe5, 0x4d, 0x55, 0xe3, 0x21, 0xf4, 0x44, 0xa6, 0xd4, 0x9e, 0xcb, 0x7c,
	0x60, 0x39, 0x96, 0xfb, 0x38, 0x39, 0xcd, 0x78, 0x06, 0x88, 0x16, 0x62, 0xc3, 0x64, 0x5a, 0xb7,
	0x82, 0x0d, 0xae, 0x1c, 0xcb, 0x7d, 0xea, 0xbb, 0xe4, 0x52, 0x6e, 0x12, 0x69, 0xc2, 0xb2, 0x15,
	0x2c, 0x01, 0x7a, 0x7a, 0x8f, 0x18, 0xf4, 0x17, 0x3a, 0x46, 0xa4, 0x4f, 0x80, 0x5f, 0x02, 0x6a,
	0x72, 0x91, 0xb2, 0x2a, 0x5b, 0x6f, 0x99, 0x71, 0xee, 0x25, 0xd0, 0xe4, 0x22, 0x36, 0x1b, 0xfc,
	0x0e, 0x6e, 0x0e, 0x15, 0xb5, 0x29, 0xf2, 0x9d, 0x73, 0x53, 0xd3, 0x8f, 0x74, 0xfd, 0xc8, 0x4a,
	0x31, 0x99, 0x68, 0xf4, 0x28, 0x81, 0x7e, 0xb4, 0x2d, 0x58, 0x55, 0x1f, 0x6d, 0x42, 0xb8, 0x35,
	0xed, 0x07, 0x96, 0x73, 0xed, 0x22, 0xff, 0xfe, 0x92, 0x8e, 0x09, 0x18, 0x57, 0xb9, 0xe0, 0x45,
	0x55, 0x27, 0x47, 0xe6, 0xfd, 0x0f, 0x80, 0x7f, 0xa5, 0x30, 0x82, 0xbb, 0xd5, 0xfc, 0xf3, 0xfc,
	0xcb, 0xb7, 0xb9, 0xfd, 0x08, 0x3f, 0x03, 0x14, 0xc4, 0x8b, 0xf4, 0x8d, 0xff, 0x31, 0x8d, 0x26,
	0xa1, 0x6d, 0x75, 0x0b, 0xff, 0xfd, 0x07, 0xbd, 0xb8, 0xc2, 0x7d, 0xe8, 0x45, 0xd3, 0x20, 0x9a,
	0x06, 0xfe, 0xd8, 0xbe, 0xc6, 0xcf, 0xe1, 0x49, 0x37, 0xa5, 0xb3, 0x78, 0xb2, 0xb4, 0x6f, 0xc2,
	0x4f, 0xe0, 0x50, 0x5e, 0x5e, 0xbc, 0x69, 0x88, 0x4c, 0x9b, 0xaf, 0x87, 0xa0, 0xdf, 0xd1, 0xd9,
	0xcf, 0xfa, 0x56, 0x87, 0x7f, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xff, 0xed, 0x11, 0xa8, 0x7b,
	0x02, 0x00, 0x00,
}