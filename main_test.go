package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func BST1() *node {
	return &node{val: 9,
		left: &node{
			val:   3,
			left:  &node{val: 0},
			right: &node{val: 7},
		},
		right: &node{
			val:   15,
			left:  &node{val: 14},
			right: &node{val: 19},
		}}
}

func TestLevelOrderTravelsal(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		want []int
	}{
		{"1", BST1(), []int{9, 3, 15, 0, 7, 14, 19}},
		// to be continued
	} {
		var got []int
		levelOrderTraversal(tc.root, func(v int) { got = append(got, v) })
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("levelOrderTraversal(#%v) = %v, want %v", tc.name, got, tc.want)
		}
	}
}

func isBST(t *node) bool {
	if (t.left == nil || t.left.val < t.val) && (t.right == nil || t.right.val > t.val) {
		return (t.left == nil || isBST(t.left)) && (t.right == nil || isBST(t.right))
	}
	return false
}

func TestDeleteNode(t *testing.T) {
	for _, tc := range []struct {
		name string
		root func() *node
	}{
		{"1", BST1},
	} {
		var elements []int
		levelOrderTraversal(tc.root(), func(v int) { elements = append(elements, v) })
		for _, element := range elements {
			tree := tc.root()
			deleteNode(tree, element)
			var new_elements []int
			levelOrderTraversal(tc.root(), func(v int) { elements = append(new_elements, v) })
			if isBST(tree) || Contains(new_elements, element) {
				var buf bytes.Buffer
				printTree(tree, &buf)
				t.Errorf("this modified tree is not BST or contains removed element %v: \n%v", element, buf.String())
			}
		}
	}
}

func printTree(n *node, w io.Writer) {
	printer(n, w, nil)
}

func printNonemptyHistory(prev []int, w io.Writer, printLast bool) {
	if len(prev) == 0 {
		panic("must not happen")
	}
	last := len(prev)
	if !printLast {
		last--
	}
	for _, p := range prev[:last] {
		if p > 0 {
			fmt.Fprint(w, "| ")
		} else {
			fmt.Fprint(w, "  ")
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

func printer(n *node, w io.Writer, prev []int) {
	if n == nil {
		return
	}
	if len(prev) > 0 {
		printNonemptyHistory(prev, w, true /*printLast*/)
		fmt.Fprintln(w)
		printNonemptyHistory(prev, w, false /*printLast*/)
		prev[len(prev)-1]--
	}
	fmt.Fprintln(w, "+->", n.val)
	prev = append(prev, directChildren(n))
	if n.left != nil {
		printer(n.left, w, prev)
	}
	if n.right != nil {
		printer(n.right, w, prev)
	}
}

func Contains(s []int, a int) bool {
	for _, v := range s {
		if a == v {
			return true
		}
	}
	return false
}
