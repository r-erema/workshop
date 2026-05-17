package maxareaofisland

func MaxAreaOfIsland(grid [][]int) int {
	visited := make(map[[2]int]struct{}, len(grid)*len(grid[0]))

	var maxArea int

	for i := range grid {
		for j := range grid[i] {
			if _, ok := visited[[2]int{i, j}]; ok || grid[i][j] == 0 {
				continue
			}

			visited[[2]int{i, j}] = struct{}{}

			currArea := Bfs(grid, [2]int{i, j}, visited)

			maxArea = max(maxArea, currArea)
		}
	}

	return maxArea
}

func Bfs(grid [][]int, startPoint [2]int, visited map[[2]int]struct{}) int {
	var (
		queue      = [][2]int{startPoint}
		curr       [2]int
		currArea   = 0
		directions = [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	)

	for len(queue) > 0 {
		currArea++

		curr, queue = queue[0], queue[1:]

		for i := range directions {
			directionPoint := [2]int{curr[0] + directions[i][0], curr[1] + directions[i][1]}

			if _, ok := visited[directionPoint]; ok {
				continue
			}

			pointValid := directionPoint[0] < len(grid) &&
				directionPoint[1] < len(grid[0]) &&
				directionPoint[0] >= 0 &&
				directionPoint[1] >= 0

			if pointValid && grid[directionPoint[0]][directionPoint[1]] == 1 {
				visited[directionPoint] = struct{}{}

				queue = append(queue, directionPoint)
			}
		}
	}

	return currArea
}
