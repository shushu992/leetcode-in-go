package maximum_depth_of_binary_tree

import "testing"

func TestMaximumDepthOfBinaryTree1(t *testing.T) {
	result := maxDepth(nil)
	want := 0

	assertDepth(t, result, want)
}

func TestMaximumDepthOfBinaryTree2(t *testing.T) {
	node := &TreeNode{
		1,
		nil,
		nil,
	}

	result := maxDepth(node)
	want := 1

	assertDepth(t, result, want)
}

func TestMaximumDepthOfBinaryTree3(t *testing.T) {
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

func TestMaximumDepthOfBinaryTree4(t *testing.T) {
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
