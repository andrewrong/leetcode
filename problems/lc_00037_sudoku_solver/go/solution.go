package solution

// SolveSudoku solves a Sudoku puzzle by filling the empty cells.
// Each of the digits 1-9 must occur exactly once in each row,
// each column, and each of the 9 3x3 sub-boxes of the grid.
// The '.' character indicates empty cells.
func SolveSudoku(board [][]byte) {
	// Track which numbers are already used in rows, columns, and boxes
	rows := make([][10]bool, 9)
	cols := make([][10]bool, 9)
	boxes := make([][10]bool, 9)

	// Initialize with existing numbers on the board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				boxIndex := (i/3)*3 + j/3
				rows[i][num] = true
				cols[j][num] = true
				boxes[boxIndex][num] = true
			}
		}
	}

	backtrack(board, rows, cols, boxes)
}

func backtrack(board [][]byte, rows, cols, boxes [][10]bool) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				boxIndex := (i/3)*3 + j/3

				// Try numbers 1-9
				for num := 1; num <= 9; num++ {
					// Check if number is valid in this position
					if !rows[i][num] && !cols[j][num] && !boxes[boxIndex][num] {
						// Place the number
						board[i][j] = byte(num + '0')
						rows[i][num] = true
						cols[j][num] = true
						boxes[boxIndex][num] = true

						// Recurse
						if backtrack(board, rows, cols, boxes) {
							return true
						}

						// Backtrack
						board[i][j] = '.'
						rows[i][num] = false
						cols[j][num] = false
						boxes[boxIndex][num] = false
					}
				}
				return false // No valid number found
			}
		}
	}
	return true // Board is complete
}