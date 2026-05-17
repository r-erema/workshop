package validatebinarysearchtree

import (
	"math"

	"github.com/r-erema/workshop/utils/data_structure/tree"
)

// IsValidBST Time O(n), since we iterate input one time
// Space O(n), due to the recursion stack, for a balanced tree, h is log(n), making the space complexity O(log(n)),
// in the worst case of a skewed tree, h is n, making the space complexity O(n).
func IsValidBST(root *tree.Node) bool {
	var dfs func(root *tree.Node, minLimit, maxLimit int) bool

	dfs = func(root *tree.Node, minLimit, maxLimit int) bool {
		if root == nil {
			return true
		}

		if root.Val <= minLimit || root.Val >= maxLimit {
			return false
		}

		return dfs(root.Left, minLimit, root.Val) && dfs(root.Right, root.Val, maxLimit)
	}

	return dfs(root, math.MinInt64, math.MaxInt64)
}
