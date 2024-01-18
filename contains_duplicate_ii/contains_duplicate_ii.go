package contains_duplicate_ii

/**
 * https://leetcode.com/problems/contains-duplicate-ii/
 *
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * 0 <= k <= 10^5
 */
func containsNearbyDuplicate(nums []int, k int) bool {
    root := &Node{}

    for i, num := range nums {
        num += 1 << 30

        node := root
        for j := 0; j < 8; j++ {
            n := num % (1 << 4)

            if node.Next[n] == nil {
                node.Next[n] = &Node{Val: i}
            } else if j == 7 {
                prev := node.Next[n]
                if i-prev.Val <= k {
                    return true
                }
                node.Next[n] = &Node{Val: i}
            }

            node = node.Next[n]
            num >>= 4
        }
    }

    return false
}

type Node struct {
    Val  int
    Next [16]*Node
}
