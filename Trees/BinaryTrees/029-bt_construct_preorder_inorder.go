package binarytrees

/*
Problem: Construct Binary Tree from Preorder and inorder
link: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
Idea:
	- preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	- If we carefully observe, if we choose 3 as root , in inorder left of 3 is entire left subtree and right is right subtree
	- from that we can consider recursively by just iterating preorder and finding it's respective in inorder
*/

func constructPreIn(preorderIdx *int, preorder []int, start int, end int, mp map[int]int) *Node {
	if start > end {
		return nil
	}

	nodeVal := preorder[*preorderIdx]
	*preorderIdx += 1
	node := Node{value: nodeVal}

	inorderIdx := mp[nodeVal]

	node.left = constructPreIn(preorderIdx, preorder, start, inorderIdx-1, mp)
	node.right = constructPreIn(preorderIdx, preorder, inorderIdx+1, end, mp)

	return &node

}

func (bt *BinaryTree) ConstructPreIn(preorder []int, inorder []int) *Node {
	mp := make(map[int]int)
	for idx, val := range inorder {
		mp[val] = idx
	}
	preorderIdx := 0
	return constructPreIn(&preorderIdx, preorder, 0, len(inorder)-1, mp)
}
