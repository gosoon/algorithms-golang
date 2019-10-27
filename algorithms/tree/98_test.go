package tree

import (
	"fmt"
	"testing"

	"github.com/gosoon/algorithms-golang/data-structures/queue"
)

func TestIsValidBST(t *testing.T) {
	nums := []int{10, 5, 15, -1, -1, 6, 20}
	q := queue.NewArrayQueue(len(nums))
	for _, v := range nums {
		q.Enqueue(v)
	}

	root := &TreeNode{}
	InitTree(root, q)
	//PrePrint(root)
	fmt.Println(isValidBST(root))
}
