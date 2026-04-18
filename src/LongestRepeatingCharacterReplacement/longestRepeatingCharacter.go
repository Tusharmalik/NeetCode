package main

import "fmt"

// Copied this code from internet, still figuring out what is happening here.
// Note: Need to code this again
func characterReplacement(s string, k int) int {
	res := 0
	runes := []rune(s)
	hash := make(map[rune]struct{})

	for _, char := range runes {
		primeChar := char
		count := 0
		l := 0

		if _, ok := hash[char]; ok {
			continue
		}
		hash[primeChar] = struct{}{}

		for r, ch := range runes {
			if primeChar == ch {
				count++
			}

			for (r-l+1)-count > k {
				if runes[l] == primeChar {
					count--
				}
				l++
			}

			res = max(res, r-l+1)
			//fmt.Println("l:", l, "r:", r, "count:", count, "res:", res, "primeChar:", primeChar, "ch:", ch)
		}
	}
	return res
}

func main() {
	input, k := "XYYX", 2

	fmt.Println(characterReplacement(input, k))

	input, k = "AAABABB", 1
	fmt.Println(characterReplacement(input, k))

	input, k = "ABBB", 2
	fmt.Println(characterReplacement(input, k))

	input, k = "AABABBB", 1
	fmt.Println(characterReplacement(input, k))
}
