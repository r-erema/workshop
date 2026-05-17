package containerwithmostwater

// MaxArea calculates the maximum area that can be trapped between heights.
// Time O(n), since we traverse each height
// Space O(1), We don't use any extra space for memorizing input, etc.
func MaxArea(heights []int) int {
	var maxAreaResult int

	left, right := 0, len(heights)-1

	for left < right {
		minSide := min(heights[left], heights[right])
		maxAreaResult = max(maxAreaResult, minSide*(right-left))

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return maxAreaResult
}
