package implementtrieprefixtree_test

import (
	"testing"

	implementtrieprefixtree "github.com/r-erema/workshop/algorithms/implement_trie_prefix_tree"
	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	t.Parallel()

	trie := implementtrieprefixtree.Constructor()
	trie.Insert("apple")
	assert.True(t, trie.Search("apple"))
	assert.True(t, trie.StartsWith("app"))
	assert.False(t, trie.Search("ape"))
	assert.False(t, trie.StartsWith("api"))
}
