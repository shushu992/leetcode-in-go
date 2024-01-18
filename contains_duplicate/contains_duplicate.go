package contains_duplicate

/**
 * https://leetcode.com/problems/contains-duplicate/
 *
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 */
func containsDuplicate(nums []int) bool {
    root := &Node{}

    for _, num := range nums {
        num += 1 << 30

        node := root
        for i := 0; i < 8; i++ {
            n := num % (1 << 4)

            if node[n] == nil {
                node[n] = &Node{}
            } else if i == 7 {
                return true
            }

            node = node[n]
            num >>= 4
        }
    }

    return false
}

type Node [16]*Node
