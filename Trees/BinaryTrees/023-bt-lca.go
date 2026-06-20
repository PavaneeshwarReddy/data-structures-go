package binarytrees

import "fmt"

/*
LCA
*/

func lca(root *Node, p *Node, q *Node, result **Node) (bool, bool) {
	if root == nil {
		return false, false
	}

	leftP, leftQ := lca(root.left, p, q, result)
	rightP, rightQ := lca(root.right, p, q, result)

	pFound := leftP || rightP || root == p
	qFound := leftQ || rightQ || root == q

	if pFound && qFound && *result == nil {
		*result = root
	}

	return pFound, qFound
}

func (bt *BinaryTree) LCA(p *Node, q *Node) {
	var result *Node
	lca(bt.Root, p, q, &result)
	fmt.Println(result)
}
