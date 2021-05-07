package dynamic_programming

/*
152. 乘积最大子数组
https://leetcode-cn.com/problems/maximum-product-subarray/

给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
*/

func maxProduct(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))
	dpMax[0], dpMin[0] = nums[0], nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		dpMax[i] = Max(dpMax[i-1]*nums[i], Max(dpMin[i-1]*nums[i], nums[i]))
		dpMin[i] = Min(dpMin[i-1]*nums[i], Min(dpMax[i-1]*nums[i], nums[i]))
		if ans < dpMax[i] {
			ans = Max(ans, dpMax[i])
		}
	}
	return ans
}

func Min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
