package same_tree

import (
    "testing"
)

func TestIsSameTree01(t *testing.T) {
    result := isSameTree(nil, nil)
    assert(t, result, true)
}

func TestIsSameTree02(t *testing.T) {
    p := &TreeNode{
        1,
        nil,
        nil,
    }
    result := isSameTree(p, nil)
    assert(t, result, false)
}

func TestIsSameTree03(t *testing.T) {
    q := &TreeNode{
        1,
        nil,
        nil,
    }
    result := isSameTree(nil, q)
    assert(t, result, false)
}

func TestIsSameTree04(t *testing.T) {
    p := &TreeNode{
        1,
        nil,
        nil,
    }
    q := &TreeNode{
        1,
        nil,
        nil,
    }
    result := isSameTree(p, q)
    assert(t, result, true)
}

func TestIsSameTree05(t *testing.T) {
    p := &TreeNode{
        1,
        nil,
        nil,
    }
    q := &TreeNode{
        2,
        nil,
        nil,
    }
    result := isSameTree(p, q)
    assert(t, result, false)
}

func TestIsSameTree06(t *testing.T) {
    p := &TreeNode{
        0,
        &TreeNode{
            -5,
            nil,
            nil,
        },
        nil,
    }
    q := &TreeNode{
        0,
        &TreeNode{
            -8,
            nil,
            nil,
        },
        nil,
    }
    result := isSameTree(p, q)
    assert(t, result, false)
}

func TestIsSameTree07(t *testing.T) {
    p := &TreeNode{
        1,
        &TreeNode{
            2,
            nil,
            nil,
        },
        &TreeNode{
            3,
            nil,
            nil,
        },
    }
    q := &TreeNode{
        1,
        &TreeNode{
            2,
            nil,
            nil,
        },
        &TreeNode{
            3,
            nil,
            nil,
        },
    }
    result := isSameTree(p, q)
    assert(t, result, true)
}

func TestIsSameTree08(t *testing.T) {
    p := &TreeNode{
        1,
        &TreeNode{
            2,
            nil,
            nil,
        },
        &TreeNode{
            3,
            nil,
            nil,
        },
    }
    q := &TreeNode{
        1,
        &TreeNode{
            2,
            nil,
            nil,
        },
        &TreeNode{
            4,
            nil,
            nil,
        },
    }
    result := isSameTree(p, q)
    assert(t, result, false)
}

func TestIsSameTree09(t *testing.T) {
    p := &TreeNode{
        1,
        &TreeNode{
            2,
            nil,
            nil,
        },
        &TreeNode{
            3,
            nil,
            nil,
        },
    }
    q := &TreeNode{
        1,
        &TreeNode{
            2,
            nil,
            nil,
        },
        nil,
    }
    result := isSameTree(p, q)
    assert(t, result, false)
}

func TestIsSameTree10(t *testing.T) {
    p := &TreeNode{
        1,
        &TreeNode{
            2,
            nil,
            nil,
        },
        &TreeNode{
            3,
            nil,
            nil,
        },
    }
    q := &TreeNode{
        1,
        &TreeNode{
            3,
            nil,
            nil,
        },
        &TreeNode{
            2,
            nil,
            nil,
        },
    }
    result := isSameTree(p, q)
    assert(t, result, false)
}

func TestIsSameTree11(t *testing.T) {
    p := &TreeNode{
        1,
        &TreeNode{
            2,
            nil,
            nil,
        },
        &TreeNode{
            3,
            nil,
            nil,
        },
    }
    q := &TreeNode{
        1,
        &TreeNode{
            2,
            nil,
            &TreeNode{
                3,
                nil,
                &TreeNode{
                    4,
                    nil,
                    nil,
                },
            },
        },
        nil,
    }
    result := isSameTree(p, q)
    assert(t, result, false)
}

func assert(t *testing.T, result bool, want bool) {
    t.Helper()

    if result != want {
        t.Fatalf("result: %v; want: %v", result, want)
    }
}
