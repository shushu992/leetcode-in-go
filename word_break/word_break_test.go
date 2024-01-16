package word_break

import "testing"

func TestWordBreak1(t *testing.T) {
    s := "c"
    wordDict := []string{"c"}

    result := wordBreak(s, wordDict)
    assert(t, result, true)
}

func TestWordBreak2(t *testing.T) {
    s := "cc"
    wordDict := []string{"c"}

    result := wordBreak(s, wordDict)
    assert(t, result, true)
}

func TestWordBreak3(t *testing.T) {
    s := "cd"
    wordDict := []string{"c"}

    result := wordBreak(s, wordDict)
    assert(t, result, false)
}

func TestWordBreak4(t *testing.T) {
    s := "c"
    wordDict := []string{"d"}

    result := wordBreak(s, wordDict)
    assert(t, result, false)
}

//goland:noinspection SpellCheckingInspection
func TestWordBreak5(t *testing.T) {
    s := "aaaaaaa"
    wordDict := []string{"aaaa", "aaa"}

    result := wordBreak(s, wordDict)
    assert(t, result, true)
}

//goland:noinspection SpellCheckingInspection
func TestWordBreak6(t *testing.T) {
    s := "leetcode"
    wordDict := []string{"leet", "code"}

    result := wordBreak(s, wordDict)
    assert(t, result, true)
}

//goland:noinspection SpellCheckingInspection
func TestWordBreak7(t *testing.T) {
    s := "applepenapple"
    wordDict := []string{"apple", "pen"}

    result := wordBreak(s, wordDict)
    assert(t, result, true)
}

//goland:noinspection SpellCheckingInspection
func TestWordBreak8(t *testing.T) {
    s := "catsandog"
    wordDict := []string{"cats", "dog", "sand", "and", "cat"}

    result := wordBreak(s, wordDict)
    assert(t, result, false)
}

func assert(t *testing.T, result bool, want bool) {
    t.Helper()

    if result != want {
        t.Fatalf("want: %v; got: %v", want, result)
    }
}
