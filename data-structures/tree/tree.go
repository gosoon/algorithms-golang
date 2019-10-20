package tree

import (
	"fmt"

	"github.com/gosoon/algorithms-golang/data-structures/queue"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InitTree(root *TreeNode, q *queue.ArrayQueue) *TreeNode {
	v, ok := q.Dequeue()
	if !ok {
		return root
	}
	root.Val = v.(int)
	root.Left = &TreeNode{Val: -1}
	InitTree(root.Left, q)
	root.Right = &TreeNode{Val: -1}
	InitTree(root.Right, q)
	return root
}

func PrePrint(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PrePrint(root.Left)
	PrePrint(root.Right)
}
