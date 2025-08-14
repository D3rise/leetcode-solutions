func backtrack(acc [][]rune, n, row int, cols, diagonals, tDiagonals map[int]struct{}, res *[][]string) {
	if row == n {
		result := make([]string, len(acc))

		for i, row := range acc {
			result[i] = string(row)
		}

		*res = append(*res, result)
		return
	}

	for col := 0; col < n; col++ {
		diag := row + col
		tDiag := row - col

		_, inCols := cols[col]
		_, inDiags := diagonals[diag]
		_, inTdiags := tDiagonals[tDiag]

		if inCols || inDiags || inTdiags {
			continue
		}

		acc[row][col] = 'Q'
		cols[col] = struct{}{}
		diagonals[diag] = struct{}{}
		tDiagonals[tDiag] = struct{}{}

		backtrack(acc, n, row+1, cols, diagonals, tDiagonals, res)

		delete(cols, col)
		delete(diagonals, diag)
		delete(tDiagonals, tDiag)
		acc[row][col] = '.'
	}
}

func solveNQueens(n int) [][]string {
	var result [][]string
	acc := make([][]rune, n)
	for i, _ := range acc {
		acc[i] = slices.Repeat([]rune{'.'}, n)
	}

	cols, diags, tDiags := make(map[int]struct{}), make(map[int]struct{}), make(map[int]struct{})
	backtrack(acc, n, 0, cols, diags, tDiags, &result)
	return result
}