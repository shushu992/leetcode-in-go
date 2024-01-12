package merge_sorted_array

import (
    "reflect"
    "testing"
)

func TestMerge1(t *testing.T) {
    nums1 := []int{1}
    m := 1
    nums2 := []int{}
    n := 0

    merge(nums1, m, nums2, n)
    want := []int{1}

    assert(t, nums1, want)
}

func TestMerge2(t *testing.T) {
    nums1 := []int{0}
    m := 0
    nums2 := []int{1}
    n := 1

    merge(nums1, m, nums2, n)
    want := []int{1}

    assert(t, nums1, want)
}

func TestMerge3(t *testing.T) {
    nums1 := []int{1, 2, 3, 0, 0, 0}
    m := 3
    nums2 := []int{2, 5, 6}
    n := 3

    merge(nums1, m, nums2, n)
    want := []int{1, 2, 2, 3, 5, 6}

    assert(t, nums1, want)
}

func TestMerge4(t *testing.T) {

}

func TestMerge5(t *testing.T) {

}

func assert(t *testing.T, result []int, want []int) {
    t.Helper()

    if !reflect.DeepEqual(result, want) {
        t.Fatalf("result: %v; want: %v", result, want)
    }
}
