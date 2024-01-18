package contains_duplicate_ii

import "testing"

func TestContainsNearbyDuplicate1(t *testing.T) {
    nums := []int{1, 2, 3, 1}
    k := 3

    result := containsNearbyDuplicate(nums, k)
    assert(t, result, true)

}

func TestContainsNearbyDuplicate2(t *testing.T) {
    nums := []int{1, 0, 1, 1}
    k := 1

    result := containsNearbyDuplicate(nums, k)
    assert(t, result, true)
}

func TestContainsNearbyDuplicate3(t *testing.T) {
    nums := []int{1, 2, 3, 1, 2, 3}
    k := 2

    result := containsNearbyDuplicate(nums, k)
    assert(t, result, false)
}

func assert(t *testing.T, result bool, want bool) {
    t.Helper()

    if result != want {
        t.Fatalf("result: %v; want: %v", result, want)
    }
}
