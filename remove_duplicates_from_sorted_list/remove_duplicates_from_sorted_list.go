package remove_duplicates_from_sorted_list

type ListNode struct {
    Val  int
    Next *ListNode
}

/**
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/
 *
 * The number of nodes in the list is in the range [0, 300].
 * -100 <= Node.val <= 100
 * The list is guaranteed to be sorted in ascending order.
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    for node := head; ; {
        if node.Next == nil {
            break
        }

        if node.Next.Val > node.Val {
            node = node.Next
            continue
        }

        node.Next = node.Next.Next
    }

    return head
}
