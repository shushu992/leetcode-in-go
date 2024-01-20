package reverse_bits

/**
 * https://leetcode.com/problems/reverse-bits/
 *
 * The input must be a binary string of length 32
 */
func reverseBits(num uint32) uint32 {
    var result uint32 = 0

    for i := 0; i < 32; i++ {
        b := num & (1 << i)

        b <<= 31 - i
        b >>= i

        result |= b
    }

    return result
}
