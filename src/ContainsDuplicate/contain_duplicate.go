package main

import "fmt"

func hasDuplicate(nums []int) bool {
	check := make(map[int]bool)
	for _, num := range nums {
		if check[num] {
			return true
		}
		check[num] = true
	}
	return false
}

func main() {
	got := hasDuplicate([]int{1, 2, 3, 4, 3})
	fmt.Println(got)
}
