package validate_binary_search_tree

import "testing"

func TestValidateBST1(t *testing.T) {
	node := &TreeNode{
		5,
		nil,
		nil,
	}

	result := isValidBST(node)
	assert(t, result, true)
}

func TestValidateBST2(t *testing.T) {
	node := &TreeNode{
		2,
		&TreeNode{
			1,
			nil,
			nil,
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}

	result := isValidBST(node)
	assert(t, result, true)
}

func TestValidateBST3(t *testing.T) {
	node := &TreeNode{
		2,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			2,
			nil,
			nil,
		},
	}

	result := isValidBST(node)
	assert(t, result, false)
}

func TestValidateBST4(t *testing.T) {
	node := &TreeNode{
		5,
		&TreeNode{
			1,
			nil,
			nil,
		},
		&TreeNode{
			4,
			&TreeNode{
				3,
				nil,
				nil,
			},
			&TreeNode{
				6,
				nil,
				nil,
			},
		},
	}

	result := isValidBST(node)
	assert(t, result, false)
}

func TestValidateBST5(t *testing.T) {
	node := &TreeNode{
		5,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			7,
			&TreeNode{
				4,
				nil,
				nil,
			},
			&TreeNode{
				9,
				nil,
				nil,
			},
		},
	}

	result := isValidBST(node)
	assert(t, result, false)
}

func assert(t *testing.T, result bool, want bool) {
	t.Helper()

	if result != want {
		t.Fatalf("want: %v; got: %v", want, result)
	}
}
