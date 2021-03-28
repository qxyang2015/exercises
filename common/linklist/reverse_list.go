package linklist

/*
反转一个单链表。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preNode, curNode := (*ListNode)(nil), head
	for curNode != nil {
		curNode.Next, preNode, curNode = preNode, curNode, curNode.Next
	}
	return preNode
}
