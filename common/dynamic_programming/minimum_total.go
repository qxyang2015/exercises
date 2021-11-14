package dynamic_programming

import (
	"math"
	"v0/common/tools"
)

/*
120. 三角形最小路径和
https://leetcode-cn.com/problems/triangle/description/

给定一个三角形 triangle ，找出自顶向下的最小路径和。
每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
*/

//自己整理思路写出
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
func minimumTotal0(triangle [][]int) int {
	n := len(triangle)
	if n <= 0 {
		return 0
	}
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		curr := i % 2
		prev := 1 - curr
		dp[curr][0] = dp[prev][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[curr][j] = tools.Min(dp[prev][j], dp[prev][j-1]) + triangle[i][j]
		}
		dp[curr][i] = dp[prev][i-1] + triangle[i][i]
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = tools.Min(ans, dp[(n-1)%2][i])
	}
	return ans
}

func minimumTotal1(triangle [][]int) int {
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
			dp[i][j] = tools.Min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	ans := math.MaxInt32
	for i := 0; i < len(triangle[len(triangle)-1]); i++ {
		ans = tools.Min(ans, dp[len(triangle)-1][i])
	}
	return ans
}
