package reverse_linked_list

import "testing"

func TestReverseLinkedList1(t *testing.T) {
	input := makeList([]int{1, 2, 3, 4, 5})
	result := reverseList(input)
	want := makeList([]int{5, 4, 3, 2, 1})
	assert(t, result, want)
}

func TestReverseLinkedList2(t *testing.T) {
	input := makeList([]int{1, 2})
	result := reverseList(input)
	want := makeList([]int{2, 1})
	assert(t, result, want)
}

func TestReverseLinkedList3(t *testing.T) {
	result := reverseList(nil)
	assert(t, result, nil)
}

func makeList(arr []int) *ListNode {
	head := &ListNode{arr[0], nil}
	tail := head

	for _, v := range arr[1:] {
		tail.Next = &ListNode{v, nil}
		tail = tail.Next
	}

	return head
}

func assert(t *testing.T, result *ListNode, want *ListNode) {
	t.Helper()

	if result == nil && want == nil {
		return
	}

	if result == nil && want != nil {
		t.Fatal()
	}

	if result != nil && want == nil {
		t.Fatal()
	}

	for {
		if result.Val != want.Val {
			t.Fatal()
		}

		if result.Next == nil && want.Next == nil {
			return
		}

		if result.Next == nil && want.Next != nil {
			t.Fatal()
		}

		if result.Next != nil && want.Next == nil {
			t.Fatal()
		}

		result = result.Next
		want = want.Next
	}
}
