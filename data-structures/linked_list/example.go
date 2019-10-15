package linked_list

/*
- 单链表反转
- 链表中环的检测
- 两个有序的链表合并
- 删除链表倒数第 n 个结点
- 求链表的中间节点
*/

// reverse linked list
// leetcode 206
func reverseList(head *ListNode) *ListNode {
	var newListNode *ListNode
	for head != nil {
		t := head.Next
		head.Next = newListNode
		newListNode = head
		head = t
	}

	return newListNode
}

// check has cycle in linked list
// leetcode 141
func hasCycle(head *ListNode) bool {
	if head != nil && head.Next != nil {
		first := head
		second := head.Next.Next

		for first != nil && second != nil {
			if first.Val == second.Val {
				return true
			}
			first = first.Next
			if second.Next == nil {
				return false
			}
			second = second.Next.Next
		}
	}
	return false
}

// merge two ordered linked list
// leetcode 21

/*
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val.(int) < l2.Val.(int) {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	pre := l3
	for l1 != nil && l2 != nil {
		if l1.Val.(int) > l2.Val.(int) {
			l3.Next = l2
			l2 = l2.Next
		} else {
			l3.Next = l1
			l1 = l1.Next
		}
		l3 = l3.Next
	}
	if l1 == nil {
		l3.Next = l2
	} else {
		l3.Next = l1
	}

	return pre.Next
}
