package leetCode

import (
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func threeSum(nums []int) [][]int {
	var targetNums [][]int
	if len(nums) < 3 {
		return targetNums
	}
	sort.Ints(nums)
	for i, a := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		k := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && a+nums[j]+nums[k] > 0 {
				k--
			}
			if j >= k {
				break
			}
			if a+nums[j]+nums[k] == 0 {
				targetNums = append(targetNums, []int{a, nums[j], nums[k]})
			}
		}
	}
	return targetNums
}
