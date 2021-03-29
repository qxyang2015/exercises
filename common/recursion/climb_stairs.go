package recursion

/*
70. 爬楼梯
https://leetcode-cn.com/problems/climbing-stairs/
*/

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}
