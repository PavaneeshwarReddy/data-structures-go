package binarytrees

import "fmt"

/*
LCA
- We cannot guarantee that these 2 nodes exisits in the tree so if root == nil we simply return false, false
- There combinations like p and q can be found on left tree or right tree or one on this and one on that
- We return pLeft, qLeft and pRight, qRight and we maintain them, if found we simply mark as true if result is not nil then we can store this value as return
- Because this can happen only once, if we found more also we need not store because we want least comming ancestor
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
