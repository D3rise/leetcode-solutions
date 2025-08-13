func backtrack(acc []rune, digits string, k int, letters map[rune]string, res *[]string) {
    if len(acc) == len(digits) {
        *res = append(*res, string(acc))
        return
    }

    for _, symbol := range letters[rune(digits[k])] {
        acc = append(acc, symbol)
        backtrack(acc, digits, k+1,letters, res)
        acc = acc[:len(acc)-1]
    }
}

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    letters := map[rune]string{
        '0': " ",
        '2': "abc", '3': "def", '4': "ghi", '5': "jkl",
        '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
    }

    var res []string

    backtrack([]rune{}, digits, 0, letters, &res)

    return res
}