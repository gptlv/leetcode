package main

import (
	"strconv"
	"strings"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := []string{}

	q := []*TreeNode{root}

	for len(q) > 0 {
		node := q[0]

		if node == nil {
			res = append(res, "nil")
		} else {
			res = append(res, strconv.Itoa(node.Val))
			q = append(q, node.Left, node.Right)
		}

		q = q[1:]
	}

	return strings.Join(res, ",")

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strSlice := strings.Split(data, ",")
	if strSlice[0] == "nil" {
		return nil
	}

	val, _ := strconv.Atoi(strSlice[0])
	root := &TreeNode{Val: val}
	q := []*TreeNode{root}
	idx := 1

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if strSlice[idx] != "nil" {
			val, _ := strconv.Atoi(strSlice[idx])
			node.Left = &TreeNode{Val: val}
			q = append(q, node.Left)
		}

		idx++

		if idx < len(strSlice) && strSlice[idx] != "nil" {
			val, _ := strconv.Atoi(strSlice[idx])
			node.Right = &TreeNode{Val: val}
			q = append(q, node.Right)
		}

		idx++

	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
