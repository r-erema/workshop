package quicksort_test

import (
	"testing"

	quicksort "github.com/r-erema/workshop/algorithms/quick_sort"
	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{
			name:  "Case 0",
			array: []int{7, 0, 4},
			want:  []int{0, 4, 7},
		},
		{
			name:  "Case 1",
			array: []int{8, 5, 2, 9, 5, 6, 3},
			want:  []int{2, 3, 5, 5, 6, 8, 9},
		},
		{
			name:  "Case 3",
			array: []int{4, 3, 1, 2, 5, 9, 7, 6, 10},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			quicksort.QuickSort(tt.array)
			assert.Equal(t, tt.want, tt.array)
		})
	}
}
