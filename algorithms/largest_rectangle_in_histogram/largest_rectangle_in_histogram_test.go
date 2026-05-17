package largestrectangleinhistogram_test

import (
	"testing"

	largestrectangleinhistogram "github.com/r-erema/workshop/algorithms/largest_rectangle_in_histogram"
	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleInHistogram(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "Normal plot",
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, largestrectangleinhistogram.LargestRectangleArea(tt.heights))
		})
	}
}
