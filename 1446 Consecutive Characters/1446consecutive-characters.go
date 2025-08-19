func maxPower(s string) int {
    result := 0
    curSeq := 0

    for i, _ := range s {
        curSeq++
        if i > 0 && s[i-1] != s[i] {
            curSeq = 1
        }
        result = max(result, curSeq)
    }

    return result
}