func removeDuplicates(nums []int) int {
    p1, p2 := 0, 0

    for p2 < len(nums) {
        if nums[p1] == nums[p2] {
            p2++
            continue
        }

        nums[p1+1] = nums[p2]
        p1++
        p2++
    }

    // increase p1 by 1 to get total length of resulting array, not the index of the last element
    return p1+1
}