func maxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1

	for left < right {
		curMaxHeight := min(height[left], height[right])
		curMaxArea := curMaxHeight * (right - left)

		if curMaxArea > maxArea {
			maxArea = curMaxArea
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}