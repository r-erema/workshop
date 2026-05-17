package lowestcommonancestorofabinarysearchtree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		switch {
		case p.Val > root.Val && q.Val > root.Val:
			root = root.Right
		case p.Val < root.Val && q.Val < root.Val:
			root = root.Left
		default:
			return root
		}
	}

	return nil
}
