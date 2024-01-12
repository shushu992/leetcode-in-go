package binary_tree_preorder_traversal

import (
    "reflect"
    "testing"
)

func TestPreorderTraversal1(t *testing.T) {
    result := preorderTraversal(nil)
    want := []int{}

    assert(t, result, want)
}

func TestPreorderTraversal2(t *testing.T) {
    node := &TreeNode{
        1,
        nil,
        nil,
    }

    result := preorderTraversal(node)
    want := []int{1}

    assert(t, result, want)
}

func TestPreorderTraversal3(t *testing.T) {
    node := &TreeNode{
        1,
        nil,
        &TreeNode{
            2,
            &TreeNode{
                3,
                nil,
                &TreeNode{
                    4,
                    nil,
                    nil,
                },
            },
            nil,
        },
    }

    result := preorderTraversal(node)
    want := []int{1, 2, 3, 4}

    assert(t, result, want)
}

func TestPreorderTraversal4(t *testing.T) {
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
                &TreeNode{
                    8,
                    nil,
                    nil,
                },
                nil,
            },
        },
    }

    result := preorderTraversal(node)
    want := []int{1, 2, 3, 4, 5, 6, 7, 8}

    assert(t, result, want)
}

func assert(t *testing.T, result []int, want []int) {
    t.Helper()

    if !reflect.DeepEqual(result, want) {
        t.Fatalf("result: %v;\nwant: %v\n", result, want)
    }
}
