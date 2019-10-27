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

func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil || node == nil {
		return
	}

	if tree.Val == 0 {
		tree.Val = node.Val
		return
	}

	if node.Val > tree.Val {
		if tree.Right == nil {
			tree.Right = &TreeNode{}
		}
		InsertNodeToTree(tree.Right, node)
	}
	if node.Val < tree.Val {
		if tree.Left == nil {
			tree.Left = &TreeNode{}
		}
		InsertNodeToTree(tree.Left, node)
	}
}

func InitTree2(Vals []int) (root *TreeNode) {
	rootNode := &TreeNode{Val: Vals[0]}
	for _, val := range Vals {
		var node *TreeNode
		if val != -1 {
			node = &TreeNode{Val: val}
		}
		InsertNodeToTree(rootNode, node)
	}
	return rootNode
}

func MidPrint(root *TreeNode) {
	if root == nil {
		return
	}
	//MidPrint(root.Left)
	fmt.Printf("%v  ", root.Val)
	MidPrint(root.Right)
}
