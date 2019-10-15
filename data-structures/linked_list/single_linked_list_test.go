package linked_list

import "testing"

func TestLinkedList(t *testing.T) {
	listNode := NewListNode(0)
	listNode.InsertToTail(1)
	listNode.InsertToTail(2)
	listNode.Print("~~0:")

	if !listNode.InsertAfter(NewListNode(1), 3) {
		t.Log("insert after nodelist failed")
	}
	listNode.Print("~~1:")

	if !listNode.InsertBefore(NewListNode(0), 3) {
		t.Log("insert before nodelist failed")
	}
	listNode.Print("~~2:")

	if _, err := listNode.FindByIndex(3); err != nil {
		t.Log(err)
	}
	listNode.Print("~~3:")
}
