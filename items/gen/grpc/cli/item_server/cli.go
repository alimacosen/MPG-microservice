// Code generated by goa v3.11.3, DO NOT EDIT.
//
// itemServer gRPC client CLI support package
//
// Command:
// $ goa gen mpg/items/design

package cli

import (
	"flag"
	"fmt"
	itemservicec "mpg/items/gen/grpc/item_service/client"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `item-service (create-item|get-item|get-all-items|update-item|delete-item)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` item-service create-item --message '{
      "damage": 5496475948197560084,
      "description": "Cupiditate et non ipsam blanditiis dolores.",
      "healing": 2130137277969925607,
      "name": "Neque saepe sed velit laboriosam dolorem.",
      "protection": 6970409012914527909
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, any, error) {
	var (
		itemServiceFlags = flag.NewFlagSet("item-service", flag.ContinueOnError)

		itemServiceCreateItemFlags       = flag.NewFlagSet("create-item", flag.ExitOnError)
		itemServiceCreateItemMessageFlag = itemServiceCreateItemFlags.String("message", "", "")

		itemServiceGetItemFlags       = flag.NewFlagSet("get-item", flag.ExitOnError)
		itemServiceGetItemMessageFlag = itemServiceGetItemFlags.String("message", "", "")

		itemServiceGetAllItemsFlags = flag.NewFlagSet("get-all-items", flag.ExitOnError)

		itemServiceUpdateItemFlags       = flag.NewFlagSet("update-item", flag.ExitOnError)
		itemServiceUpdateItemMessageFlag = itemServiceUpdateItemFlags.String("message", "", "")

		itemServiceDeleteItemFlags       = flag.NewFlagSet("delete-item", flag.ExitOnError)
		itemServiceDeleteItemMessageFlag = itemServiceDeleteItemFlags.String("message", "", "")
	)
	itemServiceFlags.Usage = itemServiceUsage
	itemServiceCreateItemFlags.Usage = itemServiceCreateItemUsage
	itemServiceGetItemFlags.Usage = itemServiceGetItemUsage
	itemServiceGetAllItemsFlags.Usage = itemServiceGetAllItemsUsage
	itemServiceUpdateItemFlags.Usage = itemServiceUpdateItemUsage
	itemServiceDeleteItemFlags.Usage = itemServiceDeleteItemUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "item-service":
			svcf = itemServiceFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "item-service":
			switch epn {
			case "create-item":
				epf = itemServiceCreateItemFlags

			case "get-item":
				epf = itemServiceGetItemFlags

			case "get-all-items":
				epf = itemServiceGetAllItemsFlags

			case "update-item":
				epf = itemServiceUpdateItemFlags

			case "delete-item":
				epf = itemServiceDeleteItemFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "item-service":
			c := itemservicec.NewClient(cc, opts...)
			switch epn {
			case "create-item":
				endpoint = c.CreateItem()
				data, err = itemservicec.BuildCreateItemPayload(*itemServiceCreateItemMessageFlag)
			case "get-item":
				endpoint = c.GetItem()
				data, err = itemservicec.BuildGetItemPayload(*itemServiceGetItemMessageFlag)
			case "get-all-items":
				endpoint = c.GetAllItems()
				data = nil
			case "update-item":
				endpoint = c.UpdateItem()
				data, err = itemservicec.BuildUpdateItemPayload(*itemServiceUpdateItemMessageFlag)
			case "delete-item":
				endpoint = c.DeleteItem()
				data, err = itemservicec.BuildDeleteItemPayload(*itemServiceDeleteItemMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// item-serviceUsage displays the usage of the item-service command and its
// subcommands.
func itemServiceUsage() {
	fmt.Fprintf(os.Stderr, `The item service performs CRUD operations for items
Usage:
    %[1]s [globalflags] item-service COMMAND [flags]

COMMAND:
    create-item: CreateItem implements createItem.
    get-item: GetItem implements getItem.
    get-all-items: GetAllItems implements getAllItems.
    update-item: UpdateItem implements updateItem.
    delete-item: DeleteItem implements deleteItem.

Additional help:
    %[1]s item-service COMMAND --help
`, os.Args[0])
}
func itemServiceCreateItemUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item-service create-item -message JSON

CreateItem implements createItem.
    -message JSON: 

Example:
    %[1]s item-service create-item --message '{
      "damage": 5496475948197560084,
      "description": "Cupiditate et non ipsam blanditiis dolores.",
      "healing": 2130137277969925607,
      "name": "Neque saepe sed velit laboriosam dolorem.",
      "protection": 6970409012914527909
   }'
`, os.Args[0])
}

func itemServiceGetItemUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item-service get-item -message JSON

GetItem implements getItem.
    -message JSON: 

Example:
    %[1]s item-service get-item --message '{
      "id": "Temporibus sed voluptates occaecati est."
   }'
`, os.Args[0])
}

func itemServiceGetAllItemsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item-service get-all-items

GetAllItems implements getAllItems.

Example:
    %[1]s item-service get-all-items
`, os.Args[0])
}

func itemServiceUpdateItemUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item-service update-item -message JSON

UpdateItem implements updateItem.
    -message JSON: 

Example:
    %[1]s item-service update-item --message '{
      "damage": 1560793895885502624,
      "description": "Earum magnam sequi.",
      "healing": 5206333268185226675,
      "id": "Et quis aperiam harum rerum.",
      "protection": 273878408864270749
   }'
`, os.Args[0])
}

func itemServiceDeleteItemUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item-service delete-item -message JSON

DeleteItem implements deleteItem.
    -message JSON: 

Example:
    %[1]s item-service delete-item --message '{
      "id": "Quod dolor."
   }'
`, os.Args[0])
}
