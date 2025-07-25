func moveZeroes(nums []int)  {
    insertIndex := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[insertIndex], nums[i] = nums[i], nums[insertIndex]
            insertIndex++
        }
    }
}