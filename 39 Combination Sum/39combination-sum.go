func backtrack(acc []int, candidates []int, index, currentSum, target int, res *[][]int) {
    if currentSum == target {
        copied := make([]int, len(acc))
        copy(copied, acc)
        *res = append(*res, copied)
        return
    }

    for i := index; i < len(candidates); i++ {
        if currentSum + candidates[i] > target {
            continue
        }

        acc = append(acc, candidates[i])
        backtrack(acc, candidates, i, currentSum + candidates[i], target, res)
        acc = acc[:len(acc)-1]
    }
}

func combinationSum(candidates []int, target int) [][]int {
    var result [][]int
    backtrack([]int{}, candidates, 0, 0, target, &result)
    return result
}