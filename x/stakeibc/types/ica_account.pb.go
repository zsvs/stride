// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/stakeibc/ica_account.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type ICAAccountType int32

const (
	ICAAccountType_DELEGATION ICAAccountType = 0
	ICAAccountType_FEE        ICAAccountType = 1
	ICAAccountType_WITHDRAWAL ICAAccountType = 2
	ICAAccountType_REDEMPTION ICAAccountType = 3
)

var ICAAccountType_name = map[int32]string{
	0: "DELEGATION",
	1: "FEE",
	2: "WITHDRAWAL",
	3: "REDEMPTION",
}

var ICAAccountType_value = map[string]int32{
	"DELEGATION": 0,
	"FEE":        1,
	"WITHDRAWAL": 2,
	"REDEMPTION": 3,
}

func (x ICAAccountType) String() string {
	return proto.EnumName(ICAAccountType_name, int32(x))
}

func (ICAAccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2976ae6e7f6ce824, []int{0}
}

func init() {
	proto.RegisterEnum("stride.stakeibc.ICAAccountType", ICAAccountType_name, ICAAccountType_value)
}

func init() { proto.RegisterFile("stride/stakeibc/ica_account.proto", fileDescriptor_2976ae6e7f6ce824) }

var fileDescriptor_2976ae6e7f6ce824 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x2e, 0x29, 0xca,
	0x4c, 0x49, 0xd5, 0x2f, 0x2e, 0x49, 0xcc, 0x4e, 0xcd, 0x4c, 0x4a, 0xd6, 0xcf, 0x4c, 0x4e, 0x8c,
	0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x87,
	0x28, 0xd1, 0x83, 0x29, 0xd1, 0xf2, 0xe4, 0xe2, 0xf3, 0x74, 0x76, 0x74, 0x84, 0x28, 0x0a, 0xa9,
	0x2c, 0x48, 0x15, 0xe2, 0xe3, 0xe2, 0x72, 0x71, 0xf5, 0x71, 0x75, 0x77, 0x0c, 0xf1, 0xf4, 0xf7,
	0x13, 0x60, 0x10, 0x62, 0xe7, 0x62, 0x76, 0x73, 0x75, 0x15, 0x60, 0x04, 0x49, 0x84, 0x7b, 0x86,
	0x78, 0xb8, 0x04, 0x39, 0x86, 0x3b, 0xfa, 0x08, 0x30, 0x81, 0xf8, 0x41, 0xae, 0x2e, 0xae, 0xbe,
	0x01, 0x60, 0x85, 0xcc, 0x4e, 0x3e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0,
	0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10,
	0x65, 0x94, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x1f, 0x0c, 0x76, 0x80,
	0xae, 0x4f, 0x62, 0x52, 0xb1, 0x3e, 0xd4, 0xbd, 0x65, 0x86, 0x66, 0xfa, 0x15, 0x08, 0x57, 0x97,
	0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x1d, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb1,
	0x6c, 0x27, 0x78, 0xd5, 0x00, 0x00, 0x00,
}
