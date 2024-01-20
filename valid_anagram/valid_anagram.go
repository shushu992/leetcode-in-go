package valid_anagram

/**
 * https://leetcode.com/problems/valid-anagram/
 *
 * 1 <= s.length, t.length <= 5 * 10^4
 * s and t consist of lowercase English letters.
 */
func isAnagram(s string, t string) bool {
    count := [26]int32{}

    byte1 := []byte(s)

    for _, b := range byte1 {
        count[b-'a']++
    }

    byte2 := []byte(t)

    for _, b := range byte2 {
        count[b-'a']--
    }

    for _, c := range count {
        if c != 0 {
            return false
        }
    }

    return true
}
