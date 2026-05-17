package wordsearch_test

import (
	"testing"

	wordsearch "github.com/r-erema/workshop/algorithms/word_search"
	"github.com/stretchr/testify/assert"
)

func TestWordSearch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			name: "word 0",
			board: [][]byte{
				{'A', 'B'},
			},
			word: "AB",
			want: true,
		},
		{
			name: "word 0",
			board: [][]byte{
				{'A'},
				{'B'},
			},
			word: "AB",
			want: true,
		},
		{
			name: "word 1",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ESE",
			want: true,
		},
		{
			name: "word 2",
			board: [][]byte{
				{'B', 'G'},
				{'G', 'S'},
				{'S', 'A'},
			},
			word: "BG",
			want: true,
		},
		{
			name: "word 3",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'E', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCESEEEFS",
			want: true,
		},
		{
			name: "word 4",
			board: [][]byte{
				{'A', 'A'},
				{'A', 'A'},
			},
			word: "AAAAA",
			want: false,
		},
		{
			name: "word 5",
			board: [][]byte{
				{'A', 'A'},
				{'A', 'A'},
			},
			word: "AAAB",
			want: false,
		},
		{
			name: "word 6",
			board: [][]byte{
				{'a', 'a', 'b', 'a', 'a', 'b'},
				{'a', 'a', 'b', 'b', 'b', 'a'},
				{'a', 'a', 'a', 'a', 'b', 'a'},
				{'b', 'a', 'b', 'b', 'a', 'b'},
				{'a', 'b', 'b', 'a', 'b', 'a'},
				{'b', 'a', 'a', 'a', 'a', 'b'},
			},
			word: "bbbaabbbbbab",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, wordsearch.Exist(tt.board, tt.word))
		})
	}
}
