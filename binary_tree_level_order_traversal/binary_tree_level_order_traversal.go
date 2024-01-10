package binary_tree_level_order_traversal

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

/**
 * https://leetcode.com/problems/binary-tree-level-order-traversal/
 *
 * The number of nodes in the tree is in the range [0, 2000].
 * -1000 <= Node.val <= 1000
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    arr := [2048][]int{}
    size := 0

    for nodes := []*TreeNode{root}; len(nodes) != 0; {
        arr1 := [1024]int{}
        size1 := 0

        nodes2 := [1024]*TreeNode{}
        size2 := 0

        for _, node := range nodes {
            arr1[size1] = node.Val
            size1++

            if node.Left != nil {
                nodes2[size2] = node.Left
                size2++
            }

            if node.Right != nil {
                nodes2[size2] = node.Right
                size2++
            }
        }

        arr[size] = arr1[:size1]
        size++

        nodes = nodes2[:size2]
    }

    return arr[:size]
}
