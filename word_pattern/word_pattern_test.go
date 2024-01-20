package word_pattern

import "testing"

func TestWordPattern1(t *testing.T) {
    pattern := "a"
    s := "dog"

    result := wordPattern(pattern, s)
    assert(t, result, true)
}

func TestWordPattern2(t *testing.T) {
    pattern := "a"
    s := "dog dog"

    result := wordPattern(pattern, s)
    assert(t, result, false)
}

func TestWordPattern3(t *testing.T) {
    pattern := "aa"
    s := "dog dog"

    result := wordPattern(pattern, s)
    assert(t, result, true)
}

func TestWordPattern4(t *testing.T) {
    pattern := "aa"
    s := "dog"

    result := wordPattern(pattern, s)
    assert(t, result, false)
}

func TestWordPattern5(t *testing.T) {
    pattern := "aa"
    s := "dog dog dog"

    result := wordPattern(pattern, s)
    assert(t, result, false)
}

func TestWordPattern6(t *testing.T) {
    pattern := "ab"
    s := "dog dog"

    result := wordPattern(pattern, s)
    assert(t, result, false)
}

func TestWordPattern7(t *testing.T) {
    pattern := "aa"
    s := "dog cat"

    result := wordPattern(pattern, s)
    assert(t, result, false)
}

func TestWordPattern8(t *testing.T) {
    pattern := "abba"
    s := "dog cat cat dog"

    result := wordPattern(pattern, s)
    assert(t, result, true)
}

func TestWordPattern9(t *testing.T) {
    pattern := "abba"
    s := "dog cat cat fish"

    result := wordPattern(pattern, s)
    assert(t, result, false)
}

//goland:noinspection SpellCheckingInspection
func TestWordPattern10(t *testing.T) {
    pattern := "aaaa"
    s := "dog cat cat dog"

    result := wordPattern(pattern, s)
    assert(t, result, true)
}

func assert(t *testing.T, result bool, want bool) {
    t.Helper()

    if result != want {
        t.Fatalf("want: %v; got: %v", want, result)
    }
}
