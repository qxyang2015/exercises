package linked_list

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

func swapPairs(head *ListNode) *ListNode {
	preHead := &ListNode{0, head}
	tmpNode := preHead
	for tmpNode.Next != nil && tmpNode.Next.Next != nil {
		node1, node2 := tmpNode.Next, tmpNode.Next.Next
		tmpNode.Next = node2
		node1.Next, node2.Next = node2.Next, node1
		tmpNode = node1
	}
	return preHead.Next
}
