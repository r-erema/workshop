package countingbits

// CountBits counts the number of 1 bits for each number up to n.
// Time O(N), since we iterate input one time
// Space O(1), since we involve array with lengths equals input number.
func CountBits(number int) []int {
	dynProg := make([]int, number+1)

	for offset, i := 1, 1; i <= number; i++ {
		if offset*2 == i {
			offset = i
		}

		dynProg[i] = 1 + dynProg[i-offset]
	}

	return dynProg
}
