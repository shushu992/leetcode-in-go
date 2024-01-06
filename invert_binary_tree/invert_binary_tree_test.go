package invert_binary_tree

import "testing"

func TestInvertBinaryTree1(t *testing.T) {
	result := invertTree(nil)

	assert(t, result, nil)
}

func TestInvertBinaryTree2(t *testing.T) {
	input := &TreeNode{1, nil, nil}

	result := invertTree(input)

	want := &TreeNode{1, nil, nil}

	assert(t, result, want)
}

func TestInvertBinaryTree3(t *testing.T) {
	input := &TreeNode{
		1,
		&TreeNode{2, nil, nil},
        &TreeNode{3, nil, nil},
	}

	result := invertTree(input)

    want := &TreeNode{
        1,
        &TreeNode{3, nil, nil},
        &TreeNode{2, nil, nil},
    }

	assert(t, result, want)
}

func TestInvertBinaryTree4(t *testing.T) {
	input := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4,
				nil,
				nil,
			},
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
		&TreeNode{
			3,
			&TreeNode{
				6,
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

	want := &TreeNode{
		1,
		&TreeNode{
			3,
			&TreeNode{
				7,
				nil,
				nil,
			},
			&TreeNode{
				6,
				nil,
				nil,
			},
		},
		&TreeNode{
			2,
			&TreeNode{
				5,
				nil,
				nil,
			},
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
	}

	result := invertTree(input)

	assert(t, result, want)
}

func assert(t *testing.T, result *TreeNode, want *TreeNode) {
	t.Helper()

	if result == nil && want == nil {
		return
	}

	if result == nil && want != nil {
		t.Fatal()
	}

	if result != nil && want == nil {
		t.Fatal()
	}

	if result == want {
		return
	}

	// assert values
	if result.Val != want.Val {
		t.Fatal()
	}

	// assert child nodes
	assert(t, result.Left, want.Left)
	assert(t, result.Right, want.Right)
}
