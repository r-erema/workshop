package numberof1bits

// HammingWeight counts the number of 1 bits in a number.
// Time O(N), since we iterate input one time
// Space O(1), since we don't involve any additional data structure.
func HammingWeight(number int) int {
	const divisor = 2

	var res int

	for number > 0 {
		res += number % divisor
		number >>= 1
	}

	return res
}

func HammingWeight2(number int) int {
	var res int

	for number > 0 {
		number &= number - 1
		res++
	}

	return res
}

func HammingWeight3(number int) int {
	var res int

	for number > 0 {
		if number&1 == 1 {
			res++
		}

		number >>= 1
	}

	return res
}
