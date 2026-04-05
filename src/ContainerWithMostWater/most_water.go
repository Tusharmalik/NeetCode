package main

import "fmt"

// We try every possible pair of lines and compute the area they form.
// For each pair (i, j), the height of the container is the shorter of
// the two lines, and the width is the distance between them.
// Time Complexity: O(n^2) where n is the number of lines.
// Space Complexity: O(1) since we only use a constant amount of extra space.
func maxAreaBruteForce(heights []int) int {
	var maxArea int

	for i := 0; i < len(heights); i++ {
		for j := i + 1; j < len(heights); j++ {
			temp := min(heights[i], heights[j]) * (j - i)
			if temp > maxArea {
				maxArea = temp
			}
		}
	}
	return maxArea
}

// Using two pointers lets us efficiently search for the maximum area without checking every pair.
// We start with the widest container (left at start, right at end).
// The height is limited by the shorter line, so to potentially increase the area, we must move
// the pointer at the shorter line inward.
// Moving the taller line never helps because it keeps the height the same but reduces the width.
// By always moving the shorter side, we explore all meaningful possibilities.
// Time Complexity: O(n) where n is the number of lines, since we traverse the array at most once.
// Space Complexity: O(1) since we only use a constant amount of extra space.
func maxAreaTwoPointer(heights []int) int {
	var maxWater int

	left, right := 0, len(heights)-1

	for left < right {
		smaller := min(heights[left], heights[right])
		distance := right - left
		temp := smaller * distance

		if temp > maxWater {
			maxWater = temp
		}

		if heights[left] == smaller {
			left++
		} else {
			right--
		}
	}
	return maxWater
}

func main() {
	input := []int{1, 7, 2, 5, 4, 7, 3, 6}
	fmt.Println(maxAreaTwoPointer(input))
	fmt.Println(maxAreaBruteForce(input))

	input = []int{2, 3, 2}
	fmt.Println(maxAreaTwoPointer(input))
	fmt.Println(maxAreaBruteForce(input))
}
