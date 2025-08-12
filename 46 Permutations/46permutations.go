func backtrack(acc []int, nums []int, res *[][]int) {
    if len(acc) == len(nums) {
        copied := make([]int, len(acc))
        copy(copied, acc)
        *res = append(*res, copied)
        return
    }

    for _, n := range nums {
        var exists bool

        for _, num := range acc {
            if num == n {
                exists = true
            }
        }

        if exists {
            continue
        }

        acc = append(acc, n)
        backtrack(acc, nums, res)
        acc = acc[:len(acc)-1]
    }
}

func permute(nums []int) [][]int {
    var res [][]int
    backtrack([]int{}, nums, &res)
    return res
}