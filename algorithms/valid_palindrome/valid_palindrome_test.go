package validpalindrome_test

import (
	"testing"

	validpalindrome "github.com/r-erema/workshop/algorithms/valid_palindrome"
	"github.com/stretchr/testify/assert"
)

func TestEncodeAndDecodeStrings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Valid palindrome",
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},
		{
			name:  "Not palindrome",
			input: "race a car",
			want:  false,
		},
		{
			name:  "Palindrome: no alphanumerical symbols",
			input: " .",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, validpalindrome.IsPalindrome(tt.input))
		})
	}
}
