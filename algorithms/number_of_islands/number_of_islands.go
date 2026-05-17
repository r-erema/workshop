package numberofislands

// NumIslands counts the number of islands in a grid.
// Time O(N*M)
// where M and N are rows and columns,
// We need to visit each cell in a grid
//
// Space O(min(M, N))
// We need to queue and deque elements from time to time.
// the max size of the queue can be min(n,m).
func NumIslands(grid [][]byte) int {
	rows, cols, islandsCount := len(grid), len(grid[0]), 0

	for i := range rows {
		for j := range cols {
			if grid[i][j] == '1' {
				BreadthFirstSearch(grid, i, j)

				islandsCount++
			}
		}
	}

	return islandsCount
}

func BreadthFirstSearch(grid [][]byte, startCoordinateX, startCoordinateY int) {
	queue := [][2]int{{startCoordinateX, startCoordinateY}}
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		startCoordinateX, startCoordinateY, queue = queue[0][0], queue[0][1], queue[1:]

		for _, direction := range directions {
			x, y := startCoordinateX+direction[0], startCoordinateY+direction[1]
			if x >= 0 && y >= 0 && len(grid) > x && len(grid[0]) > y && grid[x][y] == '1' {
				queue = append(queue, [2]int{x, y})
				grid[x][y] = '2'
			}
		}
	}
}
