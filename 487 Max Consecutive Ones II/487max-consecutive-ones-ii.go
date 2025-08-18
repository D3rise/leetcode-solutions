func findMaxConsecutiveOnes(nums []int) int {
    begin := 0
    consecutiveZeroes := 0
    result := 0

    for end, num := range nums {
        if num == 0 {
            consecutiveZeroes++
        }

        for consecutiveZeroes > 1 {
            if nums[begin] == 0 {
                consecutiveZeroes--
            }
            begin++
        }

        result = max(result, end-begin+1)
    }

    return result
}