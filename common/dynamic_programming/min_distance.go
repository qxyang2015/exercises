package dynamic_programming

import "v0/common/tools"

/*
72. 编辑距离
https://leetcode-cn.com/problems/edit-distance/

给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
*/
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1+1)
	for i := 0; i < n1+1; i++ {
		dp[i] = make([]int, n2+1)
		dp[i][0] = i
	}
	for j := 0; j < n2+1; j++ {
		dp[0][j] = j
	}
	for i := 1; i < n1+1; i++ {
		for j := 1; j < n2+1; j++ {
			if word1[i-1] != word2[j-1] {
				dp[i][j] = tools.Min(dp[i-1][j], tools.Min(dp[i][j-1], dp[i-1][j-1])) + 1
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[n1][n2]
}
