package constructbinarytreefrompreorderandinordertraversal

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// BuildTree builds a binary tree from preorder and inorder traversals.
// Time O(n), since each node in the tree is processed exactly once
// Space O(n), for the following reasons:
//   - hashmap d stores N key-value pairs representing each node's value and index in the inorder traversal
//   - the recursion stack may grow up to O(n) in the case of a skewed tree (where each node has only one child)
//   - the output structure, which is a binary tree that contains N TreeNode instances.
func BuildTree(preorder, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i := range inorder {
		inorderMap[inorder[i]] = i
	}

	var helper func(preorderStart, inorderStart, size int) *TreeNode

	helper = func(preorderStart, inorderStart, size int) *TreeNode {
		if size == 0 {
			return nil
		}

		root := &TreeNode{Val: preorder[preorderStart]}
		mid := inorderMap[root.Val]
		leftSubtreeSize := mid - inorderStart

		root.Left = helper(preorderStart+1, inorderStart, leftSubtreeSize)
		root.Right = helper(preorderStart+leftSubtreeSize+1, mid+1, size-1-leftSubtreeSize)

		return root
	}

	return helper(0, 0, len(inorder))
}
