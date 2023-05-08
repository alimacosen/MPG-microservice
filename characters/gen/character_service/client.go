// Code generated by goa v3.11.3, DO NOT EDIT.
//
// CharacterService client
//
// Command:
// $ goa gen mpg/characters/design

package characterservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "CharacterService" service client.
type Client struct {
	CreateCharacterEndpoint goa.Endpoint
	GetCharacterEndpoint    goa.Endpoint
	UpdateCharacterEndpoint goa.Endpoint
	DeleteCharacterEndpoint goa.Endpoint
}

// NewClient initializes a "CharacterService" service client given the
// endpoints.
func NewClient(createCharacter, getCharacter, updateCharacter, deleteCharacter goa.Endpoint) *Client {
	return &Client{
		CreateCharacterEndpoint: createCharacter,
		GetCharacterEndpoint:    getCharacter,
		UpdateCharacterEndpoint: updateCharacter,
		DeleteCharacterEndpoint: deleteCharacter,
	}
}

// CreateCharacter calls the "createCharacter" endpoint of the
// "CharacterService" service.
// CreateCharacter may return the following errors:
//   - "create_invalid_args" (type CreateInvalidArgs)
//   - "create_name_not_unique" (type CreateNameNotUnique)
//   - error: internal error
func (c *Client) CreateCharacter(ctx context.Context, p *CreateCharacterPayload) (res *Character, err error) {
	var ires any
	ires, err = c.CreateCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Character), nil
}

// GetCharacter calls the "getCharacter" endpoint of the "CharacterService"
// service.
// GetCharacter may return the following errors:
//   - "get_invalid_args" (type GetInvalidArgs)
//   - "get_no_match" (type GetNoMatch)
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
// "CharacterService" service.
// UpdateCharacter may return the following errors:
//   - "update_name_not_unique" (type UpdateNameNotUnique)
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
// "CharacterService" service.
// DeleteCharacter may return the following errors:
//   - "delete_invalid_args" (type DeleteInvalidArgs)
//   - "delete_no_match" (type DeleteNoMatch)
//   - error: internal error
func (c *Client) DeleteCharacter(ctx context.Context, p *DeleteCharacterPayload) (res int, err error) {
	var ires any
	ires, err = c.DeleteCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}
