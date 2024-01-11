package remove_element

/**
 * https://leetcode.com/problems/remove-element/
 *
 * 0 <= nums.length <= 100
 * 0 <= nums[i] <= 50
 * 0 <= val <= 100
 *
 */
func removeElement(nums []int, val int) int {
    k := 0

    for i := 0; i < len(nums); i++ {
        if nums[i] == val {
            continue
        }

        nums[k] = nums[i]
        k++
    }

    return k
}
