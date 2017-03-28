package main

import (
	"fmt"
)

type ListNode struct {
	value int
	next  *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		value: val,
	}
}

func reverse(head *ListNode) *ListNode {
	newList := NewListNode(head.value)
	for head.next != nil {
		tmp := newList
		newList = head.next
		head.next = head.next.next
		newList.next = tmp
	}
	return newList
}

func Insert(h, l *ListNode) bool {
	if h.next == nil {
		h.next = l
		return true
	}

	n := h
	for n.next != nil {

		if n.next.next == nil {
			n.next.next = l
			return true
		}

		n = n.next
	}
	return true
}

func main() {
	head := NewListNode(0)
	var node *ListNode
	for i := 1; i < 10; i++ {
		node = NewListNode(i)
		Insert(head, node)
	}
	n := head
	for n.next != nil {
		fmt.Println(n.value)
		n = n.next
	}
	fmt.Println("++++++++++++++++++")
	newList := reverse(head)
	next := newList.next
	fmt.Println(newList.value)
	for {
		if next != nil {
			fmt.Println(next.value)
			next = next.next
		} else {
			break
		}
	}
}
