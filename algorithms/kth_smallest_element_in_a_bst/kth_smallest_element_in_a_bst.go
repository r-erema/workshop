package kthsmallestelementinabst

import (
	"github.com/r-erema/workshop/utils/data_structure/tree"
)

// KthSmallest finds the kth smallest element in a BST.
// Time O(n), since we visit each node of the tree
// Space O(n), since the recursion stack grows as nodes count.
func KthSmallest(root *tree.Node, m int) int {
	var res int

	var inorderTraversal func(node *tree.Node)

	inorderTraversal = func(node *tree.Node) {
		if node == nil {
			return
		}

		inorderTraversal(node.Left)

		m--
		if m == 0 {
			res = node.Val

			return
		}

		inorderTraversal(node.Right)
	}

	inorderTraversal(root)

	return res
}
