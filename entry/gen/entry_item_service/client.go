// Code generated by goa v3.11.3, DO NOT EDIT.
//
// EntryItemService client
//
// Command:
// $ goa gen mpg/entry/design

package entryitemservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "EntryItemService" service client.
type Client struct {
	CreatItemEndpoint  goa.Endpoint
	GetItemEndpoint    goa.Endpoint
	UpdateItemEndpoint goa.Endpoint
	DeleteItemEndpoint goa.Endpoint
}

// NewClient initializes a "EntryItemService" service client given the
// endpoints.
func NewClient(creatItem, getItem, updateItem, deleteItem goa.Endpoint) *Client {
	return &Client{
		CreatItemEndpoint:  creatItem,
		GetItemEndpoint:    getItem,
		UpdateItemEndpoint: updateItem,
		DeleteItemEndpoint: deleteItem,
	}
}

// CreatItem calls the "creatItem" endpoint of the "EntryItemService" service.
// CreatItem may return the following errors:
//   - "create_no_criteria" (type CreateNoCriteria)
//   - "create_invalid_args" (type CreateInvalidArgs)
//   - "create_duplicated_name" (type CreateDuplicatedName)
//   - error: internal error
func (c *Client) CreatItem(ctx context.Context, p *CreatItemPayload) (res *Item, err error) {
	var ires any
	ires, err = c.CreatItemEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Item), nil
}

// GetItem calls the "getItem" endpoint of the "EntryItemService" service.
// GetItem may return the following errors:
//   - "get_no_criteria" (type GetNoCriteria)
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

// UpdateItem calls the "updateItem" endpoint of the "EntryItemService" service.
// UpdateItem may return the following errors:
//   - "update_no_criteria" (type UpdateNoCriteria)
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

// DeleteItem calls the "deleteItem" endpoint of the "EntryItemService" service.
// DeleteItem may return the following errors:
//   - "delete_no_criteria" (type DeleteNoCriteria)
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
