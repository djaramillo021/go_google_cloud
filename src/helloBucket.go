// Sample storage-quickstart creates a Google Cloud Storage bucket.
package main

import (
	"fmt"
	"io/ioutil"
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

	dataName := "1500255670b740a3b446920772a10caf69736e"

	obj := bucket.Object(dataName)

	// Read it back.
	//rc, err := obj.NewReader(ctx)

	// Read it back.
	r, err := obj.NewReader(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer r.Close()

	b, errFile := ioutil.ReadAll(r)
	if err != nil {
		//log.Fatal(err)
		fmt.Printf("%v", errFile)
	}

	fmt.Printf("%v", b)

	/*
		if _, err := io.Copy(os.Stdout, r); err != nil {
			// TODO: Handle error.
		}
		// Prints "This object contains text."
	*/
	//	fmt.Printf("Bucket %v created.\n", bucketName)
}
