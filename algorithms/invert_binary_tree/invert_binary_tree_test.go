package invertbinarytree_test

import (
	"testing"

	invertbinarytree "github.com/r-erema/workshop/algorithms/invert_binary_tree"
	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		root, want *invertbinarytree.TreeNode
	}{
		{
			name: "3-tier tree",
			root: &invertbinarytree.TreeNode{
				Val: 4,
				Left: &invertbinarytree.TreeNode{
					Val: 2,
					Left: &invertbinarytree.TreeNode{
						Val: 1,
					},
					Right: &invertbinarytree.TreeNode{
						Val: 3,
					},
				},
				Right: &invertbinarytree.TreeNode{
					Val: 7,
					Left: &invertbinarytree.TreeNode{
						Val: 6,
					},
					Right: &invertbinarytree.TreeNode{
						Val: 9,
					},
				},
			},
			want: &invertbinarytree.TreeNode{
				Val: 4,
				Left: &invertbinarytree.TreeNode{
					Val: 7,
					Left: &invertbinarytree.TreeNode{
						Val: 9,
					},
					Right: &invertbinarytree.TreeNode{
						Val: 6,
					},
				},
				Right: &invertbinarytree.TreeNode{
					Val: 2,
					Left: &invertbinarytree.TreeNode{
						Val: 3,
					},
					Right: &invertbinarytree.TreeNode{
						Val: 1,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, invertbinarytree.InvertTree(tt.root))
		})
	}
}
