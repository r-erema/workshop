package subsets_test

import (
	"testing"

	"github.com/r-erema/workshop/algorithms/subsets"
	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "8 subsets",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {3}, {2}, {2, 3}, {1}, {1, 3}, {1, 2}, {1, 2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, subsets.Subsets(tt.nums))
		})
	}
}
