package main

type node struct {
	val         int
	left, right *node
}

func levelOrderTraversal(root *node, sink func(v int)) (res []int) {
	level := make(map[int][]int)
	var visit func(root *node, level int, levelMapNode *map[int][]int)
	visit = func(root *node, level int, levelMapNode *map[int][]int) {
		if root == nil {
			return
		}
		(*levelMapNode)[level] = append((*levelMapNode)[level], root.val)
		if root.left != nil {
			visit(root.left, level+1, levelMapNode)
		}
		if root.right != nil {
			visit(root.right, level+1, levelMapNode)
		}
	}
	visit(root, 0, &level)
	output := make([]int, 0)
	for key := 0; ; key++ {
		if _, ok := level[key]; !ok {
			break
		}
		output = append(output, level[key]...)
	}
	return output

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
			r := &root.right
			for (*r).left != nil {
				r = &(*r).left
			}
			root.val = (*r).val
			*r = nil
		}
	}
	return root
}
