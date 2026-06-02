package main

import "fmt"

func main() {
	board := [][]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}

	//board = [][]int{
	//	{1, 2, 0, 0, 3, 0, 0, 0, 0},
	//	{4, 0, 0, 5, 0, 0, 0, 0, 0},
	//	{0, 9, 0, 0, 0, 0, 0, 0, 3},
	//	{5, 0, 0, 0, 6, 0, 0, 0, 4},
	//	{0, 0, 0, 8, 0, 3, 0, 0, 5},
	//	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	//	{0, 0, 0, 0, 0, 0, 2, 0, 0},
	//	{0, 0, 0, 4, 1, 9, 0, 0, 8},
	//	{0, 0, 0, 0, 8, 0, 0, 7, 9},
	//}

	if sudokuSolve(board) {
		displaySudoku(board)
	} else {
		fmt.Println("No solution exists")
	}
}

func sudokuSolve(board [][]int) bool {
	boardSize := len(board)

	row, col := -1, -1

	emptyLeft := true

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == 0 {
				row, col = i, j
				emptyLeft = false
				break
			}
		}
		// If you found some empty element in row then break
		if !emptyLeft {
			break
		}
	}

	// Sudoku is solved
	if emptyLeft {
		return true
	}

	// backtrack
	for num := 1; num <= boardSize; num++ {
		if isSafeSudoku(board, row, col, num) {
			board[row][col] = num

			if sudokuSolve(board) {
				return true
			}

			board[row][col] = 0
		}
	}

	return false
}

func displaySudoku(board [][]int) {
	for _, row := range board {
		for _, col := range row {
			fmt.Print(col, " ")
		}
		fmt.Println()
	}
	fmt.Println("==============")
}

func isSafeSudoku(board [][]int, row, col, num int) bool {
	boardSize := len(board)

	// Check row
	for i := 0; i < boardSize; i++ {
		if board[row][i] == num {
			return false
		}
	}

	// Check col
	for i := 0; i < boardSize; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// Check 3x3 box
	boxSize := 3
	boxRowStart := row - (row % boxSize)
	boxColStart := col - (col % boxSize)

	for i := boxRowStart; i < boxRowStart+boxSize; i++ {
		for j := boxColStart; j < boxColStart+boxSize; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}
