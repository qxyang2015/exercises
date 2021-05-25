package dynamic_programming

import "v0/common/tools"

/*
198. 打家劫舍
https://leetcode-cn.com/problems/house-robber/

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
*/
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return tools.Max(nums[0], nums[1])
	}
	p, q := nums[0], tools.Max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		p, q = q, tools.Max(q, p+nums[i])
	}
	return q
}

func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0], dp[1] = nums[0], tools.Max(nums[0], nums[1])
	ans := dp[1]
	for i := 2; i < n; i++ {
		dp[i] = tools.Max(dp[i-1], dp[i-2]+nums[i])
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func rob2(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}

	dp := make([][]int, 2)
	dp[0] = make([]int, n)
	dp[1] = make([]int, n)
	dp[1][0] = nums[0]
	for i := 1; i < n; i++ {
		dp[0][i] = tools.Max(dp[0][i-1], dp[1][i-1])
		dp[1][i] = dp[0][i-1] + nums[i]
	}
	return tools.Max(dp[0][n-1], dp[1][n-1])
}
