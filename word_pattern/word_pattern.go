package word_pattern

/**
 * https://leetcode.com/problems/word-pattern/
 *
 * 1 <= pattern.length <= 300
 * pattern contains only lower-case English letters.
 * 1 <= s.length <= 3000
 * s contains only lowercase English letters and spaces ' '.
 * s does not contain any leading or trailing spaces.
 * All the words in s are separated by a single space.
 */
func wordPattern(pattern string, s string) bool {
    arr := [26]int{}

    for i := 0; i < 26; i++ {
        arr[i] = -1
    }

    // todo

    return true
}
