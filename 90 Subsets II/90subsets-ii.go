func backtrack(acc, nums []int, k int, res *[][]int) {
    copied := make([]int, len(acc))
    copy(copied, acc)
    *res = append(*res, copied)

    if k == len(nums) {
        return
    }

    for i := k; i < len(nums); i++ {
        if i != k && nums[i] == nums[i-1] {
            continue
        }

        acc = append(acc, nums[i])
        backtrack(acc, nums, i+1, res)
        acc = acc[:len(acc)-1]
    }
}

func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    backtrack([]int{}, nums, 0, &result)
    return result
}