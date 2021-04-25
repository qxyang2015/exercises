package dynamic_programming

import "math"

func minimumTotal(triangle [][]int) int {
	if len(triangle) <= 0 {
		return 0
	}
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = MinInt(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	ans := math.MaxInt32
	for i := 0; i < len(triangle[len(triangle)-1]); i++ {
		ans = MinInt(ans, dp[len(triangle)-1][i])
	}
	return ans
}

func MinInt(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
