package longest_palindromic_substring_test

import (
	"testing"

	"github.com/r-erema/workshop/algorithms/longest_palindromic_substring"
	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "3 letters palindrome",
			input: "babad",
			want:  "bab",
		},
		{
			name:  "2 letters palindrome",
			input: "cbbd",
			want:  "bb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, longest_palindromic_substring.LongestPalindrome(tt.input))
		})
	}
}
