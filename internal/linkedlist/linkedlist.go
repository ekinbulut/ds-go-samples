package linkedlist

type Node struct {
	Next  *Node
	Value int
}

func NewNode(value int) *Node {
	return &Node{
		Next:  nil,
		Value: value,
	}
}

type LinkedList struct {
	Node *Node
}

func NewLinkedList(value int) *LinkedList {
	return &LinkedList{
		Node: NewNode(value),
	}
}

// Add adds a new node to the end of the list.
func (ll *LinkedList) Add(value int) {
	node := NewNode(value)
	if ll.Node == nil {
		ll.Node = node
	} else {
		current := ll.Node
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

// Print prints the list.
func (ll *LinkedList) Print() {
	current := ll.Node
	for current != nil {
		println(current.Value)
		current = current.Next
	}
}

// Size returns the size of the list.
func (ll *LinkedList) Size() int {
	size := 0
	current := ll.Node
	for current != nil {
		size++
		current = current.Next
	}
	return size
}

// Get returns the value of the node at the given index.
func (ll *LinkedList) Get(index int) int {
	current := ll.Node
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}
