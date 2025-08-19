func numOfUnplacedFruits(fruits []int, baskets []int) int {
    unplaced := len(fruits)
    occupiedBaskets := make(map[int]struct{})

    for _, fruit := range fruits {
        for i, basket := range baskets {
            _, occupied := occupiedBaskets[i]

            if !occupied && basket >= fruit {
                unplaced--
                occupiedBaskets[i] = struct{}{}
                break
            }
        }
    }

    return unplaced
}