package main

import (
	"fmt"
	"sort"
)

// The brute-force approach simply tries every possible triplet.
// Since we check all combinations (i, j, k) with i < j < k, we are guaranteed
// to find all sets of three numbers that sum to zero.
// Sorting helps keep the triplets in order and makes it easier to avoid
// duplicates by storing them in a set.
// Time Complexity: O(n^3) - We have three nested loops, each running up to n times.
// Space Complexity: O(n) - We use a set to store unique triplets, which can take up to O(n) space in the worst case.
func threeSumBruteForce(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	res := map[[3]int]struct{}{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
				}
			}
		}
	}

	for k := range res {
		result = append(result, []int{k[0], k[1], k[2]})
	}
	return result
}

// After sorting the array, we can fix two numbers and look for the third number that completes the triplet.
// To do this efficiently, we use a hash map that stores how many times each number appears.
// As we pick the first and second numbers, we temporarily reduce their counts in the map so we don't reuse them.
// Then we check whether the needed third value still exists in the map.
// Sorting also helps us skip duplicates easily so we only add unique triplets.
// Time Complexity: O(n^2) - We have two nested loops, and the inner loop runs in O(1) time due to the hash map lookups.
// Space Complexity: O(n) - We use a hash map to store the counts of each number, which can take up to O(n) space in the worst case.
func threeSumHashMap(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	count := map[int]int{}
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	for i := 0; i < len(nums); i++ {
		count[nums[i]]--
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			count[nums[j]]--

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			target := -(nums[i] + nums[j])

			if count[target] > 0 {
				result = append(result, []int{nums[i], nums[j], target})
			}
		}
		for j := i + 1; j < len(nums); j++ {
			count[nums[j]]++
		}
	}
	return result
}

// After sorting the array, we can fix one number and then search for the
// other two using the two-pointer technique.
// Time Complexity: O(n^2) - We have two nested loops, and the inner loop runs in O(n) time.
// Space Complexity: O(1) - We only use a constant amount of extra space for the pointers and the result list (not counting the space used for the output).
func threeSumTwoPointer(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := range nums {
		if nums[i] >= 1 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			total := nums[i] + nums[left] + nums[right]

			if total == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else if total > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSumBruteForce(input))

	input = []int{0, 1, 1}
	fmt.Println(threeSumBruteForce(input))

	input = []int{0, 0, 0}
	fmt.Println(threeSumBruteForce(input))

	// input = []int{-1, 0, 1, 2, -1, -4}
	// fmt.Println(threeSumTwoPointer(input))

	input = []int{-1, 0, 1, 2, -1, -4, -1, 0, 1, 2, -1, -4, -1, 0, 1, 2, -1, -4}
	fmt.Println(threeSumTwoPointer(input))
	fmt.Println(threeSumHashMap(input))

}
