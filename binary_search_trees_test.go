package main

import (
	"reflect"
	"testing"
)

func TestLevelOrderTraversal(t *testing.T) {
	for _, tc := range []struct {
		name string
		tree *node
		want []int
	}{
		{"tree 1", tree1(), []int{7, 5, 9, 1, 6, 8, 10}},
		{"tree 2", tree2(), []int{7, 5, 9, 1, 10}},
		{"single el tree", singleElTree(), []int{7}},
		{"nil tree", nilTree(), nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			LevelOrderTraversal(tc.tree, func(v int) { got = append(got, v) })
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestDeleteNode(t *testing.T) {
	for _, tc := range []struct {
		name string
		tree *node
		k    int
		want []int
	}{
		{"leaf node", tree1(), 1, []int{7, 5, 9, 6, 8, 10}},
		{"left children", tree2(), 5, []int{7, 1, 9, 10}},
		{"right children", tree2(), 9, []int{7, 5, 10, 1}},
		{"2 children", tree1(), 9, []int{7, 5, 10, 1, 6, 8}},
		{"root node", tree1(), 7, []int{8, 5, 9, 1, 6, 10}},
		{"single el", singleElTree(), 7, nil},
		{"nil", nilTree(), 7, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tree := DeleteNode(tc.tree, tc.k)
			var got []int
			LevelOrderTraversal(tree, func(v int) { got = append(got, v) })
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
