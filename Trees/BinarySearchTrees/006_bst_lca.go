package binarysearchtrees

/*
LCA
- There is a chance that both p and q can be found on same, or opposite side
- For every iteration we need to check for both on both sides
- The moment we encounter that 2 values are found then we can set the result
- First result result mill be the minimum, as it look like post order

*/

func lca(root *Node, p *Node, q *Node, result **Node) (bool, bool) {
	if root == nil {
		return false, false
	}

	leftP, leftQ := lca(root.Left, p, q, result)
	rightP, rightQ := lca(root.Right, p, q, result)

	pFound := leftP || rightP || root == p
	qFound := leftQ || rightQ || root == q

	if pFound && qFound && *result == nil {
		*result = root
	}

	return pFound, qFound
}

func lowestCommonAncestor(root, p, q *Node) *Node {
	var result *Node
	lca(root, p, q, &result)
	return result
}
