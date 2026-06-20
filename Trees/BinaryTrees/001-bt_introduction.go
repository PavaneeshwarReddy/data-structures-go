package binarytrees

type Node struct {
	left  *Node
	value int
	right *Node
}

type BinaryTree struct {
	Root *Node
}

type BinaryTreeInterface interface {
	// operations
	InsertRoot(value int)
	InsertNodeLeft(node *Node, value int)
	InsertNodeRight(node *Node, value int)

	// traversals
	PreorderTraversal(node *Node)
	InorderTraversal(node *Node)
	PostorderTraversal(node *Node)
	LevelorderTraversal(node *Node)
	PreorderIterative()
	InorderIterative()
	PostorderIterative()
	AllTraversal()

	// Problems
	MaxDepth(node *Node, depth int) int
	CheckBalancedTree() bool
	Diameter() int
	MaxpathSum()
	CheckSimilarity(root1 *Node, root2 *Node) bool
	SpiralTraversal()
	BoundaryTraversal()
}

func (bt *BinaryTree) InsertRoot(value int) {
	bt.Root = &Node{value: value}
}

func (bt *BinaryTree) InsertNodeLeft(node *Node, value int) {
	node.left = &Node{value: value}
}

func (bt *BinaryTree) InsertNodeRight(node *Node, value int) {
	node.right = &Node{value: value}
}
