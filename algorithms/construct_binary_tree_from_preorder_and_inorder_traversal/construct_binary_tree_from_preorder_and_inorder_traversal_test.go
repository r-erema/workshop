package constructbinarytreefrompreorderandinordertraversal_test

import (
	"testing"

	constructbinarytreefrompreorderandinordertraversal "github.com/r-erema/workshop/algorithms/construct_binary_tree_from_preorder_and_inorder_traversal"
	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		preorder, inorder []int
		want              *constructbinarytreefrompreorderandinordertraversal.TreeNode
	}{
		{
			name:     "Normal tree",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want: &constructbinarytreefrompreorderandinordertraversal.TreeNode{
				Val:  3,
				Left: &constructbinarytreefrompreorderandinordertraversal.TreeNode{Val: 9},
				Right: &constructbinarytreefrompreorderandinordertraversal.TreeNode{
					Val:   20,
					Left:  &constructbinarytreefrompreorderandinordertraversal.TreeNode{Val: 15},
					Right: &constructbinarytreefrompreorderandinordertraversal.TreeNode{Val: 7},
				},
			},
		},
		{
			name:     "Normal tree 2",
			preorder: []int{2, 8, 11, 7, 6, 12, 3, 5},
			inorder:  []int{11, 8, 6, 7, 2, 3, 5, 12},
			want: &constructbinarytreefrompreorderandinordertraversal.TreeNode{
				Val: 2,
				Left: &constructbinarytreefrompreorderandinordertraversal.TreeNode{
					Val:  8,
					Left: &constructbinarytreefrompreorderandinordertraversal.TreeNode{Val: 11},
					Right: &constructbinarytreefrompreorderandinordertraversal.TreeNode{
						Val:  7,
						Left: &constructbinarytreefrompreorderandinordertraversal.TreeNode{Val: 6},
					},
				},
				Right: &constructbinarytreefrompreorderandinordertraversal.TreeNode{
					Val: 12,
					Left: &constructbinarytreefrompreorderandinordertraversal.TreeNode{
						Val:   3,
						Right: &constructbinarytreefrompreorderandinordertraversal.TreeNode{Val: 5},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(
				t,
				tt.want,
				constructbinarytreefrompreorderandinordertraversal.BuildTree(tt.preorder, tt.inorder),
			)
		})
	}
}
