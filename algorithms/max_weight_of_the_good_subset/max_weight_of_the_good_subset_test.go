package maxweightofthegoodsubset_test

import (
	"testing"

	maxweightofthegoodsubset "github.com/r-erema/workshop/algorithms/max_weight_of_the_good_subset"
	"github.com/stretchr/testify/assert"
)

func TestMaxWeight(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "subset [1 3 3 3 3]",
			arr:  []int{3, 3, 3, 1, 3, 7, 1},
			want: 13,
		},
		{
			name: "subset [7 15]",
			arr:  []int{1, 7, 3, 15, 2, 5, 2, 1, 4},
			want: 22,
		},
		{
			name: "subset [3 4 5 6 7]",
			arr:  []int{6, 2, 5, 1, 7, 4, 3},
			want: 25,
		},
		{
			name: "subset []",
			arr:  []int{},
			want: 0,
		},
		{
			name: "subset [4 11]",
			arr:  []int{4, 11},
			want: 15,
		},
		{
			name: "subset [4]",
			arr:  []int{4},
			want: 4,
		},
		{
			name: "subset [1 1 1 1 1 1 1 1 1 1 1 1 1 2]",
			arr:  []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 5, 6},
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, maxweightofthegoodsubset.MaxWeight(tt.arr))
		})
	}
}
