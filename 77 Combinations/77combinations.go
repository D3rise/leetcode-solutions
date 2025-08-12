func backtrack(acc []int, res *[][]int, k, n, m int) {
    if len(acc) == k {
        copied := make([]int, len(acc))
        copy(copied, acc)
        *res = append(*res, copied)
        return
    }

    for i := m; i <= n; i++ {
        acc = append(acc, i)
        backtrack(acc, res, k, n, i+1)
        acc = acc[:len(acc)-1]
    }
}

func combine(n int, k int) [][]int {
    var res [][]int
    backtrack([]int{}, &res, k, n, 1)
    return res
}