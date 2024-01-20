package number_of_1_bits

/**
 * https://leetcode.com/problems/number-of-1-bits/
 *
 * The input must be a binary string of length 32.
 */
func hammingWeight(num uint32) int {
    var count uint32 = 0

    for i := 0; i < 32; i++ {
        count += num & 1
        num >>= 1
    }

    return int(count)
}
