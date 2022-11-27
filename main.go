package main

import "fmt"

type node struct {
	val         int
	left, right *node
}

func inorderTraversal(root *node, elements *map[int][]int, level int) int {
	if root == nil {
		return level - 1
	}
	(*elements)[level] = append((*elements)[level], root.val)
	a := inorderTraversal(root.left, elements, level+1)
	b := inorderTraversal(root.right, elements, level+1)
	if a > b {
		return a
	}
	return b
}

func levelOrderTraversal(root *node, sink func(v int)) {
	elements := make(map[int][]int)
	level := inorderTraversal(root, &elements, 0)
	for i := 0; i <= level; i++ {
		for _, num := range elements[i] {
			sink(num)
		}
	}
}

func deleteNode(root *node, k int) *node {
	if root == nil {
		return nil
	}
	if k < root.val {
		deleteNode(root.left, k)
		return root
	}
	if k > root.val {
		deleteNode(root.right, k)
		return root
	}
	if root.left == nil && root.right == nil {
		root = nil
		return root
	}
	if root.left == nil {
		root = root.right
		return root
	}
	if root.right == nil {
		root = root.left
		return root
	}
	tmp := &root.right
	for (*tmp).left != nil {
		tmp = &(*tmp).left
	}
	root.val = (*tmp).val
	*tmp = (*tmp).right
	return root
}

func main() {
	tree := &node{val: 9, left: &node{val: 3, left: &node{val: 0}, right: &node{val: 7}}, right: &node{val: 15, left: &node{val: 14}, right: &node{val: 19}}}
	levelOrderTraversal(tree, func(v int) { fmt.Println(v) })
	deleteNode(tree, 15)
	levelOrderTraversal(tree, func(v int) { fmt.Println(v) })

}
