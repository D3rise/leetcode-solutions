func longestValidParentheses(s string) int {
    stack := []int{-1}
    result := 0

    for i, c := range s {
        if c == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                stack = append(stack, i)
            } else {
                result = max(result, i - stack[len(stack)-1])
            }
        }
    }

    return result
}