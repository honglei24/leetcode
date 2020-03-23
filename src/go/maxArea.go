func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1
	for left < right {
		maxArea = max(maxArea, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}