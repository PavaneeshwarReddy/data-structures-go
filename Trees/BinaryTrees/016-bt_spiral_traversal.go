package binarytrees

import "fmt"

/*
Spiral Traversal
*/

func (bt *BinaryTree) SpiralTraversal() {
	nodes := []*Node{bt.Root}
	isReversed := false
	result := [][]int{}

	for len(nodes) > 0 {
		tempResult := []int{}
		newNodes := []*Node{}

		for i := 0; i < len(nodes); i++ {
			if isReversed {
				tempResult = append(tempResult, nodes[len(nodes)-1-i].value)
			} else {
				tempResult = append(tempResult, nodes[i].value)
			}
			if nodes[i].left != nil {
				newNodes = append(newNodes, nodes[i].left)
			}
			if nodes[i].right != nil {
				newNodes = append(newNodes, nodes[i].right)
			}
		}
		nodes = newNodes
		result = append(result, tempResult)
		isReversed = !isReversed
	}

	for _, v1 := range result {
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}

}
