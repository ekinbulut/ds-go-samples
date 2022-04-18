package main

import (
	hservice "github.com/ekinbulut-yemeksepeti/samples/app/hashtable"
)

func main() {

	// Create a new HashTable.
	ht := hservice.NewHashTable(10)

	// Add a new key-value pair.
	ht.Add("key", "value")
	ht.Add("key2", "value2")
	ht.Add("key2", "value3")

	ht.Print()

}
