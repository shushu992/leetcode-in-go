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

        num1 := (num >> 24) % (1 << 8)
        num2 := (num >> 16) % (1 << 8)
        num3 := (num >> 8) % (1 << 8)
        num4 := (num >> 0) % (1 << 8)

        if root[num1] == nil {
            root[num1] = &Node{}
        }

        node1 := root[num1]

        if node1[num2] == nil {
            node1[num2] = &Node{}
        }

        node2 := node1[num2]

        if node2[num3] == nil {
            node2[num3] = &Node{}
        }

        node3 := node2[num3]

        if node3[num4] == nil {
            node3[num4] = &Node{}
            continue
        }

        return true
    }

    return false
}

type Node [256]*Node
