package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * https://leetcode.com/problems/merge-two-sorted-lists/
 *
 * Constraints:
 * The number of nodes in both lists is in the range [0, 50].
 * -100 <= Node.val <= 100
 * Both list1 and list2 are sorted in non-decreasing order.
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil && list2 != nil {
		return list2
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	head := &ListNode{0, nil}
	last := head

	if list1.Val <= list2.Val {
		last.Val = list1.Val
		list1 = list1.Next
	} else {
		last.Val = list2.Val
		list2 = list2.Next
	}

	for {
		if list1 == nil && list2 == nil {
			break
		}

		if list1 == nil && list2 != nil {
			last.Next = &ListNode{list2.Val, nil}
			last = last.Next
			list2 = list2.Next
			continue
		}

		if list1 != nil && list2 == nil {
			last.Next = &ListNode{list1.Val, nil}
			last = last.Next
			list1 = list1.Next
			continue
		}

		if list1.Val <= list2.Val {
			last.Next = &ListNode{list1.Val, nil}
			last = last.Next
			list1 = list1.Next
		} else {
			last.Next = &ListNode{list2.Val, nil}
			last = last.Next
			list2 = list2.Next
		}
	}

	return head
}
