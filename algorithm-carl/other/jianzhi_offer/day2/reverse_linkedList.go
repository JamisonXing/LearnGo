package day2

type ListNode01 struct {
	Val  int
	Next *ListNode01
}

func reverseList(head *ListNode01) *ListNode01 {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
