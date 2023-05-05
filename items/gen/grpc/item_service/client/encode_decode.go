// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ItemService gRPC client encoders and decoders
//
// Command:
// $ goa gen mpg/items/design

package client

import (
	"context"
	item_servicepb "mpg/items/gen/grpc/item_service/pb"
	itemservice "mpg/items/gen/item_service"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildCreateItemFunc builds the remote method to invoke for "ItemService"
// service "createItem" endpoint.
func BuildCreateItemFunc(grpccli item_servicepb.ItemServiceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.CreateItem(ctx, reqpb.(*item_servicepb.CreateItemRequest), opts...)
		}
		return grpccli.CreateItem(ctx, &item_servicepb.CreateItemRequest{}, opts...)
	}
}

// EncodeCreateItemRequest encodes requests sent to ItemService createItem
// endpoint.
func EncodeCreateItemRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*itemservice.CreateItemPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "createItem", "*itemservice.CreateItemPayload", v)
	}
	return NewProtoCreateItemRequest(payload), nil
}

// DecodeCreateItemResponse decodes responses from the ItemService createItem
// endpoint.
func DecodeCreateItemResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*item_servicepb.CreateItemResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "createItem", "*item_servicepb.CreateItemResponse", v)
	}
	res := NewCreateItemResult(message)
	return res, nil
}

// BuildGetItemFunc builds the remote method to invoke for "ItemService"
// service "getItem" endpoint.
func BuildGetItemFunc(grpccli item_servicepb.ItemServiceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.GetItem(ctx, reqpb.(*item_servicepb.GetItemRequest), opts...)
		}
		return grpccli.GetItem(ctx, &item_servicepb.GetItemRequest{}, opts...)
	}
}

// EncodeGetItemRequest encodes requests sent to ItemService getItem endpoint.
func EncodeGetItemRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*itemservice.GetItemPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "getItem", "*itemservice.GetItemPayload", v)
	}
	return NewProtoGetItemRequest(payload), nil
}

// DecodeGetItemResponse decodes responses from the ItemService getItem
// endpoint.
func DecodeGetItemResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*item_servicepb.GetItemResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "getItem", "*item_servicepb.GetItemResponse", v)
	}
	res := NewGetItemResult(message)
	return res, nil
}

// BuildUpdateItemFunc builds the remote method to invoke for "ItemService"
// service "updateItem" endpoint.
func BuildUpdateItemFunc(grpccli item_servicepb.ItemServiceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.UpdateItem(ctx, reqpb.(*item_servicepb.UpdateItemRequest), opts...)
		}
		return grpccli.UpdateItem(ctx, &item_servicepb.UpdateItemRequest{}, opts...)
	}
}

// EncodeUpdateItemRequest encodes requests sent to ItemService updateItem
// endpoint.
func EncodeUpdateItemRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*itemservice.UpdateItemPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "updateItem", "*itemservice.UpdateItemPayload", v)
	}
	return NewProtoUpdateItemRequest(payload), nil
}

// DecodeUpdateItemResponse decodes responses from the ItemService updateItem
// endpoint.
func DecodeUpdateItemResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*item_servicepb.UpdateItemResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "updateItem", "*item_servicepb.UpdateItemResponse", v)
	}
	res := NewUpdateItemResult(message)
	return res, nil
}

// BuildDeleteItemFunc builds the remote method to invoke for "ItemService"
// service "deleteItem" endpoint.
func BuildDeleteItemFunc(grpccli item_servicepb.ItemServiceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.DeleteItem(ctx, reqpb.(*item_servicepb.DeleteItemRequest), opts...)
		}
		return grpccli.DeleteItem(ctx, &item_servicepb.DeleteItemRequest{}, opts...)
	}
}

// EncodeDeleteItemRequest encodes requests sent to ItemService deleteItem
// endpoint.
func EncodeDeleteItemRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*itemservice.DeleteItemPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "deleteItem", "*itemservice.DeleteItemPayload", v)
	}
	return NewProtoDeleteItemRequest(payload), nil
}

// DecodeDeleteItemResponse decodes responses from the ItemService deleteItem
// endpoint.
func DecodeDeleteItemResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*item_servicepb.DeleteItemResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "deleteItem", "*item_servicepb.DeleteItemResponse", v)
	}
	res := NewDeleteItemResult(message)
	return res, nil
}
