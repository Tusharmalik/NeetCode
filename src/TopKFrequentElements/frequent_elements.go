package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	final := make([]int, 0, k)

	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	for i := 0; i < k; i++ {
		maxNum := 0
		maxCount := 0

		for num, c := range count {
			if c > maxCount {
				maxNum = num
				maxCount = c
			}
		}

		final = append(final, maxNum)
		delete(count, maxNum)
	}

	return final
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
