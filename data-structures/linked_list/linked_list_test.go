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

}
