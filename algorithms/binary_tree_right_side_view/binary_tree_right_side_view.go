package binarytreerightsideview

import (
	"github.com/r-erema/workshop/utils/data_structure/tree"
)

// RightSideView returns the right side view of a binary tree.
// Time O(n), since we iterate input one time
// Space O(n), due to the recursion stack.
func RightSideView(root *tree.Node) []int {
	var res []int

	var preorderTraversal func(root *tree.Node, level int)

	preorderTraversal = func(root *tree.Node, level int) {
		if root == nil {
			return
		}

		if len(res) == level {
			res = append(res, root.Val)
		}

		preorderTraversal(root.Right, level+1)
		preorderTraversal(root.Left, level+1)
	}

	preorderTraversal(root, 0)

	return res
}
