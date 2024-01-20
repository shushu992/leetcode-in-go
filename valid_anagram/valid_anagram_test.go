package valid_anagram

import "testing"

func TestIsAnagram1(t *testing.T) {
    s1 := "a"
    s2 := "a"

    result := isAnagram(s1, s2)
    assert(t, result, true)
}

func TestIsAnagram2(t *testing.T) {
    s1 := "a"
    s2 := "b"

    result := isAnagram(s1, s2)
    assert(t, result, false)
}

//goland:noinspection SpellCheckingInspection
func TestIsAnagram3(t *testing.T) {
    s1 := "anagram"
    s2 := "nagaram"

    result := isAnagram(s1, s2)
    assert(t, result, true)
}
func TestIsAnagram4(t *testing.T) {
    s1 := "rat"
    s2 := "car"

    result := isAnagram(s1, s2)
    assert(t, result, false)
}

func assert(t *testing.T, result bool, want bool) {
    t.Helper()

    if result != want {
        t.Fatalf("want: %v; got:  %v", want, result)
    }
}
