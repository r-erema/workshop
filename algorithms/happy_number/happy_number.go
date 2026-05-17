package happynumber

// IsHappy checks if a number is a happy number.
// Time O(log(n)) where n is the input number, as the sequence of transformations will be at most O(log n)
// before repeating or reaching 1
// Space O(n), since we use map.
func IsHappy(n int) bool {
	nums := make(map[int]struct{})

	const decimalBase = 10

	for _, ok := nums[n]; !ok; _, ok = nums[n] {
		nums[n] = struct{}{}

		sum := 0

		for n != 0 {
			digit := n % decimalBase
			sum += digit * digit
			n /= decimalBase
		}

		if sum == 1 {
			return true
		}

		n = sum
	}

	return false
}
