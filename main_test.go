package main

import (
	"reflect"
	"testing"
)

func TestLevelOrderTraversal(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		want []int
	}{
		{"case-1", nil, nil},
		{"case-2", &node{val: 7}, []int{7}},
		{"case-3", &node{val: 7, left: &node{val: 5}}, []int{7, 5}},
		{"case-4", &node{val: 7, right: &node{val: 9}}, []int{7, 9}},
		{"case-5", &node{val: 7, left: &node{val: 5}, right: &node{val: 9}}, []int{7, 5, 9}},
		{"case-6", &node{val: 7, left: &node{val: 5, left: &node{val: 1}, right: &node{val: 6}}, right: &node{val: 9}}, []int{7, 5, 9, 1, 6}},
		{"case-7", &node{val: 7, left: &node{val: 5, left: &node{val: 1}, right: &node{val: 6}}, right: &node{val: 9, left: &node{val: 8}, right: &node{val: 10}}}, []int{7, 5, 9, 1, 6, 8, 10}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			levelOrderTraversal(tc.root, func(v int) { got = append(got, v) })
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestDeleteNode(t *testing.T) {
	for _, tc := range []struct {
		name       string
		k          int
		root, want *node
	}{
		{"case-1", 1, tree(), &node{val: 7, left: &node{val: 5}, right: &node{val: 9, left: &node{val: 8}, right: &node{val: 10}}}},
		{"case-2", 7, tree(), &node{val: 8, left: &node{val: 5, left: &node{val: 1}}, right: &node{val: 9, right: &node{val: 10}}}},
		{"case-3", 5, tree(), &node{val: 7, left: &node{val: 1}, right: &node{val: 9, left: &node{val: 8}, right: &node{val: 10}}}},
		{"case-4", 9, tree(), &node{val: 7, left: &node{val: 5, left: &node{val: 1}}, right: &node{val: 10, left: &node{val: 8}}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := deleteNode(tc.root, tc.k); !equal(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func tree() *node {
	return &node{val: 7, left: &node{val: 5, left: &node{val: 1}}, right: &node{val: 9, left: &node{val: 8}, right: &node{val: 10}}}
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
