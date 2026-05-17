package dailytemperatures

// DailyTemperatures returns for each day the number of days until a warmer temperature.
// Time: O(n), since we handle each element only within a stack which can't be greater than the input
// Memory: O(n), since we involve a stack that can't be greater than the input.
func DailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([][2]int, 0)

	var stackIndex int

	for i, temperature := range temperatures {
		for len(stack) > 0 && temperature > stack[len(stack)-1][1] {
			stackIndex, stack = stack[len(stack)-1][0], stack[:len(stack)-1]
			res[stackIndex] = i - stackIndex
		}

		stack = append(stack, [2]int{i, temperature})
	}

	return res
}
