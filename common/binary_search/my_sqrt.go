package binary_search

import "math"

/*
69. x 的平方根
https://leetcode-cn.com/problems/sqrtx/

实现int sqrt(int x)函数。
计算并返回x的平方根，其中x 是非负整数。
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
*/

//牛顿迭代法
//Xi+1=(Xi+C/Xi)
func mySqrt(x int) int {
	xi := x
	for xi*xi > x {
		xi = (xi + x/xi) / 2
	}
	return xi
}

func mySqrt0(x int) int {
	if x == 0 {
		return 0
	}
	xi, c := float64(x), float64(x)
	for {
		xi1 := 0.5 * (xi + c/xi)
		if math.Abs(xi1-xi) < 1e-7 {
			break
		}
		xi = xi1
	}
	return int(xi)
}

//二分法
func mySqrt1(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
