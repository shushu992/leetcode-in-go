package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * https://leetcode.com/problems/reverse-linked-list/
 *
 * Constraints:
 * The number of nodes in the list is the range [0, 5000].
 * -5000 <= Node.val <= 5000
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	rHead := &ListNode{head.Val, nil}

	for {
		head = head.Next
		if head == nil {
			break
		}

		rHead = &ListNode{head.Val, rHead}
	}

	return rHead
}
