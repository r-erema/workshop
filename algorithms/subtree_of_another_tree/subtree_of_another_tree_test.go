package subtreeofanothertree_test

import (
	"testing"

	subtreeofanothertree "github.com/r-erema/workshop/algorithms/subtree_of_another_tree"
	"github.com/stretchr/testify/assert"
)

func TestIsSubtree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		root, subRoot *subtreeofanothertree.TreeNode
		want          bool
	}{
		{
			name: "There is subtree",
			root: &subtreeofanothertree.TreeNode{
				Val: 3,
				Left: &subtreeofanothertree.TreeNode{
					Val: 4,
					Left: &subtreeofanothertree.TreeNode{
						Val: 1,
					},
					Right: &subtreeofanothertree.TreeNode{
						Val: 2,
					},
				},
				Right: &subtreeofanothertree.TreeNode{
					Val: 5,
				},
			},
			subRoot: &subtreeofanothertree.TreeNode{
				Val: 4,
				Left: &subtreeofanothertree.TreeNode{
					Val: 1,
				},
				Right: &subtreeofanothertree.TreeNode{
					Val: 2,
				},
			},
			want: true,
		},
		{
			name: "There is no subtree",
			root: &subtreeofanothertree.TreeNode{
				Val: 3,
				Left: &subtreeofanothertree.TreeNode{
					Val: 4,
					Left: &subtreeofanothertree.TreeNode{
						Val: 1,
					},
					Right: &subtreeofanothertree.TreeNode{
						Val: 2,
						Left: &subtreeofanothertree.TreeNode{
							Val: 0,
						},
					},
				},
				Right: &subtreeofanothertree.TreeNode{
					Val: 5,
				},
			},
			subRoot: &subtreeofanothertree.TreeNode{
				Val: 4,
				Left: &subtreeofanothertree.TreeNode{
					Val: 1,
				},
				Right: &subtreeofanothertree.TreeNode{
					Val: 2,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, subtreeofanothertree.IsSubtree(tt.root, tt.subRoot))
		})
	}
}
