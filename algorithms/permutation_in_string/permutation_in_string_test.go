package permutationinstring_test

import (
	"testing"

	permutationinstring "github.com/r-erema/workshop/algorithms/permutation_in_string"
	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input1, input2 string
		want           bool
	}{
		{
			name:   "permutation exists",
			input1: "a",
			input2: "ab",
			want:   true,
		},
		{
			name:   "permutation does not exist",
			input1: "ab",
			input2: "eidboaoo",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, permutationinstring.CheckInclusion(tt.input1, tt.input2))
		})
	}
}
