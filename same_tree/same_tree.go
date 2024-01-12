package same_tree

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

/**
 * https://leetcode.com/problems/same-tree/
 *
 * The number of nodes in both trees is in the range [0, 100].
 * -10^4 <= Node.val <= 10^4
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }

    if p == nil && q != nil {
        return false
    }

    if p != nil && q == nil {
        return false
    }

    if p.Val != q.Val {
        return false
    }

    if !isSameTree(p.Left, q.Left) {
        return false
    }

    if !isSameTree(p.Right, q.Right) {
        return false
    }

    return true
}
