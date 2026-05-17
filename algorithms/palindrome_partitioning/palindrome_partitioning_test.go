package palindromepartitioning_test

import (
	"testing"

	palindromepartitioning "github.com/r-erema/workshop/algorithms/palindrome_partitioning"
	"github.com/stretchr/testify/assert"
)

func TestPalindromePartitioning(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{
			name: "1 palindrome case",
			s:    "a",
			want: [][]string{{"a"}},
		},
		{
			name: "2 palindromes case",
			s:    "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, palindromepartitioning.Partition(tt.s))
		})
	}
}
