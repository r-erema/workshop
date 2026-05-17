package validsudoku

// IsValidSudoku Time O(1) or O(9^2) we iterate consequently on each cell
// Time O(1) we allocate memory for maps equals to input.
func IsValidSudoku(board [][]byte) bool {
	const subSquareDivisor = 3

	rows := make([]map[byte]struct{}, len(board))
	cols := make([]map[byte]struct{}, len(board))
	subSquares := make(map[[2]int]map[byte]struct{}, len(board)*len(board[0])/subSquareDivisor)

	for i := range board {
		for j := range len(board[0]) {
			if board[i][j] == '.' {
				continue
			}

			_, duplicateInRow := rows[i][board[i][j]]
			_, duplicateInCol := cols[j][board[i][j]]
			_, duplicateInSubSquare := subSquares[[2]int{i / 3, j / 3}][board[i][j]]

			if duplicateInRow || duplicateInCol || duplicateInSubSquare {
				return false
			}

			if rows[i] == nil {
				rows[i] = make(map[byte]struct{})
			}

			if cols[j] == nil {
				cols[j] = make(map[byte]struct{})
			}

			if subSquares[[2]int{i / 3, j / 3}] == nil {
				subSquares[[2]int{i / 3, j / 3}] = make(map[byte]struct{})
			}

			rows[i][board[i][j]] = struct{}{}
			cols[j][board[i][j]] = struct{}{}
			subSquares[[2]int{i / 3, j / 3}][board[i][j]] = struct{}{}
		}
	}

	return true
}
