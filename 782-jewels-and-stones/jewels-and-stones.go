func numJewelsInStones(jewels string, stones string) int {
    jewelsMap := make(map[rune]struct{})
    result := 0

    for _, j := range jewels {
        jewelsMap[j] = struct{}{}
    }

    for _, s := range stones {
        if _, isJewel := jewelsMap[s]; isJewel {
            result++
        }
    }

    return result
}