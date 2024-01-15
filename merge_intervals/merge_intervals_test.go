package merge_intervals

import (
    "reflect"
    "testing"
)

func TestMergeInterval1(t *testing.T) {
    input := [][]int{{0, 0}}

    result := merge(input)
    want := [][]int{{0, 0}}

    assert(t, result, want)
}

func TestMergeInterval2(t *testing.T) {
    input := [][]int{{0, 1}}

    result := merge(input)
    want := [][]int{{0, 1}}

    assert(t, result, want)
}

func TestMergeInterval3(t *testing.T) {
    input := [][]int{{1, 1}}

    result := merge(input)
    want := [][]int{{1, 1}}

    assert(t, result, want)
}

func TestMergeInterval4(t *testing.T) {
    input := [][]int{{1, 1}, {0, 1}}

    result := merge(input)
    want := [][]int{{0, 1}}

    assert(t, result, want)
}

func TestMergeInterval5(t *testing.T) {
    input := [][]int{{1, 1}, {1, 1}, {1, 1}}

    result := merge(input)
    want := [][]int{{1, 1}}

    assert(t, result, want)
}

func TestMergeInterval6(t *testing.T) {
    input := [][]int{{1, 4}, {4, 5}}

    result := merge(input)
    want := [][]int{{1, 5}}

    assert(t, result, want)
}

func TestMergeInterval7(t *testing.T) {
    input := [][]int{{1, 3}, {1, 7}, {1, 5}}

    result := merge(input)
    want := [][]int{{1, 7}}

    assert(t, result, want)
}

func TestMergeInterval8(t *testing.T) {
    input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

    result := merge(input)
    want := [][]int{{1, 6}, {8, 10}, {15, 18}}

    assert(t, result, want)
}

func TestMergeInterval9(t *testing.T) {
    input := [][]int{{1, 4}, {5, 6}}

    result := merge(input)
    want := [][]int{{1, 4}, {5, 6}}

    assert(t, result, want)
}

func assert(t *testing.T, result [][]int, want [][]int) {
    t.Helper()

    if !reflect.DeepEqual(result, want) {
        t.Fatalf("result: %v; want: %v", result, want)
    }
}
