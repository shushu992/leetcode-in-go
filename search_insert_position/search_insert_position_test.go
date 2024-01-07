package search_insert_position

import (
	"testing"
)

func TestSearchInsert1(t *testing.T) {
	nums := []int{3}
	target := 2

	result := searchInsert(nums, target)
	want := 0

	assert(t, result, want)
}

func TestSearchInsert2(t *testing.T) {
	nums := []int{3}
	target := 4

	result := searchInsert(nums, target)
	want := 1

	assert(t, result, want)
}

func TestSearchInsert3(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5

	result := searchInsert(nums, target)
	want := 2

	assert(t, result, want)
}

func TestSearchInsert4(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2

	result := searchInsert(nums, target)
	want := 1

	assert(t, result, want)
}

func TestSearchInsert5(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 7

	result := searchInsert(nums, target)
	want := 4

	assert(t, result, want)
}

func assert(t *testing.T, result int, want int) {
	t.Helper()

	if result != want {
		t.Fatalf("result: %d; want: %d", result, want)
	}
}
