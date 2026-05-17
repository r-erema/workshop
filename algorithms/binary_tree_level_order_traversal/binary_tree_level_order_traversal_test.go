package binarytreelevelordertraversal_test

import (
	"testing"

	binarytreelevelordertraversal "github.com/r-erema/workshop/algorithms/binary_tree_level_order_traversal"
	"github.com/r-erema/workshop/utils/data_structure/tree"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		root *tree.Node
		want [][]int
	}{
		{
			name: "normal case",
			root: &tree.Node{
				Val:  3,
				Left: &tree.Node{Val: 9},
				Right: &tree.Node{
					Val:   20,
					Left:  &tree.Node{Val: 15},
					Right: &tree.Node{Val: 7},
				},
			},
			want: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			name: "1 node",
			root: &tree.Node{
				Val: 1,
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "no nodes",
			root: nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, binarytreelevelordertraversal.LevelOrder(tt.root))
		})
	}
}
