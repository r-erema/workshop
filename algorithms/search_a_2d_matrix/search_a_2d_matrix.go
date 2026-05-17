package searcha2dmatrix

// SearchMatrix searches for a target in a 2D matrix.
// Time O(log(N*M)), since we apply binary search to matrix rows and then apply binary search in the found row
// Space O(1), we don't involve any additional data structures.
func SearchMatrix(matrix [][]int, target int) bool {
	const divisor = 2

	var row int

	top, bottom := 0, len(matrix)-1

search:
	for top <= bottom {
		row = (top + bottom) / divisor

		switch {
		case target > matrix[row][len(matrix[0])-1]:
			top = row + 1
		case target < matrix[row][0]:
			bottom = row - 1
		default:
			break search
		}
	}

	left, right := 0, len(matrix[row])-1
	for left <= right {
		pointer := (left + right) / divisor

		switch cmp := matrix[row][pointer]; {
		case cmp < target:
			left = pointer + 1
		case cmp > target:
			right = pointer - 1
		default:
			return true
		}
	}

	return false
}
