package binarytrees

/*
Right / Left view
*/

func (bt *BinaryTree) LeftView() {
	rightTreeTraversal(bt.Root)
}
