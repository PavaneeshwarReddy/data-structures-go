package binarytrees

/*
Problem: Minimum time taken to burn the tree from the target node
Link: https://takeuforward.org/data-structure/minimum-time-taken-to-burn-the-binary-tree-from-a-node
Idea:
	- It's similar to find nodes at k distance.
	- We need to create a map child - parent.
	- From the target node we need to bfs, where nodes are left, right, parent
	- We need to mark nodes as visited inorder to not visit them again
	- If there are any new nodes then we can simply increment
*/

func (bt *BinaryTree) MinTimeToBurn(target *Node) int {
	mp := make(map[*Node]*Node)
	mapChildParent(bt.Root, nil, &mp)
	t := int(0)
	nodes := []*Node{target}
	visited := make(map[*Node]bool)

	for len(nodes) > 0 {
		t += 1
		newNodes := []*Node{}
		for _, node := range nodes {
			if node.left != nil {
				if _, ok := visited[node.left]; !ok {
					newNodes = append(newNodes, node.left)
				}
			}

			if node.right != nil {
				if _, ok := visited[node.right]; !ok {
					newNodes = append(newNodes, node.right)
				}
			}

			if par, ok := mp[node]; ok {
				if _, ok := visited[par]; !ok {
					newNodes = append(newNodes, par)
				}
			}
			visited[node] = true
		}
		nodes = newNodes
	}
	return t
}
