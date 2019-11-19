package tree

import (
	"testing"

	"github.com/gosoon/algorithms-golang/data-structures/queue"
	"github.com/gosoon/algorithms-golang/data-structures/stack"
)

func TestDFS(t *testing.T) {
	nums := []int{10, 5, 15, -1, -1, 6, 20}
	q := queue.NewArrayQueue(len(nums))
	for _, v := range nums {
		q.Enqueue(v)
	}
	root := &TreeNode{}
	root = InitTree(root, q)
	//PrePrint(root)

	s := stack.NewArrayStack(100)
	DFSPre(root, s)
	s.Print("msg:")
	DFSMid(root, s)
}
