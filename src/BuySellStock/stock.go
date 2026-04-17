package main

import "fmt"

// As we scan through the prices, we keep track of two things
// 1.The lowest price so far → this is the best day to buy.
// 2.The best profit so far → selling today minus the lowest buy price seen earlier.
// At each price, we imagine selling on that day.
// The profit would be:
// current price – lowest price seen so far
// We then update the maximum profit,
// and the lowest price if we find a cheaper one.
// This way, we make the optimal buy–sell decision in one simple pass.
//
// Time Complexity: O(n) since we traverse the list of prices once.
// Space Complexity: O(1)
func maxProfitDynamicProgramming(prices []int) int {
	lowest := prices[0]
	res := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < lowest {
			lowest = prices[i]
		} else {
			profit := prices[i] - lowest
			if profit > res {
				res = profit
			}
		}
	}
	return res
}

// We want to buy at a low price and sell at a higher price that comes after it.
// Using two pointers helps us track this efficiently:
//
// l is the buy day (looking for the lowest price)
// r is the sell day (looking for a higher price)
// If the price at r is higher than at l, we can make a profit — so we update the maximum.
// If the price at r is lower, then r becomes the new l because a cheaper buying price is always better.
//
// By moving the pointers this way, we scan the list once and always keep the best buying opportunity.
// Time Complexity: O(n), Space Complexity: O(1)
func maxProfitTwoPointer(prices []int) int {
	l, r := 0, 1
	maxP := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			if profit > maxP {
				maxP = profit
			}
		} else {
			l = r
		}
		r++
	}
	return maxP
}

// This function is irrelevant since we can only buy and sell stock one time and looking for max profit
// with multiple trades
func maxProfitMutiBuySell(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	res := 0

	for i := 1; i < len(prices)-1; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if prices[i] > prices[i+1] {
			res += prices[i] - minPrice
			minPrice = prices[i+1]
		}
	}

	if prices[len(prices)-1] > minPrice {
		res += prices[len(prices)-1] - minPrice
	}
	return res
}

// The brute-force approach checks every possible buy–sell pair.
// For each day, we pretend to buy the stock, and then we look at all the future days to see
// what the best-selling price would be.
// Among all these profits, we keep the highest one.
// Time Complexity: O(n^2) because we have two nested loops, Space Complexity: O(1)
func maxProfitBruteForce(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0

	for i := 0; i < len(prices); i++ {
		buy := prices[i]

		for j := i + 1; j < len(prices); j++ {
			sell := prices[j]
			res = max(res, sell-buy)
		}
	}
	return res
}

func main() {
	input := []int{10, 1, 5, 6, 7, 1}
	fmt.Println(maxProfitDynamicProgramming(input))

	input = []int{10, 8, 7, 5, 2}
	fmt.Println(maxProfitDynamicProgramming(input))

	input = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfitDynamicProgramming(input))

	input = []int{5, 10}
	fmt.Println(maxProfitBruteForce(input))
	fmt.Println(maxProfitTwoPointer(input))
	fmt.Println(maxProfitDynamicProgramming(input))
}
