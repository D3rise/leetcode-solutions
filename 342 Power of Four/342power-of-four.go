func isPowerOfFour(n int) bool {
    if n == 0 {
        return false
    }

    x := math.Log(float64(n)) / math.Log(4.0)
    if x == math.Trunc(x) {
        return true
    }

    return false
}