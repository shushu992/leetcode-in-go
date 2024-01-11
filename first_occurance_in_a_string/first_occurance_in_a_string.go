package first_occurance_in_a_string

/**
 * https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
 *
 * Constraints:
 * 1 <= haystack.length, needle.length <= 10^4
 * haystack and needle consist of only lowercase English characters.
 *
 */
func strStr(haystack string, needle string) int {
loop:
    for n1 := 0; n1 <= len(haystack)-len(needle); n1++ {
        for n2 := 0; n2 < len(needle); n2++ {
            if haystack[n1+n2] != needle[n2] {
                continue loop
            }
        }
        return n1
    }

    return -1
}
