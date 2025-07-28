func backspaceCompare(s string, t string) bool {
    sStack := processStringIntoStack(s)
    tStack := processStringIntoStack(t)

    if len(sStack) != len(tStack) {
        return false
    }

    for i, v := range sStack {
        if tStack[i] != v {
            return false
        }
    }

    return true
}

func processStringIntoStack(s string) []rune {
    stack := []rune{}

    for _, v := range s {
        if v == '#' {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
            continue
        }
        stack = append(stack, v)
    }

    return stack
}