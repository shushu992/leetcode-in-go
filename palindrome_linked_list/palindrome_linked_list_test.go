package palindrome_linked_list

import "testing"

func TestPalindromeLinkedList1(t *testing.T) {
    input := makeList(1)
    result := isPalindrome(input)
    assert(t, result, true)
}

func TestPalindromeLinkedList2(t *testing.T) {
    input := makeList(1, 1)
    result := isPalindrome(input)
    assert(t, result, true)
}

func TestPalindromeLinkedList3(t *testing.T) {
    input := makeList(1, 2)
    result := isPalindrome(input)
    assert(t, result, false)
}

func TestPalindromeLinkedList4(t *testing.T) {
    input := makeList(1, 1, 1)
    result := isPalindrome(input)
    assert(t, result, true)
}

func TestPalindromeLinkedList5(t *testing.T) {
    input := makeList(1, 1, 2)
    result := isPalindrome(input)
    assert(t, result, false)
}

func TestPalindromeLinkedList6(t *testing.T) {
    input := makeList(1, 2, 2, 1)
    result := isPalindrome(input)
    assert(t, result, true)
}

func TestPalindromeLinkedList7(t *testing.T) {
    input := makeList(1, 2, 1, 2)
    result := isPalindrome(input)
    assert(t, result, false)
}

func TestPalindromeLinkedList8(t *testing.T) {
    input := makeList(1, 3, 0, 2)
    result := isPalindrome(input)
    assert(t, result, false)
}

func TestPalindromeLinkedList9(t *testing.T) {
    input := makeList(8, 0, 7, 1, 7, 7, 9, 7, 5, 2, 9, 1, 7, 3, 7, 0,
        6, 5, 1, 7, 7, 9, 3, 8, 1, 5, 7, 7, 8, 4, 0, 9,
        3, 7, 3, 4, 5, 7, 4, 8, 8, 5, 8, 9, 8, 5, 8, 8,
        4, 7, 5, 4, 3, 7, 3, 9, 0, 4, 8, 7, 7, 5, 1, 8,
        3, 9, 7, 7, 1, 5, 6, 0, 7, 3, 7, 1, 9, 2, 5, 7,
        9, 7, 7, 1, 7, 0, 8)
    result := isPalindrome(input)
    assert(t, result, true)
}

func makeList(arr ...int) *ListNode {
    head := &ListNode{arr[0], nil}
    tail := head

    for _, v := range arr[1:] {
        tail.Next = &ListNode{v, nil}
        tail = tail.Next
    }

    return head
}

func assert(t *testing.T, result bool, want bool) {
    t.Helper()

    if result != want {
        t.Fatalf("want: %v; result: %v", want, result)
    }
}
