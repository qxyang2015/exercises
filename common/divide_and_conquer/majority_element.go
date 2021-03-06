package divide_and_conquer

/*
169. 多数元素
https://leetcode-cn.com/problems/majority-element/

给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/

func majorityElement(nums []int) int {
	major, cnt := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if cnt <= 0 {
			major = nums[i]
		}
		if major == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return major
}
