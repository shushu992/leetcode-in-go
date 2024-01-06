package rotate_array

/**
 * https://leetcode.com/problems/rotate-array/
 *
 * 1 <= nums.length <= 10^5
 * -2^31 <= nums[i] <= 2^31 - 1
 * 0 <= k <= 10^5
 */
func rotate(nums []int, k int) {
	size := len(nums)

	k = k % size
	if k == 0 {
		return
	}
	// 1 <= k < len(nums)

	for i := 0; i < (size-k)/2; i++ {
		tmp := nums[i]
		nums[i] = nums[size-k-1-i]
		nums[size-k-1-i] = tmp
	}

	for i := size - k; i < size-k/2; i++ {
		tmp := nums[i]
		nums[i] = nums[2*size-k-i-1]
		nums[2*size-k-i-1] = tmp
	}

	for i := 0; i < size/2; i++ {
		tmp := nums[i]
		nums[i] = nums[size-1-i]
		nums[size-1-i] = tmp
	}
}
