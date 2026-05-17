package mincostclimbingstairs_test

import (
	"testing"

	mincostclimbingstairs "github.com/r-erema/workshop/algorithms/min_cost_climbing_stairs"
	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		cost []int
		want int
	}{
		{
			name: "3 costs",
			cost: []int{10, 15, 20},
			want: 15,
		},
		{
			name: "10 costs",
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, mincostclimbingstairs.MinCostClimbingStairs(tt.cost))
		})
	}
}
