package linked_list

import "fmt"

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

	newListNode := NewListNode(v)
	if p.Next == nil {
		p.Next = newListNode
	} else {
		newListNode.Next = p.Next
		p.Next = newListNode
	}
	return true
}

func (l *ListNode) InsertBefore(p *ListNode, v interface{}) bool {
	return true

}

func (l *ListNode) InsertToHead() error {
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

	return 1, nil
}
