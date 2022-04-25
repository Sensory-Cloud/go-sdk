// sensory.api.event

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: v1/event/event.proto

package event

import (
	common "github.com/Sensory-Cloud/go-sdk/pkg/api/common"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request to create a new usage event
type PublishUsageEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of events to publish
	Events []*UsageEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *PublishUsageEventsRequest) Reset() {
	*x = PublishUsageEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_event_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishUsageEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishUsageEventsRequest) ProtoMessage() {}

func (x *PublishUsageEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_event_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishUsageEventsRequest.ProtoReflect.Descriptor instead.
func (*PublishUsageEventsRequest) Descriptor() ([]byte, []int) {
	return file_v1_event_event_proto_rawDescGZIP(), []int{0}
}

func (x *PublishUsageEventsRequest) GetEvents() []*UsageEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

type UsageEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The UTC timestamp
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The duration of the event in milliseconds
	Duration int64 `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// UUID representing the unique event
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// OAuth ClientID related to the event
	ClientId string `protobuf:"bytes,4,opt,name=clientId,proto3" json:"clientId,omitempty"`
	// Type of event to create
	Type common.UsageEventType `protobuf:"varint,5,opt,name=type,proto3,enum=sensory.api.common.UsageEventType" json:"type,omitempty"`
	// The specific route (endpoint) that was accessed
	Route string `protobuf:"bytes,6,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *UsageEvent) Reset() {
	*x = UsageEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_event_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsageEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageEvent) ProtoMessage() {}

func (x *UsageEvent) ProtoReflect() protoreflect.Message {
	mi := &file_v1_event_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageEvent.ProtoReflect.Descriptor instead.
func (*UsageEvent) Descriptor() ([]byte, []int) {
	return file_v1_event_event_proto_rawDescGZIP(), []int{1}
}

func (x *UsageEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *UsageEvent) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *UsageEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UsageEvent) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *UsageEvent) GetType() common.UsageEventType {
	if x != nil {
		return x.Type
	}
	return common.UsageEventType_AUTHENTICATION
}

func (x *UsageEvent) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

// Response to a publish events request
type PublishUsageEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishUsageEventsResponse) Reset() {
	*x = PublishUsageEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_event_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishUsageEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishUsageEventsResponse) ProtoMessage() {}

func (x *PublishUsageEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_event_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishUsageEventsResponse.ProtoReflect.Descriptor instead.
func (*PublishUsageEventsResponse) Descriptor() ([]byte, []int) {
	return file_v1_event_event_proto_rawDescGZIP(), []int{2}
}

var File_v1_event_event_proto protoreflect.FileDescriptor

var file_v1_event_event_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x19, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x55, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x9a, 0x02, 0x0a, 0x0a, 0x55, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x42, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0xb2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18,
	0x7f, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x03, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22,
	0x1c, 0x0a, 0x1a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x55, 0x73, 0x61, 0x67, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x89, 0x01,
	0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x79,
	0x0a, 0x12, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x55, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x2f, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x55, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x55, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x65, 0x0a, 0x17, 0x69, 0x6f, 0x2e,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x42, 0x16, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x41, 0x70, 0x69,
	0x56, 0x31, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x79, 0x2d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_event_event_proto_rawDescOnce sync.Once
	file_v1_event_event_proto_rawDescData = file_v1_event_event_proto_rawDesc
)

func file_v1_event_event_proto_rawDescGZIP() []byte {
	file_v1_event_event_proto_rawDescOnce.Do(func() {
		file_v1_event_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_event_event_proto_rawDescData)
	})
	return file_v1_event_event_proto_rawDescData
}

var file_v1_event_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_event_event_proto_goTypes = []interface{}{
	(*PublishUsageEventsRequest)(nil),  // 0: sensory.api.v1.event.PublishUsageEventsRequest
	(*UsageEvent)(nil),                 // 1: sensory.api.v1.event.UsageEvent
	(*PublishUsageEventsResponse)(nil), // 2: sensory.api.v1.event.PublishUsageEventsResponse
	(*timestamppb.Timestamp)(nil),      // 3: google.protobuf.Timestamp
	(common.UsageEventType)(0),         // 4: sensory.api.common.UsageEventType
}
var file_v1_event_event_proto_depIdxs = []int32{
	1, // 0: sensory.api.v1.event.PublishUsageEventsRequest.events:type_name -> sensory.api.v1.event.UsageEvent
	3, // 1: sensory.api.v1.event.UsageEvent.timestamp:type_name -> google.protobuf.Timestamp
	4, // 2: sensory.api.v1.event.UsageEvent.type:type_name -> sensory.api.common.UsageEventType
	0, // 3: sensory.api.v1.event.EventService.PublishUsageEvents:input_type -> sensory.api.v1.event.PublishUsageEventsRequest
	2, // 4: sensory.api.v1.event.EventService.PublishUsageEvents:output_type -> sensory.api.v1.event.PublishUsageEventsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_event_event_proto_init() }
func file_v1_event_event_proto_init() {
	if File_v1_event_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_event_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishUsageEventsRequest); i {
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
		file_v1_event_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsageEvent); i {
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
		file_v1_event_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishUsageEventsResponse); i {
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
			RawDescriptor: file_v1_event_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_event_event_proto_goTypes,
		DependencyIndexes: file_v1_event_event_proto_depIdxs,
		MessageInfos:      file_v1_event_event_proto_msgTypes,
	}.Build()
	File_v1_event_event_proto = out.File
	file_v1_event_event_proto_rawDesc = nil
	file_v1_event_event_proto_goTypes = nil
	file_v1_event_event_proto_depIdxs = nil
}
