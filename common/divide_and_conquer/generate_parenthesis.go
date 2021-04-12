package divide_and_conquer

/*
22. 括号生成
https://leetcode-cn.com/problems/generate-parentheses/

数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
*/
func generateParenthesis(n int) (targets []string) {
	if n < 1 {
		return targets
	}
	var generateFunc func(left, right int, path string)
	generateFunc = func(left, right int, path string) {
		if left == n && right == n {
			targets = append(targets, path)
			return
		}
		if left < n {
			generateFunc(left+1, right, path+"(")
		}
		if right < left {
			generateFunc(left, right+1, path+")")
		}
	}
	generateFunc(0, 0, "")
	return
}
