package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else if head.Val == head.Next.Val {
		head = deleteDuplicates(head.Next)
		return head
	} else {
		head.Next = deleteDuplicates(head.Next)
		return head
	}
}
