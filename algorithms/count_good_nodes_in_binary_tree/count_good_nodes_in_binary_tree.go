package countgoodnodesinbinarytree

import (
	"github.com/r-erema/workshop/utils/data_structure/tree"
)

// GoodNodes counts the number of good nodes in a binary tree.
// Time O(n), since we iterate input one time
// Space O(n), due to the recursion stack.
func GoodNodes(root *tree.Node) int {
	var (
		good int
		dfs  func(root *tree.Node, maxVal int)
	)

	dfs = func(root *tree.Node, maxVal int) {
		if root == nil {
			return
		}

		if root.Val >= maxVal {
			good++
			maxVal = root.Val
		}

		dfs(root.Left, maxVal)
		dfs(root.Right, maxVal)
	}
	dfs(root, root.Val)

	return good
}
