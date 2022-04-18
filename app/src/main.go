package main

import (
	aservice "github.com/ekinbulut-yemeksepeti/samples/app/arraylist"
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

	// create a new arraylist
	list := aservice.NewArrayList()

	list.Add("value")
	list.Add("value2")
	list.Add("value3")

	list.Print()

	s := list.Size()

	println("Size: ", s)

	println("Removing: ")
	list.Remove(1)

	println("Size: ", s)

	list.Print()

}
