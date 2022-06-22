package trees

// TreeNode represents a node in a tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode creates a new TreeNode
func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		Value: value,
	}
}

// Tree represents a tree
type Tree struct {
	Root *TreeNode
}

// NewTree creates a new Tree
func NewTree() *Tree {
	return &Tree{}
}

// Add adds a value to the tree
func (t *Tree) Add(value int) {
	if t.Root == nil {
		t.Root = NewTreeNode(value)
		return
	}

	t.Root.add(value)
}

// Traversal traverses the tree in-order
func (t *Tree) Traversal() {
	t.traversal(t.Root)
}

// traversal traverses the tree in-order
func (t *Tree) traversal(node *TreeNode) {
	if node == nil {
		return
	}

	t.traversal(node.Left)
	println(node.Value)
	t.traversal(node.Right)
}

// add adds a value to the tree
func (n *TreeNode) add(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = NewTreeNode(value)
			return
		}

		n.Left.add(value)
	} else {
		if n.Right == nil {
			n.Right = NewTreeNode(value)
			return
		}

		n.Right.add(value)
	}
}
