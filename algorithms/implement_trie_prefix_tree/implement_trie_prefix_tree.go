package implementtrieprefixtree

type Trie struct {
	children  [26]*Trie
	endOfWord bool
}

func Constructor() Trie {
	return Trie{}
}

// Insert inserts a word into the trie.
// Time O(n), since we need iterate entire tree
// Space O(n), since we add a node at any letter.
func (trie *Trie) Insert(word string) {
	currNode := trie

	for i := range word {
		index := word[i] - 'a'
		if currNode.children[index] == nil {
			newChild := Constructor()
			currNode.children[index] = &newChild
			currNode = &newChild
		} else {
			currNode = currNode.children[index]
		}
	}

	currNode.endOfWord = true
}

// Search searches for a word in the trie.
// Time O(n), since we need iterate entire tree
// Space O(1), since we don't allocate additional memory.
func (trie *Trie) Search(word string) bool {
	currNode := trie

	for i := range word {
		index := word[i] - 'a'
		if currNode.children[index] == nil {
			return false
		}

		currNode = currNode.children[index]
	}

	return currNode.endOfWord
}

// StartsWith checks if any word in the trie starts with the given prefix.
// Time O(n), since we need iterate entire tree
// Space O(1), since we don't allocate additional memory.
func (trie *Trie) StartsWith(prefix string) bool {
	currNode := trie

	for i := range prefix {
		index := prefix[i] - 'a'
		if currNode.children[index] == nil {
			return false
		}

		currNode = currNode.children[index]
	}

	return true
}
