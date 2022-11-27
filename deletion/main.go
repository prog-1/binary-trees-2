package main

import "fmt"

type node struct {
	val         int
	left, right *node
}

func insert(t *node, a int) *node {
	if t == nil {
		return &node{val: a}
	}
	if a < t.val {
		t.left = insert(t.left, a)

	} else {
		t.right = insert(t.right, a)
	}
	return t
}

func find(t, parent *node, a int) (*node, *node) {
	if t == nil || t.val == a {
		return t, parent
	}
	if a < t.val {
		return find(t.left, t, a)
	}
	return find(t.right, t, a)
}

func nextNode(t *node) *node {
	if t.left == nil && t.right == nil {
		return nil
	} else if t.left == nil {
		return t.right
	} else if t.right == nil {
		return t.left
	} else {
		tmp := t.right
		var parent *node
		for ; tmp.left != nil; tmp, parent = tmp.left, tmp {
		}
		t.val = tmp.val
		if t.val == t.right.val {
			t.right = nil
		} else {
			parent.left = nil
		}
		return t
	}
}

func deleteNode(root *node, k int) *node {
	el, parent := find(root, root, k)
	if parent == el {
		nextNode(el)
	} else if parent.left == el {
		parent.left = nextNode(el)
	} else {
		parent.right = nextNode(el)
	}
	return root
}

func main() {
	t := insert(nil, 7)
	t = insert(t, 5)
	t = insert(t, 9)
	t = insert(t, 1)
	t = insert(t, 6)
	t = insert(t, 8)
	t = insert(t, 10)
	printTree(t)
	fmt.Println()
	t = deleteNode(t, 5)
	printTree(t)
}

func printTree(n *node) {
	printer(n, nil)
}

func printNonemptyHistory(prev []int, printLast bool) {
	if len(prev) == 0 {
		panic("must not happen")
	}
	last := len(prev)
	if !printLast {
		last--
	}
	for _, p := range prev[:last] {
		if p > 0 {
			fmt.Print("| ")
		} else {
			fmt.Print("  ")
		}
	}
}

func directChildren(n *node) (num int) {
	if n == nil {
		return 0
	}
	if n.left != nil {
		num++
	}
	if n.right != nil {
		num++
	}
	return num
}

func printer(n *node, prev []int) {
	if n == nil {
		return
	}
	if len(prev) > 0 {
		printNonemptyHistory(prev, true /*printLast*/)
		fmt.Println()
		printNonemptyHistory(prev, false /*printLast*/)
		prev[len(prev)-1]--
	}
	fmt.Println("+->", n.val)
	prev = append(prev, directChildren(n))
	if n.left != nil {
		printer(n.left, prev)
	}
	if n.right != nil {
		printer(n.right, prev)
	}
}
