package day2

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var stack []int
	var res []int
	if head == nil {
		return nil
	}

	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	for len(stack) > 0 {
		res = append(res, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return res
}
