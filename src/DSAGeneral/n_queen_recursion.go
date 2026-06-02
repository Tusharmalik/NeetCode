package main

import "fmt"

type Point struct {
	row, col int
}

func main() {
	boardSize := 4

	board := make([][]bool, boardSize)
	obstacles := map[Point]int{}
	//var results [][][]bool

	for i := range board {
		board[i] = make([]bool, boardSize)
	}

	//fmt.Println(board)

	//nQueenRecursion(board, obstacles, 0)

	//fmt.Println(nQueenRecursionCount(board, 0))

	//_, count := nQueenRecursion2(board, obstacles, 0)
	//fmt.Println(count)

	nQueenRecursion3(board, obstacles, 0, 0, 0)
}

func nQueenRecursion(board [][]bool, obstacles map[Point]int, queenCount int) {
	boardSize := len(board)

	if queenCount == boardSize {
		display(board)
		return
	}

	for r := range board {
		tempCount := 0
		for c := range board[r] {
			if _, exists := obstacles[Point{r, c}]; !exists {
				board[r][c] = true
				obstacles = addObstacles(r, c, boardSize, obstacles)
				nQueenRecursion(board, obstacles, queenCount+1)
				board[r][c] = false
				obstacles = removeObstacles(r, c, boardSize, obstacles)
			}
			if !board[r][c] {
				tempCount++
			}
		}
		if tempCount == len(board[r]) {
			break
		}
	}
	return
}

func nQueenRecursionCount(board [][]bool, row int) int {
	if row == len(board) {
		display(board)
		return 1
	}

	count := 0

	for col := range board[row] {
		if isSafe(board, row, col) {
			board[row][col] = true
			count += nQueenRecursionCount(board, row+1)
			board[row][col] = false
		}
	}
	return count
}

func nQueenRecursion2(board [][]bool, obstacles map[Point]int, row int) ([][][]bool, int) {
	var results [][][]bool
	boardSize := len(board)
	count := 0

	if row == len(board) {
		display(board)
		results = append(results, board)
		return results, 1
	}

	for col := range board[row] {
		if _, exists := obstacles[Point{row, col}]; !exists {
			board[row][col] = true
			obstacles = addObstacles(row, col, boardSize, obstacles)
			response, localCount := nQueenRecursion2(board, obstacles, row+1)
			results = append(results, response...)
			count += localCount
			board[row][col] = false
			obstacles = removeObstacles(row, col, boardSize, obstacles)
		}
	}
	return results, count
}

func nQueenRecursion3(board [][]bool, obstacles map[Point]int, row int, col int, target int) {
	boardSize := len(board)

	if row == boardSize-1 && col == boardSize {
		return
	}

	if col == boardSize {
		nQueenRecursion3(board, obstacles, row+1, 0, target)
		return
	}

	if target == boardSize {
		display(board)
		return
	}

	if _, exists := obstacles[Point{row, col}]; !exists {
		board[row][col] = true
		obstacles = addObstacles(row, col, boardSize, obstacles)
		nQueenRecursion3(board, obstacles, row, col+1, target+1)
		board[row][col] = false
		obstacles = removeObstacles(row, col, boardSize, obstacles)
	}

	nQueenRecursion3(board, obstacles, row, col+1, target)
}

func addObstacles(row, col, boardSize int, obstacles map[Point]int) map[Point]int {

	for i := 0; i < boardSize; i++ {
		obstacles[Point{row, i}]++
		obstacles[Point{i, col}]++

		if row+i < boardSize && col+i < boardSize {
			obstacles[Point{row + i, col + i}]++
		}

		if row-i >= 0 && col-i >= 0 {
			obstacles[Point{row - i, col - i}]++
		}

		if row+i < boardSize && col-i >= 0 {
			obstacles[Point{row + i, col - i}]++
		}

		if row-i >= 0 && col+i < boardSize {
			obstacles[Point{row - i, col + i}]++
		}
	}
	return obstacles
}

func removeObstacles(row, col, boardSize int, obstacles map[Point]int) map[Point]int {
	for i := 0; i < boardSize; i++ {
		obstacles[Point{row, i}]--
		obstacles[Point{i, col}]--

		if row+i < boardSize && col+i < boardSize {
			obstacles[Point{row + i, col + i}]--
		}

		if row-i >= 0 && col-i >= 0 {
			obstacles[Point{row - i, col - i}]--
		}

		if row+i < boardSize && col-i >= 0 {
			obstacles[Point{row + i, col - i}]--
		}

		if row-i >= 0 && col+i < boardSize {
			obstacles[Point{row - i, col + i}]--
		}
	}

	for key, val := range obstacles {
		if val == 0 {
			delete(obstacles, key)
		}
	}
	return obstacles
}

func display(board [][]bool) {
	for _, row := range board {
		for _, col := range row {
			if !col {
				fmt.Printf(". ")
			} else {
				fmt.Printf("Q ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("------------------")
}

func isSafe(board [][]bool, row, col int) bool {
	// Check column
	for i := 0; i < row; i++ {
		if board[i][col] {
			return false
		}
	}

	// Check upper left diagonal
	maxLeft := min(row, col)
	for i := 1; i <= maxLeft; i++ {
		if board[row-i][col-i] {
			return false
		}
	}

	// Check upper right diagonal
	maxRight := min(row, len(board)-col-1)
	for i := 1; i <= maxRight; i++ {
		if board[row-i][col+i] {
			return false
		}
	}

	return true
}
