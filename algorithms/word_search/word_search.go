package wordsearch

// Exist Time O(m * 4^n), since we can start from any cell and explore all 4 directions
// Space O(n), since we can store the whole word in the visited map.
func Exist(board [][]byte, word string) bool {
	if len(word) > len(board)*len(board[0]) {
		return false
	}

	var (
		res     bool
		dfs     func(i, dirX, dirY int)
		visited = make(map[[2]int]struct{})
	)

	dfs = func(i, dirX, dirY int) {
		if res || !IsValidPosition(dirX, dirY, board) || word[i] != board[dirX][dirY] {
			return
		}

		key := [2]int{dirX, dirY}
		if _, exists := visited[key]; exists {
			return
		}

		visited[key] = struct{}{}
		defer delete(visited, key)

		if i == len(word)-1 {
			res = true

			return
		}

		neighbors := [][]int{
			{dirX, dirY + 1},
			{dirX - 1, dirY},
			{dirX, dirY - 1},
			{dirX + 1, dirY},
		}
		for _, neighbor := range neighbors {
			dfs(i+1, neighbor[0], neighbor[1])
		}
	}

	for x := range board {
		for y := range board[x] {
			dfs(0, x, y)
		}
	}

	return res
}

func IsValidPosition(dirX, dirY int, board [][]byte) bool {
	return dirX >= 0 && dirX < len(board) && dirY >= 0 && dirY < len(board[0])
}
