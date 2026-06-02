package main

import "fmt"

func main() {
	//fmt.Println(possiblePathCount(2, 2))
	//fmt.Println(possiblePath(2, 2, ""))
	//fmt.Println(possiblePathObstacle(2, 2, "", [][]int{{1, 1}}))

	//maze := [][]bool{{true, true, true}, {true, true, true}, {true, true, true}}
	mazeSize := []int{3, 3}
	start := []int{0, 0}
	end := []int{2, 2}

	maze := make([][]bool, mazeSize[0])
	stepPath := make([][]int, mazeSize[0])

	for i := range maze {
		maze[i] = make([]bool, mazeSize[1])
		for j := range maze[i] {
			maze[i][j] = true
		}
	}

	for i := range stepPath {
		stepPath[i] = make([]int, mazeSize[1])
	}

	fmt.Println(maze, len(maze), len(maze[0]))

	fmt.Println(backtracking(start, end, maze, "", stepPath, 1))
}

func possiblePathCount(row, col int) int {
	count := 0
	if row == 0 || col == 0 {
		return 1
	}

	count = count + possiblePathCount(row-1, col)
	count = count + possiblePathCount(row, col-1)

	return count
}

func possiblePath(row, col int, result string) []string {
	var path []string

	if row == 0 && col == 0 {
		path = append(path, result)
		return path
	}

	if col > 0 && row > 0 {
		path = append(path, possiblePath(row-1, col-1, result+"Diagonal")...)
	}

	if col > 0 {
		path = append(path, possiblePath(row, col-1, result+"Right")...)
	}

	if row > 0 {
		path = append(path, possiblePath(row-1, col, result+"Down")...)
	}

	return path
}

func possiblePathObstacle(row, col int, result string, obstacle [][]int) []string {
	var path []string

	if row == 0 && col == 0 {
		path = append(path, result)
		return path
	}

	move := true
	for _, obs := range obstacle {
		if row == obs[0] && col == obs[1] {
			move = false
			break
		}
	}

	if !move {
		return path
	}

	if col > 0 && row > 0 {
		path = append(path, possiblePathObstacle(row-1, col-1, result+"Diagonal", obstacle)...)
	}

	if col > 0 {
		path = append(path, possiblePathObstacle(row, col-1, result+"Right", obstacle)...)
	}

	if row > 0 {
		path = append(path, possiblePathObstacle(row-1, col, result+"Down", obstacle)...)
	}

	return path
}

func backtracking(start, end []int, maze [][]bool, result string, stepPath [][]int, step int) []string {
	var path []string
	currentRow, currentCol := start[0], start[1]
	targetRow, targetCol := end[0], end[1]
	totalRow, totalCol := len(maze)-1, len(maze[0])-1

	if currentRow == targetRow && currentCol == targetCol {
		stepPath[currentRow][currentCol] = step
		for _, col := range stepPath {
			fmt.Println(col)
		}
		fmt.Println(result, step)
		fmt.Println()
		path = append(path, result)
		return path
	}

	if !maze[currentRow][currentCol] {
		return path
	}
	maze[currentRow][currentCol] = false
	stepPath[currentRow][currentCol] = step

	if currentRow < totalRow {
		localPath := backtracking([]int{currentRow + 1, currentCol}, end, maze, result+"D", stepPath, step+1)
		path = append(path, localPath...)
	}

	if currentCol < totalCol {
		localPath := backtracking([]int{currentRow, currentCol + 1}, end, maze, result+"R", stepPath, step+1)
		path = append(path, localPath...)
	}

	if currentRow > 0 {
		localPath := backtracking([]int{currentRow - 1, currentCol}, end, maze, result+"U", stepPath, step+1)
		path = append(path, localPath...)
	}

	if currentCol > 0 {
		localPath := backtracking([]int{currentRow, currentCol - 1}, end, maze, result+"L", stepPath, step+1)
		path = append(path, localPath...)
	}
	maze[currentRow][currentCol] = true
	stepPath[currentRow][currentCol] = 0

	return path
}
