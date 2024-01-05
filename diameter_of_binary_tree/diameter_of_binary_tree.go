package diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * https://leetcode.com/problems/diameter-of-binary-tree/
 *
 * The number of nodes in the tree is in the range [1, 10^4].
 * -100 <= Node.val <= 100
 */
func diameterOfBinaryTree(root *TreeNode) int {
	return findLongest(root, 0)
}

func findLongest(node *TreeNode, top int) int {
	left := 0
	if node.Left != nil {
		left = findDepth(node.Left, 1)
	}

	right := 0
	if node.Right != nil {
		right = findDepth(node.Right, 1)
	}

	if top >= left && left >= right {
		return top + left
	}
	if top >= right && right >= left {
		return top + right
	}

	left2 := 0
	if node.Left != nil && top >= right {
		left2 = findLongest(node.Left, 1+top)
	}
	if node.Left != nil && top < right {
		left2 = findLongest(node.Left, 1+right)
	}

	right2 := 0
	if node.Right != nil && top >= left {
		right2 = findLongest(node.Right, 1+top)
	}
	if node.Right != nil && top < left {
		right2 = findLongest(node.Right, 1+left)
	}

	if left2 >= right2 {
		return left2
	}

	return right2
}

func findDepth(node *TreeNode, sum int) int {
	l := sum
	if node.Left != nil {
		l = findDepth(node.Left, sum+1)
	}

	r := sum
	if node.Right != nil {
		r = findDepth(node.Right, sum+1)
	}

	if l >= r {
		return l
	}

	return r
}
