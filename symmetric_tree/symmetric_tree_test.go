package symmetric_tree

import "testing"

func TestSymmetricTree1(t *testing.T) {
	node := &TreeNode{
		1,
		nil,
		nil,
	}

	result := isSymmetric(node)

	assertResult(t, result, true)
}

func TestSymmetricTree2(t *testing.T) {
	node := &TreeNode{
		1,
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

	result := isSymmetric(node)

	assertResult(t, result, true)
}

func TestSymmetricTree3(t *testing.T) {
	node := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				3,
				nil,
				nil,
			},
			&TreeNode{
				4,
				nil,
				nil,
			},
		},
		&TreeNode{
			2,
			&TreeNode{
				4,
				nil,
				nil,
			},
			&TreeNode{
				3,
				nil,
				nil,
			},
		},
	}

	result := isSymmetric(node)

	assertResult(t, result, true)
}

func TestSymmetricTree4(t *testing.T) {
	node := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			&TreeNode{
				3,
				nil,
				nil,
			},
		},
		&TreeNode{
			2,
			nil,
			&TreeNode{
				3,
				nil,
				nil,
			},
		},
	}

	result := isSymmetric(node)

	assertResult(t, result, false)
}

func TestSymmetricTree5(t *testing.T) {
	node := &TreeNode{
		3,
		&TreeNode{
			81,
			nil,
			&TreeNode{
				67,
				nil,
				&TreeNode{
					-87,
					nil,
					nil,
				},
			},
		},
		&TreeNode{
			81,
			&TreeNode{
				67,
				nil,
				&TreeNode{
					-85,
					nil,
					nil,
				},
			},
			nil,
		},
	}

	result := isSymmetric(node)

	assertResult(t, result, false)
}

func assertResult(t *testing.T, result bool, want bool) {
	t.Helper()

	if result != want {
		t.Fatalf("want: %v; got: %v", want, result)
	}
}
