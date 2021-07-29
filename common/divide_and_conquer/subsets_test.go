package divide_and_conquer

import "testing"

func TestSubSets(t *testing.T) {
	mArea := subsets([]int{1, 2, 3})
	if len(mArea) != 5 {
		t.Errorf("subsets([]int{1,8,6,2,5,4,8,3,7}) = %d; want 49", mArea)
	}
}
