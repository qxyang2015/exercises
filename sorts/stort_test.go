package Sort

import (
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{2, 1, 3, 4, 5}
	ans := QuickSort(arr)
	if len(ans) != 5 {
		t.Errorf("BubbleSort(arr) = %d; want 5", ans)
	} else {
		t.Logf("BubbleSort(arr) = %d; want 5", ans)
	}
}
