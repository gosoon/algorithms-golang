package tree

import (
	"fmt"
	"testing"

	"github.com/gosoon/algorithms-golang/data-structures/queue"
)

func TestLowestCommonAncestor_235(t *testing.T) {
	nums := []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5}
	q := queue.NewArrayQueue(len(nums))
	for _, v := range nums {
		q.Enqueue(v)
	}

	root := &TreeNode{}
	InitTree(root, q)
	fmt.Println(lowestCommonAncestor_235(root, &TreeNode{Val: 2}, &TreeNode{Val: 8}))
}
