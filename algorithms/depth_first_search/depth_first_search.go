package depth_first_search

// Time O(N), since we iterate each node one time
// Space O(N), since we store every visited node.

import "slices"

func DFSRecursive(graph [][]int, startIndex int) map[int]int {
	var (
		recursive    func(int)
		visitedOrder = make(map[int]int)
		step         = 0
	)

	recursive = func(index int) {
		if _, ok := visitedOrder[index]; ok {
			return
		}

		step++

		visitedOrder[index] = step
		for i := range graph[index] {
			recursive(graph[index][i])
		}
	}

	recursive(startIndex)

	return visitedOrder
}

func DFSStack(graph [][]int, startIndex int) map[int]int {
	var (
		curr         int
		step         = 0
		stack        = []int{startIndex}
		visitedOrder = make(map[int]int)
	)

	for len(stack) > 0 {
		curr, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if _, ok := visitedOrder[curr]; ok {
			continue
		}

		step++
		visitedOrder[curr] = step

		for _, v := range slices.Backward(graph[curr]) {
			stack = append(stack, v)
		}
	}

	return visitedOrder
}
