func totalFruit(fruits []int) int {
    baskets := make(map[int]int)
    begin := 0
    result := 0

    for end, num := range fruits {
        baskets[num]++

        for len(baskets) > 2 {
            baskets[fruits[begin]]--
            if baskets[fruits[begin]] == 0 {
                delete(baskets, fruits[begin])
            }
            begin++
        }

        result = max(result, end-begin+1)
    }

    return result
}