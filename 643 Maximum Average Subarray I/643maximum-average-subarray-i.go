import "math"

func findMaxAverage(nums []int, k int) float64 {
    result := -math.MaxFloat64
    windowState := 0
    begin := 0
    for end, num := range nums {
        windowState += num
        if end - begin + 1 == k {
            result = max(result, float64(windowState))
            windowState -= nums[begin]
            begin += 1
        }
    }

    return result / float64(k)
}