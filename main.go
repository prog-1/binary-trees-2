package main

import "fmt"

type node struct {
	val         int
	left, right *node
}

func n(v int) *node             { return &node{val: v} }
func (n *node) l(l *node) *node { n.left = l; return n }
func (n *node) r(r *node) *node { n.right = r; return n }

func levelOrderTraversal(root *node, sink func(v int)) {
	// Requirements: Space/Time Complexity = O(n)

	if root == nil {
		return
	}

	var s []*node
	s = append(s, root)
	for len(s) != 0 {
		cur := &s[0]
		sink((*cur).val)
		s = s[1:] // deleting first element

		if (*cur).left != nil {
			s = append(s, (*cur).left)
		}
		if (*cur).right != nil {
			s = append(s, (*cur).right)
		}
	}
}

func deleteNode(root *node, k int) *node {
	if root == nil {
		return nil
	}
	// How to delete node:
	// 1. Find the node in the tree
	// 2. If it's a leaf node, then set pointer to it to nil
	// 3. If it has the only child, then replace pointer to it with a pointer to a child
	// 4. If it has two children, then replace it with child from the right part with the smallest value

	p := &root
	for (*p) != nil && k < (*p).val {
		p = &(*p).left
	}
	for (*p) != nil && k > (*p).val {
		p = &(*p).right
	}
	if (*p) == nil {
		return nil
	}

	if (*p).left == nil && (*p).right == nil {
		(*p) = nil
		return root
	}
	if (*p).left != nil && (*p).right == nil {
		(*p) = (*p).left
		return root
	}
	if (*p).left == nil && (*p).right != nil {
		(*p) = (*p).right
		return root
	}

	sr := &(*p).right // sr - smallest right
	for (*sr).left != nil {
		sr = &(*sr).left
	}
	(*p).val = (*sr).val
	(*sr) = nil
	return root
}

func print(root *node) {
	if root == nil {
		return
	}
	fmt.Println(root.val)
	print(root.left)
	print(root.right)
}

func main() {
	root := n(10).
		l(n(5).
			l(n(2)).
			r(n(7)))
	// r(n(15).
	// 	l(n(1)).
	// 	r(n(4)))
	print(root)
	fmt.Println()
	deleteNode(root, 5)
	print(root) // 10 7 2
}
