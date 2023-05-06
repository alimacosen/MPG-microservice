// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ItemService gRPC client types
//
// Command:
// $ goa gen mpg/items/design

package client

import (
	item_servicepb "mpg/items/gen/grpc/item_service/pb"
	itemservice "mpg/items/gen/item_service"
)

// NewProtoCreateItemRequest builds the gRPC request type from the payload of
// the "createItem" endpoint of the "ItemService" service.
func NewProtoCreateItemRequest(payload *itemservice.CreateItemPayload) *item_servicepb.CreateItemRequest {
	message := &item_servicepb.CreateItemRequest{
		Name:        payload.Name,
		Description: payload.Description,
		Damage:      int32(payload.Damage),
		Healing:     int32(payload.Healing),
		Protection:  int32(payload.Protection),
	}
	return message
}

// NewCreateItemResult builds the result type of the "createItem" endpoint of
// the "ItemService" service from the gRPC response type.
func NewCreateItemResult(message *item_servicepb.CreateItemResponse) *itemservice.Item {
	result := &itemservice.Item{
		ID:          message.Id,
		Name:        message.Name,
		Description: message.Description,
		Damage:      int(message.Damage),
		Healing:     int(message.Healing),
		Protection:  int(message.Protection),
	}
	return result
}

// NewProtoGetItemRequest builds the gRPC request type from the payload of the
// "getItem" endpoint of the "ItemService" service.
func NewProtoGetItemRequest(payload *itemservice.GetItemPayload) *item_servicepb.GetItemRequest {
	message := &item_servicepb.GetItemRequest{
		Id: payload.ID,
	}
	return message
}

// NewGetItemResult builds the result type of the "getItem" endpoint of the
// "ItemService" service from the gRPC response type.
func NewGetItemResult(message *item_servicepb.GetItemResponse) *itemservice.Item {
	result := &itemservice.Item{
		ID:          message.Id,
		Name:        message.Name,
		Description: message.Description,
		Damage:      int(message.Damage),
		Healing:     int(message.Healing),
		Protection:  int(message.Protection),
	}
	return result
}

// NewProtoGetAllItemsRequest builds the gRPC request type from the payload of
// the "getAllItems" endpoint of the "ItemService" service.
func NewProtoGetAllItemsRequest() *item_servicepb.GetAllItemsRequest {
	message := &item_servicepb.GetAllItemsRequest{}
	return message
}

// NewGetAllItemsResult builds the result type of the "getAllItems" endpoint of
// the "ItemService" service from the gRPC response type.
func NewGetAllItemsResult(message *item_servicepb.GetAllItemsResponse) []*itemservice.Item {
	result := make([]*itemservice.Item, len(message.Field))
	for i, val := range message.Field {
		result[i] = &itemservice.Item{
			ID:          val.Id,
			Name:        val.Name,
			Description: val.Description,
			Damage:      int(val.Damage),
			Healing:     int(val.Healing),
			Protection:  int(val.Protection),
		}
	}
	return result
}

// NewProtoUpdateItemRequest builds the gRPC request type from the payload of
// the "updateItem" endpoint of the "ItemService" service.
func NewProtoUpdateItemRequest(payload *itemservice.UpdateItemPayload) *item_servicepb.UpdateItemRequest {
	message := &item_servicepb.UpdateItemRequest{
		Id:          payload.ID,
		Description: payload.Description,
	}
	if payload.Damage != nil {
		damage := int32(*payload.Damage)
		message.Damage = &damage
	}
	if payload.Healing != nil {
		healing := int32(*payload.Healing)
		message.Healing = &healing
	}
	if payload.Protection != nil {
		protection := int32(*payload.Protection)
		message.Protection = &protection
	}
	return message
}

// NewUpdateItemResult builds the result type of the "updateItem" endpoint of
// the "ItemService" service from the gRPC response type.
func NewUpdateItemResult(message *item_servicepb.UpdateItemResponse) int {
	result := int(message.Field)
	return result
}

// NewProtoDeleteItemRequest builds the gRPC request type from the payload of
// the "deleteItem" endpoint of the "ItemService" service.
func NewProtoDeleteItemRequest(payload *itemservice.DeleteItemPayload) *item_servicepb.DeleteItemRequest {
	message := &item_servicepb.DeleteItemRequest{
		Id: payload.ID,
	}
	return message
}

// NewDeleteItemResult builds the result type of the "deleteItem" endpoint of
// the "ItemService" service from the gRPC response type.
func NewDeleteItemResult(message *item_servicepb.DeleteItemResponse) int {
	result := int(message.Field)
	return result
}
