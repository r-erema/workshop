package combinationsumii_test

import (
	"testing"

	combinationsumii "github.com/r-erema/workshop/algorithms/combination_sum_ii"
	"github.com/stretchr/testify/assert"
)

func TestCombinationSum2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "1 result sets",
			candidates: []int{1},
			target:     1,
			want:       [][]int{{1}},
		},
		{
			name:       "1 result sets",
			candidates: []int{1, 2},
			target:     2,
			want:       [][]int{{2}},
		},
		{
			name:       "2 result sets",
			candidates: []int{1, 2, 3, 2},
			target:     3,
			want:       [][]int{{1, 2}, {3}},
		},
		{
			name:       "2 result sets",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			want:       [][]int{{1, 2, 2}, {5}},
		},
		{
			name:       "2 result sets",
			candidates: []int{1, 2, 0},
			target:     2,
			want:       [][]int{{0, 2}, {2}},
		},
		{
			name:       "4 result sets",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			want:       [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, combinationsumii.CombinationSum2(tt.candidates, tt.target))
		})
	}
}
