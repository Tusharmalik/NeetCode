package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(diceProblem(6, ""))
}

func letterCombinations(digits string) []string {
	result := letterCombinations2(digits, "")
	return result
}

func letterCombinations2(digits string, result string) []string {
	var combinationsList []string

	digitsMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	if digits == "" {
		combinationsList = append(combinationsList, result)
		return combinationsList
	}

	for _, val := range digitsMap[string(digits[0])] {
		localList := letterCombinations2(digits[1:], result+val)
		combinationsList = append(combinationsList, localList...)
	}
	return combinationsList
}

func diceProblem(number int, result string) []string {
	possibleList := make([]string, 0)
	if number == 0 {
		possibleList = append(possibleList, result)
		return possibleList
	}

	if number < 0 {
		return possibleList
	}

	possibleNumbers := []int{1, 2, 3, 4, 5, 6}

	for _, val := range possibleNumbers {
		if val <= number {
			innerList := diceProblem(number-val, result+strconv.Itoa(val))
			possibleList = append(possibleList, innerList...)
		}
	}
	return possibleList
}
