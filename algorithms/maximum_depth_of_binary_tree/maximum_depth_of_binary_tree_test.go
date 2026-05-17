package maximumdepthofbinarytree_test

import (
	"testing"

	maximumdepthofbinarytree "github.com/r-erema/workshop/algorithms/maximum_depth_of_binary_tree"
	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		root *maximumdepthofbinarytree.TreeNode
		want int
	}{
		{
			name: "3-tier tree",
			root: &maximumdepthofbinarytree.TreeNode{
				Val: 3,
				Left: &maximumdepthofbinarytree.TreeNode{
					Val: 9,
				},
				Right: &maximumdepthofbinarytree.TreeNode{
					Val: 20,
					Left: &maximumdepthofbinarytree.TreeNode{
						Val: 15,
					},
					Right: &maximumdepthofbinarytree.TreeNode{
						Val: 7,
					},
				},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, maximumdepthofbinarytree.MaxDepthDFS(tt.root))
			assert.Equal(t, tt.want, maximumdepthofbinarytree.MaxDepthBFS(tt.root))
			assert.Equal(t, tt.want, maximumdepthofbinarytree.MaxDepthRecursive(tt.root))
		})
	}
}
