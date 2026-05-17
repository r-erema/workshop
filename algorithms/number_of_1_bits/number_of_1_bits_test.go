package numberof1bits_test

import (
	"testing"

	numberof1bits "github.com/r-erema/workshop/algorithms/number_of_1_bits"
	"github.com/stretchr/testify/assert"
)

func TestHammingWeight(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input, want int
	}{
		{
			name:  "3 ones",
			input: 11,
			want:  3,
		},
		{
			name:  "1 one",
			input: 128,
			want:  1,
		},
		{
			name:  "30 ones",
			input: 2147483645,
			want:  30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, numberof1bits.HammingWeight(tt.input))
			assert.Equal(t, tt.want, numberof1bits.HammingWeight2(tt.input))
			assert.Equal(t, tt.want, numberof1bits.HammingWeight3(tt.input))
		})
	}
}
