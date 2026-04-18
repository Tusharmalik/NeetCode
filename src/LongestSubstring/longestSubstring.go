package main

import (
	"fmt"
)

func lengthOfLongestSubstringSlidingWindowOptimal(s string) int {
	res := 0
	l := 0
	hash := make(map[int32]int)
	runes := []rune(s)

	for r, char := range runes {
		if idx, ok := hash[char]; ok {
			l = max(l, idx+1)
		}

		hash[char] = r
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res
}

func lengthOfLongestSubstringSlidingWindow(s string) int {
	hash := make(map[int32]struct{})
	res := 0
	runes := []rune(s)
	l, r := 0, 0

	for _, char := range runes {
		if _, ok := hash[char]; !ok {
			hash[char] = struct{}{}
			r++
		} else {
			for l < r {
				if runes[l] != char {
					delete(hash, runes[l])
					l++
				} else {
					l++
					r++
					break
				}
			}
		}

		if r-l > res {
			res = r - l
		}

	}
	return res
}

func lengthOfLongestSubstringBruteForce(s string) int {
	res := 0

	for i := range s {
		start := 0
		hash := make(map[int32]struct{})
		for _, ch := range s[i:] {
			if _, ok := hash[ch]; !ok {
				hash[ch] = struct{}{}
				start += 1
			} else {
				break
			}
		}
		if start > res {
			res = start
		}
	}
	return res
}

func main() {
	input := "dvdf"
	//input = "zxyzxyz"
	fmt.Println(lengthOfLongestSubstringBruteForce(input))

	input = "pwwkew"
	fmt.Println(lengthOfLongestSubstringBruteForce(input))

	input = " "
	fmt.Println(lengthOfLongestSubstringBruteForce(input))

	input = "dvdf"
	fmt.Println(lengthOfLongestSubstringSlidingWindowOptimal(input))

	input = "pwwkew"
	fmt.Println(lengthOfLongestSubstringSlidingWindowOptimal(input))

	input = " "
	fmt.Println(lengthOfLongestSubstringSlidingWindowOptimal(input))
}
