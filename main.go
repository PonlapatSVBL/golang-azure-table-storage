package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/storage"
)

var (
	tableCli storage.TableServiceClient
)

const (
	account      = ""
	key          = ""
	fullmetadata = "application/json;odata=fullmetadata"
	tablename    = ""
)

func main() {
	query()
}

func query() {
	client, err := storage.NewBasicClient(account, key)

	if err != nil {
		fmt.Printf("%s: \n", err)
	}

	tableCli = client.GetTableService()

	// fmt.Println(tableCli)

	table := tableCli.GetTableReference(tablename)

	// timeout, metatadalevel, options
	entities, err := table.QueryEntities(2, fullmetadata, nil)

	if err != nil {
		fmt.Println(err)
	}

	for _, entity := range entities.Entities {
		// Convert entity to JSON
		entityJSON, err := json.Marshal(entity)
		if err != nil {
			log.Fatalf("Failed to marshal entity to JSON: %v", err)
		}
		fmt.Println(string(entityJSON))
		fmt.Println()
	}
}
