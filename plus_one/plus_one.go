package plus_one

/**
 * https://leetcode.com/problems/plus-one/
 *
 * Constraints:
 * 1 <= digits.length <= 100
 * 0 <= digits[i] <= 9
 * digits does not contain any leading 0's.
 */
func plusOne(digits []int) []int {
    carry := true

    for i := len(digits) - 1; i >= 0; i-- {
        if !carry {
            continue
        }

        if digits[i] == 9 {
            digits[i] = 0
            carry = true
        } else {
            digits[i]++
            carry = false
        }
    }

    if !carry {
        return digits
    }

    arr := make([]int, len(digits)+1)
    arr[0] = 1
    copy(arr[1:], digits)

    return arr
}
