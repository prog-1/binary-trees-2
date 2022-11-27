package main

type node struct {
	val         int
	left, right *node
}

func LevelOrderTraversal(root *node, sink func(v int)) {
	if root == nil {
		return
	}
	s := []*node{}
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
		s = s[1:]
	}

}

func DeleteNode(root *node, k int) *node {
	if root == nil {
		return nil
	}
	if k < root.val {
		root.left = DeleteNode(root.left, k)
	} else if k > root.val {
		root.right = DeleteNode(root.right, k)
	} else {
		if root.left == nil && root.right == nil {
			root = nil
		} else if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			tmp := findMinNode(root.right)
			root.val = tmp.val
			root.right = DeleteNode(root.right, tmp.val)
		}
	}
	return root
}

func findMinNode(root *node) *node {
	for ; root.left != nil; root = root.left {
	}
	return root
}

func main() {
}
