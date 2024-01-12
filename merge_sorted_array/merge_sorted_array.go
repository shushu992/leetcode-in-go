package merge_sorted_array

/**
 * https://leetcode.com/problems/merge-sorted-array/
 *
 * Constraints:
 * nums1.length == m + n
 * nums2.length == n
 * 0 <= m, n <= 200
 * 1 <= m + n <= 200
 * -10^9 <= nums1[i], nums2[j] <= 10^9
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
    m--
    n--

    for z := m + n + 1; z >= 0; z-- {
        if m < 0 {
            nums1[z] = nums2[n]
            n--
        } else if n < 0 {
            nums1[z] = nums1[m]
            m--
        } else if nums1[m] >= nums2[n] {
            nums1[z] = nums1[m]
            m--
        } else {
            nums1[z] = nums2[n]
            n--
        }
    }
}
