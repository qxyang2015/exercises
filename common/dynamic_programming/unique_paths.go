package dynamic_programming

/*
62. 不同路径
https://leetcode-cn.com/problems/unique-paths/

一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？
*/

//自己写的滚动数组法，空间复杂度更低
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

//滚动数组法，空间复杂度更低
func uniquePaths0(m int, n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if c > 0 {
				dp[c] += dp[c-1]
			}
		}
	}
	return dp[n-1]
}

func uniquePaths1(m int, n int) int {
	dp := make([][]int, m)
	for r := 0; r < m; r++ {
		dp[r] = make([]int, n)
		dp[r][0] = 1
	}
	for c := 0; c < n; c++ {
		dp[0][c] = 1
	}
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			dp[r][c] = dp[r][c-1] + dp[r-1][c]
		}
	}
	return dp[m-1][n-1]
}
