package binarytrees

/*
Problem: Construct Binary Tree from  Inorder and Postorder
link: https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
Idea:
	- inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
	- If we carefully observe, if we choose 3 as root , in inorder left of 3 is entire left subtree and right is right subtree
	- When the pointer in postorder shifts actually it's right side of subtree
	- from that we can consider recursively by just iterating preorder and finding it's respective in inorder
*/

func constructInPo(postorderIdx *int, postorder []int, start int, end int, mp map[int]int) *Node {
	if start > end {
		return nil
	}

	nodeVal := postorder[*postorderIdx]
	*postorderIdx -= 1
	node := Node{value: nodeVal}

	inorderIdx := mp[nodeVal]

	node.right = constructPreIn(postorderIdx, postorder, inorderIdx+1, end, mp)
	node.left = constructPreIn(postorderIdx, postorder, start, inorderIdx-1, mp)

	return &node

}

func (bt *BinaryTree) ConstructInPos(postorder []int, inorder []int) *Node {
	mp := make(map[int]int)
	for idx, val := range inorder {
		mp[val] = idx
	}
	postorderIdx := len(postorder) - 1
	return constructPreIn(&postorderIdx, postorder, 0, len(inorder)-1, mp)
}
