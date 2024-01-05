package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * https://leetcode.com/problems/linked-list-cycle/description/
 *
 * Constraints:
 * The number of the nodes in the list is in the range [0, 10^4].
 * -10^5 <= Node.val <= 10^5
 * pos is -1 or a valid index in the linked-list.
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head.Next
	if slow == nil {
		return false
	}

	fast := slow.Next

	for {
		if fast == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next

		if fast == nil {
			return false
		}

		if slow == fast {
			return true
		}

		fast = fast.Next

		if slow == fast {
			return true
		}
	}
}
