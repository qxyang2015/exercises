package dynamic_programming

import "testing"

func TestCoinChange(t *testing.T) {
	nums := []int{1, 2, 5}
	ans := coinChange(nums, 11)
	if ans != 3 {
		t.Errorf(" coinChange(nums,11)= %d; want 3", ans)
	}
}
