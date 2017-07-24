// Sample storage-quickstart creates a Google Cloud Storage bucket.
package main

import (
	"fmt"
	"log"

	// Imports the Google Cloud Storage client package.
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	//projectID := "geth2-173819"

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new bucket.
	bucketName := "raw-blockchain"

	// Creates a Bucket instance.
	bucket := client.Bucket(bucketName)

	dataName := "datazombi2"

	obj := bucket.Object(dataName)

	// Read it back.
	//rc, err := obj.NewReader(ctx)

	// Write something to obj.
	// w implements io.Writer.
	w := obj.NewWriter(ctx)
	// Write some text to obj. This will overwrite whatever is there.

	if _, err := w.Write([]byte("abcde\n")); err != nil {
		fmt.Printf("createFile: unable to write data to bucket %q, file %q: %v", bucket, dataName, err)
	}

	// Close, just like writing a file.
	if err := w.Close(); err != nil {
		// TODO: Handle error.
		fmt.Printf("Bucket %v created.\n", bucketName)
	}

}
