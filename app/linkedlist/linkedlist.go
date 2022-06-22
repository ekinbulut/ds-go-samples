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
