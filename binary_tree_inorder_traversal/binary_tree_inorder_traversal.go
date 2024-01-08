package binary_tree_inorder_traversal

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type LinkedNode struct {
    Val  int
    Next *LinkedNode
}

/**
 * https://leetcode.com/problems/binary-tree-inorder-traversal/
 *
 * The number of nodes in the tree is in the range [0, 100].
 * -100 <= Node.val <= 100
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    head := makeLinkedNode(root)

    arr := [100]int{}
    size := 0

    for {
        if head == nil {
            break
        }

        arr[size] = head.Val

        head = head.Next
        size++
    }

    return arr[:size]
}

// todo return both of head and tail to save time
func makeLinkedNode(root *TreeNode) *LinkedNode {
    head := &LinkedNode{root.Val, nil}

    if root.Right != nil {
        head.Next = makeLinkedNode(root.Right)
    }

    if root.Left != nil {
        lHead := makeLinkedNode(root.Left)

        lTail := lHead
        for {
            if lTail.Next == nil {
                break
            }
            lTail = lTail.Next
        }
        lTail.Next = head

        head = lHead
    }

    return head
}
