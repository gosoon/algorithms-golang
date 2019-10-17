package linked_list

import "testing"

func TestReverseList(t *testing.T) {
	listNode := NewListNode(0)
	listNode.InsertToTail(1)
	listNode.InsertToTail(2)

	listNode.Print("first:")

	reverse := reverseList(listNode)
	reverse.Print("~~reverse:")
}

func TestHasCycle(t *testing.T) {
	listNode := NewListNode(0)
	listNode.InsertToTail(1)
	listNode.InsertToTail(2)

	m := listNode.Next.Next
	m.Next = listNode

	if hasCycle(listNode) {
		t.Log("has cycle")
	}
}

func TestMergeTwoLists(t *testing.T) {
	l1 := NewListNode(0)
	l1.InsertToTail(2)
	l1.InsertToTail(5)
	l1.Print("l1")

	l2 := NewListNode(1)
	l2.InsertToTail(2)
	l2.InsertToTail(4)
	l2.Print("l2")

	l3 := mergeTwoLists(l1, l2)
	l3.Print("l3")
}

func TestRemoveNthFromEnd(t *testing.T) {
	l1 := NewListNode(1)
	//l1.InsertToTail(2)
	//l1.InsertToTail(3)
	//l1.InsertToTail(4)
	//l1.InsertToTail(5)

	l1.Print("l1")

	l2 := removeNthFromEnd(l1, 1)
	l2.Print("l2")
}

func TestMiddleNode(t *testing.T) {
	l1 := NewListNode(1)
	l1.InsertToTail(2)
	l1.InsertToTail(3)
	l1.InsertToTail(4)
	l1.InsertToTail(5)
	l1.InsertToTail(6)

	l1.Print("l1")

	l2 := middleNode(l1)
	l2.Print("l2")
}
