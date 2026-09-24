package binarytrees

/*
Right / Left view
- If we require right we traverse completely to the right or left
*/

func (bt *BinaryTree) LeftView() {
	rightTreeTraversal(bt.Root)
}
