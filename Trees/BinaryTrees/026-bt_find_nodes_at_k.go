package binarytrees

/*
Find nodes at distance k from target node
*/

func mapChildParent(root *Node, pNode *Node, mp *map[*Node]*Node) {
	if root == nil {
		return
	}

	mapChildParent(root.left, root, mp)
	if pNode != nil {
		(*mp)[root] = pNode
	}
	mapChildParent(root.right, root, mp)
}

func findK(currNode *Node, prevNode *Node, k int, mp map[*Node]*Node, result *[]int) {
	if currNode == nil {
		return
	}
	if k == 0 {
		*result = append(*result, currNode.value)
	}

	if currNode.left != prevNode {
		findK(currNode.left, currNode, k-1, mp, result)
	}
	if currNode.right != prevNode {
		findK(currNode.right, currNode, k-1, mp, result)
	}
	if parentNode, ok := mp[currNode]; ok && parentNode != currNode {
		findK(parentNode, currNode, k-1, mp, result)
	}
}

func (bt *BinaryTree) FindDistanceKNodes(target *Node, k int) []int {
	mp := make(map[*Node]*Node)
	mapChildParent(bt.Root, nil, &mp)
	result := []int{}
	findK(target, nil, k, mp, &result)
	return result
}
