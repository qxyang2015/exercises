package array

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	mArea := 0
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * MinInt(height[left], height[right])
		if area > mArea {
			mArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return mArea
}

func MinInt(v1, v2 int) int {
	if v2 > v1 {
		return v2
	}
	return v1
}
