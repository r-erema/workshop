package topkfrequentwords

import (
	"sort"
)

// TopKFrequent Time O(N*logN)
// O(NlogN) where N is the length of words.
// We count the frequency of each word in O(N) time,
// then we sort the given words in O(NlogN) time.
//
// Memory O(n)
// We have a hash table that contains a number of rows equal to the input.
func TopKFrequent(words []string, k int) []string {
	wordsCount := make(map[string]int)

	var keys []string

	for _, word := range words {
		if _, ok := wordsCount[word]; !ok {
			keys = append(keys, word)
		}

		wordsCount[word]++
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if wordsCount[keys[i]] == wordsCount[keys[j]] {
			return keys[i] < keys[j]
		}

		return wordsCount[keys[i]] > wordsCount[keys[j]]
	})

	return keys[:k]
}
