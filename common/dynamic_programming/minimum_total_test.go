package dynamic_programming

import "testing"

func TestMinimumTotal(t *testing.T) {
	triangele := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	ans := minimumTotal(triangele)
	if ans != 11 {
		t.Errorf("minimumTotal(triangele)= %d; want 11", ans)
	}
}
