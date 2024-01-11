package remove_duplicates_from_sorted_list

import "testing"

func TestDeleteDuplicates1(t *testing.T) {
    result := deleteDuplicates(nil)
    assert(t, result, nil)
}

func TestDeleteDuplicates2(t *testing.T) {
    input := makeList([]int{1, 1, 1, 1})

    result := deleteDuplicates(input)
    want := makeList([]int{1})

    assert(t, result, want)
}

func TestDeleteDuplicates3(t *testing.T) {
    input := makeList([]int{1, 2})

    result := deleteDuplicates(input)
    want := makeList([]int{1, 2})

    assert(t, result, want)
}

func TestDeleteDuplicates4(t *testing.T) {
    input := makeList([]int{1, 1, 2})

    result := deleteDuplicates(input)
    want := makeList([]int{1, 2})

    assert(t, result, want)
}

func TestDeleteDuplicates5(t *testing.T) {
    input := makeList([]int{1, 1, 2, 2})

    result := deleteDuplicates(input)
    want := makeList([]int{1, 2})

    assert(t, result, want)
}

func TestDeleteDuplicates6(t *testing.T) {
    input := makeList([]int{1, 1, 1, 2, 3, 3})

    result := deleteDuplicates(input)
    want := makeList([]int{1, 2, 3})

    assert(t, result, want)
}

func TestDeleteDuplicates7(t *testing.T) {
    input := makeList([]int{1, 1, 1, 2, 3, 3, 4, 5, 5, 5})

    result := deleteDuplicates(input)
    want := makeList([]int{1, 2, 3, 4, 5})

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
