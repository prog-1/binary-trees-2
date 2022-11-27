package main

type node struct {
	val         int
	left, right *node
}

func levelOrderTraversal(root *node, sink func(v int)) {
	if root == nil {
		return
	}
	var s []*node
	s = append(s, root)
	for len(s) != 0 {
		n := s[0]
		s = s[1:]
		sink(n.val)
		if n.left != nil {
			s = append(s, n.left)
		}
		if n.right != nil {
			s = append(s, n.right)
		}
	}
}

func deleteNode(root *node, k int) *node {
	if root == nil {
		return nil
	}
	if k < root.val {
		root.left = deleteNode(root.left, k)
	} else if k > root.val {
		root.right = deleteNode(root.right, k)
	} else {
		if root.left == nil && root.right == nil {
			return nil
		} else if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			tmp := root.right
			for tmp.left != nil {
				tmp = tmp.left
			}
			root.val = tmp.val
			root.right = deleteNode(root.right, root.val)
		}
	}
	return root
}
