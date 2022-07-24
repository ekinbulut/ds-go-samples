package arraylist_test

// Imports
import (
	"testing"

	"samples/app/arraylist"
)

// TestNewArrayList
func TestNewArrayList(t *testing.T) {

	// Create a new ArrayList.
	list := arraylist.NewArrayList()

	// Test
	if list == nil {
		t.Error("NewArrayList() should not return nil")
	}
}

// TestAdd
func TestAdd(t *testing.T) {

	// Create a new ArrayList.
	list := arraylist.NewArrayList()

	// Add an element to the end of the list.
	list.Add("value")

	// Test
	if list.Size() != 1 {
		t.Error("Add() should increase the size of the list")
	}
}

// TestGet
func TestGet(t *testing.T) {

	// Create a new ArrayList.
	list := arraylist.NewArrayList()

	// Add an element to the end of the list.
	list.Add("value")

	// Test
	if list.Get(0) != "value" {
		t.Error("Get() should return the element at the specified index")
	}
}

// TestSet
func TestSet(t *testing.T) {

	// Create a new ArrayList.
	list := arraylist.NewArrayList()

	// Add an element to the end of the list.
	list.Add("value")

	// Test
	if list.Get(0) != "value" {
		t.Error("Get() should return the element at the specified index")
	}

	// Set the element at the specified index.
	list.Set(0, "value2")

	// Test
	if list.Get(0) != "value2" {
		t.Error("Set() should set the element at the specified index")
	}
}

// TestRemove
func TestRemove(t *testing.T) {

	// Create a new ArrayList.
	list := arraylist.NewArrayList()

	// Add an element to the end of the list.
	list.Add("value")

	// Test
	if list.Size() != 1 {
		t.Error("Add() should increase the size of the list")
	}

	// Remove the element at the specified index.
	list.Remove(0)

	// Test
	if list.Size() != 0 {
		t.Error("Remove() should decrease the size of the list")
	}
}

// TestSize
func TestSize(t *testing.T) {

	// Create a new ArrayList.
	list := arraylist.NewArrayList()

	// Add an element to the end of the list.
	list.Add("value")

	// Test
	if list.Size() != 1 {
		t.Error("Add() should increase the size of the list")
	}
}

// TestClear
func TestClear(t *testing.T) {

	// Create a new ArrayList.
	list := arraylist.NewArrayList()

	// Add an element to the end of the list.
	list.Add("value")

	// Test
	if list.Size() != 1 {
		t.Error("Add() should increase the size of the list")
	}

	// Clear the list.
	list.Clear()

	// Test
	if list.Size() != 0 {
		t.Error("Clear() should decrease the size of the list")
	}
}

// TestAddMultiple
func TestAddMultiple(t *testing.T) {

	// Create a new ArrayList.
	list := arraylist.NewArrayList()

	// Add multiple elements to the end of the list.
	list.Add("value1")
	list.Add("value2")
	list.Add("value3")

	// Test
	if list.Size() != 3 {
		t.Error("Add() should increase the size of the list")
	}
}

// TestGetMultiple
func TestGetMultiple(t *testing.T) {

	// Create a new ArrayList.
	list := arraylist.NewArrayList()

	// Add multiple elements to the end of the list.
	list.Add("value1")
	list.Add("value2")
	list.Add("value3")

	// Test
	if list.Get(0) != "value1" {
		t.Error("Get() should return the element at the specified index")
	}
	if list.Get(1) != "value2" {
		t.Error("Get() should return the element at the specified index")
	}
	if list.Get(2) != "value3" {
		t.Error("Get() should return the element at the specified index")
	}
}

// TestSetMultiple
func TestSetMultiple(t *testing.T) {

	// Create a new ArrayList.
	list := arraylist.NewArrayList()

	// Add multiple elements to the end of the list.
	list.Add("value1")
	list.Add("value2")
	list.Add("value3")

	// Test
	if list.Get(0) != "value1" {
		t.Error("Get() should return the element at the specified index")
	}
	if list.Get(1) != "value2" {
		t.Error("Get() should return the element at the specified index")
	}
	if list.Get(2) != "value3" {
		t.Error("Get() should return the element at the specified index")
	}

	// Set the element at the specified index.
	list.Set(0, "value4")
	list.Set(1, "value5")
	list.Set(2, "value6")

	// Test
	if list.Get(0) != "value4" {
		t.Error("Set() should set the element at the specified index")
	}
	if list.Get(1) != "value5" {
		t.Error("Set() should set the element at the specified index")
	}
	if list.Get(2) != "value6" {
		t.Error("Set() should set the element at the specified index")
	}
}

// TestRemoveMultiple
func TestRemoveMultiple(t *testing.T) {

	// Create a new ArrayList.
	list := arraylist.NewArrayList()

	// Add multiple elements to the end of the list.
	list.Add("value1")
	list.Add("value2")
	list.Add("value3")

	// Test
	if list.Size() != 3 {
		t.Error("Add() should increase the size of the list")
	}

	// Remove the element at the specified index.
	list.Remove(0)
	list.Remove(0)
	list.Remove(0)

	// Test
	if list.Size() != 0 {
		t.Error("Remove() should decrease the size of the list")
	}
}

// TestClearMultiple
func TestClearMultiple(t *testing.T) {

	// Create a new ArrayList.
	list := arraylist.NewArrayList()

	// Add multiple elements to the end of the list.
	list.Add("value1")
	list.Add("value2")
	list.Add("value3")

	// Test
	if list.Size() != 3 {
		t.Error("Add() should increase the size of the list")
	}

	// Clear the list.
	list.Clear()

	// Test
	if list.Size() != 0 {
		t.Error("Clear() should decrease the size of the list")
	}
}
