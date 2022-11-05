package main

import "fmt"

type ListNodeDemo struct {
	Val  int
	Next *ListNodeDemo
}

//https://leetcode.cn/problems/remove-linked-list-elements/
/*
给你一个链表的头节点 head 和一个整数 val ，
请你删除链表中所有满足 Node.val == val 的节点，
并返回 新的头节点 。
*/

// 设置一个虚拟节点，https://www.bilibili.com/video/BV18B4y1s7R9/?vd_source=1761deea5b8c69ccd1cd089a400b2c54
func removeElements(head *ListNodeDemo, val int) *ListNodeDemo {
	dummyHead := &ListNodeDemo{
		Next: head,
	}
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

func main() {
	b := ListNodeDemo{
		Val:  3,
		Next: nil,
	}
	a := ListNodeDemo{
		Val:  2,
		Next: &b,
	}
	head := ListNodeDemo{
		Val:  1,
		Next: &a,
	}
	res := removeElements(&head, 2)
	fmt.Println(res.Val, "->", res.Next.Val)
}
