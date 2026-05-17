package longestrepeatingcharacterreplacement

// CharacterReplacement finds the longest substring with character replacement.
// Time O(n), since we iterate input one time
// Space O(1), since we use only additional fixed array with 26 elements.
func CharacterReplacement(str string, m int) int {
	const (
		alphabetSize = 26
		uppercaseA   = 65
	)

	var (
		res, left, right, maxFrequentChar int
		charsCount                        [alphabetSize]int
	)

	for right = range str {
		charIndexRight := str[right] - uppercaseA
		charsCount[charIndexRight]++

		maxFrequentChar = max(maxFrequentChar, charsCount[charIndexRight])

		for (right - left + 1 - maxFrequentChar) > m {
			charIndexLeft := str[left] - uppercaseA
			charsCount[charIndexLeft]--
			left++
		}

		res = max(res, right-left+1)
	}

	return res
}
