package main

import (
	"reflect"
	"testing"
)

func tree1() (root, a, b, nca *node) {
	a = nil
	b = nil
	nca = nil
	root = nil
	return
}
func tree2() (root, a, b, nca *node) {
	a = &node{val: 5, left: &node{val: 9, left: &node{val: 7}}, right: &node{val: 10}}
	b = &node{val: 8, left: &node{val: 123, left: &node{val: 23}, right: &node{val: 10}}, right: &node{val: 10}}
	nca = &node{val: 2, left: &node{val: 4, right: b}, right: a}
	root = &node{val: 1, left: nca, right: &node{val: 3}}
	return
}

func tree3() (root, a, b, nca *node) {
	a = &node{val: 5, left: &node{val: 9}, right: &node{val: 10}}
	b = &node{val: 8, left: &node{val: 7, left: a, right: &node{val: 7}}, right: &node{val: 3}}
	nca = b
	root = &node{val: 1, left: &node{val: 1, left: &node{val: 9}, right: nca}, right: &node{val: 3}}
	return
}
func TestLevelOrderTravelsal(t *testing.T) {
	for _, tc := range []struct {
		name string
		fnc  func() (*node, *node, *node, *node)
		want []int
	}{
		{"test1", tree1, []int{}},
		{"test2", tree2, []int{1, 2, 3, 4, 5, 8, 9, 10, 123, 10, 7, 23, 10}},
		{"test3", tree3, []int{1, 1, 3, 9, 8, 7, 3, 5, 7, 9, 10}},
	} {

		t.Run(tc.name, func(t *testing.T) {
			root, _, _, _ := tc.fnc()
			if got := levelOrderTraversal(root, func(v int) {}); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
