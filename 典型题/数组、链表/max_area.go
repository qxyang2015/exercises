package leetCode

func MaxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	mArea := 0
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * MaxInt(height[left], height[right])
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

func MaxInt(v1, v2 int) int {
	if v2 > v1 {
		return v2
	}
	return v1
}
