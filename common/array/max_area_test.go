package array

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	mArea := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if mArea != 49 {
		t.Errorf("MaxArea([]int{1,8,6,2,5,4,8,3,7}) = %d; want 49", mArea)
	}
}
