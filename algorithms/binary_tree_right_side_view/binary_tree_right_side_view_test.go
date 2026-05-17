package binarytreerightsideview_test

import (
	"testing"

	binarytreerightsideview "github.com/r-erema/workshop/algorithms/binary_tree_right_side_view"
	"github.com/r-erema/workshop/utils/data_structure/tree"
	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		root *tree.Node
		want []int
	}{
		{
			name: "see 3 nodes",
			root: &tree.Node{
				Val: 1,
				Left: &tree.Node{
					Val:   2,
					Right: &tree.Node{Val: 5},
				},
				Right: &tree.Node{
					Val:   3,
					Right: &tree.Node{Val: 4},
				},
			},
			want: []int{1, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, binarytreerightsideview.RightSideView(tt.root))
		})
	}
}
