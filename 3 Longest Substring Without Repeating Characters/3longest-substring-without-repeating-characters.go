func lengthOfLongestSubstring(s string) int {
    result := 0
    begin := 0
    charsSet := make(map[rune]int)

    for end, c := range s {
        charsSet[c]++

        for charsSet[c] > 1 {
            charsSet[rune(s[begin])]--
            begin++
        }

        result = max(result, end-begin+1)
    }

    return result
}