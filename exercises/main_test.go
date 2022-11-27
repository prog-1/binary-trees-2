package main

import (
	"reflect"
	"testing"
)

var (
	tree  = &node{val: 3, left: &node{val: 1, right: &node{val: 2}}, right: &node{val: 5, left: &node{val: 4}}}
	tree2 = &node{val: 4, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 6, left: &node{val: 5}, right: &node{val: 7}}}
)

func TestLevelOrderTraversal(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		want []int
	}{
		{"nil", nil, nil},
		{"small", &node{val: 1, right: &node{val: 2}}, []int{1, 2}},
		{"small2", &node{val: 1, left: &node{val: 2}, right: &node{val: 3}}, []int{1, 2, 3}},
		{"large", tree2, []int{4, 2, 6, 1, 3, 5, 7}},
		{"incomplete", tree, []int{3, 1, 5, 2, 4}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			levelOrderTraversal(tc.root, func(v int) { got = append(got, v) })
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Test: %v, got = %v, want = %v", tc.name, got, tc.want)
			}
		})
	}
}

func TestDeleteNode(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		k    int
		want *node
	}{
		{"nil", nil, 0, nil},
		{"one node", &node{val: 1}, 1, nil},
		{"delete first node", tree, 3, &node{val: 4, left: &node{val: 1, right: &node{val: 2}}, right: &node{val: 5}}},
		{"delete first node 2", tree2, 4, &node{val: 5, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 6, right: &node{val: 7}}}},
		{"last node", tree2, 7, &node{val: 4, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 6, left: &node{val: 5}}}},
		{"middle", tree2, 6, &node{val: 4, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 7, left: &node{val: 5}}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := deleteNode(tc.root, tc.k); !equal(got, tc.want) {
				t.Errorf("Test: %v, got = %v, want = %v", tc.name, got, tc.want)
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
