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

func main() {
	tree := &node{val: 9, left: &node{val: 3, left: &node{val: 0}, right: &node{val: 7}}, right: &node{val: 15, left: &node{val: 14}, right: &node{val: 19}}}
	levelOrderTraversal(tree, func(v int) { fmt.Println(v) })
}
