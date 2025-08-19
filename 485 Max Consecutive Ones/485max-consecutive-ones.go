func findMaxConsecutiveOnes(nums []int) int {
    k := 0
    begin := 0
    result := 0

    for end, num := range nums {
        if num == 0 {
            k++
        }

        for k > 0 {
            if nums[begin] == 0 {
                k--
            }
            begin++
        }

        result = max(result, end - begin + 1)
    }

    return result
}