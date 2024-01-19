package reverse_words_in_a_string

import "testing"

func TestReverseWords1(t *testing.T) {
    input := "a"

    result := reverseWords(input)
    want := "a"

    assert(t, result, want)
}

func TestReverseWords2(t *testing.T) {
    input := "aB"

    result := reverseWords(input)
    want := "aB"

    assert(t, result, want)
}

func TestReverseWords3(t *testing.T) {
    input := "A bC"

    result := reverseWords(input)
    want := "bC A"

    assert(t, result, want)
}

func TestReverseWords4(t *testing.T) {
    input := "  hello world  "

    result := reverseWords(input)
    want := "world hello"

    assert(t, result, want)
}

func TestReverseWords5(t *testing.T) {
    input := "a good   example"

    result := reverseWords(input)
    want := "example good a"

    assert(t, result, want)
}

func assert(t *testing.T, result string, want string) {
    t.Helper()

    if result != want {
        t.Fatalf("result: %s; want: %s", result, want)
    }
}
