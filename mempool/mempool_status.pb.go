// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mempool_status.proto

package mempool

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

type MempoolAddTransactionStatus int32

const (
	// Transaction was sent to Mempool
	MempoolAddTransactionStatus_Valid MempoolAddTransactionStatus = 0
	// The sender does not have enough balance for the transaction.
	MempoolAddTransactionStatus_InsufficientBalance MempoolAddTransactionStatus = 1
	// Sequence number is old, etc.
	MempoolAddTransactionStatus_InvalidSeqNumber MempoolAddTransactionStatus = 2
	// Mempool is full (reached max global capacity)
	MempoolAddTransactionStatus_MempoolIsFull MempoolAddTransactionStatus = 3
	// Account reached max capacity per account
	MempoolAddTransactionStatus_TooManyTransactions MempoolAddTransactionStatus = 4
	// Invalid update. Only gas price increase is allowed
	MempoolAddTransactionStatus_InvalidUpdate MempoolAddTransactionStatus = 5
)

var MempoolAddTransactionStatus_name = map[int32]string{
	0: "Valid",
	1: "InsufficientBalance",
	2: "InvalidSeqNumber",
	3: "MempoolIsFull",
	4: "TooManyTransactions",
	5: "InvalidUpdate",
}

var MempoolAddTransactionStatus_value = map[string]int32{
	"Valid":               0,
	"InsufficientBalance": 1,
	"InvalidSeqNumber":    2,
	"MempoolIsFull":       3,
	"TooManyTransactions": 4,
	"InvalidUpdate":       5,
}

func (x MempoolAddTransactionStatus) String() string {
	return proto.EnumName(MempoolAddTransactionStatus_name, int32(x))
}

func (MempoolAddTransactionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cad4a86f8a5465be, []int{0}
}

func init() {
	proto.RegisterEnum("mempool.MempoolAddTransactionStatus", MempoolAddTransactionStatus_name, MempoolAddTransactionStatus_value)
}

func init() { proto.RegisterFile("mempool_status.proto", fileDescriptor_cad4a86f8a5465be) }

var fileDescriptor_cad4a86f8a5465be = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xbf, 0x4a, 0x04, 0x31,
	0x10, 0x87, 0x3d, 0xf5, 0x14, 0x03, 0x42, 0x8c, 0x07, 0x16, 0x3e, 0x80, 0x68, 0x71, 0x5b, 0x58,
	0x59, 0x7a, 0x85, 0xb0, 0xc5, 0xd9, 0xdc, 0x69, 0x61, 0x73, 0x4c, 0x92, 0x39, 0x37, 0x30, 0xc9,
	0xc4, 0xfc, 0x11, 0x7d, 0x11, 0x9f, 0x57, 0x76, 0x4d, 0x61, 0x37, 0x0c, 0x1f, 0xdf, 0x8f, 0x4f,
	0x2c, 0x3c, 0xfa, 0xc8, 0x4c, 0xbb, 0x5c, 0xa0, 0xd4, 0xbc, 0x8c, 0x89, 0x0b, 0xab, 0xd3, 0xf6,
	0xbd, 0xfb, 0x99, 0x89, 0xeb, 0xf5, 0xdf, 0xfd, 0x68, 0xed, 0x36, 0x41, 0xc8, 0x60, 0x8a, 0xe3,
	0xb0, 0x99, 0x70, 0x75, 0x26, 0xe6, 0xaf, 0x40, 0xce, 0xca, 0x03, 0x75, 0x25, 0x2e, 0xfb, 0x90,
	0xeb, 0x7e, 0xef, 0x8c, 0xc3, 0x50, 0x56, 0x40, 0x10, 0x0c, 0xca, 0x99, 0x5a, 0x08, 0xd9, 0x87,
	0xcf, 0x91, 0xda, 0xe0, 0xc7, 0x73, 0xf5, 0x1a, 0x93, 0x3c, 0x54, 0x17, 0xe2, 0xbc, 0x89, 0xfb,
	0xfc, 0x54, 0x89, 0xe4, 0xd1, 0x68, 0xd8, 0x32, 0xaf, 0x21, 0x7c, 0xff, 0x1b, 0xca, 0xf2, 0x78,
	0x64, 0x9b, 0xe1, 0x25, 0x5a, 0x28, 0x28, 0xe7, 0xab, 0xdb, 0xb7, 0x9b, 0x77, 0x57, 0x86, 0xaa,
	0x97, 0x86, 0x7d, 0x17, 0x07, 0x72, 0xf1, 0xa1, 0x23, 0xa7, 0x13, 0xec, 0xf0, 0x0b, 0x7c, 0x24,
	0xec, 0x5a, 0x83, 0x3e, 0x99, 0x9a, 0xee, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x7e, 0x66,
	0xd4, 0xeb, 0x00, 0x00, 0x00,
}