package tree

import (
	"testing"

	"github.com/gosoon/algorithms-golang/data-structures/queue"
)

func TestTree(t *testing.T) {
	nums := []int{10, 5, 15, -1, -1, 6, 20}
	q := queue.NewArrayQueue(len(nums))
	for _, v := range nums {
		q.Enqueue(v)
	}
	root := &TreeNode{}
	root = InitTree(root, q)
	PrePrint(root)
}
