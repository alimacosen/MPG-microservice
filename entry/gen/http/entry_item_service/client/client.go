// Code generated by goa v3.11.3, DO NOT EDIT.
//
// EntryItemService client HTTP transport
//
// Command:
// $ goa gen mpg/entry/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the EntryItemService service endpoint HTTP clients.
type Client struct {
	// CreatItem Doer is the HTTP client used to make requests to the creatItem
	// endpoint.
	CreatItemDoer goahttp.Doer

	// GetItem Doer is the HTTP client used to make requests to the getItem
	// endpoint.
	GetItemDoer goahttp.Doer

	// UpdateItem Doer is the HTTP client used to make requests to the updateItem
	// endpoint.
	UpdateItemDoer goahttp.Doer

	// DeleteItem Doer is the HTTP client used to make requests to the deleteItem
	// endpoint.
	DeleteItemDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the EntryItemService service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		CreatItemDoer:       doer,
		GetItemDoer:         doer,
		UpdateItemDoer:      doer,
		DeleteItemDoer:      doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// CreatItem returns an endpoint that makes HTTP requests to the
// EntryItemService service creatItem server.
func (c *Client) CreatItem() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreatItemRequest(c.encoder)
		decodeResponse = DecodeCreatItemResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildCreatItemRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreatItemDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("EntryItemService", "creatItem", err)
		}
		return decodeResponse(resp)
	}
}

// GetItem returns an endpoint that makes HTTP requests to the EntryItemService
// service getItem server.
func (c *Client) GetItem() goa.Endpoint {
	var (
		decodeResponse = DecodeGetItemResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildGetItemRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetItemDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("EntryItemService", "getItem", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateItem returns an endpoint that makes HTTP requests to the
// EntryItemService service updateItem server.
func (c *Client) UpdateItem() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateItemRequest(c.encoder)
		decodeResponse = DecodeUpdateItemResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildUpdateItemRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateItemDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("EntryItemService", "updateItem", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteItem returns an endpoint that makes HTTP requests to the
// EntryItemService service deleteItem server.
func (c *Client) DeleteItem() goa.Endpoint {
	var (
		decodeResponse = DecodeDeleteItemResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildDeleteItemRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteItemDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("EntryItemService", "deleteItem", err)
		}
		return decodeResponse(resp)
	}
}
