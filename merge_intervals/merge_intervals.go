package merge_intervals

/**
 * https://leetcode.com/problems/merge-intervals/
 *
 * Constraints:
 * 1 <= intervals.length <= 10^4
 * intervals[i].length == 2
 * 0 <= start#i <= end#i <= 10^4
 */
func merge(intervals [][]int) [][]int {
    arr := make([]int, 10240)

    for i, _ := range arr {
        arr[i] = -1
    }

    for _, interval := range intervals {
        if arr[interval[0]] < interval[1] {
            arr[interval[0]] = interval[1]
        }
    }

    result := make([][]int, 10240)
    size := 0

    l := -1
    r := -1

    for i := 0; i < 10240; {
        j := arr[i]

        if l == -1 {
            if j >= 0 {
                l = i
                r = j
            }

            i++
            continue
        }

        if i > r {
            result[size] = []int{l, r}
            size++

            l = -1
            j = -1

            continue
        }

        if r < j {
            r = j
        }
        i++
    }

    return result[0:size]
}
