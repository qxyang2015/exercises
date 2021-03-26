package linked_list

/*
1. 合并两个有序链表
https://leetcode-cn.com/problems/merge-two-sorted-lists/

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	target := &ListNode{}
	targetPtr := target
	ptrl1, ptrl2 := l1, l2
	for ptrl1 != nil && ptrl2 != nil {
		if ptrl1.Val < ptrl2.Val {
			targetPtr.Next = ptrl1
			ptrl1 = ptrl1.Next
		} else {
			targetPtr.Next = ptrl2
			ptrl2 = ptrl2.Next
		}
		targetPtr = targetPtr.Next
	}
	if ptrl1 != nil {
		targetPtr.Next = ptrl1
	}
	if ptrl2 != nil {
		targetPtr.Next = ptrl2
	}
	return target.Next
}
