package main

import "fmt"

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
