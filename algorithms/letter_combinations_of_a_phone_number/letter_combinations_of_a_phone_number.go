package lettercombinationsofaphonenumber

// LetterCombinations returns all possible letter combinations for phone number digits.
// Space O(4^n), because we store all the possible combinations of letters.
func LetterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	digitsToChar := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var (
		backtrack func(i int, curStr string)
		res       []string
	)

	backtrack = func(i int, curStr string) {
		if len(curStr) == len(digits) {
			res = append(res, curStr)

			return
		}

		for _, letter := range digitsToChar[digits[i]] {
			backtrack(i+1, curStr+letter)
		}
	}

	backtrack(0, "")

	return res
}
