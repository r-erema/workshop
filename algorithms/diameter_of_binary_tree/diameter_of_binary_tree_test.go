package diameterofbinarytree_test

import (
	"testing"

	diameterofbinarytree "github.com/r-erema/workshop/algorithms/diameter_of_binary_tree"
	"github.com/stretchr/testify/assert"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		root *diameterofbinarytree.TreeNode
		want int
	}{
		{
			name: "Diameter 3",
			root: &diameterofbinarytree.TreeNode{
				Val: 1,
				Left: &diameterofbinarytree.TreeNode{
					Val: 2,
					Left: &diameterofbinarytree.TreeNode{
						Val: 4,
					},
					Right: &diameterofbinarytree.TreeNode{
						Val: 5,
					},
				},
				Right: &diameterofbinarytree.TreeNode{
					Val: 3,
				},
			},
			want: 3,
		},
		{
			name: "Diameter 2",
			root: &diameterofbinarytree.TreeNode{
				Val: 2,
				Left: &diameterofbinarytree.TreeNode{
					Val: 3,
					Left: &diameterofbinarytree.TreeNode{
						Val: 1,
					},
					Right: nil,
				},
				Right: nil,
			},
			want: 2,
		},
		{
			name: "Diameter 1",
			root: &diameterofbinarytree.TreeNode{
				Val: 1,
				Left: &diameterofbinarytree.TreeNode{
					Val: 2,
				},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, diameterofbinarytree.DiameterOfBinaryTree(tt.root))
		})
	}
}
