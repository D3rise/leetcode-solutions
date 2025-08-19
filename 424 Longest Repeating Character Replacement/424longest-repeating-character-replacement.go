func maxFromMap(m map[rune]int) int {
    var maximum int
    for _, v := range m {
        maximum = max(maximum, v)
    }
    return maximum
}

func characterReplacement(s string, k int) int {
    chars := make(map[rune]int)
    result := 0
    begin := 0

    for end, c := range s {
        chars[c]++

        for (end-begin+1) - maxFromMap(chars) > k {
            chars[rune(s[begin])]--
            begin++
        }

        result = max(result, end-begin+1)
    }

    return result
}