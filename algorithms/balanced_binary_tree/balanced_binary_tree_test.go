package balancedbinarytree_test

import (
	"testing"

	balancedbinarytree "github.com/r-erema/workshop/algorithms/balanced_binary_tree"
	"github.com/stretchr/testify/assert"
)

func TestBalancedTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		root *balancedbinarytree.TreeNode
		want bool
	}{
		{
			name: "Balanced tree",
			root: &balancedbinarytree.TreeNode{
				Val: 3,
				Left: &balancedbinarytree.TreeNode{
					Val: 9,
				},
				Right: &balancedbinarytree.TreeNode{
					Val: 20,
					Left: &balancedbinarytree.TreeNode{
						Val: 15,
					},
					Right: &balancedbinarytree.TreeNode{
						Val: 7,
					},
				},
			},
			want: true,
		},
		{
			name: "Non-balanced tree",
			root: &balancedbinarytree.TreeNode{
				Val: 1,
				Left: &balancedbinarytree.TreeNode{
					Val: 2,
					Left: &balancedbinarytree.TreeNode{
						Val: 3,
						Left: &balancedbinarytree.TreeNode{
							Val: 4,
						},
						Right: &balancedbinarytree.TreeNode{
							Val: 4,
						},
					},
					Right: &balancedbinarytree.TreeNode{
						Val: 3,
					},
				},
				Right: &balancedbinarytree.TreeNode{
					Val: 2,
				},
			},
			want: false,
		},
		{
			name: "Non-balanced tree, only 1 branch",
			root: &balancedbinarytree.TreeNode{
				Val: 1,
				Right: &balancedbinarytree.TreeNode{
					Val: 2,
					Right: &balancedbinarytree.TreeNode{
						Val: 3,
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, balancedbinarytree.IsBalanced(tt.root))
		})
	}
}
