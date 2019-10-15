package linked_list

import (
	"errors"
	"fmt"
)

/*
  单链表基本操作
*/

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{Val: v}
}

func (l *ListNode) Print(info string) {
	fmt.Println(info)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

// insert data after a node
func (l *ListNode) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}

	// find local p and insert v
	for l != nil {
		if l.Val == p.Val {
			newListNode := NewListNode(v)
			if l.Next == nil {
				l.Next = newListNode
			} else {
				newListNode.Next = l.Next
				l.Next = newListNode
			}
			return true
		}
		l = l.Next
	}

	return false
}

func (l *ListNode) InsertBefore(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}

	newListNode := NewListNode(v)
	// find local p and insert v
	for l != nil {
		if l.Next != nil && l.Next.Val == p.Val {
			newListNode.Next = l.Next
			l.Next = newListNode
			return true
		}
		l = l.Next
	}
	return false
}

func (l *ListNode) InsertToHead(v interface{}) error {
	return nil
}

// InsertToTail
func (l *ListNode) InsertToTail(v interface{}) error {
	for l.Next != nil {
		l = l.Next
	}
	l.Next = NewListNode(v)
	return nil
}

func (l *ListNode) FindByIndex(index uint) (interface{}, error) {
	for i := uint(1); i <= index; i++ {
		if l == nil {
			return nil, errors.New("out of index range")
		}

		if i == index {
			return l.Val, nil
		}
		l = l.Next
	}

	return l.Val, nil
}
