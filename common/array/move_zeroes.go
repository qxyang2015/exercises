package array

/*
283. 移动零
https://leetcode-cn.com/problems/move-zeroes/
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
*/

func moveZeroes(nums []int) {
	zeroPostion := 0
	for idx, v := range nums {
		if v != 0 {
			nums[zeroPostion], nums[idx] = nums[idx], nums[zeroPostion]
			zeroPostion++
		}
	}
}
