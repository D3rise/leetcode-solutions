func isSubsequence(s string, t string) bool {
    sequencePointer := 0

    if len(s) == 0 {
        return true
    }

    for _, targetSymbol := range t {
        if targetSymbol == rune(s[sequencePointer]) {
            sequencePointer++
        }

        if sequencePointer == len(s) {
            return true
        }
    }

    return false
}