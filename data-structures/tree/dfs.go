package tree

import (
	"fmt"

	"github.com/gosoon/algorithms-golang/data-structures/stack"
)

func DFSPre(root *TreeNode, s *stack.ArrayStack) {
	if root == nil {
		return
	}
	s.Push(root)

	for !s.IsEmpty() {
		cur, _ := s.Pop()
		c := cur.(*TreeNode)
		fmt.Println(c.Val)

		if c.Right != nil {
			s.Push(c.Right)
		}
		if c.Left != nil {
			s.Push(c.Left)
		}
	}
}

func DFSMid(root *TreeNode, s *stack.ArrayStack) {
	if root == nil {
		return
	}

	cur := root
	for {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}
		if s.IsEmpty() {
			return
		}

		cur, _ := s.Pop()
		c := cur.(*TreeNode)
		fmt.Println(c.Val)
		cur = c.Right
	}
}
