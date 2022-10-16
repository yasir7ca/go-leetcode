package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.com/problems/coin-change/
*/
func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
}

func coinChange(coins []int, amount int) int {
	dp := make([]float64, amount+1) // dp[k] denotes minimum number of coins to make amount k from given coin denominations
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.Inf(+1)
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if math.IsInf(dp[amount], +1) {
		return -1
	}
	return int(dp[amount])
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}
