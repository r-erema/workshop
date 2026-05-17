package groupanagrams

// GroupAnagrams groups strings that are anagrams of each other.
// Time O(M * N) - we need to increment each symbol(M) in each word(N)
// Space O(N) - we need to persist each word(N) in hash map.
func GroupAnagrams(strings []string) [][]string {
	anagramsMap := make(map[[26]byte][]string)

	for _, word := range strings {
		var wordASCIICount [26]byte

		for _, letterByte := range word {
			wordASCIICount[letterByte-'a']++
		}

		anagramsMap[wordASCIICount] = append(anagramsMap[wordASCIICount], word)
	}

	result := make([][]string, 0, len(anagramsMap))
	for _, group := range anagramsMap {
		result = append(result, group)
	}

	return result
}
