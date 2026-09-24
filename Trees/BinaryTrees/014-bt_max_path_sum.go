package binarytrees

import "fmt"

/*
Max path sum
- This has to be the same thought as diameter, we can either choose left sum or right sum if it negative we don't choose at all
- We calculate and return, we don't return the result, we only return the sum possible
*/

func maxSum(root *Node, result *int) int {
	if root == nil {
		return 0
	}

	leftSum := max(0, maxSum(root.left, result))
	rightSum := max(0, maxSum(root.right, result))

	*result = max(*result, leftSum+rightSum+root.value)

	return max(leftSum, rightSum) + root.value
}

func (bt *BinaryTree) MaxpathSum() {
	result := -100000
	maxSum(bt.Root, &result)
	fmt.Println(result)
}
