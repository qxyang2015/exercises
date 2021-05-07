package dynamic_programming

import "testing"

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ans := maxSubArray(nums)
	if ans != 6 {
		t.Errorf("maxSubArray(nums)= %d; want 6", ans)
	}
}
