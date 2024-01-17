package unique_number_of_occurences

/**
 * https://leetcode.com/problems/unique-number-of-occurrences/?envType=daily-question&envId=2024-01-17
 *
 * 1 <= arr.length <= 1000
 * -1000 <= arr[i] <= 1000
 */
func uniqueOccurrences(arr []int) bool {
    count := [2048]uint16{}

    for _, n := range arr {
        count[n+1024]++
    }

    bit := [2048]bool{}

    for _, c := range count {
        if c == 0 {
            continue
        }

        if bit[c] {
            return false
        }

        bit[c] = true
    }

    return true
}
