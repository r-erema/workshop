package invertbinarytree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// InvertTree inverts a binary tree by swapping left and right children.
// Time O(n), we have to visit every node in the tree once to swap its left and right children
// Space O(n), due to the maximum depth of the recursion stack.
func InvertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	}

	return root
}
