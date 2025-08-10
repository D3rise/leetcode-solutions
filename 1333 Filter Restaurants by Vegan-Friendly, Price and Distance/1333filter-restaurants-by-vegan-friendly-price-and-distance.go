func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	result := []int{}
    idsMap := map[int]int{}

	for i, restaurant := range restaurants {
		if veganFriendly != 0 && restaurant[2] != veganFriendly {
			continue
		}

		if maxPrice >= restaurant[3] && maxDistance >= restaurant[4] {
			result = append(result, restaurant[0])
            idsMap[restaurant[0]] = i
		}
	}

    sort.Slice(result, func (i, j int) bool {
        restaurantI := restaurants[idsMap[result[i]]]
        restaurantJ := restaurants[idsMap[result[j]]]

        if restaurantI[1] != restaurantJ[1] {
            return restaurantI[1] > restaurantJ[1]
        }

        return restaurantI[0] > restaurantJ[0]
    })

    return result
}