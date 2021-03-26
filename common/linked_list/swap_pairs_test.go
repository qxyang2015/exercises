package linked_list

import "testing"

func TestSwapPairs(t *testing.T) {
	p1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	p2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	p1.Next = p2
	p3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	p2.Next = p3
	p4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	p3.Next = p4

	v := swapPairs(p1)
	if v.Val == 2 {
		t.Logf("查找环正确")
	} else {
		t.Errorf("detectCycle([1,2,3],2) = %v; want 2", v)
	}
}
