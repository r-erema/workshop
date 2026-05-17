package merge2nondecreasingarrays_test

import (
	"testing"

	merge2nondecreasingarrays "github.com/r-erema/workshop/algorithms/merge_2_non_decreasing_arrays"
	"github.com/stretchr/testify/assert"
)

func TestMergeArrays(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		arr1, arr2 []int
		want       []int
	}{
		{
			name: "end on the first array",
			arr1: []int{-2, 3, 3, 22},
			arr2: []int{-5, 0},
			want: []int{-5, -2, 0, 3, 3, 22},
		},
		{
			name: "end on the second array",
			arr1: []int{-2, 3},
			arr2: []int{0, 8},
			want: []int{-2, 0, 3, 8},
		},
		{
			name: "first array is empty",
			arr1: []int{},
			arr2: []int{0},
			want: []int{0},
		},
		{
			name: "second array is empty",
			arr1: []int{-1},
			arr2: []int{},
			want: []int{-1},
		},
		{
			name: "both arrays are empty",
			arr1: []int{},
			arr2: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, merge2nondecreasingarrays.MergeArrays(tt.arr1, tt.arr2))
		})
	}
}
