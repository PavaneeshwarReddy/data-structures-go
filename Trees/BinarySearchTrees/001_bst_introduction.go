package binarysearchtrees

import "fmt"

/*
BST
- Everything in LEFT subtree  <  node  <  Everything in RIGHT subtree
- Inorder gives the ascending order

Insert:
- Donot insert duplicate
- If root is greater than current key move to the left or else right
- Find is simply searching whole tree for the value
*/

type Node struct {
	Left  *Node
	Val   int
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (bst *BinarySearchTree) Insert(root *Node, key int) *Node {

	if root == nil {
		newNode := Node{Val: key}
		if bst.Root == nil {
			bst.Root = &newNode
		}
		return &newNode
	}

	if root.Val == key {
		return root
	}

	if root.Val > key {
		root.Left = bst.Insert(root.Left, key)
	} else {
		root.Right = bst.Insert(root.Right, key)
	}
	return root
}

func getSuccessor(root *Node) *Node {
	curr := root.Right
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

func (bst *BinarySearchTree) Delete(root *Node, key int) *Node {
	if root == nil {
		return root
	}

	if root.Val > key {
		root.Left = bst.Delete(root.Left, key)
	} else if root.Val < key {
		root.Right = bst.Delete(root.Right, key)
	} else {
		if root.Left == nil {
			rightNode := root.Right
			root.Right = nil
			return rightNode
		}
		if root.Right == nil {
			leftNode := root.Left
			root.Left = nil
			return leftNode
		}

		successor := getSuccessor(root)
		root.Val = successor.Val
		root.Right = bst.Delete(root.Right, root.Val)
	}

	return root
}

func (bst *BinarySearchTree) Search(root *Node, key int) bool {
	if root == nil {
		return false
	}

	if root.Val == key {
		return true
	}

	return bst.Search(root.Left, key) && bst.Search(root.Right, key)
}

func (bst *BinarySearchTree) InorderTraversal(root *Node) {
	if root == nil {
		return
	}

	bst.InorderTraversal(root.Left)
	fmt.Println(root.Val)
	bst.InorderTraversal(root.Right)
}
