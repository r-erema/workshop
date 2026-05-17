package validanagram

// IsAnagram Time O(n), since we need to iterate each element in a word
// Space O(1) - we use an array with length 26 to be able to store all 26 ASCII characters,
//
//	where indexes are difference of ASCII code of lower case letter and minimum possible code of lower case letter,
//	i.e. 'a' == 97, e.g. if we have letter 'z', it's code 122, but we can't add it into 26 length array,
//	so we subtract 122-97=25 and can put it into array, i.e. array[25]++
func IsAnagram(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	arr := [26]int{}

	for i := range len(word1) {
		arr[word1[i]-'a']++
		arr[word2[i]-'a']--
	}

	return arr == [26]int{}
}
