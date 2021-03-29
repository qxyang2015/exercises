package divide_and_conquer

/*
50. Pow(x, n)
https://leetcode-cn.com/problems/powx-n/

实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。
*/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n > 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	val := quickMul(x, n/2)
	if n%2 == 0 {
		return val * val
	}
	return val * val * x
}
