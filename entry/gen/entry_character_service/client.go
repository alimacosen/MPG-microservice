// Code generated by goa v3.11.3, DO NOT EDIT.
//
// EntryCharacterService client
//
// Command:
// $ goa gen mpg/entry/design

package entrycharacterservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "EntryCharacterService" service client.
type Client struct {
	CreatCharacterEndpoint  goa.Endpoint
	GetCharacterEndpoint    goa.Endpoint
	UpdateCharacterEndpoint goa.Endpoint
	DeleteCharacterEndpoint goa.Endpoint
}

// NewClient initializes a "EntryCharacterService" service client given the
// endpoints.
func NewClient(creatCharacter, getCharacter, updateCharacter, deleteCharacter goa.Endpoint) *Client {
	return &Client{
		CreatCharacterEndpoint:  creatCharacter,
		GetCharacterEndpoint:    getCharacter,
		UpdateCharacterEndpoint: updateCharacter,
		DeleteCharacterEndpoint: deleteCharacter,
	}
}

// CreatCharacter calls the "creatCharacter" endpoint of the
// "EntryCharacterService" service.
// CreatCharacter may return the following errors:
//   - "create_no_criteria" (type CreateNoCriteria)
//   - "create_invalid_args" (type CreateInvalidArgs)
//   - error: internal error
func (c *Client) CreatCharacter(ctx context.Context, p *CreatCharacterPayload) (res *Character, err error) {
	var ires any
	ires, err = c.CreatCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Character), nil
}

// GetCharacter calls the "getCharacter" endpoint of the
// "EntryCharacterService" service.
// GetCharacter may return the following errors:
//   - "get_invalid_args" (type GetInvalidArgs)
//   - "get_no_match" (type GetNoMatch)
//   - "get_no_criteria" (type GetNoCriteria)
//   - error: internal error
func (c *Client) GetCharacter(ctx context.Context, p *GetCharacterPayload) (res *Character, err error) {
	var ires any
	ires, err = c.GetCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Character), nil
}

// UpdateCharacter calls the "updateCharacter" endpoint of the
// "EntryCharacterService" service.
// UpdateCharacter may return the following errors:
//   - "no_criteria" (type NoCriteria)
//   - "update_invalid_args" (type UpdateInvalidArgs)
//   - "update_no_match" (type UpdateNoMatch)
//   - error: internal error
func (c *Client) UpdateCharacter(ctx context.Context, p *UpdateCharacterPayload) (res int, err error) {
	var ires any
	ires, err = c.UpdateCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// DeleteCharacter calls the "deleteCharacter" endpoint of the
// "EntryCharacterService" service.
// DeleteCharacter may return the following errors:
//   - "delete_invalid_args" (type DeleteInvalidArgs)
//   - "delete_no_match" (type DeleteNoMatch)
//   - "no_criteria" (type NoCriteria)
//   - error: internal error
func (c *Client) DeleteCharacter(ctx context.Context, p *DeleteCharacterPayload) (res int, err error) {
	var ires any
	ires, err = c.DeleteCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}