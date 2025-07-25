import "math"

func maxSum(nums []int) int {
    max := math.MinInt32
    seenNumbers := make(map[int]struct{})
    sum := 0

    for _, val := range nums {
        if val > max {
            max = val
        }

        if _, seen := seenNumbers[val]; val > 0 && !seen {
            sum += val
            seenNumbers[val] = struct{}{}
        }
    }

    if sum == 0 {
        return max
    }

    return sum
}