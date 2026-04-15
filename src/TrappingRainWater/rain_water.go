package main

import "fmt"

// For each position, the water trapped above it depends on the tallest bar to its left and the tallest
// bar to its right. The water trapped at that position is the minimum of these two heights minus
// the height of the bar at that position. If this value is negative, it means no water is trapped
// at that position.
// min(leftMax, rightMax) - height[i]
// Time Complexity: O(n^2) - For each element, we are traversing the left and right sides to find the maximum height.
// Space Complexity: O(1) - We are using only a constant amount of extra space.
func trapBruteForce(height []int) int {
	var result int
	if len(height) == 0 {
		return result
	}

	n := len(height)
	for i := 0; i < n; i++ {
		leftMax, rightMax := height[i], height[i]

		for j := 0; j < i; j++ {
			if height[j] > leftMax {
				leftMax = height[j]
			}
		}

		for j := i + 1; j < n; j++ {
			if height[j] > rightMax {
				rightMax = height[j]
			}
		}

		result = result + min(leftMax, rightMax) - height[i]
	}
	return result
}

// Instead of recomputing the tallest bar to the left and right for every index, we can precompute these values once.
// This removes the repeated work from the brute-force approach and makes the solution
// more efficient and easier to understand.
// Time Complexity: O(n) - We traverse the height array three times: once to fill the leftMax array,
// once to fill the rightMax array, and once to calculate the trapped water.
// Space Complexity: O(n) - We are using two additional arrays of size n to store the leftMax and rightMax values.
func trapPrefixSuffix(height []int) int {
	var result int
	if len(height) == 0 {
		return result
	}

	n := len(height)
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i := range height {
		result = result + min(leftMax[i], rightMax[i]) - height[i]
	}

	return result
}

// This approach uses two pointers to traverse the height array from both ends towards the center.
// We maintain two variables, leftMax and rightMax, to keep track of the maximum height encountered
// so far from the left and right sides, respectively. At each step, we compare leftMax and rightMax.
// If leftMax is smaller, it means that the water trapped at the current left pointer depends on leftMax,
// and we can safely move the left pointer inward. If rightMax is smaller or equal, it means that the
// water trapped at the current right pointer depends on rightMax, and we can safely move the
// right pointer inward.
// Time Complexity: O(n) - We traverse the height array once with two pointers.
// Space Complexity: O(1) - We are using only a constant amount of extra space for the pointers and max variables.
func trapTwoPointer(height []int) int {
	var result int
	if len(height) == 0 {
		return result
	}
	n := len(height)
	l, r := 0, n-1
	leftMax, rightMax := height[l], height[r]

	for l < r {
		if leftMax < rightMax {
			l++
			leftMax = max(leftMax, height[l])
			result += leftMax - height[l]
		} else {
			r--
			rightMax = max(rightMax, height[r])
			result += rightMax - height[r]
		}
	}
	return result
}

func main() {
	input := []int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1}
	fmt.Println(trapBruteForce(input))

	input = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trapBruteForce(input))

	input = []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trapBruteForce(input))

	input = []int{0}
	fmt.Println(trapBruteForce(input))

	input = []int{5, 4, 1, 2}
	fmt.Println(trapBruteForce(input))

	input = []int{9, 6, 8, 8, 5, 6, 3}
	fmt.Println(trapTwoPointer(input))

	input = []int{5}
	fmt.Println(trapTwoPointer(input))
}
