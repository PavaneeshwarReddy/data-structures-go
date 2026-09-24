package binarytrees

/*
Find nodes at distance k from target node
- We can get easily the child nodes but we won't get the parent nodes
- To get that we need a map that marks the node -> parent node

- After creating this map, there still an issue, there is a chance of visiting the same node again
- So we need to check with the previous node, whether we can pickup left and right of the current node
- Traverse until k == 0 then you can add to the result array
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
