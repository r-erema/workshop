package lettercombinationsofaphonenumber_test

import (
	"testing"

	lettercombinationsofaphonenumber "github.com/r-erema/workshop/algorithms/letter_combinations_of_a_phone_number"
	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{
			name:   "2 digits",
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, lettercombinationsofaphonenumber.LetterCombinations(tt.digits))
		})
	}
}
