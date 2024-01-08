package binary_tree_inorder_traversal

import (
    "reflect"
    "testing"
)

func TestInorderTraversal1(t *testing.T) {
    result := inorderTraversal(nil)

    assert(t, result, []int{})
}

func TestInorderTraversal2(t *testing.T) {
    node := &TreeNode{
        1,
        nil,
        nil,
    }

    result := inorderTraversal(node)
    want := []int{1}

    assert(t, result, want)
}

func TestInorderTraversal3(t *testing.T) {
    node := &TreeNode{
        1,
        nil,
        &TreeNode{
            2,
            &TreeNode{
                3,
                nil,
                nil,
            },
            nil,
        },
    }

    result := inorderTraversal(node)
    want := []int{1, 3, 2}

    assert(t, result, want)
}

func TestInorderTraversal4(t *testing.T) {
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
            5,
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

    result := inorderTraversal(node)
    want := []int{3, 2, 4, 1, 6, 5, 7}

    assert(t, result, want)
}

func assert(t *testing.T, result []int, want []int) {
    t.Helper()

    if !reflect.DeepEqual(result, want) {
        t.Fatalf("want: %v; result: %v", want, result)
    }
}
