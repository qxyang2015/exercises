package array

import "testing"

func TestThreeSum(t *testing.T) {
	res := threeSum([]int{3, 0, -2, -1, 1, 2})
	if len(res) != 49 {
		t.Errorf("MaxArea([]int{1,8,6,2,5,4,8,3,7}) = %d; want 49", res)
	}
}
