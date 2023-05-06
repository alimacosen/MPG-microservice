// Code generated by goa v3.11.3, DO NOT EDIT.
//
// EntryInventoryService HTTP client CLI support package
//
// Command:
// $ goa gen mpg/entry/design

package client

import (
	"encoding/json"
	"fmt"
	entryinventoryservice "mpg/entry/gen/entry_inventory_service"

	goa "goa.design/goa/v3/pkg"
)

// BuildGetInventoryPayload builds the payload for the EntryInventoryService
// getInventory endpoint from CLI flags.
func BuildGetInventoryPayload(entryInventoryServiceGetInventoryID string) (*entryinventoryservice.GetInventoryPayload, error) {
	var id string
	{
		id = entryInventoryServiceGetInventoryID
	}
	v := &entryinventoryservice.GetInventoryPayload{}
	v.ID = id

	return v, nil
}

// BuildUpdateInventoryPayload builds the payload for the EntryInventoryService
// updateInventory endpoint from CLI flags.
func BuildUpdateInventoryPayload(entryInventoryServiceUpdateInventoryBody string, entryInventoryServiceUpdateInventoryID string) (*entryinventoryservice.UpdateInventoryPayload, error) {
	var err error
	var body UpdateInventoryRequestBody
	{
		err = json.Unmarshal([]byte(entryInventoryServiceUpdateInventoryBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"itemsId\": [\n         \"Quibusdam nobis natus a.\",\n         \"Tempore qui.\"\n      ]\n   }'")
		}
		if body.ItemsID == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("itemsId", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = entryInventoryServiceUpdateInventoryID
	}
	v := &entryinventoryservice.UpdateInventoryPayload{}
	if body.ItemsID != nil {
		v.ItemsID = make([]string, len(body.ItemsID))
		for i, val := range body.ItemsID {
			v.ItemsID[i] = val
		}
	}
	v.ID = id

	return v, nil
}