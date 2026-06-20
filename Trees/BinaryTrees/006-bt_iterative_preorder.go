package binarytrees
import "fmt"


/*
Iterative Preorder Traversal
*/


func (bt *BinaryTree) PreorderIterative() {
	nodes := []*Node{bt.Root}

	for len(nodes) > 0 {
		topNode := nodes[len(nodes)-1]
		nodes = nodes[0:len(nodes)-1]

		fmt.Print(topNode.value, " ")

		if topNode.right != nil {
			nodes = append(nodes,topNode.right)
		}
		if topNode.left != nil {
			nodes = append(nodes, topNode.left)
		}
	}
	fmt.Println()
}

