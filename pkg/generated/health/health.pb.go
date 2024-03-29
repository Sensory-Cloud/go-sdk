// sensory.api.health

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: health/health.proto

package health

import (
	common "github.com/Sensory-Cloud/go-sdk/pkg/generated/common"
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

// Request to obtain a Health response
type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_health_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_health_health_proto_rawDescGZIP(), []int{0}
}

var File_health_health_proto protoreflect.FileDescriptor

var file_health_health_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f,
	0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32,
	0x6b, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x21, 0x2e,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x76, 0x0a, 0x1a,
	0x61, 0x69, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x42, 0x15, 0x53, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x79, 0x41, 0x70, 0x69, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x74, 0x69, 0x74, 0x61, 0x6e, 0x2e, 0x67, 0x69, 0x74, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0xa2, 0x02, 0x04,
	0x53, 0x45, 0x4e, 0x47, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_health_health_proto_rawDescOnce sync.Once
	file_health_health_proto_rawDescData = file_health_health_proto_rawDesc
)

func file_health_health_proto_rawDescGZIP() []byte {
	file_health_health_proto_rawDescOnce.Do(func() {
		file_health_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_health_proto_rawDescData)
	})
	return file_health_health_proto_rawDescData
}

var file_health_health_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_health_health_proto_goTypes = []interface{}{
	(*HealthRequest)(nil),               // 0: sensory.api.health.HealthRequest
	(*common.ServerHealthResponse)(nil), // 1: sensory.api.common.ServerHealthResponse
}
var file_health_health_proto_depIdxs = []int32{
	0, // 0: sensory.api.health.HealthService.GetHealth:input_type -> sensory.api.health.HealthRequest
	1, // 1: sensory.api.health.HealthService.GetHealth:output_type -> sensory.api.common.ServerHealthResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_health_health_proto_init() }
func file_health_health_proto_init() {
	if File_health_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_health_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
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
			RawDescriptor: file_health_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_health_proto_goTypes,
		DependencyIndexes: file_health_health_proto_depIdxs,
		MessageInfos:      file_health_health_proto_msgTypes,
	}.Build()
	File_health_health_proto = out.File
	file_health_health_proto_rawDesc = nil
	file_health_health_proto_goTypes = nil
	file_health_health_proto_depIdxs = nil
}
