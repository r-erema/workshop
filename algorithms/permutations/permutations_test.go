package permutations_test

import (
	"testing"

	"github.com/r-erema/workshop/algorithms/permutations"
	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "2 subsets",
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "3 subsets",
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, permutations.Permute(tt.nums))
		})
	}
}
