package tree

import (
	"fmt"

	"github.com/gosoon/algorithms-golang/data-structures/queue"
)

func BFS(root *TreeNode) {
	queue := queue.NewArrayQueue(100)
	visit(root, queue)
	fmt.Println(queue)
}

func visit(root *TreeNode, queue *queue.ArrayQueue) {
	if root == nil {
		return
	}
	queue.Enqueue(root.Val)
	if root.Left != nil {
		visit(root.Left, queue)
	}

	if root.Right != nil {
		visit(root.Right, queue)
	}
}
