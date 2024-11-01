// chittychat.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: proto/chittychat.proto

package chittychat

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

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	mi := &file_proto_chittychat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chittychat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_proto_chittychat_proto_rawDescGZIP(), []int{0}
}

func (x *JoinRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	mi := &file_proto_chittychat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chittychat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_proto_chittychat_proto_rawDescGZIP(), []int{1}
}

func (x *PublishRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *PublishRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	mi := &file_proto_chittychat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chittychat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_proto_chittychat_proto_rawDescGZIP(), []int{2}
}

func (x *PublishResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PublishResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type LeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *LeaveRequest) Reset() {
	*x = LeaveRequest{}
	mi := &file_proto_chittychat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRequest) ProtoMessage() {}

func (x *LeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chittychat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRequest.ProtoReflect.Descriptor instead.
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return file_proto_chittychat_proto_rawDescGZIP(), []int{3}
}

func (x *LeaveRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type LeaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *LeaveResponse) Reset() {
	*x = LeaveResponse{}
	mi := &file_proto_chittychat_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LeaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveResponse) ProtoMessage() {}

func (x *LeaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chittychat_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveResponse.ProtoReflect.Descriptor instead.
func (*LeaveResponse) Descriptor() ([]byte, []int) {
	return file_proto_chittychat_proto_rawDescGZIP(), []int{4}
}

func (x *LeaveResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LeaveResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type BroadcastMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId    string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Message     string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	LamportTime int64  `protobuf:"varint,3,opt,name=lamport_time,json=lamportTime,proto3" json:"lamport_time,omitempty"`
}

func (x *BroadcastMessage) Reset() {
	*x = BroadcastMessage{}
	mi := &file_proto_chittychat_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BroadcastMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastMessage) ProtoMessage() {}

func (x *BroadcastMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chittychat_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastMessage.ProtoReflect.Descriptor instead.
func (*BroadcastMessage) Descriptor() ([]byte, []int) {
	return file_proto_chittychat_proto_rawDescGZIP(), []int{5}
}

func (x *BroadcastMessage) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *BroadcastMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BroadcastMessage) GetLamportTime() int64 {
	if x != nil {
		return x.LamportTime
	}
	return 0
}

var File_proto_chittychat_proto protoreflect.FileDescriptor

var file_proto_chittychat_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79,
	0x63, 0x68, 0x61, 0x74, 0x22, 0x2a, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x47, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x41, 0x0a, 0x0f, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2b, 0x0a, 0x0c,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x0d, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x6c, 0x0a, 0x10, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xde, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x69,
	0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x12, 0x43, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x43,
	0x68, 0x61, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63,
	0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x49, 0x0a, 0x0e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x68, 0x69,
	0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x63, 0x68, 0x69,
	0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68,
	0x69, 0x74, 0x74, 0x79, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_chittychat_proto_rawDescOnce sync.Once
	file_proto_chittychat_proto_rawDescData = file_proto_chittychat_proto_rawDesc
)

func file_proto_chittychat_proto_rawDescGZIP() []byte {
	file_proto_chittychat_proto_rawDescOnce.Do(func() {
		file_proto_chittychat_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chittychat_proto_rawDescData)
	})
	return file_proto_chittychat_proto_rawDescData
}

var file_proto_chittychat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_chittychat_proto_goTypes = []any{
	(*JoinRequest)(nil),      // 0: chittychat.JoinRequest
	(*PublishRequest)(nil),   // 1: chittychat.PublishRequest
	(*PublishResponse)(nil),  // 2: chittychat.PublishResponse
	(*LeaveRequest)(nil),     // 3: chittychat.LeaveRequest
	(*LeaveResponse)(nil),    // 4: chittychat.LeaveResponse
	(*BroadcastMessage)(nil), // 5: chittychat.BroadcastMessage
}
var file_proto_chittychat_proto_depIdxs = []int32{
	0, // 0: chittychat.ChittyChat.JoinChat:input_type -> chittychat.JoinRequest
	1, // 1: chittychat.ChittyChat.PublishMessage:input_type -> chittychat.PublishRequest
	3, // 2: chittychat.ChittyChat.LeaveChat:input_type -> chittychat.LeaveRequest
	5, // 3: chittychat.ChittyChat.JoinChat:output_type -> chittychat.BroadcastMessage
	2, // 4: chittychat.ChittyChat.PublishMessage:output_type -> chittychat.PublishResponse
	4, // 5: chittychat.ChittyChat.LeaveChat:output_type -> chittychat.LeaveResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_chittychat_proto_init() }
func file_proto_chittychat_proto_init() {
	if File_proto_chittychat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_chittychat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chittychat_proto_goTypes,
		DependencyIndexes: file_proto_chittychat_proto_depIdxs,
		MessageInfos:      file_proto_chittychat_proto_msgTypes,
	}.Build()
	File_proto_chittychat_proto = out.File
	file_proto_chittychat_proto_rawDesc = nil
	file_proto_chittychat_proto_goTypes = nil
	file_proto_chittychat_proto_depIdxs = nil
}
