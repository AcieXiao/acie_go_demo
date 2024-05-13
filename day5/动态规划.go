package main

import "fmt"

func maxProfile(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func main() {
	list := []int{65, 456, 6, 54, 65, 646, 4656, 4, 66, 123, 1654, 1651, 53, 1651, 651, 1, 321, 3}
	fmt.Println(maxProfile(list))
}
