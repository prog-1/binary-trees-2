package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	name    string
	getTree func() (root *node)
	want    []int
}

var testCases = []testCase{
	{"nil", func() (root *node) {
		return nil
	}, nil},
	{"incomplete tree", func() *node {
		root := n(10).
			l(n(5).
				l(n(2)).
				r(n(7))).
			r(n(15))

		return root
	}, []int{10, 5, 15, 2, 7}},
	{"complete tree", func() *node {
		root := n(10).
			l(n(5).
				l(n(2)).
				r(n(7))).
			r(n(15).
				l(n(1)).
				r(n(4)))

		return root
	}, []int{10, 5, 15, 2, 7, 1, 4}},
}

func TestNCA(t *testing.T) {
	for _, tc := range testCases {
		root := tc.getTree()
		var s []int
		levelOrderTraversal(root, func(v int) { s = append(s, v) })
		if !reflect.DeepEqual(s, tc.want) {
			t.Errorf("got = %v, want = %v", s, tc.want)
		}
	}
}
