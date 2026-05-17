package validanagram_test

import (
	"testing"

	validanagram "github.com/r-erema/workshop/algorithms/valid_anagram"
	"github.com/stretchr/testify/assert"
)

func TestValidAnagram(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name, word1, word2 string
		want               bool
	}{
		{
			name:  "anagram",
			word1: "anagram",
			word2: "nagaram",
			want:  true,
		},
		{
			name:  "not anagram",
			word1: "rat",
			word2: "car",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, validanagram.IsAnagram(tt.word1, tt.word2))
		})
	}
}
