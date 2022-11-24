# Binary search trees

In all tasks, assume that a binary search tree is defined as a composition of
nodes represented by the following structure:

```golang
type node struct {
    val int
    left, right *node
}
```

# Level order traversal

Implement a function `func levelOrderTraversal(root *node, sink func(v int))` (+tests) that
performs a level order traversal, i.e.

1. a root node (a node with height 0) is visited;
2. nodes with height 1 are visited from left to right;
3. nodes with height 2 are visited from left to right;
4. etc.

`levelOrderTraversal` must have the worst-time $O(n)$ time complexity and must use
$O(n)$ memory.

## Example

For the following tree

```
       7
    /     \
   5       9
 /   \   /   \
1     6 8    10
```

the level traversal is `7 5 9 1 6 8 10`.


# Deletion

Implement a function `func deleteNode(root *node, k int) *node` (+tests) that deletes a
node with a given key `k` from a binary search tree with a root node `root`.

`deleteNode` returns a node that is a root of the resulting tree. After the
deletion, the tree must represent a binary search tree.

`deleteNode` must have the worst-time $O(\log n)$ time complexity and must use
$O(\log n)$ memory.

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