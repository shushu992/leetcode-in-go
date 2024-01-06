package rotate_array

import (
	"reflect"
	"testing"
)

func TestRotateArray1(t *testing.T) {
	arr := []int{-1, -100, 3, 99}
	rotate(arr, 2)
	want := []int{3, 99, -1, -100}
	assert(t, arr, want)
}

func TestRotateArray2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 3)
	want := []int{5, 6, 7, 1, 2, 3, 4}
	assert(t, arr, want)
}

func TestRotateArray3(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 15)
	want := []int{7, 1, 2, 3, 4, 5, 6}
	assert(t, arr, want)
}

func assert(t *testing.T, result []int, want []int) {
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("result: %v; want: %v", result, want)
	}
}
