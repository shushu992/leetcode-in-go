package contains_duplicate

import "testing"

func TestContainsDuplicate1(t *testing.T) {
    input := []int{1}
    result := containsDuplicate(input)
    assert(t, result, false)
}

func TestContainsDuplicate2(t *testing.T) {
    input := []int{1, 1}
    result := containsDuplicate(input)
    assert(t, result, true)
}

func TestContainsDuplicate3(t *testing.T) {
    input := []int{1, 2, 969018, 1}
    result := containsDuplicate(input)
    assert(t, result, true)
}

func TestContainsDuplicate4(t *testing.T) {
    input := []int{1, 2, 3, 4}
    result := containsDuplicate(input)
    assert(t, result, false)
}

func TestContainsDuplicate5(t *testing.T) {
    input := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
    result := containsDuplicate(input)
    assert(t, result, true)
}

func TestContainsDuplicate6(t *testing.T) {
    input := []int{601205, 435287, 372065, 369781}
    result := containsDuplicate(input)
    assert(t, result, false)
}

func assert(t *testing.T, result bool, want bool) {
    t.Helper()

    if result != want {
        t.Fatalf("result: %v; want: %v", result, want)
    }
}
