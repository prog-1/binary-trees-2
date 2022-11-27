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

func main() {
	root := n(10).
		l(n(5).
			l(n(2)).
			r(n(7))).
		r(n(15).
			l(n(1)).
			r(n(4)))

	levelOrderTraversal(root, func(v int) { fmt.Println(v) }) // 10 5 15 2 7 1 4
}
