package leetCode

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func moveZeroes(nums []int) {
	zeroPostion := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[zeroPostion], nums[i] = nums[i], nums[zeroPostion]
			zeroPostion++
		}
	}
}
