// Sample datastore-quickstart fetches an entity from Google Cloud Datastore.
package main

import (
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	// Imports the Google Cloud Datastore client package.
	"google.golang.org/api/iterator"
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
	kind := "Task"
	// Sets the name/ID for the new entity.

	//1500218125fc345b1df08ecd716819b1dcab0e

	q := datastore.NewQuery(kind).
		Filter("KeyBlokchain =", []byte("dogr3"))

	for t := client.Run(ctx, q); ; {

		var x KeyValueBlockchain
		key, err := t.Next(&x)
		if err == iterator.Done {
			fmt.Printf("pase break")
			break
		}
		if err != nil {
			// Handle error.
			fmt.Printf("error %v", err)
		}
		fmt.Printf("Key=%v\nWidget=%#v\n\n", key, x)
		//entre = true
	}
	/*
		fmt.Printf("pase %v", entre)
		if entre == false {
			fmt.Printf("es nuevo")
		}
	*/

}
