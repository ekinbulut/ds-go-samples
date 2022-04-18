package hashtable_test

import (
	"testing"

	"github.com/ekinbulut-yemeksepeti/samples/app/hashtable"
)

// TestNewHashTable
func TestNewHashTable(t *testing.T) {

	// Create a new HashTable.
	ht := hashtable.NewHashTable(10)

	// Test
	if ht == nil {
		t.Error("NewHashTable() should not return nil")
	}
}

// TestAdd
func TestAdd(t *testing.T) {

	// Create a new HashTable.
	ht := hashtable.NewHashTable(10)

	// Add an element to the end of the list.
	ht.Add("key", "value")

	// Test
	if ht.Size() != 10 {
		t.Error("Add() should increase the size of the list")
	}
}

// TestGet
func TestGet(t *testing.T) {

	// Create a new HashTable.
	ht := hashtable.NewHashTable(10)

	// Add an element to the end of the list.
	ht.Add("key", "value")

	// Test
	if ht.Get("key") != "value" {
		t.Error("Get() should return the element at the specified index")
	}
}

// TestSize
func TestSize(t *testing.T) {

	// Create a new HashTable.
	ht := hashtable.NewHashTable(10)

	// Add an element to the end of the list.
	ht.Add("key", "value")

	// Test
	if ht.Size() != 10 {
		t.Error("Size() should return the size of the list")
	}
}
