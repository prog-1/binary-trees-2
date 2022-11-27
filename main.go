package main

import (
	"fmt"
)

type node struct {
	val         int
	left, right *node
}

func levelOrderTraversal(root *node, sink func(v int)) {
	for i := 0; i < height(root)+1; i++ {
		traverse(root, i, sink)
	}
}

func traverse(root *node, i int, sink func(v int)) {
	if root == nil {
		return
	}
	if i == 0 {
		sink(root.val)
	}
	if i > 0 {
		traverse(root.left, i-1, sink)
		traverse(root.right, i-1, sink)
	}
}

func height(root *node) int {
	if root == nil || (root.left == nil && root.right == nil) {
		return 0
	}

	if height(root.left) > height(root.right) {
		return height(root.left) + 1
	} else {
		return height(root.right) + 1
	}
}
func main() {
	tree := &node{val: 9, left: &node{val: 3, left: &node{val: 0}, right: &node{val: 7}}, right: &node{val: 15, left: &node{val: 14}, right: &node{val: 19}}}
	levelOrderTraversal(tree, func(v int) { fmt.Println(v) })
}
