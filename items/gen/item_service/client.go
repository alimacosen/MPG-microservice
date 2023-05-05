// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ItemService client
//
// Command:
// $ goa gen mpg/items/design

package itemservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "ItemService" service client.
type Client struct {
	CreateItemEndpoint goa.Endpoint
	GetItemEndpoint    goa.Endpoint
	UpdateItemEndpoint goa.Endpoint
	DeleteItemEndpoint goa.Endpoint
}

// NewClient initializes a "ItemService" service client given the endpoints.
func NewClient(createItem, getItem, updateItem, deleteItem goa.Endpoint) *Client {
	return &Client{
		CreateItemEndpoint: createItem,
		GetItemEndpoint:    getItem,
		UpdateItemEndpoint: updateItem,
		DeleteItemEndpoint: deleteItem,
	}
}

// CreateItem calls the "createItem" endpoint of the "ItemService" service.
// CreateItem may return the following errors:
//   - "create_invalid_args" (type CreateInvalidArgs)
//   - "create_duplicated_name" (type CreateDuplicatedName)
//   - error: internal error
func (c *Client) CreateItem(ctx context.Context, p *CreateItemPayload) (res *Item, err error) {
	var ires any
	ires, err = c.CreateItemEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Item), nil
}

// GetItem calls the "getItem" endpoint of the "ItemService" service.
// GetItem may return the following errors:
//   - "get_invalid_args" (type GetInvalidArgs)
//   - "get_no_match" (type GetNoMatch)
//   - error: internal error
func (c *Client) GetItem(ctx context.Context, p *GetItemPayload) (res *Item, err error) {
	var ires any
	ires, err = c.GetItemEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Item), nil
}

// UpdateItem calls the "updateItem" endpoint of the "ItemService" service.
// UpdateItem may return the following errors:
//   - "update_invalid_args" (type UpdateInvalidArgs)
//   - "update_no_match" (type UpdateNoMatch)
//   - error: internal error
func (c *Client) UpdateItem(ctx context.Context, p *UpdateItemPayload) (res int, err error) {
	var ires any
	ires, err = c.UpdateItemEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// DeleteItem calls the "deleteItem" endpoint of the "ItemService" service.
// DeleteItem may return the following errors:
//   - "delete_invalid_args" (type DeleteInvalidArgs)
//   - "delete_no_match" (type DeleteNoMatch)
//   - error: internal error
func (c *Client) DeleteItem(ctx context.Context, p *DeleteItemPayload) (res int, err error) {
	var ires any
	ires, err = c.DeleteItemEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}
