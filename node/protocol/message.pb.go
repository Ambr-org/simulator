// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package protocol

/*
protoc --go_out=. message.proto
*/

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

type Payload struct {
	// address base58 encoded
	Creator              string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Previous             string   `protobuf:"bytes,2,opt,name=previous,proto3" json:"previous,omitempty"`
	Balance              int64    `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_10fb5d1ba8601988, []int{0}
}
func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (dst *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(dst, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Payload) GetPrevious() string {
	if m != nil {
		return m.Previous
	}
	return ""
}

func (m *Payload) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Payload) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type SendUnit struct {
	Payload              *Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendUnit) Reset()         { *m = SendUnit{} }
func (m *SendUnit) String() string { return proto.CompactTextString(m) }
func (*SendUnit) ProtoMessage()    {}
func (*SendUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_10fb5d1ba8601988, []int{1}
}
func (m *SendUnit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendUnit.Unmarshal(m, b)
}
func (m *SendUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendUnit.Marshal(b, m, deterministic)
}
func (dst *SendUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendUnit.Merge(dst, src)
}
func (m *SendUnit) XXX_Size() int {
	return xxx_messageInfo_SendUnit.Size(m)
}
func (m *SendUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_SendUnit.DiscardUnknown(m)
}

var xxx_messageInfo_SendUnit proto.InternalMessageInfo

func (m *SendUnit) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

type RecvUnit struct {
	Payload              *Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Other                string   `protobuf:"bytes,2,opt,name=other,proto3" json:"other,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecvUnit) Reset()         { *m = RecvUnit{} }
func (m *RecvUnit) String() string { return proto.CompactTextString(m) }
func (*RecvUnit) ProtoMessage()    {}
func (*RecvUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_10fb5d1ba8601988, []int{2}
}
func (m *RecvUnit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecvUnit.Unmarshal(m, b)
}
func (m *RecvUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecvUnit.Marshal(b, m, deterministic)
}
func (dst *RecvUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecvUnit.Merge(dst, src)
}
func (m *RecvUnit) XXX_Size() int {
	return xxx_messageInfo_RecvUnit.Size(m)
}
func (m *RecvUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_RecvUnit.DiscardUnknown(m)
}

var xxx_messageInfo_RecvUnit proto.InternalMessageInfo

func (m *RecvUnit) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *RecvUnit) GetOther() string {
	if m != nil {
		return m.Other
	}
	return ""
}

// for test usage
type VoteRequest struct {
	ConflictUnit1        string   `protobuf:"bytes,1,opt,name=conflictUnit1,proto3" json:"conflictUnit1,omitempty"`
	ConfilictUnit2       string   `protobuf:"bytes,2,opt,name=confilictUnit2,proto3" json:"confilictUnit2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteRequest) Reset()         { *m = VoteRequest{} }
func (m *VoteRequest) String() string { return proto.CompactTextString(m) }
func (*VoteRequest) ProtoMessage()    {}
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_10fb5d1ba8601988, []int{3}
}
func (m *VoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteRequest.Unmarshal(m, b)
}
func (m *VoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteRequest.Marshal(b, m, deterministic)
}
func (dst *VoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteRequest.Merge(dst, src)
}
func (m *VoteRequest) XXX_Size() int {
	return xxx_messageInfo_VoteRequest.Size(m)
}
func (m *VoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VoteRequest proto.InternalMessageInfo

func (m *VoteRequest) GetConflictUnit1() string {
	if m != nil {
		return m.ConflictUnit1
	}
	return ""
}

func (m *VoteRequest) GetConfilictUnit2() string {
	if m != nil {
		return m.ConfilictUnit2
	}
	return ""
}

type VoteResponse struct {
	ConflictUnit1        string   `protobuf:"bytes,1,opt,name=conflictUnit1,proto3" json:"conflictUnit1,omitempty"`
	ConfilictUnit2       string   `protobuf:"bytes,2,opt,name=confilictUnit2,proto3" json:"confilictUnit2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteResponse) Reset()         { *m = VoteResponse{} }
func (m *VoteResponse) String() string { return proto.CompactTextString(m) }
func (*VoteResponse) ProtoMessage()    {}
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_10fb5d1ba8601988, []int{4}
}
func (m *VoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteResponse.Unmarshal(m, b)
}
func (m *VoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteResponse.Marshal(b, m, deterministic)
}
func (dst *VoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteResponse.Merge(dst, src)
}
func (m *VoteResponse) XXX_Size() int {
	return xxx_messageInfo_VoteResponse.Size(m)
}
func (m *VoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VoteResponse proto.InternalMessageInfo

func (m *VoteResponse) GetConflictUnit1() string {
	if m != nil {
		return m.ConflictUnit1
	}
	return ""
}

func (m *VoteResponse) GetConfilictUnit2() string {
	if m != nil {
		return m.ConfilictUnit2
	}
	return ""
}

type ReplicationRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplicationRequest) Reset()         { *m = ReplicationRequest{} }
func (m *ReplicationRequest) String() string { return proto.CompactTextString(m) }
func (*ReplicationRequest) ProtoMessage()    {}
func (*ReplicationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_10fb5d1ba8601988, []int{5}
}
func (m *ReplicationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplicationRequest.Unmarshal(m, b)
}
func (m *ReplicationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplicationRequest.Marshal(b, m, deterministic)
}
func (dst *ReplicationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplicationRequest.Merge(dst, src)
}
func (m *ReplicationRequest) XXX_Size() int {
	return xxx_messageInfo_ReplicationRequest.Size(m)
}
func (m *ReplicationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplicationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReplicationRequest proto.InternalMessageInfo

type ReplicationResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplicationResponse) Reset()         { *m = ReplicationResponse{} }
func (m *ReplicationResponse) String() string { return proto.CompactTextString(m) }
func (*ReplicationResponse) ProtoMessage()    {}
func (*ReplicationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_10fb5d1ba8601988, []int{6}
}
func (m *ReplicationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplicationResponse.Unmarshal(m, b)
}
func (m *ReplicationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplicationResponse.Marshal(b, m, deterministic)
}
func (dst *ReplicationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplicationResponse.Merge(dst, src)
}
func (m *ReplicationResponse) XXX_Size() int {
	return xxx_messageInfo_ReplicationResponse.Size(m)
}
func (m *ReplicationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplicationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReplicationResponse proto.InternalMessageInfo

type HeartbeatRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatRequest) Reset()         { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string { return proto.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()    {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_10fb5d1ba8601988, []int{7}
}
func (m *HeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatRequest.Unmarshal(m, b)
}
func (m *HeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatRequest.Marshal(b, m, deterministic)
}
func (dst *HeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatRequest.Merge(dst, src)
}
func (m *HeartbeatRequest) XXX_Size() int {
	return xxx_messageInfo_HeartbeatRequest.Size(m)
}
func (m *HeartbeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatRequest proto.InternalMessageInfo

type HeartbeatResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatResponse) Reset()         { *m = HeartbeatResponse{} }
func (m *HeartbeatResponse) String() string { return proto.CompactTextString(m) }
func (*HeartbeatResponse) ProtoMessage()    {}
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_10fb5d1ba8601988, []int{8}
}
func (m *HeartbeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatResponse.Unmarshal(m, b)
}
func (m *HeartbeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatResponse.Marshal(b, m, deterministic)
}
func (dst *HeartbeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatResponse.Merge(dst, src)
}
func (m *HeartbeatResponse) XXX_Size() int {
	return xxx_messageInfo_HeartbeatResponse.Size(m)
}
func (m *HeartbeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatResponse proto.InternalMessageInfo

// for test usage
type TestMessage struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Length               int32    `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	Cnt                  int32    `protobuf:"varint,3,opt,name=cnt,proto3" json:"cnt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestMessage) Reset()         { *m = TestMessage{} }
func (m *TestMessage) String() string { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()    {}
func (*TestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_10fb5d1ba8601988, []int{9}
}
func (m *TestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMessage.Unmarshal(m, b)
}
func (m *TestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMessage.Marshal(b, m, deterministic)
}
func (dst *TestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessage.Merge(dst, src)
}
func (m *TestMessage) XXX_Size() int {
	return xxx_messageInfo_TestMessage.Size(m)
}
func (m *TestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessage proto.InternalMessageInfo

func (m *TestMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TestMessage) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *TestMessage) GetCnt() int32 {
	if m != nil {
		return m.Cnt
	}
	return 0
}

func init() {
	proto.RegisterType((*Payload)(nil), "protocol.Payload")
	proto.RegisterType((*SendUnit)(nil), "protocol.SendUnit")
	proto.RegisterType((*RecvUnit)(nil), "protocol.RecvUnit")
	proto.RegisterType((*VoteRequest)(nil), "protocol.VoteRequest")
	proto.RegisterType((*VoteResponse)(nil), "protocol.VoteResponse")
	proto.RegisterType((*ReplicationRequest)(nil), "protocol.ReplicationRequest")
	proto.RegisterType((*ReplicationResponse)(nil), "protocol.ReplicationResponse")
	proto.RegisterType((*HeartbeatRequest)(nil), "protocol.HeartbeatRequest")
	proto.RegisterType((*HeartbeatResponse)(nil), "protocol.HeartbeatResponse")
	proto.RegisterType((*TestMessage)(nil), "protocol.TestMessage")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_10fb5d1ba8601988) }

var fileDescriptor_message_10fb5d1ba8601988 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x4f, 0x2a, 0x31,
	0x14, 0xcd, 0x3c, 0x1e, 0x30, 0x5c, 0xe0, 0x05, 0x0a, 0xcf, 0x4c, 0x8c, 0x0b, 0xd2, 0x18, 0x43,
	0x62, 0x42, 0x22, 0x2e, 0xfc, 0x0b, 0x6e, 0x48, 0xb4, 0x7e, 0x6c, 0x74, 0x53, 0xca, 0x15, 0x9a,
	0x8c, 0xed, 0xd8, 0x5e, 0x30, 0xfe, 0x7b, 0x33, 0x33, 0xad, 0x8a, 0x3b, 0x13, 0x57, 0xed, 0x39,
	0xbd, 0x1f, 0xe7, 0xf6, 0x1e, 0xe8, 0x3f, 0xa3, 0xf7, 0x72, 0x8d, 0xb3, 0xc2, 0x59, 0xb2, 0x2c,
	0xad, 0x0e, 0x65, 0x73, 0xfe, 0x0a, 0xed, 0x2b, 0xf9, 0x96, 0x5b, 0xb9, 0x62, 0x19, 0xb4, 0x95,
	0x43, 0x49, 0xd6, 0x65, 0xc9, 0x24, 0x99, 0x76, 0x44, 0x84, 0xec, 0x10, 0xd2, 0xc2, 0xe1, 0x4e,
	0xdb, 0xad, 0xcf, 0xfe, 0x54, 0x4f, 0x1f, 0xb8, 0xcc, 0x5a, 0xca, 0x5c, 0x1a, 0x85, 0x59, 0x63,
	0x92, 0x4c, 0x1b, 0x22, 0x42, 0x76, 0x04, 0x1d, 0xaf, 0xd7, 0x46, 0xd2, 0xd6, 0x61, 0xf6, 0x77,
	0x92, 0x4c, 0x7b, 0xe2, 0x93, 0xe0, 0x17, 0x90, 0xde, 0xa0, 0x59, 0xdd, 0x19, 0x4d, 0xec, 0x14,
	0xda, 0x45, 0x2d, 0xa2, 0xea, 0xdc, 0x9d, 0x0f, 0x67, 0x51, 0xe0, 0x2c, 0xa8, 0x13, 0x31, 0x82,
	0x2f, 0x20, 0x15, 0xa8, 0x76, 0x3f, 0x4e, 0x64, 0x63, 0x68, 0x5a, 0xda, 0xa0, 0x0b, 0x23, 0xd4,
	0x80, 0x3f, 0x40, 0xf7, 0xde, 0x12, 0x0a, 0x7c, 0xd9, 0xa2, 0x27, 0x76, 0x0c, 0x7d, 0x65, 0xcd,
	0x53, 0xae, 0x15, 0x95, 0x1d, 0xce, 0xc2, 0x57, 0xec, 0x93, 0xec, 0x04, 0xfe, 0x95, 0x84, 0x8e,
	0xcc, 0x3c, 0xd4, 0xfc, 0xc6, 0xf2, 0x47, 0xe8, 0xd5, 0xc5, 0x7d, 0x61, 0x8d, 0xc7, 0x5f, 0xae,
	0x3e, 0x06, 0x26, 0xb0, 0xc8, 0xb5, 0x92, 0xa4, 0xad, 0x09, 0x13, 0xf0, 0xff, 0x30, 0xda, 0x63,
	0xeb, 0xd6, 0x9c, 0xc1, 0xe0, 0x12, 0xa5, 0xa3, 0x25, 0x4a, 0x8a, 0xa1, 0x23, 0x18, 0x7e, 0xe1,
	0x42, 0xe0, 0x35, 0x74, 0x6f, 0xd1, 0xd3, 0xa2, 0x36, 0x4c, 0xb9, 0xdf, 0xe0, 0x9d, 0xe8, 0x8a,
	0x00, 0xd9, 0x01, 0xb4, 0x72, 0x34, 0x6b, 0xda, 0x54, 0xf2, 0x9a, 0x22, 0x20, 0x36, 0x80, 0x86,
	0x32, 0x54, 0xb9, 0xa1, 0x29, 0xca, 0xeb, 0xb2, 0x55, 0x2d, 0xe5, 0xfc, 0x3d, 0x00, 0x00, 0xff,
	0xff, 0xce, 0xae, 0x6f, 0x13, 0x86, 0x02, 0x00, 0x00,
}
