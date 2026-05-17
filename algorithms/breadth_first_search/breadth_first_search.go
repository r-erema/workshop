package breadthfirstsearch

func Bfs(nodeIndex int, edges [][]int, vertices []int) {
	queue := []int{nodeIndex}

	for len(queue) > 0 {
		nodeIndex, queue = queue[0], queue[1:]
		vertices[nodeIndex] = 1

		if nodeIndex < len(edges) {
			for _, edge := range edges[nodeIndex] {
				if vertices[edge] == 0 {
					queue = append(queue, edge)
				}
			}
		}
	}
}
