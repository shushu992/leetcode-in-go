package maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * https://leetcode.com/problems/maximum-depth-of-binary-tree/
 *
 * Constraints:
 * The number of nodes in the tree is in the range [0, 104].
 * -100 <= Node.val <= 100
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return findDepth(root, 1)
}

func findDepth(root *TreeNode, sum int) int {
	l := sum
	if root.Left != nil {
		l = findDepth(root.Left, sum+1)
	}

	r := sum
	if root.Right != nil {
		r = findDepth(root.Right, sum+1)
	}

	if l >= r {
		return l
	}

	return r
}
