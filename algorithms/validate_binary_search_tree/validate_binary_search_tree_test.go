package validatebinarysearchtree_test

import (
	"testing"

	validatebinarysearchtree "github.com/r-erema/workshop/algorithms/validate_binary_search_tree"
	"github.com/r-erema/workshop/utils/data_structure/tree"
	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		root *tree.Node
		want bool
	}{
		{
			name: "valid tree",
			root: &tree.Node{
				Val: 2,
				Left: &tree.Node{
					Val: 1,
				},
				Right: &tree.Node{
					Val: 3,
				},
			},
			want: true,
		},
		{
			name: "invalid tree",
			root: &tree.Node{
				Val: 5,
				Left: &tree.Node{
					Val: 1,
				},
				Right: &tree.Node{
					Val: 4,
					Left: &tree.Node{
						Val: 3,
					},
					Right: &tree.Node{
						Val: 6,
					},
				},
			},
			want: false,
		},
		{
			name: "invalid tree 2",
			root: &tree.Node{
				Val: 1,
				Left: &tree.Node{
					Val: 1,
				},
			},
			want: false,
		},
		{
			name: "invalid tree 3",
			root: &tree.Node{
				Val: 5,
				Left: &tree.Node{
					Val: 4,
				},
				Right: &tree.Node{
					Val:   6,
					Left:  &tree.Node{Val: 3},
					Right: &tree.Node{Val: 7},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, validatebinarysearchtree.IsValidBST(tt.root))
		})
	}
}
