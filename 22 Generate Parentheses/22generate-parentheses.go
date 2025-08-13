func backtrack(result *[]string, n, op, cl int, acc []string) {
    if len(acc) == n*2 {
        *result = append(*result, strings.Join(acc, ""))
        return
    }

    if op < n {
        acc = append(acc, "(")
        backtrack(result, n, op+1, cl, acc)
        acc = acc[:len(acc)-1]
    }

    if cl < op {
        acc = append(acc, ")")
        backtrack(result, n, op, cl+1, acc)
        acc = acc[:len(acc)-1]
    }
}

func generateParenthesis(n int) []string {
    var result []string
    backtrack(&result, n, 0, 0, []string{})
    return result
}