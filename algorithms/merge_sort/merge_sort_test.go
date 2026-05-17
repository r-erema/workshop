package mergesort_test

import (
	"testing"

	mergesort "github.com/r-erema/workshop/algorithms/merge_sort"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "simple array",
			arr:  []int{5, 3},
			want: []int{3, 5},
		},
		{
			name: "simple array",
			arr:  []int{5, 3, 2, 1},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "array has not unique values",
			arr:  []int{5, 1, 1, 2, 0, 0},
			want: []int{0, 0, 1, 1, 2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, mergesort.MergeSort(tt.arr))
		})
	}
}
