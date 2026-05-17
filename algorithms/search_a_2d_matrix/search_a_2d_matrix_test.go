package searcha2dmatrix_test

import (
	"testing"

	searcha2dmatrix "github.com/r-erema/workshop/algorithms/search_a_2d_matrix"
	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name:   "target exists in matrix",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		{
			name:   "target does not exist in matrix",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			want:   false,
		},
		{
			name:   "target does not exist in 1 element matrix",
			matrix: [][]int{{1}},
			target: 1,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, searcha2dmatrix.SearchMatrix(tt.matrix, tt.target))
		})
	}
}
