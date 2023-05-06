// Code generated by goa v3.11.3, DO NOT EDIT.
//
// EntryInventoryService HTTP client encoders and decoders
//
// Command:
// $ goa gen mpg/entry/design

package client

import (
	"bytes"
	"context"
	"io"
	entryinventoryservice "mpg/entry/gen/entry_inventory_service"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetInventoryRequest instantiates a HTTP request object with method and
// path set to call the "EntryInventoryService" service "getInventory" endpoint
func (c *Client) BuildGetInventoryRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*entryinventoryservice.GetInventoryPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("EntryInventoryService", "getInventory", "*entryinventoryservice.GetInventoryPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetInventoryEntryInventoryServicePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("EntryInventoryService", "getInventory", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetInventoryResponse returns a decoder for responses returned by the
// EntryInventoryService getInventory endpoint. restoreBody controls whether
// the response body should be restored after having been read.
// DecodeGetInventoryResponse may return the following errors:
//   - "get_invalid_args" (type entryinventoryservice.GetInvalidArgs): http.StatusBadRequest
//   - "get_no_criteria" (type entryinventoryservice.GetNoCriteria): http.StatusBadRequest
//   - "get_no_match" (type entryinventoryservice.GetNoMatch): http.StatusNotFound
//   - error: internal error
func DecodeGetInventoryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetInventoryResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("EntryInventoryService", "getInventory", err)
			}
			err = ValidateGetInventoryResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("EntryInventoryService", "getInventory", err)
			}
			res := NewGetInventoryInventoryOK(&body)
			return res, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "get_invalid_args":
				var (
					body string
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("EntryInventoryService", "getInventory", err)
				}
				return nil, NewGetInventoryGetInvalidArgs(body)
			case "get_no_criteria":
				var (
					body string
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("EntryInventoryService", "getInventory", err)
				}
				return nil, NewGetInventoryGetNoCriteria(body)
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("EntryInventoryService", "getInventory", resp.StatusCode, string(body))
			}
		case http.StatusNotFound:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("EntryInventoryService", "getInventory", err)
			}
			return nil, NewGetInventoryGetNoMatch(body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("EntryInventoryService", "getInventory", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateInventoryRequest instantiates a HTTP request object with method
// and path set to call the "EntryInventoryService" service "updateInventory"
// endpoint
func (c *Client) BuildUpdateInventoryRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*entryinventoryservice.UpdateInventoryPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("EntryInventoryService", "updateInventory", "*entryinventoryservice.UpdateInventoryPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateInventoryEntryInventoryServicePath(id)}
	req, err := http.NewRequest("PATCH", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("EntryInventoryService", "updateInventory", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateInventoryRequest returns an encoder for requests sent to the
// EntryInventoryService updateInventory server.
func EncodeUpdateInventoryRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*entryinventoryservice.UpdateInventoryPayload)
		if !ok {
			return goahttp.ErrInvalidType("EntryInventoryService", "updateInventory", "*entryinventoryservice.UpdateInventoryPayload", v)
		}
		body := NewUpdateInventoryRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("EntryInventoryService", "updateInventory", err)
		}
		return nil
	}
}

// DecodeUpdateInventoryResponse returns a decoder for responses returned by
// the EntryInventoryService updateInventory endpoint. restoreBody controls
// whether the response body should be restored after having been read.
// DecodeUpdateInventoryResponse may return the following errors:
//   - "update_invalid_args" (type entryinventoryservice.UpdateInvalidArgs): http.StatusBadRequest
//   - "update_no_criteria" (type entryinventoryservice.UpdateNoCriteria): http.StatusBadRequest
//   - "update_no_match" (type entryinventoryservice.UpdateNoMatch): http.StatusNotFound
//   - error: internal error
func DecodeUpdateInventoryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("EntryInventoryService", "updateInventory", err)
			}
			return body, nil
		case http.StatusBadRequest:
			en := resp.Header.Get("goa-error")
			switch en {
			case "update_invalid_args":
				var (
					body string
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("EntryInventoryService", "updateInventory", err)
				}
				return nil, NewUpdateInventoryUpdateInvalidArgs(body)
			case "update_no_criteria":
				var (
					body string
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("EntryInventoryService", "updateInventory", err)
				}
				return nil, NewUpdateInventoryUpdateNoCriteria(body)
			default:
				body, _ := io.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("EntryInventoryService", "updateInventory", resp.StatusCode, string(body))
			}
		case http.StatusNotFound:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("EntryInventoryService", "updateInventory", err)
			}
			return nil, NewUpdateInventoryUpdateNoMatch(body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("EntryInventoryService", "updateInventory", resp.StatusCode, string(body))
		}
	}
}