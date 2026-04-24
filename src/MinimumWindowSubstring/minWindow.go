package main

import (
	"fmt"
	"math"
)

// We want the smallest window in s that contains all characters of t (with the right counts).
// Instead of checking all substrings, we use a sliding window:
//
// - Expand the window by moving the right pointer r and adding characters into a window map.
// - Once the window has all required characters (i.e., it "covers" t), we try to shrink it from the left with pointer l to make it as small as possible while still valid.
//
// During this process, we keep track of the best (smallest) window seen so far.
// This way, we only scan each character at most two times, making it efficient and still easy to follow.
// Time Complexity: O(n + m) where n is the length of s and m is the total number of unique characters in the strings s and t.
// Space Complexity: O(n + m) for the character count maps.
func minWindowSlidingWindow(s string, t string) string {
	result := ""
	l := 0

	if len(s) < len(t) {
		return result
	}

	sRunes := []rune(s)
	tRunes := []rune(t)

	sCount := make(map[rune]int, len(sRunes))
	tCount := make(map[rune]int, len(sRunes))

	for i := 0; i < len(tRunes); i++ {
		tCount[tRunes[i]]++
	}

	have, need := 0, len(tCount)

	diff := math.MaxInt32

	for r := 0; r < len(sRunes); r++ {
		index := sRunes[r]

		sCount[index]++

		if tCount[index] > 0 && sCount[index] == tCount[index] {
			have++
		}

		//fmt.Println("l:", l, "r:", r, "have:", have, "need:", need, "diff:", diff, "result:", result, "s:", s, "t:", t)

		for have == need {
			if r-l < diff {
				diff = r - l
				result = string(sRunes[l : r+1])
			}

			sCount[sRunes[l]]--
			if tCount[sRunes[l]] > 0 && sCount[sRunes[l]] < tCount[sRunes[l]] {
				have--
			}
			l++
		}
	}

	return result
}

func IsSubset(subset, superset map[rune]int) bool {
	// A subset cannot be larger than the superset
	if len(subset) > len(superset) {
		return false
	}

	for key, val := range subset {
		superVal, ok := superset[key]
		// Key must exist and value must match
		if !ok || superVal < val {
			return false
		}
	}
	return true
}

// We want the smallest substring of s that contains all characters of t (with the right counts).
// The brute-force way is to try every possible substring of s and check whether it covers all the characters in t.
// For each starting index, we expand the end index and keep a frequency map for the current substring.
// Whenever the substring has all required characters, we see if it's the smallest one so far.
// This is simple to understand but very slow because we check many overlapping substrings.
// Time Complexity: O(n^2 * m) where n is the length of s and m is the total number of unique characters in the strings s and t
// Space Complexity: O(n + m) for the character count maps.
func minWindowBruteForce(s string, t string) string {
	result := ""

	if len(s) < len(t) {
		return result
	}

	sRunes := []rune(s)
	tRunes := []rune(t)

	tCount := make(map[rune]int, len(sRunes))

	for _, char := range tRunes {
		tCount[char]++
	}

	diff := len(sRunes) + 1

	for l := 0; l < len(sRunes); l++ {
		sCount := make(map[rune]int, len(sRunes))

		for r := l; r < len(sRunes); r++ {
			//fmt.Println("l:", l, "r:", r, "diff:", diff, "result:", result, "s:", s, "t:", t)
			sCount[sRunes[r]]++

			if IsSubset(tCount, sCount) && diff > r-l {
				diff = r - l
				result = string(sRunes[l : r+1])
			}

		}
	}

	return result
}

func main() {

	s, t := "OUZODYXAZV", "XYZ"
	fmt.Println(minWindowBruteForce(s, t))
	fmt.Println(minWindowSlidingWindow(s, t))

	s, t = "xyz", "xyz"
	fmt.Println(minWindowBruteForce(s, t))
	fmt.Println(minWindowSlidingWindow(s, t))

	s, t = "x", "xy"
	fmt.Println(minWindowBruteForce(s, t))
	fmt.Println(minWindowSlidingWindow(s, t))

	s, t = "ADOBECODEBANC", "ABC"
	fmt.Println(minWindowBruteForce(s, t))
	fmt.Println(minWindowSlidingWindow(s, t))

	s, t = "a", "a"
	fmt.Println(minWindowBruteForce(s, t))
	fmt.Println(minWindowSlidingWindow(s, t))

	s, t = "a", "aa"
	fmt.Println(minWindowBruteForce(s, t))
	fmt.Println(minWindowSlidingWindow(s, t))

	s, t = "aaaaaaaaaaaabbbbbcdd", "abcdd"
	fmt.Println(minWindowBruteForce(s, t))
	fmt.Println(minWindowSlidingWindow(s, t))
}
