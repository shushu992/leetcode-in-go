package power_of_two

/**
 * https://leetcode.com/problems/power-of-two/
 *
 * -2^31 <= n <= 2^31 - 1
 */
func isPowerOfTwo(n int) bool {
    for i := 0; i < 32; i++ {
        if n == 1<<i {
            return true
        }
    }

    return false
}
