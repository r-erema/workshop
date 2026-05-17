package plusone_test

import (
	"testing"

	plusone "github.com/r-erema/workshop/algorithms/plus_one"
	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input, want []int
	}{
		{
			name:  "no increasing the next digit",
			input: []int{1, 2, 3},
			want:  []int{1, 2, 4},
		},
		{
			name:  "increasing the next digit",
			input: []int{4, 3, 9, 9},
			want:  []int{4, 4, 0, 0},
		},
		{
			name:  "adding one more digit",
			input: []int{9, 9},
			want:  []int{1, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, plusone.PlusOne(tt.input))
		})
	}
}
