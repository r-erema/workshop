package nodedepth

import (
	"github.com/r-erema/workshop/utils/data_structure/bst"
)

// NodeDepthRecursively calculates the sum of all node depths recursively.
// Worst: O(n) time | O(h) space, h - height if the binary tree.
func NodeDepthRecursively(bst *bst.BST) int {
	return RecursionHelper(bst, 0)
}

func RecursionHelper(node *bst.BST, currentDepth int) int {
	var depthLeft, depthRight int

	if node.Left() != nil {
		depthLeft = RecursionHelper(node.Left(), currentDepth+1)
	}

	if node.Right() != nil {
		depthRight = RecursionHelper(node.Right(), currentDepth+1)
	}

	return currentDepth + depthLeft + depthRight
}

// NodeDepthIterative calculates the sum of all node depths iteratively.
// Worst: O(n) time | O(h) space, h - height if the binary tree.
func NodeDepthIterative(bstNode *bst.BST) int {
	var depth int

	depthNodes := [][]*bst.BST{{bstNode}}

	for level := 0; len(depthNodes) > level; level++ {
		var (
			levelNodes      []*bst.BST
			levelNodesCount int
		)

		for _, node := range depthNodes[level] {
			if left := node.Left(); left != nil {
				levelNodes = append(levelNodes, left)
				levelNodesCount++
			}

			if right := node.Right(); right != nil {
				levelNodes = append(levelNodes, right)
				levelNodesCount++
			}
		}

		if len(levelNodes) > 0 {
			depthNodes = append(depthNodes, levelNodes)
			depth += (level + 1) * levelNodesCount
		}
	}

	return depth
}

// NodeDepthIterative2 calculates the sum of all node depths iteratively (second implementation).
// Worst: O(n) time | O(h) space, h - height if the binary tree.
func NodeDepthIterative2(bstNode *bst.BST) int {
	type stackItem struct {
		node  *bst.BST
		depth int
	}

	var totalDepth int

	stack := []stackItem{{node: bstNode, depth: 0}}

	for len(stack) > 0 {
		item := stack[0]
		stack = stack[1:]
		node, depth := item.node, item.depth

		if node == nil {
			continue
		}

		stack = append(
			stack,
			stackItem{node: node.Left(), depth: depth + 1},
			stackItem{node: node.Right(), depth: depth + 1},
		)

		totalDepth += depth
	}

	return totalDepth
}
