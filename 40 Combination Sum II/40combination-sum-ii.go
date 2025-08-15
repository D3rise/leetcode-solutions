func backtrack(acc, candidates []int, used map[int]struct{}, target, currentSum, n int, res *[][]int) {
    if currentSum == target {
        copied := make([]int, len(acc))
        copy(copied, acc)
        *res = append(*res, copied)
        return
    }

    for i := n; i < len(candidates); i++ {
        if i != n && candidates[i] == candidates[i-1] {
            continue
        }

        if currentSum + candidates[i] > target {
            continue
        }

        if _, alreadyUsed := used[i]; alreadyUsed {
            continue
        }

        used[i] = struct{}{}
        acc = append(acc, candidates[i])

        backtrack(acc, candidates, used, target, currentSum + candidates[i], i+1, res)

        acc = acc[:len(acc)-1]
        delete(used, i)
    }
}

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    used := make(map[int]struct{})
    var result [][]int

    backtrack([]int{}, candidates, used, target, 0, 0, &result)

    return result
}