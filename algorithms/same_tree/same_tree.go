package sametree

import (
	"github.com/r-erema/workshop/utils/data_structure/tree"
)

// IsSameTree checks if two trees are identical.
// Time O(p+q), since we should iterate recursively through the both trees
// Space O(n), in case of degenerate tree.
func IsSameTree(tree1, tree2 *tree.Node) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 == nil || tree2 == nil || tree1.Val != tree2.Val {
		return false
	}

	return IsSameTree(tree1.Left, tree2.Left) && IsSameTree(tree1.Right, tree2.Right)
}
