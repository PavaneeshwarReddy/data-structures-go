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
	LevelorderTraversal(nodes []*Node)
	PreorderIterative()
	InorderIterative()
	PostorderTraversalOneStack()
	PostordrerIterativeTwoStacks()
	AllTraversal()

	// views
	LeftView()
	RightView()
	TopView()
	BottomView()
	VerticalOrderTraversal() [][]int

	// Problems
	MaxDepth(node *Node, depth int) int
	CheckBalancedTree() bool
	Diameter() int
	MaxpathSum()
	CheckSimilarity(root1 *Node, root2 *Node) bool
	SpiralTraversal()
	BoundaryTraversal()
	Symmetric() bool
	LCA(p *Node, q *Node)
	FindDistanceKNodes(target *Node, k int) []int
	ChildrenSumProperty()
	MaxWidth() int
	CountNodes() int
	ConstructPreIn(preorder []int, inorder []int) *Node
	ConstructInPos(postorder []int, inorder []int) *Node
	FlattenTree()
	MinTimeToBurn(target *Node) int
	MorrisPreorder() []int
	MorrisInorder() []int
	RootToLeafPath(key int)
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
