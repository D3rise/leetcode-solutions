func longestSubarray(nums []int) int {
    k := 1
    begin := 0
    windowState := 0
    result := 0

    for end, num := range nums {
        if num == 0 {
            windowState++
        }

        for windowState > k {
            if nums[begin] == 0 {
                windowState--
            }
            begin++
        }

        result = max(result, end - begin + 1)
    }

    return result - 1
}