// Sample datastore-quickstart fetches an entity from Google Cloud Datastore.
package main

import (
	"fmt"
	"log"
	"time"
	// Imports the Google Cloud Datastore client package.
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

type KeyValueBlockchain struct {
	KeyBlokchain []byte
	DataGCS      string
}

func main() {
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
	name := int64(time.Now().Unix()) //time.Now() // "sampletask3"
	// Creates a Key instance.
	//taskKey := datastore.NameKey("Task", "sampletask", nil)
	taskKey := datastore.IDKey(kind, name, nil)

	// Creates a Task instance.
	task := KeyValueBlockchain{
		KeyBlokchain: []byte("dogr3"),
		DataGCS:      "url gcs22",
	}

	// Saves the new entity.
	if _, err := client.Put(ctx, taskKey, &task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}

	fmt.Printf("Saved %v: %v\n %v\n", taskKey, task.KeyBlokchain, task.DataGCS)
}
