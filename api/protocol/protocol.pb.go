// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: api/protocol/protocol.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Operation int32

const (
	Operation_UNSPECIFIED Operation = 0
	// 连接
	Operation_Connect Operation = 1
	// 连接确认
	Operation_ConnectAck Operation = 2
	// 断开连接
	Operation_Disconnect Operation = 3
	// ping
	Operation_PingReq Operation = 4
	// pong
	Operation_PingResp Operation = 5
	// 发布消息
	Operation_Publish Operation = 6
	// 发布消息确认
	Operation_PublishAck Operation = 7
	// 订阅主题
	Operation_Subscribe Operation = 8
	// 订阅主题确认
	Operation_SubscribeAck Operation = 9
	// 取消订阅主题
	Operation_Unsubscribe Operation = 10
	// 取消订阅主题确认
	Operation_UnsubscribeAck Operation = 11
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "Connect",
		2:  "ConnectAck",
		3:  "Disconnect",
		4:  "PingReq",
		5:  "PingResp",
		6:  "Publish",
		7:  "PublishAck",
		8:  "Subscribe",
		9:  "SubscribeAck",
		10: "Unsubscribe",
		11: "UnsubscribeAck",
	}
	Operation_value = map[string]int32{
		"UNSPECIFIED":    0,
		"Connect":        1,
		"ConnectAck":     2,
		"Disconnect":     3,
		"PingReq":        4,
		"PingResp":       5,
		"Publish":        6,
		"PublishAck":     7,
		"Subscribe":      8,
		"SubscribeAck":   9,
		"Unsubscribe":    10,
		"UnsubscribeAck": 11,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_api_protocol_protocol_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_api_protocol_protocol_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{0}
}

type Proto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 消息指令
	Operation Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=realtimex.protocol.Operation" json:"operation,omitempty"`
	// 消息主体
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Proto) Reset() {
	*x = Proto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocol_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proto) ProtoMessage() {}

func (x *Proto) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocol_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proto.ProtoReflect.Descriptor instead.
func (*Proto) Descriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *Proto) GetOperation() Operation {
	if x != nil {
		return x.Operation
	}
	return Operation_UNSPECIFIED
}

func (x *Proto) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type ConnectPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 客户端ID
	ClientId string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	// 协议版本
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ConnectPacket) Reset() {
	*x = ConnectPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocol_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectPacket) ProtoMessage() {}

func (x *ConnectPacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocol_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectPacket.ProtoReflect.Descriptor instead.
func (*ConnectPacket) Descriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectPacket) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ConnectPacket) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ConnectAckPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectAckPacket) Reset() {
	*x = ConnectAckPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocol_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectAckPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectAckPacket) ProtoMessage() {}

func (x *ConnectAckPacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocol_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectAckPacket.ProtoReflect.Descriptor instead.
func (*ConnectAckPacket) Descriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{2}
}

type DisconnectPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisconnectPacket) Reset() {
	*x = DisconnectPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocol_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectPacket) ProtoMessage() {}

func (x *DisconnectPacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocol_protocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectPacket.ProtoReflect.Descriptor instead.
func (*DisconnectPacket) Descriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{3}
}

type DisconnectAckPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisconnectAckPacket) Reset() {
	*x = DisconnectAckPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocol_protocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectAckPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectAckPacket) ProtoMessage() {}

func (x *DisconnectAckPacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocol_protocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectAckPacket.ProtoReflect.Descriptor instead.
func (*DisconnectAckPacket) Descriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{4}
}

type PingReqPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingReqPacket) Reset() {
	*x = PingReqPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocol_protocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReqPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReqPacket) ProtoMessage() {}

func (x *PingReqPacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocol_protocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReqPacket.ProtoReflect.Descriptor instead.
func (*PingReqPacket) Descriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{5}
}

type PingRespPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRespPacket) Reset() {
	*x = PingRespPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocol_protocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRespPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRespPacket) ProtoMessage() {}

func (x *PingRespPacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocol_protocol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRespPacket.ProtoReflect.Descriptor instead.
func (*PingRespPacket) Descriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{6}
}

type PublishPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// 主题
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	// 消息ID
	MsgId uint64 `protobuf:"varint,3,opt,name=msgId,proto3" json:"msgId,omitempty"`
	// 消息内容
	Payload []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PublishPacket) Reset() {
	*x = PublishPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocol_protocol_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishPacket) ProtoMessage() {}

func (x *PublishPacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocol_protocol_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishPacket.ProtoReflect.Descriptor instead.
func (*PublishPacket) Descriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{7}
}

func (x *PublishPacket) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *PublishPacket) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PublishPacket) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *PublishPacket) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type PublishAckPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 消息ID
	MsgId uint64 `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
}

func (x *PublishAckPacket) Reset() {
	*x = PublishAckPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocol_protocol_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishAckPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishAckPacket) ProtoMessage() {}

func (x *PublishAckPacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocol_protocol_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishAckPacket.ProtoReflect.Descriptor instead.
func (*PublishAckPacket) Descriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{8}
}

func (x *PublishAckPacket) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

type SubscribePacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 订阅主题
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// 消息ID
	MsgId uint64 `protobuf:"varint,2,opt,name=msgId,proto3" json:"msgId,omitempty"`
}

func (x *SubscribePacket) Reset() {
	*x = SubscribePacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocol_protocol_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribePacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribePacket) ProtoMessage() {}

func (x *SubscribePacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocol_protocol_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribePacket.ProtoReflect.Descriptor instead.
func (*SubscribePacket) Descriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{9}
}

func (x *SubscribePacket) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *SubscribePacket) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

type SubscribeAckPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 消息ID
	MsgId uint64 `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
}

func (x *SubscribeAckPacket) Reset() {
	*x = SubscribeAckPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocol_protocol_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeAckPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeAckPacket) ProtoMessage() {}

func (x *SubscribeAckPacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocol_protocol_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeAckPacket.ProtoReflect.Descriptor instead.
func (*SubscribeAckPacket) Descriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{10}
}

func (x *SubscribeAckPacket) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

type UnsubscribePacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 取消订阅主题
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// 消息ID
	MsgId uint64 `protobuf:"varint,2,opt,name=msgId,proto3" json:"msgId,omitempty"`
}

func (x *UnsubscribePacket) Reset() {
	*x = UnsubscribePacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocol_protocol_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribePacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribePacket) ProtoMessage() {}

func (x *UnsubscribePacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocol_protocol_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribePacket.ProtoReflect.Descriptor instead.
func (*UnsubscribePacket) Descriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{11}
}

func (x *UnsubscribePacket) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *UnsubscribePacket) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

type UnsubscribeAckPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 消息ID
	MsgId uint64 `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
}

func (x *UnsubscribeAckPacket) Reset() {
	*x = UnsubscribeAckPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protocol_protocol_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeAckPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeAckPacket) ProtoMessage() {}

func (x *UnsubscribeAckPacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_protocol_protocol_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeAckPacket.ProtoReflect.Descriptor instead.
func (*UnsubscribeAckPacket) Descriptor() ([]byte, []int) {
	return file_api_protocol_protocol_proto_rawDescGZIP(), []int{12}
}

func (x *UnsubscribeAckPacket) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

var File_api_protocol_protocol_proto protoreflect.FileDescriptor

var file_api_protocol_protocol_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72,
	0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x22, 0x58, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e,
	0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x45, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x41, 0x63, 0x6b,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x41, 0x63, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x22, 0x67, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73,
	0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x28, 0x0a,
	0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x41, 0x63, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73, 0x67,
	0x49, 0x64, 0x22, 0x3f, 0x0a, 0x11, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73,
	0x67, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x14, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x41, 0x63, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x2a, 0xc7, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0x01, 0x12, 0x0e, 0x0a,
	0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x41, 0x63, 0x6b, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x41, 0x63, 0x6b, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x41, 0x63, 0x6b, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x6e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63, 0x6b, 0x10, 0x0b, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x6f, 0x75, 0x74, 0x68,
	0x7a, 0x7a, 0x7a, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x78, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protocol_protocol_proto_rawDescOnce sync.Once
	file_api_protocol_protocol_proto_rawDescData = file_api_protocol_protocol_proto_rawDesc
)

func file_api_protocol_protocol_proto_rawDescGZIP() []byte {
	file_api_protocol_protocol_proto_rawDescOnce.Do(func() {
		file_api_protocol_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protocol_protocol_proto_rawDescData)
	})
	return file_api_protocol_protocol_proto_rawDescData
}

var file_api_protocol_protocol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_protocol_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_protocol_protocol_proto_goTypes = []interface{}{
	(Operation)(0),               // 0: realtimex.protocol.Operation
	(*Proto)(nil),                // 1: realtimex.protocol.Proto
	(*ConnectPacket)(nil),        // 2: realtimex.protocol.ConnectPacket
	(*ConnectAckPacket)(nil),     // 3: realtimex.protocol.ConnectAckPacket
	(*DisconnectPacket)(nil),     // 4: realtimex.protocol.DisconnectPacket
	(*DisconnectAckPacket)(nil),  // 5: realtimex.protocol.DisconnectAckPacket
	(*PingReqPacket)(nil),        // 6: realtimex.protocol.PingReqPacket
	(*PingRespPacket)(nil),       // 7: realtimex.protocol.PingRespPacket
	(*PublishPacket)(nil),        // 8: realtimex.protocol.PublishPacket
	(*PublishAckPacket)(nil),     // 9: realtimex.protocol.PublishAckPacket
	(*SubscribePacket)(nil),      // 10: realtimex.protocol.SubscribePacket
	(*SubscribeAckPacket)(nil),   // 11: realtimex.protocol.SubscribeAckPacket
	(*UnsubscribePacket)(nil),    // 12: realtimex.protocol.UnsubscribePacket
	(*UnsubscribeAckPacket)(nil), // 13: realtimex.protocol.UnsubscribeAckPacket
}
var file_api_protocol_protocol_proto_depIdxs = []int32{
	0, // 0: realtimex.protocol.Proto.operation:type_name -> realtimex.protocol.Operation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_protocol_protocol_proto_init() }
func file_api_protocol_protocol_proto_init() {
	if File_api_protocol_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protocol_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protocol_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protocol_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectAckPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protocol_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protocol_protocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectAckPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protocol_protocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReqPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protocol_protocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRespPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protocol_protocol_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protocol_protocol_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishAckPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protocol_protocol_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribePacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protocol_protocol_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeAckPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protocol_protocol_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribePacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protocol_protocol_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeAckPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_protocol_protocol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_protocol_protocol_proto_goTypes,
		DependencyIndexes: file_api_protocol_protocol_proto_depIdxs,
		EnumInfos:         file_api_protocol_protocol_proto_enumTypes,
		MessageInfos:      file_api_protocol_protocol_proto_msgTypes,
	}.Build()
	File_api_protocol_protocol_proto = out.File
	file_api_protocol_protocol_proto_rawDesc = nil
	file_api_protocol_protocol_proto_goTypes = nil
	file_api_protocol_protocol_proto_depIdxs = nil
}
