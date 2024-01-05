package leetcode_in_go

import (
	"testing"
)

func TestMergeTwoLists1(t *testing.T) {
	list1 := makeList([]int{1, 2, 4})
	list2 := makeList([]int{1, 3, 4})

	merged := mergeTwoLists(list1, list2)
	expected := makeList([]int{1, 1, 2, 3, 4, 4})

	assertList(t, merged, expected)
}

func TestMergeTwoLists2(t *testing.T) {
	list1 := makeList([]int{})
	list2 := makeList([]int{})

	merged := mergeTwoLists(list1, list2)
	expected := makeList([]int{})

	assertList(t, merged, expected)
}

func TestMergeTwoLists3(t *testing.T) {
	list1 := makeList([]int{})
	list2 := makeList([]int{0})

	merged := mergeTwoLists(list1, list2)
	expected := makeList([]int{0})

	assertList(t, merged, expected)
}

func TestMergeTwoLists4(t *testing.T) {
	list1 := makeList([]int{0})
	list2 := makeList([]int{})

	merged := mergeTwoLists(list1, list2)
	expected := makeList([]int{0})

	assertList(t, merged, expected)
}

func makeList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	var head = &ListNode{arr[0], nil}

	var tail = head
	for _, val := range arr[1:] {
		tail.Next = &ListNode{val, nil}
		tail = tail.Next
	}

	return head
}

func assertList(t *testing.T, result *ListNode, expected *ListNode) {
	t.Helper()

	if result == nil && expected == nil {
		return
	}

	if result == nil && expected != nil {
		t.FailNow()
	}

	if result != nil && expected == nil {
		t.FailNow()
	}

	for {
		if result.Val != expected.Val {
			t.FailNow()
		}

		if result.Next == nil && expected.Next == nil {
			return
		}

		if result.Next == nil && expected.Next != nil {
			t.FailNow()
		}

		if result.Next != nil && expected.Next == nil {
			t.FailNow()
		}

		result = result.Next
		expected = expected.Next
	}
}
