package remove_nth_node_from_end_of_list

import "testing"

func TestRemoveNthFromEnd1(t *testing.T) {
    head := makeList([]int{1})

    result := removeNthFromEnd(head, 1)

    assert(t, result, nil)
}

func TestRemoveNthFromEnd2(t *testing.T) {
    head := makeList([]int{1, 2})

    result := removeNthFromEnd(head, 1)
    want := makeList([]int{1})

    assert(t, result, want)
}

func TestRemoveNthFromEnd3(t *testing.T) {
    head := makeList([]int{1, 2})

    result := removeNthFromEnd(head, 2)
    want := makeList([]int{2})

    assert(t, result, want)
}

func TestRemoveNthFromEnd4(t *testing.T) {
    head := makeList([]int{1, 2, 3, 4, 5})

    result := removeNthFromEnd(head, 1)
    want := makeList([]int{1, 2, 3, 4})

    assert(t, result, want)
}

func TestRemoveNthFromEnd5(t *testing.T) {
    head := makeList([]int{1, 2, 3, 4, 5})

    result := removeNthFromEnd(head, 2)
    want := makeList([]int{1, 2, 3, 5})

    assert(t, result, want)
}

func TestRemoveNthFromEnd6(t *testing.T) {
    head := makeList([]int{1, 2, 3, 4, 5})

    result := removeNthFromEnd(head, 5)
    want := makeList([]int{2, 3, 4, 5})

    assert(t, result, want)
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
