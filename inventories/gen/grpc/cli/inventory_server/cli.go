// Code generated by goa v3.11.3, DO NOT EDIT.
//
// inventoryServer gRPC client CLI support package
//
// Command:
// $ goa gen mpg/inventories/design

package cli

import (
	"flag"
	"fmt"
	inventoryservicec "mpg/inventories/gen/grpc/inventory_service/client"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `inventory-service (create-inventory|get-inventory|update-inventory|delete-inventory)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` inventory-service create-inventory --message '{
      "userId": "Est in ab ut aut."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, any, error) {
	var (
		inventoryServiceFlags = flag.NewFlagSet("inventory-service", flag.ContinueOnError)

		inventoryServiceCreateInventoryFlags       = flag.NewFlagSet("create-inventory", flag.ExitOnError)
		inventoryServiceCreateInventoryMessageFlag = inventoryServiceCreateInventoryFlags.String("message", "", "")

		inventoryServiceGetInventoryFlags       = flag.NewFlagSet("get-inventory", flag.ExitOnError)
		inventoryServiceGetInventoryMessageFlag = inventoryServiceGetInventoryFlags.String("message", "", "")

		inventoryServiceUpdateInventoryFlags       = flag.NewFlagSet("update-inventory", flag.ExitOnError)
		inventoryServiceUpdateInventoryMessageFlag = inventoryServiceUpdateInventoryFlags.String("message", "", "")

		inventoryServiceDeleteInventoryFlags       = flag.NewFlagSet("delete-inventory", flag.ExitOnError)
		inventoryServiceDeleteInventoryMessageFlag = inventoryServiceDeleteInventoryFlags.String("message", "", "")
	)
	inventoryServiceFlags.Usage = inventoryServiceUsage
	inventoryServiceCreateInventoryFlags.Usage = inventoryServiceCreateInventoryUsage
	inventoryServiceGetInventoryFlags.Usage = inventoryServiceGetInventoryUsage
	inventoryServiceUpdateInventoryFlags.Usage = inventoryServiceUpdateInventoryUsage
	inventoryServiceDeleteInventoryFlags.Usage = inventoryServiceDeleteInventoryUsage

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
		case "inventory-service":
			svcf = inventoryServiceFlags
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
		case "inventory-service":
			switch epn {
			case "create-inventory":
				epf = inventoryServiceCreateInventoryFlags

			case "get-inventory":
				epf = inventoryServiceGetInventoryFlags

			case "update-inventory":
				epf = inventoryServiceUpdateInventoryFlags

			case "delete-inventory":
				epf = inventoryServiceDeleteInventoryFlags

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
		case "inventory-service":
			c := inventoryservicec.NewClient(cc, opts...)
			switch epn {
			case "create-inventory":
				endpoint = c.CreateInventory()
				data, err = inventoryservicec.BuildCreateInventoryPayload(*inventoryServiceCreateInventoryMessageFlag)
			case "get-inventory":
				endpoint = c.GetInventory()
				data, err = inventoryservicec.BuildGetInventoryPayload(*inventoryServiceGetInventoryMessageFlag)
			case "update-inventory":
				endpoint = c.UpdateInventory()
				data, err = inventoryservicec.BuildUpdateInventoryPayload(*inventoryServiceUpdateInventoryMessageFlag)
			case "delete-inventory":
				endpoint = c.DeleteInventory()
				data, err = inventoryservicec.BuildDeleteInventoryPayload(*inventoryServiceDeleteInventoryMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// inventory-serviceUsage displays the usage of the inventory-service command
// and its subcommands.
func inventoryServiceUsage() {
	fmt.Fprintf(os.Stderr, `The inventory service performs CRUD operations for inventories
Usage:
    %[1]s [globalflags] inventory-service COMMAND [flags]

COMMAND:
    create-inventory: CreateInventory implements createInventory.
    get-inventory: GetInventory implements getInventory.
    update-inventory: UpdateInventory implements updateInventory.
    delete-inventory: DeleteInventory implements deleteInventory.

Additional help:
    %[1]s inventory-service COMMAND --help
`, os.Args[0])
}
func inventoryServiceCreateInventoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory-service create-inventory -message JSON

CreateInventory implements createInventory.
    -message JSON: 

Example:
    %[1]s inventory-service create-inventory --message '{
      "userId": "Est in ab ut aut."
   }'
`, os.Args[0])
}

func inventoryServiceGetInventoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory-service get-inventory -message JSON

GetInventory implements getInventory.
    -message JSON: 

Example:
    %[1]s inventory-service get-inventory --message '{
      "id": "Aut eveniet quo."
   }'
`, os.Args[0])
}

func inventoryServiceUpdateInventoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory-service update-inventory -message JSON

UpdateInventory implements updateInventory.
    -message JSON: 

Example:
    %[1]s inventory-service update-inventory --message '{
      "id": "Mollitia quis sint omnis similique at sint.",
      "itemsId": [
         "Debitis sit sit veritatis omnis.",
         "Suscipit sunt minus voluptate corporis enim."
      ]
   }'
`, os.Args[0])
}

func inventoryServiceDeleteInventoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory-service delete-inventory -message JSON

DeleteInventory implements deleteInventory.
    -message JSON: 

Example:
    %[1]s inventory-service delete-inventory --message '{
      "id": "Laboriosam hic aliquid excepturi."
   }'
`, os.Args[0])
}
