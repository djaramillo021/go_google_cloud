// Sample datastore-quickstart fetches an entity from Google Cloud Datastore.
package main

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"golang.org/x/crypto/sha3"
)

func main() {

	s := "sha1 this string"
	h := sha3.New224()
	h.Write([]byte(s))
	sha3_hash := hex.EncodeToString(h.Sum(nil))
	sha3_hash = strconv.Itoa(123) + sha3_hash[0:len(sha3_hash)/2] // sha3_hash

	fmt.Println(s, sha3_hash)
}
