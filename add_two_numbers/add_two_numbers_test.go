package add_two_numbers

import "testing"

func TestAddTwoNumbers1(t *testing.T) {
	l1 := &ListNode{0, nil}
	l2 := &ListNode{0, nil}

	result := addTwoNumbers(l1, l2)
	want := &ListNode{0, nil}

	assert(t, result, want)
}

func TestAddTwoNumbers2(t *testing.T) {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{9, nil}

	result := addTwoNumbers(l1, l2)
	want := &ListNode{0, &ListNode{1, nil}}

	assert(t, result, want)
}

func TestAddTwoNumbers3(t *testing.T) {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{9, &ListNode{9, nil}}

	result := addTwoNumbers(l1, l2)
	want := &ListNode{0, &ListNode{0, &ListNode{1, nil}}}

	assert(t, result, want)
}

func TestAddTwoNumbers4(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	result := addTwoNumbers(l1, l2)
	want := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}

	assert(t, result, want)
}

func assert(t *testing.T, result *ListNode, want *ListNode) {
	t.Helper()

	for {
		if result == nil && want == nil {
			return
		}

		if result == nil && want != nil {
			t.Fatal()
		}

		if result != nil && want == nil {
			t.Fatal()
		}

		// assert values
		if result.Val != want.Val {
			t.Fatal()
		}

		result = result.Next
		want = want.Next
	}
}
