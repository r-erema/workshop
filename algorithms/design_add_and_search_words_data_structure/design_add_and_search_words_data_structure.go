package designaddandsearchwordsdatastructure

type WordDictionary struct {
	endOfWord bool
	children  [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

// AddWord adds a word to the word dictionary.
// Time O(n), since we need iterate entire tree
// Space O(t+n), where n is the length of the string and t is the total number of TrieNodes created in the Trie.
func (trie *WordDictionary) AddWord(word string) {
	curr := trie

	for _, i := range word {
		idx := i - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &WordDictionary{}
		}

		curr = curr.children[idx]
	}

	curr.endOfWord = true
}

// Search searches for a word in the word dictionary.
// Time O(n), since we need iterate entire tree
// Space O(t+n), where n is the length of the string and t is the total number of TrieNodes created in the Trie.
func (trie *WordDictionary) Search(word string) bool {
	var dfs func(i int, node *WordDictionary) bool

	dfs = func(i int, node *WordDictionary) bool {
		for ; i < len(word); i++ {
			char := word[i]

			if char != '.' {
				idx := char - 'a'
				if node == nil || node.children[idx] == nil {
					return false
				}

				node = node.children[idx]

				continue
			}

			for _, child := range node.children {
				if child != nil && dfs(i+1, child) {
					return true
				}
			}

			return false
		}

		return node.endOfWord
	}

	return dfs(0, trie)
}
