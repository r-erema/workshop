package topkfrequentwords_test

import (
	"testing"

	topkfrequentwords "github.com/r-erema/workshop/algorithms/top_k_frequent_words"
	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		words []string
		k     int
		want  []string
	}{
		{
			name:  "Two words",
			words: []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:     2,
			want:  []string{"i", "love"},
		},
		{
			name:  "Three words",
			words: []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:     3,
			want:  []string{"i", "love", "coding"},
		},
		{
			name:  "Four words",
			words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
			k:     4,
			want:  []string{"the", "is", "sunny", "day"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, topkfrequentwords.TopKFrequent(tt.words, tt.k))
		})
	}
}
