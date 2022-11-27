package main

type node struct {
	val         int
	left, right *node
}

func levelOrderTraversal(root *node, sink func(v int)) {
	for i := 0; i < height(root)+1; i++ {
		traverse(root, i, sink)
	}
}

func traverse(root *node, i int, sink func(v int)) {
	if root == nil {
		return
	}
	if i == 0 {
		sink(root.val)
	}
	if i > 0 {
		traverse(root.left, i-1, sink)
		traverse(root.right, i-1, sink)
	}
}

func height(root *node) int {
	if root == nil || (root.left == nil && root.right == nil) {
		return 0
	}

	if height(root.left) > height(root.right) {
		return height(root.left) + 1
	} else {
		return height(root.right) + 1
	}
}

func deleteNode(root *node, k int) *node {
	if root == nil {
		return root
	}
	if k < root.val {
		root.left = deleteNode(root.left, k)
	} else if k > root.val {
		root.right = deleteNode(root.right, k)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			root.val = min(root.right)
			root.right = deleteNode(root.right, root.val)
		}
	}
	return root
}

func min(root *node) int {
	for root.left != nil {
		root = root.left
	}
	return root.val
}
