package binarytrees
import "fmt"

/*
Levelorder traversal
Print all the nodes of each level
*/


func (bt *BinaryTree) LevelorderTraversal(nodes []*Node) {
	if len(nodes) == 0 {
		return
	}

	newNodes := []*Node{}
	
	for _,node := range nodes {
		fmt.Print(node.value, " ")
		if node.left != nil {
			newNodes = append(newNodes, node.left)
		}
		if node.right != nil {
			newNodes = append(newNodes, node.right)
		}
	}
	fmt.Println()

	bt.LevelorderTraversal(newNodes)
}
