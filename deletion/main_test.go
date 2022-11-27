package main

import (
	"testing"
)

func tree(s []int) *node { // creates Binary Search Tree
	if len(s) == 0 {
		return nil
	}
	t := insert(nil, s[0])
	for i := 1; i < len(s); i++ {
		t = insert(t, s[i])
	}
	return t
}

func TestDeleteNode(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		k    int
		want *node
	}{
		{name: "remove the only element",
			root: tree([]int{1}),
			k:    1,
			want: tree([]int{})},
		{name: "remove one of the two elements #1",
			root: tree([]int{1, 2}),
			k:    1,
			want: tree([]int{2})},
		{name: "remove one of the two elements #2",
			root: tree([]int{1, 2}),
			k:    2,
			want: tree([]int{1})},
		{name: "example 1",
			root: tree([]int{7, 5, 9, 1, 6, 8, 10}),
			k:    1,
			want: tree([]int{7, 5, 9, 6, 8, 10})},
		{name: "example 2",
			root: tree([]int{7, 5, 9, 1, 6, 8, 10}),
			k:    7,
			want: tree([]int{8, 5, 9, 1, 6, 10})},
		{name: "remove an element somewhere in the middle",
			root: tree([]int{7, 5, 9, 1, 6, 8, 10}),
			k:    5,
			want: tree([]int{7, 6, 9, 1, 8, 10})},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := deleteNode(tc.root, tc.k); !equal(got, tc.root) {
				t.Error("returned unexpected node")
			}
		})
	}
}

func equal(a, b *node) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else if a.val != b.val {
		return false
	} else if equal(a.left, b.left) && equal(a.right, b.right) {
		return true
	}
	return false
}
