// Code generated by goa v3.11.3, DO NOT EDIT.
//
// EntryItemService HTTP server types
//
// Command:
// $ goa gen mpg/entry/design

package server

import (
	entryitemservice "mpg/entry/gen/entry_item_service"

	goa "goa.design/goa/v3/pkg"
)

// CreatItemRequestBody is the type of the "EntryItemService" service
// "creatItem" endpoint HTTP request body.
type CreatItemRequestBody struct {
	// name of the item
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the item
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The amount of damage the item can do
	Damage *int `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// The amount of healing the item can do
	Healing *int `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// The amount of protection the item can do
	Protection *int `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// UpdateItemRequestBody is the type of the "EntryItemService" service
// "updateItem" endpoint HTTP request body.
type UpdateItemRequestBody struct {
	// Description of the item
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// The amount of damage the item can do
	Damage *int `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// The amount of healing the item can do
	Healing *int `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// The amount of protection the item can do
	Protection *int `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// CreatItemResponseBody is the type of the "EntryItemService" service
// "creatItem" endpoint HTTP response body.
type CreatItemResponseBody struct {
	// UUId of the item
	ID string `form:"id" json:"id" xml:"id"`
	// name of the item
	Name string `form:"name" json:"name" xml:"name"`
	// Description of the item
	Description string `form:"description" json:"description" xml:"description"`
	// The amount of damage the item can do
	Damage int `form:"damage" json:"damage" xml:"damage"`
	// The amount of healing the item can do
	Healing int `form:"healing" json:"healing" xml:"healing"`
	// The amount of protection the item can do
	Protection int `form:"protection" json:"protection" xml:"protection"`
}

// GetItemResponseBody is the type of the "EntryItemService" service "getItem"
// endpoint HTTP response body.
type GetItemResponseBody struct {
	// UUId of the item
	ID string `form:"id" json:"id" xml:"id"`
	// name of the item
	Name string `form:"name" json:"name" xml:"name"`
	// Description of the item
	Description string `form:"description" json:"description" xml:"description"`
	// The amount of damage the item can do
	Damage int `form:"damage" json:"damage" xml:"damage"`
	// The amount of healing the item can do
	Healing int `form:"healing" json:"healing" xml:"healing"`
	// The amount of protection the item can do
	Protection int `form:"protection" json:"protection" xml:"protection"`
}

// NewCreatItemResponseBody builds the HTTP response body from the result of
// the "creatItem" endpoint of the "EntryItemService" service.
func NewCreatItemResponseBody(res *entryitemservice.Item) *CreatItemResponseBody {
	body := &CreatItemResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Damage:      res.Damage,
		Healing:     res.Healing,
		Protection:  res.Protection,
	}
	return body
}

// NewGetItemResponseBody builds the HTTP response body from the result of the
// "getItem" endpoint of the "EntryItemService" service.
func NewGetItemResponseBody(res *entryitemservice.Item) *GetItemResponseBody {
	body := &GetItemResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Damage:      res.Damage,
		Healing:     res.Healing,
		Protection:  res.Protection,
	}
	return body
}

// NewCreatItemPayload builds a EntryItemService service creatItem endpoint
// payload.
func NewCreatItemPayload(body *CreatItemRequestBody) *entryitemservice.CreatItemPayload {
	v := &entryitemservice.CreatItemPayload{
		Name:        *body.Name,
		Description: *body.Description,
		Damage:      *body.Damage,
		Healing:     *body.Healing,
		Protection:  *body.Protection,
	}

	return v
}

// NewGetItemPayload builds a EntryItemService service getItem endpoint payload.
func NewGetItemPayload(id string) *entryitemservice.GetItemPayload {
	v := &entryitemservice.GetItemPayload{}
	v.ID = id

	return v
}

// NewUpdateItemPayload builds a EntryItemService service updateItem endpoint
// payload.
func NewUpdateItemPayload(body *UpdateItemRequestBody, id string) *entryitemservice.UpdateItemPayload {
	v := &entryitemservice.UpdateItemPayload{
		Description: body.Description,
		Damage:      body.Damage,
		Healing:     body.Healing,
		Protection:  body.Protection,
	}
	v.ID = id

	return v
}

// NewDeleteItemPayload builds a EntryItemService service deleteItem endpoint
// payload.
func NewDeleteItemPayload(id string) *entryitemservice.DeleteItemPayload {
	v := &entryitemservice.DeleteItemPayload{}
	v.ID = id

	return v
}

// ValidateCreatItemRequestBody runs the validations defined on
// CreatItemRequestBody
func ValidateCreatItemRequestBody(body *CreatItemRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Damage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("damage", "body"))
	}
	if body.Healing == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("healing", "body"))
	}
	if body.Protection == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("protection", "body"))
	}
	return
}
