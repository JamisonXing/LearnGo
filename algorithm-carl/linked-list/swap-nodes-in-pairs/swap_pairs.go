package swap_nodes_in_pairs

//https://programmercarl.com/0024.%E4%B8%A4%E4%B8%A4%E4%BA%A4%E6%8D%A2%E9%93%BE%E8%A1%A8%E4%B8%AD%E7%9A%84%E8%8A%82%E7%82%B9.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{
		Next: head,
	}
	cur := dummyNode

	for head != nil && head.Next != nil { //cur.next.next不能写前面，可能会造成空指针
		temp1 := head.Next.Next

		cur.Next = head.Next
		cur.Next.Next = head
		head.Next = temp1

		//移动cur至一对节点前面
		cur = cur.Next.Next
		head = head.Next
	}
	return dummyNode.Next
}
