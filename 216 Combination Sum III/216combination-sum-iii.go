func backtrack(acc []int, currentSum, currentNum, k, n int, result *[][]int) {
    if len(acc) == k {
        if currentSum == n {
            copied := make([]int, len(acc))
            copy(copied, acc)
            *result = append(*result, copied)
        }
        return
    }

    if currentSum > n {
        return
    }

    for num := currentNum; num <= 9; num++ {
        acc = append(acc, num)
        backtrack(acc, currentSum + num, num+1, k, n, result)
        acc = acc[:len(acc)-1]
    }
}

func combinationSum3(k int, n int) [][]int {
    var result [][]int
    backtrack([]int{}, 0, 1, k, n, &result)
    return result
}