func lengthOfLongestSubstringKDistinct(s string, k int) int {
    chars := make(map[rune]int)

    begin := 0
    result := 0

    for end, c := range s {
        chars[c]++

        for len(chars) > k {
            chars[rune(s[begin])]--
            if chars[rune(s[begin])] == 0 {
                delete(chars, rune(s[begin]))
            }
            begin++
        }

        result = max(result, end-begin+1)
    }

    return result
}