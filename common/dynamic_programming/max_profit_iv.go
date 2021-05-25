package dynamic_programming

import "v0/common/tools"

func maxProfit5(k int, prices []int) int {
	n := len(prices)
	if n < 1 {

	}
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k)
	}
	dp[0][0][1] = -prices[0]
	for i := 1; i < n; i++ {
		for j := 1; j < k; j++ {
			dp[i][j][0] = tools.Max(dp[i-1][k][0], dp[i-1][k-1][1]+prices[i])
			dp[i][j][1] = tools.Max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][k-1][0]
}
