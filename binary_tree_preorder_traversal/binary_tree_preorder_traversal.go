package binary_tree_preorder_traversal

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

/**
 * https://leetcode.com/problems/binary-tree-preorder-traversal/
 *
 * Definition:
 * https://www.geeksforgeeks.org/preorder-traversal-of-binary-tree/
 *
 * Constraints:
 * The number of nodes in the tree is in the range [0, 100].
 * -100 <= Node.val <= 100
 */
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    nums := [128]int{}
    size1 := 0

    nodes := [256]*TreeNode{root}
    size2 := 1

    for {
        if size2 == 0 {
            break
        }

        size2--
        node := nodes[size2]

        nums[size1] = node.Val
        size1++

        if node.Right != nil {
            nodes[size2] = node.Right
            size2++
        }

        if node.Left != nil {
            nodes[size2] = node.Left
            size2++
        }
    }

    return nums[0:size1]
}
