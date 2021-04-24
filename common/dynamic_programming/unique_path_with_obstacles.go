package dynamic_programming

/*
63. 不同路径 II
https://leetcode-cn.com/problems/unique-paths-ii/

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
*/

//滚动数组的方式，空间复杂度更低
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, col)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if obstacleGrid[r][c] == 1 {
				dp[c] = 0
				continue
			}
			if c > 0 && obstacleGrid[r][c-1] == 0 {
				dp[c] += dp[c-1]
			}
		}
	}
	return dp[col-1]
}

func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	if len(obstacleGrid) <= 0 {
		return 0
	}
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, len(obstacleGrid))
	for r := 0; r < row; r++ {
		dp[r] = make([]int, len(obstacleGrid[r]))
		if r == 0 {
			dp[0][0] = 1
			continue
		}
		if obstacleGrid[r][0] == 0 && dp[r-1][0] == 0 {
			dp[r][0] = 1
		}
	}
	for c := 1; c < col; c++ {
		if obstacleGrid[0][c] == 0 && dp[0][c-1] == 0 {
			dp[0][c] = 1
		}
	}
	for r := 1; r < row; r++ {
		for c := 1; c < col; c++ {
			if obstacleGrid[r][c] == 0 && (dp[r][c-1] != 0 || dp[r-1][c] != 0) {
				dp[r][c] = dp[r][c-1] + dp[r-1][c]
			}
		}
	}
	return dp[row-1][col-1]
}
