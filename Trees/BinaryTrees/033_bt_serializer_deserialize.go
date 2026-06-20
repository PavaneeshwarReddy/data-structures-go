package binarytrees

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	Seperator string
	NilStr    string
}

func Constructor() Codec {
	return Codec{
		Seperator: "@",
		NilStr:    "$",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *Node) string {
	nodes := []*Node{root}
	result := ""

	if root == nil {
		return result
	}

	for len(nodes) > 0 {
		newNodes := []*Node{}
		for _, node := range nodes {

			if result != "" {
				result += this.Seperator
			}
			if node == nil {
				result += this.NilStr
				continue
			}
			result += strconv.Itoa(node.value)
			newNodes = append(newNodes, node.left, node.right)

		}
		nodes = newNodes
	}
	return result
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *Node {

	if data == "" {
		return nil
	}

	var root *Node
	splits := strings.Split(data, this.Seperator)
	num64, _ := strconv.ParseInt(splits[0], 10, 64)
	root = &Node{value: int(num64)}
	nodes := []*Node{root}
	idx := 1
	fmt.Println(data)
	for len(nodes) > 0 {
		newNodes := []*Node{}
		for _, node := range nodes {

			if idx >= len(splits) {
				break
			}

			if splits[idx] != this.NilStr {
				num64, _ := strconv.ParseInt(splits[idx], 10, 64)
				node.left = &Node{value: int(num64)}
				newNodes = append(newNodes, node.left)
			}
			idx += 1
			if idx >= len(splits) {
				break
			}
			if splits[idx] != this.NilStr {
				num64, _ := strconv.ParseInt(splits[idx], 10, 64)
				node.right = &Node{value: int(num64)}
				newNodes = append(newNodes, node.right)
			}
			idx += 1
		}
		nodes = newNodes
	}

	return root

}
