func minSubArrayLen(target int, nums []int) int {
    begin := 0
    windowSum := 0
    result := math.MaxInt32
    
    for end, num := range nums {
        windowSum += num

        for windowSum >= target {
            result = min(result, end - begin + 1)
            windowSum -= nums[begin]
            begin++
        }
    }

    if result == math.MaxInt32 {
        return 0
    }

    return result
}