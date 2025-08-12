func backtrack(acc, nums []int, res *[][]int, k int) {
    if len(acc) == k {
        return
    }

    for i, num := range nums {
        acc = append(acc, num)
        copied := make([]int, len(acc))
        copy(copied, acc)
        *res = append(*res, copied)
        backtrack(acc, nums[i+1:], res, k)
        acc = acc[:len(acc)-1]
    }
}

func subsets(nums []int) [][]int {
    result := [][]int{{}}
    backtrack([]int{}, nums, &result, len(nums))
    return result
}