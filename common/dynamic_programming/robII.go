package dynamic_programming

import "v0/common/tools"

/*
213. 打家劫舍 II
https://leetcode-cn.com/problems/house-robber-ii/

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
*/

func robII(nums []int) int {
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
	return tools.Max(robMethod(nums[:n-1]), robMethod(nums[1:]))
}

func robMethod(nums []int) int {
	p, q := nums[0], tools.Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		p, q = q, tools.Max(q, p+nums[i])
	}
	return q
}
