package main

import "fmt"

// Optimized approach: Uses prefix and suffix products to calculate the result for each element without using division.
// Time Complexity: O(n) and Space Complexity: O(n) for the output array and prefix/suffix arrays.
func productExceptSelfPrefixSuffixUnoptimised(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	prefix := make([]int, n)
	suffix := make([]int, n)

	prefix[0] = 1
	suffix[n-1] = 1

	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		result[i] = prefix[i] * suffix[i]
	}

	return result
}

// Optimized approach: Uses prefix and suffix products to calculate the result for
// each element without using division.
// Time Complexity: O(n) and Space Complexity: O(1) for the prefix and suffix
// variables, O(n) for the output array.
func productExceptSelfPrefixSuffixOptimised(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	prefix, suffix := 1, 1

	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

// Solved in O(n) time with O(1) extra space along with O(n) space for the output array.
// Division approach: Calculate the total product of all elements and then
// divide by each element to get the result for that position.
func productExceptSelfDivide(nums []int) []int {
	n := len(nums)
	totalProduct := 1
	zeroCount := 0
	result := make([]int, n)

	for _, num := range nums {
		if num == 0 {
			zeroCount++
			continue
		}
		totalProduct *= num
	}

	if zeroCount > 1 {
		return result
	}

	for i := 0; i < n; i++ {
		if zeroCount > 0 {
			if nums[i] == 0 {
				result[i] = totalProduct
			} else {
				result[i] = 0
			}
		} else {
			result[i] = totalProduct / nums[i]
		}
	}

	return result
}

// Unoptimized approach: Uses a nested loop to calculate the product for each element by
// multiplying all other elements.
// Time Complexity: O(n^2) and Space Complexity: O(n) for the output array.
func productExceptSelfUnoptimized(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		product := 1
		for j := 0; j < n; j++ {
			if i != j {
				product *= nums[j]
			}
		}
		result[i] = product
	}

	return result
}

func main() {

	//input := []int{1, 2, 4, 6}
	input := []int{-1, 0, 1, 2, 3}
	result := productExceptSelfPrefixSuffixUnoptimised(input)
	fmt.Println(result)
	result = productExceptSelfPrefixSuffixOptimised(input)
	fmt.Println(result)
	result = productExceptSelfDivide(input)
	fmt.Println(result)
	result = productExceptSelfUnoptimized(input)
	fmt.Println(result)
}
