package pow

/**
 * https://leetcode.com/problems/powx-n/
 *
 * Constraints:
 *
 * -100.0 < x < 100.0
 * -2^31 <= n <= 2^31-1
 * n is an integer.
 * Either x is not zero or n > 0.
 * -10^4 <= x^n <= 10^4
 */
func myPow(x float64, n int) float64 {
    if x == 0 {
        return 0
    }

    if n == 0 {
        return 1
    }

    if n < 0 {
        n = -n
        x = 1 / x
    }

    arr := [32]float64{x}

    for i := 0; i <= 30; i++ {
        arr[i+1] = arr[i] * arr[i]
    }

    result := x
    n--

    for i := 31; ; {
        if n == 0 {
            break
        }

        if n < 1<<i {
            i--
            continue
        }

        result *= arr[i]
        n -= 1 << i
    }

    return result
}
