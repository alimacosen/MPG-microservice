// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ItemService gRPC client
//
// Command:
// $ goa gen mpg/items/design

package client

import (
	"context"
	item_servicepb "mpg/items/gen/grpc/item_service/pb"

	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli item_servicepb.ItemServiceClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the ItemService service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: item_servicepb.NewItemServiceClient(cc),
		opts:    opts,
	}
}

// CreateItem calls the "CreateItem" function in
// item_servicepb.ItemServiceClient interface.
func (c *Client) CreateItem() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildCreateItemFunc(c.grpccli, c.opts...),
			EncodeCreateItemRequest,
			DecodeCreateItemResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// GetItems calls the "GetItems" function in item_servicepb.ItemServiceClient
// interface.
func (c *Client) GetItems() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildGetItemsFunc(c.grpccli, c.opts...),
			EncodeGetItemsRequest,
			DecodeGetItemsResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// GetAllItems calls the "GetAllItems" function in
// item_servicepb.ItemServiceClient interface.
func (c *Client) GetAllItems() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildGetAllItemsFunc(c.grpccli, c.opts...),
			nil,
			DecodeGetAllItemsResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// UpdateItem calls the "UpdateItem" function in
// item_servicepb.ItemServiceClient interface.
func (c *Client) UpdateItem() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildUpdateItemFunc(c.grpccli, c.opts...),
			EncodeUpdateItemRequest,
			DecodeUpdateItemResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// DeleteItem calls the "DeleteItem" function in
// item_servicepb.ItemServiceClient interface.
func (c *Client) DeleteItem() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildDeleteItemFunc(c.grpccli, c.opts...),
			EncodeDeleteItemRequest,
			DecodeDeleteItemResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}
