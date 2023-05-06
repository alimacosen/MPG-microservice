// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ItemService service
//
// Command:
// $ goa gen mpg/items/design

package itemservice

import (
	"context"
)

// The item service performs CRUD operations for items
type Service interface {
	// CreateItem implements createItem.
	CreateItem(context.Context, *CreateItemPayload) (res *Item, err error)
	// GetItem implements getItem.
	GetItem(context.Context, *GetItemPayload) (res *Item, err error)
	// GetAllItems implements getAllItems.
	GetAllItems(context.Context) (res []*Item, err error)
	// UpdateItem implements updateItem.
	UpdateItem(context.Context, *UpdateItemPayload) (res int, err error)
	// DeleteItem implements deleteItem.
	DeleteItem(context.Context, *DeleteItemPayload) (res int, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "ItemService"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"createItem", "getItem", "getAllItems", "updateItem", "deleteItem"}

// CreateItemPayload is the payload type of the ItemService service createItem
// method.
type CreateItemPayload struct {
	// name of the item
	Name string
	// Description of the item
	Description string
	// The amount of damage the item can do
	Damage int
	// The amount of healing the item can do
	Healing int
	// The amount of protection the item can do
	Protection int
}

// DeleteItemPayload is the payload type of the ItemService service deleteItem
// method.
type DeleteItemPayload struct {
	// UUId of the item
	ID string
}

// GetItemPayload is the payload type of the ItemService service getItem method.
type GetItemPayload struct {
	// UUId of the item
	ID string
}

// Item is the result type of the ItemService service createItem method.
type Item struct {
	// UUId of the item
	ID string
	// name of the item
	Name string
	// Description of the item
	Description string
	// The amount of damage the item can do
	Damage int
	// The amount of healing the item can do
	Healing int
	// The amount of protection the item can do
	Protection int
}

// UpdateItemPayload is the payload type of the ItemService service updateItem
// method.
type UpdateItemPayload struct {
	// UUId of the item
	ID string
	// Description of the item
	Description *string
	// The amount of damage the item can do
	Damage *int
	// The amount of healing the item can do
	Healing *int
	// The amount of protection the item can do
	Protection *int
}

// Duplicated name. This item name already exists
type CreateDuplicatedName string

// Invalid arguments. Required: name, description, damage, healing, protection
type CreateInvalidArgs string

// Invalid arguments. Required: id
type DeleteInvalidArgs string

// No item matched given criteria
type DeleteNoMatch string

// Invalid arguments. Required: id
type GetInvalidArgs string

// No item matched given criteria
type GetNoMatch string

// Invalid arguments. Required: id, itemsId
type UpdateInvalidArgs string

// No item matched given criteria
type UpdateNoMatch string

// Error returns an error description.
func (e CreateDuplicatedName) Error() string {
	return "Duplicated name. This item name already exists "
}

// ErrorName returns "create_duplicated_name".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e CreateDuplicatedName) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "create_duplicated_name".
func (e CreateDuplicatedName) GoaErrorName() string {
	return "create_duplicated_name"
}

// Error returns an error description.
func (e CreateInvalidArgs) Error() string {
	return "Invalid arguments. Required: name, description, damage, healing, protection "
}

// ErrorName returns "create_invalid_args".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e CreateInvalidArgs) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "create_invalid_args".
func (e CreateInvalidArgs) GoaErrorName() string {
	return "create_invalid_args"
}

// Error returns an error description.
func (e DeleteInvalidArgs) Error() string {
	return "Invalid arguments. Required: id "
}

// ErrorName returns "delete_invalid_args".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e DeleteInvalidArgs) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "delete_invalid_args".
func (e DeleteInvalidArgs) GoaErrorName() string {
	return "delete_invalid_args"
}

// Error returns an error description.
func (e DeleteNoMatch) Error() string {
	return "No item matched given criteria"
}

// ErrorName returns "delete_no_match".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e DeleteNoMatch) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "delete_no_match".
func (e DeleteNoMatch) GoaErrorName() string {
	return "delete_no_match"
}

// Error returns an error description.
func (e GetInvalidArgs) Error() string {
	return "Invalid arguments. Required: id "
}

// ErrorName returns "get_invalid_args".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e GetInvalidArgs) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "get_invalid_args".
func (e GetInvalidArgs) GoaErrorName() string {
	return "get_invalid_args"
}

// Error returns an error description.
func (e GetNoMatch) Error() string {
	return "No item matched given criteria"
}

// ErrorName returns "get_no_match".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e GetNoMatch) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "get_no_match".
func (e GetNoMatch) GoaErrorName() string {
	return "get_no_match"
}

// Error returns an error description.
func (e UpdateInvalidArgs) Error() string {
	return "Invalid arguments. Required: id, itemsId "
}

// ErrorName returns "update_invalid_args".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e UpdateInvalidArgs) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "update_invalid_args".
func (e UpdateInvalidArgs) GoaErrorName() string {
	return "update_invalid_args"
}

// Error returns an error description.
func (e UpdateNoMatch) Error() string {
	return "No item matched given criteria"
}

// ErrorName returns "update_no_match".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e UpdateNoMatch) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "update_no_match".
func (e UpdateNoMatch) GoaErrorName() string {
	return "update_no_match"
}
