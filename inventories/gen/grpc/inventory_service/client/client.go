// Code generated by goa v3.11.3, DO NOT EDIT.
//
// InventoryService gRPC client
//
// Command:
// $ goa gen mpg/inventories/design

package client

import (
	"context"
	inventory_servicepb "mpg/inventories/gen/grpc/inventory_service/pb"

	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli inventory_servicepb.InventoryServiceClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the InventoryService service
// servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: inventory_servicepb.NewInventoryServiceClient(cc),
		opts:    opts,
	}
}

// CreateInventory calls the "CreateInventory" function in
// inventory_servicepb.InventoryServiceClient interface.
func (c *Client) CreateInventory() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildCreateInventoryFunc(c.grpccli, c.opts...),
			EncodeCreateInventoryRequest,
			DecodeCreateInventoryResponse)
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

// GetInventory calls the "GetInventory" function in
// inventory_servicepb.InventoryServiceClient interface.
func (c *Client) GetInventory() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildGetInventoryFunc(c.grpccli, c.opts...),
			EncodeGetInventoryRequest,
			DecodeGetInventoryResponse)
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

// UpdateInventory calls the "UpdateInventory" function in
// inventory_servicepb.InventoryServiceClient interface.
func (c *Client) UpdateInventory() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildUpdateInventoryFunc(c.grpccli, c.opts...),
			EncodeUpdateInventoryRequest,
			DecodeUpdateInventoryResponse)
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

// DeleteInventory calls the "DeleteInventory" function in
// inventory_servicepb.InventoryServiceClient interface.
func (c *Client) DeleteInventory() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildDeleteInventoryFunc(c.grpccli, c.opts...),
			EncodeDeleteInventoryRequest,
			DecodeDeleteInventoryResponse)
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
