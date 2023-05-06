// Code generated by goa v3.11.3, DO NOT EDIT.
//
// HTTP request path constructors for the EntryInventoryService service.
//
// Command:
// $ goa gen mpg/entry/design

package client

import (
	"fmt"
)

// GetInventoryEntryInventoryServicePath returns the URL path to the EntryInventoryService service getInventory HTTP endpoint.
func GetInventoryEntryInventoryServicePath(id string) string {
	return fmt.Sprintf("/inventory/%v", id)
}

// UpdateInventoryEntryInventoryServicePath returns the URL path to the EntryInventoryService service updateInventory HTTP endpoint.
func UpdateInventoryEntryInventoryServicePath(id string) string {
	return fmt.Sprintf("/inventory/%v", id)
}