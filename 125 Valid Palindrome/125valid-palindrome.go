func isPalindrome(s string) bool {
    p1 := 0
    p2 := len(s)-1

    for p1 < p2 {
        if !isAsciiLetter(s[p1]) {
            p1++
            continue
        }

        if !isAsciiLetter(s[p2]) {
            p2--
            continue
        }

        if lowerLetterIfUppercase(s[p1]) != lowerLetterIfUppercase(s[p2]) {
            return false
        }

        p1++
        p2--
        continue
    }

    return true
}

func isAsciiLetter(r byte) bool {
    isNumber := r >= '0' && r <= '9'
    isUppercaseLetter := r >= 'A' && r <= 'Z'
    isLowercaseLetter := r >= 'a' && r <= 'z'
    return isNumber || isUppercaseLetter || isLowercaseLetter
}

func lowerLetterIfUppercase(r byte) byte {
    if r >= 'A' && r <= 'Z' {
        return r + 32
    } else {
        return r
    }
}