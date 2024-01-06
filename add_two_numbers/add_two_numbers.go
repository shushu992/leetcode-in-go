package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * https://leetcode.com/problems/add-two-numbers/
 * The number of nodes in each linked list is in the range [1, 100].
 * 0 <= Node.val <= 9
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{l1.Val + l2.Val, nil}

	carry := 0
	tail := result

	if tail.Val >= 10 {
        tail.Val = tail.Val - 10
		carry = 1
	}

	for {
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil && carry == 0 {
			break
		}

		tail.Next = &ListNode{0, nil}
		tail = tail.Next

		if l1 != nil {
			tail.Val += l1.Val
		}

		if l2 != nil {
			tail.Val += l2.Val
		}

		tail.Val += carry
		carry = 0

		if tail.Val >= 10 {
			tail.Val = tail.Val - 10
			carry = 1
		}
	}

	return result
}
