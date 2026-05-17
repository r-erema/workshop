package missingnumber_test

import (
	"testing"

	missingnumber "github.com/r-erema/workshop/algorithms/missing_number"
	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "number 2 is missed",
			input: []int{3, 0, 1},
			want:  2,
		},
		{
			name:  "number 1 is missed",
			input: []int{0},
			want:  1,
		},
		{
			name:  "number 8 is missed",
			input: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want:  8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, missingnumber.MissingNumber(tt.input))
			assert.Equal(t, tt.want, missingnumber.MissingNumber2(tt.input))
		})
	}
}
