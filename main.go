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
