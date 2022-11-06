package design_linked_list

import "math"

//https://leetcode.cn/problems/design-linked-list/

//单链表

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&ListNode{}, 0}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	cur := l.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (l *MyLinkedList) AddAtIndex(index, val int) {
	if index > l.size {
		return
	}
	index = int(math.Max(float64(index), 0))
	l.size++
	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	newNode := &ListNode{val, pred.Next}
	pred.Next = newNode
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > l.size {
		return
	}
	l.size--

	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	pred.Next = pred.Next.Next
}
