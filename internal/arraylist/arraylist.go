package arraylist

import "fmt"

type ArrayList struct {
	data []interface{}
	size int
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		data: make([]interface{}, 0),
		size: 0,
	}
}

// add an element to the end of the list
func (list *ArrayList) Add(element interface{}) {
	list.data = append(list.data, element)
	list.size++
}

// get the element at the specified index
func (list *ArrayList) Get(index int) interface{} {
	if index < 0 || index >= list.size {
		return nil
	}
	return list.data[index]
}

// set the element at the specified index
func (list *ArrayList) Set(index int, element interface{}) {
	if index < 0 || index >= list.size {
		return
	}
	list.data[index] = element
}

// remove the element at the specified index
func (list *ArrayList) Remove(index int) interface{} {
	if index < 0 || index >= list.size {
		return nil
	}
	element := list.data[index]
	list.data = append(list.data[:index], list.data[index+1:]...)
	list.size--
	return element
}

// get the size of the list
func (list *ArrayList) Size() int {
	return list.size
}

// clear the list
func (list *ArrayList) Clear() {
	list.data = make([]interface{}, 0)
	list.size = 0
}

// get the index of the first occurrence of the specified element
func (list *ArrayList) IndexOf(element interface{}) int {
	for i, e := range list.data {
		if e == element {
			return i
		}
	}
	return -1
}

// get the index of the last occurrence of the specified element
func (list *ArrayList) LastIndexOf(element interface{}) int {
	for i := list.size - 1; i >= 0; i-- {
		if list.data[i] == element {
			return i
		}
	}
	return -1
}

// check if the list contains the specified element
func (list *ArrayList) Contains(element interface{}) bool {
	return list.IndexOf(element) != -1
}

// check if the list contains all of the elements in the specified collection
func (list *ArrayList) ContainsAll(collection []interface{}) bool {
	for _, element := range collection {
		if !list.Contains(element) {
			return false
		}
	}
	return true
}

// check if the list contains any of the elements in the specified collection
func (list *ArrayList) ContainsAny(collection []interface{}) bool {
	for _, element := range collection {
		if list.Contains(element) {
			return true
		}
	}
	return false
}

// check if the list is empty
func (list *ArrayList) IsEmpty() bool {
	return list.size == 0
}

// print the list
func (list *ArrayList) Print() {
	for _, element := range list.data {
		fmt.Println(element)
	}
}
