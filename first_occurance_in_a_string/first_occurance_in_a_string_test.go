package first_occurance_in_a_string

import (
    "testing"
)

func TestStrStr1(t *testing.T) {
    result := strStr("a", "b")
    assert(t, result, -1)
}

func TestStrStr2(t *testing.T) {
    result := strStr("a", "a")
    assert(t, result, 0)
}

func TestStrStr3(t *testing.T) {
    result := strStr("a", "ab")
    assert(t, result, -1)
}

func TestStrStr4(t *testing.T) {
    result := strStr("ab", "ab")
    assert(t, result, 0)
}

func TestStrStr5(t *testing.T) {
    //goland:noinspection SpellCheckingInspection
    result := strStr("abcd", "bc")
    assert(t, result, 1)
}

func TestStrStr6(t *testing.T) {
    //goland:noinspection SpellCheckingInspection
    result := strStr("abcd", "cd")
    assert(t, result, 2)
}

func assert(t *testing.T, result int, want int) {
    t.Helper()

    if result != want {
        t.Fatalf("result: %d; want: %d", result, want)
    }
}
