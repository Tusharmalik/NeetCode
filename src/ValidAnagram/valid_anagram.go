package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	for _, char := range s {
		count[char] = count[char] + 1
	}

	for _, char := range t {
		if count[char] >= 1 {
			count[char] = count[char] - 1
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("bbcc", "ccbc"))
}
