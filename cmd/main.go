package main

import (
	"fmt"
)

func maxProfit(n int, prices []int) int {
	if n == 0 || len(prices) == 0 {
		return 0
	}

	// Initialize the buy and sell arrays
	buy := make([]int, n+1)
	sell := make([]int, n+1)
	for i := 0; i <= n; i++ {
		buy[i] = -prices[0]
	}

	// Update the buy and sell arrays for each price
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= n; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}

	// Note: elements in buy and sell array represent the maximum profit value by
	// selling, buying or doing nothing and these profit values are updated each day
	// depending on the price.

	// Return the maximum achievable profit
	return sell[n]
}

// max returns the maximum value between a and b
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{6, 4, 7, 2, 3}
	n := 2
	fmt.Println(maxProfit(n, prices))
}
