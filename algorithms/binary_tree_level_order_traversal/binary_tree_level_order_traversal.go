package binarytreelevelordertraversal

import (
	"github.com/r-erema/workshop/utils/data_structure/tree"
)

// LevelOrder performs level order traversal on a binary tree.
// Time O(n), since we iterate input one time
// Space O(n), since we need an array containing elements from input tree.
func LevelOrder(root *tree.Node) [][]int {
	if root == nil {
		return nil
	}

	var (
		res        [][]int
		levelQueue []*tree.Node
	)

	queue := [][]*tree.Node{{root}}

	for len(queue) > 0 {
		levelQueue, queue = queue[0], queue[1:]

		var (
			levelVals     []int
			newLevelQueue []*tree.Node
			node          *tree.Node
		)

		for len(levelQueue) > 0 {
			node, levelQueue = levelQueue[0], levelQueue[1:]
			levelVals = append(levelVals, node.Val)

			if node.Left != nil {
				newLevelQueue = append(newLevelQueue, node.Left)
			}

			if node.Right != nil {
				newLevelQueue = append(newLevelQueue, node.Right)
			}
		}

		res = append(res, levelVals)

		if len(newLevelQueue) > 0 {
			queue = append(queue, newLevelQueue)
		}
	}

	return res
}
