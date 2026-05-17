package subsetsii_test

import (
	"testing"

	subsetsii "github.com/r-erema/workshop/algorithms/subsets_ii"
	"github.com/stretchr/testify/assert"
)

func TestSubsetsII(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "2 subsets",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			name: "4 subsets",
			nums: []int{1, 2},
			want: [][]int{{}, {1}, {1, 2}, {2}},
		},
		{
			name: "6 subsets",
			nums: []int{1, 2, 2},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, subsetsii.SubsetsWithDup(tt.nums))
		})
	}
}
