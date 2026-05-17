package closestvalueinbst

import (
	"math"
)

type TreeNode struct {
	Val         float64
	Left, Right *TreeNode
}

// ClosestValueInBST finds the closest value to target in a BST.
// Time O(log(N)), since we explore a particular subtree
// Space O(1), since we don't involve any additional data structure.
func ClosestValueInBST(root *TreeNode, target float64) float64 {
	closest := root.Val

loop:
	for root != nil {
		if math.Abs(root.Val-target) < math.Abs(closest-target) {
			closest = root.Val
		}

		switch {
		case target < root.Val:
			root = root.Left
		case target > root.Val:
			root = root.Right
		default:
			break loop
		}
	}

	return closest
}
