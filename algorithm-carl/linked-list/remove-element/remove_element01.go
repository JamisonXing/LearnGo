package main

import "fmt"

type ListNodeDemo01 struct {
	Val  int
	Next *ListNodeDemo01
}

// RemoveElements01 不使用虚拟节点, 头节点的处理不一样
func RemoveElements01(head *ListNodeDemo01, val int) *ListNodeDemo01 {
	for head != nil && head.Val == val && head.Next != nil { //这里不是if
		head = head.Next
	}

	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {
	b := ListNodeDemo01{
		Val:  3,
		Next: nil,
	}
	a := ListNodeDemo01{
		Val:  2,
		Next: &b,
	}
	head := ListNodeDemo01{
		Val:  1,
		Next: &a,
	}
	res := RemoveElements01(&head, 2)
	fmt.Println(res.Val, "->", res.Next.Val)
}
