package diameterofbinarytree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// DiameterOfBinaryTree calculates the diameter of a binary tree.
// Time O(N), since we iterate input one time
// Space O(N), due to the maximum depth of the recursion stack.
func DiameterOfBinaryTree(root *TreeNode) int {
	var diameter int

	var dfs func(root *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left, right := dfs(node.Left), dfs(node.Right)
		diameter = max(diameter, left+right)

		return 1 + max(left, right)
	}
	dfs(root)

	return diameter
}
