package array

/*
66. 加一
https://leetcode-cn.com/problems/plus-one/

给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
*/

func plusOne(digits []int) []int {
	val := digits[len(digits)-1] + 1
	q := val / 10
	digits[len(digits)-1] = val % 10
	for i := len(digits) - 2; i >= 0; i-- {
		val = (digits[i] + q)
		q = val / 10
		digits[i] = val % 10
	}
	if q > 0 {
		return append([]int{1}, digits...)
	}
	return digits
}
