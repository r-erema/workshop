package countingbits_test

import (
	"testing"

	countingbits "github.com/r-erema/workshop/algorithms/counting_bits"
	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		number int
		want   []int
	}{
		{
			name:   "number 2",
			number: 2,
			want:   []int{0, 1, 1},
		},
		{
			name:   "number 5",
			number: 5,
			want:   []int{0, 1, 1, 2, 1, 2},
		},
		{
			name:   "number 8",
			number: 8,
			want:   []int{0, 1, 1, 2, 1, 2, 2, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, countingbits.CountBits(tt.number))
		})
	}
}
