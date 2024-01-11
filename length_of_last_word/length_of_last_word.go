package length_of_last_word

/**
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * 1 <= s.length <= 10^4
 * s consists of only English letters and spaces ' '.
 * There will be at least one word in s.
 */
func lengthOfLastWord(s string) int {
    count := 0
    space := false

    for _, c := range s {
        if c == ' ' {
            space = true
            continue
        }

        if space {
            count = 1
            space = false
            continue
        }

        count++
    }

    return count
}
