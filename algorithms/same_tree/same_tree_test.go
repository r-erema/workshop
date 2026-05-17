package sametree_test

import (
	"testing"

	sametree "github.com/r-erema/workshop/algorithms/same_tree"
	"github.com/r-erema/workshop/utils/data_structure/tree"
	"github.com/stretchr/testify/assert"
)

func TestSameTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		tree1, tree2 *tree.Node
		want         bool
	}{
		{
			name: "same tree",
			tree1: &tree.Node{
				Val: 1,
				Left: &tree.Node{
					Val: 2,
				},
				Right: &tree.Node{
					Val: 3,
				},
			},
			tree2: &tree.Node{
				Val: 1,
				Left: &tree.Node{
					Val: 2,
				},
				Right: &tree.Node{
					Val: 3,
				},
			},
			want: true,
		},
		{
			name: "not same tree with 2 nodes",
			tree1: &tree.Node{
				Val: 1,
				Left: &tree.Node{
					Val: 2,
				},
			},
			tree2: &tree.Node{
				Val: 1,
				Right: &tree.Node{
					Val: 2,
				},
			},
			want: false,
		},
		{
			name: "not same tree with 3 nodes",
			tree1: &tree.Node{
				Val: 1,
				Left: &tree.Node{
					Val: 2,
				},
				Right: &tree.Node{
					Val: 1,
				},
			},
			tree2: &tree.Node{
				Val: 1,
				Left: &tree.Node{
					Val: 1,
				},
				Right: &tree.Node{
					Val: 2,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, sametree.IsSameTree(tt.tree1, tt.tree2))
		})
	}
}
