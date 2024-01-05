package leetcode_in_go

import "testing"

func TestMaximunDepthOfBinaryTree1(t *testing.T) {
	result := maxDepth(nil)
	want := 0

	assertDepth(t, result, want)
}

func TestMaximunDepthOfBinaryTree2(t *testing.T) {
	node := &TreeNode{
		1,
		nil,
		nil,
	}

	result := maxDepth(node)
	want := 1

	assertDepth(t, result, want)
}

func TestMaximunDepthOfBinaryTree3(t *testing.T) {
	node := &TreeNode{
		3,
		&TreeNode{
			9,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{
				15,
				nil,
				nil,
			},
			&TreeNode{
				7,
				nil,
				nil,
			},
		},
	}

	result := maxDepth(node)
	want := 3

	assertDepth(t, result, want)
}

func TestMaximunDepthOfBinaryTree4(t *testing.T) {
	node := &TreeNode{
		1,
		nil,
		&TreeNode{
			2,
			nil,
			nil,
		},
	}

	result := maxDepth(node)
	want := 2

	assertDepth(t, result, want)
}

func assertDepth(t *testing.T, want int, result int) {
	t.Helper()

	if result != want {
		t.Fatalf("want: %d; got: %d", want, result)
	}
}
