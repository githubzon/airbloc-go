// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rpc/v1/types.proto

package v1

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

type Result struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_f56f1664f486e671, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

type Results struct {
	Results              []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Results) Reset()         { *m = Results{} }
func (m *Results) String() string { return proto.CompactTextString(m) }
func (*Results) ProtoMessage()    {}
func (*Results) Descriptor() ([]byte, []int) {
	return fileDescriptor_f56f1664f486e671, []int{1}
}

func (m *Results) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Results.Unmarshal(m, b)
}
func (m *Results) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Results.Marshal(b, m, deterministic)
}
func (m *Results) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Results.Merge(m, src)
}
func (m *Results) XXX_Size() int {
	return xxx_messageInfo_Results.Size(m)
}
func (m *Results) XXX_DiscardUnknown() {
	xxx_messageInfo_Results.DiscardUnknown(m)
}

var xxx_messageInfo_Results proto.InternalMessageInfo

func (m *Results) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type Address struct {
	ChainId              int32    `protobuf:"varint,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
	Address              []byte   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_f56f1664f486e671, []int{2}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetChainId() int32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Address) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type Hash struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hash) Reset()         { *m = Hash{} }
func (m *Hash) String() string { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()    {}
func (*Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_f56f1664f486e671, []int{3}
}

func (m *Hash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hash.Unmarshal(m, b)
}
func (m *Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hash.Marshal(b, m, deterministic)
}
func (m *Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hash.Merge(m, src)
}
func (m *Hash) XXX_Size() int {
	return xxx_messageInfo_Hash.Size(m)
}
func (m *Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_Hash.DiscardUnknown(m)
}

var xxx_messageInfo_Hash proto.InternalMessageInfo

func (m *Hash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func init() {
	proto.RegisterType((*Result)(nil), "airbloc.rpc.v1.Result")
	proto.RegisterType((*Results)(nil), "airbloc.rpc.v1.Results")
	proto.RegisterType((*Address)(nil), "airbloc.rpc.v1.Address")
	proto.RegisterType((*Hash)(nil), "airbloc.rpc.v1.Hash")
}

func init() { proto.RegisterFile("proto/rpc/v1/types.proto", fileDescriptor_f56f1664f486e671) }

var fileDescriptor_f56f1664f486e671 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x03,
	0x0b, 0x09, 0xf1, 0x25, 0x66, 0x16, 0x25, 0xe5, 0xe4, 0x27, 0xeb, 0x15, 0x15, 0x24, 0xeb, 0x95,
	0x19, 0x2a, 0x71, 0x70, 0xb1, 0x05, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x28, 0x59, 0x73, 0xb1, 0x43,
	0x58, 0xc5, 0x42, 0x06, 0x5c, 0xec, 0x45, 0x10, 0xa6, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x91,
	0x98, 0x1e, 0xaa, 0x36, 0x3d, 0x88, 0xca, 0x20, 0x98, 0x32, 0x25, 0x5b, 0x2e, 0x76, 0xc7, 0x94,
	0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0x8c, 0xc4, 0xcc, 0x3c, 0xcf, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x18, 0x17, 0x24, 0x93, 0x08, 0x51, 0x24, 0xc1, 0xa4,
	0xc0, 0xa8, 0xc1, 0x13, 0x04, 0xe3, 0x2a, 0x49, 0x71, 0xb1, 0x78, 0x24, 0x16, 0x67, 0x08, 0x09,
	0x71, 0xb1, 0x64, 0x24, 0x16, 0x67, 0x80, 0x35, 0xf2, 0x04, 0x81, 0xd9, 0x4e, 0x3a, 0x51, 0x5a,
	0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x50, 0x77, 0xc0, 0x68, 0xdd,
	0xf4, 0x7c, 0x7d, 0x64, 0xbf, 0x26, 0xb1, 0x81, 0x79, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xff, 0x7f, 0xeb, 0x8a, 0x02, 0x01, 0x00, 0x00,
}
