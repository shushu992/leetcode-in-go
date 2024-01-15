package remove_nth_node_from_end_of_list

type ListNode struct {
    Val  int
    Next *ListNode
}

/**
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/
 *
 * 1 <= n <= sz <= 30
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var x *ListNode = nil // parent of the target node to be removed

    interval := 1
    for y := head; y.Next != nil; y = y.Next {
        if interval < n {
            interval++
            continue
        }

        if x == nil {
            x = head
        } else {
            x = x.Next
        }
    }

    if x == nil {
        return head.Next
    }

    if x.Next.Next == nil {
        x.Next = nil
    } else {
        x.Next = x.Next.Next
    }

    return head
}
