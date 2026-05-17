package maxareaofisland_test

import (
	"testing"

	maxareaofisland "github.com/r-erema/workshop/algorithms/max_area_of_island"
	"github.com/stretchr/testify/assert"
)

func TestMaxAreaOfIsland(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "6 islands",
			grid: [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, maxareaofisland.MaxAreaOfIsland(tt.grid))
		})
	}
}
