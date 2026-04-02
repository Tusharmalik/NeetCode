package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ValidPalindrome(s string) bool {
	reg, _ := regexp.Compile(`[^a-zA-Z0-9]`)
	s = reg.ReplaceAllString(s, "")

	s = strings.ToLower(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true

}

func main() {
	input := "Was it a car or a cat I saw?"
	fmt.Println(ValidPalindrome(input))

	input = "tab a cat"
	fmt.Println(ValidPalindrome(input))
}
