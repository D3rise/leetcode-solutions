func backtrack(acc []int, res *[][]int, counter map[int]int, maxLen int) {
    if len(acc) == maxLen {
        copied := make([]int, len(acc))
        copy(copied, acc)
        *res = append(*res, copied)
        return
    }

    for num, count := range counter {
        if count == 0 {
            continue
        }

        counter[num]--
        acc = append(acc, num)
        backtrack(acc, res, counter, maxLen)
        acc = acc[:len(acc)-1]
        counter[num]++
    }
}

func permuteUnique(nums []int) [][]int {
    var result [][]int
    counter := make(map[int]int)
    for _, n := range nums {
        counter[n]++
    }

    backtrack([]int{}, &result, counter, len(nums))
    return result
}