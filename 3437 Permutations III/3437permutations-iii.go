func backtrack(acc []int, n int, usedSet map[int]struct{}, res *[][]int) {
	if len(acc) == n {
		copied := make([]int, len(acc))
		copy(copied, acc)
		*res = append(*res, copied)
		return
	}

	for i := 1; i <= n; i++ {
        if _, used := usedSet[i]; used {
            continue
        }

		if len(acc) > 0 && acc[len(acc)-1] % 2 == i % 2 {
            continue
        }

		acc = append(acc, i)
        usedSet[i] = struct{}{}
		backtrack(acc, n, usedSet, res)
        delete(usedSet, i)
		acc = acc[:len(acc)-1]
	}
}

func permute(n int) [][]int {
	var result [][]int
	backtrack([]int{}, n, make(map[int]struct{}), &result)
	return result
}