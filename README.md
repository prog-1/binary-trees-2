# Binary search trees

In all tasks, assume that a binary search tree is defined as a composition of
nodes represented by the following structure:

```golang
type node struct {
    val int
    left, right *node
}
```

# Postorder traversal

Implement a function `func postorderTraversal(root *node, sink func(v int))` (+tests) that
performs a preorder traversal, i.e.

1. visits the left subtree,
2. visits the right subtree,
3. visits the current node.

## Example

For the following tree

```
       7
    /     \
   5       9
 /   \   /   \
1     6 8    10
```

the postorder traversal is `1 6 5 8 10 9 7`.


# Deletion

Implement a function `func deleteNode(root *node, k int) *node` (+tests) that deletes a
node with a given key `k` from a binary search tree with a root node `root`.

`deleteNode` returns a node that is a root of the resulting tree. After the
deletion, the tree must represent a binary search tree.

`deleteNode` must have the worst-time $O(\log n)$ time complexity and must use
$O(1)$ memory.

## Examples

```
       7                                7
    /     \       deleteNode(1)      /     \
   5       9       --------->       5       9
 /   \   /   \                       \    /   \ 
1     6 8    10                       6  8    10
```

```
       7                              8
    /     \       deleteNode(7)    /     \
   5       9       --------->     5       9
 /   \   /   \                  /   \       \ 
1     6 8    10                1     6      10
```