package subtreeofanothertree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// IsSubtree Time O(n*m), where n, m sizes of both trees
// Space O(1), we don't allocate additional memory.
func IsSubtree(root, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}

	if root == nil {
		return false
	}

	if SameTree(root, subRoot) {
		return true
	}

	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

func SameTree(tree1, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 != nil && tree2 != nil && tree1.Val == tree2.Val {
		return SameTree(tree1.Left, tree2.Left) && SameTree(tree1.Right, tree2.Right)
	}

	return false
}
