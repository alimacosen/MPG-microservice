// Code generated with goa v3.11.3, DO NOT EDIT.
//
// InventoryService protocol buffer definition
//
// Command:
// $ goa gen inventories/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: goadesign_goagen_inventory_service.proto

package inventory_servicepb

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

type CreateInventoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUId of the character that this inventory belongs to
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateInventoryRequest) Reset() {
	*x = CreateInventoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInventoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInventoryRequest) ProtoMessage() {}

func (x *CreateInventoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInventoryRequest.ProtoReflect.Descriptor instead.
func (*CreateInventoryRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_inventory_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateInventoryRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateInventoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUId of the inventory
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// UUId of the character that this inventory belongs to
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Array of uuid of items
	ItemsId []string `protobuf:"bytes,3,rep,name=items_id,json=itemsId,proto3" json:"items_id,omitempty"`
}

func (x *CreateInventoryResponse) Reset() {
	*x = CreateInventoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInventoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInventoryResponse) ProtoMessage() {}

func (x *CreateInventoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInventoryResponse.ProtoReflect.Descriptor instead.
func (*CreateInventoryResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_inventory_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateInventoryResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateInventoryResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateInventoryResponse) GetItemsId() []string {
	if x != nil {
		return x.ItemsId
	}
	return nil
}

type GetInventoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUId of the inventory
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetInventoryRequest) Reset() {
	*x = GetInventoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInventoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInventoryRequest) ProtoMessage() {}

func (x *GetInventoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInventoryRequest.ProtoReflect.Descriptor instead.
func (*GetInventoryRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_inventory_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetInventoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetInventoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUId of the inventory
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// UUId of the character that this inventory belongs to
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Array of uuid of items
	ItemsId []string `protobuf:"bytes,3,rep,name=items_id,json=itemsId,proto3" json:"items_id,omitempty"`
}

func (x *GetInventoryResponse) Reset() {
	*x = GetInventoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInventoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInventoryResponse) ProtoMessage() {}

func (x *GetInventoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInventoryResponse.ProtoReflect.Descriptor instead.
func (*GetInventoryResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_inventory_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetInventoryResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetInventoryResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetInventoryResponse) GetItemsId() []string {
	if x != nil {
		return x.ItemsId
	}
	return nil
}

type UpdateInventoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUId of the Inventory
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Array of uuid of items
	ItemsId []string `protobuf:"bytes,2,rep,name=items_id,json=itemsId,proto3" json:"items_id,omitempty"`
}

func (x *UpdateInventoryRequest) Reset() {
	*x = UpdateInventoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInventoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInventoryRequest) ProtoMessage() {}

func (x *UpdateInventoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInventoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateInventoryRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_inventory_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateInventoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateInventoryRequest) GetItemsId() []string {
	if x != nil {
		return x.ItemsId
	}
	return nil
}

type UpdateInventoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field int32 `protobuf:"zigzag32,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *UpdateInventoryResponse) Reset() {
	*x = UpdateInventoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInventoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInventoryResponse) ProtoMessage() {}

func (x *UpdateInventoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInventoryResponse.ProtoReflect.Descriptor instead.
func (*UpdateInventoryResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_inventory_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateInventoryResponse) GetField() int32 {
	if x != nil {
		return x.Field
	}
	return 0
}

type DeleteInventoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUId of the Inventory
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteInventoryRequest) Reset() {
	*x = DeleteInventoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInventoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInventoryRequest) ProtoMessage() {}

func (x *DeleteInventoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInventoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteInventoryRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_inventory_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteInventoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteInventoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field int32 `protobuf:"zigzag32,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *DeleteInventoryResponse) Reset() {
	*x = DeleteInventoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInventoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInventoryResponse) ProtoMessage() {}

func (x *DeleteInventoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_inventory_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInventoryResponse.ProtoReflect.Descriptor instead.
func (*DeleteInventoryResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_inventory_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteInventoryResponse) GetField() int32 {
	if x != nil {
		return x.Field
	}
	return 0
}

var File_goadesign_goagen_inventory_service_proto protoreflect.FileDescriptor

var file_goadesign_goagen_inventory_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x31, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x5d, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x64, 0x22,
	0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x49, 0x64, 0x22, 0x43, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2f, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0xb1, 0x03, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x29, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x29, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a,
	0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x29, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x2f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_inventory_service_proto_rawDescOnce sync.Once
	file_goadesign_goagen_inventory_service_proto_rawDescData = file_goadesign_goagen_inventory_service_proto_rawDesc
)

func file_goadesign_goagen_inventory_service_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_inventory_service_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_inventory_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_inventory_service_proto_rawDescData)
	})
	return file_goadesign_goagen_inventory_service_proto_rawDescData
}

var file_goadesign_goagen_inventory_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_goadesign_goagen_inventory_service_proto_goTypes = []interface{}{
	(*CreateInventoryRequest)(nil),  // 0: inventory_service.CreateInventoryRequest
	(*CreateInventoryResponse)(nil), // 1: inventory_service.CreateInventoryResponse
	(*GetInventoryRequest)(nil),     // 2: inventory_service.GetInventoryRequest
	(*GetInventoryResponse)(nil),    // 3: inventory_service.GetInventoryResponse
	(*UpdateInventoryRequest)(nil),  // 4: inventory_service.UpdateInventoryRequest
	(*UpdateInventoryResponse)(nil), // 5: inventory_service.UpdateInventoryResponse
	(*DeleteInventoryRequest)(nil),  // 6: inventory_service.DeleteInventoryRequest
	(*DeleteInventoryResponse)(nil), // 7: inventory_service.DeleteInventoryResponse
}
var file_goadesign_goagen_inventory_service_proto_depIdxs = []int32{
	0, // 0: inventory_service.InventoryService.CreateInventory:input_type -> inventory_service.CreateInventoryRequest
	2, // 1: inventory_service.InventoryService.GetInventory:input_type -> inventory_service.GetInventoryRequest
	4, // 2: inventory_service.InventoryService.UpdateInventory:input_type -> inventory_service.UpdateInventoryRequest
	6, // 3: inventory_service.InventoryService.DeleteInventory:input_type -> inventory_service.DeleteInventoryRequest
	1, // 4: inventory_service.InventoryService.CreateInventory:output_type -> inventory_service.CreateInventoryResponse
	3, // 5: inventory_service.InventoryService.GetInventory:output_type -> inventory_service.GetInventoryResponse
	5, // 6: inventory_service.InventoryService.UpdateInventory:output_type -> inventory_service.UpdateInventoryResponse
	7, // 7: inventory_service.InventoryService.DeleteInventory:output_type -> inventory_service.DeleteInventoryResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_inventory_service_proto_init() }
func file_goadesign_goagen_inventory_service_proto_init() {
	if File_goadesign_goagen_inventory_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_inventory_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInventoryRequest); i {
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
		file_goadesign_goagen_inventory_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInventoryResponse); i {
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
		file_goadesign_goagen_inventory_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInventoryRequest); i {
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
		file_goadesign_goagen_inventory_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInventoryResponse); i {
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
		file_goadesign_goagen_inventory_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInventoryRequest); i {
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
		file_goadesign_goagen_inventory_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInventoryResponse); i {
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
		file_goadesign_goagen_inventory_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInventoryRequest); i {
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
		file_goadesign_goagen_inventory_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInventoryResponse); i {
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
			RawDescriptor: file_goadesign_goagen_inventory_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_inventory_service_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_inventory_service_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_inventory_service_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_inventory_service_proto = out.File
	file_goadesign_goagen_inventory_service_proto_rawDesc = nil
	file_goadesign_goagen_inventory_service_proto_goTypes = nil
	file_goadesign_goagen_inventory_service_proto_depIdxs = nil
}