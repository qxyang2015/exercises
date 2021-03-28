package linklist

/*
给定一个链表，判断链表中是否有环。
如果链表中存在环，则返回 true 。 否则，返回 false 。
*/

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return false
		}
		slow, fast = slow.Next, fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
