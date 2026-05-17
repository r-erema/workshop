package branchsums_test

import (
	"testing"

	branchsums "github.com/r-erema/workshop/algorithms/branch_sums"
	"github.com/r-erema/workshop/utils/data_structure/tree"
	"github.com/stretchr/testify/assert"
)

func TestFindClosestValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		bst  *tree.Node
		want []int
	}{
		{
			name: "Short tree",
			bst: &tree.Node{
				Val:  5,
				Left: &tree.Node{Val: 2},
				Right: &tree.Node{
					Val:   10,
					Left:  &tree.Node{Val: 8},
					Right: &tree.Node{Val: 34},
				},
			},
			want: []int{7, 23, 49},
		},
		{
			name: "Long tree",
			bst: &tree.Node{
				Val: 9,
				Left: &tree.Node{
					Val:  4,
					Left: &tree.Node{Val: 3},
					Right: &tree.Node{
						Val:   6,
						Left:  &tree.Node{Val: 5},
						Right: &tree.Node{Val: 7},
					},
				},
				Right: &tree.Node{
					Val: 17,
					Right: &tree.Node{
						Val:   22,
						Left:  &tree.Node{Val: 20},
						Right: &tree.Node{Val: 23},
					},
				},
			},
			want: []int{16, 24, 26, 68, 71},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sums := branchsums.BranchSums(tt.bst)
			assert.Equal(t, tt.want, sums)
		})
	}
}
