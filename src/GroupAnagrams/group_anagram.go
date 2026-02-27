package main

import "fmt"

// groupAnagrams groups strings that are anagrams of each other.
// It returns a 2D slice where each inner slice contains strings that are anagrams.
//
// Algorithm: Uses a nested loop approach to compare each string with all remaining strings.
// When anagrams are found, they are added to the same group and removed from the array
// to prevent duplicate processing.
//
// Time Complexity: O(n^2 * m) where n is the number of strings and m is the average string length
// Space Complexity: O(n) for storing the result
//
// Example:
//
//	Input: []string{"eat", "tea", "tan", "ate", "nat", "bat"}
//	Output: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
func groupAnagrams(strs []string) [][]string {
	// Initialize result slice to store all anagram groups
	var combos [][]string

	// Outer loop: iterate through each string as the base of a new group
	for i := 0; i < len(strs); i++ {
		// Start a new group with the current string
		groupList := []string{strs[i]}

		// Inner loop: compare current string with all strings that come after it
		for j := i + 1; j < len(strs); j++ {
			// Check if the current string at position j is an anagram of strs[i]
			if isAnagram(strs[i], strs[j]) {
				// Add the anagram to the current group
				groupList = append(groupList, strs[j])
				// Remove the matched string from the array to avoid processing it again
				strs = append(strs[:j], strs[j+1:]...)
				// Decrement j because we just removed an element at position j
				// The next iteration will check the new string that shifted into position j
				j--
			}
		}

		// Add the completed group to the result
		combos = append(combos, groupList)
	}

	return combos
}

// isAnagram checks if two strings are anagrams of each other.
// Two strings are anagrams if they contain the same characters with the same frequencies.
// Time Complexity: O(m) where m is the length of the strings
// Space Complexity: O(k) where k is the number of unique characters
func isAnagram(a string, b string) bool {
	// Quick check: anagrams must have the same length
	if len(a) != len(b) {
		return false
	}

	// Map to track character frequencies
	count := make(map[rune]int)

	// Count frequency of each character in string 'a'
	for _, char := range a {
		count[char]++
	}

	// Decrement count for each character in string 'b'
	for _, char := range b {
		count[char]--
		// If count goes negative, 'b' has more of this character than 'a'
		if count[char] < 0 {
			return false
		}
	}

	// If we reach here, both strings have the same character frequencies
	return true
}

func main() {
	// Test case 1: Multiple anagram groups
	// Expected: [["eat","tea","ate"], ["tan","nat"], ["bat"]]
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(result)

	// Test case 2: Three anagram groups
	// Expected: [["act","cat"], ["pots","tops","stop"], ["hat"]]
	result = groupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"})
	fmt.Println(result)

	// Test case 3: No anagrams (different lengths)
	// Expected: [["abct"], ["cat"]]
	result = groupAnagrams([]string{"abct", "cat"})
	fmt.Println(result)

}
