package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * https://leetcode.com/problems/symmetric-tree/
 *
 * Constraints:
 * The number of nodes in the tree is in the range [1, 1000].
 * -100 <= Node.val <= 100
 */
func isSymmetric(root *TreeNode) bool {
	arr := [2048]*TreeNode{root}
	size := 1

	for depth := 2; depth <= 10; depth++ {
		for i := size - 1; i >= 0; i-- {
			if arr[i] == nil {
				arr[i*2+1] = nil
				arr[i*2] = nil
			} else {
				arr[i*2+1] = arr[i].Right
				arr[i*2] = arr[i].Left
			}
		}

		size <<= 1

		for i := 0; i < size/2; i++ {
			if arr[i] == nil && arr[size-i-1] == nil {
				continue
			}

			if arr[i] == nil && arr[size-i-1] != nil {
				return false
			}

			if arr[i] != nil && arr[size-i-1] == nil {
				return false
			}

			if arr[i].Val != arr[size-i-1].Val {
				return false
			}
		}
	}

	return true
}
