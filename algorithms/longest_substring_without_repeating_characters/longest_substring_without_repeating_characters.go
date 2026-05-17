package longestsubstringwithoutrepeatingcharacters

// LengthOfLongestSubstring finds the length of the longest substring without repeating characters.
// Time O(n), since we should iterate all the input
// Space O(n), we may allocate memory in map equal to input chars count.
func LengthOfLongestSubstring(str string) int {
	var left, res int

	chars := make(map[byte]int)

	for right := range str {
		if lastPosition, ok := chars[str[right]]; ok && lastPosition >= left {
			left = lastPosition + 1
		}

		chars[str[right]] = right
		res = max(res, right-left+1)
	}

	return res
}
