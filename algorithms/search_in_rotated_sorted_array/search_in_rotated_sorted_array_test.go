package searchinrotatedsortedarray_test

import (
	"testing"

	searchinrotatedsortedarray "github.com/r-erema/workshop/algorithms/search_in_rotated_sorted_array"
	"github.com/stretchr/testify/assert"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  []int
		target int
		want   int
	}{
		{
			name:   "target exists",
			input:  []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "target does not exist",
			input:  []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "target does not exist, input with 1 value",
			input:  []int{1},
			target: 0,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, searchinrotatedsortedarray.Search(tt.input, tt.target))
		})
	}
}
