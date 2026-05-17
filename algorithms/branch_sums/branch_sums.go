package branchsums

import (
	"github.com/r-erema/workshop/utils/data_structure/tree"
)

// BranchSums calculates the sum of all branch paths in a binary tree.
// Time O(n), since we visit each node of the tree
// Space O(n), since the recursion stack grows as nodes count.
func BranchSums(bst *tree.Node) []int {
	var sums []int

	var helper func(node *tree.Node)

	helper = func(node *tree.Node) {
		if node == nil {
			return
		}

		if node.Left != nil {
			node.Left.Val += node.Val
		}

		if node.Right != nil {
			node.Right.Val += node.Val
		}

		if node.Left == nil && node.Right == nil {
			sums = append(sums, node.Val)
		}

		helper(node.Left)
		helper(node.Right)
	}

	helper(bst)

	return sums
}
