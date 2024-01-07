package search_insert_position

/**
 * https://leetcode.com/problems/search-insert-position/
 *
 * 1 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums contains distinct values sorted in ascending order.
 * -10^4 <= target <= 10^4
 */
func searchInsert(nums []int, target int) int {
	start := 0
	size := len(nums)

	for {
		if size == 1 {
			if nums[start] < target {
				return start + 1
			}
			return start
		}

		if nums[start+size/2] < target {
			start += size / 2
			size -= size / 2
		} else {
			size /= 2
		}
	}
}
