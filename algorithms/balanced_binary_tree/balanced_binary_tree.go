package balancedbinarytree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsBalanced checks if a binary tree is balanced.
// Time O(N), since we iterate input one time
// Space O(1), since we don't involve any additional data structure.
func IsBalanced(root *TreeNode) bool {
	type res struct {
		isBalanced bool
		level      float64
	}

	var dfs func(root *TreeNode) res

	dfs = func(root *TreeNode) res {
		if root == nil {
			return res{true, 0}
		}

		left, right := dfs(root.Left), dfs(root.Right)
		balanced := left.isBalanced && right.isBalanced && math.Abs(left.level-right.level) <= 1

		return res{balanced, 1 + max(left.level, right.level)}
	}

	return dfs(root).isBalanced
}
