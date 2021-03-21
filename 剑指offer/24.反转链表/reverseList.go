package ReverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preNode, curNode := (*ListNode)(nil), head
	for curNode != nil {
		curNode.Next, preNode, curNode = preNode, curNode, curNode.Next
	}
	return preNode
}
