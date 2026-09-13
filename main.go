package main

import (
	bst "dsa-go/Trees/BinarySearchTrees"
)

func main() {

	bst := bst.BinarySearchTree{}
	bst.Insert(bst.Root, -2)
	bst.Insert(bst.Root, 10)
	bst.Insert(bst.Root, -5)

	bst.Delete(bst.Root, -2)
	bst.InorderTraversal(bst.Root)

}
