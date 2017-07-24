// Sample datastore-quickstart fetches an entity from Google Cloud Datastore.
package main

import (
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	// Imports the Google Cloud Datastore client package.
)

type KeyValueBlockchain struct {
	KeyBlokchain []byte
	DataGCS      string
}

func main() {
	//entre := false

	ctx := context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := "geth2-173819"

	// Creates a client.

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the kind for the new entity.
	kind := "BlockchainEthereum"
	// Sets the name/ID for the new entity.

	//1500218125fc345b1df08ecd716819b1dcab0e

	key := datastore.NameKey(kind, "QmxvY2tjaGFpblZlcnNpb24=", nil)
	kvBlockchain := &KeyValueBlockchain{}
	if err := client.Get(ctx, key, kvBlockchain); err != nil {
		// TODO: Handle error.
		fmt.Printf("erro %v\n", err)
	} else {
		fmt.Printf("Key=%v\n kvBlockchain =%#v\n\n", key, kvBlockchain)
		fmt.Printf("Key=%v\n kvBlockchain =%#v\n\n", key, kvBlockchain.DataGCS)
	}

	key2 := datastore.NameKey(kind, "SDZuMfPCzUFNdAHn5mcgrWpBzK1XOomiHsxLh2dM4Zlfk", nil)
	kvBlockchain2 := &KeyValueBlockchain{}
	if err := client.Get(ctx, key2, kvBlockchain2); err != nil {
		if err != datastore.ErrNoSuchEntity {
			fmt.Printf("erro %v\n", err)
		} else {
			fmt.Printf("no existe e elemento")
		}

	} else {
		fmt.Printf("Key=%v\n kvBlockchain =%#v\n\n", key2, kvBlockchain2)
	}

}
