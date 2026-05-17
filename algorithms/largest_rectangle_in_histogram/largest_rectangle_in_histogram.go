package largestrectangleinhistogram

// LargestRectangleArea finds the largest rectangle area in a histogram.
// Time O(n), since we iterate the input once
// Space O(n), since we use only stack structure which couldn't be greater then input.
func LargestRectangleArea(heights []int) int {
	var (
		stack   = make([][2]int, 0)
		row     [2]int
		maxArea int
	)

	for i, height := range heights {
		start := i

		for len(stack) > 0 && height < stack[len(stack)-1][1] {
			row, stack = stack[len(stack)-1], stack[:len(stack)-1]
			maxArea = max(maxArea, row[1]*(i-row[0]))
			start = row[0]
		}

		stack = append(stack, [2]int{start, height})
	}

	for _, row = range stack {
		maxArea = max(maxArea, row[1]*(len(heights)-row[0]))
	}

	return maxArea
}
