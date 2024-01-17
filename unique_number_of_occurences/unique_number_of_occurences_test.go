package unique_number_of_occurences

import "testing"

func TestUniqueOccurrences1(t *testing.T) {
    input := []int{1}
    result := uniqueOccurrences(input)
    assert(t, result, true)
}

func TestUniqueOccurrences2(t *testing.T) {
    input := []int{1, 1}
    result := uniqueOccurrences(input)
    assert(t, result, true)
}

func TestUniqueOccurrences3(t *testing.T) {
    input := []int{1, 2}
    result := uniqueOccurrences(input)
    assert(t, result, false)
}

func TestUniqueOccurrences4(t *testing.T) {
    input := []int{1, 1, 2}
    result := uniqueOccurrences(input)
    assert(t, result, true)
}

func TestUniqueOccurrences5(t *testing.T) {
    input := []int{1, 1, 2, 2}
    result := uniqueOccurrences(input)
    assert(t, result, false)
}

func TestUniqueOccurrences6(t *testing.T) {
    input := []int{1, 2, 2, 1, 1, 3}
    result := uniqueOccurrences(input)
    assert(t, result, true)
}

func assert(t *testing.T, result bool, want bool) {
    t.Helper()

    if result != want {
        t.Fatalf("result: %v; want: %v", result, want)
    }
}
