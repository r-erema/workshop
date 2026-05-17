package validpalindrome

// IsPalindrome Time O(n), since we move consequently and narrow a search area
// Time O(1), since we don't involve any additional data structure.
func IsPalindrome(str string) bool {
	left, right := 0, len(str)-1

	for left < right {
		if !IsAlphanumeric(str[left]) {
			left++

			continue
		}

		if !IsAlphanumeric(str[right]) {
			right--

			continue
		}

		if ToLower(str[left]) != ToLower(str[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func IsAlphanumeric(symbol byte) bool {
	return symbol >= 'a' && symbol <= 'z' || symbol >= 'A' && symbol <= 'Z' ||
		symbol >= '0' && symbol <= '9'
}

func ToLower(symbol byte) byte {
	const caseDifference = 32

	if symbol >= 'A' && symbol <= 'Z' {
		return symbol + caseDifference
	}

	return symbol
}
