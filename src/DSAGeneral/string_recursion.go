package main

import "fmt"

func main() {
	input := "ba💙ccad💙"
	input = "abcd"
	//fmt.Println(removeChar(input, 'a', 0, len(input), ""))
	//fmt.Println(removeChar1(input, '💙', ""))
	//possibleSubstrings(input, "")
	fmt.Println(possibleSubsets(input, ""))
}

func removeChar(input string, remove byte, start, end int, result string) string {
	localResult := result

	if start == end {
		return localResult
	}

	if input[start] != remove {
		localResult += string(input[start])
	}
	return removeChar(input, remove, start+1, end, localResult)
}

func removeChar1(input string, remove rune, result string) string {
	localResult := result
	input1 := []rune(input)

	if input == "" {
		return localResult
	}

	if input1[0] != remove {
		localResult += string(input1[0])
	}
	return removeChar1(string(input1[1:]), remove, localResult)
}

func possibleSubstrings(input string, result string) {
	if input == "" {
		if result != "" {
			fmt.Println(result)
		}
		return
	}

	input1 := []rune(input)

	possibleSubstrings(string(input1[1:]), result)
	result += string(input1[0])
	possibleSubstrings(string(input1[1:]), result)
}

func possibleSubsets(input string, result string) ([]string, int) {
	var list []string
	count := 0
	if input == "" {
		list = append(list, result)
		return list, 1
	}

	input1 := []rune(input)

	for i := 0; i < len(result)+1; i++ {
		response, localCount := possibleSubsets(string(input1[1:]), result[0:i]+string(input1[0])+result[i:])
		list = append(list, response...)
		count += localCount
	}
	return list, count
}
