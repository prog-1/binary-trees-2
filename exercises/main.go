package main

import "fmt"

type node struct {
	val         int
	left, right *node
}

func levelOrderTraversal(root *node, sink func(v int)) {
	//
}

func deleteNode(root *node, k int) *node {
	if root == nil {
		return root
	}
	if root.val > k {
		root.left = deleteNode(root.left, k)
		return root
	} else if root.val < k {
		root.right = deleteNode(root.right, k)
		return root
	}
	if root.left == nil {
		return root.right
	} else if root.right == nil {
		return root.left
	}
	parent := root
	res := root.right
	for res.left != nil {
		parent = res
		res = res.left
	}
	if parent != root {
		parent.left = res.right
	} else {
		parent.right = res.right
	}
	root.val = res.val
	return root
}

func main() {
	tree := &node{val: 4, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 6, left: &node{val: 5}, right: &node{val: 7}}}
	levelOrderTraversal(tree, func(v int) { fmt.Println(v) })
	res := deleteNode(tree, 6)
	levelOrderTraversal(res, func(v int) { fmt.Println(v) })
}
