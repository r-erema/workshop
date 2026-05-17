package plusone

// PlusOne increments a number represented as an array of digits by one.
// Time O(n), since we iterate each element of input
// Space O(1), since we don't involve any additional data structure.
func PlusOne(digits []int) []int {
	digits[len(digits)-1]++
	for i := len(digits) - 1; digits[i] == 10; i-- {
		digits[i] = 0
		if i-1 < 0 {
			return append([]int{1}, digits...)
		}

		digits[i-1]++
	}

	return digits
}
