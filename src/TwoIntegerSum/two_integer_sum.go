package main

import "fmt"

// Brute force ignores the ordering and simply checks every possible pair.
// For each index i, we look at every index j > i and check whether their sum equals the target.
// Time Complexity: O(n^2), Space Complexity: O(1)
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return nil
}

// Because the array is already sorted, we don’t need to check every pair.
// For each number at index i, we know exactly what value we need to find:
// target - numbers[i].
// Since the array is sorted, we can efficiently search for this value using
// binary search instead of scanning linearly.
// Time Complexity: O(n log n), Space Complexity: O(1)
func twoSumBinarySearch(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		left, right := i+1, len(nums)-1

		for left <= right {
			mid := left + (right-left)/2

			if nums[mid] == temp {
				return []int{i + 1, mid + 1}
			} else if nums[mid] < temp {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return nil
}

// As we scan through the list, we compute the needed complement for each number.
// If that complement has already been seen earlier (stored in the hash map),
// then we have found the required pair.
// Time Complexity: O(n), Space Complexity: O(n)
func twoSumHashMap(nums []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]

		if val, ok := mp[temp]; ok {
			return []int{val, i + 1}
		}

		mp[nums[i]] = i + 1
	}
	return nil
}

// Because the array is sorted, we can use two pointers to adjust the sum efficiently.
// If the current sum is too big, moving the right pointer left makes the sum smaller.
// If the sum is too small, moving the left pointer right makes the sum larger.
// This lets us quickly close in on the target without checking every pair.
// Time Complexity: O(n), Space Complexity: O(1)
func twoSumTwoPointer(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func main() {
	input := []int{1, 2, 3}
	fmt.Println(twoSumBruteForce(input, 3))
	fmt.Println(twoSumBinarySearch(input, 3))
	fmt.Println(twoSumHashMap(input, 3))
	fmt.Println(twoSumTwoPointer(input, 3))
}
