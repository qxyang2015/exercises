package binary_search

/*
367. 有效的完全平方数
https://leetcode-cn.com/problems/valid-perfect-square/

给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
进阶：不要 使用任何内置的库函数，如sqrt 。
*/

func isPerfectSquare(num int) bool {
	if num <= 1 {
		return true
	}
	x := num
	for x*x > num {
		x = (x + num/x) / 2
	}
	return x*x == num
}
