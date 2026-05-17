package productsum

// ProductSum calculates the product sum of a nested array.
// Time O(n), since we walk trough the n 1 time
// Space O(1), we don't use any extra space.
func ProductSum(array []any) int {
	return Helper(array, 1)
}

func Helper(array []any, depth int) int {
	var sum int

	for _, element := range array {
		switch v := element.(type) {
		case []any:
			sum += Helper(v, depth+1)
		case int:
			sum += v
		}
	}

	return sum * depth
}
