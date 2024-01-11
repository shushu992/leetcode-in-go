package length_of_last_word

import "testing"

func TestLengthOfLastWord1(t *testing.T) {
    input := "h"

    result := lengthOfLastWord(input)
    want := 1

    assert(t, input, want, result)
}

func TestLengthOfLastWord2(t *testing.T) {
    input := "he"

    result := lengthOfLastWord(input)
    want := 2

    assert(t, input, want, result)
}

func TestLengthOfLastWord3(t *testing.T) {
    input := "Hello World"

    result := lengthOfLastWord(input)
    want := 5

    assert(t, input, want, result)
}

func TestLengthOfLastWord4(t *testing.T) {
    input := "   fly me   to   the moon  "

    result := lengthOfLastWord(input)
    want := 4

    assert(t, input, want, result)
}

func TestLengthOfLastWord5(t *testing.T) {
    //goland:noinspection SpellCheckingInspection
    input :=  "luffy is still joyboy"

    result := lengthOfLastWord(input)
    want := 6

    assert(t, input, want, result)
}

func assert(t *testing.T, input string, want int, result int) {
    t.Helper()

    if want != result {
        t.Fatalf("input: %s; want %d; result %d;", input, want, result)
    }
}
