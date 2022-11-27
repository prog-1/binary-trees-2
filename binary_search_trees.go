package main

import "fmt"

type node struct {
	val         int
	left, right *node
}

func LevelOrderTraversal(root *node, sink func(v int)) {
	//edge case if tree is empty
	if root == nil {
		return
	}

	s := []*node{} //declaring node queue
	s = append(s, root)
	for len(s) != 0 {
		cur := s[0]
		sink(cur.val)
		if cur.left != nil {
			s = append(s, cur.left)
		}
		if cur.right != nil {
			s = append(s, cur.right)
		}
		s = s[1:] //remove first el from queue
	}

}

func DeleteNode(root *node, k int) *node {

	//edge case
	if root == nil {
		return nil
	}
	//start := root// store the first node in the tree to return it

	//find node with this value (there couldn't be nodes with same value, because it's binary search tree)
	//check if node has childrens
	//delete if no childrens
	//if there are childrens - replace the value on the next bigger value and delete previous

	//Find node with input value
	if k < root.val { //node is in left side
		root.left = DeleteNode(root.left, k)
	} else if k > root.val { // done is in right side
		root.right = DeleteNode(root.right, k)
	} else { // k == root.val (we found the node to delete)

		//if it is leaf node with no childrens
		if root.left == nil && root.right == nil {
			//just delete the node
			root = nil
		} else if root.left == nil { // if we have only right children
			//change parent of the right node to parent of current root node
			//delete current root node
			root = root.right // change current node on child node
		} else if root.right == nil { // if only left children
			root = root.left
		} else { // root.left != nil && root.right != nil
			// if we have left and right childrens
			tmp := findMinNode(root.right)                   //find min right value from that node
			root.val = tmp.val                           // copy the value in current root
			root.right = DeleteNode(root.right, tmp.val) // delete copied node

		}
	}

	return root
}

//find the node with min value in binary search tree
func findMinNode(root *node) *node {
	for ; root.left != nil; root = root.left {
	}
	return root
}

func main() {
	tree := tree1()
	DeleteNode(tree, 1)
	DeleteNode(tree, 7)
	LevelOrderTraversal(tree, func(v int) { fmt.Print(v, " ") })
}

//7 5 9 1 6 8 10
func tree1() *node {
	return &node{val: 7,
		left: &node{val: 5,
			left:  &node{val: 1},
			right: &node{val: 6}},
		right: &node{val: 9,
			left:  &node{val: 8},
			right: &node{val: 10}}}
}

//7 5 9 1 10
func tree2() *node {
	return &node{val: 7,
		left: &node{val: 5,
			left: &node{val: 1}},
		right: &node{val: 9,
			right: &node{val: 10}}}
}

// 7
func singleElTree() *node {
	return &node{val: 7}
}

// nil
func nilTree() *node {
	return nil
}
