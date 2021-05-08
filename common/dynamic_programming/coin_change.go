package dynamic_programming

import "math"

/*
322. 零钱兑换
https://leetcode-cn.com/problems/coin-change/

给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
你可以认为每种硬币的数量是无限的。
*/

//动态规划
func coinChange1(coins []int, amount int) int {
	if len(coins) < 1 {
		return -1
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = Min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

//BFS解法
func coinChange(coins []int, amount int) int {
	if len(coins) < 1 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	ans := 0
	p := make([]int, len(coins))
	for i := 0; i < len(coins); i++ {
		p[i] = amount
	}
	for len(p) > 0 {
		var q []int
		for i := 0; i < len(p); i++ {
			for j := 0; j < len(coins); j++ {
				v := p[i] - coins[j]
				if v == 0 {
					return ans + 1
				} else if v > 0 {
					q = append(q, v)
				} else {
					continue
				}
			}
		}
		p = q
		ans++
	}
	return -1
}
