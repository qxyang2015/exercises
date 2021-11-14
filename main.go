package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	nums := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(nums)
	fmt.Println(minimumTotal(nums))

	fmt.Println("done!")
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	dp := make([][]int, 2)
	dp[0], dp[1] = make([]int, n), make([]int, n)
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[1][j] = dp[0][j] + triangle[i][j]
				continue
			}
			if i == j {
				dp[1][j] = dp[0][j-1] + triangle[i][j]
				continue
			}
			dp[1][j] = maxInt(dp[0][j-1], dp[0][j]) + triangle[i][j]
		}
		dp[0], dp[1] = dp[1], dp[0]
	}
	minPath := dp[0][0]
	for i := 1; i < n; i++ {
		minPath = maxInt(minPath, dp[0][i])
	}
	return minPath
}

func maxInt(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
