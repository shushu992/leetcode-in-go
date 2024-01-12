package add_binary

/**
 * https://leetcode.com/problems/add-binary/
 *
 * Constraints:
 * 1 <= a.length, b.length <= 10^4
 * a and b consist only of '0' or '1' characters.
 * Each string does not contain leading zeros except for the zero itself.
 */
func addBinary(a string, b string) string {
    arr := make([]byte, 10240)

    carry := 0
    i := 0

    for ; i < len(arr); i++ {
        if carry == 0 && i >= len(a) && i >= len(b) {
            break
        }

        sum := carry
        carry = 0

        if i < len(a) {
            if a[len(a)-i-1] == '1' {
                sum++
            }
        }

        if i < len(b) {
            if b[len(b)-i-1] == '1' {
                sum++
            }
        }

        if sum == 0 {
            arr[len(arr)-i-1] = '0'
        } else if sum == 1 {
            arr[len(arr)-i-1] = '1'
        } else if sum%2 == 0 {
            arr[len(arr)-i-1] = '0'
            carry = 1
        } else {
            arr[len(arr)-i-1] = '1'
            carry = 1
        }
    }

    return string(arr[len(arr)-i:])
}
