package combinationsum_test

import (
	"testing"

	combinationsum "github.com/r-erema/workshop/algorithms/combination_sum"
	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "basic case",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "basic case 2",
			candidates: []int{2, 4, 8},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 2, 4}, {4, 4}, {8}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, combinationsum.CombinationSum(tt.candidates, tt.target))
		})
	}
}
