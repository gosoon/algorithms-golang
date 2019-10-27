package tree

//  https://leetcode-cn.com/problems/validate-binary-search-tree/submissions/

func isValidBST(root *TreeNode) bool {
	min, max := -1<<63, 1<<63-1
	return recur(root, min, max)
}

func recur(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	return root.Val > min && root.Val < max &&
		recur(root.Left, min, root.Val) &&
		recur(root.Right, root.Val, max)

}
