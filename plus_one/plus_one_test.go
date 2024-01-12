package plus_one

import (
    "reflect"
    "testing"
)

func TestPlusOne1(t *testing.T) {
    input := []int{0}

    result := plusOne(input)
    want := []int{1}

    assert(t, result, want)
}

func TestPlusOne2(t *testing.T) {
    input := []int{9}

    result := plusOne(input)
    want := []int{1, 0}

    assert(t, result, want)
}

func TestPlusOne3(t *testing.T) {
    input := []int{1, 2, 3}

    result := plusOne(input)
    want := []int{1, 2, 4}

    assert(t, result, want)
}

func TestPlusOne4(t *testing.T) {
    input := []int{9, 9, 9}

    result := plusOne(input)
    want := []int{1, 0, 0, 0}

    assert(t, result, want)
}

func assert(t *testing.T, result []int, want []int) {
    t.Helper()

    if !reflect.DeepEqual(result, want) {
        t.Fatalf("result: %v; want: %v", result, want)
    }
}
