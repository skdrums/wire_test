// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/apis/model_service.proto

package tensorflow_serving

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

func init() {
	proto.RegisterFile("tensorflow_serving/apis/model_service.proto", fileDescriptor_59da5a1442bba4ef)
}

var fileDescriptor_59da5a1442bba4ef = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xcf, 0xbd, 0x0a, 0xc2, 0x30,
	0x10, 0xc0, 0x71, 0x8a, 0xe0, 0x10, 0xc4, 0x21, 0x9b, 0x7d, 0x01, 0xbf, 0x20, 0x05, 0x7d, 0x03,
	0x1d, 0x74, 0x71, 0xa9, 0x0f, 0x50, 0xa2, 0xbd, 0x86, 0x42, 0x7a, 0x57, 0x9b, 0xab, 0x3e, 0xb2,
	0xaf, 0xe0, 0x28, 0x36, 0x05, 0x11, 0x5b, 0x74, 0xbd, 0xfc, 0xf8, 0xe7, 0x4e, 0x2c, 0x19, 0xd0,
	0x51, 0x95, 0x59, 0xba, 0x25, 0x0e, 0xaa, 0x6b, 0x8e, 0x26, 0xd2, 0x65, 0xee, 0xa2, 0x82, 0x52,
	0xb0, 0x7e, 0x74, 0x06, 0x55, 0x56, 0xc4, 0x24, 0xe5, 0x1b, 0xab, 0x16, 0x87, 0xaa, 0x2f, 0x60,
	0x80, 0x93, 0x36, 0xc2, 0x9a, 0x6b, 0xe7, 0x1b, 0xfd, 0xde, 0xdb, 0x42, 0xa3, 0x36, 0x50, 0x00,
	0xb2, 0xf7, 0xab, 0x7b, 0x20, 0x46, 0x87, 0xd7, 0xd3, 0xd1, 0xaf, 0x22, 0x8d, 0x18, 0xef, 0x80,
	0xfd, 0xa8, 0x09, 0xcb, 0xb9, 0xfa, 0xde, 0x4b, 0x7d, 0x9a, 0x18, 0x2e, 0x35, 0x38, 0x0e, 0x17,
	0xff, 0x50, 0x57, 0x12, 0x3a, 0x90, 0x28, 0x26, 0x7b, 0x8d, 0xa9, 0x85, 0x18, 0x2c, 0xe9, 0x74,
	0x4b, 0x98, 0xe5, 0xa6, 0x0d, 0xc9, 0x69, 0x57, 0xa8, 0x03, 0x86, 0xb3, 0xdf, 0xd0, 0xff, 0xb7,
	0x19, 0x3c, 0x82, 0xe0, 0x34, 0x6c, 0xae, 0x5e, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xc5,
	0x56, 0x18, 0x98, 0x01, 0x00, 0x00,
}