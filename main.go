package main

import (
	binaryTrees "dsa-go/Trees/BinaryTrees"
	"fmt"
)

func main() {
	bt := binaryTrees.BinaryTree{}

	bt.InsertRoot(2)

	bt.InsertNodeLeft(bt.Root, 23)
	bt.InsertNodeRight(bt.Root, 54)

	bt.PreorderTraversal(bt.Root)

	fmt.Println()

	bt.InorderTraversal(bt.Root)

	fmt.Println()

	bt.PostorderTraversal(bt.Root)

	fmt.Println()

	bt.LevelorderTraversal([]*binaryTrees.Node{bt.Root})

	fmt.Println()

	bt.PreorderIterative()

	fmt.Println()

	bt.InorderIterative()

	fmt.Println()

	bt.PostordrerIterativeTwoStacks()

	fmt.Println()

	bt.AllTraversal()

	fmt.Println()

	fmt.Println(bt.MaxDepth(bt.Root, 0))

	fmt.Println(bt.Diameter())

	bt.MaxpathSum()

	fmt.Println()

	bt.SpiralTraversal()

}
