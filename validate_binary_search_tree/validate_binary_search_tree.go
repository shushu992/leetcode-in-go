package validate_binary_search_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * https://leetcode.com/problems/validate-binary-search-tree/
 *
 * The number of nodes in the tree is in the range [1, 10^4].
 * -2^31 <= Node.val <= 2^31 - 1
 *
 */
func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt32, math.MaxInt32)
}

func validate(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val < min {
		return false
	}

	if root.Val > max {
		return false
	}

    if !validate(root.Left, min, root.Val -1) {
        return false
    }

    if !validate(root.Right, root.Val +1, max){
        return false
    }

    return true
}
