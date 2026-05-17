package maximumdepthofbinarytree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func MaxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(MaxDepthDFS(root.Left), MaxDepthDFS(root.Right))
}

func MaxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		queueLen := len(queue)
		for i := range queueLen {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[queueLen:]
		level++
	}

	return level
}

func MaxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type stackRow struct {
		node  *TreeNode
		depth int
	}

	stack := []stackRow{
		{
			node:  root,
			depth: 1,
		},
	}

	var result int

	var row stackRow

	for len(stack) > 0 {
		row, stack = stack[0], stack[1:]

		if row.node != nil {
			result = max(result, row.depth)
			stack = append(
				stack,
				stackRow{node: row.node.Left, depth: row.depth + 1},
				stackRow{node: row.node.Right, depth: row.depth + 1},
			)
		}
	}

	return result
}
