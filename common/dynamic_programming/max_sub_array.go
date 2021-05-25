package dynamic_programming

import "v0/common/tools"

/*
53. 最大子序和
https://leetcode-cn.com/problems/maximum-subarray/

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
*/

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = tools.Max(dp[i-1], 0) + nums[i]
	}

	ans := nums[0]
	for _, v := range dp {
		ans = tools.Max(ans, v)
	}
	return ans
}
