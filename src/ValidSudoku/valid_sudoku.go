package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	hashMap1 := make(map[string]map[int]map[byte]struct{})

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {

			ok := setNestedHashmap(hashMap1, "rows", i, board[i][j])
			if !ok {
				return false
			}
			ok = setNestedHashmap(hashMap1, "cols", i, board[j][i])
			if !ok {
				return false
			}
			ok = setNestedHashmap(hashMap1, "boxes", i, board[i/3*3+j/3][i%3*3+j%3])
			if !ok {
				return false
			}
		}
	}

	return true
}

func setNestedHashmap(m1 map[string]map[int]map[byte]struct{}, key1 string, key2 int, key3 byte) bool {
	if key3 == '.' {
		return true
	}

	if _, ok := m1[key1]; !ok {
		m1[key1] = make(map[int]map[byte]struct{})
	}

	if _, ok := m1[key1][key2]; !ok {
		m1[key1][key2] = make(map[byte]struct{})
	}

	if _, ok := m1[key1][key2][key3]; ok {
		return false
	}

	m1[key1][key2][key3] = struct{}{}

	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '8', '8', '6', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board = [][]byte{
		{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
		{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '9', '1', '.', '.', '.', '.', '.', '3'},
		{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
		{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board = [][]byte{
		{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
		{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '.', '3'},
		{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
		{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	for _, row := range board {
		for _, col := range row {
			fmt.Printf("%v, ", string(col))
		}
		fmt.Println()
	}

	fmt.Println("============")

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%v, ", string(board[j][i]))
		}
		fmt.Println()
	}

	fmt.Println("============")

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			//fmt.Println(i, j, i/3*3+j/3, i%3*3+j%3)
			fmt.Printf("%v, ", string(board[i/3*3+j/3][i%3*3+j%3]))
		}
		fmt.Println()
	}
	fmt.Println("============")

	fmt.Println(isValidSudoku(board))
}
