func sortedSquares(nums []int) []int {
    left := 0
    right := len(nums)-1

    result := make([]int, len(nums))
    nextNum := right

    for left <= right {
        sqLeft := nums[left] * nums[left]
        sqRight := nums[right] * nums[right]

        if sqRight >= sqLeft {
            result[nextNum] = sqRight
            right--
        } else {
            result[nextNum] = sqLeft
            left++
        }

        nextNum--
    }

    return result
}