// sensory.api.management

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: v1/management/device.proto

package management

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Request to enroll a new device in the system. Devices can then enroll and associate users with themselves.
type EnrollDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The friendly name you'd like to associate with this device
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The unique internal identifier for this device. Ideally, this value is static for the lifetime of the device.
	// A typical value would be a hardware serial number.
	DeviceId string `protobuf:"bytes,2,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	// Unique Tenant ID as UUID
	TenantId string `protobuf:"bytes,3,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// OAuth client-specific request details
	Client *common.GenericClient `protobuf:"bytes,4,opt,name=client,proto3" json:"client,omitempty"`
	// If required, the credential to be validated by the server upon enrollment.
	// Possible values are a shared secret or signed JWT.
	Credential string `protobuf:"bytes,5,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (x *EnrollDeviceRequest) Reset() {
	*x = EnrollDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_management_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollDeviceRequest) ProtoMessage() {}

func (x *EnrollDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_management_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollDeviceRequest.ProtoReflect.Descriptor instead.
func (*EnrollDeviceRequest) Descriptor() ([]byte, []int) {
	return file_v1_management_device_proto_rawDescGZIP(), []int{0}
}

func (x *EnrollDeviceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnrollDeviceRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *EnrollDeviceRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *EnrollDeviceRequest) GetClient() *common.GenericClient {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *EnrollDeviceRequest) GetCredential() string {
	if x != nil {
		return x.Credential
	}
	return ""
}

// Request to renew a device's credential
type RenewDeviceCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique internal identifier for this device. Ideally, this value is static for the lifetime of the device.
	// A typical value would be a hardware serial number.
	DeviceId string `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	// Unique Client ID as UUID
	ClientId string `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	// Unique Tenant ID as UUID
	TenantId string `protobuf:"bytes,3,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// The credential to be validated by the server upon enrollment.
	// Possible values are a shared secret or signed JWT.
	Credential string `protobuf:"bytes,4,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (x *RenewDeviceCredentialRequest) Reset() {
	*x = RenewDeviceCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_management_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenewDeviceCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenewDeviceCredentialRequest) ProtoMessage() {}

func (x *RenewDeviceCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_management_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenewDeviceCredentialRequest.ProtoReflect.Descriptor instead.
func (*RenewDeviceCredentialRequest) Descriptor() ([]byte, []int) {
	return file_v1_management_device_proto_rawDescGZIP(), []int{1}
}

func (x *RenewDeviceCredentialRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *RenewDeviceCredentialRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *RenewDeviceCredentialRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *RenewDeviceCredentialRequest) GetCredential() string {
	if x != nil {
		return x.Credential
	}
	return ""
}

// A request to obtain information about the calling device
type DeviceGetWhoAmIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeviceGetWhoAmIRequest) Reset() {
	*x = DeviceGetWhoAmIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_management_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceGetWhoAmIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceGetWhoAmIRequest) ProtoMessage() {}

func (x *DeviceGetWhoAmIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_management_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceGetWhoAmIRequest.ProtoReflect.Descriptor instead.
func (*DeviceGetWhoAmIRequest) Descriptor() ([]byte, []int) {
	return file_v1_management_device_proto_rawDescGZIP(), []int{2}
}

// A request for a device
type DeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique internal identifier for this device. Ideally, this value is static for the lifetime of the device.
	DeviceId string `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
}

func (x *DeviceRequest) Reset() {
	*x = DeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_management_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRequest) ProtoMessage() {}

func (x *DeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_management_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRequest.ProtoReflect.Descriptor instead.
func (*DeviceRequest) Descriptor() ([]byte, []int) {
	return file_v1_management_device_proto_rawDescGZIP(), []int{3}
}

func (x *DeviceRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type GetDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TenantID to filter on. TenantId validation is handled within the tenant_filter_middleware.
	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// Metadata about how to paginate the response
	Pagination *common.PaginationOptions `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// Optional user id to filter devices by
	UserId string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetDevicesRequest) Reset() {
	*x = GetDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_management_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDevicesRequest) ProtoMessage() {}

func (x *GetDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_management_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDevicesRequest.ProtoReflect.Descriptor instead.
func (*GetDevicesRequest) Descriptor() ([]byte, []int) {
	return file_v1_management_device_proto_rawDescGZIP(), []int{4}
}

func (x *GetDevicesRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *GetDevicesRequest) GetPagination() *common.PaginationOptions {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *GetDevicesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// A request to update the name of a device
type UpdateDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique internal identifier for this device. Ideally, this value is static for the lifetime of the device.
	DeviceId string `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	// The new name of the device
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateDeviceRequest) Reset() {
	*x = UpdateDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_management_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeviceRequest) ProtoMessage() {}

func (x *UpdateDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_management_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeviceRequest.ProtoReflect.Descriptor instead.
func (*UpdateDeviceRequest) Descriptor() ([]byte, []int) {
	return file_v1_management_device_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDeviceRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *UpdateDeviceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A response containing information about a device
type DeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The friendly name of the device
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The unique internal identifier for this device. Ideally, this value is static for the lifetime of the device.
	DeviceId string `protobuf:"bytes,2,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
}

func (x *DeviceResponse) Reset() {
	*x = DeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_management_device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceResponse) ProtoMessage() {}

func (x *DeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_management_device_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceResponse.ProtoReflect.Descriptor instead.
func (*DeviceResponse) Descriptor() ([]byte, []int) {
	return file_v1_management_device_proto_rawDescGZIP(), []int{6}
}

func (x *DeviceResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceResponse) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

// A response containing information about a device and associated users
type GetDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The friendly name of the device
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The unique internal identifier for this device. Ideally, this value is static for the lifetime of the device.
	DeviceId string `protobuf:"bytes,2,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	// The count of total users that are linked to the device
	UserCount int64 `protobuf:"varint,3,opt,name=userCount,proto3" json:"userCount,omitempty"`
}

func (x *GetDeviceResponse) Reset() {
	*x = GetDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_management_device_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceResponse) ProtoMessage() {}

func (x *GetDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_management_device_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceResponse) Descriptor() ([]byte, []int) {
	return file_v1_management_device_proto_rawDescGZIP(), []int{7}
}

func (x *GetDeviceResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetDeviceResponse) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *GetDeviceResponse) GetUserCount() int64 {
	if x != nil {
		return x.UserCount
	}
	return 0
}

// A response containing multiple devices
type DeviceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of devices
	Devices []*DeviceResponse `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
	// Metadata about how the response has been paginated
	Pagination *common.PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *DeviceListResponse) Reset() {
	*x = DeviceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_management_device_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceListResponse) ProtoMessage() {}

func (x *DeviceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_management_device_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceListResponse.ProtoReflect.Descriptor instead.
func (*DeviceListResponse) Descriptor() ([]byte, []int) {
	return file_v1_management_device_proto_rawDescGZIP(), []int{8}
}

func (x *DeviceListResponse) GetDevices() []*DeviceResponse {
	if x != nil {
		return x.Devices
	}
	return nil
}

func (x *DeviceListResponse) GetPagination() *common.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_v1_management_device_proto protoreflect.FileDescriptor

var file_v1_management_device_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x13, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x01, 0x18, 0x7f, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x7f, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x06, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0xbd,
	0x01, 0x0a, 0x1c, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x7f, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0,
	0x01, 0x01, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x2a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18,
	0xff, 0x01, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x18,
	0x0a, 0x16, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x57, 0x68, 0x6f, 0x41, 0x6d,
	0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x01, 0x18, 0x7f, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x22, 0x97, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x45, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x18, 0x7f, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x7f, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18,
	0x7f, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x0e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa1, 0x01, 0x0a,
	0x12, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x32, 0x90, 0x06, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x6b, 0x0a, 0x0c, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2e, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45,
	0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x7d, 0x0a, 0x15, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x37, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x12, 0x31, 0x2e, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65,
	0x74, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x6b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x2c, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x6b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2e, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x73,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x90, 0x01, 0x0a, 0x21, 0x61, 0x69, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x21, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x79, 0x41, 0x70, 0x69, 0x56, 0x31, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x79, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x74, 0x69, 0x74, 0x61, 0x6e, 0x2e, 0x67, 0x69, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0xa2,
	0x02, 0x04, 0x53, 0x45, 0x4e, 0x47, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_management_device_proto_rawDescOnce sync.Once
	file_v1_management_device_proto_rawDescData = file_v1_management_device_proto_rawDesc
)

func file_v1_management_device_proto_rawDescGZIP() []byte {
	file_v1_management_device_proto_rawDescOnce.Do(func() {
		file_v1_management_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_management_device_proto_rawDescData)
	})
	return file_v1_management_device_proto_rawDescData
}

var file_v1_management_device_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v1_management_device_proto_goTypes = []interface{}{
	(*EnrollDeviceRequest)(nil),          // 0: sensory.api.v1.management.EnrollDeviceRequest
	(*RenewDeviceCredentialRequest)(nil), // 1: sensory.api.v1.management.RenewDeviceCredentialRequest
	(*DeviceGetWhoAmIRequest)(nil),       // 2: sensory.api.v1.management.DeviceGetWhoAmIRequest
	(*DeviceRequest)(nil),                // 3: sensory.api.v1.management.DeviceRequest
	(*GetDevicesRequest)(nil),            // 4: sensory.api.v1.management.GetDevicesRequest
	(*UpdateDeviceRequest)(nil),          // 5: sensory.api.v1.management.UpdateDeviceRequest
	(*DeviceResponse)(nil),               // 6: sensory.api.v1.management.DeviceResponse
	(*GetDeviceResponse)(nil),            // 7: sensory.api.v1.management.GetDeviceResponse
	(*DeviceListResponse)(nil),           // 8: sensory.api.v1.management.DeviceListResponse
	(*common.GenericClient)(nil),         // 9: sensory.api.common.GenericClient
	(*common.PaginationOptions)(nil),     // 10: sensory.api.common.PaginationOptions
	(*common.PaginationResponse)(nil),    // 11: sensory.api.common.PaginationResponse
}
var file_v1_management_device_proto_depIdxs = []int32{
	9,  // 0: sensory.api.v1.management.EnrollDeviceRequest.client:type_name -> sensory.api.common.GenericClient
	10, // 1: sensory.api.v1.management.GetDevicesRequest.pagination:type_name -> sensory.api.common.PaginationOptions
	6,  // 2: sensory.api.v1.management.DeviceListResponse.devices:type_name -> sensory.api.v1.management.DeviceResponse
	11, // 3: sensory.api.v1.management.DeviceListResponse.pagination:type_name -> sensory.api.common.PaginationResponse
	0,  // 4: sensory.api.v1.management.DeviceService.EnrollDevice:input_type -> sensory.api.v1.management.EnrollDeviceRequest
	1,  // 5: sensory.api.v1.management.DeviceService.RenewDeviceCredential:input_type -> sensory.api.v1.management.RenewDeviceCredentialRequest
	2,  // 6: sensory.api.v1.management.DeviceService.GetWhoAmI:input_type -> sensory.api.v1.management.DeviceGetWhoAmIRequest
	3,  // 7: sensory.api.v1.management.DeviceService.GetDevice:input_type -> sensory.api.v1.management.DeviceRequest
	4,  // 8: sensory.api.v1.management.DeviceService.GetDevices:input_type -> sensory.api.v1.management.GetDevicesRequest
	5,  // 9: sensory.api.v1.management.DeviceService.UpdateDevice:input_type -> sensory.api.v1.management.UpdateDeviceRequest
	3,  // 10: sensory.api.v1.management.DeviceService.DeleteDevice:input_type -> sensory.api.v1.management.DeviceRequest
	6,  // 11: sensory.api.v1.management.DeviceService.EnrollDevice:output_type -> sensory.api.v1.management.DeviceResponse
	6,  // 12: sensory.api.v1.management.DeviceService.RenewDeviceCredential:output_type -> sensory.api.v1.management.DeviceResponse
	6,  // 13: sensory.api.v1.management.DeviceService.GetWhoAmI:output_type -> sensory.api.v1.management.DeviceResponse
	7,  // 14: sensory.api.v1.management.DeviceService.GetDevice:output_type -> sensory.api.v1.management.GetDeviceResponse
	8,  // 15: sensory.api.v1.management.DeviceService.GetDevices:output_type -> sensory.api.v1.management.DeviceListResponse
	6,  // 16: sensory.api.v1.management.DeviceService.UpdateDevice:output_type -> sensory.api.v1.management.DeviceResponse
	6,  // 17: sensory.api.v1.management.DeviceService.DeleteDevice:output_type -> sensory.api.v1.management.DeviceResponse
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_v1_management_device_proto_init() }
func file_v1_management_device_proto_init() {
	if File_v1_management_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_management_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollDeviceRequest); i {
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
		file_v1_management_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenewDeviceCredentialRequest); i {
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
		file_v1_management_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceGetWhoAmIRequest); i {
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
		file_v1_management_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceRequest); i {
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
		file_v1_management_device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDevicesRequest); i {
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
		file_v1_management_device_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeviceRequest); i {
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
		file_v1_management_device_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceResponse); i {
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
		file_v1_management_device_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceResponse); i {
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
		file_v1_management_device_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceListResponse); i {
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
			RawDescriptor: file_v1_management_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_management_device_proto_goTypes,
		DependencyIndexes: file_v1_management_device_proto_depIdxs,
		MessageInfos:      file_v1_management_device_proto_msgTypes,
	}.Build()
	File_v1_management_device_proto = out.File
	file_v1_management_device_proto_rawDesc = nil
	file_v1_management_device_proto_goTypes = nil
	file_v1_management_device_proto_depIdxs = nil
}
