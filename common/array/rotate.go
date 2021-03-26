package array

/*
189. 旋转数组
https://leetcode-cn.com/problems/rotate-array/
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
*/

func rotate(nums []int, k int) {
	k %= len(nums)
	Reverse(nums)
	Reverse(nums[:k])
	Reverse(nums[k:])
}

func Reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
