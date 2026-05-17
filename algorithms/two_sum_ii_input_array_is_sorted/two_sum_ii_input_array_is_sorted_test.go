package twosumiiinputarrayissorted_test

import (
	"testing"

	twosumiiinputarrayissorted "github.com/r-erema/workshop/algorithms/two_sum_ii_input_array_is_sorted"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "Simple array",
			numbers: []int{1, 2, 10, 17},
			target:  12,
			want:    []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, twosumiiinputarrayissorted.TwoSum(tt.numbers, tt.target))
		})
	}
}
