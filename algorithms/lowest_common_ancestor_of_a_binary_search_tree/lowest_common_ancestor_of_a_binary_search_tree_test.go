package lowestcommonancestorofabinarysearchtree_test

import (
	"testing"

	lowestcommonancestorofabinarysearchtree "github.com/r-erema/workshop/algorithms/lowest_common_ancestor_of_a_binary_search_tree"
	"github.com/stretchr/testify/assert"
)

// Time O(log n), we don't have to visit every node in the tree, but only 1 node in a level
// Space O(1), we don't allocate additional memory.
func TestLowestCommonAncestor(t *testing.T) {
	t.Parallel()

	commonTree := &lowestcommonancestorofabinarysearchtree.TreeNode{
		Val: 6,
		Left: &lowestcommonancestorofabinarysearchtree.TreeNode{
			Val:  2,
			Left: &lowestcommonancestorofabinarysearchtree.TreeNode{Val: 0},
			Right: &lowestcommonancestorofabinarysearchtree.TreeNode{
				Val:   4,
				Left:  &lowestcommonancestorofabinarysearchtree.TreeNode{Val: 3},
				Right: &lowestcommonancestorofabinarysearchtree.TreeNode{Val: 5},
			},
		},
		Right: &lowestcommonancestorofabinarysearchtree.TreeNode{
			Val:   8,
			Left:  &lowestcommonancestorofabinarysearchtree.TreeNode{Val: 7},
			Right: &lowestcommonancestorofabinarysearchtree.TreeNode{Val: 9},
		},
	}

	tests := []struct {
		name string
		p, q *lowestcommonancestorofabinarysearchtree.TreeNode
		want int
	}{
		{
			name: "lca is root node",
			p:    &lowestcommonancestorofabinarysearchtree.TreeNode{Val: 2},
			q:    &lowestcommonancestorofabinarysearchtree.TreeNode{Val: 8},
			want: 6,
		},
		{
			name: "lca is node itself",
			p:    &lowestcommonancestorofabinarysearchtree.TreeNode{Val: 2},
			q:    &lowestcommonancestorofabinarysearchtree.TreeNode{Val: 4},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(
				t,
				tt.want,
				lowestcommonancestorofabinarysearchtree.LowestCommonAncestor(commonTree, tt.p, tt.q).Val,
			)
		})
	}
}
