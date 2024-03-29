// sensory.api.assistant

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: v1/assistant/assistant.proto

package assistant

import (
	_ "github.com/Sensory-Cloud/go-sdk/pkg/generated/common"
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

type ChatRole int32

const (
	ChatRole_SYSTEM    ChatRole = 0
	ChatRole_USER      ChatRole = 1
	ChatRole_ASSISTANT ChatRole = 2
)

// Enum value maps for ChatRole.
var (
	ChatRole_name = map[int32]string{
		0: "SYSTEM",
		1: "USER",
		2: "ASSISTANT",
	}
	ChatRole_value = map[string]int32{
		"SYSTEM":    0,
		"USER":      1,
		"ASSISTANT": 2,
	}
)

func (x ChatRole) Enum() *ChatRole {
	p := new(ChatRole)
	*p = x
	return p
}

func (x ChatRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatRole) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_assistant_assistant_proto_enumTypes[0].Descriptor()
}

func (ChatRole) Type() protoreflect.EnumType {
	return &file_v1_assistant_assistant_proto_enumTypes[0]
}

func (x ChatRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatRole.Descriptor instead.
func (ChatRole) EnumDescriptor() ([]byte, []int) {
	return file_v1_assistant_assistant_proto_rawDescGZIP(), []int{0}
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role    ChatRole `protobuf:"varint,1,opt,name=role,proto3,enum=sensory.api.v1.assistant.ChatRole" json:"role,omitempty"`
	Content string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_assistant_assistant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_v1_assistant_assistant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_v1_assistant_assistant_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetRole() ChatRole {
	if x != nil {
		return x.Role
	}
	return ChatRole_SYSTEM
}

func (x *ChatMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type TextChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelName string         `protobuf:"bytes,1,opt,name=modelName,proto3" json:"modelName,omitempty"`
	Messages  []*ChatMessage `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *TextChatRequest) Reset() {
	*x = TextChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_assistant_assistant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextChatRequest) ProtoMessage() {}

func (x *TextChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_assistant_assistant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextChatRequest.ProtoReflect.Descriptor instead.
func (*TextChatRequest) Descriptor() ([]byte, []int) {
	return file_v1_assistant_assistant_proto_rawDescGZIP(), []int{1}
}

func (x *TextChatRequest) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *TextChatRequest) GetMessages() []*ChatMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

type TextChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *ChatMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TextChatResponse) Reset() {
	*x = TextChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_assistant_assistant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextChatResponse) ProtoMessage() {}

func (x *TextChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_assistant_assistant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextChatResponse.ProtoReflect.Descriptor instead.
func (*TextChatResponse) Descriptor() ([]byte, []int) {
	return file_v1_assistant_assistant_proto_rawDescGZIP(), []int{2}
}

func (x *TextChatResponse) GetMessage() *ChatMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_v1_assistant_assistant_proto protoreflect.FileDescriptor

var file_v1_assistant_assistant_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2f, 0x61,
	0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61,
	0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a,
	0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x72,
	0x0a, 0x0f, 0x54, 0x65, 0x78, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x41, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x53, 0x0a, 0x10, 0x54, 0x65, 0x78, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x2f, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x53, 0x53,
	0x49, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x10, 0x02, 0x32, 0x77, 0x0a, 0x10, 0x41, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x08,
	0x54, 0x65, 0x78, 0x74, 0x43, 0x68, 0x61, 0x74, 0x12, 0x29, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x54,
	0x65, 0x78, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x87, 0x01, 0x0a, 0x20, 0x61, 0x69, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x73,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x42, 0x21, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x41,
	0x70, 0x69, 0x56, 0x31, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x74, 0x69, 0x74,
	0x61, 0x6e, 0x2e, 0x67, 0x69, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_assistant_assistant_proto_rawDescOnce sync.Once
	file_v1_assistant_assistant_proto_rawDescData = file_v1_assistant_assistant_proto_rawDesc
)

func file_v1_assistant_assistant_proto_rawDescGZIP() []byte {
	file_v1_assistant_assistant_proto_rawDescOnce.Do(func() {
		file_v1_assistant_assistant_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_assistant_assistant_proto_rawDescData)
	})
	return file_v1_assistant_assistant_proto_rawDescData
}

var file_v1_assistant_assistant_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_assistant_assistant_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_assistant_assistant_proto_goTypes = []interface{}{
	(ChatRole)(0),            // 0: sensory.api.v1.assistant.ChatRole
	(*ChatMessage)(nil),      // 1: sensory.api.v1.assistant.ChatMessage
	(*TextChatRequest)(nil),  // 2: sensory.api.v1.assistant.TextChatRequest
	(*TextChatResponse)(nil), // 3: sensory.api.v1.assistant.TextChatResponse
}
var file_v1_assistant_assistant_proto_depIdxs = []int32{
	0, // 0: sensory.api.v1.assistant.ChatMessage.role:type_name -> sensory.api.v1.assistant.ChatRole
	1, // 1: sensory.api.v1.assistant.TextChatRequest.messages:type_name -> sensory.api.v1.assistant.ChatMessage
	1, // 2: sensory.api.v1.assistant.TextChatResponse.message:type_name -> sensory.api.v1.assistant.ChatMessage
	2, // 3: sensory.api.v1.assistant.AssistantService.TextChat:input_type -> sensory.api.v1.assistant.TextChatRequest
	3, // 4: sensory.api.v1.assistant.AssistantService.TextChat:output_type -> sensory.api.v1.assistant.TextChatResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_assistant_assistant_proto_init() }
func file_v1_assistant_assistant_proto_init() {
	if File_v1_assistant_assistant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_assistant_assistant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
		file_v1_assistant_assistant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextChatRequest); i {
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
		file_v1_assistant_assistant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextChatResponse); i {
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
			RawDescriptor: file_v1_assistant_assistant_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_assistant_assistant_proto_goTypes,
		DependencyIndexes: file_v1_assistant_assistant_proto_depIdxs,
		EnumInfos:         file_v1_assistant_assistant_proto_enumTypes,
		MessageInfos:      file_v1_assistant_assistant_proto_msgTypes,
	}.Build()
	File_v1_assistant_assistant_proto = out.File
	file_v1_assistant_assistant_proto_rawDesc = nil
	file_v1_assistant_assistant_proto_goTypes = nil
	file_v1_assistant_assistant_proto_depIdxs = nil
}
