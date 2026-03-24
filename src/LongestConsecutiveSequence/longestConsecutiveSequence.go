package main

import (
	"fmt"
	"sort"
)

// Brute Force method for finding the longest consecutive sequence in an array of integers.
// Time complexity is O(n^3) and space complexity is O(1).
func longestConsecutiveBruteForce(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	longest := 0

	for i := 0; i < len(nums); i++ {
		length := 1
		num := nums[i]
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				if num+1 == nums[k] {
					length++
					num = num + 1
					break
				}
			}
		}
		if length > longest {
			longest = length
		}
	}
	return longest
}

// Brute Force method with Hash Map for finding the longest consecutive sequence in an array of integers.
// Time Complexity is O(n^2) and space complexity is O(n).
func longestConsecutiveBruteForceHashMap(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	storedHash := make(map[int]struct{})
	longest := 1

	for _, num := range nums {
		storedHash[num] = struct{}{}
	}

	for _, num := range nums {
		length, curr := 0, num
		for _, ok := storedHash[curr]; ok; _, ok = storedHash[curr] {
			length++
			curr++
		}

		if length > longest {
			longest = length
		}
	}
	return longest
}

// Sort the array of integers before finding longest consecutive sequence.
// Time Complexity is O(n log n) and space complexity is O(1).
func longestConsecutiveSorting(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	longest := 1
	current, length := nums[0], 1

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] == current {
			continue
		} else if nums[i] == current+1 {
			length++
			current = nums[i]

			if length > longest {
				longest = length
			}
		} else {
			current = nums[i]
			length = 1
		}
	}

	return longest
}

// Hash Map method where length of left and right is stored in order to find the longest Consecutive Sequence.
// Time Complexity is O(n) and Space Complexity is O(n)
func longestConsecutiveHashMap(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	storedHash := make(map[int]int)
	longest := 0

	for _, num := range nums {
		if storedHash[num] == 0 {
			left, right := storedHash[num-1], storedHash[num+1]
			length := left + right + 1

			storedHash[num] = length
			storedHash[num-left] = length
			storedHash[num+right] = length

			if length > longest {
				longest = length
			}
		}
	}
	return longest
}

// Hash Map Method where we only start where num-1 is not found in the hash map.
// Time Complexity is O(n) and Space Complexity is O(n).
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	storedHash := make(map[int]struct{})
	longest := 1

	for _, num := range nums {
		storedHash[num] = struct{}{}
	}

	for _, num := range nums {
		length := 1
		if _, found := storedHash[num-1]; !found {
			for {
				if _, exist := storedHash[num+length]; exist {
					length++
				} else {
					break
				}
			}
		}

		if length > longest {
			longest = length
		}
	}

	return longest
}

func main() {
	nums := []int{2, 20, 4, 10, 3, 4, 5}
	fmt.Println(longestConsecutive(nums))

	nums = []int{0, 3, 2, 5, 4, 6, -1, 1}
	fmt.Println(longestConsecutive(nums))

	nums = []int{0}
	fmt.Println(longestConsecutive(nums))

	nums = []int{2, 20, 4, 10, 3, 4, 5}
	fmt.Println(longestConsecutiveBruteForce(nums))

	nums = []int{0, 3, 2, 5, 4, 6, -1, 1}
	fmt.Println(longestConsecutiveHashMap(nums))

	nums = []int{0}
	fmt.Println(longestConsecutiveBruteForce(nums))
}
