package linked_list_cycle

import "testing"

func TestLinkedListCycle1(t *testing.T) {
	result := hasCycle(nil)
	assert(t, result, false)
}

func TestLinkedListCycle2(t *testing.T) {
	input := makeList([]int{1}, -1)
	result := hasCycle(input)
	assert(t, result, false)
}

func TestLinkedListCycle3(t *testing.T) {
	input := makeList([]int{1}, 0)
	result := hasCycle(input)
	assert(t, result, true)
}

func TestLinkedListCycle4(t *testing.T) {
	input := makeList([]int{3, 2, 0, -4}, -1)
	result := hasCycle(input)
	assert(t, result, false)
}

func TestLinkedListCycle5(t *testing.T) {
	input := makeList([]int{3, 2, 0, -4}, 0)
	result := hasCycle(input)
	assert(t, result, true)
}

func TestLinkedListCycle6(t *testing.T) {
	input := makeList([]int{5, 4, -9, 3}, 1)
	result := hasCycle(input)
	assert(t, result, true)
}

func TestLinkedListCycle7(t *testing.T) {
	input := makeList([]int{5, 4, -9, 3}, 3)
	result := hasCycle(input)
	assert(t, result, true)
}

func TestLinkedListCycle8(t *testing.T) {
	input := makeList([]int{1, 2}, -1)
	result := hasCycle(input)
	assert(t, result, false)
}

func TestLinkedListCycle9(t *testing.T) {
	input := makeList([]int{1, 2}, 0)
	result := hasCycle(input)
	assert(t, result, true)
}

func TestLinkedListCycle10(t *testing.T) {
	input := makeList([]int{1, 2}, 1)
	result := hasCycle(input)
	assert(t, result, true)
}

func makeList(arr []int, pos int) *ListNode {
	head := &ListNode{arr[0], nil}
	tail := head

	for _, v := range arr[1:] {
		tail.Next = &ListNode{v, nil}
		tail = tail.Next
	}

	if pos >= 0 {
		tail.Next = head
		for i := 1; i <= pos; i++ {
			tail.Next = tail.Next.Next
		}
	}

	return head
}

func assert(t *testing.T, result bool, want bool) {
	t.Helper()

	if result != want {
		t.Fatalf("want: %v; got: %v", want, result)
	}
}
