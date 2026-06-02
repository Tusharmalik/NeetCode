package main

import "fmt"

func main() {
	boardSize := 4

	board := make([][]bool, boardSize)

	for i := range board {
		board[i] = make([]bool, boardSize)
	}

	nKnight(board, 0, 0, 4)

}

func nKnight(board [][]bool, row, col, target int) {

	if row == len(board)-1 && col == len(board[0]) {
		return
	}

	if col == len(board) {
		nKnight(board, row+1, 0, target)
		return
	}

	if target == 0 {
		displayKnight(board)
		return
	}

	if isSafeKnight(board, row, col) {
		board[row][col] = true
		nKnight(board, row, col+1, target-1)
		board[row][col] = false
	}

	nKnight(board, row, col+1, target)

}

func displayKnight(board [][]bool) {
	for _, row := range board {
		for _, col := range row {
			if !col {
				fmt.Printf(".  ")
			} else {
				fmt.Printf("K ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("------------------")
}

func isSafeKnight(board [][]bool, row, col int) bool {
	if row-1 >= 0 && col-2 >= 0 && board[row-1][col-2] {
		return false
	}

	if row-1 >= 0 && col+2 < len(board) && board[row-1][col+2] {
		return false
	}

	if row-2 >= 0 && col-1 >= 0 && board[row-2][col-1] {
		return false
	}

	if row-2 >= 0 && col+1 < len(board) && board[row-2][col+1] {
		return false
	}

	if row+1 < len(board) && col+2 < len(board[0]) && board[row+1][col+2] {
		return false
	}

	if row+1 < len(board) && col-2 >= 0 && board[row+1][col-2] {
		return false
	}

	if row+2 < len(board) && col+1 < len(board[0]) && board[row+2][col+1] {
		return false
	}

	if row+2 < len(board) && col-1 >= 0 && board[row+2][col-1] {
		return false
	}

	return true
}
