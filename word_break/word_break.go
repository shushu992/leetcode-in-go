package word_break

/**
 * https://leetcode.com/problems/word-break/
 * todo
 *
 * Constraints:
 *
 * 1 <= s.length <= 300
 * 1 <= wordDict.length <= 1000
 * 1 <= wordDict[i].length <= 20
 * s and wordDict[i] consist of only lowercase English letters.
 * All the strings of wordDict are unique.
 */
func wordBreak(s string, wordDict []string) bool {

    buckets := [256][1024]string{}
    sizes := [256]int{}

    for _, word := range wordDict {
        hash := byte(0)

        for _, b := range []byte(word) {
            hash += b
        }

        buckets[hash][sizes[hash]] = word
        sizes[hash]++
    }

    bytes := []byte(s)
    hash := byte(0)

    x := 0
    y := 0

loop:
    for {
        if y == len(bytes) {
            return x == y
        }

        hash += bytes[y]

        bucket := buckets[hash]
        size := sizes[hash]

        segment := string(bytes[x : y+1])

        for i := 0; i < size; i++ {
            if bucket[i] == segment {
                y++
                x = y
                hash = 0
                continue loop
            }
        }

        y++
    }
}
