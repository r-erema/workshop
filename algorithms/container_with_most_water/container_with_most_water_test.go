package containerwithmostwater_test

import (
	"testing"

	containerwithmostwater "github.com/r-erema/workshop/algorithms/container_with_most_water"
	"github.com/stretchr/testify/assert"
)

func TestContainerWithMostWater(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "9 heights",
			heights: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:    49,
		},
		{
			name:    "2 heights",
			heights: []int{1, 1},
			want:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, containerwithmostwater.MaxArea(tt.heights))
		})
	}
}
