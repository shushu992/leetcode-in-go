package binary_tree_level_order_traversal

import (
    "reflect"
    "testing"
)

func TestLevelOrderTraversal1(t *testing.T) {
    result := levelOrder(nil)
    want := [][]int{}

    assert(t, result, want)
}

func TestLevelOrderTraversal2(t *testing.T) {
    node := &TreeNode{
        1,
        nil,
        nil,
    }

    result := levelOrder(node)
    want := [][]int{{1}}

    assert(t, result, want)
}

func TestLevelOrderTraversal3(t *testing.T) {
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

    result := levelOrder(node)
    want := [][]int{{1}, {2}, {3}}

    assert(t, result, want)
}

func TestLevelOrderTraversal4(t *testing.T) {
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

    result := levelOrder(node)
    want := [][]int{{1}, {2, 5}, {3, 4, 6, 7}}

    assert(t, result, want)
}

func assert(t *testing.T, result [][]int, want [][]int) {
    t.Helper()

    if !reflect.DeepEqual(result, want) {
        t.Fatalf("want: %v; result: %v", want, result)
    }
}
