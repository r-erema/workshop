package insertionsort_test

import (
	"testing"

	insertionsort "github.com/r-erema/workshop/algorithms/insertion_sort"
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			insertionsort.InsertionSort(tt.array)
			assert.Equal(t, tt.want, tt.array)
		})
	}
}
