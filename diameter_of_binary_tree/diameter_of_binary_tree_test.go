package diameter_of_binary_tree

import "testing"

func TestDiameterOfBinaryTree1(t *testing.T) {
	node := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		nil,
	}

	result := diameterOfBinaryTree(node)

	assert(t, result, 1)
}

func TestDiameterOfBinaryTree2(t *testing.T) {
	node := &TreeNode{
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
			nil,
			nil,
		},
	}

	result := diameterOfBinaryTree(node)

	assert(t, result, 3)
}

func TestDiameterOfBinaryTree3(t *testing.T) {
	node := &TreeNode{
		4,
		&TreeNode{
			1,
			&TreeNode{
				2,
				&TreeNode{
					3,
					nil,
					nil,
				},
				nil,
			},
			nil,
		},
		nil,
	}

	result := diameterOfBinaryTree(node)

	assert(t, result, 3)
}

func TestDiameterOfBinaryTree4(t *testing.T) {
	node := &TreeNode{
		5,
		&TreeNode{
			2,
			&TreeNode{
				4,
				&TreeNode{
					1,
					nil,
					nil,
				},
				nil,
			},
			&TreeNode{
				3,
				nil,
				nil,
			},
		},
		nil,
	}

	result := diameterOfBinaryTree(node)

	assert(t, result, 3)
}

func assert(t *testing.T, result int, want int) {
	t.Helper()

	if result != want {
		t.Fatalf("want: %d; got: %d", want, result)
	}
}
