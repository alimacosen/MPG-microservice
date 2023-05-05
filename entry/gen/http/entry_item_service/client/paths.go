// Code generated by goa v3.11.3, DO NOT EDIT.
//
// HTTP request path constructors for the EntryItemService service.
//
// Command:
// $ goa gen mpg/entry/design

package client

import (
	"fmt"
)

// CreatItemEntryItemServicePath returns the URL path to the EntryItemService service creatItem HTTP endpoint.
func CreatItemEntryItemServicePath() string {
	return "/item"
}

// GetItemEntryItemServicePath returns the URL path to the EntryItemService service getItem HTTP endpoint.
func GetItemEntryItemServicePath(id string) string {
	return fmt.Sprintf("/item/%v", id)
}

// UpdateItemEntryItemServicePath returns the URL path to the EntryItemService service updateItem HTTP endpoint.
func UpdateItemEntryItemServicePath(id string) string {
	return fmt.Sprintf("/item/%v", id)
}

// DeleteItemEntryItemServicePath returns the URL path to the EntryItemService service deleteItem HTTP endpoint.
func DeleteItemEntryItemServicePath(id string) string {
	return fmt.Sprintf("/item/%v", id)
}
