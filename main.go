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

	type list struct {
		val  *node
		next *list
	}
	insertHead := func(l *list, val *node) *list {
		n := &list{val, l}
		return n
	}
	removeHead := func(l *list) *list {
		l = l.next
		return l
	}

	if root == nil {
		return
	}

	q := insertHead(nil, root)

	cur := root
	for q != nil {
		cur = q.val
		sink(q.val.val)
		q = removeHead(q)

		if cur.right != nil {
			q = insertHead(q, cur.right)
		}
		if cur.left != nil {
			q = insertHead(q, cur.left)
		}
	}

}

func main() {
	root := n(10).
		l(n(5).
			l(n(2)).
			r(n(7))).
		r(n(15))

	levelOrderTraversal(root, func(v int) { fmt.Println(v) }) // 10 5 15 2 7
}
