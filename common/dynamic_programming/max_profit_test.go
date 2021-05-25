package dynamic_programming

import "testing"

func TestMaxProfit(t *testing.T) {
	nums := []int{7, 1, 5, 3, 6, 4}
	ans := maxProfit121(nums)
	if ans != 5 {
		t.Errorf("maxSubArray(nums)= %d; want 5", ans)
	}
}
