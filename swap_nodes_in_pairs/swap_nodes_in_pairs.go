package swap_nodes_in_pairs

type ListNode struct {
    Val  int
    Next *ListNode
}

/**
 * https://leetcode.com/problems/swap-nodes-in-pairs/
 *
 * Constraints:
 * The number of nodes in the list is in the range [0, 100].
 * 0 <= Node.val <= 100
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    node1 := head.Next
    head.Next = head.Next.Next
    node1.Next = head

    node2 := head

    for {
        if node2.Next == nil || node2.Next.Next == nil {
            break
        }

        node3 := node2.Next
        node2.Next = node2.Next.Next

        node4 := node2.Next.Next
        node2.Next.Next = node3
        node3.Next = node4

        node2 = node2.Next.Next
    }

    return node1
}
