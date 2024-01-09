package palindrome_linked_list

type ListNode struct {
    Val  int
    Next *ListNode
}

/**
 * https://leetcode.com/problems/palindrome-linked-list/
 *
 * Constraints:
 * The number of nodes in the list is in the range [1, 10^5].
 * 0 <= Node.val <= 9
 *
 * Follow up:
 * Could you do it in O(n) time and O(1) space?
 */
func isPalindrome(head *ListNode) bool {
    sum := 0
    mul := 10
    // todo non-liner sequence of 10^5 items

    for tail := head; tail != nil; tail = tail.Next {
        sum += tail.Val * mul
        mul++
    }

    for tail := head; tail != nil; tail = tail.Next {
        mul--
        sum -= tail.Val * mul
    }

    return sum == 0
}
