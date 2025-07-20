func reverseString(s []byte)  {
    p1 := 0
    p2 := len(s) - 1

    for p1 < p2 {
        s[p1], s[p2] = s[p2], s[p1]
        p1++
        p2--
    }
}