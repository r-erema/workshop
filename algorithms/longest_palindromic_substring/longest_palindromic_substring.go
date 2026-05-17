package longest_palindromic_substring

func LongestPalindrome(input string) string {
	var res string

	helper := func(left, right int, currRes string) {
		for left >= 0 && right < len(input) && input[left] == input[right] {
			currRes = input[left : right+1]
			left--
			right++
		}

		if len(res) < len(currRes) {
			res = currRes
		}
	}

	for i := range input {
		helper(i, i, "")
		helper(i, i+1, "")
	}

	return res
}
