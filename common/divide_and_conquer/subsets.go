package divide_and_conquer

import "fmt"

/*
78. 子集
https://leetcode-cn.com/problems/subsets/

给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/

func subsets(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 != 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, set)
	}
	return ans
}

//依然不能理解啊
func subsets_1(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			fmt.Println(set)
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}
